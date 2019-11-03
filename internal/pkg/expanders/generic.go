package expanders

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// DefaultExpander expands RGs under a subscription
type DefaultExpander struct{}

// DefaultExpanderInstance provides an instance of the default expander for use
var DefaultExpanderInstance DefaultExpander

// Name returns the name of the expander
func (e *DefaultExpander) Name() string {
	return "GenericExpander"
}

// DoesExpand checks if this handler can expand this item
func (e *DefaultExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ExpandURL == ExpandURLNotSupported || currentItem.Metadata["SuppressGenericExpand"] == "true" {
		return false, nil
	}
	return true, nil
}

// Expand returns resource
func (e *DefaultExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"

	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
			SourceDescription: "Resource Group Request",
		}
	}

	var resource armclient.Resource
	err = json.Unmarshal([]byte(data), &resource)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
			SourceDescription: "Resource Group Request",
		}
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
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		SourceDescription: "Resource Group Request",
	}
}
