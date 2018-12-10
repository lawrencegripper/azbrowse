package main

import (
	"encoding/json"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowser/armclient"
)

const (
	SubscriptionType  = "subscription"
	ResourceGroupType = "resourcegroup"
	ResourceType      = "resource"
)

type treeNode struct {
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

	selected int
}

func NewListWidget(x, y, w, h int, items []string, selected int, contentView *ItemWidget) *ListWidget {
	return &ListWidget{x: x, y: y, w: w, h: h, contentView: contentView}
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

	w.items = newList
}

func (w *ListWidget) ExpandCurrentSelection() {
	currentItem := w.items[w.selected]

	data, err := armclient.DoRequest(currentItem.expandURL)

	w.contentView.Content = data

	if currentItem.expandReturnType == ResourceGroupType {
		var rgResponse armclient.ResourceGroupResponse
		err = json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		newItems := []treeNode{}
		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, treeNode{
				name:             rg.Name,
				id:               rg.ID,
				expandURL:        rg.ID + "/resources?api-version=2017-05-10",
				expandReturnType: ResourceType,
				itemType:         ResourceGroupType,
			})
		}
		w.items = newItems
		w.selected = 0
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
				name:             "[" + rg.Type + "] - " + rg.Name,
				id:               rg.ID,
				expandURL:        rg.ID + "?api-version=latest", //TOD: This won't work we need to query the provider api to get the resource api version https://zimmergren.net/developing-with-azure-resource-manager-part-5-tip-get-the-available-api-version-for-the-arm-endpoints/
				expandReturnType: "none",
				itemType:         ResourceType,
			})
		}
		w.items = newItems
		w.selected = 0

	}
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
