package views

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
)

// StatusbarWidget controls the statusbar
type StatusbarWidget struct {
	name            string
	x, y            int
	w               int
	hideGuids       bool
	messages        map[string]eventing.StatusEvent
	currentMessage  *eventing.StatusEvent
	messageAddition string
}

// NewStatusbarWidget create new instance and start go routine for spinner
func NewStatusbarWidget(x, y, w int, hideGuids bool, g *gocui.Gui) *StatusbarWidget {
	widget := &StatusbarWidget{
		name:      "statusBarWidget",
		x:         x,
		y:         y,
		w:         w,
		hideGuids: hideGuids,
		messages:  map[string]eventing.StatusEvent{},
	}

	widget.currentMessage = &eventing.StatusEvent{}

	newEvents := eventing.SubscribeToStatusEvents()
	// Start loop for showing loading in statusbar
	go func() {
		for {
			// Wait for a second to see if we have any new messages
			timeout := time.After(time.Second)
			select {
			case eventObj := <-newEvents:
				// See if we have any new events
				event := eventObj.(eventing.StatusEvent)
				widget.messages[event.ID()] = event
				// Favour the most recent message
				widget.currentMessage = &event
			case <-timeout:
				// Update the UI
				continue
			}

			for _, message := range widget.messages {
				// Remove any that have now expired
				if message.HasExpired() {
					delete(widget.messages, message.ID())
					continue
				}
			}

			// Set the current message to a non-expired message favour in-progress messages
			if !widget.currentMessage.HasExpired() || !widget.currentMessage.InProgress {
				foundInProgress := false
				for _, message := range widget.messages {
					if message.InProgress {
						foundInProgress = true
						widget.currentMessage = &message
						break
					}
				}

				if !foundInProgress {
					for _, message := range widget.messages {
						widget.currentMessage = &message
						break
					}
				}
			}

			g.Update(func(gui *gocui.Gui) error {
				if widget.currentMessage.InProgress {
					widget.messageAddition = widget.messageAddition + "."
				} else {
					widget.messageAddition = ""
				}
				return nil
			})

			time.Sleep(time.Second)

		}
	}()
	return widget
}

// Layout draws the widget in the gocui view
func (w *StatusbarWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+3)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	v.Title = `Status [CTRL+I -> Help]`
	v.Wrap = true

	if w.hideGuids {
		w.currentMessage.Message = stripSecretVals(w.currentMessage.Message)
	}

	if w.currentMessage.InProgress {
		fmt.Fprint(v, style.Loading("⏳  "+w.currentMessage.Message))
	} else if w.currentMessage.Failure {
		fmt.Fprint(v, style.Loading("☠ "+w.currentMessage.Message))
	} else {
		fmt.Fprint(v, style.Completed("✓ "+w.currentMessage.Message))
	}
	fmt.Fprint(v, w.messageAddition)

	return nil
}

// Status updates the message in the status bar and whether to show loading indicator
func (w *StatusbarWidget) Status(message string, loading bool) func() {
	_, done := eventing.SendStatusEvent(eventing.StatusEvent{
		Message:    message,
		InProgress: loading,
		Timeout:    time.Duration(time.Second * 3),
	})
	return done
}
