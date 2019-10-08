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
)

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y  int
	w, h  int
	g     *gocui.Gui
	items []*handlers.TreeNode

	filteredItems []*handlers.TreeNode
	filterString  string

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
func NewListWidget(ctx context.Context, x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget, enableTracing bool, title string, g *gocui.Gui) *ListWidget {
	listWidget := &ListWidget{ctx: ctx, x: x, y: y, w: w, h: h, contentView: contentView, statusView: status, enableTracing: enableTracing, lastTopIndex: 0, filterString: "", title: title, g: g}
	go func() {
		filterChannel := eventing.SubscribeToTopic("filter")
		for {
			filterStringInterface := <-filterChannel
			filterString := strings.ToLower(strings.TrimSpace(strings.Replace(filterStringInterface.(string), "/", "", 1)))

			filteredItems := []*handlers.TreeNode{}
			for _, item := range listWidget.items {
				if strings.Contains(strings.ToLower(item.Display), listWidget.filterString) {
					filteredItems = append(filteredItems, item)
				}
			}

			listWidget.selected = 0
			listWidget.filterString = filterString
			listWidget.filteredItems = filteredItems

			g.Update(func(gui *gocui.Gui) error {
				return nil
			})
		}
	}()
	return listWidget
}

func (w *ListWidget) itemCount() int {
	if w.filterString == "" {
		return len(w.items)
	}

	return len(w.filteredItems)
}

// ClearFilter clears a filter if applied
func (w *ListWidget) ClearFilter() {
	w.filterString = ""
}

func (w *ListWidget) itemsToShow() []*handlers.TreeNode {
	if w.filterString == "" {
		return w.items
	}

	return w.filteredItems
}

func highlightText(displayText string, highlight string) string {
	if highlight == "" {
		return displayText
	}
	index := strings.Index(strings.ToLower(displayText), highlight)
	if index < 0 {
		return displayText
	}
	return displayText[:index] + style.Highlight(displayText[index:index+len(highlight)]) + displayText[index+len(highlight):]
}

// Layout draws the widget in the gocui view
func (w *ListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("listWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	w.view = v

	if w.itemCount() < 1 {
		return nil
	}

	linesUsedCount := 0
	renderedItems := make([]string, 0, w.itemCount())

	renderedItems = append(renderedItems, style.Separator("  ---\n"))

	for i, s := range w.itemsToShow() {
		var itemToShow string
		if i == w.selected {
			itemToShow = "▶ "
		} else {
			itemToShow = "  "
		}

		itemToShow = itemToShow + highlightText(s.Display, w.filterString) + " " + s.StatusIndicator + "\n" + style.Separator("  ---") + "\n"

		linesUsedCount = linesUsedCount + strings.Count(itemToShow, "\n")
		renderedItems = append(renderedItems, itemToShow)
	}

	linesPerItem := linesUsedCount / w.itemCount()
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
	if bottomIndex > len(renderedItems) {
		bottomIndex = len(renderedItems) - 1
	}

	for index := topIndex; index < bottomIndex+1; index++ {
		if index < len(renderedItems) {
			item := renderedItems[index]
			fmt.Fprint(v, item)
		}
	}

	// If the title is getting too long trim things
	// down from the front
	title := w.title
	if w.filterString != "" {
		title += "[filter=" + w.filterString + "]"
	}
	if len(title) > w.w {
		trimLength := len(title) - w.w + 5 // Add five for spacing and elipsis
		title = ".." + title[trimLength:]
	}

	w.view.Title = title

	return nil
}

// Refresh refreshes the current view
func (w *ListWidget) Refresh() {
	w.statusView.Status("Refreshing", true)
	currentSelection := w.CurrentSelection()

	w.GoBack()
	w.ExpandCurrentSelection()

	w.ChangeSelection(currentSelection)
	w.statusView.Status("Done refreshing", false)

	w.ClearFilter()
}

// GoBack takes the user back to preview view
func (w *ListWidget) GoBack() {
	if w.filterString != "" {
		// initial Back action is to clear filter, subsequent Back actions are normal
		w.ClearFilter()
		return
	}
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

	currentItem := w.CurrentItem()
	w.ClearFilter()

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
	var newContent string
	var newTitle string

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
				newContent = done.Response
				newTitle = fmt.Sprintf("[%s-> Fullscreen|%s -> Actions] %s", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding), currentItem.Name)
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
		newContent = result.Response
		newTitle = fmt.Sprintf("[%s -> Fullscreen|%s -> Actions] %s", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding), currentItem.Name)
	}

	w.Navigate(newItems, newContent, newTitle)

	if !observedError {
		done()
	}
}

// Navigate updates the currently selected list nodes, title and details content
func (w *ListWidget) Navigate(nodes []*handlers.TreeNode, content string, title string) {
	currentItem := w.CurrentItem()
	if len(nodes) > 0 {
		w.SetNodes(nodes)
	}
	w.expandedNodeItem = currentItem
	w.contentView.SetContent(content, title)
	if currentItem != nil {
		w.title = w.title + ">" + currentItem.Name
	}

	eventing.Publish("list.navigated", nodes)
}

// SetNodes allows others to set the list nodes
func (w *ListWidget) SetNodes(nodes []*handlers.TreeNode) {
	w.selected = 0

	// Capture current view to navstack
	if w.HasCurrentItem() {
		w.navStack.Push(&Page{
			Data:             w.contentView.GetContent(),
			Value:            w.items,
			Title:            w.title,
			Selection:        w.selected,
			ExpandedNodeItem: w.CurrentItem(),
		})

		currentID := w.CurrentItem().ID
		for _, node := range nodes {
			if node.ID == currentID {
				panic(fmt.Errorf("ids must be unique or the navigate command breaks"))
			}
		}
	}

	w.items = nodes
	w.ClearFilter()
}

// ChangeSelection updates the selected item
func (w *ListWidget) ChangeSelection(i int) {
	if i >= w.itemCount() {
		i = w.itemCount() - 1
	} else if i < 0 {
		i = 0
	}
	w.selected = i
}

// CurrentSelection returns the current selection int
func (w *ListWidget) CurrentSelection() int {
	return w.selected
}

// HasCurrentItem indicates whether there is a current item
func (w *ListWidget) HasCurrentItem() bool {
	return w.selected < len(w.items)
}

// CurrentItem returns the selected item as a treenode
func (w *ListWidget) CurrentItem() *handlers.TreeNode {
	if w.HasCurrentItem() {
		return w.itemsToShow()[w.selected]
	}
	return nil
}

// CurrentExpandedItem returns the currently expanded item as a treenode
func (w *ListWidget) CurrentExpandedItem() *handlers.TreeNode {
	return w.expandedNodeItem
}

// MovePageDown changes the selection to be a page further down the list
func (w *ListWidget) MovePageDown() {
	i := w.selected

	for remainingLinesToPage := w.h; remainingLinesToPage > 0 && i < w.itemCount(); i++ {
		item := w.itemsToShow()[i]
		remainingLinesToPage -= strings.Count(item.Display, "\n") + 1 // +1 as there is an implicit newline
		remainingLinesToPage--                                        // separator
	}

	w.ChangeSelection(i)
}

// MovePageUp changes the selection to be a page further up the list
func (w *ListWidget) MovePageUp() {
	i := w.selected

	for remainingLinesToPage := w.h; remainingLinesToPage > 0 && i >= 0; i-- {
		item := w.itemsToShow()[i]
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
	w.ChangeSelection(w.itemCount() - 1)
}

// MoveUp moves the selection up one item
func (w *ListWidget) MoveUp() {
	w.ChangeSelection(w.CurrentSelection() - 1)
}

// MoveDown moves the selection down one item
func (w *ListWidget) MoveDown() {
	w.ChangeSelection(w.CurrentSelection() + 1)
}
