package keybindings

import "github.com/jroimartin/gocui"

// A list of available handlers and their associated array index
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
}

// The KeyHandler interface that all key handlers must implement
type KeyHandler interface {
	Id() string
	Fn() func(g *gocui.Gui, v *gocui.View) error
	Widget() string
	DefaultKey() gocui.Key
}

// A base structure that will return the associated handler id from
// the HandlersId array and a default key for the handler.
type KeyHandlerBase struct {
	Index uint16
}

func (h KeyHandlerBase) Id() string {
	return HandlerIds[h.Index]
}

func (h KeyHandlerBase) DefaultKey() gocui.Key {
	return DefaultKeys[h.Id()]
}

// List handler is a parent struct for all key handlers tied to the
// list widget view
type ListHandler struct {
	KeyHandlerBase
}

func (h ListHandler) Widget() string {
	return "listWidget"
}

// Item handler is a parent struct for all key handlers tied to the
// item widget view
type ItemHandler struct {
	KeyHandlerBase
}

func (h ItemHandler) Widget() string {
	return "itemWidget"
}

// Global handler is a parent struct for all key handlers not tied to
// a specific view.
type GlobalHandler struct {
	KeyHandlerBase
}

func (h GlobalHandler) Widget() string {
	return ""
}
