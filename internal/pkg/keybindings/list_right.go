package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListRightHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListRightHandler(list *views.ListWidget, editModeEnabled *bool) *ListRightHandler {
	handler := &ListRightHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.id = HandlerIDListRight
	return handler
}

func (h ListRightHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = true
		g.SetCurrentView("itemWidget")
		return nil
	}
}
