package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListPageDownHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListPageDownHandler(list *views.ListWidget) *ListPageDownHandler {
	handler := &ListPageDownHandler{
		List: list,
	}
	handler.id = HandlerIDListPageDown
	return handler
}

func (h ListPageDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MovePageDown()
		return nil
	}
}
