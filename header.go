package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type HeaderWidget struct {
	x, y int
	w, h int
}

func NewHeaderWidget(x, y, w, h int) *HeaderWidget {
	return &HeaderWidget{x: x, y: y, w: w, h: h}
}

func (w *HeaderWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("headerWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	fmt.Fprint(v, `
	
    _       ___                         ↑/↓:     Select resource
   /_\   __| _ )_ _ _____ __ _____ ___  ENTER:   Expand/View resource
  / _ \ |_ / _ \ '_/ _ \ V  V (_-</ -_) Backspace: Go back
 /_/ \_\/__|___/_| \___/\_/\_//__/\___| F5:      Refresh 
                                        
 Interactive CLI for browsing Azure resources`)

	return nil
}
