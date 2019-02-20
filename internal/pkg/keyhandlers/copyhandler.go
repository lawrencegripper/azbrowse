package keyhandlers

import (
	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const CopyId = "Copy"

type CopyHandler struct {
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

func (h CopyHandler) Id() string {
	return CopyId
}

func (h CopyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(h.Content.GetContent())
		h.StatusBar.Status("Current resource's JSON copied to clipboard", false)
		return nil
	}
}

func (h CopyHandler) Widget() string {
	return ""
}

func (h CopyHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlS
}
