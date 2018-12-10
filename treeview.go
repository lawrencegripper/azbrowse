package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type ListWidget struct {
	x, y     int
	w, h     int
	items    []string
	selected int
}

func NewListWidget(x, y, w, h int, items []string, selected int) *ListWidget {
	return &ListWidget{x: x, y: y, w: w, h: h, items: items, selected: selected}
}

func (w *ListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView("listWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	for i, s := range w.items {
		if i == w.selected {
			fmt.Fprintf(v, "*  ")
		} else {
			fmt.Fprintf(v, "   ")
		}
		fmt.Fprintf(v, s+"\n")
	}
	return nil
}

func (w *ListWidget) ChangeSelection(i int) {
	if i >= len(w.items) || i < 0 {
		return
	}
	w.selected = i
}

func (w *ListWidget) CurrentSelection() int {
	return w.selected
}
