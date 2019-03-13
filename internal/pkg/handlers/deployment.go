package handlers

import (
	"context"
	"encoding/json"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// DeploymentExpander expands a deployment to list operations
type DeploymentExpander struct{}

// Name returns the name of the expander
func (e *DeploymentExpander) Name() string {
	return "DeploymentExpander"
}

// DoesExpand checks if this is an RG
func (e *DeploymentExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == deploymentType {
		return true, nil
	}
	return false, nil
}

// Expand returns Resources in the RG
func (e *DeploymentExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "",
			SourceDescription: "Deployment Subdeployment",
			IsPrimaryResponse: true,
		}
	}
	newItems := []*TreeNode{}

	var operations armclient.DeploymentOperationsResponse
	err = json.Unmarshal([]byte(data), &operations)
	if err != nil {
		panic(err)
	}

	value, err := fastJSONParser.Parse(data)
	if err != nil {
		panic(err)
	}

	for i, operation := range operations.Value {
		// Update the existing state as we have more up-to-date info
		objectJSON := string(value.GetArray("value")[i].MarshalTo([]byte("")))

		display := operation.OperationID + "\n   " + style.Subtle("Started:"+operation.Properties.Timestamp) + "\n   " + style.Subtle("Duration: "+operation.Properties.Duration) + "\n   " + style.Subtle("DeploymentStatus: "+operation.Properties.ProvisioningState+"")
		if operation.Properties.TargetResource.ResourceType != "" {
			display += "\n   " + style.Subtle("ResourceType:"+operation.Properties.TargetResource.ResourceType) +
				"\n   " + style.Subtle("ResourceName:"+operation.Properties.TargetResource.ResourceName)
		}
		newItems = append(newItems, &TreeNode{
			Name:           operation.OperationID,
			Display:        display,
			ID:             operation.ID,
			Parentid:       currentItem.ID,
			ExpandURL:      ExpandURLNotSupported,
			ItemType:       deploymentOperationType,
			SubscriptionID: currentItem.SubscriptionID,
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
		IsPrimaryResponse: false,
	}
}
