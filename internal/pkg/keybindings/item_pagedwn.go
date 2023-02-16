package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ItemViewPageDownHandler struct {
	ItemHandler
	ItemView *views.ItemWidget
}

func NewItemViewPageDownHandler(itemView *views.ItemWidget) *ItemViewPageDownHandler {
	handler := &ItemViewPageDownHandler{
		ItemView: itemView,
	}
	handler.id = HandlerIDItemPageDown
	return handler
}

func (h ItemViewPageDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.PageDown()
		return nil
	}
}
