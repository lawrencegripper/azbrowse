package expanders

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

// Check interface
var _ Expander = &ContainerInstanceExpander{}

// ContainerInstanceExpander expands the data-plane aspects of a Container Instance
type ContainerInstanceExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *ContainerInstanceExpander) setClient(c *armclient.Client) {
	e.client = c
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
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "ContainerInstanceExpander request",
	}
}

func (e *ContainerInstanceExpander) expandContainers(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	resourceAPIVersion, err := armclient.GetAPIVersion(currentItem.ArmType)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			Failure: true,
			Message: "Failed to get resouceVersion for the Type:" + currentItem.ArmType,
			Timeout: time.Duration(time.Second * 5),
		})
	}
	containersListURL := currentItem.ID + "?api-version=" + resourceAPIVersion

	data, err := e.client.DoRequest(ctx, "GET", containersListURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: ""},
			SourceDescription: "expandContainers list",
			IsPrimaryResponse: false,
		}
	}

	newItems := []*TreeNode{}

	var containerGroupResponse ContainerGroupResponse
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

	data, err := e.client.DoRequest(ctx, "GET", containersLogURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: ""},
			SourceDescription: "expandContainers logs",
			IsPrimaryResponse: true,
		}
	}

	var containerLogResponse ContainerLogResponse
	err = json.Unmarshal([]byte(data), &containerLogResponse)
	if err != nil {
		panic(err)
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Response:          ExpanderResponse{Response: containerLogResponse.Content},
		SourceDescription: "ContainerInstanceExpander request",
	}
}

// ContainerLogResponse for container logs
type ContainerLogResponse struct {
	Content string `json:"content"`
}

// ContainerGroupResponse is the response to a get request on a container group
type ContainerGroupResponse struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Containers []struct {
			Name       string `json:"name"`
			Properties struct {
				Command              []interface{} `json:"command"`
				EnvironmentVariables []interface{} `json:"environmentVariables"`
				Image                string        `json:"image"`
				Ports                []struct {
					Port int `json:"port"`
				} `json:"ports"`
				InstanceView struct {
					RestartCount int `json:"restartCount"`
					CurrentState struct {
						State        string    `json:"state"`
						StartTime    time.Time `json:"startTime"`
						DetailStatus string    `json:"detailStatus"`
					} `json:"currentState"`
					Events []struct {
						Count          int       `json:"count"`
						FirstTimestamp time.Time `json:"firstTimestamp"`
						LastTimestamp  time.Time `json:"lastTimestamp"`
						Name           string    `json:"name"`
						Message        string    `json:"message"`
						Type           string    `json:"type"`
					} `json:"events"`
				} `json:"instanceView"`
				Resources struct {
					Requests struct {
						CPU        float64 `json:"cpu"`
						MemoryInGB float64 `json:"memoryInGB"`
					} `json:"requests"`
				} `json:"resources"`
				VolumeMounts []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly"`
				} `json:"volumeMounts"`
			} `json:"properties"`
		} `json:"containers"`
		ImageRegistryCredentials []struct {
			Server   string `json:"server"`
			Username string `json:"username"`
		} `json:"imageRegistryCredentials"`
		IPAddress struct {
			IP    string `json:"ip"`
			Ports []struct {
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
			Type string `json:"type"`
		} `json:"ipAddress"`
		OsType            string `json:"osType"`
		ProvisioningState string `json:"provisioningState"`
		Volumes           []struct {
			AzureFile struct {
				ReadOnly           bool   `json:"readOnly"`
				ShareName          string `json:"shareName"`
				StorageAccountName string `json:"storageAccountName"`
			} `json:"azureFile"`
			Name string `json:"name"`
		} `json:"volumes"`
	} `json:"properties"`
	Type string `json:"type"`
}

func (e *ContainerInstanceExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
