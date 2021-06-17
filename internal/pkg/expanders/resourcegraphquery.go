package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Check interface
var _ Expander = &ResourceGraphQueryExpander{}

const subNameMapCacheKey = "SubNameMap"

// TenantExpander expands the subscriptions under a tenant
type ResourceGraphQueryExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *ResourceGraphQueryExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *ResourceGraphQueryExpander) Name() string {
	return "ResourceGraphQueryExpander"
}

// DoesExpand checks if this is an RG
func (e *ResourceGraphQueryExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == ResourceGraphQueryType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *ResourceGraphQueryExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	span, ctx := tracing.StartSpanFromContext(ctx, "expand:subs")
	defer span.Finish()

	// Get subs to query
	subs := strings.Split(currentItem.Metadata["subscriptions"], ",")
	// Run the query
	data, err := e.client.DoResourceGraphQueryReturningObjectArray(ctx, subs, currentItem.Metadata["query"])
	if err != nil {
		return ExpanderResult{
			SourceDescription: e.Name(),
			Err:               err,
			Response: ExpanderResponse{
				Response:     data,
				ResponseType: interfaces.ResponsePlainText,
			},
			IsPrimaryResponse: true,
		}
	}

	var queryResponse QueryResponse
	err = json.Unmarshal([]byte(data), &queryResponse)
	if err != nil {
		return ExpanderResult{
			SourceDescription: e.Name(),
			Err:               fmt.Errorf("Failed to load subscriptions: %s", err),
			Response: ExpanderResponse{
				Response:     data,
				ResponseType: interfaces.ResponsePlainText,
			},
			IsPrimaryResponse: true,
		}
	}

	// Get sub ID -> Names mapping
	subNameMapJson, err := storage.GetCache(subNameMapCacheKey)
	if err != nil {
		return ExpanderResult{
			SourceDescription: e.Name(),
			Err:               fmt.Errorf("Failed to load subscription to name mapping from cache: %s", err),
		}
	}
	subNameMap := map[string]string{}
	err = json.Unmarshal([]byte(subNameMapJson), &subNameMap)
	if err != nil {
		return ExpanderResult{
			SourceDescription: e.Name(),
			Err:               fmt.Errorf("Failed to deserialise subscription to name mapping from cache: %s", err),
		}
	}

	newList := []*TreeNode{}
	for _, item := range queryResponse.Data {
		var expandUrl string
		var itemType string
		// Build the expand URL based off the resource type in the query response
		if item.Type == "microsoft.resources/subscriptions/resourcegroups" {
			expandUrl = item.ID + "/resources?api-version=2017-05-10"
			itemType = resourceGroupType
		} else {
			resourceAPIVersion, err := armclient.GetAPIVersion(item.Type)
			if err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					Failure: true,
					Message: "Failed to get resourceVersion for the Type:" + item.Type,
					Timeout: time.Duration(time.Second * 5),
				})
			}
			expandUrl = item.ID + "?api-version=" + resourceAPIVersion
			itemType = ResourceType
		}

		// Add it to the list
		newList = append(newList, &TreeNode{
			Display: style.Subtle("["+item.SubscriptionID+"]") +
				"\n  " +
				style.Subtitle("["+subNameMap[item.SubscriptionID]+"]") +
				"\n  " +
				item.Name,
			Name:           item.Name,
			ID:             item.ID,
			ExpandURL:      expandUrl,
			ItemType:       itemType,
			SubscriptionID: item.SubscriptionID,
		})
	}

	return ExpanderResult{
		SourceDescription: e.Name(),
		IsPrimaryResponse: true,
		Nodes:             newList,
		Response: ExpanderResponse{
			Response:     data,
			ResponseType: interfaces.ResponseJSON,
		},
	}
}

type QueryResponse struct {
	Totalrecords int `json:"totalRecords"`
	Count        int `json:"count"`
	Data         []struct {
		Name           string `json:"name"`
		Type           string `json:"type"`
		Location       string `json:"location"`
		SubscriptionID string `json:"subscriptionId"`
		ID             string `json:"id"`
	} `json:"data"`
	Facets          []interface{} `json:"facets"`
	Resulttruncated string        `json:"resultTruncated"`
}
