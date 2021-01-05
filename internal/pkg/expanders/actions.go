package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// GetActions returns the available actions for a TreeNode
func GetActions(ctx context.Context, item *TreeNode) ([]*TreeNode, error) {
	statusEvent, _ := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Getting available actions",
		InProgress: true,
	})
	defer statusEvent.Done()

	var namespace string
	var armType string

	if item.ItemType == ResourceType {
		namespace = item.Namespace
		armType = item.ArmType
	}

	if namespace == "" || armType == "" {
		return []*TreeNode{}, nil
	}

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+item.Name, tracing.SetTag("item", item))
	defer span.Finish()

	data, err := armclient.LegacyInstance.DoRequest(ctx, "GET", "/providers/Microsoft.Authorization/providerOperations/"+namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		return []*TreeNode{}, fmt.Errorf("Failed to get actions: %s", err)
	}
	var opsRequest OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	items := []*TreeNode{}
	for _, resOps := range opsRequest.ResourceTypes {
		if resOps.Name == strings.Split(armType, "/")[1] {
			for _, op := range resOps.Operations {
				resourceAPIVersion, err := armclient.GetAPIVersion(item.ArmType)
				if err != nil {
					return []*TreeNode{}, fmt.Errorf("Failed to find an api version: %s", err)
				}
				stripArmType := strings.Replace(op.Name, item.ArmType, "", -1)
				actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + resourceAPIVersion
				items = append(items, &TreeNode{
					Name:             op.DisplayName,
					Display:          op.DisplayName,
					ExpandURL:        item.ID + "/" + actionURL,
					ExpandReturnType: ActionType,
					ItemType:         "action",
					ID:               item.ID + "/" + actionURL,
				})
			}
		}
	}

	return items, nil
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
