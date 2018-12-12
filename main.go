package main

import (
	"encoding/json"
	"github.com/lawrencegripper/azbrowse/style"
	"log"
	"os"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/armclient"
	open "github.com/skratchdot/open-golang/open"
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

	status := NewStatusbarWidget(1, maxY-2, maxX, g)
	header := NewHeaderWidget(1, 1, 70, 8)
	content := NewItemWidget(70+2, 1, ((maxX/4)*3)-3, maxY-4, "")
	list := NewListWidget(1, 10, 70, maxY-13, []string{"Loading..."}, 0, content, status)

	g.SetManager(status, content, list, header)

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

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ExpandCurrentSelection()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyF2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		item := list.CurrentItem()
		protalURL := os.Getenv("AZURE_PORTAL_URL")
		if protalURL == "" {
			protalURL = "https://portal.azure.com"
		}
		open.Run(protalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.parentid + "/overview")
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// HACK: To prevent accidental deletes this method requires del to be pressed twice on a resource
	// before it will proceed
	var deleteConfirmItemID string
	var deleteConfirmCount int
	if err := g.SetKeybinding("", gocui.KeyDelete, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		item := list.CurrentItem()
		if deleteConfirmItemID != item.id {
			deleteConfirmItemID = item.id
			deleteConfirmCount = 0
		}
		status.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.deleteURL, true)
		deleteConfirmCount++

		if deleteConfirmCount > 1 {
			status.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.deleteURL, true)

			res, err := armclient.DoRequest("DELETE", item.deleteURL)
			if err != nil {
				panic(err)
			}
			content.Content = style.Title("Delete response for item:"+item.deleteURL+"\n ------------------------------- \n") + res
			status.Status("Delete request sent successfully: "+item.deleteURL, false)

			deleteConfirmItemID = ""

		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	go func() {
		time.Sleep(time.Second * 1)

		status.Status("Fetching Subscriptions", true)

		// Get Subscriptions
		data, err := armclient.DoRequest("GET", "/subscriptions?api-version=2018-01-01")
		if err != nil {
			panic(err)
		}

		var subRequest armclient.SubResponse
		err = json.Unmarshal([]byte(data), &subRequest)
		if err != nil {
			panic(err)
		}

		g.Update(func(gui *gocui.Gui) error {

			list.SetSubscriptions(subRequest)

			if err != nil {
				content.Content = err.Error()
				return nil
			}
			content.Content = data
			return nil
		})

		status.Status("Fetching Subscriptions: Completed", false)

	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
