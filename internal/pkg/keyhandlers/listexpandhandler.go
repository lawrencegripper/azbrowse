package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listExpandId = 14

type ListExpandHandler struct {
	List *views.ListWidget
}

func (h ListExpandHandler) Id() string {
	return HandlerIds[listExpandId]
}

func (h ListExpandHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ExpandCurrentSelection()
		return nil
	}
}

func (h ListExpandHandler) Widget() string {
	return "listWidget"
}

func (h ListExpandHandler) DefaultKey() gocui.Key {
	return gocui.KeyEnter
}
