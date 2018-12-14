package main

import (
	"encoding/json"
	"fmt"
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
	providerCacheKey  = "providerCache"
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

	selected int

	resourceAPIVersionLookup map[string]string
}

// NewListWidget creates a new instance
func NewListWidget(x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget) *ListWidget {
	return &ListWidget{x: x, y: y, w: w, h: h, contentView: contentView, statusView: status}
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

	go w.PopulateResourceAPILookup()

	w.title = "Subscriptions"
	w.items = newList
}

// GoBack takes the user back to preview view
func (w *ListWidget) GoBack() {
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		return
	}
	w.contentView.Content = previousPage.Data
	w.selected = 0
	w.items = previousPage.Value
	w.title = previousPage.Title
	w.selected = previousPage.Selection
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {
	currentItem := w.items[w.selected]
	if currentItem.expandReturnType != "none" {
		// Capture current view to navstack
		w.navStack.Push(&Page{
			Data:      w.contentView.Content,
			Value:     w.items,
			Title:     w.title,
			Selection: w.selected,
		})
	}

	w.statusView.Status("Fetching item:"+currentItem.expandURL, true)

	data, err := armclient.DoRequest("GET", currentItem.expandURL)

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
			newItems = append(newItems, TreeNode{
				name:             style.Subtle("["+rg.Type+"] \n   ") + rg.Name,
				parentid:         currentItem.id,
				namespace:        strings.Split(rg.Type, "/")[0], // We just want the namespace not the subresource
				armType:          rg.Type,
				id:               rg.ID,
				expandURL:        rg.ID + "?api-version=" + w.resourceAPIVersionLookup[rg.Type],
				expandReturnType: "none",
				itemType:         resourceType,
				deleteURL:        rg.ID + "?api-version=" + w.resourceAPIVersionLookup[rg.Type],
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = w.title + ">" + currentItem.name
	}

	if currentItem.expandReturnType == "none" {
		w.title = w.title + ">" + currentItem.name
	}

	w.statusView.Status("Fetching item completed:"+currentItem.expandURL, false)
	w.contentView.Content = style.Title(w.title) + "\n-------------------------------------------------------\n" + data

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

// PopulateResourceAPILookup is used to build a cache of resourcetypes -> api versions
// this is needed when requesting details from a resource as APIVersion isn't known and is required
func (w *ListWidget) PopulateResourceAPILookup() {
	if w.resourceAPIVersionLookup == nil {
		w.statusView.Status("Getting provider data from cache", true)
		// Get data from cache
		providerData, err := get(providerCacheKey)

		w.statusView.Status("Getting provider data from cache: Completed", false)

		if err != nil || providerData == "" {
			w.statusView.Status("Getting provider data from API", true)

			// Get Subscriptions
			data, err := armclient.DoRequest("GET", "/providers?api-version=2017-05-10")
			if err != nil {
				panic(err)
			}
			var providerResponse armclient.ProvidersResponse
			err = json.Unmarshal([]byte(data), &providerResponse)
			if err != nil {
				panic(err)
			}

			w.resourceAPIVersionLookup = make(map[string]string)
			for _, provider := range providerResponse.Providers {
				for _, resourceType := range provider.ResourceTypes {
					w.resourceAPIVersionLookup[provider.Namespace+"/"+resourceType.ResourceType] = resourceType.APIVersions[0]
				}
			}

			bytes, err := json.Marshal(w.resourceAPIVersionLookup)
			if err != nil {
				panic(err)
			}
			providerData = string(bytes)

			put(providerCacheKey, providerData)
			w.statusView.Status("Getting provider data from API: Completed", false)

		} else {
			var providerCache map[string]string
			err = json.Unmarshal([]byte(providerData), &providerCache)
			if err != nil {
				panic(err)
			}
			w.resourceAPIVersionLookup = providerCache
			w.statusView.Status("Getting provider data from cache: Completed", false)

		}

	}
}
