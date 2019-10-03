package keybindings

import "github.com/jroimartin/gocui"

// HandlerID is used as the ID for Key Handlers
type HandlerID string

const (
	HandlerIDQuit                HandlerID = "quit"                //nolint:golint
	HandlerIDCopy                HandlerID = "copy"                //nolint:golint
	HandlerIDListDelete          HandlerID = "listdelete"          //nolint:golint
	HandlerIDFullScreen          HandlerID = "fullscreen"          //nolint:golint
	HandlerIDHelp                HandlerID = "help"                //nolint:golint
	HandlerIDItemBack            HandlerID = "itemback"            //nolint:golint
	HandlerIDItemLeft            HandlerID = "itemleft"            //nolint:golint
	HandlerIDListActions         HandlerID = "listactions"         //nolint:golint
	HandlerIDListBack            HandlerID = "listback"            //nolint:golint
	HandlerIDListBackLegacy      HandlerID = "listbacklegacy"      //nolint:golint
	HandlerIDListDown            HandlerID = "listdown"            //nolint:golint
	HandlerIDListUp              HandlerID = "listup"              //nolint:golint
	HandlerIDListRight           HandlerID = "listright"           //nolint:golint
	HandlerIDListEdit            HandlerID = "listedit"            //nolint:golint
	HandlerIDListExpand          HandlerID = "listexpand"          //nolint:golint
	HandlerIDListOpen            HandlerID = "listopen"            //nolint:golint
	HandlerIDListRefresh         HandlerID = "listrefresh"         //nolint:golint
	HandlerIDListUpdate          HandlerID = "listupdate"          //nolint:golint
	HandlerIDListPageDown        HandlerID = "listpagedown"        //nolint:golint
	HandlerIDListPageUp          HandlerID = "listpageup"          //nolint:golint
	HandlerIDListEnd             HandlerID = "listend"             //nolint:golint
	HandlerIDListHome            HandlerID = "listhome"            //nolint:golint
	HandlerIDConfirmDelete       HandlerID = "confirmdelete"       //nolint:golint
	HandlerIDClearPendingDeletes HandlerID = "clearpendingdeletes" //nolint:golint
	HandlerIDItemPageDown        HandlerID = "itempagedown"        //nolint:golint
	HandlerIDItemPageUp          HandlerID = "itempageup"          //nolint:golint
	HandlerIDToggleCommandPanel  HandlerID = "commandpanel"        //nolint:golint
	HandlerIDFilter              HandlerID = "filter"              //nolint:golint
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
