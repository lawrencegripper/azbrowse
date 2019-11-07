package keybindings

import (
	"github.com/stuartleeks/gocui"
)

// DefaultKeys are the default key bindings for each handler.
var DefaultKeys = map[string]interface{}{
	"quit":                gocui.KeyCtrlC,
	"copy":                gocui.KeyCtrlS,
	"listdelete":          gocui.KeyDelete,
	"fullscreen":          gocui.KeyF11,
	"help":                gocui.KeyCtrlI,
	"itemback":            gocui.KeyBackspace2,
	"itemleft":            gocui.KeyArrowLeft,
	"listactions":         gocui.KeyCtrlA,
	"listback":            gocui.KeyBackspace2,
	"listbacklegacy":      gocui.KeyBackspace,
	"listdown":            gocui.KeyArrowDown,
	"listup":              gocui.KeyArrowUp,
	"listright":           gocui.KeyArrowRight,
	"listedit":            gocui.KeyCtrlE,
	"listexpand":          gocui.KeyEnter,
	"listopen":            gocui.KeyCtrlO,
	"listrefresh":         gocui.KeyF5,
	"listupdate":          gocui.KeyCtrlU,
	"listpagedown":        gocui.KeyPgdn,
	"listpageup":          gocui.KeyPgup,
	"listend":             gocui.KeyEnd,
	"listhome":            gocui.KeyHome,
	"listclearfilter":     gocui.KeyEsc,
	"confirmdelete":       gocui.KeyCtrlY,
	"clearpendingdeletes": gocui.KeyCtrlN,
	"itempagedown":        gocui.KeyPgdn,
	"itempageup":          gocui.KeyPgup,
	"commandpanelopen":    gocui.KeyCtrlP,
	"filter":              rune("/"[0]),
	"commandpanelclose":   gocui.KeyEsc,
	"azuresearchquery":    gocui.KeyCtrlSlash,
}
