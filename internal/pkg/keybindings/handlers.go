package keybindings

import "github.com/jroimartin/gocui"

type Handler struct {
	Id     string
	Fn     func(g *gocui.Gui, v *gocui.View) error
	Widget string
}

var handlers []Handler

func RegisterHandler(hnd *Handler) {
	handlers = append(handlers, *hnd)
}

func BindHandlersToKeys(g *gocui.Gui, keyMap KeyMap) error {
	for _, hnd := range handlers {
		if err := bindHandlerToKey(g, keyMap[hnd.Id], hnd); err != nil {
			return err
		}
	}

	return nil
}

func bindHandlerToKey(g *gocui.Gui, key gocui.Key, hnd Handler) error {
	return g.SetKeybinding(hnd.Widget, key, gocui.ModNone, hnd.Fn)
}
