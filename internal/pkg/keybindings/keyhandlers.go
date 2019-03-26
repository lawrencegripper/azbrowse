package keybindings

import "github.com/jroimartin/gocui"

// HandlerIds A list of available handlers and their associated array index
var HandlerIds = []string{
	"quit",           // 0
	"copy",           // 1
	"listdelete",     // 2
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
	"listupdate",     // 17
	"listpagedown",   // 18
	"listpageup",     // 19
	"listend",        // 20
	"listhome",       // 21
}

// KeyHandler is an interface that all key handlers must implement
type KeyHandler interface {
	ID() string
	Fn() func(g *gocui.Gui, v *gocui.View) error
	Widget() string

	DefaultKey() gocui.Key
}

// KeyHandlerBase A base structure that will return the associated handler id from
// the HandlersId array and a default key for the handler.
type KeyHandlerBase struct {
	Index uint16
}

// ID returns the name of this item for example "listup"
func (h KeyHandlerBase) ID() string {
	return HandlerIds[h.Index]
}

// DefaultKey returns the default key mapped to the handler
func (h KeyHandlerBase) DefaultKey() gocui.Key {
	return DefaultKeys[h.ID()]
}

// ListHandler is a parent struct for all key handlers tied to the
// list widget view
type ListHandler struct {
	KeyHandlerBase
}

// Widget returns the name of the widget this handler binds to
func (h ListHandler) Widget() string {
	return "listWidget"
}

// ItemHandler is a parent struct for all key handlers tied to the
// item widget view
type ItemHandler struct {
	KeyHandlerBase
}

// Widget returns the name of the widget this handler binds to
func (h ItemHandler) Widget() string {
	return "itemWidget"
}

// GlobalHandler is a parent struct for all key handlers not tied to
// a specific view.
type GlobalHandler struct {
	KeyHandlerBase
}

// Widget returns the name of the widget this handler binds to
func (h GlobalHandler) Widget() string {
	return ""
}
