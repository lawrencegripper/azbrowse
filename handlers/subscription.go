package handlers

import (
	"context"
	"encoding/json"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/eventing"
)

// SubscriptionExpander expands RGs under a subscription
type SubscriptionExpander struct{}

// Name returns the name of the expander
func (e *SubscriptionExpander) Name() string {
	return "SubscriptionExpander"
}

// DoesExpand checks if this is an RG
func (e *SubscriptionExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
	if currentItem.ItemType == subscriptionType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *SubscriptionExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {
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
	var rgResponse armclient.ResourceGroupResponse
	err = json.Unmarshal([]byte(data), &rgResponse)
	if err != nil {
		panic(err)
	}

	newItems := []TreeNode{}
	for _, rg := range rgResponse.Groups {
		newItems = append(newItems, TreeNode{
			Name:             rg.Name,
			Display:          rg.Name + " " + drawStatus(rg.Properties.ProvisioningState),
			ID:               rg.ID,
			Parentid:         currentItem.ID,
			ExpandURL:        rg.ID + "/resources?api-version=2017-05-10",
			ExpandReturnType: resourceType,
			ItemType:         resourceGroupType,
			DeleteURL:        rg.ID + "?api-version=2017-05-10",
		})
	}

	return ExpanderResult{
		Err:               err,
		Nodes:             &newItems,
		Response:          string(data),
		SourceDescription: "Resource Group Request",
	}
}
