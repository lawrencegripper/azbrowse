package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ItemViewPageUpHandler struct {
	ItemHandler
	ItemView *views.ItemWidget
}

func NewItemViewPageUpHandler(itemView *views.ItemWidget) *ItemViewPageUpHandler {
	handler := &ItemViewPageUpHandler{
		ItemView: itemView,
	}
	handler.id = HandlerIDItemPageUp
	return handler
}

func (h ItemViewPageUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.PageUp()
		return nil
	}
}
