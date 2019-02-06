package handlers

import (
	"context"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/eventing"
)

// ActionExpander handles actions
type ActionExpander struct{}

// Name returns the name of the expander
func (e *ActionExpander) Name() string {
	return "ActionExpander"
}

// DoesExpand checks if it is an action
func (e *ActionExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
	if currentItem.ItemType == ActionType {
		return true, nil
	}

	return false, nil
}

// Expand performs the action
func (e *ActionExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {
	method := "POST"

	status, done := eventing.SendStatusEvent(eventing.StatusEvent{
		InProgress: true,
		Message:    "Requestion Resource Groups for subscription",
	})
	defer done()

	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		status.Message = "Failed" + err.Error() + currentItem.ExpandURL
		status.Update()
	}

	return ExpanderResult{
		Err:               err,
		Response:          string(data),
		SourceDescription: "Resource Group Request",
		IsPrimaryResponse: true,
	}
}
