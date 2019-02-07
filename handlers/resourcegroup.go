package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lawrencegripper/azbrowse/eventing"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/style"
	"github.com/lawrencegripper/azbrowse/tracing"
	"github.com/valyala/fastjson"
)

var fastJSONParser fastjson.Parser

func init() {
	fastJSONParser = fastjson.Parser{}
}

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

	queryDoneChan := make(chan map[string]string)
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
	newItems := []TreeNode{}
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

	// Get the latest from the ARM API
	method := "GET"
	responseChan := armclient.DoRequestAsync(ctx, method, currentItem.ExpandURL)

	stateMap := <-queryDoneChan
	armResponse := <-responseChan

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
			panic(err)
		}
		item := TreeNode{
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
		}

		state, exists := stateMap[item.ID]
		if exists {
			item.Display = item.Display + " " + drawStatus(state)
		}

		newItems = append(newItems, item)
	}

	return ExpanderResult{
		Nodes:             &newItems,
		Response:          armResponse.Result,
		SourceDescription: "Resources Request",
		IsPrimaryResponse: true,
	}
}
