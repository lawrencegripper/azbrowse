package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

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
