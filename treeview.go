package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lawrencegripper/azbrowse/tracing"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/handlers"
	"github.com/lawrencegripper/azbrowse/style"
)

const (
	subscriptionType  = "subscription"
	resourceGroupType = "resourcegroup"
	resourceType      = "resource"
	deploymentType    = "deployment"
	actionType        = "action"
)

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y        int
	w, h        int
	items       []handlers.TreeNode
	contentView *ItemWidget
	statusView  *StatusbarWidget
	navStack    Stack
	title       string
	ctx         context.Context
	selected    int
	view        *gocui.View
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
	w.view = v

	if len(w.items) < 1 {
		return nil
	}

	linesUsedCount := 0
	allItems := make([]string, 0, len(w.items))

	allItems = append(allItems, style.Separator("  ---\n"))

	for i, s := range w.items {
		var itemToShow string
		if i == w.selected {
			itemToShow = "▶ "
		} else {
			itemToShow = "  "
		}
		itemToShow = itemToShow + s.Display + "\n" + style.Separator("  ---") + "\n"

		linesUsedCount = linesUsedCount + strings.Count(itemToShow, "\n")
		allItems = append(allItems, itemToShow)
	}

	linesPerItem := linesUsedCount / len(w.items)
	maxItemsCanShow := (w.h / linesPerItem) - 1 // minus 1 to be on the safe side

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
func (w *ListWidget) SetNodes(nodes []handlers.TreeNode) {
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
	newList := []handlers.TreeNode{}
	for _, sub := range subs.Subs {
		newList = append(newList, handlers.TreeNode{
			Display:          sub.DisplayName,
			Name:             sub.DisplayName,
			ID:               sub.ID,
			ExpandURL:        sub.ID + "/resourceGroups?api-version=2018-05-01",
			ItemType:         subscriptionType,
			ExpandReturnType: resourceGroupType,
		})
	}

	w.title = "Subscriptions"
	w.items = newList
	w.view.Title = w.title
}

// GoBack takes the user back to preview view
func (w *ListWidget) GoBack() {
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		return
	}
	w.contentView.SetContent(previousPage.Data, "Response")
	w.selected = 0
	w.items = previousPage.Value
	w.title = previousPage.Title
	w.selected = previousPage.Selection
	w.view.Title = w.title
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {

	currentItem := w.items[w.selected]
	if currentItem.ExpandReturnType != "none" && currentItem.ExpandReturnType != actionType {
		// Capture current view to navstack
		w.navStack.Push(&Page{
			Data:      w.contentView.GetContent(),
			Value:     w.items,
			Title:     w.title,
			Selection: w.selected,
		})
	}
	span, ctx := tracing.StartSpanFromContext(w.ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	method := "GET"
	if currentItem.ExpandReturnType == actionType {
		method = "POST"
	}
	w.statusView.Status("Requesting:"+currentItem.ExpandURL, true)

	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		w.statusView.Status("Failed"+err.Error()+currentItem.ExpandURL, false)
	} else if currentItem.ExpandReturnType == actionType {
		w.title = "Action Succeeded: " + currentItem.ExpandURL
	}

	if currentItem.ExpandReturnType == resourceGroupType {
		var rgResponse armclient.ResourceGroupResponse
		err := json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		newItems := []handlers.TreeNode{}
		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, handlers.TreeNode{
				Name:             rg.Name,
				Display:          rg.Name + " " + drawStatus(rg.Properties.ProvisioningState),
				ID:               rg.ID,
				Parentid:         currentItem.ID,
				ExpandURL:        rg.ID + "/resources?api-version=2017-05-10",
				ExpandReturnType: resourceType,
				ItemType:         resourceGroupType,
				DeleteURL:        rg.ID + "?api-version=2017-05-10",
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = currentItem.Name + ">Resource Groups"
	}

	if currentItem.ExpandReturnType == resourceType {
		var resourceResponse armclient.ResourceReseponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			panic(err)
		}

		newItems := []handlers.TreeNode{}
		// Add Deployments
		if currentItem.ItemType == resourceGroupType {
			newItems = append(newItems, handlers.TreeNode{
				Parentid:         currentItem.ID,
				Namespace:        "None",
				Display:          style.Subtle("[Microsoft.Resources]") + "\n  Deployments",
				Name:             "Deployments",
				ID:               currentItem.ID,
				ExpandURL:        currentItem.ID + "/providers/Microsoft.Resources/deployments?api-version=2017-05-10",
				ExpandReturnType: deploymentType,
				ItemType:         resourceType,
				DeleteURL:        "NotSupported",
			})
		}
		for _, rg := range resourceResponse.Resources {
			resourceAPIVersion, err := armclient.GetAPIVersion(rg.Type)
			if err != nil {
				w.statusView.Status("Failed to find an api version: "+err.Error(), false)
			}
			newItems = append(newItems, handlers.TreeNode{
				Display:          style.Subtle("["+rg.Type+"] \n  ") + rg.Name,
				Name:             rg.Name,
				Parentid:         currentItem.ID,
				Namespace:        strings.Split(rg.Type, "/")[0], // We just want the namespace not the subresource
				ArmType:          rg.Type,
				ID:               rg.ID,
				ExpandURL:        rg.ID + "?api-version=" + resourceAPIVersion,
				ExpandReturnType: "none",
				ItemType:         resourceType,
				DeleteURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = w.title + ">" + currentItem.Name
	}

	if currentItem.ExpandReturnType == "none" {
		w.title = w.title + ">" + currentItem.Name
		w.contentView.SetContent(data, "[CTRL+F -> Fullscreen|CTRL+A -> Actions] "+currentItem.Name)
		w.view.Title = w.title
	} else {
		w.contentView.SetContent(data, "[CTRL+F -> Fullscreen] "+currentItem.Name)
	}

	if err == nil {
		w.statusView.Status("Fetching item completed:"+currentItem.ExpandURL, false)
	}

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
func (w *ListWidget) CurrentItem() *handlers.TreeNode {
	return &w.items[w.selected]
}

func drawStatus(s string) string {
	switch s {
	case "Deleting":
		return "☠"
	case "Updating":
		return "⚙️"
	case "Resuming":
		return "⚙️"
	case "Starting":
		return "⚙️"
	case "Provisioning":
		return "⌛"
	case "Creating":
		return "🧱"
	case "Preparing":
		return "🧱"
	case "Scaling":
		return "𝄩"
	case "Suspended":
		return "⛔"
	case "Suspending":
		return "⛔"
	}
	return ""
}
