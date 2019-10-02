package views

import (
	"github.com/jroimartin/gocui"
)

// CommandPanelWidget controls the notifications windows in the top right
type CommandPanelWidget struct {
	name    string
	x, y    int
	w       int
	visible bool
	gui     *gocui.Gui
}

// NewCommandPanelWidget create new instance and start go routine for spinner
func NewCommandPanelWidget(x, y, w int, g *gocui.Gui) *CommandPanelWidget {
	widget := &CommandPanelWidget{
		name: "commandPanelWidget",
		x:    x,
		y:    y,
		w:    w,
		gui:  g,
	}
	return widget
}

// ToggleShowHide shows and hides the command panel
func (w *CommandPanelWidget) ToggleShowHide() {
	w.visible = !w.visible
}

// Layout draws the widget in the gocui view
func (w *CommandPanelWidget) Layout(g *gocui.Gui) error {
	// Don't draw anything if no pending deletes
	if !w.visible {
		g.DeleteView(w.name)
		g.Cursor = false
		return nil
	}

	height := 3

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+height)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Editable = true
	g.Cursor = true
	v.Title = "Command Pallet"
	v.Wrap = false

	// panelContent := v.Buffer()
	// if len(panelContent) > 0 {
	// 	eventing.SendStatusEvent(eventing.StatusEvent{
	// 		Failure: false,
	// 		Message: "Panel contains: " + panelContent,
	// 		Timeout: time.Second * 5,
	// 	})
	// }

	_, err = w.gui.SetCurrentView(w.name)
	if err != nil {
		panic("No view " + w.name + " found: " + err.Error())
	}

	return nil
}
