package keybindings

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ItemCopyItemIDHandler struct {
	ListHandler
	Item      *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &ItemCopyItemIDHandler{}

func NewItemCopyItemIDHandler(item *views.ItemWidget, statusBar *views.StatusbarWidget) *ItemCopyItemIDHandler {
	handler := &ItemCopyItemIDHandler{
		Item:      item,
		StatusBar: statusBar,
	}
	handler.id = HandlerIDListCopyItemID
	return handler
}

func (h ItemCopyItemIDHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ItemCopyItemIDHandler) DisplayText() string {
	return "Copy current resource ID"
}
func (h *ItemCopyItemIDHandler) IsEnabled() bool {
	return h.Item.GetNode() != nil
}
func (h *ItemCopyItemIDHandler) Invoke() error {
	item := h.Item.GetNode()
	if item != nil {
		if err := copyToClipboard(item.ID); err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to copy resource ID to clipboard: %s", err.Error()), false)
			return nil
		}
		h.StatusBar.Status("Current resource ID copied to clipboard", false)
		return nil
	}
	return nil
}
