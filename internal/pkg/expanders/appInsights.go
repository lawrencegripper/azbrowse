package expanders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type analyticsItem struct {
	Content      string `json:"Content"`
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Scope        string `json:"Scope"`
	TimeCreated  string `json:"TimeCreated"`
	TimeModified string `json:"TimeModified"`
	Type         string `json:"Type"`
	Version      string `json:"Version"`
}

// Check interface
var _ Expander = &AppInsightsExpander{}

// AppInsightsExpander expands aspects of App Insights that don't naturally flow from the api spec
type AppInsightsExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *AppInsightsExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *AppInsightsExpander) Name() string {
	return "AppInsightsExpander"
}

// DoesExpand checks if this is a node we should extend
func (e *AppInsightsExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "AppInsights" {
		return true, nil
	}
	return false, nil
}

// Expand returns nodes for App Insights
func (e *AppInsightsExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "AppInsights" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}" {
		newItems := []*TreeNode{}
		resourceAPIVersion, err := armclient.GetAPIVersion(currentItem.ArmType)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				Nodes:             newItems,
				SourceDescription: "AppInsightsExpander Request",
				IsPrimaryResponse: true,
			}
		}

		newItems = append(newItems, &TreeNode{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/analyticsItems",
			Namespace:             "AppInsights",
			Name:                  "Analytics Items",
			Display:               "Analytics Items",
			ItemType:              "AppInsights.AnalyticsItems",
			ExpandURL:             currentItem.ID + "/analyticsItems?api-version=" + resourceAPIVersion,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"ResourceAPIVersion": resourceAPIVersion,
				"AppInsightsID":      currentItem.ID,
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "AppInsightsExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.ItemType == "AppInsights.AnalyticsItems" {
		return e.expandAnalyticsItems(ctx, currentItem)
	} else if currentItem.ItemType == "AppInsights.AnalyticsItem" {
		return e.expandAnalyticsItem(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "AppInsightsExpander request",
	}
}

func (e *AppInsightsExpander) expandAnalyticsItems(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	newItems := []*TreeNode{}

	data, err := e.client.DoRequest(ctx, "GET", currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Nodes:             newItems,
			SourceDescription: "AppInsightsExpander Request",
			IsPrimaryResponse: true,
		}
	}

	resourceAPIVersion := currentItem.Metadata["ResourceAPIVersion"]
	appInsightsID := currentItem.Metadata["AppInsightsID"]

	var items []analyticsItem
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling analytics items response: %s, %s", err, data)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "AppInsightsExpander request",
		}
	}

	for _, item := range items {
		var collectionName string
		if item.Scope == "user" {
			collectionName = "myanalyticsItems"
		} else {
			collectionName = "analyticsItems"
		}
		newItem := TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "AppInsights",
			ItemType:  "AppInsights.AnalyticsItem",
			Name:      item.Name,
			ExpandURL: appInsightsID + "/" + collectionName + "/item?api-version=" + resourceAPIVersion + "&id=" + item.ID,
			DeleteURL: appInsightsID + "/" + collectionName + "/item?api-version=" + resourceAPIVersion + "&id=" + item.ID,
			Display:   style.Subtle("["+item.Type+" - "+item.Scope+"]") + "\n " + item.Name,
		}
		newItems = append(newItems, &newItem)
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "AppInsightsExpander request",
	}
}

func (e *AppInsightsExpander) expandAnalyticsItem(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	newItems := []*TreeNode{}

	data, err := e.client.DoRequest(ctx, "GET", currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Nodes:             newItems,
			SourceDescription: "AppInsightsExpander Request",
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "AppInsightsExpander request",
	}
}
