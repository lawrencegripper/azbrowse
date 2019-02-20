package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listDownId = "ListDown"

type ListDownHandler struct {
	List *views.ListWidget
}

func (h ListDownHandler) Id() string {
	return listDownId
}

func (h ListDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() + 1)
		return nil
	}
}

func (h ListDownHandler) Widget() string {
	return "listWidget"
}

func (h ListDownHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowDown
}
