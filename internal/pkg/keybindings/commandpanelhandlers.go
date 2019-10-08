package keybindings

import (
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/stuartleeks/gocui"
)

type CloseCommandPanelHandler struct {
	CommandPanelHandler
	commandPanelWidget *views.CommandPanelWidget
}

func NewCloseCommandPanelHandler(commandPanelWidget *views.CommandPanelWidget) *CloseCommandPanelHandler {
	handler := &CloseCommandPanelHandler{
		commandPanelWidget: commandPanelWidget,
	}
	handler.id = HandlerIDToggleCloseCommandPanel
	return handler
}

func (h *CloseCommandPanelHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.commandPanelWidget.ToggleShowHide()
		return nil
	}
}
