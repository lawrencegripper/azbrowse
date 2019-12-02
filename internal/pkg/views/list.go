package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y  int
	w, h  int
	g     *gocui.Gui
	items []*expanders.TreeNode

	filteredItems []*expanders.TreeNode
	filterString  string

	contentView          *ItemWidget
	statusView           *StatusbarWidget
	navStack             Stack
	title                string
	ctx                  context.Context
	selected             int
	expandedNodeItem     *expanders.TreeNode
	view                 *gocui.View
	enableTracing        bool
	FullscreenKeyBinding string
	ActionKeyBinding     string
	lastTopIndex         int
}

// ListNavigatedEventState captures the state when raising a `list.navigated` event
type ListNavigatedEventState struct {
	Success      bool                  // True if this was a successful navigation
	NewNodes     []*expanders.TreeNode // If Success==true this contains the new nodes
	ParentNodeID string                // This is the ID of the item expanded.
	NodeID       string                // The current nodes id
	IsBack       bool                  // Was this a navigation back?
}

// NewListWidget creates a new instance
func NewListWidget(ctx context.Context, x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget, enableTracing bool, title string, g *gocui.Gui) *ListWidget {
	listWidget := &ListWidget{ctx: ctx, x: x, y: y, w: w, h: h, contentView: contentView, statusView: status, enableTracing: enableTracing, lastTopIndex: 0, filterString: "", title: title, g: g}
	return listWidget
}

func (w *ListWidget) itemCount() int {
	if w.filterString == "" {
		return len(w.items)
	}

	return len(w.filteredItems)
}

// SetFilter sets the filter to be applied to list items
func (w *ListWidget) SetFilter(filterString string) {
	filteredItems := []*expanders.TreeNode{}
	for _, item := range w.items {
		if strings.Contains(strings.ToLower(item.Display), filterString) {
			filteredItems = append(filteredItems, item)
		}
	}

	w.selected = 0
	w.filterString = filterString
	w.filteredItems = filteredItems

	w.g.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

// ClearFilter clears a filter if applied
func (w *ListWidget) ClearFilter() {
	w.filterString = ""
}

func (w *ListWidget) itemsToShow() []*expanders.TreeNode {
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
	eventing.Publish("list.prenavigate", "GOBACK")

	if w.filterString != "" {
		// initial Back action is to clear filter, subsequent Back actions are normal
		w.ClearFilter()
		return
	}
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		eventing.Publish("list.navigated", ListNavigatedEventState{Success: false})
		return
	}
	w.contentView.SetContent(previousPage.Data, previousPage.DataType, "Response")
	w.selected = 0
	w.items = previousPage.Value
	w.title = previousPage.Title
	w.selected = previousPage.Selection
	w.expandedNodeItem = previousPage.ExpandedNodeItem

	eventing.Publish("list.navigated", ListNavigatedEventState{
		Success:      true,
		NewNodes:     w.items,
		ParentNodeID: w.expandedNodeItem.Parentid,
		IsBack:       true,
	})
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {

	if w.title == "Subscriptions" {
		w.title = ""
	}

	currentItem := w.CurrentItem()

	newTitle := fmt.Sprintf("[%s-> Fullscreen|%s -> Actions] %s", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding), currentItem.Name)

	eventing.Publish("list.prenavigate", currentItem.ID)

	newContent, newItems, err := expanders.ExpandItem(w.ctx, currentItem)
	if err != nil { // Don't need to display error as expander emits status event on error
		// Set parameters to trigger non-successful `list.navigated` event
		newItems = []*expanders.TreeNode{}
		newContent = nil
		newTitle = ""
	}
	w.Navigate(newItems, newContent, newTitle)
}

// Navigate updates the currently selected list nodes, title and details content
func (w *ListWidget) Navigate(nodes []*expanders.TreeNode, content *expanders.ExpanderResponse, title string) {

	if len(nodes) == 0 && content == nil && title == "" {
		eventing.Publish("list.navigated", ListNavigatedEventState{Success: false})
	}
	currentItem := w.CurrentItem()
	if len(nodes) > 0 {
		w.expandedNodeItem = currentItem
		w.SetNodes(nodes)
	}

	w.contentView.SetContent(content.Response, content.ResponseType, title)
	if currentItem != nil {
		w.title = w.title + ">" + currentItem.Name
	}

	if w.expandedNodeItem != nil {
		eventing.Publish("list.navigated", ListNavigatedEventState{
			Success:      true,
			NewNodes:     nodes,
			ParentNodeID: w.expandedNodeItem.ID,
			NodeID:       currentItem.ID,
		})
	} else {
		eventing.Publish("list.navigated", ListNavigatedEventState{
			Success:      true,
			NewNodes:     nodes,
			ParentNodeID: "root",
			NodeID:       "root",
		})
	}
}

// GetNodes returns the currently listed nodes
func (w *ListWidget) GetNodes() []*expanders.TreeNode {
	return w.items
}

// SetNodes allows others to set the list nodes
func (w *ListWidget) SetNodes(nodes []*expanders.TreeNode) {

	// Capture current view to navstack
	if w.HasCurrentItem() {
		w.navStack.Push(&Page{
			Data:             w.contentView.GetContent(),
			DataType:         w.contentView.GetContentType(),
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

	w.selected = 0
	w.items = nodes
	w.ClearFilter()
}

// ChangeSelection updates the selected item
func (w *ListWidget) ChangeSelection(i int) {
	if i >= w.itemCount() {
		// panic("invalid changeSelection out of range")
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
func (w *ListWidget) CurrentItem() *expanders.TreeNode {
	if w.HasCurrentItem() {
		return w.itemsToShow()[w.selected]
	}
	return nil
}

// CurrentExpandedItem returns the currently expanded item as a treenode
func (w *ListWidget) CurrentExpandedItem() *expanders.TreeNode {
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
