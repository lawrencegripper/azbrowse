package handlers

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// ActivityLogExpander expands activity logs under an RG
type ActivityLogExpander struct{}

// Name returns the name of the expander
func (e *ActivityLogExpander) Name() string {
	return "ActivityLogExpander"
}

// DoesExpand checks if this is an RG
func (e *ActivityLogExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == activityLogType {
		return true, nil
	}
	return false, nil
}

// Expand returns Resources in the RG
func (e *ActivityLogExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "",
			SourceDescription: "Activity Log",
			IsPrimaryResponse: true,
		}
	}
	newItems := []*TreeNode{}

	var activityLogs armclient.ActivityLogResource
	err = json.Unmarshal([]byte(data), &activityLogs)
	if err != nil {
		panic(err)
	}

	value, err := fastJSONParser.Parse(data)
	if err != nil {
		panic(err)
	}

	for i, log := range activityLogs.Value {
		// Update the existing state as we have more up-to-date info
		objectJSON := string(value.GetArray("value")[i].MarshalTo([]byte("")))

		newItems = append(newItems, &TreeNode{
			Name:            log.OperationName.Value,
			Display:         log.OperationName.LocalizedValue + "\n   " + style.Subtle("At:  "+log.EventTimestamp.String()) + "\n   " + style.Subtle("ResourceType: "+log.ResourceType.Value) + "\n   " + style.Subtle("Status: "+log.Status.Value+""),
			ID:              log.ID,
			Parentid:        currentItem.ID,
			ExpandURL:       ExpandURLNotSupported,
			ItemType:        subActivityLogType,
			SubscriptionID:  currentItem.SubscriptionID,
			StatusIndicator: DrawStatus(log.Status.Value),
			Metadata: map[string]string{
				"jsonItem": objectJSON,
			},
		})
	}

	return ExpanderResult{
		Err:               err,
		Response:          string(data),
		SourceDescription: "Deployments request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

// GetActivityLogExpandURL gets the urls which should be used to get activity logs
func GetActivityLogExpandURL(subscriptionID, resourceName string) string {
	queryString := `eventTimestamp ge '` + time.Now().AddDate(0, 0, -30).Format("2006-01-02T15:04:05Z07:00") + `' and eventTimestamp le '` +
		time.Now().Format("2006-01-02T15:04:05Z07:00") + `' and eventChannels eq 'Admin, Operation' and resourceGroupName eq '` +
		resourceName + `' and levels eq 'Critical,Error,Warning,Informational'`
	return `/subscriptions/` + subscriptionID + `/providers/microsoft.insights/eventtypes/management/values?api-version=2017-03-01-preview&$filter=` +
		url.QueryEscape(queryString)
}
