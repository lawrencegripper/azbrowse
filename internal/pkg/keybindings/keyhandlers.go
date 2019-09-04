package keybindings

import "github.com/jroimartin/gocui"

type HandlerID string

const (
	HandlerID_Quit                HandlerID = "quit"
	HandlerID_Copy                HandlerID = "copy"
	HandlerID_ListDelete          HandlerID = "listdelete"
	HandlerID_FullScreen          HandlerID = "fullscreen"
	HandlerID_Help                HandlerID = "help"
	HandlerID_ItemBack            HandlerID = "itemback"
	HandlerID_ItemLeft            HandlerID = "itemleft"
	HandlerID_ListActions         HandlerID = "listactions"
	HandlerID_ListBack            HandlerID = "listback"
	HandlerID_ListBackLegacy      HandlerID = "listbacklegacy"
	HandlerID_ListDown            HandlerID = "listdown"
	HandlerID_ListUp              HandlerID = "listup"
	HandlerID_ListRight           HandlerID = "listright"
	HandlerID_ListEdit            HandlerID = "listedit"
	HandlerID_ListExpand          HandlerID = "listexpand"
	HandlerID_ListOpen            HandlerID = "listopen"
	HandlerID_ListRefresh         HandlerID = "listrefresh"
	HandlerID_ListUpdate          HandlerID = "listupdate"
	HandlerID_ListPageDown        HandlerID = "listpagedown"
	HandlerID_ListPageUp          HandlerID = "listpageup"
	HandlerID_ListEnd             HandlerID = "listend"
	HandlerID_ListHome            HandlerID = "listhome"
	HandlerID_ConfirmDelete       HandlerID = "confirmdelete"
	HandlerID_ClearPendingDeletes HandlerID = "clearpendingdeletes"
	HandlerID_ItemPageDown        HandlerID = "itempagedown"
	HandlerID_ItemPageUp          HandlerID = "itempageup"
)

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
	tempID HandlerID
	xIndex uint16
}

// // ID returns the name of this item for example "listup"
// func (h KeyHandlerBase) ID() string {
// 	return HandlerIds[h.Index]
// }
// ID returns the name of this item for example "listup"
func (h KeyHandlerBase) ID() string {
	return string(h.tempID)
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
