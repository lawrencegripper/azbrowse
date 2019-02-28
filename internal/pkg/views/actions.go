package views

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/handlers"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// TODO: Migrate to a handler...maybe

// LoadActionsView Shows available actions for the current resource
func LoadActionsView(ctx context.Context, list *ListWidget) error {
	list.statusView.Status("Getting available Actions", true)

	var namespace string
	var armType string
	currentItem := list.CurrentItem()
	if currentItem.ItemType == handlers.ResourceType {
		namespace = currentItem.Namespace
		armType = currentItem.ArmType
	}

	currentExpandedItem := list.CurrentExpandedItem()
	if currentExpandedItem.ItemType == handlers.ResourceType {
		namespace = currentExpandedItem.Namespace
		armType = currentExpandedItem.ArmType
	}

	if namespace == "" || armType == "" {
		return nil
	}

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	data, err := armclient.DoRequest(ctx, "GET", "/providers/Microsoft.Authorization/providerOperations/"+namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		list.statusView.Status("Failed to get actions: "+err.Error(), false)
	}
	var opsRequest armclient.OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	items := []*handlers.TreeNode{}
	for _, resOps := range opsRequest.ResourceTypes {
		if resOps.Name == strings.Split(armType, "/")[1] {
			for _, op := range resOps.Operations {
				resourceAPIVersion, err := armclient.GetAPIVersion(currentItem.ArmType)
				if err != nil {
					list.statusView.Status("Failed to find an api version: "+err.Error(), false)
				}
				stripArmType := strings.Replace(op.Name, currentItem.ArmType, "", -1)
				actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + resourceAPIVersion
				items = append(items, &handlers.TreeNode{
					Name:             op.DisplayName,
					Display:          op.DisplayName,
					ExpandURL:        currentItem.ID + "/" + actionURL,
					ExpandReturnType: handlers.ActionType,
					ItemType:         "action",
					ID:               currentItem.ID,
				})
			}
		}
	}
	if len(items) > 1 {
		list.SetNodes(items)
	}
	list.statusView.Status("Fetched available Actions", false)

	return nil
}
