package expanders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

const (
	// TentantItemType a TreeNode item representing a tenant
	TentantItemType = "tentantItemType"
)

// Check interface
var _ Expander = &TenantExpander{}

// TenantExpander expands the subscriptions under a tenant
type TenantExpander struct {
	client *armclient.Client
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
				ResponseType: ResponsePlainText,
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
				ResponseType: ResponsePlainText,
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
	var newContentType ExpanderResponseType
	if err != nil {
		newContent = err.Error()
		newContentType = ResponsePlainText
	} else {
		newContent = data
		newContentType = ResponseJSON
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
