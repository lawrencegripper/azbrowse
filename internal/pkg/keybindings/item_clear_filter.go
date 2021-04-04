package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ItemClearFilterHandler struct {
	ItemHandler
	ItemView *views.ItemWidget
}

func NewItemClearFilterHandler(itemView *views.ItemWidget) *ItemClearFilterHandler {
	handler := &ItemClearFilterHandler{
		ItemView: itemView,
	}
	handler.id = HandlerIDListClearFilter
	return handler
}

func (h ItemClearFilterHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.ItemView.ClearFilter()
		return nil
	}
}
