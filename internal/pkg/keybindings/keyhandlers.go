package keybindings

import "github.com/jroimartin/gocui"

type HandlerID string

const (
	HandlerIDQuit                HandlerID = "quit"
	HandlerIDCopy                HandlerID = "copy"
	HandlerIDListDelete          HandlerID = "listdelete"
	HandlerIDFullScreen          HandlerID = "fullscreen"
	HandlerIDHelp                HandlerID = "help"
	HandlerIDItemBack            HandlerID = "itemback"
	HandlerIDItemLeft            HandlerID = "itemleft"
	HandlerIDListActions         HandlerID = "listactions"
	HandlerIDListBack            HandlerID = "listback"
	HandlerIDListBackLegacy      HandlerID = "listbacklegacy"
	HandlerIDListDown            HandlerID = "listdown"
	HandlerIDListUp              HandlerID = "listup"
	HandlerIDListRight           HandlerID = "listright"
	HandlerIDListEdit            HandlerID = "listedit"
	HandlerIDListExpand          HandlerID = "listexpand"
	HandlerIDListOpen            HandlerID = "listopen"
	HandlerIDListRefresh         HandlerID = "listrefresh"
	HandlerIDListUpdate          HandlerID = "listupdate"
	HandlerIDListPageDown        HandlerID = "listpagedown"
	HandlerIDListPageUp          HandlerID = "listpageup"
	HandlerIDListEnd             HandlerID = "listend"
	HandlerIDListHome            HandlerID = "listhome"
	HandlerIDConfirmDelete       HandlerID = "confirmdelete"
	HandlerIDClearPendingDeletes HandlerID = "clearpendingdeletes"
	HandlerIDItemPageDown        HandlerID = "itempagedown"
	HandlerIDItemPageUp          HandlerID = "itempageup"
	HandlerIDListLetterS         HandlerID = "listletter-s"
)

// KeyHandler is an interface that all key handlers must implement
type KeyHandler interface {
	ID() string
	Fn() func(g *gocui.Gui, v *gocui.View) error
	Widget() string

	DefaultKey() interface{}
}

// KeyHandlerBase A base structure that will return the associated handler id from
// the HandlersId array and a default key for the handler.
type KeyHandlerBase struct {
	id HandlerID
}

// // ID returns the name of this item for example "listup"
// func (h KeyHandlerBase) ID() string {
// 	return HandlerIds[h.Index]
// }
// ID returns the name of this item for example "listup"
func (h KeyHandlerBase) ID() string {
	return string(h.id)
}

// DefaultKey returns the default key mapped to the handler
func (h KeyHandlerBase) DefaultKey() interface{} {
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
