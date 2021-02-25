package keybindings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/stuartleeks/gocui"
)

////////////////////////////////////////////////////////////////////
type QuitHandler struct {
	GlobalHandler
}

func NewQuitHandler() *QuitHandler {
	handler := &QuitHandler{}
	handler.id = HandlerIDQuit
	return handler
}

func (h QuitHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type CopyHandler struct {
	GlobalHandler
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &CopyHandler{}

func NewCopyHandler(content *views.ItemWidget, statusbar *views.StatusbarWidget) *CopyHandler {
	handler := &CopyHandler{
		Content:   content,
		StatusBar: statusbar,
	}
	handler.id = HandlerIDCopy
	return handler
}

func (h CopyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *CopyHandler) DisplayText() string {
	return "Copy content"
}

func (h *CopyHandler) IsEnabled() bool {
	return true
}

func (h *CopyHandler) Invoke() error {
	var err error
	contentType := h.Content.GetContentType()
	content := h.Content.GetContent()

	var formattedContent string
	switch contentType {
	case interfaces.ResponseJSON:
		if !json.Valid([]byte(content)) {
			h.StatusBar.Status("Resource content is not valid JSON", false)
			return fmt.Errorf("Resource content is not valid JSON: %s", content)
		}

		var formattedBuf bytes.Buffer
		err = json.Indent(&formattedBuf, []byte(content), "", "  ")
		if err != nil {
			h.StatusBar.Status(fmt.Sprintf("Error formatting JSON for editor: %s", err), false)
			return err
		}

		formattedContent = formattedBuf.String()
	case interfaces.ResponseYAML:
		formattedContent = content // TODO: add YAML formatter

	default:
		formattedContent = content
	}

	if err := copyToClipboard(formattedContent); err != nil {
		h.StatusBar.Status(fmt.Sprintf("Failed to copy to clipboard: %s", err.Error()), false)
		return nil
	}
	h.StatusBar.Status("Current resource's content copied to clipboard", false)
	return nil
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type FullscreenHandler struct {
	GlobalHandler
	List         *views.ListWidget
	IsFullscreen *bool
	Content      *views.ItemWidget
}

func NewFullscreenHandler(list *views.ListWidget, content *views.ItemWidget, isFullscreen *bool) *FullscreenHandler {
	handler := &FullscreenHandler{
		List:         list,
		Content:      content,
		IsFullscreen: isFullscreen,
	}
	handler.id = HandlerIDFullScreen
	return handler
}

func (h FullscreenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.IsFullscreen)
		h.IsFullscreen = &tmp // memory leak?
		if *h.IsFullscreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			keyBindings := GetKeyBindingsAsStrings()
			v.Title = fmt.Sprintf("JSON Response - Fullscreen (%s to exit)", strings.ToUpper(strings.Join(keyBindings["fullscreen"], ",")))

			content := h.Content.GetContent()
			fmt.Fprint(v, style.ColorJSON(content))

			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type HelpHandler struct {
	GlobalHandler
	ShowHelp *bool
}

func NewHelpHandler(showHelp *bool) *HelpHandler {
	handler := &HelpHandler{
		ShowHelp: showHelp,
	}
	handler.id = HandlerIDHelp
	return handler
}

func (h HelpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.ShowHelp)
		h.ShowHelp = &tmp // memory leak?

		// If we're up and running clear and redraw the view
		// if w.g != nil {
		if *h.ShowHelp {
			v, err := g.SetView("helppopup", 1, 1, 145, 45)
			g.SetCurrentView("helppopup")
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			keyBindings := GetKeyBindingsAsStrings()
			views.DrawHelp(keyBindings, v)
		} else {
			g.DeleteView("helppopup")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ConfirmDeleteHandler struct {
	GlobalHandler
	notificationWidget *views.NotificationWidget
}

func NewConfirmDeleteHandler(notificationWidget *views.NotificationWidget) *ConfirmDeleteHandler {
	handler := &ConfirmDeleteHandler{
		notificationWidget: notificationWidget,
	}
	handler.id = HandlerIDConfirmDelete
	return handler
}

func (h *ConfirmDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.notificationWidget.ConfirmDelete()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ClearPendingDeleteHandler struct {
	GlobalHandler
	notificationWidget *views.NotificationWidget
}

func NewClearPendingDeleteHandler(notificationWidget *views.NotificationWidget) *ClearPendingDeleteHandler {
	handler := &ClearPendingDeleteHandler{
		notificationWidget: notificationWidget,
	}
	handler.id = HandlerIDClearPendingDeletes
	return handler
}

func (h *ClearPendingDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.notificationWidget.ClearPendingDeletes()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////
