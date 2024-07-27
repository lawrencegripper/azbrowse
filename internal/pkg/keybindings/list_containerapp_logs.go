package keybindings

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type CommandPanelContainerAppLogsHandler struct {
	ListHandler
	commandPanelWidget *views.CommandPanelWidget
	list               *views.ListWidget
	content            *views.ItemWidget
}

var _ Command = &CommandPanelContainerAppLogsHandler{}

func NewCommandPanelContainerAppLogsHandler(commandPanelWidget *views.CommandPanelWidget, content *views.ItemWidget, list *views.ListWidget) *CommandPanelContainerAppLogsHandler {
	handler := &CommandPanelContainerAppLogsHandler{
		commandPanelWidget: commandPanelWidget,
		content:            content,
		list:               list,
	}
	handler.id = HandlerIDContainerAppLogs

	return handler
}

func (h *CommandPanelContainerAppLogsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if h.IsEnabled() {
			return h.Invoke()
		}
		return nil
	}
}

func (h *CommandPanelContainerAppLogsHandler) DisplayText() string {
	return "Container Apps: Show logs"
}

func (h *CommandPanelContainerAppLogsHandler) IsEnabled() bool {
	currentItem := h.list.CurrentItem()
	if currentItem != nil && currentItem.Metadata["ContainerAppNodeType"] == expanders.ContainerAppNode_RevisionReplicaContainer {
		return true
	}
	return false
}

func (h *CommandPanelContainerAppLogsHandler) Invoke() error {

	currentItem := h.list.CurrentItem()
	logStreamEndpoint := currentItem.Metadata["LogStreamEndpoint"]
	if logStreamEndpoint == "" {
		return fmt.Errorf("no log stream endpoint found")
	}

	containerAppExpander, ok := (currentItem.Expander).(expanders.ContainerAppExpanderInterface)
	if !ok {
		return fmt.Errorf("current item is not a ContainerAppExpanderInterface")
	}

	ctx, cancel := context.WithCancel(context.Background())
	authToken, err := containerAppExpander.GetAuthToken(ctx, currentItem)
	if err != nil {
		cancel()
		return fmt.Errorf("failed to get auth token: %s", err)
	}

	_ = authToken

	go func() {
		// Wait for the user to navigate away
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")
		<-navigatedChannel
		// Clean up subscription
		eventing.Unsubscribe(navigatedChannel)
		// Cancel log context
		cancel()
	}()

	return h.getLogs(ctx, logStreamEndpoint, authToken)

}

func (h *CommandPanelContainerAppLogsHandler) getLogs(ctx context.Context, logStreamEndpoint string, authToken string) error {
	// TODO - actually get the logs!

	url, err := url.Parse(logStreamEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse log stream endpoint: %s", err)
	}
	query := url.Query()
	query.Add("output", "text")
	query.Add("follow", "true")
	query.Add("taillines", "50")
	url.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}
	request.Header.Set("Authorization", "Bearer "+authToken)

	httpClient := http.DefaultClient

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %s", err)
	}

	scanner := bufio.NewScanner(response.Body)
	go func() {
		defer response.Body.Close() //nolint: errcheck

		content := ""
		for scanner.Scan() {
			content += scanner.Text() + "\n"
			h.content.SetContent(content, interfaces.ResponsePlainText, "Logs")
		}

		content += "!!Connection closed"
		h.content.SetContent(content, interfaces.ResponsePlainText, "Logs")
	}()
	return nil
}
