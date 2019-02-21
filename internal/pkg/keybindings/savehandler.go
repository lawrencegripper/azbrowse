package keybindings

import (
	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const saveId = 17

type SaveHandler struct {
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

func (h SaveHandler) Id() string {
	return HandlerIds[saveId]
}

func (h SaveHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(h.Content.GetContent())
		h.StatusBar.Status("Current resource's JSON copied to clipboard", false)
		return nil
	}
}

func (h SaveHandler) Widget() string {
	return ""
}

func (h SaveHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlS
}
