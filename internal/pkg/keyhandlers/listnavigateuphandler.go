package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listNavigateUpId = "ListNavigateUp"

type ListNavigateUpHandler struct {
	List *views.ListWidget
}

func (h ListNavigateUpHandler) Id() string {
	return listNavigateUpId
}

func (h ListNavigateUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() - 1)
		return nil
	}
}

func (h ListNavigateUpHandler) Widget() string {
	return "listWidget"
}

func (h ListNavigateUpHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowUp
}
