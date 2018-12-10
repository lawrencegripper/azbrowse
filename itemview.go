package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type ItemWidget struct {
	x, y    int
	w, h    int
	Content string
}

func NewItemWidget(x, y, w, h int, content string) *ItemWidget {
	return &ItemWidget{x: x, y: y, w: w, h: h, Content: content}
}

func (w *ItemWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("itemWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	fmt.Fprint(v, w.Content)

	return nil
}
