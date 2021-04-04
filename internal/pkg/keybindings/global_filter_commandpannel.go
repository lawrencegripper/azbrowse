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
	itemView           *views.ItemWidget
}

var _ Command = &CommandPanelFilterHandler{}

func NewCommandPanelFilterHandler(
	commandPanelWidget *views.CommandPanelWidget,
	list *views.ListWidget,
	itemView *views.ItemWidget) *CommandPanelFilterHandler {

	handler := &CommandPanelFilterHandler{
		commandPanelWidget: commandPanelWidget,
		list:               list,
		itemView:           itemView,
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
	switch h.commandPanelWidget.PreviousViewName {
	case "listWidget":
		h.list.SetFilter(state.CurrentText)
	case "itemWidget":
		h.itemView.SetFilter(state.CurrentText)
	}
	if state.EnterPressed {
		h.commandPanelWidget.Hide()
	}
}
