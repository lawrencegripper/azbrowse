package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

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
