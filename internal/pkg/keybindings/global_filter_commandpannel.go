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

func NewCommandPanelFilterHandler(commandPanelWidget *views.CommandPanelWidget) *CommandPanelFilterHandler {

	handler := &CommandPanelFilterHandler{
		commandPanelWidget: commandPanelWidget,
	}
	handler.id = HandlerIDFilter
	return handler
}

// Hack to work around circular import
func (h *CommandPanelFilterHandler) SetItemWidget(w *views.ItemWidget) {
	h.itemView = w
}

// Hack #2 to work around circular import
func (h *CommandPanelFilterHandler) SetListWidget(w *views.ListWidget) {
	h.list = w
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
func (h *CommandPanelFilterHandler) InvokeWithStartString(s string) error {
	h.commandPanelWidget.ShowWithText("Filter", s, nil, h.CommandPanelNotification)
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
