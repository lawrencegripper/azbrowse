package views

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
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
	if currentItem.ItemType == expanders.ResourceType {
		namespace = currentItem.Namespace
		armType = currentItem.ArmType
	}

	currentExpandedItem := list.CurrentExpandedItem()
	if currentExpandedItem.ItemType == expanders.ResourceType {
		namespace = currentExpandedItem.Namespace
		armType = currentExpandedItem.ArmType
	}

	if namespace == "" || armType == "" {
		return nil
	}

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	data, err := armclient.LegacyInstance.DoRequest(ctx, "GET", "/providers/Microsoft.Authorization/providerOperations/"+namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		list.statusView.Status("Failed to get actions: "+err.Error(), false)
	}
	var opsRequest OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	items := []*expanders.TreeNode{}
	for _, resOps := range opsRequest.ResourceTypes {
		if resOps.Name == strings.Split(armType, "/")[1] {
			for _, op := range resOps.Operations {
				resourceAPIVersion, err := armclient.GetAPIVersion(currentItem.ArmType)
				if err != nil {
					list.statusView.Status("Failed to find an api version: "+err.Error(), false)
				}
				stripArmType := strings.Replace(op.Name, currentItem.ArmType, "", -1)
				actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + resourceAPIVersion
				items = append(items, &expanders.TreeNode{
					Name:             op.DisplayName,
					Display:          op.DisplayName,
					ExpandURL:        currentItem.ID + "/" + actionURL,
					ExpandReturnType: expanders.ActionType,
					ItemType:         "action",
					ID:               currentItem.ID + "/" + actionURL,
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
