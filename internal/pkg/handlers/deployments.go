package handlers

import (
	"context"
	"encoding/json"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// DeploymentsExpander expands RGs under a subscription
type DeploymentsExpander struct{}

// Name returns the name of the expander
func (e *DeploymentsExpander) Name() string {
	return "DeploymentsExpander"
}

// DoesExpand checks if this is an RG
func (e *DeploymentsExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == deploymentsType {
		return true, nil
	}
	return false, nil
}

// Expand returns Resources in the RG
func (e *DeploymentsExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "",
			SourceDescription: "Deployments Subdeployment",
			IsPrimaryResponse: true,
		}
	}
	newItems := []*TreeNode{}

	var deployments armclient.DeploymentsResponse
	err = json.Unmarshal([]byte(data), &deployments)
	if err != nil {
		panic(err)
	}

	value, err := fastJSONParser.Parse(data)
	if err != nil {
		panic(err)
	}
	for i, dep := range deployments.Value {
		// Update the existing state as we have more up-to-date info
		objectJSON := string(value.GetArray("value")[i].MarshalTo([]byte("")))
		newItems = append(newItems, &TreeNode{
			Name:            dep.Name,
			Display:         dep.Name + "\n   " + style.Subtle("Started:  "+dep.Properties.Timestamp) + "\n   " + style.Subtle("Duration: "+dep.Properties.Duration) + "\n   " + style.Subtle("DeploymentStatus: "+dep.Properties.ProvisioningState+""),
			ID:              dep.ID,
			Parentid:        currentItem.ID,
			ExpandURL:       dep.ID + "/operations/?api-version=2017-05-10",
			ItemType:        deploymentType,
			DeleteURL:       dep.ID + "?api-version=2017-05-10",
			SubscriptionID:  currentItem.SubscriptionID,
			StatusIndicator: DrawStatus(dep.Properties.ProvisioningState),
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
