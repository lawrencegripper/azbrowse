package keybindings

import (
	"context"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

////////////////////////////////////////////////////////////////////
type QuitHandler struct {
	GlobalHandler
}

func NewQuitHandler() *QuitHandler {
	handler := &QuitHandler{}
	handler.Index = 0
	return handler
}

func (h QuitHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type CopyHandler struct {
	GlobalHandler
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

func NewCopyHandler(content *views.ItemWidget, statusbar *views.StatusbarWidget) *CopyHandler {
	handler := &CopyHandler{
		Content:   content,
		StatusBar: statusbar,
	}
	handler.Index = 1
	return handler
}

func (h CopyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(h.Content.GetContent())
		h.StatusBar.Status("Current resource's JSON copied to clipboard", false)
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type DeleteHandler struct {
	GlobalHandler
	DeleteConfirmItemID string
	DeleteConfirmCount  int
	StatusBar           *views.StatusbarWidget
	Content             *views.ItemWidget
	List                *views.ListWidget
	Context             context.Context
}

func NewDeleteHandler(content *views.ItemWidget,
	statusbar *views.StatusbarWidget,
	list *views.ListWidget,
	deleteConfirmItemId string,
	deleteConfirmCount int,
	context context.Context) *DeleteHandler {
	handler := &DeleteHandler{
		Content:             content,
		StatusBar:           statusbar,
		List:                list,
		DeleteConfirmCount:  deleteConfirmCount,
		DeleteConfirmItemID: deleteConfirmItemId,
		Context:             context,
	}
	handler.Index = 2
	return handler
}

func (h DeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		if h.DeleteConfirmItemID != item.ID {
			h.DeleteConfirmItemID = item.ID
			h.DeleteConfirmCount = 0
		}
		done := h.StatusBar.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.DeleteURL, true)
		h.DeleteConfirmCount++

		if h.DeleteConfirmCount > 1 {
			done()
			doneDelete := h.StatusBar.Status("Deleting item: "+item.DeleteURL, true)

			res, err := armclient.DoRequest(h.Context, "DELETE", item.DeleteURL)
			if err != nil {
				panic(err)
			}
			h.List.Refresh()
			h.Content.SetContent(res, "Delete response>"+item.Name)
			doneDelete()
			h.DeleteConfirmItemID = ""
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type FullscreenHandler struct {
	GlobalHandler
	List         *views.ListWidget
	IsFullscreen *bool
	Content      *views.ItemWidget
}

func NewFullscreenHandler(list *views.ListWidget, content *views.ItemWidget, isFullscreen *bool) *FullscreenHandler {
	handler := &FullscreenHandler{
		List:         list,
		Content:      content,
		IsFullscreen: isFullscreen,
	}
	handler.Index = 3
	return handler
}

func (h FullscreenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.IsFullscreen)
		h.IsFullscreen = &tmp // memory leak?
		if *h.IsFullscreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			keyBindings := GetKeyBindingsAsStrings()
			v.Title = fmt.Sprintf("JSON Response - Fullscreen (%s to exit)", strings.ToUpper(keyBindings["fullscreen"]))
			fmt.Fprintf(v, h.Content.GetContent())
			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type HelpHandler struct {
	GlobalHandler
	ShowHelp *bool
}

func NewHelpHandler(showHelp *bool) *HelpHandler {
	handler := &HelpHandler{
		ShowHelp: showHelp,
	}
	handler.Index = 4
	return handler
}

func (h HelpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.ShowHelp)
		h.ShowHelp = &tmp // memory leak?

		// If we're up and running clear and redraw the view
		// if w.g != nil {
		if *h.ShowHelp {
			v, err := g.SetView("helppopup", 1, 1, 140, 38)
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			keyBindings := GetKeyBindingsAsStrings()
			views.DrawHelp(keyBindings, v)
		} else {
			g.DeleteView("helppopup")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////
