package views

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/stuartleeks/gocui"
)

// NotificationWidget controls the notifications windows in the top right
type NotificationWidget struct {
	ConfirmDeleteKeyBinding       string
	ClearPendingDeletesKeyBinding string
	name                          string
	x, y                          int
	w                             int
	pendingDeletes                []*expanders.TreeNode
	toastNotifications            map[string]*eventing.StatusEvent
	deleteMutex                   sync.Mutex // ensure delete occurs only once
	deleteInProgress              bool
	gui                           *gocui.Gui
	client                        *armclient.Client
}

// AddPendingDelete queues deletes for
// delete once confirmed
func (w *NotificationWidget) AddPendingDelete(item *expanders.TreeNode) {
	if w.deleteInProgress {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			Failure: true,
			Message: "Delete already in progress. Please wait for completion.",
			Timeout: time.Second * 5,
		})
		return
	}

	if item.DeleteURL == "" {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			Failure: true,
			Message: "Item `" + item.Name + "` doesn't support delete",
			Timeout: time.Second * 5,
		})
		return
	}

	// Don't add more items than we can draw on the
	// current terminal size
	_, yMax := w.gui.Size()
	if len(w.pendingDeletes) > (yMax - 12) {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			Failure: true,
			Message: "Can't add `" + item.Name + "` run out of space to draw the `Pending delete` list!",
			Timeout: time.Second * 5,
		})
		return
	}

	w.deleteMutex.Lock()
	defer w.deleteMutex.Unlock()

	for _, i := range w.pendingDeletes {
		if i.DeleteURL == item.DeleteURL {
			eventing.SendStatusEvent(&eventing.StatusEvent{
				Failure: true,
				Message: "Item already `" + item.Name + "` in pending delete list",
				Timeout: time.Second * 5,
			})
			return
		}
	}

	w.pendingDeletes = append(w.pendingDeletes, item)
}

// ConfirmDelete delete all queued/pending deletes
func (w *NotificationWidget) ConfirmDelete() {
	if w.deleteInProgress {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			Failure: true,
			Message: "Delete already in progress. Please wait for completion.",
			Timeout: time.Second * 5,
		})
		return
	}

	w.deleteMutex.Lock()
	w.deleteInProgress = true

	// Take a copy of the current pending deletes
	pending := make([]*expanders.TreeNode, len(w.pendingDeletes))
	copy(pending, w.pendingDeletes)

	w.deleteMutex.Unlock()

	go func() {
		// recover from panic, if one occurrs, and leave terminal usable
		defer errorhandling.RecoveryWithCleanup()

		// unlock and mark delete as not in progress
		defer func() {
			w.deleteInProgress = false
		}()

		event, _ := eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: true,
			Message:    "Starting to delete items",
			Timeout:    time.Second * 15,
		})

		// Deleting resources is an async operation so this timeout only
		// applies to issuing 'n' deletion requests. Rather than the time
		// the cloud takes to actually delete the resources.
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		for _, i := range pending {
			var err error
			fallback := true
			if i.Expander != nil {
				deleted, err := i.Expander.Delete(ctx, i)
				fallback = (err == nil && !deleted)
			}
			if fallback {
				// fallback to ARM request to delete
				_, err = w.client.DoRequest(ctx, "DELETE", i.DeleteURL)
			}
			if err != nil {
				event.Failure = true
				event.InProgress = false
				event.Message = "Failed to delete `" + i.Name + "` with error:" + err.Error()
				event.Update()

				w.pendingDeletes = []*expanders.TreeNode{}
				// In the event that a delete fails in the
				// batch of pending deletes lets give up on the rest
				// as something might have gone wrong and best
				// to be cautious
				return
			}

			event.Message = "Deleted: " + i.Name
			event.Update()
		}

		event.Message = "Delete request sent"
		event.InProgress = false
		event.Update()

		w.pendingDeletes = []*expanders.TreeNode{}
	}()
}

// ClearPendingDeletes removes all pending deletes
func (w *NotificationWidget) ClearPendingDeletes() {
	w.deleteMutex.Lock()
	w.gui.Update(func(g *gocui.Gui) error {

		_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: true,
			Message:    "Clearing pending deletes",
			Timeout:    time.Second * 2,
		})

		w.pendingDeletes = []*expanders.TreeNode{}
		w.deleteMutex.Unlock()
		done()

		return nil
	})
}

// NewNotificationWidget create new instance and start go routine for spinner
func NewNotificationWidget(x, y, w int, g *gocui.Gui, client *armclient.Client) *NotificationWidget {
	widget := &NotificationWidget{
		name:               "notificationWidget",
		x:                  x,
		y:                  y,
		w:                  w,
		gui:                g,
		pendingDeletes:     []*expanders.TreeNode{},
		toastNotifications: map[string]*eventing.StatusEvent{},
		client:             client,
	}

	newEvents := eventing.SubscribeToStatusEvents()
	// Start loop for showing loading in statusbar
	go func() {
		// recover from panic, if one occurrs, and leave terminal usable
		defer errorhandling.RecoveryWithCleanup()

		for {
			// Wait for a second to see if we have any new messages
			timeout := time.After(time.Second)
			select {
			case eventObjRaw := <-newEvents:
				eventObj := eventObjRaw.(*eventing.StatusEvent)
				if eventObj.IsToast {
					widget.toastNotifications[eventObj.ID()] = eventObj
				}
			case <-timeout:
				// Update the UI
			}
			for ID, toast := range widget.toastNotifications {
				if !toast.HasExpired() {
					delete(widget.toastNotifications, ID)
				}
			}
		}
	}()

	_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "bob doing things",
		InProgress: true,
		IsToast:    true,
		Timeout:    time.Second * 45,
	})

	done()

	return widget
}

// Layout draws the widget in the gocui view
func (w *NotificationWidget) Layout(g *gocui.Gui) error {
	// Don't draw anything if no pending deletes
	if len(w.pendingDeletes) < 1 && len(w.toastNotifications) < 1 {
		g.DeleteView(w.name)
		return nil
	}

	height := len(w.pendingDeletes) + len(w.toastNotifications)
	if len(w.pendingDeletes) > 0 {
		// Add padding for extra lines
		height = height + 7
	}
	if len(w.toastNotifications) > 0 {
		height = height + 3
	}

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
	toasts := w.toastNotifications

	if len(toasts) > 0 {
		fmt.Fprintln(v, style.Title("Notifications:"))
		for _, i := range toasts {
			fmt.Fprintln(v, i.Icon()+" "+i.Message)
		}
		fmt.Fprintln(v, "")
	}

	if len(pending) > 0 {
		fmt.Fprintln(v, style.Title("Pending Deletes:"))
		for _, i := range pending {
			fmt.Fprintln(v, " - "+i.Name)
		}
		fmt.Fprintln(v, "")
		fmt.Fprintln(v, "Do you want to delete these items?")
		fmt.Fprintln(v, style.Warning("Press "+strings.ToUpper(w.ConfirmDeleteKeyBinding)+" to DELETE"))
		fmt.Fprintln(v, style.Highlight("Press "+strings.ToUpper(w.ClearPendingDeletesKeyBinding)+" to CANCEL"))
	}

	return nil
}
