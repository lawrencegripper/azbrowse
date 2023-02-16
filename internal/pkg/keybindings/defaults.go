package keybindings

import (
	"github.com/awesome-gocui/gocui"
)

// DefaultKeys are the default key bindings for each handler.
var DefaultKeys = map[string]interface{}{
	"quit":                gocui.KeyCtrlC,
	"copy":                gocui.KeyCtrlS,
	"listdelete":          gocui.KeyDelete,
	"fullscreen":          gocui.KeyF11,
	"help":                gocui.KeyCtrlI,
	"itemback":            gocui.KeyBackspace2,
	"itemleft":            []interface{}{gocui.KeyArrowLeft, rune('h')},
	"listactions":         gocui.KeyCtrlA,
	"listback":            gocui.KeyBackspace2,
	"listbacklegacy":      gocui.KeyBackspace,
	"listdown":            []interface{}{gocui.KeyArrowDown, rune('j')},
	"listup":              []interface{}{gocui.KeyArrowUp, rune('k')},
	"listright":           []interface{}{gocui.KeyArrowRight, rune('l')},
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
	"itemclearfilter":     gocui.KeyEsc,
	"commandpanelopen":    gocui.KeyCtrlP,
	"commandpaneldown":    gocui.KeyArrowDown,
	"commandpanelup":      gocui.KeyArrowUp,
	"commandpanelenter":   gocui.KeyEnter,
	"filter":              rune('/'),
	"commandpanelclose":   gocui.KeyEsc,
	"azuresearchquery":    gocui.KeyCtrlR,
}
