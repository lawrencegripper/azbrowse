package main

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/style"
)

// StatusbarWidget controls the statusbar
type StatusbarWidget struct {
	name            string
	x, y            int
	w               int
	message         string
	messageAddition string
	loading         bool
}

// NewStatusbarWidget create new instance and start go routine for spinner
func NewStatusbarWidget(x, y, w int, g *gocui.Gui) *StatusbarWidget {
	widget := &StatusbarWidget{name: "statusBarWidget", x: x, y: y, w: w}
	// Start loop for showing loading in statusbar
	go func() {
		for {
			time.Sleep(time.Second)
			g.Update(func(gui *gocui.Gui) error {
				if widget.loading {
					widget.messageAddition = widget.messageAddition + "."
				} else {
					widget.messageAddition = ""
				}
				return nil
			})

		}
	}()
	return widget
}

// Layout draws the widget in the gocui view
func (w *StatusbarWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+3)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	v.Title = `Status [CTRL+I -> Help]`
	v.Wrap = true

	if w.loading {
		fmt.Fprint(v, style.Loading("⏳  "+w.message))
	} else {
		fmt.Fprint(v, style.Completed("✓ "+w.message))
	}
	fmt.Fprint(v, w.messageAddition)

	return nil
}

// Status updates the message in the status bar and whether to show loading indicator
func (w *StatusbarWidget) Status(message string, loading bool) {
	w.message = message
	w.loading = loading
}
