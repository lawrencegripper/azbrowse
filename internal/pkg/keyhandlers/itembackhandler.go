package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const itemBackId = "ItemBack"

type ItemBackHandler struct {
	List *views.ListWidget
}

func (h ItemBackHandler) Id() string {
	return listBackId
}

func (h ItemBackHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView("listWidget")
		g.Cursor = false
		h.List.GoBack()
		return nil
	}
}

func (h ItemBackHandler) Widget() string {
	return "itemWidget"
}

func (h ItemBackHandler) DefaultKey() gocui.Key {
	return gocui.KeyBackspace2
}
