package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/eventing"

	"github.com/lawrencegripper/azbrowse/tracing"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/handlers"
	"github.com/lawrencegripper/azbrowse/style"
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

	// If the title is getting too long trim things
	// down from the front
	if len(w.title) > w.w {
		trimLength := len(w.title) - w.w + 5 // Add five for spacing and elipsis
		w.title = ".." + w.title[trimLength:]
	}

	w.view.Title = w.title

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
	//Todo: Evaluate moving this to a handler
	newList := []handlers.TreeNode{}
	for _, sub := range subs.Subs {
		newList = append(newList, handlers.TreeNode{
			Display:   sub.DisplayName,
			Name:      sub.DisplayName,
			ID:        sub.ID,
			ExpandURL: sub.ID + "/resourceGroups?api-version=2018-05-01",
			ItemType:  handlers.SubscriptionType,
		})
	}

	w.title = "Subscriptions"
	w.items = newList
}

// Refresh refreshes the current view
func (w *ListWidget) Refresh() {
	w.statusView.Status("Refreshing", true)
	currentSelection := w.CurrentSelection()

	w.GoBack()
	w.ExpandCurrentSelection()

	w.ChangeSelection(currentSelection)
	w.statusView.Status("Done refreshing", false)
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
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {
	if w.title == "Subscriptions" {
		w.title = ""
	}

	currentItem := w.items[w.selected]

	_, done := eventing.SendStatusEvent(eventing.StatusEvent{
		InProgress: true,
		Message:    "Opening: " + currentItem.ID,
	})

	if currentItem.ExpandReturnType != "none" && currentItem.ExpandReturnType != handlers.ActionType {
		// Capture current view to navstack
		w.navStack.Push(&Page{
			Data:      w.contentView.GetContent(),
			Value:     w.items,
			Title:     w.title,
			Selection: w.selected,
		})
	}

	newItems := []handlers.TreeNode{}

	span, ctx := tracing.StartSpanFromContext(w.ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	// New handler approach
	handlerExpanding := 0
	completedExpands := make(chan handlers.ExpanderResult)

	// Check which expanders are interested and kick them off
	spanQuery, _ := tracing.StartSpanFromContext(ctx, "querexpanders", tracing.SetTag("item", currentItem))
	for _, h := range handlers.Register {
		doesExpand, err := h.DoesExpand(w.ctx, currentItem)
		spanQuery.SetTag(h.Name(), doesExpand)
		if err != nil {
			panic(err)
		}
		if !doesExpand {
			continue
		}

		// Fire each handler in parallel
		hCurrent := h // capture current iteration variable
		go func() {
			completedExpands <- hCurrent.Expand(ctx, currentItem)
		}()

		handlerExpanding = handlerExpanding + 1
	}
	spanQuery.Finish()

	// Lets give all the expanders 45secs to completed (unless debugging)
	hasPrimaryResponse := false
	var timeout <-chan time.Time
	if enableTracing {
		timeout = time.After(time.Second * 600)
	} else {
		timeout = time.After(time.Second * 45)
	}
	for index := 0; index < handlerExpanding; index++ {
		select {
		case done := <-completedExpands:
			span, _ := tracing.StartSpanFromContext(ctx, "subexpand:"+done.SourceDescription, tracing.SetTag("result", done))
			// Did it fail?
			if done.Err != nil {
				eventing.SendStatusEvent(eventing.StatusEvent{
					Failure: true,
					Message: "Expander '" + done.SourceDescription + "' failed on resource: " + currentItem.ID,
					Timeout: time.Duration(time.Second * 15),
				})
			}
			if done.IsPrimaryResponse {
				if hasPrimaryResponse {
					panic("Two handlers returned a primary response for this item... failing")
				}
				// Log that we have a primary response
				hasPrimaryResponse = true
				w.contentView.SetContent(done.Response, "[CTRL+F -> Fullscreen|CTRL+A -> Actions] "+currentItem.Name)
			}
			if done.Nodes == nil {
				continue
			}
			// Add the items it found
			newItems = append(newItems, *done.Nodes...)
			span.Finish()
		case <-timeout:
			eventing.SendStatusEvent(eventing.StatusEvent{
				Failure: true,
				Message: "Timed out opening:" + currentItem.ID,
				Timeout: time.Duration(time.Second * 10),
			})
			return
		}
	}

	// Update the list if we have sub items from the expanders
	// or return the default experience for and unknown item
	if hasPrimaryResponse || len(newItems) > 0 {
		w.items = newItems
		w.selected = 0
	}

	// Use the default handler to get the resource JSON for display
	if !hasPrimaryResponse {
		defaultHandler := handlers.DefaultExpander{}
		result := defaultHandler.Expand(ctx, currentItem)
		if result.Err != nil {
			panic(result.Err)
		}
		w.contentView.SetContent(result.Response, "[CTRL+F -> Fullscreen|CTRL+A -> Actions] "+currentItem.Name)
	}

	w.title = w.title + ">" + currentItem.Name
	done()
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
