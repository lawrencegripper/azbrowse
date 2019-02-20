package keyhandlers

import (
	"github.com/jroimartin/gocui"
)

const QuitId = "Quit"

type QuitHandler struct {
}

func (h QuitHandler) Id() string {
	return QuitId
}

func (h QuitHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}
}

func (h QuitHandler) Widget() string {
	return ""
}

func (h QuitHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlC
}
