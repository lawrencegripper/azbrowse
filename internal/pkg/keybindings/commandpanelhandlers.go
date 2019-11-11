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
		h.commandPanelWidget.Hide()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////

type CommandPanelDownHandler struct {
	CommandPanelHandler
	commandPanelWidget *views.CommandPanelWidget
}

func NewCommandPanelDownHandler(commandPanelWidget *views.CommandPanelWidget) *CommandPanelDownHandler {
	handler := &CommandPanelDownHandler{
		commandPanelWidget: commandPanelWidget,
	}
	handler.id = HandlerIDCommandPanelDown
	return handler
}

func (h *CommandPanelDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.commandPanelWidget.MoveDown()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////

type CommandPanelUpHandler struct {
	CommandPanelHandler
	commandPanelWidget *views.CommandPanelWidget
}

func NewCommandPanelUpHandler(commandPanelWidget *views.CommandPanelWidget) *CommandPanelDownHandler {
	handler := &CommandPanelDownHandler{
		commandPanelWidget: commandPanelWidget,
	}
	handler.id = HandlerIDCommandPanelUp
	return handler
}

func (h *CommandPanelUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.commandPanelWidget.MoveUp()
		return nil
	}
}
