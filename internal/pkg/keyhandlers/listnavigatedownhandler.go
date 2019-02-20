package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listNavigateDownId = "ListNavigateDown"

type ListNavigateDownHandler struct {
	List *views.ListWidget
}

func (h ListNavigateDownHandler) Id() string {
	return listNavigateDownId
}

func (h ListNavigateDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() + 1)
		return nil
	}
}

func (h ListNavigateDownHandler) Widget() string {
	return "listWidget"
}

func (h ListNavigateDownHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowDown
}
