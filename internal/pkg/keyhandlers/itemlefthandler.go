package keyhandlers

import (
	"github.com/jroimartin/gocui"
)

const listLeftId = "ItemLeft"

type ItemLeftHandler struct {
	EditModeEnabled *bool
}

func (h ItemLeftHandler) Id() string {
	return listLeftId
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
