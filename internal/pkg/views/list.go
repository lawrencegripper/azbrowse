package views

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
)

// hitbox is used to enable mouse support on the list
// it tracks which item is drawn on which lines in the list
type hitbox struct {
	start     int
	end       int
	itemIndex int
}

// ListWidget hosts the left panel showing resources and controls the navigation
type ListWidget struct {
	x, y int
	w, h int
	g    *gocui.Gui

	contentView          *ItemWidget
	statusView           *StatusbarWidget
	navStack             Stack
	currentPage          *Page
	ctx                  context.Context
	view                 *gocui.View
	enableTracing        bool
	shouldRender         bool
	lastCalculatedHeight int
	hitBoxes             []hitbox
	// To avoid blocking the "UI" thread opening items is done in a go routine. This lock prevents duplicates.
	navLock      sync.Mutex
	isNavigating bool
	refreshLock  sync.Mutex
	isRefreshing bool
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
func NewListWidget(ctx context.Context, x, y, w, h int, items []string, selected int, contentView *ItemWidget, status *StatusbarWidget, enableTracing bool, title string, shouldRender bool, g *gocui.Gui) *ListWidget {
	listWidget := &ListWidget{ctx: ctx, x: x, y: y, w: w, h: h, contentView: contentView, statusView: status, enableTracing: enableTracing, shouldRender: shouldRender, g: g}
	return listWidget
}

func (w *ListWidget) itemCount() int {
	if w.currentPage == nil {
		return 0
	}
	if w.currentPage.FilterString == "" {
		return len(w.currentPage.Items)
	}

	return len(w.currentPage.FilteredItems)
}

// SetFilter sets the filter to be applied to list items
func (w *ListWidget) SetFilter(filterString string) {
	if w.currentPage == nil {
		return
	}
	filteredItems := []*expanders.TreeNode{}
	for _, item := range w.currentPage.Items {
		if strings.Contains(strings.ToLower(item.Display), filterString) {
			filteredItems = append(filteredItems, item)
		}
	}

	w.currentPage.Selection = 0
	w.currentPage.FilterString = filterString
	w.currentPage.FilteredItems = filteredItems

	w.g.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

// ClearFilter clears a filter if applied
func (w *ListWidget) ClearFilter() {
	if w.currentPage != nil {
		w.currentPage.FilterString = ""
	}
}

func (w *ListWidget) itemsToShow() []*expanders.TreeNode {
	if w.currentPage == nil {
		return []*expanders.TreeNode{}
	}
	if w.currentPage.FilterString == "" {
		return w.currentPage.Items
	}

	return w.currentPage.FilteredItems
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
	x0, y0, x1, y1 := getViewBounds(g, w.x, w.y, w.w, w.h)
	v, err := g.SetView("listWidget", x0, y0, x1, y1, 0)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	w.view = v

	width := x1 - x0 + 1
	height := y1 - y0 + 1
	w.lastCalculatedHeight = height
	if w.shouldRender {
		if w.itemCount() < 1 {
			return nil
		}

		linesUsedCount := 0
		completeString := strings.Builder{}
		completeString.WriteString(style.Separator("  ---\n"))
		for i, s := range w.itemsToShow() {
			// Calculate the start and end line of this item
			// this is used to translate the x,y of a mouse click
			// into an item index to expand
			itemHitbox := hitbox{
				start: linesUsedCount,
			}
			var itemToShow string
			if i == w.currentPage.Selection {
				itemToShow = "▶ "
			} else {
				itemToShow = "  "
			}

			itemToShow = itemToShow + highlightText(s.Display, w.currentPage.FilterString) + " " + s.StatusIndicator + "\n" + style.Separator("  ---") + "\n"

			linesUsedCount += strings.Count(itemToShow, "\n")
			itemHitbox.end = linesUsedCount
			itemHitbox.itemIndex = i
			completeString.WriteString(itemToShow) //nolint: errcheck
			w.hitBoxes = append(w.hitBoxes, itemHitbox)
		}

		// Hack, find the line currently selected
		lines := strings.Split(completeString.String(), "\n")
		var selectedLine int
		for i, line := range lines {
			if strings.HasPrefix(line, "▶ ") {
				selectedLine = i
				break
			}
		}
		// Calculate the last line of the selected item
		// ie. With a 3 line item find the index of it's last line
		_, linesUsed := w.FindItemIndexOnLine(0, selectedLine)
		selectedLine = selectedLine + linesUsed
		// Find out how many lines we have available in the list view
		_, y0, _, y1 := v.Dimensions()
		availableLines := y1 - y0
		// Configure the top and bottom indexs for the view
		topIndex := 0
		bottomIndex := availableLines

		// Is the selected line off the bottom of the view?
		if selectedLine >= bottomIndex {
			// need to adjust down
			diff := selectedLine - bottomIndex + 1
			topIndex += diff
			bottomIndex += diff
		}

		// Is the selected line off the top of the view?
		if selectedLine < topIndex {
			// need to adjust up
			diff := topIndex - selectedLine
			topIndex -= diff
			bottomIndex -= diff
		}

		// Draw the lines which should be shown in the view based on the top and bottom index
		for index := topIndex; index < bottomIndex+1; index++ {
			if index < len(lines) {
				item := lines[index]
				fmt.Fprint(v, item+"\n")
			}
		}

		// If the title is getting too long trim things
		// down from the front
		title := w.currentPage.Title
		if w.currentPage.FilterString != "" {
			title += "[filter=" + w.currentPage.FilterString + "]"
		}
		if len(title) > width {
			trimLength := len(title) - width + 5 // Add five for spacing and elipsis
			title = ".." + title[trimLength:]
		}

		w.view.Title = title
	}

	return nil
}

// FindItemIndexOnLine finds index of the item at a particular line on the list
// eg. w.itemsToShow()[!indexHere!] it also returns how man lines are used by that item in the list
// if the item's display name has 3 new lines in it it'll linesUsed = 3
func (w *ListWidget) FindItemIndexOnLine(x, y int) (itemIndex, linesUsed int) {
	// Currently mouse is only bound to this view so don't have to worry about
	// the x coordinate being inside the view, just which item y is over
	for _, item := range w.hitBoxes {
		if y >= item.start && y <= item.end {
			return item.itemIndex, item.end - item.start
		}
	}
	return 0, 0
}

// MouseClick expands an item given it's x,y coordinates
// Todo: May not work with scrolled lists
func (w *ListWidget) MouseClick(x, y int) {
	itemIndex, _ := w.FindItemIndexOnLine(x, y)
	w.ChangeSelection(itemIndex)
	w.ExpandCurrentSelection()
}

// Refresh refreshes the current view
func (w *ListWidget) Refresh() {

	if w.isRefreshing {
		// already refreshing!
		return
	}
	w.refreshLock.Lock()
	if w.isRefreshing {
		return
	}
	// Mark refresh as in-progress and disable UI update
	w.isRefreshing = true
	w.SetShouldRender(false)

	currentExpandedItem := w.CurrentExpandedItem()
	currentSelection := w.CurrentSelection()
	done := w.statusView.Status("Refreshing", true)

	// capture current state
	sorted := false
	filterString := ""
	if w.currentPage != nil {
		if w.currentPage.Sorted {
			sorted = true
		}
		if w.currentPage.FilterString != "" {
			filterString = w.currentPage.FilterString
			w.ClearFilter() // clear filter so that `GoBack` actually navigates back
		}
	}

	w.GoBack()
	w.expandItem(currentExpandedItem)

	// wait for navigation before resetting previous selection
	go func() {
		defer errorhandling.RecoveryWithCleanup()
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")
		<-navigatedChannel
		eventing.Unsubscribe(navigatedChannel)

		w.ChangeSelection(currentSelection)

		w.ClearFilter()

		// reapply state
		if sorted {
			w.SortItems()
		}
		if filterString != "" {
			w.SetFilter(filterString)
		}

		// Clear isRefreshing and re-enable UI updates
		w.SetShouldRender(true)
		w.isRefreshing = false
		w.refreshLock.Unlock()
		done()
	}()
}

// GoBack takes the user back to preview view
func (w *ListWidget) GoBack() {
	eventing.Publish("list.prenavigate", "GOBACK")

	if w.currentPage == nil {
		return
	}
	if w.currentPage.FilterString != "" {
		// initial Back action is to clear filter, subsequent Back actions are normal
		w.ClearFilter()
		return
	}
	previousPage := w.navStack.Pop()
	if previousPage == nil {
		eventing.Publish("list.navigated", ListNavigatedEventState{Success: false})
		return
	}
	w.contentView.SetContentWithNode(previousPage.ExpandedNodeItem, previousPage.Data, previousPage.DataType, "Response")
	w.currentPage = previousPage

	if w.currentPage.ExpandedNodeItem == nil {
		w.currentPage.ExpandedNodeItem = &expanders.TreeNode{}
	}

	eventing.Publish("list.navigated", ListNavigatedEventState{
		Success:      true,
		NewNodes:     w.currentPage.Items,
		ParentNodeID: w.currentPage.ExpandedNodeItem.Parentid,
		IsBack:       true,
	})
}

// ExpandCurrentSelection opens the resource Sub->RG for example
func (w *ListWidget) ExpandCurrentSelection() {
	w.expandItem(w.CurrentItem())
}

// expandItem opens the specified resource Sub->RG for example
func (w *ListWidget) expandItem(item *expanders.TreeNode) {
	if w.isNavigating {
		// Skip if a navigation is already in progress
		return
	}
	w.navLock.Lock()
	if w.isNavigating { //double-check pattern
		// Skip if a navigation is already in progress
		return
	}
	w.isNavigating = true

	suppressPreviousTitle := false
	if w.currentPage != nil && w.currentPage.Title == "Subscriptions" {
		suppressPreviousTitle = true
	}

	newTitle := item.Name

	eventing.Publish("list.prenavigate", item.ID)

	go func() {
		defer func() {
			w.isNavigating = false
			w.navLock.Unlock()
		}()
		newContent, newItems, err := expanders.ExpandItem(w.ctx, item)
		if err != nil { // Don't need to display error as expander emits status event on error
			// Set parameters to trigger non-successful `list.navigated` event
			newItems = []*expanders.TreeNode{}
			newContent = nil
			newTitle = ""
		}
		w.Navigate(newItems, newContent, newTitle, suppressPreviousTitle)

		// Force UI to re-render to pickup
		w.g.Update(func(g *gocui.Gui) error {
			return nil
		})

	}()
}

// Navigate updates the currently selected list nodes, title and details content
func (w *ListWidget) Navigate(nodes []*expanders.TreeNode, content *expanders.ExpanderResponse, title string, suppressPreviousTitle bool) {

	titlePrefix := ""
	if !suppressPreviousTitle && w.currentPage != nil {
		titlePrefix = w.currentPage.Title + ">"
	}
	if len(nodes) == 0 && content == nil && title == "" {
		eventing.Publish("list.navigated", ListNavigatedEventState{Success: false})
	}
	currentItem := w.CurrentItem()

	if len(nodes) > 0 {

		if currentItem != nil && currentItem.ExpandInPlace {
			w.replaceNode(currentItem, nodes)
		} else {
			// Saves to current page to the nav stack and creates new current page.
			w.SetNewNodes(nodes)
		}

		// Build out new current page item
		w.currentPage.DataType = content.ResponseType
		w.currentPage.Data = content.Response
		if currentItem != nil {
			w.currentPage.Title = titlePrefix + currentItem.Name
			w.currentPage.ExpandedNodeItem = currentItem
		}
	}
	if content != nil {
		w.contentView.SetContentWithNode(currentItem, content.Response, content.ResponseType, title)
	}

	parentNodeID := "root"
	nodeID := "root"
	if w.currentPage != nil && w.currentPage.ExpandedNodeItem != nil {
		parentNodeID = w.currentPage.ExpandedNodeItem.ID
		nodeID = currentItem.ID
	}

	eventing.Publish("list.navigated", ListNavigatedEventState{
		Success:      true,
		NewNodes:     nodes,
		ParentNodeID: parentNodeID,
		NodeID:       nodeID,
	})
}

// GetNodes returns the currently listed nodes
func (w *ListWidget) GetNodes() []*expanders.TreeNode {
	if w.currentPage == nil {
		return []*expanders.TreeNode{}
	}
	return w.currentPage.Items
}

func (w *ListWidget) replaceNode(nodeToReplace *expanders.TreeNode, nodes []*expanders.TreeNode) {
	existingNodes := w.GetNodes()
	if existingNodes[len(existingNodes)-1] != nodeToReplace {
		panic("replaceNode can only be used on the last node in the list")
	}
	existingNodes = existingNodes[0 : len(existingNodes)-1]
	existingNodes = append(existingNodes, nodes...)

	w.currentPage.Items = existingNodes
}

// SetNewNodes allows others to set the list nodes
func (w *ListWidget) SetNewNodes(nodes []*expanders.TreeNode) {

	// Capture current view to navstack
	if w.HasCurrentItem() {
		w.navStack.Push(w.currentPage)

		currentID := w.CurrentItem().ID
		for _, node := range nodes {
			if node.ID == currentID {
				panic(fmt.Errorf("IDs must be unique or the navigate command breaks. ID: %s", currentID))
			}
		}
	}
	w.currentPage = &Page{
		Selection: 0,
		Items:     nodes,
	}
	w.ClearFilter()
}

// ChangeSelection updates the selected item
func (w *ListWidget) ChangeSelection(i int) {
	if i >= w.itemCount() {
		i = w.itemCount() - 1
	} else if i < 0 {
		i = 0
	}
	if w.currentPage != nil {
		w.currentPage.Selection = i
	}
}

// CurrentSelection returns the current selection int
func (w *ListWidget) CurrentSelection() int {
	if w.currentPage == nil {
		return -1
	}
	return w.currentPage.Selection
}

// HasCurrentItem indicates whether there is a current item
func (w *ListWidget) HasCurrentItem() bool {
	if w.currentPage == nil {
		return false
	}
	return w.currentPage.Selection < len(w.currentPage.Items)
}

// CurrentItem returns the selected item as a treenode
func (w *ListWidget) CurrentItem() *expanders.TreeNode {
	if w.HasCurrentItem() {
		return w.itemsToShow()[w.currentPage.Selection]
	}
	return nil
}

// CurrentExpandedItem returns the currently expanded item as a treenode
func (w *ListWidget) CurrentExpandedItem() *expanders.TreeNode {
	if w.currentPage == nil {
		return nil
	}
	return w.currentPage.ExpandedNodeItem
}

// MovePageDown changes the selection to be a page further down the list
func (w *ListWidget) MovePageDown() {
	if w.currentPage == nil {
		return
	}
	i := w.currentPage.Selection

	for remainingLinesToPage := w.lastCalculatedHeight; remainingLinesToPage > 0 && i < w.itemCount(); i++ {
		item := w.itemsToShow()[i]
		remainingLinesToPage -= strings.Count(item.Display, "\n") + 1 // +1 as there is an implicit newline
		remainingLinesToPage--                                        // separator
	}

	w.ChangeSelection(i)
}

// MovePageUp changes the selection to be a page further up the list
func (w *ListWidget) MovePageUp() {
	if w.currentPage == nil {
		return
	}
	i := w.currentPage.Selection

	for remainingLinesToPage := w.lastCalculatedHeight; remainingLinesToPage > 0 && i >= 0; i-- {
		item := w.itemsToShow()[i]
		remainingLinesToPage -= strings.Count(item.Display, "\n") + 1 // +1 as there is an implicit newline
		remainingLinesToPage--                                        // separator
	}

	w.ChangeSelection(i)
}

// MoveHome changes the selection to the top of the list
func (w *ListWidget) MoveHome() {
	if w.isNavigating {
		return
	}
	w.ChangeSelection(0)
}

// MoveEnd changes the selection to the bottom of the list
func (w *ListWidget) MoveEnd() {
	if w.isNavigating {
		return
	}
	w.ChangeSelection(w.itemCount() - 1)
}

// MoveUp moves the selection up one item
func (w *ListWidget) MoveUp() {
	if w.isNavigating {
		return
	}
	w.ChangeSelection(w.CurrentSelection() - 1)
}

// MoveDown moves the selection down one item
func (w *ListWidget) MoveDown() {
	if w.isNavigating {
		return
	}
	w.ChangeSelection(w.CurrentSelection() + 1)
}

// SetShouldRender sets the shouldRender value for the list and the current item
func (w *ListWidget) SetShouldRender(val bool) {
	w.shouldRender = val
	w.contentView.SetShouldRender(val)
}

// SortItems sorts the current list items by Name
func (w *ListWidget) SortItems() {
	if w.currentPage == nil {
		return
	}
	getSortName := func(itemName string) string {
		return strings.ToLower(itemName)
	}
	sortFunc := func(i, j int) bool {
		iValue := getSortName(w.currentPage.Items[i].Name)
		jValue := getSortName(w.currentPage.Items[j].Name)
		return iValue < jValue
	}

	sort.Slice(w.currentPage.Items, sortFunc)
	w.currentPage.Sorted = true
}
