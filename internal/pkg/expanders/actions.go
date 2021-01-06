package expanders

import (
	"context"
	"fmt"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

type expanderAndListActionsResponse struct {
	Expander          Expander
	ListActionsResult ListActionsResult
}

// GetActions returns the available actions for a TreeNode
func GetActions(ctx context.Context, item *TreeNode) ([]*TreeNode, error) {
	newItems := []*TreeNode{}
	statusEvent, _ := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Getting available actions",
		InProgress: true,
	})
	defer statusEvent.Done()

	// TODO - working here: copy ExpandNodes pattern and invoke expanders to get actions

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+item.ItemType+":"+item.Name, tracing.SetTag("item", item))
	defer span.Finish()

	// New handler approach
	handlerExpanding := 0

	completedExpands := make(chan expanderAndListActionsResponse)

	// Check which expanders are interested and kick them off
	spanQuery, _ := tracing.StartSpanFromContext(ctx, "querexpanders", tracing.SetTag("item", item))
	expanders := append(getRegisteredExpanders(), GetDefaultExpander())
	for _, h := range expanders {
		doesExpand, err := h.HasActions(ctx, item)
		spanQuery.SetTag(h.Name(), doesExpand)
		if err != nil {
			panic(err)
		}
		if !doesExpand {
			continue
		}

		// Fire each handler in parallel
		hCurrent := h // capture current iteration variable
		go func() {
			// recover from panic, if one occurrs, and leave terminal usable
			defer errorhandling.RecoveryWithCleanup()

			completedExpands <- expanderAndListActionsResponse{
				Expander:          hCurrent,
				ListActionsResult: hCurrent.ListActions(ctx, item),
			}
		}()

		handlerExpanding = handlerExpanding + 1
	}
	spanQuery.Finish()

	// Lets give all the expanders 45secs to completed (unless debugging)
	hasPrimaryResponse := false
	timeout := time.After(time.Second * 45)

	for index := 0; index < handlerExpanding; index++ {
		select {
		case done := <-completedExpands:
			result := done.ListActionsResult
			span, _ := tracing.StartSpanFromContext(ctx, "subexpand:"+result.SourceDescription, tracing.SetTag("result", done))
			// Did it fail?
			if result.Err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					Failure: true,
					Message: "Expander '" + result.SourceDescription + "' failed on resource: " + item.ID + "Err: " + result.Err.Error(),
					Timeout: time.Duration(time.Second * 15),
				})
			}
			if result.IsPrimaryResponse {
				if hasPrimaryResponse {
					panic(fmt.Sprintf("Two handlers returned a primary response for this item... failing. ID: %s EXPANDER: %s", item.ID, result.SourceDescription))
				}
				// Log that we have a primary response
				hasPrimaryResponse = true
			}
			if result.Nodes == nil {
				continue
			}
			for _, node := range result.Nodes {
				node.Expander = done.Expander
				node.ItemType = ActionType
			}
			// Add the items it found
			if result.IsPrimaryResponse {
				newItems = append(result.Nodes, newItems...)
			} else {
				newItems = append(newItems, result.Nodes...)
			}
			span.Finish()
		case <-timeout:
			eventing.SendStatusEvent(&eventing.StatusEvent{
				Failure: true,
				Message: "Timed out opening:" + item.ID,
				Timeout: time.Duration(time.Second * 10),
			})
			return nil, fmt.Errorf("Timed out opening: %s", item.ID)
		}
	}
	return newItems, nil

	// var namespace string
	// var armType string

	// if item.ItemType == ResourceType {
	// 	namespace = item.Namespace
	// 	armType = item.ArmType
	// }

	// if namespace == "" || armType == "" {
	// 	return []*TreeNode{}, nil
	// }

	// span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+item.Name, tracing.SetTag("item", item))
	// defer span.Finish()

	// data, err := armclient.LegacyInstance.DoRequest(ctx, "GET", "/providers/Microsoft.Authorization/providerOperations/"+namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	// if err != nil {
	// 	return []*TreeNode{}, fmt.Errorf("Failed to get actions: %s", err)
	// }
	// var opsRequest OperationsRequest
	// err = json.Unmarshal([]byte(data), &opsRequest)
	// if err != nil {
	// 	panic(err)
	// }

	// items := []*TreeNode{}
	// for _, resOps := range opsRequest.ResourceTypes {
	// 	if resOps.Name == strings.Split(armType, "/")[1] {
	// 		for _, op := range resOps.Operations {
	// 			resourceAPIVersion, err := armclient.GetAPIVersion(item.ArmType)
	// 			if err != nil {
	// 				return []*TreeNode{}, fmt.Errorf("Failed to find an api version: %s", err)
	// 			}
	// 			stripArmType := strings.Replace(op.Name, item.ArmType, "", -1)
	// 			actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + resourceAPIVersion
	// 			items = append(items, &TreeNode{
	// 				Name:             op.DisplayName,
	// 				Display:          op.DisplayName,
	// 				ExpandURL:        item.ID + "/" + actionURL,
	// 				ExpandReturnType: ActionType,
	// 				ItemType:         "action",
	// 				ID:               item.ID + "/" + actionURL,
	// 			})
	// 		}
	// 	}
	// }
	//
	// return items, nil
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
