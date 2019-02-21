package keybindings

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listEditId = 13

type ListEditHandler struct {
	List            *views.ListWidget
	EditModeEnabled *bool
}

func (h ListEditHandler) Id() string {
	return HandlerIds[listEditId]
}

func (h ListEditHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.EditModeEnabled)
		h.EditModeEnabled = &tmp // memory leak?
		if *h.EditModeEnabled {
			g.Cursor = true
			g.SetCurrentView("itemWidget")
		} else {
			g.Cursor = false
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

func (h ListEditHandler) Widget() string {
	return "listWidget"
}

func (h ListEditHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlE
}
