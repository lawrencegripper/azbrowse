package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListHomeHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListHomeHandler(list *views.ListWidget) *ListHomeHandler {
	handler := &ListHomeHandler{
		List: list,
	}
	handler.id = HandlerIDListHome
	return handler
}

func (h ListHomeHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveHome()
		return nil
	}
}
