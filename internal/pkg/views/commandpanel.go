package views

import (
	"strings"

	// "github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/stuartleeks/gocui"
)

type CommandPanelListOption interface {
	ID() string
	DisplayText() string
}

type CommandPanelNotification struct {
	CurrentText  string
	SelectedID   string
	EnterPressed bool
}
type CommandPanelNotificationHandler func(state CommandPanelNotification)

// CommandPanelWidget controls the notifications windows in the top right
type CommandPanelWidget struct {
	name                string
	x, y                int
	w                   int
	visible             bool
	gui                 *gocui.Gui
	panelContent        string
	newPanelContent     string
	prepopulate         string
	previousViewName    string
	title               string
	options             *[]CommandPanelListOption
	filteredOptions     *[]CommandPanelListOption
	selectedIndex       int
	notificationHandler CommandPanelNotificationHandler
}

// NewCommandPanelWidget create new instance and start go routine for spinner
func NewCommandPanelWidget(x, y, w int, g *gocui.Gui) *CommandPanelWidget {
	widget := &CommandPanelWidget{
		name:    "commandPanelWidget",
		x:       x,
		y:       y,
		w:       w,
		gui:     g,
		visible: false,
	}
	return widget
}

// Hide hides the command panel
func (w *CommandPanelWidget) Hide() {
	// hide
	w.prepopulate = ""
	w.visible = false
}

// ShowWithText launches the command panel pre-populated with some text
func (w *CommandPanelWidget) ShowWithText(title string, s string, options *[]CommandPanelListOption, handler CommandPanelNotificationHandler) {
	// Ensure we put things back how we found them before the panel was launched
	w.trackPreviousView()

	w.gui.DeleteView(w.name) //nolint: errcheck
	w.title = title
	w.gui.Cursor = false
	w.prepopulate = s
	w.options = options
	w.filteredOptions = options
	w.notificationHandler = handler
	w.visible = true
	w.selectedIndex = -1
}

// MoveDown moves down a list item if options are displayed
func (w *CommandPanelWidget) MoveDown() {
	if w.filteredOptions != nil && len(*w.filteredOptions) > 0 {
		w.selectedIndex += 1
		if w.selectedIndex >= len(*w.filteredOptions) {
			w.selectedIndex = len(*w.filteredOptions) - 1
		}
	}
}

// MoveDown moves up a list item if options are displayed
func (w *CommandPanelWidget) MoveUp() {
	if w.filteredOptions != nil && len(*w.filteredOptions) > 0 {
		w.selectedIndex -= 1
		if w.selectedIndex < 0 {
			w.selectedIndex = 0
		}
	}
}
// EnterPressed is used to communicate that the enter key was pressed but a handler received it
func (w *CommandPanelWidget) EnterPressed() {
	// the handler was added to invoke this method as Enter without any input failed to trigger the update
	w.panelChanged(w.panelContent + "\n\n") // ensure newlines to trigger EnterPressed logic
}

func (w *CommandPanelWidget) trackPreviousView() {
	if view := w.gui.CurrentView(); view != nil {
		w.previousViewName = view.Name()
	}
}

// Layout draws the widget in the gocui view
func (w *CommandPanelWidget) Layout(g *gocui.Gui) error {

	inputViewName := w.name
	optionsViewName := w.name + "Options"

	// If we're not updating an existing view then
	// set the content to the value from prepopulate
	viewExists := true
	_, err := g.View(inputViewName)
	if err == gocui.ErrUnknownView {
		viewExists = false
	}

	// If we're not visible then do any clean-up needed
	if !w.visible {
		if _, err := g.View(optionsViewName); err != gocui.ErrUnknownView {
			g.DeleteView(optionsViewName)
		}
		if viewExists {
			g.DeleteView(inputViewName)
			g.SetCurrentView(w.previousViewName) //nolint: errcheck
			g.Cursor = false
		}
		return nil
	}

	if w.options == nil {
		// delete options view if now options
		if _, err := g.View(optionsViewName); err != gocui.ErrUnknownView {
			g.DeleteView(optionsViewName)
		}
	}

	height := 2

	var vList *gocui.View
	if w.options != nil {
		vList, err = g.SetView(optionsViewName, w.x, w.y+2, w.x+w.w, w.y+13)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}

	v, err := g.SetView(inputViewName, w.x, w.y, w.x+w.w, w.y+height)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Editable = true
	g.Cursor = true
	v.Title = w.title
	v.Wrap = false

	if w.options != nil {
		vList.Clear()
		for i, option := range *w.filteredOptions {
			itemToShow := ""
			if i == w.selectedIndex {
				itemToShow = "▶ "
			} else {
				itemToShow = "  "
			}
			itemToShow += option.DisplayText() + "\n"

			vList.Write([]byte(itemToShow))
		}
	}

	// Is this a new view?
	if !viewExists {
		// It is lets prepopulate it with content like `/` if it was launched from the filter handler
		v.Write([]byte(w.prepopulate))     //nolint: errcheck
		v.SetCursor(len(w.prepopulate), 0) //nolint: errcheck

	} else if w.newPanelContent != "" {
		// update panel contents
		v.Clear()
		if _, err := v.Write([]byte(w.newPanelContent)); err != nil {
			panic("Failed to write to view")
		}
		if err := v.SetCursor(len(w.newPanelContent), 0); err != nil {
			panic("Failed to set cursor position")
		}
		if err := v.SetOrigin(0, 0); err != nil {
			panic("Failed to set view origin")
		}
		w.panelContent = w.newPanelContent
		w.newPanelContent = ""
	} else if w.panelContent != v.Buffer() {
		// Has something been typed?
		w.panelContent = v.Buffer()
		// Handle this change and take action
		w.panelChanged(w.panelContent)
	}

	_, err = w.gui.SetCurrentView(inputViewName)
	if err != nil {
		panic("No view " + inputViewName + " found: " + err.Error())
	}

	return nil
}

// This is fired every time the content of the command panel has changed.
func (w *CommandPanelWidget) panelChanged(content string) {

	if len(content) < 1 && w.selectedIndex < 0 {
		return
	}

	state := CommandPanelNotification{
		EnterPressed: w.contentHasNewLine(content),
		CurrentText:  strings.ReplaceAll(content, "\n", ""),
	}

	triggerLayout := false
	// Clear newlines to allow usage with repeated enter presses (e.g. search)
	if state.EnterPressed {
		w.newPanelContent = state.CurrentText
		triggerLayout = true
	}
	if w.options != nil {
		// apply filter, re-selecting the current item
		selectedID := ""
		if w.selectedIndex >= 0 && len(*w.filteredOptions) > 0 {
			selectedID = (*w.filteredOptions)[w.selectedIndex].ID()
		}
		w.selectedIndex = -1
		filterOptions := []CommandPanelListOption{}
		loweredCurrentText := strings.ToLower(state.CurrentText)
		for i, option := range *w.options {
			if strings.Contains(strings.ToLower(option.DisplayText()), loweredCurrentText) {
				if option.ID() == selectedID {
					w.selectedIndex = i
				}
				filterOptions = append(filterOptions, option)
			}
		}
		w.filteredOptions = &filterOptions
		triggerLayout = true
	}

	if state.EnterPressed {
		if w.filteredOptions != nil {
			if w.selectedIndex >= 0 && len(*w.filteredOptions) > 0 {
				state.SelectedID = (*w.filteredOptions)[w.selectedIndex].ID()
			} else if len(*w.filteredOptions) == 1 {
				// no selected item but only one left - pretend it was selected
				state.SelectedID = (*w.filteredOptions)[0].ID()
			}
		}
	}

	if triggerLayout {
		if err := w.Layout(w.gui); err != nil {
			panic("Layout failed")
		}
	}

	w.gui.Update(func(gui *gocui.Gui) error {
		w.notificationHandler(state)
		return nil
	})

}

func (w *CommandPanelWidget) contentHasNewLine(content string) bool {
	return strings.Count(content, "\n") > 1 //Warning: By default gocui is adding a newline so we're checking for existence of 2
}
