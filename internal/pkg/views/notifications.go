package views

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
)

func init() {
	pendingDeletes = []pendingDelete{}
}

type pendingDelete struct {
	display string
	url     string
}

var pendingDeletes []pendingDelete
var deleteMutex sync.Mutex // ensure delete occurs only once
var gui *gocui.Gui

// AddPendingDelete queues deletes for
// delete once confirmed
func AddPendingDelete(display, url string) {
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
	_, yMax := gui.Size()
	if len(pendingDeletes) > (yMax - 12) {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Can't add `" + display + "` run out of space to draw the `Pending delete` list!",
			Timeout: time.Second * 5,
		})
		return
	}

	gui.Update(func(g *gocui.Gui) error {
		deleteMutex.Lock()
		defer deleteMutex.Unlock()

		for _, i := range pendingDeletes {
			if i.url == url {
				eventing.SendStatusEvent(eventing.StatusEvent{
					Failure: true,
					Message: "Item already `" + display + "` in pending delete list",
					Timeout: time.Second * 5,
				})
				return nil
			}
		}

		pendingDeletes = append(pendingDeletes, pendingDelete{
			url:     url,
			display: display,
		})

		return nil
	})
}

// ConfirmDelete delete all queued/pending deletes
func ConfirmDelete() {
	go func() {
		deleteMutex.Lock()
		defer deleteMutex.Unlock()

		pending := pendingDeletes
		event, _ := eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: true,
			Message:    "Starting to delete items",
			Timeout:    time.Second * 15,
		})

		for _, i := range pending {
			_, err := armclient.DoRequest(context.Background(), "DELETE", i.url)
			if err != nil {
				event.Failure = true
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

		pendingDeletes = []pendingDelete{}
	}()
}

// ClearPendingDeletes removes all pending deletes
func ClearPendingDeletes() {
	deleteMutex.Lock()
	gui.Update(func(g *gocui.Gui) error {

		eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: true,
			Message:    "Clearing pending deletes",
			Timeout:    time.Second * 2,
		})

		pendingDeletes = []pendingDelete{}
		deleteMutex.Unlock()

		return nil
	})
}

// NotificationWidget controls the statusbar
type NotificationWidget struct {
	ConfirmDeleteKeyBinding       string
	ClearPendingDeletesKeyBinding string
	name                          string
	x, y                          int
	w                             int
}

// NewNotificationWidget create new instance and start go routine for spinner
func NewNotificationWidget(x, y, w int, hideGuids bool, g *gocui.Gui) *NotificationWidget {
	widget := &NotificationWidget{
		name: "notificationWidget",
		x:    x,
		y:    y,
		w:    w,
	}
	gui = g
	return widget
}

// Layout draws the widget in the gocui view
func (w *NotificationWidget) Layout(g *gocui.Gui) error {
	// Don't draw anything if no pending deletes
	if len(pendingDeletes) < 1 {
		g.DeleteView(w.name)
		return nil
	}

	height := len(pendingDeletes)*1 + 7

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, height)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Clear()
	v.Title = "Notifications [ESC to clear]"
	v.Wrap = false

	pending := pendingDeletes

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
