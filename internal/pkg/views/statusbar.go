package views

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
)

// StatusbarWidget controls the statusbar
type StatusbarWidget struct {
	name            string
	x, y            int
	w               int
	hideGuids       bool
	messages        map[string]*eventing.StatusEvent
	currentMessage  *eventing.StatusEvent
	messageAddition string
	HelpKeyBinding  string
}

// NewStatusbarWidget create new instance and start go routine for spinner
func NewStatusbarWidget(x, y, w int, hideGuids bool, g *gocui.Gui) *StatusbarWidget {
	widget := &StatusbarWidget{
		name:      "statusBarWidget",
		x:         x,
		y:         y,
		w:         w,
		hideGuids: hideGuids,
		messages:  map[string]*eventing.StatusEvent{},
	}

	widget.currentMessage = &eventing.StatusEvent{}

	newEvents := eventing.SubscribeToStatusEvents()
	// Start loop for showing loading in statusbar
	go func() {
		// recover from panic, if one occurrs, and leave terminal usable
		defer errorhandling.RecoveryWithCleanup()

		for {
			// Track if we need to redraw UI
			changesMade := false
			// Wait for a second to see if we have any new messages
			timeout := time.After(time.Second)
			select {
			case eventObj := <-newEvents:
				widget.addStatusEvent(eventObj)
			case <-timeout:
				// Update the UI
			}

			// Seeing as we're about to process and thats quite a bit of effort
			// lets grab everything off the channel, if there is stuff
			// stacked up waiting for us
			itemsInChan := len(newEvents)
			for index := 0; index < itemsInChan; index++ {
				eventObj := <-newEvents
				widget.addStatusEvent(eventObj)
			}

			for _, message := range widget.messages {
				// Remove any that have now expired
				if message.HasExpired() {
					delete(widget.messages, message.ID())
					continue
				}
			}

			messages := make([]*eventing.StatusEvent, 0, len(widget.messages))
			for _, message := range widget.messages {
				messages = append(messages, message)
			}

			sort.Slice(messages, func(i, j int) bool {
				return messages[i].CreatedAt().After(messages[j].CreatedAt())
			})

			// Find the first message which isn't past it's timeout
			// and is in progress
			haveUpdatedMessage := false
			for _, message := range messages {
				if message.HasExpired() {
					continue
				}
				if !message.InProgress {
					continue
				}

				widget.currentMessage = message
				widget.messageAddition = ""
				haveUpdatedMessage = true
				changesMade = true
				break
			}

			// If we didn't find one of those pick the most recent message
			if !haveUpdatedMessage && len(messages) > 0 && widget.currentMessage != messages[0] {
				widget.currentMessage = messages[0]
				widget.messageAddition = ""
				changesMade = true
			}

			if widget.currentMessage.InProgress {
				widget.messageAddition = widget.messageAddition + "."
				changesMade = true
			}

			// If we're made some changes redraw the UI
			if changesMade {
				g.Update(func(gui *gocui.Gui) error {
					return nil
				})
			}
		}
	}()
	return widget
}

func (w *StatusbarWidget) addStatusEvent(eventObj interface{}) {
	// See if we have any new events
	event := eventObj.(*eventing.StatusEvent)
	if event.IsToast {
		// Leave toast notifications for the notifications panel
		return
	}
	w.messages[event.ID()] = event
	// Favor the most recent message
	w.currentMessage = event
}

// Layout draws the widget in the gocui view
func (w *StatusbarWidget) Layout(g *gocui.Gui) error {
	x0, y0, x1, y1 := getViewBounds(g, w.x, w.y, w.w, 3)

	v, err := g.SetView(w.name, x0, y0, x1, y1, 0)
	if err != nil && errors.Is(err, gocui.ErrUnknownView) {
		return err
	}
	v.Clear()
	v.Title = fmt.Sprintf(`Status [%s -> Help]`, strings.ToUpper(w.HelpKeyBinding))
	v.Wrap = true

	if w.hideGuids {
		w.currentMessage.Message = StripSecretVals(w.currentMessage.Message)
	}

	if w.currentMessage.InProgress || w.currentMessage.Failure {
		fmt.Fprint(v, style.Loading(w.currentMessage.Icon()+"  "+w.currentMessage.Message))
	} else {
		fmt.Fprint(v, style.Completed(w.currentMessage.Icon()+" "+w.currentMessage.Message))
	}

	fmt.Fprint(v, w.messageAddition)

	return nil
}

// Status updates the message in the status bar and whether to show loading indicator
func (w *StatusbarWidget) Status(message string, loading bool) func() {
	_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    message,
		InProgress: loading,
		Timeout:    time.Duration(time.Second * 3),
	})
	return done
}

// SetHideGuids sets the HideGuids option
func (w *StatusbarWidget) SetHideGuids(value bool) {
	w.hideGuids = value
}
