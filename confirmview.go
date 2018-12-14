package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/style"
)

// ConfirmWidget is response for showing the text response from the Rest requests
type ConfirmWidget struct {
	Content string
	Action  func() error
	view    *gocui.View
}

// Layout draws the widget in the gocui view
func (w *ConfirmWidget) Layout(g *gocui.Gui) error {
	// maxX, maxY := g.Size()
	v, err := g.SetView("ConfirmWidget", 10, 14, 80, 22)
	w.view = v
	g.SetCurrentView("ConfirmWidget")

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	g.SetKeybinding("ConfirmWidget", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.Update(func(g *gocui.Gui) error {
			w.Content = "Update in progress..... please wait"
			err := w.Action()
			if err != nil {
				panic(err)
			}
			g.SetCurrentView("listWidget")

			return g.DeleteView("ConfirmWidget")
		})
		return nil
	})

	g.SetKeybinding("ConfirmWidget", gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.Update(func(g *gocui.Gui) error {
			g.SetCurrentView("listWidget")
			return g.DeleteView("ConfirmWidget")
		})
		return nil
	})

	fmt.Fprint(v, style.Header(w.Content))
	fmt.Fprint(v)

	return nil
}
