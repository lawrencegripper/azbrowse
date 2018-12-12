package main

import (
	"encoding/json"
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowser/armclient"
	"github.com/lawrencegripper/azbrowser/style"
)

const (
	SubscriptionType  = "subscription"
	ResourceGroupType = "resourcegroup"
	ResourceType      = "resource"
	ProviderCacheKey  = "providerCache"
)

type treeNode struct {
	parentid         string
	id               string
	name             string
	expandURL        string
	itemType         string
	expandReturnType string
}

type ListWidget struct {
	x, y        int
	w, h        int
	items       []treeNode
	contentView *ItemWidget
	statusView  *StatusbarWidget
	navStack    Stack
	title       string

	selected int

	resourceApiVersionLookup map[string]string
}

func NewListWidget(x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget) *ListWidget {
	return &ListWidget{x: x, y: y, w: w, h: h, contentView: contentView, statusView: status}
}

func (w *ListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("listWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	for i, s := range w.items {
		if i == w.selected {
			fmt.Fprintf(v, "▶  ")
		} else {
			fmt.Fprintf(v, "   ")
		}
		fmt.Fprintf(v, s.name+"\n")
		// if s.isExpanded {
		// 	for _, sub
		// }
	}
	return nil
}

func (w *ListWidget) SetSubscriptions(subs armclient.SubResponse) {
	newList := []treeNode{}
	for _, sub := range subs.Subs {
		newList = append(newList, treeNode{
			name:             sub.DisplayName,
			id:               sub.ID,
			expandURL:        sub.ID + "/resourceGroups?api-version=2014-04-01",
			itemType:         SubscriptionType,
			expandReturnType: ResourceGroupType,
		})
	}

	go w.PopulateResourceAPILookup()

	w.title = "Subscriptions"
	w.items = newList
}

func (w *ListWidget) GoBack() {
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		return
	}
	w.contentView.Content = previousPage.Data
	w.selected = 0
	w.items = previousPage.Value
	w.title = previousPage.Title
}

func (w *ListWidget) ExpandCurrentSelection() {
	currentItem := w.items[w.selected]
	if currentItem.expandReturnType != "none" {
		// Capture current view to navstack
		w.navStack.Push(&Page{
			Data:  w.contentView.Content,
			Value: w.items,
			Title: w.title,
		})
	}

	w.statusView.Status("Fetching item:"+currentItem.expandURL, true)

	data, err := armclient.DoRequest(currentItem.expandURL)

	if currentItem.expandReturnType == ResourceGroupType {
		var rgResponse armclient.ResourceGroupResponse
		err := json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		newItems := []treeNode{}
		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, treeNode{
				name:             rg.Name,
				id:               rg.ID,
				parentid:         currentItem.id,
				expandURL:        rg.ID + "/resources?api-version=2017-05-10",
				expandReturnType: ResourceType,
				itemType:         ResourceGroupType,
			})
		}
		w.items = newItems
		w.selected = 0
		w.title = currentItem.name + ">Resource Groups"
	}

	if currentItem.expandReturnType == ResourceType {
		var resourceResponse armclient.ResourceReseponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			panic(err)
		}

		newItems := []treeNode{}
		for _, rg := range resourceResponse.Resources {
			newItems = append(newItems, treeNode{
				name:             style.Subtle("["+rg.Type+"] \n   ") + rg.Name,
				parentid:         currentItem.id,
				id:               rg.ID,
				expandURL:        rg.ID + "?api-version=" + w.resourceApiVersionLookup[rg.Type],
				expandReturnType: "none",
				itemType:         ResourceType,
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

func (w *ListWidget) ChangeSelection(i int) {
	if i >= len(w.items) || i < 0 {
		return
	}
	w.selected = i
}

func (w *ListWidget) CurrentSelection() int {
	return w.selected
}

func (w *ListWidget) CurrentItem() *treeNode {
	return &w.items[w.selected]
}

func (w *ListWidget) PopulateResourceAPILookup() {
	if w.resourceApiVersionLookup == nil {
		w.statusView.Status("Getting provider data from cache", true)
		// Get data from cache
		providerData, err := get(ProviderCacheKey)

		w.statusView.Status("Getting provider data from cache: Completed", false)

		if err != nil || providerData == "" {
			w.statusView.Status("Getting provider data from API", true)

			// Get Subscriptions
			data, err := armclient.DoRequest("/providers?api-version=2017-05-10")
			if err != nil {
				panic(err)
			}
			var providerResponse armclient.ProvidersResponse
			err = json.Unmarshal([]byte(data), &providerResponse)
			if err != nil {
				panic(err)
			}

			w.resourceApiVersionLookup = make(map[string]string)
			for _, provider := range providerResponse.Providers {
				for _, resourceType := range provider.ResourceTypes {
					w.resourceApiVersionLookup[provider.Namespace+"/"+resourceType.ResourceType] = resourceType.APIVersions[0]
				}
			}

			bytes, err := json.Marshal(w.resourceApiVersionLookup)
			if err != nil {
				panic(err)
			}
			providerData = string(bytes)

			put(ProviderCacheKey, providerData)
			w.statusView.Status("Getting provider data from API: Completed", false)

		} else {
			var providerCache map[string]string
			err = json.Unmarshal([]byte(providerData), &providerCache)
			if err != nil {
				panic(err)
			}
			w.resourceApiVersionLookup = providerCache
			w.statusView.Status("Getting provider data from cache: Completed", false)

		}

	}
}
