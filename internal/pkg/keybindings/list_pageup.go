package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListPageUpHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListPageUpHandler(list *views.ListWidget) *ListPageUpHandler {
	handler := &ListPageUpHandler{
		List: list,
	}
	handler.id = HandlerIDListPageUp
	return handler
}

func (h ListPageUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MovePageUp()
		return nil
	}
}
