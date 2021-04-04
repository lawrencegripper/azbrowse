package keybindings

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type ListSortHandler struct {
	ListHandler
	List *views.ListWidget
}

var _ Command = &ListSortHandler{}

func NewListSortHandler(list *views.ListWidget) *ListSortHandler {
	handler := &ListSortHandler{
		List: list,
	}
	handler.id = HandlerIDListSort
	return handler
}

func (h ListSortHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListSortHandler) DisplayText() string {
	return "Sort list items"
}
func (h *ListSortHandler) IsEnabled() bool {
	return true
}
func (h *ListSortHandler) Invoke() error {
	h.List.SortItems()
	return nil
}
