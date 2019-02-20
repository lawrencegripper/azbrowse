package keyhandlers

import (
	"context"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// HACK: To prevent accidental deletes this method requires del to be pressed twice on a resource
// before it will proceed

const deleteId = 2

type DeleteHandler struct {
	DeleteConfirmItemID string
	DeleteConfirmCount  int
	StatusBar           *views.StatusbarWidget
	Content             *views.ItemWidget
	List                *views.ListWidget
	Context             context.Context
}

func (h DeleteHandler) Id() string {
	return HandlerIds[deleteId]
}

func (h DeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		if h.DeleteConfirmItemID != item.ID {
			h.DeleteConfirmItemID = item.ID
			h.DeleteConfirmCount = 0
		}
		done := h.StatusBar.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.DeleteURL, true)
		h.DeleteConfirmCount++

		if h.DeleteConfirmCount > 1 {
			done()
			doneDelete := h.StatusBar.Status("Deleting item: "+item.DeleteURL, true)

			res, err := armclient.DoRequest(h.Context, "DELETE", item.DeleteURL)
			if err != nil {
				panic(err)
			}
			h.List.Refresh()
			h.Content.SetContent(res, "Delete response>"+item.Name)
			doneDelete()
			h.DeleteConfirmItemID = ""
		}
		return nil
	}
}

func (h DeleteHandler) Widget() string {
	return ""
}

func (h DeleteHandler) DefaultKey() gocui.Key {
	return gocui.KeyDelete
}
