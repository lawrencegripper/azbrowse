package keybindings

import (
	"context"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listActionsId = 7

type ListActionsHandler struct {
	List    *views.ListWidget
	Context context.Context
}

func (h ListActionsHandler) Id() string {
	return HandlerIds[listActionsId]
}

func (h ListActionsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return views.LoadActionsView(h.Context, h.List)
	}
}

func (h ListActionsHandler) Widget() string {
	return "listWidget"
}

func (h ListActionsHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlA
}
