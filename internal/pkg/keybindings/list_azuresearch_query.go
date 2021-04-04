package keybindings

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type CommandPanelAzureSearchQueryHandler struct {
	ListHandler
	commandPanelWidget *views.CommandPanelWidget
	list               *views.ListWidget
	content            *views.ItemWidget
}

var _ Command = &CommandPanelAzureSearchQueryHandler{}

func NewCommandPanelAzureSearchQueryHandler(commandPanelWidget *views.CommandPanelWidget, content *views.ItemWidget, list *views.ListWidget) *CommandPanelAzureSearchQueryHandler {
	handler := &CommandPanelAzureSearchQueryHandler{
		commandPanelWidget: commandPanelWidget,
		content:            content,
		list:               list,
	}
	handler.id = HandlerIDAzureSearchQuery

	return handler
}

func (h *CommandPanelAzureSearchQueryHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if h.IsEnabled() {
			return h.Invoke()
		}
		return nil
	}
}

func (h *CommandPanelAzureSearchQueryHandler) DisplayText() string {
	return "Azure search query"
}

func (h *CommandPanelAzureSearchQueryHandler) IsEnabled() bool {
	currentItem := h.list.CurrentExpandedItem()
	if currentItem != nil && currentItem.SwaggerResourceType != nil && currentItem.SwaggerResourceType.Endpoint.TemplateURL == "/indexes('{indexName}')/docs" {
		return true
	}
	return false
}

func (h *CommandPanelAzureSearchQueryHandler) Invoke() error {
	h.commandPanelWidget.ShowWithText("search query:", "search=", nil, h.CommandPanelNotification)
	return nil
}

func (h *CommandPanelAzureSearchQueryHandler) CommandPanelNotification(state interfaces.CommandPanelNotification) {

	if state.EnterPressed {
		queryString := state.CurrentText
		currentItem := h.list.CurrentExpandedItem()

		apiSetID := currentItem.Metadata["SwaggerAPISetID"]
		apiSetPtr := expanders.GetSwaggerResourceExpander().GetAPISet(apiSetID)
		if apiSetPtr == nil {
			return
		}
		apiSet := *apiSetPtr
		searchApiSet := apiSet.(expanders.SwaggerAPISetSearch)

		data, err := searchApiSet.DoRequest("GET", currentItem.ExpandURL+"&"+queryString)
		if err != nil {
			h.content.SetContent(fmt.Sprintf("%s", err), interfaces.ResponseJSON, queryString)
		} else {
			h.content.SetContent(data, interfaces.ResponseJSON, queryString)
		}
	}
}
