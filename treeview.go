package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lawrencegripper/azbrowse/tracing"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/style"
)

const (
	subscriptionType  = "subscription"
	resourceGroupType = "resourcegroup"
	resourceType      = "resource"
	deploymentType    = "deployment"
	actionType        = "action"
)

// TreeNode is an item in the ListWidget
type TreeNode struct {
	parentid         string
	id               string
	name             string
	expandURL        string
	itemType         string
	expandReturnType string
	deleteURL        string
	namespace        string
	armType          string
}

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y        int
	w, h        int
	items       []TreeNode
	contentView *ItemWidget
	statusView  *StatusbarWidget
	navStack    Stack
	title       string
	ctx         context.Context
	selected    int
}

// NewListWidget creates a new instance
func NewListWidget(ctx context.Context, x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget) *ListWidget {
	return &ListWidget{ctx: ctx, x: x, y: y, w: w, h: h, contentView: contentView, statusView: status}
}

// Layout draws the widget in the gocui view
func (w *ListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("listWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	if len(w.items) < 1 {
		return nil
	}

	linesUsedCount := 0
	allItems := make([]string, 0, len(w.items))

	for i, s := range w.items {
		var itemToShow string
		if i == w.selected {
			itemToShow = "▶  "
		} else {
			itemToShow = "   "
		}
		itemToShow = itemToShow + s.name + "\n"

		linesUsedCount = linesUsedCount + strings.Count(itemToShow, "\n")
		allItems = append(allItems, itemToShow)
	}

	linesPerItem := linesUsedCount / len(w.items)
	maxItemsCanShow := w.h / linesPerItem

	for i, item := range allItems {
		// Skip items above the selection to allow scrolling
		if w.selected > maxItemsCanShow && i < w.selected {
			continue
		}
		fmt.Fprintf(v, item)
	}

	return nil
}

// SetNodes allows others to set the list nodes
func (w *ListWidget) SetNodes(nodes []TreeNode) {
	w.selected = 0
	// Capture current view to navstack
	w.navStack.Push(&Page{
		Data:      w.contentView.GetContent(),
		Value:     w.items,
		Title:     w.title,
		Selection: w.selected,
	})

	w.items = nodes
}

// SetSubscriptions starts vaidation with the subs found
func (w *ListWidget) SetSubscriptions(subs armclient.SubResponse) {
	newList := []TreeNode{}
	for _, sub := range subs.Subs {
		newList = append(newList, TreeNode{
			name:             sub.DisplayName,
			id:               sub.ID,
			expandURL:        sub.ID + "/resourceGroups?api-version=2014-04-01",
			itemType:         subscriptionType,
			expandReturnType: resourceGroupType,
		})
	}

	w.title = "Subscriptions"
	w.items = newList
}

// GoBack takes the user back to preview view
func (w *ListWidget) GoBack() {
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		return
	}
	w.contentView.SetContent(previousPage.Data)
	w.selected = 0
	w.items = previousPage.Value
	w.title = previousPage.Title
	w.selected = previousPage.Selection
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {

	currentItem := w.items[w.selected]
	if currentItem.expandReturnType != "none" && currentItem.expandReturnType != actionType {
		// Capture current view to navstack
		w.navStack.Push(&Page{
			Data:      w.contentView.GetContent(),
			Value:     w.items,
			Title:     w.title,
			Selection: w.selected,
		})
	}
	span, ctx := tracing.StartSpanFromContext(w.ctx, "expandItem-"+currentItem.itemType, tracing.SetTag("item", currentItem))
	defer span.Finish()
	w.ctx = ctx

	method := "GET"
	if currentItem.expandReturnType == actionType {
		method = "POST"
	}
	w.statusView.Status("Requesting:"+currentItem.expandURL, true)

	data, err := armclient.DoRequest(ctx, method, currentItem.expandURL)
	if err != nil {
		w.statusView.Status("Failed"+err.Error()+currentItem.expandURL, false)
	} else if currentItem.expandReturnType == actionType {
		w.title = "Action Succeeded: " + currentItem.expandURL
	}

	if currentItem.expandReturnType == resourceGroupType {
		var rgResponse armclient.ResourceGroupResponse
		err := json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		newItems := []TreeNode{}
		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, TreeNode{
				name:             rg.Name,
				id:               rg.ID,
				parentid:         currentItem.id,
				expandURL:        rg.ID + "/resources?api-version=2017-05-10",
				expandReturnType: resourceType,
				itemType:         resourceGroupType,
				deleteURL:        rg.ID + "?api-version=2017-05-10",
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = currentItem.name + ">Resource Groups"
	}

	if currentItem.expandReturnType == resourceType {
		var resourceResponse armclient.ResourceReseponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			panic(err)
		}

		newItems := []TreeNode{}
		// Add Deployments
		if currentItem.itemType == resourceGroupType {
			newItems = append(newItems, TreeNode{
				parentid:         currentItem.id,
				namespace:        "None",
				name:             style.Subtle("[Microsoft.Resources]") + "\n   Deployments",
				id:               currentItem.id,
				expandURL:        currentItem.id + "/providers/Microsoft.Resources/deployments?api-version=2017-05-10",
				expandReturnType: deploymentType,
				itemType:         resourceType,
				deleteURL:        "NotSupported",
			})
		}
		for _, rg := range resourceResponse.Resources {
			resourceAPIVersion, err := armclient.GetAPIVersion(rg.Type)
			if err != nil {
				w.statusView.Status("Failed to find an api version: "+err.Error(), false)
			}
			newItems = append(newItems, TreeNode{
				name:             style.Subtle("["+rg.Type+"] \n   ") + rg.Name,
				parentid:         currentItem.id,
				namespace:        strings.Split(rg.Type, "/")[0], // We just want the namespace not the subresource
				armType:          rg.Type,
				id:               rg.ID,
				expandURL:        rg.ID + "?api-version=" + resourceAPIVersion,
				expandReturnType: "none",
				itemType:         resourceType,
				deleteURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = w.title + ">" + currentItem.name
	}

	if currentItem.expandReturnType == "none" {
		w.title = w.title + ">" + currentItem.name
	}
	if err == nil {
		w.statusView.Status("Fetching item completed:"+currentItem.expandURL, false)
	}
	w.contentView.SetContent(style.Title(w.title) + "\n-------------------------------------------------------\n" + data)

}

// ChangeSelection updates the selected item
func (w *ListWidget) ChangeSelection(i int) {
	if i >= len(w.items) || i < 0 {
		return
	}
	w.selected = i
}

// CurrentSelection returns the current selection int
func (w *ListWidget) CurrentSelection() int {
	return w.selected
}

// CurrentItem returns the selected item as a treenode
func (w *ListWidget) CurrentItem() *TreeNode {
	return &w.items[w.selected]
}
