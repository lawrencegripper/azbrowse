package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/style"
	"github.com/lawrencegripper/azbrowse/tracing"
)

// ResourceGroupResourceExpander expands resource under an RG
type ResourceGroupResourceExpander struct{}

// Name returns the name of the expander
func (e *ResourceGroupResourceExpander) Name() string {
	return "ResourceGroupExpander"
}

// DoesExpand checks if this is an RG
func (e *ResourceGroupResourceExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
	if currentItem.ItemType == resourceGroupType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *ResourceGroupResourceExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {

	span, ctx := tracing.StartSpanFromContext(ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	query := "where resourceGroup=='" + currentItem.Name + "' | project name, id, sku, kind, location, tags, properties.provisioningState"
	data, err := armclient.DoResourceGraphQuery(ctx, currentItem.SubscriptionID, query)
	if err != nil {
		return ExpanderResult{
			Nodes:             nil,
			Response:          currentItem.Name + query + string(data),
			IsPrimaryResponse: true,
			Err:               fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}

	var resourceResponse []armclient.Resource
	err = json.Unmarshal([]byte(data), &resourceResponse)
	if err != nil {

		return ExpanderResult{
			Nodes:             nil,
			Response:          string(data),
			IsPrimaryResponse: true,
			Err:               fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}

	newItems := []TreeNode{}
	// Add Deployments
	if currentItem.ItemType == resourceGroupType {
		newItems = append(newItems, TreeNode{
			Parentid:         currentItem.ID,
			Namespace:        "None",
			Display:          style.Subtle("[Microsoft.Resources]") + "\n  Deployments",
			Name:             "Deployments",
			ID:               currentItem.ID,
			ExpandURL:        currentItem.ID + "/providers/Microsoft.Resources/deployments?api-version=2017-05-10",
			ExpandReturnType: deploymentType,
			ItemType:         resourceType,
			DeleteURL:        "NotSupported",
			SubscriptionID:   currentItem.SubscriptionID,
		})
	}
	for _, rg := range resourceResponse {
		resourceAPIVersion, err := armclient.GetAPIVersion(rg.Type)
		if err != nil {
			panic(err)
		}
		newItems = append(newItems, TreeNode{
			Display:          style.Subtle("["+rg.Type+"] \n  ") + rg.Name,
			Name:             rg.Name,
			Parentid:         currentItem.ID,
			Namespace:        strings.Split(rg.Type, "/")[0], // We just want the namespace not the subresource
			ArmType:          rg.Type,
			ID:               rg.ID,
			ExpandURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			ExpandReturnType: "none",
			ItemType:         resourceType,
			DeleteURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			SubscriptionID:   currentItem.SubscriptionID,
		})
	}

	return ExpanderResult{
		Nodes:             &newItems,
		Response:          currentItem.Name + query + string(data),
		SourceDescription: "Resources Request",
		IsPrimaryResponse: true,
	}
}
