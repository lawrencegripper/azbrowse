package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listRefreshId = "ListRefresh"

type ListRefreshHandler struct {
	List *views.ListWidget
}

func (h ListRefreshHandler) Id() string {
	return listRefreshId
}

func (h ListRefreshHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.Refresh()
		return nil
	}
}

func (h ListRefreshHandler) Widget() string {
	return "listWidget"
}

func (h ListRefreshHandler) DefaultKey() gocui.Key {
	return gocui.KeyF5
}
