package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListEditHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListEditHandler(list *views.ListWidget, editModeEnabled *bool) *ListEditHandler {
	handler := &ListEditHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.id = HandlerIDListEdit
	return handler
}

func (h ListEditHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.EditModeEnabled)
		h.EditModeEnabled = &tmp // memory leak?
		if *h.EditModeEnabled {
			g.Cursor = true
			g.SetCurrentView("itemWidget")
		} else {
			g.Cursor = false
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}
