package keybindings

import (
	"context"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListActionsHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

var _ Command = &ListActionsHandler{}

func NewListActionsHandler(list *views.ListWidget, context context.Context) *ListActionsHandler {
	handler := &ListActionsHandler{
		Context: context,
		List:    list,
	}
	handler.id = HandlerIDListActions
	return handler
}

func (h ListActionsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}
func (h *ListActionsHandler) DisplayText() string {
	return "Show Actions"
}
func (h *ListActionsHandler) IsEnabled() bool {
	return h.List.CurrentExpandedItem() != nil
}
func (h *ListActionsHandler) Invoke() error {
	return views.LoadActionsView(h.Context, h.List)
}
