package keybindings

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ToggleDemoModeHandler struct {
	GlobalHandler
	Settings *config.Settings
	List     *views.ListWidget
	Status   *views.StatusbarWidget
	Content  *views.ItemWidget
}

var _ Command = &ToggleDemoModeHandler{}

func NewToggleDemoModeHandler(settings *config.Settings, list *views.ListWidget, status *views.StatusbarWidget, content *views.ItemWidget) *ToggleDemoModeHandler {
	handler := &ToggleDemoModeHandler{
		Settings: settings,
		List:     list,
		Status:   status,
		Content:  content,
	}
	handler.id = HandlerIDToggleDemoMode
	return handler
}

func (h ToggleDemoModeHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ToggleDemoModeHandler) DisplayText() string {
	status := "off"
	if h.Settings.HideGuids {
		status = "on"
	}
	return fmt.Sprintf("Toggle Demo mode (currently %s)", status)
}

func (h *ToggleDemoModeHandler) IsEnabled() bool {
	return true
}

func (h *ToggleDemoModeHandler) Invoke() error {
	h.Settings.HideGuids = !h.Settings.HideGuids
	h.Status.SetHideGuids(h.Settings.HideGuids)
	h.Content.SetHideGuids(h.Settings.HideGuids)
	h.List.Refresh()
	return nil
}
