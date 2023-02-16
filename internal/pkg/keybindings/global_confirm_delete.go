package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

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
