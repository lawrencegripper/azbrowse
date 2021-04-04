package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListClearFilterHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListClearFilterHandler(list *views.ListWidget) *ListClearFilterHandler {
	handler := &ListClearFilterHandler{
		List: list,
	}
	handler.id = HandlerIDListClearFilter
	return handler
}

func (h ListClearFilterHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ClearFilter()
		return nil
	}
}
