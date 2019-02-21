package keybindings

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listBackId = 8

type ListBackHandler struct {
	List *views.ListWidget
}

func (h ListBackHandler) Id() string {
	return HandlerIds[listBackId]
}

func (h ListBackHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

func (h ListBackHandler) Widget() string {
	return "listWidget"
}

func (h ListBackHandler) DefaultKey() gocui.Key {
	return gocui.KeyBackspace2
}
