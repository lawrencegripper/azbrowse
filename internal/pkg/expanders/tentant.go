package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/nbio/st"
)

const (
	// TentantItemType a TreeNode item representing a tenant
	TentantItemType = "tentantItemType"
)

// Check interface
var _ Expander = &TenantExpander{}

// TenantExpander expands the subscriptions under a tenant
type TenantExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *TenantExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *TenantExpander) Name() string {
	return "TenantExpander"
}

// DoesExpand checks if this is an RG
func (e *TenantExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == TentantItemType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *TenantExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	span, ctx := tracing.StartSpanFromContext(ctx, "expand:subs")
	defer span.Finish()

	// Get Subscriptions
	data, err := e.getSubscriptionsWithPaging(ctx)
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

	var subRequest SubResponse
	err = json.Unmarshal([]byte(data), &subRequest)
	if err != nil {
		return ExpanderResult{
			SourceDescription: e.Name(),
			Err:               fmt.Errorf("failed to load subscriptions: %w", err),
			Response: ExpanderResponse{
				Response:     data,
				ResponseType: interfaces.ResponsePlainText,
			},
			IsPrimaryResponse: true,
		}
	}

	newList := []*TreeNode{}
	newList = append(newList, &TreeNode{
		Display:        "MS Graph",
		Namespace:      "graph",
		Name:           "MS Graph",
		ID:             "graph",
		ExpandURL:      ExpandURLNotSupported,
		ItemType:       GraphType,
		SubscriptionID: "",
	})

	subIds := make([]string, 0, len(subRequest.Subs))
	subNameMap := map[string]string{}
	for _, sub := range subRequest.Subs {
		subIds = append(subIds, strings.Replace(sub.ID, "/subscriptions/", "", 1))
		subNameMap[sub.SubscriptionID] = sub.DisplayName
	}
	subNameMapJson, err := json.Marshal(subNameMap)
	if err != nil {
		panic("Failed to marshal map to json for subnames")
	}
	err = storage.PutCache(subNameMapCacheKey, string(subNameMapJson)) //nolint: errcheck
	if err != nil {
		panic("Failed to save json subnames to cache")
	}
	// Load any custom graph queryies
	queries, err := config.GetCustomResourceGraphQueries()
	if err != nil {
		eventing.SendFailureStatus(fmt.Sprintf("Failed to load custom resource graph queries %q", err))
	}
	for _, query := range queries {
		newList = append(newList, &TreeNode{
			Display:        style.Subtle("[Microsoft.ResourceGraph]") + "\n  Query: " + query.Name,
			Name:           query.Name,
			ID:             query.Name,
			ExpandURL:      ExpandURLNotSupported,
			ItemType:       ResourceGraphQueryType,
			SubscriptionID: NotSupported,
			Metadata: map[string]string{
				"subscriptions": strings.Join(subIds, ","),
				"query":         query.Query,
			},
		})
	}

	// Add each subscription in tenant
	for _, sub := range subRequest.Subs {
		newList = append(newList, &TreeNode{
			Display:        sub.DisplayName,
			Name:           sub.DisplayName,
			ID:             sub.ID,
			ExpandURL:      sub.ID + "/resourceGroups?api-version=2018-05-01",
			ItemType:       SubscriptionType,
			SubscriptionID: sub.SubscriptionID,
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

// SubResponse Subscriptions REST type
type SubResponse struct {
	Subs []struct {
		ID                   string `json:"id"`
		SubscriptionID       string `json:"subscriptionId"`
		DisplayName          string `json:"displayName"`
		State                string `json:"state"`
		SubscriptionPolicies struct {
			LocationPlacementID string `json:"locationPlacementId"`
			QuotaID             string `json:"quotaId"`
			SpendingLimit       string `json:"spendingLimit"`
		} `json:"subscriptionPolicies"`
	} `json:"value"`
	NextLink string `json:"nextLink,omitempty"`
}

// getSubscriptionsWithPaging retrieves all subscriptions using pagination
func (e *TenantExpander) getSubscriptionsWithPaging(ctx context.Context) (string, error) {
	// Start with the initial request
	url := "/subscriptions?api-version=2018-01-01"
	data, err := e.client.DoRequest(ctx, "GET", url)
	if err != nil {
		return data, err
	}

	// Parse initial response to check for nextLink
	var initialResponse SubResponse
	err = json.Unmarshal([]byte(data), &initialResponse)
	if err != nil {
		return data, fmt.Errorf("failed to parse subscription response: %w", err)
	}

	// If there's no nextLink, no need for pagination
	if initialResponse.NextLink == "" {
		return data, nil
	}

	// We need to aggregate all subscriptions from all pages
	var allSubs SubResponse
	allSubs.Subs = initialResponse.Subs

	// Follow nextLinks until there are no more pages
	nextLink := initialResponse.NextLink
	for nextLink != "" {
		// The nextLink is a full URL, we need to extract just the path+query part for DoRequest
		// Remove the ARM endpoint part (https://management.azure.com) from the URL
		path := strings.TrimPrefix(nextLink, "https://management.azure.com")

		// Make the request to the next page
		pageData, err := e.client.DoRequest(ctx, "GET", path)
		if err != nil {
			// Return what we have so far along with the error
			return "", fmt.Errorf("error retrieving additional subscription pages: %w", err)
		}

		// Parse this page's response
		var pageResponse SubResponse
		err = json.Unmarshal([]byte(pageData), &pageResponse)
		if err != nil {
			// Return what we have so far along with the error
			return "", fmt.Errorf("failed to parse additional subscription page: %w", err)
		}

		// Add this page's subscriptions to our aggregated list
		allSubs.Subs = append(allSubs.Subs, pageResponse.Subs...)

		// Update nextLink for the next iteration
		nextLink = pageResponse.NextLink
	}

	// Convert the aggregated response back to JSON
	combinedData, err := json.Marshal(allSubs)
	if err != nil {
		return "", fmt.Errorf("failed to marshal aggregated subscription list: %w", err)
	}

	return string(combinedData), nil
}

func (e *TenantExpander) testCases() (bool, *[]expanderTestCase) {
	treeNode := &TreeNode{
		ItemType:  TentantItemType,
		ID:        "AvailableSubscriptions",
		ExpandURL: ExpandURLNotSupported,
	}
	return true, &[]expanderTestCase{
		{
			name:         "Tenant->Subs",
			nodeToExpand: treeNode,
			urlPath:      "subscriptions",
			responseFile: "./testdata/armsamples/subscriptions/response.json",
			statusCode:   200,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)
				st.Expect(t, len(r.Nodes), 4)

				// Validate content
				st.Expect(t, r.Nodes[1].Display, "1testsub")
				st.Expect(t, r.Nodes[1].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01")
			},
		},
		{
			name: "Tenant->Subs500Response",
			nodeToExpand: &TreeNode{
				ItemType:  TentantItemType,
				ID:        "AvailableSubscriptions",
				ExpandURL: ExpandURLNotSupported,
			},
			urlPath:    "subscriptions",
			statusCode: 500,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				if r.Err == nil {
					t.Error("Failed expanding resource. Should have errored and didn't", r)
				}
			},
		},
	}
}
