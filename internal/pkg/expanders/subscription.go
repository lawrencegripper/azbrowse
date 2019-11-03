package expanders

import (
	"context"
	"encoding/json"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Check interface
var _ Expander = &SubscriptionExpander{}

// SubscriptionExpander expands RGs under a subscription
type SubscriptionExpander struct{}

// Name returns the name of the expander
func (e *SubscriptionExpander) Name() string {
	return "SubscriptionExpander"
}

// DoesExpand checks if this is an RG
func (e *SubscriptionExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == SubscriptionType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *SubscriptionExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"

	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	newItems := []*TreeNode{}

	//    \/ It's not the usual ... look out
	if err == nil {
		var rgResponse armclient.ResourceGroupResponse
		err = json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, &TreeNode{
				Name:             rg.Name,
				Display:          rg.Name,
				ID:               rg.ID,
				Parentid:         currentItem.ID,
				ExpandURL:        rg.ID + "/resources?api-version=2017-05-10",
				ExpandReturnType: ResourceType,
				ItemType:         resourceGroupType,
				DeleteURL:        rg.ID + "?api-version=2017-05-10",
				SubscriptionID:   currentItem.SubscriptionID,
				StatusIndicator:  DrawStatus(rg.Properties.ProvisioningState),
			})
		}
	}

	return ExpanderResult{
		Err:               err,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		SourceDescription: "Resource Group Request",
		IsPrimaryResponse: true,
	}
}
