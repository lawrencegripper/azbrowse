package keybindings

import (
	"encoding/json"
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListDebugCopyItemDataHandler struct {
	ListHandler
	List      *views.ListWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &ListDebugCopyItemDataHandler{}

func NewListDebugCopyItemDataHandler(list *views.ListWidget, statusBar *views.StatusbarWidget) *ListDebugCopyItemDataHandler {
	handler := &ListDebugCopyItemDataHandler{
		List:      list,
		StatusBar: statusBar,
	}
	handler.id = HandlerIDListDebugCopyItemData
	return handler
}

func (h ListDebugCopyItemDataHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListDebugCopyItemDataHandler) DisplayText() string {
	return "DEBUG: Copy current node data"
}
func (h *ListDebugCopyItemDataHandler) IsEnabled() bool {
	return h.List.CurrentExpandedItem() != nil
}
func (h *ListDebugCopyItemDataHandler) Invoke() error {
	item := h.List.CurrentExpandedItem()
	if item != nil {
		buf, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to generate node summary: %s", err.Error()), false)
			return nil
		}
		if err := copyToClipboard(string(buf)); err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to copy node data to clipboard: %s", err.Error()), false)
			return nil
		}
		h.StatusBar.Status("Current node data copied to clipboard", false)
		return nil
	}
	return nil
}
