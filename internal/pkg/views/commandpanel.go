package views

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
)

// CommandPanelWidget controls the notifications windows in the top right
type CommandPanelWidget struct {
	name             string
	x, y             int
	w                int
	visible          bool
	gui              *gocui.Gui
	panelContent     string
	prepopulate      string
	previousViewName string
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

// ToggleShowHide shows and hides the command panel
func (w *CommandPanelWidget) ToggleShowHide() {
	// Ensure we put things back how we found them before the panel was launched
	if !w.visible {
		w.trackPreviousView()
	}

	// Show or hide
	w.visible = !w.visible
}

// ShowWithText launches the command panel pre-populated with some text
func (w *CommandPanelWidget) ShowWithText(s string) {
	// Ensure we put things back how we found them before the panel was launched
	w.trackPreviousView()

	w.gui.DeleteView(w.name) //nolint: errcheck
	w.gui.Cursor = false
	w.prepopulate = s
	w.visible = true
}

func (w *CommandPanelWidget) trackPreviousView() {
	if view := w.gui.CurrentView(); view != nil {
		w.previousViewName = view.Name()
	}
}

// Layout draws the widget in the gocui view
func (w *CommandPanelWidget) Layout(g *gocui.Gui) error {
	// If we're not updating an existing view then
	// set the content to the value from prepopulate
	viewExists := true
	_, err := g.View(w.name)
	if err == gocui.ErrUnknownView {
		viewExists = false
	}

	// If we have a view but we're not meant to clean up
	if viewExists && !w.visible {
		g.DeleteView(w.name)
		g.SetCurrentView(w.previousViewName) //nolint: errcheck
		g.Cursor = false
		return nil
	}

	// If the view doesn't exists and we're not visible do nothing
	if !w.visible {
		return nil
	}

	height := 2

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+height)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Editable = true
	g.Cursor = true
	v.Title = "Command Pallet"
	v.Wrap = false

	// Is this a new view?
	if !viewExists {
		// It is lets prepopulate it with content like `/` if it was launched from the filter handler
		v.Write([]byte(w.prepopulate))     //nolint: errcheck
		v.SetCursor(len(w.prepopulate), 0) //nolint: errcheck
	} else if w.panelContent != v.Buffer() {
		// Has something been typed?
		w.panelContent = v.Buffer()
		// Handle this change and take action
		w.panelChanged()
	}

	_, err = w.gui.SetCurrentView(w.name)
	if err != nil {
		panic("No view " + w.name + " found: " + err.Error())
	}

	return nil
}

// This is fired every time the content of the command panel has changed.
func (w *CommandPanelWidget) panelChanged() {
	content := w.panelContent

	if len(content) < 1 {
		return
	}

	// What command is being used?

	// `/` -> Filter command
	if strings.HasPrefix(content, "/") {
		eventing.Publish("filter", content)
	}

	// Handle return by closing panel... events should handle dealing with whats needed
	if strings.Count(content, "\n") > 1 { //Warning: By default gocui is adding a newline so we're checking for existence of 2
		w.ToggleShowHide()
	}
}
