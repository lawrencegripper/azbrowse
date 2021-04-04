package keybindings

import (
	"github.com/awesome-gocui/gocui"
)

type QuitHandler struct {
	GlobalHandler
}

func NewQuitHandler() *QuitHandler {
	handler := &QuitHandler{}
	handler.id = HandlerIDQuit
	return handler
}

func (h QuitHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}
}
