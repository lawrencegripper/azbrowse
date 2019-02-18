package handlers

import (
	"context"
	"encoding/json"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"time"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// DefaultExpander expands RGs under a subscription
type DefaultExpander struct{}

// Name returns the name of the expander
func (e *DefaultExpander) Name() string {
	return "DefaultExpander"
}

// DoesExpand checks if this is an RG
func (e *DefaultExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	return true, nil
}

// Expand returns Resources in the RG
func (e *DefaultExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"

	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)

	var resource armclient.Resource
	err = json.Unmarshal([]byte(data), &resource)
	if err != nil {
		panic(err)
	}

	// Update the existing state as we have more up-to-date info
	newStatus := DrawStatus(resource.Properties.ProvisioningState)
	if newStatus != currentItem.StatusIndicator {
		eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: true,
			Message:    "Updated resource status -> " + DrawStatus(resource.Properties.ProvisioningState),
			Timeout:    time.Duration(time.Second * 3),
		})
	}

	return ExpanderResult{
		Err:               err,
		Response:          string(data),
		SourceDescription: "Resource Group Request",
	}
}
