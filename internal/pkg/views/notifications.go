package views

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
)

type pendingDelete struct {
	display string
	url     string
}

// NotificationWidget controls the notifications windows in the top right
type NotificationWidget struct {
	ConfirmDeleteKeyBinding       string
	ClearPendingDeletesKeyBinding string
	name                          string
	x, y                          int
	w                             int
	pendingDeletes                []pendingDelete
	deleteMutex                   sync.Mutex // ensure delete occurs only once
	deleteInProgress              bool
	gui                           *gocui.Gui
}

// AddPendingDelete queues deletes for
// delete once confirmed
func (w *NotificationWidget) AddPendingDelete(display, url string) {
	if w.deleteInProgress {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Delete already in progress. Please wait for completion.",
			Timeout: time.Second * 5,
		})
		return
	}

	if url == "" {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Item `" + display + "` doesn't support delete",
			Timeout: time.Second * 5,
		})
		return
	}

	// Don't add more items than we can draw on the
	// current terminal size
	_, yMax := w.gui.Size()
	if len(w.pendingDeletes) > (yMax - 12) {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Can't add `" + display + "` run out of space to draw the `Pending delete` list!",
			Timeout: time.Second * 5,
		})
		return
	}

	w.deleteMutex.Lock()
	defer w.deleteMutex.Unlock()

	for _, i := range w.pendingDeletes {
		if i.url == url {
			eventing.SendStatusEvent(eventing.StatusEvent{
				Failure: true,
				Message: "Item already `" + display + "` in pending delete list",
				Timeout: time.Second * 5,
			})
			return
		}
	}

	w.pendingDeletes = append(w.pendingDeletes, pendingDelete{
		url:     url,
		display: display,
	})
}

// ConfirmDelete delete all queued/pending deletes
func (w *NotificationWidget) ConfirmDelete() {
	if w.deleteInProgress {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Delete already in progress. Please wait for completion.",
			Timeout: time.Second * 5,
		})
		return
	}

	w.deleteMutex.Lock()
	w.deleteInProgress = true

	go func() {
		// unlock and mark delete as not in progress
		defer w.deleteMutex.Unlock()
		defer func() {
			w.deleteInProgress = false
		}()

		pending := w.pendingDeletes
		event, _ := eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: true,
			Message:    "Starting to delete items",
			Timeout:    time.Second * 15,
		})

		for _, i := range pending {
			_, err := armclient.DoRequest(context.Background(), "DELETE", i.url)
			if err != nil {
				event.Failure = true
				event.InProgress = false
				event.Message = "Failed to delete `" + i.display + "` with error:" + err.Error()
				event.Update()

				// In the event that a delete fails in the
				// batch of pending deletes lets give up on the rest
				// as something might have gone wrong and best
				// to be cautious
				return
			}

			event.Message = "Deleted: " + i.display
			event.Update()
		}

		event.Message = "Delete request sent"
		event.InProgress = false
		event.Update()

		w.pendingDeletes = []pendingDelete{}
	}()
}

// ClearPendingDeletes removes all pending deletes
func (w *NotificationWidget) ClearPendingDeletes() {
	w.deleteMutex.Lock()
	w.gui.Update(func(g *gocui.Gui) error {

		eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: true,
			Message:    "Clearing pending deletes",
			Timeout:    time.Second * 2,
		})

		w.pendingDeletes = []pendingDelete{}
		w.deleteMutex.Unlock()

		return nil
	})
}

// NewNotificationWidget create new instance and start go routine for spinner
func NewNotificationWidget(x, y, w int, hideGuids bool, g *gocui.Gui) *NotificationWidget {
	widget := &NotificationWidget{
		name:           "notificationWidget",
		x:              x,
		y:              y,
		w:              w,
		gui:            g,
		pendingDeletes: []pendingDelete{},
	}
	return widget
}

// Layout draws the widget in the gocui view
func (w *NotificationWidget) Layout(g *gocui.Gui) error {
	// Don't draw anything if no pending deletes
	if len(w.pendingDeletes) < 1 {
		g.DeleteView(w.name)
		return nil
	}

	height := len(w.pendingDeletes)*1 + 7

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, height)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Clear()
	v.Title = "Notifications [ESC to clear]"
	v.Wrap = false

	return w.layoutInternal(v)
}

func (w *NotificationWidget) layoutInternal(v io.Writer) error {
	pending := w.pendingDeletes

	fmt.Fprintln(v, style.Title("Pending Deletes:"))
	for _, i := range pending {
		fmt.Fprintln(v, " - "+i.display)
	}
	fmt.Fprintln(v, "")
	fmt.Fprintln(v, "Do you want to delete these items?")
	fmt.Fprintln(v, style.Warning("Press "+strings.ToUpper(w.ConfirmDeleteKeyBinding)+" to DELETE"))
	fmt.Fprintln(v, style.Highlight("Press "+strings.ToUpper(w.ClearPendingDeletesKeyBinding)+" to CANCEL"))

	return nil
}
