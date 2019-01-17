package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// ItemWidget is response for showing the text response from the Rest requests
type ItemWidget struct {
	x, y    int
	w, h    int
	content string
	view    *gocui.View
	g       *gocui.Gui
}

// NewItemWidget creates a new instance of ItemWidget
func NewItemWidget(x, y, w, h int, content string) *ItemWidget {
	return &ItemWidget{x: x, y: y, w: w, h: h, content: content}
}

// Layout draws the widget in the gocui view
func (w *ItemWidget) Layout(g *gocui.Gui) error {
	w.g = g
	v, err := g.SetView("itemWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Wrap = true

	w.view = v
	v.Clear()

	fmt.Fprint(v, w.content)

	return nil
}

// SetContent displays the string in the itemview
func (w *ItemWidget) SetContent(content string, title string) {
	w.g.Update(func(g *gocui.Gui) error {
		w.content = content
		// Reset the cursor and origin (scroll poisition)
		// so we don't start at the bottom of a new doc
		w.view.SetCursor(0, 0)
		w.view.SetOrigin(0, 0)
		w.view.Title = title
		return nil
	})
}

// GetContent returns the current content
func (w *ItemWidget) GetContent() string {
	return w.content
}
