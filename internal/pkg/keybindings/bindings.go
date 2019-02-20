package keybindings

import (
	"os"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/keyhandlers"
)

type KeyMap map[string]gocui.Key

// SemanticKeyMap properties should always
// match the IDs of the associated handlers
type SemanticKeyMap struct {
	ListDown       string `json:"ListDown,omitempty,omitempty"`
	ListUp         string `json:"ListUp,omitempty"`
	ListBack       string `json:"ListBack,omitempty"`
	ListBackLegacy string `json:"ListBackLegacy,omitempty"`
	ListEdit       string `json:"ListEdit,omitempty"`
	ListExpand     string `json:"ListExpand,omitempty"`
	ListOpen       string `json:"ListOpen,omitempty"`
	ListRefresh    string `json:"ListRefresh,omitempty"`
	ListRight      string `json:"ListRight,omitempty"`
	ListActions    string `json:"ListActions,omitempty"`
	ItemBack       string `json:"ItemBack,omitempty"`
	ItemLeft       string `json:"ItemLeft,omitempty"`
	Help           string `json:"Help,omitempty"`
	Copy           string `json:"Copy,omitempty"`
	Fullscreen     string `json:"Fullscreen,omitempty"`
	Delete         string `json:"Delete,omitempty"`
	Quit           string `json:"Quit,omitempty"`
}

var handlers []keyhandlers.KeyHandler
var overrides KeyMap

func Bind(g *gocui.Gui) error {
	defaultFilePath := "bindings.json"
	return BindWithFileOverrides(g, defaultFilePath)
}

func BindWithFileOverrides(g *gocui.Gui, filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		err = initializeOverrides(filePath)
		if err != nil {
			return err
		}
	} // ignore file overrides if error
	if err := bindHandlersToKeys(g); err != nil {
		return err
	}
	return nil
}

func AddHandler(hnd keyhandlers.KeyHandler) {
	handlers = append(handlers, hnd)
}

func bindHandlersToKeys(g *gocui.Gui) error {
	for _, hnd := range handlers {
		if err := bindHandlerToKey(g, hnd); err != nil {
			return err
		}
	}
	return nil
}

func bindHandlerToKey(g *gocui.Gui, hnd keyhandlers.KeyHandler) error {
	var key gocui.Key
	if k, ok := overrides[hnd.Id()]; ok {
		key = k
	} else {
		key = hnd.DefaultKey()
	}
	return g.SetKeybinding(hnd.Widget(), key, gocui.ModNone, hnd.Fn())
}
