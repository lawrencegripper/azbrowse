package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
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
	data, err := e.client.DoRequest(ctx, "GET", "/subscriptions?api-version=2018-01-01")
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
		// return armclient.SubResponse{}, "", fmt.Errorf("Failed to load subscriptions: %s", err)
	}

	var subRequest SubResponse
	err = json.Unmarshal([]byte(data), &subRequest)
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

	newList := []*TreeNode{}
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

	var newContent string
	var newContentType interfaces.ExpanderResponseType
	if err != nil {
		newContent = err.Error()
		newContentType = interfaces.ResponsePlainText
	} else {
		newContent = data
		newContentType = interfaces.ResponseJSON
	}

	return ExpanderResult{
		SourceDescription: e.Name(),
		IsPrimaryResponse: true,
		Nodes:             newList,
		Response: ExpanderResponse{
			Response:     newContent,
			ResponseType: newContentType,
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
				st.Expect(t, len(r.Nodes), 3)

				// Validate content
				st.Expect(t, r.Nodes[0].Display, "1testsub")
				st.Expect(t, r.Nodes[0].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01")
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
