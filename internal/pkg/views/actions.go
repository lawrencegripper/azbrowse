package views

import (
	"context"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
)

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
	if len(actionItems) > 0 {
		list.SetNewNodes(actionItems)
	}

	return nil
}
