package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

const containerInstanceTemplate = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
const containerInstanceNamespace = "containerInstance"

// ContainerInstanceExpander expands the data-plane aspects of a Container Instance
type ContainerInstanceExpander struct {
}

// Name returns the name of the expander
func (e *ContainerInstanceExpander) Name() string {
	return "ContainerInstanceExpander"
}

// DoesExpand checks if this is a storage account
func (e *ContainerInstanceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == containerInstanceTemplate {
			return true, nil
		}
	}
	if currentItem.ExpandReturnType == "containerInstance.logs" {
		return true, nil
	}
	return false, nil
}

// Expand adds items for container instance items to the list
func (e *ContainerInstanceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	if currentItem.ExpandReturnType == "containerInstance.logs" {
		return e.expandLogs(ctx, currentItem)
	}

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != containerInstanceNamespace &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == containerInstanceTemplate {

		return e.expandContainers(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          "Error!",
		SourceDescription: "ContainerInstanceExpander request",
	}
}

func (e *ContainerInstanceExpander) expandContainers(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	resourceAPIVersion, err := armclient.GetAPIVersion(currentItem.ArmType)
	if err != nil {
		eventing.SendStatusEvent(eventing.StatusEvent{
			Failure: true,
			Message: "Failed to get resouceVersion for the Type:" + currentItem.ArmType,
			Timeout: time.Duration(time.Second * 5),
		})
	}
	containersListURL := currentItem.ID + "?api-version=" + resourceAPIVersion

	data, err := armclient.DoRequest(ctx, "GET", containersListURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "",
			SourceDescription: "expandContainers list",
			IsPrimaryResponse: false,
		}
	}

	newItems := []*TreeNode{}

	var containerGroupResponse armclient.ContainerGroupResponse
	err = json.Unmarshal([]byte(data), &containerGroupResponse)
	if err != nil {
		panic(err)
	}

	for _, container := range containerGroupResponse.Properties.Containers {
		newItems = append(newItems, &TreeNode{
			Name:             container.Name,
			Display:          container.Name + "\n   " + style.Subtle("Status:  "+container.Properties.InstanceView.CurrentState.State) + "\n   " + style.Subtle("Restarts:  "+strconv.Itoa(container.Properties.InstanceView.RestartCount)) + "\n   " + style.Subtle("Image: "+container.Properties.Image),
			ID:               currentItem.ID + "/" + container.Name,
			Parentid:         currentItem.ID,
			ExpandURL:        currentItem.ID + "/containers/" + container.Name + "/logs?tail=400&api-version=" + resourceAPIVersion,
			ExpandReturnType: "containerInstance.logs",
			SubscriptionID:   currentItem.SubscriptionID,
			StatusIndicator:  DrawStatus(container.Properties.InstanceView.CurrentState.State),
			Metadata: map[string]string{
				"ContainerName":         container.Name,
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "false",
			},
		})
	}

	return ExpanderResult{
		IsPrimaryResponse: false,
		Nodes:             newItems,
		SourceDescription: "ContainerInstanceExpander request",
	}
}

func (e *ContainerInstanceExpander) expandLogs(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	containersLogURL := currentItem.ExpandURL

	data, err := armclient.DoRequest(ctx, "GET", containersLogURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "",
			SourceDescription: "expandContainers logs",
			IsPrimaryResponse: true,
		}
	}

	var containerLogResponse armclient.ContainerLogResponse
	err = json.Unmarshal([]byte(data), &containerLogResponse)
	if err != nil {
		panic(err)
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Response:          containerLogResponse.Content,
		SourceDescription: "ContainerInstanceExpander request",
	}
}
