package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListUpHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListUpHandler(list *views.ListWidget) *ListUpHandler {
	handler := &ListUpHandler{
		List: list,
	}
	handler.id = HandlerIDListUp
	return handler
}

func (h ListUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveUp()
		return nil
	}
}
