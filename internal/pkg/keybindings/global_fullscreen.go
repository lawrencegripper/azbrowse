package keybindings

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

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
	handler.id = HandlerIDFullScreen
	return handler
}

func (h FullscreenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.IsFullscreen)
		h.IsFullscreen = &tmp // memory leak?
		if *h.IsFullscreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY, 0)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			keyBindings := GetKeyBindingsAsStrings()
			v.Title = fmt.Sprintf("JSON Response - Fullscreen (%s to exit)", strings.ToUpper(strings.Join(keyBindings["fullscreen"], ",")))

			content := h.Content.GetContent()
			fmt.Fprint(v, style.ColorJSON(content))

			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}
