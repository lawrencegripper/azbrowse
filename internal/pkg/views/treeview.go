package views

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/handlers"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y                 int
	w, h                 int
	items                []*handlers.TreeNode
	contentView          *ItemWidget
	statusView           *StatusbarWidget
	navStack             Stack
	title                string
	ctx                  context.Context
	selected             int
	expandedNodeItem     *handlers.TreeNode
	view                 *gocui.View
	enableTracing        bool
	FullscreenKeyBinding string
	ActionKeyBinding     string
	lastTopIndex         int
}

// NewListWidget creates a new instance
func NewListWidget(ctx context.Context, x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget, enableTracing bool) *ListWidget {
	return &ListWidget{ctx: ctx, x: x, y: y, w: w, h: h, contentView: contentView, statusView: status, enableTracing: enableTracing, lastTopIndex: 0}
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
		itemToShow = itemToShow + s.Display + " " + s.StatusIndicator + "\n" + style.Separator("  ---") + "\n"

		linesUsedCount = linesUsedCount + strings.Count(itemToShow, "\n")
		allItems = append(allItems, itemToShow)
	}

	linesPerItem := linesUsedCount / len(w.items)
	maxItemsCanShow := (w.h / linesPerItem) - 1 // minus 1 to be on the safe side

	topIndex := w.lastTopIndex
	bottomIndex := w.lastTopIndex + maxItemsCanShow

	if w.selected >= bottomIndex {
		// need to adjust down
		diff := w.selected - bottomIndex + 1
		topIndex += diff
		bottomIndex += diff
	}
	if w.selected < topIndex {
		// need to adjust up
		diff := topIndex - w.selected
		topIndex -= diff
		bottomIndex -= diff
	}
	w.lastTopIndex = topIndex
	if bottomIndex > len(allItems) {
		bottomIndex = len(allItems) - 1
	}

	for index := topIndex; index < bottomIndex+1; index++ {
		if index < len(allItems) {
			item := allItems[index]
			fmt.Fprint(v, item)
		}
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
func (w *ListWidget) SetNodes(nodes []*handlers.TreeNode) {
	w.selected = 0
	// Capture current view to navstack
	w.navStack.Push(&Page{
		Data:             w.contentView.GetContent(),
		Value:            w.items,
		Title:            w.title,
		Selection:        w.selected,
		ExpandedNodeItem: w.CurrentItem(),
	})
	w.items = nodes
}

// SetSubscriptions starts vaidation with the subs found
func (w *ListWidget) SetSubscriptions(subs armclient.SubResponse) {
	//Todo: Evaluate moving this to a handler
	newList := []*handlers.TreeNode{}
	for _, sub := range subs.Subs {
		newList = append(newList, &handlers.TreeNode{
			Display:        sub.DisplayName,
			Name:           sub.DisplayName,
			ID:             sub.ID,
			ExpandURL:      sub.ID + "/resourceGroups?api-version=2018-05-01",
			ItemType:       handlers.SubscriptionType,
			SubscriptionID: sub.SubscriptionID,
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
	w.expandedNodeItem = previousPage.ExpandedNodeItem
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

	newItems := []*handlers.TreeNode{}

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
	if w.enableTracing {
		timeout = time.After(time.Second * 600)
	} else {
		timeout = time.After(time.Second * 45)
	}

	observedError := false
	for index := 0; index < handlerExpanding; index++ {
		select {
		case done := <-completedExpands:
			span, _ := tracing.StartSpanFromContext(ctx, "subexpand:"+done.SourceDescription, tracing.SetTag("result", done))
			// Did it fail?
			if done.Err != nil {
				eventing.SendStatusEvent(eventing.StatusEvent{
					Failure: true,
					Message: "Expander '" + done.SourceDescription + "' failed on resource: " + currentItem.ID + "Err: " + done.Err.Error(),
					Timeout: time.Duration(time.Second * 15),
				})
				observedError = true
			}
			if done.IsPrimaryResponse {
				if hasPrimaryResponse {
					panic("Two handlers returned a primary response for this item... failing")
				}
				// Log that we have a primary response
				hasPrimaryResponse = true
				w.contentView.SetContent(done.Response, fmt.Sprintf("[%s-> Fullscreen|%s -> Actions] %s", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding), currentItem.Name))
			}
			if done.Nodes == nil {
				continue
			}
			// Add the items it found
			newItems = append(newItems, done.Nodes...)
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

	if len(newItems) > 0 {
		// Capture current view to navstack as we're viewing an item with children
		w.navStack.Push(&Page{
			Data:             w.contentView.GetContent(),
			Value:            w.items,
			Title:            w.title,
			Selection:        w.selected,
			ExpandedNodeItem: w.CurrentItem(),
		})
		// Show new items and move cursor to top
		w.items = newItems
		w.selected = 0
	}
	w.expandedNodeItem = currentItem

	// Use the default handler to get the resource JSON for display
	defaultExpanderWorksOnThisItem, _ := handlers.DefaultExpanderInstance.DoesExpand(ctx, currentItem)
	if !hasPrimaryResponse && defaultExpanderWorksOnThisItem {
		result := handlers.DefaultExpanderInstance.Expand(ctx, currentItem)
		if result.Err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: true,
				Message:    "Failed to expand resource: " + result.Err.Error(),
				Timeout:    time.Duration(time.Second * 3),
			})
		}
		w.contentView.SetContent(result.Response, fmt.Sprintf("[%s -> Fullscreen|%s -> Actions] %s", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding), currentItem.Name))
	}

	w.title = w.title + ">" + currentItem.Name
	if !observedError {
		done()
	}
}

// ChangeSelection updates the selected item
func (w *ListWidget) ChangeSelection(i int) {
	if i >= len(w.items) {
		i = len(w.items) - 1
	} else if i < 0 {
		i = 0
	}
	w.selected = i
}

// CurrentSelection returns the current selection int
func (w *ListWidget) CurrentSelection() int {
	return w.selected
}

// CurrentItem returns the selected item as a treenode
func (w *ListWidget) CurrentItem() *handlers.TreeNode {
	return w.items[w.selected]
}

// CurrentExpandedItem returns the currently expanded item as a treenode
func (w *ListWidget) CurrentExpandedItem() *handlers.TreeNode {
	return w.expandedNodeItem
}

// MovePageDown changes the selection to be a page further down the list
func (w *ListWidget) MovePageDown() {
	i := w.selected

	for remainingLinesToPage := w.h; remainingLinesToPage > 0 && i < len(w.items); i++ {
		item := w.items[i]
		remainingLinesToPage -= strings.Count(item.Display, "\n") + 1 // +1 as there is an implicit newline
		remainingLinesToPage--                                        // separator
	}

	w.ChangeSelection(i)
}

// MovePageUp changes the selection to be a page further up the list
func (w *ListWidget) MovePageUp() {
	i := w.selected

	for remainingLinesToPage := w.h; remainingLinesToPage > 0 && i >= 0; i-- {
		item := w.items[i]
		remainingLinesToPage -= strings.Count(item.Display, "\n") + 1 // +1 as there is an implicit newline
		remainingLinesToPage--                                        // separator
	}

	w.ChangeSelection(i)
}

// MoveHome changes the selection to the top of the list
func (w *ListWidget) MoveHome() {
	w.ChangeSelection(0)
}

// MoveEnd changes the selection to the bottom of the list
func (w *ListWidget) MoveEnd() {
	w.ChangeSelection(len(w.items) - 1)
}

// MoveUp moves the selection up one item
func (w *ListWidget) MoveUp() {
	w.ChangeSelection(w.CurrentSelection() - 1)
}

// MoveDown moves the selection down one item
func (w *ListWidget) MoveDown() {
	w.ChangeSelection(w.CurrentSelection() + 1)
}
