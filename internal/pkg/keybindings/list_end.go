package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListEndHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListEndHandler(list *views.ListWidget) *ListEndHandler {
	handler := &ListEndHandler{
		List: list,
	}
	handler.id = HandlerIDListEnd
	return handler
}

func (h ListEndHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveEnd()
		return nil
	}
}
