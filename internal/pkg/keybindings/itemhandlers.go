package keybindings

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

////////////////////////////////////////////////////////////////////
type ItemBackHandler struct {
	ItemHandler
	List *views.ListWidget
}

func NewItemBackHandler(list *views.ListWidget) *ItemBackHandler {
	handler := &ItemBackHandler{
		List: list,
	}
	handler.id = HandlerIDItemBack
	return handler
}

func (h ItemBackHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView("listWidget")
		g.Cursor = false
		h.List.GoBack()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ItemLeftHandler struct {
	ItemHandler
	EditModeEnabled *bool
}

func NewItemLeftHandler(editModeEnabled *bool) *ItemLeftHandler {
	handler := &ItemLeftHandler{
		EditModeEnabled: editModeEnabled,
	}
	handler.id = HandlerIDItemLeft
	return handler
}

func (h ItemLeftHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = false
		g.SetCurrentView("listWidget")
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ItemViewPageDownHandler struct {
	ItemHandler
	ItemView *views.ItemWidget
}

func NewItemViewPageDownHandler(itemView *views.ItemWidget) *ItemViewPageDownHandler {
	handler := &ItemViewPageDownHandler{
		ItemView: itemView,
	}
	handler.id = HandlerIDItemPageDown
	return handler
}

func (h ItemViewPageDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.PageDown()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ItemViewPageUpHandler struct {
	ItemHandler
	ItemView *views.ItemWidget
}

func NewItemViewPageUpHandler(itemView *views.ItemWidget) *ItemViewPageUpHandler {
	handler := &ItemViewPageUpHandler{
		ItemView: itemView,
	}
	handler.id = HandlerIDItemPageUp
	return handler
}

func (h ItemViewPageUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.PageUp()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ItemCopyItemIDHandler struct {
	ListHandler
	Item      *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &ItemCopyItemIDHandler{}

func NewItemCopyItemIDHandler(item *views.ItemWidget, statusBar *views.StatusbarWidget) *ItemCopyItemIDHandler {
	handler := &ItemCopyItemIDHandler{
		Item:      item,
		StatusBar: statusBar,
	}
	handler.id = HandlerIDListCopyItemID
	return handler
}

func (h ItemCopyItemIDHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ItemCopyItemIDHandler) DisplayText() string {
	return "Copy current resource ID"
}
func (h *ItemCopyItemIDHandler) IsEnabled() bool {
	return h.Item.GetNode() != nil
}
func (h *ItemCopyItemIDHandler) Invoke() error {
	item := h.Item.GetNode()
	if item != nil {
		if err := copyToClipboard(item.ID); err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to copy resource ID to clipboard: %s", err.Error()), false)
			return nil
		}
		h.StatusBar.Status("Current resource ID copied to clipboard", false)
		return nil
	}
	return nil
}

////////////////////////////////////////////////////////////////////
