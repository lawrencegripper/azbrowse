package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/style"
)

// HeaderWidget controls the header for the cli interface
type HeaderWidget struct {
	x, y int
	w, h int
}

// NewHeaderWidget creates a new header instance
func NewHeaderWidget(x, y, w, h int) *HeaderWidget {
	return &HeaderWidget{x: x, y: y, w: w, h: h}
}

// Layout draws the widget in the gocui view
func (w *HeaderWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("headerWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	fmt.Fprint(v, style.Header(`
                                        CTLT+E:  Toggle browse JSON    
                                        CTLT+F:  Toggle Fullscreen                                        
    _       ___                         ↑/↓:     Select resource              
   /_\   __| _ )_ _ _____ __ _____ ___  ENTER:   Expand/View resource
  / _ \ |_ / _ \ '_/ _ \ V  V (_-</ -_) Backspace: Go back           
 /_/ \_\/__|___/_| \___/\_/\_//__/\___| CTRL+O:  Open Portal             
                                        DEL:     Delete resource                             
 Interactive CLI for browsing Azure resources                         
                                                                       `))

	return nil
}
