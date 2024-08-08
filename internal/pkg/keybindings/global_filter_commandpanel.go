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
	fuzzyFilter        bool
}

var _ Command = &CommandPanelFilterHandler{}

func NewCommandPanelFilterHandler(commandPanelWidget *views.CommandPanelWidget, fuzzyFilter bool) *CommandPanelFilterHandler {

	handler := &CommandPanelFilterHandler{
		commandPanelWidget: commandPanelWidget,
		fuzzyFilter:        fuzzyFilter,
	}
	if fuzzyFilter {
		handler.id = HandlerIDFilterFuzzy
	} else {
		handler.id = HandlerIDFilter
	}
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
	if h.fuzzyFilter {
		return "Filter (Fuzzy)"
	} else {
		return "Filter"
	}
}
func (h *CommandPanelFilterHandler) IsEnabled() bool {
	return true
}
func (h *CommandPanelFilterHandler) getTitle() string {
	if h.fuzzyFilter {
		return "Filter (fuzzy)"
	} else {
		return "Filter"
	}
}
func (h *CommandPanelFilterHandler) Invoke() error {
	h.commandPanelWidget.ShowWithText(h.getTitle(), "", nil, h.CommandPanelNotification)
	return nil
}

func (h *CommandPanelFilterHandler) InvokeWithStartString(s string) error {
	h.commandPanelWidget.ShowWithText(h.getTitle(), s, nil, h.CommandPanelNotification)
	return nil
}
func (h *CommandPanelFilterHandler) CommandPanelNotification(state interfaces.CommandPanelNotification) {
	switch h.commandPanelWidget.PreviousViewName {
	case "listWidget":
		h.list.SetFilter(state.CurrentText, h.fuzzyFilter)
	case "itemWidget":
		h.itemView.SetFilter(state.CurrentText, h.fuzzyFilter)
	}
	if state.EnterPressed {
		h.commandPanelWidget.Hide()
	}
}
