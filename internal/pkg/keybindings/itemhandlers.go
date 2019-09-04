package keybindings

import (
	"github.com/jroimartin/gocui"
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
	handler.Index = 5
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
	handler.Index = 6
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
	handler.Index = 24
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
	handler.Index = 25
	return handler
}

func (h ItemViewPageUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.PageUp()
		return nil
	}
}

////////////////////////////////////////////////////////////////////
