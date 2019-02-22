package keybindings

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/skratchdot/open-golang/open"
)

////////////////////////////////////////////////////////////////////
type ListActionsHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

func NewListActionsHandler(list *views.ListWidget, context context.Context) *ListActionsHandler {
	handler := &ListActionsHandler{
		Context: context,
		List:    list,
	}
	handler.Index = 7
	return handler
}

func (h ListActionsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return views.LoadActionsView(h.Context, h.List)
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListBackHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListBackHandler(list *views.ListWidget) *ListBackHandler {
	handler := &ListBackHandler{
		List: list,
	}
	handler.Index = 8
	return handler
}

func (h ListBackHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListBackLegacyHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListBackLegacyHandler(list *views.ListWidget) *ListBackLegacyHandler {
	handler := &ListBackLegacyHandler{
		List: list,
	}
	handler.Index = 9
	return handler
}

func (h ListBackLegacyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListDownHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListDownHandler(list *views.ListWidget) *ListDownHandler {
	handler := &ListDownHandler{
		List: list,
	}
	handler.Index = 10
	return handler
}

func (h ListDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() + 1)
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListUpHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListUpHandler(list *views.ListWidget) *ListUpHandler {
	handler := &ListUpHandler{
		List: list,
	}
	handler.Index = 11
	return handler
}

func (h ListUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ChangeSelection(h.List.CurrentSelection() - 1)
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListRightHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListRightHandler(list *views.ListWidget, editModeEnabled *bool) *ListRightHandler {
	handler := &ListRightHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.Index = 12
	return handler
}

func (h ListRightHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = true
		g.SetCurrentView("itemWidget")
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListEditHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListEditHandler(list *views.ListWidget, editModeEnabled *bool) *ListEditHandler {
	handler := &ListEditHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.Index = 13
	return handler
}

func (h ListEditHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.EditModeEnabled)
		h.EditModeEnabled = &tmp // memory leak?
		if *h.EditModeEnabled {
			g.Cursor = true
			g.SetCurrentView("itemWidget")
		} else {
			g.Cursor = false
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListExpandHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListExpandHandler(list *views.ListWidget) *ListExpandHandler {
	handler := &ListExpandHandler{
		List: list,
	}
	handler.Index = 14
	return handler
}

func (h ListExpandHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ExpandCurrentSelection()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListOpenHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

func NewListOpenHandler(list *views.ListWidget, context context.Context) *ListOpenHandler {
	handler := &ListOpenHandler{
		List:    list,
		Context: context,
	}
	handler.Index = 15
	return handler
}

func (h ListOpenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		portalURL := os.Getenv("AZURE_PORTAL_URL")
		if portalURL == "" {
			portalURL = "https://portal.azure.com"
		}
		url := portalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.ID
		span, _ := tracing.StartSpanFromContext(h.Context, "openportal:url")
		open.Run(url)
		span.Finish()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListRefreshHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListRefreshHandler(list *views.ListWidget) *ListRefreshHandler {
	handler := &ListRefreshHandler{
		List: list,
	}
	handler.Index = 16
	return handler
}

func (h ListRefreshHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.Refresh()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListDeleteHandler struct {
	ListHandler
	DeleteConfirmItemID string
	DeleteConfirmCount  int
	StatusBar           *views.StatusbarWidget
	Content             *views.ItemWidget
	List                *views.ListWidget
	Context             context.Context
}

func NewListDeleteHandler(content *views.ItemWidget,
	statusbar *views.StatusbarWidget,
	list *views.ListWidget,
	deleteConfirmItemId string,
	deleteConfirmCount int,
	context context.Context) *ListDeleteHandler {
	handler := &ListDeleteHandler{
		Content:             content,
		StatusBar:           statusbar,
		List:                list,
		DeleteConfirmCount:  deleteConfirmCount,
		DeleteConfirmItemID: deleteConfirmItemId,
		Context:             context,
	}
	handler.Index = 2
	return handler
}

func (h ListDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		if h.DeleteConfirmItemID != item.ID {
			h.DeleteConfirmItemID = item.ID
			h.DeleteConfirmCount = 0
		}
		keyBindings := GetKeyBindingsAsStrings()
		done := h.StatusBar.Status(fmt.Sprintf("Delete item? Really? PRESS %s TO CONFIRM: %s", strings.ToUpper(keyBindings["listdelete"]), item.DeleteURL), true)
		h.DeleteConfirmCount++

		if h.DeleteConfirmCount > 1 {
			done()
			doneDelete := h.StatusBar.Status("Deleting item: "+item.DeleteURL, true)

			h.DeleteConfirmItemID = ""

			// Run in the background
			go func() {
				res, err := armclient.DoRequest(h.Context, "DELETE", item.DeleteURL)
				if err != nil {
					panic(err)
				}
				// list.Refresh()
				h.Content.SetContent(res, "Delete response>"+item.Name)
				doneDelete()
			}()
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////
