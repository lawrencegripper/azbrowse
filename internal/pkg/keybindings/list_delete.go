package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListDeleteHandler struct {
	ListHandler
	List               *views.ListWidget
	NotificationWidget *views.NotificationWidget
}

func NewListDeleteHandler(list *views.ListWidget, notificationWidget *views.NotificationWidget) *ListDeleteHandler {
	handler := &ListDeleteHandler{
		List:               list,
		NotificationWidget: notificationWidget,
	}
	handler.id = HandlerIDListDelete
	return handler
}

func (h ListDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		h.NotificationWidget.AddPendingDelete(item)
		return nil
	}
}
