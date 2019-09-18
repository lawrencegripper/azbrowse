package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type AnalyticsItem struct {
	Content      string `json:"Content"`
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	Scope        string `json:"Scope"`
	TimeCreated  string `json:"TimeCreated"`
	TimeModified string `json:"TimeModified"`
	Type         string `json:"Type"`
	Version      string `json:"Version"`
}

// AppInsightsExpander expands aspects of App Insights that don't naturally flow from the api spec
type AppInsightsExpander struct {
	client *http.Client
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
			Parentid:  currentItem.ID,
			Namespace: "AppInsights",
			Name:      "Analytics Items",
			Display:   "Analytics Items",
			ItemType:  "AppInsights.AnalyticsItems",
			ExpandURL: currentItem.ID + "/analyticsItems?api-version=" + resourceAPIVersion,
			Metadata: map[string]string{
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          "", // Swagger expander will supply the response
			SourceDescription: "AppInsightsExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.ItemType == "AppInsights.AnalyticsItems" {
		return e.expandAnalyticsItems(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          "Error!",
		SourceDescription: "AppInsightsExpander request",
	}
}

func (e *AppInsightsExpander) expandAnalyticsItems(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	newItems := []*TreeNode{}
	data, err := armclient.DoRequest(ctx, "GET", currentItem.ExpandURL)

	if err != nil {
		return ExpanderResult{
			Err:               err,
			Nodes:             newItems,
			SourceDescription: "AppInsightsExpander Request",
			IsPrimaryResponse: true,
		}
	}

	var items []AnalyticsItem
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling analytics items response: %s, %s", err, data)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "AppInsightsExpander request",
		}
	}

	for _, item := range items {
		newItem := TreeNode{
			Name:    item.Name,
			Display: style.Subtle("[" + item.Type + "]") + "\n " + item.Name,
		}
		newItems = append(newItems, &newItem)
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             newItems,
		Response:          data,
		SourceDescription: "AppInsightsExpander request",
	}
}