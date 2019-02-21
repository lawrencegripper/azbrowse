package keybindings

import (
	"github.com/jroimartin/gocui"
)

const itemLeftId = 6

type ItemLeftHandler struct {
	EditModeEnabled *bool
}

func (h ItemLeftHandler) Id() string {
	return HandlerIds[itemLeftId]
}

func (h ItemLeftHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = false
		g.SetCurrentView("listWidget")
		return nil
	}
}

func (h ItemLeftHandler) Widget() string {
	return "itemWidget"
}

func (h ItemLeftHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowLeft
}
