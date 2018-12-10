package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"time"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	// help := NewHelpWidget("help", 1, 1, helpText)
	maxX, maxY := g.Size()
	// Padding
	maxX = maxX - 2
	maxY = maxY - 2

	status := NewStatusbarWidget("status", 1, maxY-2, maxX)
	list := NewListWidget(1, 1, maxX/4, maxY-4, []string{"thing", "another", "somemore", "stuff"}, 0)
	contentStart := maxX / 4
	content := NewItemWidget(contentStart+4, 1, ((maxX/4)*3)-3, maxY-4, "This is a thing")

	g.SetManager(status, list, content)
	statusSet(status, 0.1)

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ChangeSelection(list.CurrentSelection() + 1)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ChangeSelection(list.CurrentSelection() - 1)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	go func() {
		time.Sleep(time.Second * 1)
		g.Update(func(gui *gocui.Gui) error {
			statusSet(status, 0.9)
			return nil
		})
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
