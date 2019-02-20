package keyhandlers

import "github.com/jroimartin/gocui"

type KeyHandler interface {
	Id() string
	Fn() func(g *gocui.Gui, v *gocui.View) error
	Widget() string
	DefaultKey() gocui.Key
}
