package keybindings

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type OpenCommandPanelHandler struct {
	GlobalHandler
	gui                *gocui.Gui
	commandPanelWidget *views.CommandPanelWidget
	commands           []Command
}

func NewOpenCommandPanelHandler(gui *gocui.Gui, commandPanelWidget *views.CommandPanelWidget, commands []Command) *OpenCommandPanelHandler {
	handler := &OpenCommandPanelHandler{
		gui:                gui,
		commandPanelWidget: commandPanelWidget,
		commands:           commands,
	}
	handler.id = HandlerIDToggleOpenCommandPanel
	return handler
}

func (h *OpenCommandPanelHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		keyBindings := GetKeyBindingsAsStrings()

		paletteWidth := h.commandPanelWidget.Width() - 4
		options := []interfaces.CommandPanelListOption{}
		for _, command := range h.commands {
			if command.IsEnabled() {
				commandID := command.ID()
				commandDisplayText := command.DisplayText()

				// ensure binding is in upper-case
				binding := keyBindings[commandID]
				for i, b := range binding {
					binding[i] = strings.ToUpper(b)
				}

				// calculate padding to right-align shortcut
				bindingString := ""
				if len(binding) > 0 {
					bindingString = fmt.Sprintf("%s", binding)
				}
				padAmount := paletteWidth - len(commandDisplayText) - len(bindingString)
				if padAmount < 0 {
					padAmount = 0 // TODO - we should also look at truncating the DisplayText
				}

				option := interfaces.CommandPanelListOption{
					ID:          commandID,
					DisplayText: command.DisplayText() + strings.Repeat(" ", padAmount) + bindingString,
				}

				options = append(options, option)
			}
		}
		h.commandPanelWidget.ShowWithText("Command Palette", "", &options, h.CommandPanelNotification)
		return nil
	}
}

func (h *OpenCommandPanelHandler) CommandPanelNotification(state interfaces.CommandPanelNotification) {
	if state.EnterPressed {
		h.commandPanelWidget.Hide()
		for _, command := range h.commands {
			if command.ID() == state.SelectedID {
				// invoke via Update to allow Hide to restore preview view state
				h.gui.Update(func(gui *gocui.Gui) error {
					return command.Invoke()
				})
				return
			}
		}
	}
}
