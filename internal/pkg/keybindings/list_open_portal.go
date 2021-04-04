package keybindings

import (
	"context"
	"os"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/internal/pkg/wsl"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/skratchdot/open-golang/open"
)

type ListOpenHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

var _ Command = &ListOpenHandler{}

func NewListOpenHandler(list *views.ListWidget, context context.Context) *ListOpenHandler {
	handler := &ListOpenHandler{
		List:    list,
		Context: context,
	}
	handler.id = HandlerIDListOpen
	return handler
}

func (h ListOpenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListOpenHandler) DisplayText() string {
	return "Open in Portal"
}
func (h *ListOpenHandler) IsEnabled() bool {
	return true // TODO - filter to Azure resource nodes
}
func (h *ListOpenHandler) Invoke() error {
	item := h.List.CurrentItem()
	portalURL := os.Getenv("AZURE_PORTAL_URL")
	if portalURL == "" {
		portalURL = "https://portal.azure.com"
	}
	url := portalURL + "/#@" + armclient.LegacyInstance.GetTenantID() + "/resource/" + item.ID
	span, _ := tracing.StartSpanFromContext(h.Context, "openportal:url")
	var err error
	if wsl.IsWSL() {
		err = wsl.TryLaunchBrowser(url)
	} else {
		err = open.Run(url)
	}
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed opening resources in browser: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
		return nil
	}
	span.Finish()
	return nil
}
