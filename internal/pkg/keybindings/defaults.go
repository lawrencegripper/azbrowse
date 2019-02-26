package keybindings

import "github.com/jroimartin/gocui"

// Theses are the default key bindings for each handler.
var DefaultKeys = map[string]gocui.Key{
	"quit":           gocui.KeyCtrlC,
	"copy":           gocui.KeyCtrlS,
	"listdelete":     gocui.KeyDelete,
	"fullscreen":     gocui.KeyCtrlF,
	"help":           gocui.KeyCtrlI,
	"itemback":       gocui.KeyBackspace2,
	"itemleft":       gocui.KeyArrowLeft,
	"listactions":    gocui.KeyCtrlA,
	"listback":       gocui.KeyBackspace2,
	"listbacklegacy": gocui.KeyBackspace,
	"listdown":       gocui.KeyArrowDown,
	"listup":         gocui.KeyArrowUp,
	"listright":      gocui.KeyArrowRight,
	"listedit":       gocui.KeyCtrlE,
	"listexpand":     gocui.KeyEnter,
	"listopen":       gocui.KeyCtrlO,
	"listrefresh":    gocui.KeyF5,
	"listupdate":     gocui.KeyCtrlU,
}
