package views

import (
	"context"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
)

// TODO: Migrate to a handler...maybe

// LoadActionsView Shows available actions for the current resource
func LoadActionsView(ctx context.Context, list *ListWidget) error {
	statusEvent, _ := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Getting available actions",
		InProgress: true,
	})
	defer statusEvent.Done()

	currentItem := list.CurrentItem()

	actionItems, err := expanders.GetActions(ctx, currentItem)
	if err != nil {
		list.statusView.Status("Failed to get actions:"+err.Error(), false)
		return err
	}
	if len(actionItems) > 1 {
		list.SetNewNodes(actionItems)
	}

	return nil
}

// OperationsRequest list the actions that can be performed
type OperationsRequest struct {
	DisplayName string `json:"displayName"`
	Operations  []struct {
		Name         string      `json:"name"`
		DisplayName  string      `json:"displayName"`
		Description  string      `json:"description"`
		Origin       interface{} `json:"origin"`
		Properties   interface{} `json:"properties"`
		IsDataAction bool        `json:"isDataAction"`
	} `json:"operations"`
	ResourceTypes []struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Operations  []struct {
			Name         string      `json:"name"`
			DisplayName  string      `json:"displayName"`
			Description  string      `json:"description"`
			Origin       interface{} `json:"origin"`
			Properties   interface{} `json:"properties"`
			IsDataAction bool        `json:"isDataAction"`
		} `json:"operations"`
	} `json:"resourceTypes"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
