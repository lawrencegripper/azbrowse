package keyhandlers

import (
	"context"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/skratchdot/open-golang/open"
)

const listOpenId = "ListOpen"

type ListOpenHandler struct {
	List    *views.ListWidget
	Context context.Context
}

func (h ListOpenHandler) Id() string {
	return listOpenId
}

func (h ListOpenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		portalURL := os.Getenv("AZURE_PORTAL_URL")
		if portalURL == "" {
			portalURL = "https://portal.azure.com"
		}
		url := portalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.Parentid + "/overview"
		span, _ := tracing.StartSpanFromContext(h.Context, "openportal:url")
		open.Run(url)
		span.Finish()
		return nil
	}
}

func (h ListOpenHandler) Widget() string {
	return "listWidget"
}

func (h ListOpenHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlO
}
