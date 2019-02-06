package handlers

import (
	"context"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/eventing"
)

// DefaultExpander expands RGs under a subscription
type DefaultExpander struct{}

// Name returns the name of the expander
func (e *DefaultExpander) Name() string {
	return "DefaultExpander"
}

// DoesExpand checks if this is an RG
func (e *DefaultExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
	return true, nil
}

// Expand returns Resources in the RG
func (e *DefaultExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {
	method := "GET"

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
	}
}
