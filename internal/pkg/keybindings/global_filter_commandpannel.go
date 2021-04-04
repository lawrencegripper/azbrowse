package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type CommandPanelFilterHandler struct {
	GlobalHandler
	commandPanelWidget *views.CommandPanelWidget
	list               *views.ListWidget
}

var _ Command = &CommandPanelFilterHandler{}

func NewCommandPanelFilterHandler(commandPanelWidget *views.CommandPanelWidget, list *views.ListWidget) *CommandPanelFilterHandler {
	handler := &CommandPanelFilterHandler{
		commandPanelWidget: commandPanelWidget,
		list:               list,
	}
	handler.id = HandlerIDFilter
	return handler
}

func (h *CommandPanelFilterHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}
func (h *CommandPanelFilterHandler) DisplayText() string {
	return "Filter"
}
func (h *CommandPanelFilterHandler) IsEnabled() bool {
	return true
}
func (h *CommandPanelFilterHandler) Invoke() error {
	h.commandPanelWidget.ShowWithText("Filter", "", nil, h.CommandPanelNotification)
	return nil
}
func (h *CommandPanelFilterHandler) CommandPanelNotification(state interfaces.CommandPanelNotification) {
	h.list.SetFilter(state.CurrentText)
	if state.EnterPressed {
		h.commandPanelWidget.Hide()
	}
}
