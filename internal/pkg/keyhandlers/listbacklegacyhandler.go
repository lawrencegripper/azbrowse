package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

// Handle backspace for out-dated terminals
// A side-effect is that this key combination clashes with CTRL+H so we can't use that combination for help... oh well.
// https://superuser.com/questions/375864/ctrlh-causing-backspace-instead-of-help-in-emacs-on-cygwin

const listBackLegacyId = 9

type ListBackLegacyHandler struct {
	List *views.ListWidget
}

func (h ListBackLegacyHandler) Id() string {
	return HandlerIds[listBackLegacyId]
}

func (h ListBackLegacyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

func (h ListBackLegacyHandler) Widget() string {
	return "listWidget"
}

func (h ListBackLegacyHandler) DefaultKey() gocui.Key {
	return gocui.KeyBackspace
}
