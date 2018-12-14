package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// ItemWidget is response for showing the text response from the Rest requests
type ItemWidget struct {
	x, y    int
	w, h    int
	Content string
	view    *gocui.View
}

// NewItemWidget creates a new instance of ItemWidget
func NewItemWidget(x, y, w, h int, content string) *ItemWidget {
	return &ItemWidget{x: x, y: y, w: w, h: h, Content: content}
}

// Layout draws the widget in the gocui view
func (w *ItemWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("itemWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	v.Editable = true
	// v.Wrap = true

	w.view = v
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	fmt.Fprint(v, w.Content)

	return nil
}
