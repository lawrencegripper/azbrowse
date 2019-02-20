package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const helpId = 4

type HelpHandler struct {
	ShowHelp *bool
}

func (h HelpHandler) Id() string {
	return HandlerIds[helpId]
}

func (h HelpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.ShowHelp)
		h.ShowHelp = &tmp // memory leak?

		// If we're up and running clear and redraw the view
		// if w.g != nil {
		if *h.ShowHelp {
			v, err := g.SetView("helppopup", 1, 1, 140, 38)
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			views.DrawHelp(v)
		} else {
			g.DeleteView("helppopup")
		}
		return nil
	}
}

func (h HelpHandler) Widget() string {
	return ""
}

func (h HelpHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlI
}
