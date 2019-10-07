package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// ResourceGroupResourceExpander expands resource under an RG
type ResourceGroupResourceExpander struct{}

// Name returns the name of the expander
func (e *ResourceGroupResourceExpander) Name() string {
	return "ResourceGroupExpander"
}

// DoesExpand checks if this is an RG
func (e *ResourceGroupResourceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == resourceGroupType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *ResourceGroupResourceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	span, ctx := tracing.StartSpanFromContext(ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	queryDoneChan := make(chan map[string]string)
	// Refactor this into DoResourceGraphQueryAync
	go func() {
		// Use resource graph to enrich response
		query := "where resourceGroup=='" + currentItem.Name + "' | project name, id, sku, kind, location, tags, properties.provisioningState"
		queryData, err := armclient.DoResourceGraphQuery(ctx, currentItem.SubscriptionID, query)
		span.SetTag("queryResponse", queryData)
		span.SetTag("queryError", err)
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: false,
				Failure:    true,
				Message:    "Getting query response: " + query + " " + err.Error(),
				Timeout:    time.Duration(time.Second * 4),
			})
		}

		queryValue, err := fastJSONParser.Parse(string(queryData))
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: false,
				Failure:    true,
				Message:    "Parsing query response: " + query,
				Timeout:    time.Duration(time.Second * 4),
			})
		}

		stateMap := map[string]string{}

		for _, row := range queryValue.Get("data").GetArray("rows") {
			span.SetTag("Row", row)
			rowValues, err := row.Array()
			if err != nil {
				panic(err)
			}
			currentState := string(rowValues[6].GetStringBytes())
			itemID := string(rowValues[1].GetStringBytes())
			stateMap[itemID] = currentState
		}

		queryDoneChan <- stateMap
		span.SetTag("stateMap", stateMap)
	}()

	// Add deployment item
	newItems := []*TreeNode{}
	newItems = append(newItems, &TreeNode{
		Parentid:       currentItem.ID,
		Namespace:      "None",
		Display:        style.Subtle("[Microsoft.Resources]") + "\n  Deployments",
		Name:           "Deployments",
		ID:             currentItem.ID + "/providers/Microsoft.Resources/deployments",
		ExpandURL:      currentItem.ID + "/providers/Microsoft.Resources/deployments?api-version=2017-05-10",
		ItemType:       deploymentsType,
		DeleteURL:      "",
		SubscriptionID: currentItem.SubscriptionID,
	})

	// Add Activity Log item
	newItems = append(newItems, &TreeNode{
		Parentid:       currentItem.ID,
		Namespace:      "None",
		Display:        style.Subtle("[Microsoft.Insights]") + "\n  Activity Log",
		Name:           "Activity Log",
		ID:             currentItem.ID + "/<activitylog>",
		ExpandURL:      GetActivityLogExpandURL(currentItem.SubscriptionID, currentItem.Name),
		ItemType:       activityLogType,
		DeleteURL:      "",
		SubscriptionID: currentItem.SubscriptionID,
	})

	// Get the latest from the ARM API
	method := "GET"
	responseChan := armclient.DoRequestAsync(ctx, method, currentItem.ExpandURL)

	stateMap := map[string]string{}
	armResponse := &armclient.RequestResult{}

	// Here be dragons.....
	// This block does the following:
	// 1. Wait for the ARM request (gets latest resource in the group) to complete
	// 2. Then give the GraphQuery (gets resource status) another second to complete
	// or give up on it.
	// This is because the status information is a value add and we don't want
	// to slow down browsing as a result of the graph query going slowly
	wg := &sync.WaitGroup{}
	wg.Add(2)
	timeoutGraphQuery := make(chan bool)
	go func() {
		result := <-responseChan
		armResponse = &result
		wg.Done()
		//Give the graph query an extra second to complete
		<-time.After(time.Second * 2)
		timeoutGraphQuery <- true
	}()

	// Give the graphQuery some time to complete or timeout
	select {
	case stateMap = <-queryDoneChan:
		wg.Done()
	case <-timeoutGraphQuery:
		span.SetTag("graphQueryTimedout", true)
		wg.Done()
	}
	wg.Wait()
	// .....

	err := armResponse.Error

	if err != nil {
		return ExpanderResult{
			Nodes:    nil,
			Response: armResponse.Result,
			Err:      fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}
	var resourceResponse armclient.ResourceResponse
	err = json.Unmarshal([]byte(armResponse.Result), &resourceResponse)
	if err != nil {

		return ExpanderResult{
			Nodes:             nil,
			Response:          armResponse.Result,
			IsPrimaryResponse: true,
			Err:               fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}

	for _, rg := range resourceResponse.Resources {
		resourceAPIVersion, err := armclient.GetAPIVersion(rg.Type)
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				Failure: true,
				Message: "Failed to get resouceVersion for the Type:" + rg.Type,
				Timeout: time.Duration(time.Second * 5),
			})
		}
		item := &TreeNode{
			Display:          style.Subtle("["+rg.Type+"] \n  ") + rg.Name,
			Name:             rg.Name,
			Parentid:         currentItem.ID,
			Namespace:        getNamespaceFromARMType(rg.Type), // We just want the namespace not the subresource
			ArmType:          rg.Type,
			ID:               rg.ID,
			ExpandURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			ExpandReturnType: "none",
			ItemType:         ResourceType,
			DeleteURL:        rg.ID + "?api-version=" + resourceAPIVersion,
			SubscriptionID:   currentItem.SubscriptionID,
		}

		state, exists := stateMap[item.ID]
		if exists {
			item.StatusIndicator = DrawStatus(state)
		}

		newItems = append(newItems, item)
	}

	return ExpanderResult{
		Nodes:             newItems,
		Response:          armResponse.Result,
		SourceDescription: "Resources Request",
		IsPrimaryResponse: true,
	}
}
