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
	ListNavigateDown string `json:"ListNavigateDown"`
	ListNavigateUp   string `json:"ListNavigateUp"`
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
