package keybindings

import "github.com/jroimartin/gocui"

type KeyHandler interface {
	Id() string
	Fn() func(g *gocui.Gui, v *gocui.View) error
	Widget() string
	DefaultKey() gocui.Key
}

var HandlerIds = []string{
	"quit",           // 0
	"copy",           // 1
	"delete",         // 2
	"fullscreen",     // 3
	"help",           // 4
	"itemback",       // 5
	"itemleft",       // 6
	"listactions",    // 7
	"listback",       // 8
	"listbacklegacy", // 9
	"listdown",       // 10
	"listup",         // 11
	"listright",      // 12
	"listedit",       // 13
	"listexpand",     // 14
	"listopen",       // 15
	"listrefresh",    // 16
	"save",           // 17
}
