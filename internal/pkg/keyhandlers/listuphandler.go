package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listUpId = 11

type ListUpHandler struct {
	List *views.ListWidget
}

func (h ListUpHandler) Id() string {
	return HandlerIds[listUpId]
}

func (h ListUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() - 1)
		return nil
	}
}

func (h ListUpHandler) Widget() string {
	return "listWidget"
}

func (h ListUpHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowUp
}
