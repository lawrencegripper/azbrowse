package expanders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

const containerAppRevisionReplicaTemplate = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/revisions/{revisionName}/replicas/{replicaName}"

// ContainerAppNodeType is used to indicate the type of ContainerApp node
type ContainerAppNodeType string

const (
	ContainerAppNode_RevisionReplicaContainers = "containerApp.revision.replica.containers"
	ContainerAppNode_RevisionReplicaContainer  = "containerApp.revision.replica.container"
)

var _ Expander = &ContainerAppExpander{}

type ContainerAppExpanderInterface interface {
	GetAuthToken(ctx context.Context, currentItem *TreeNode) (string, error)
}

var _ ContainerAppExpanderInterface = &ContainerAppExpander{}

type ContainerAppExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *ContainerAppExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *ContainerAppExpander) Name() string {
	return "ContainerAppExpander"
}

// DoesExpand checks if we can expand this node
func (e *ContainerAppExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "subResource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == containerAppRevisionReplicaTemplate {
			return true, nil
		}
	}
	if currentItem.ExpandReturnType == ContainerAppNode_RevisionReplicaContainers {
		return true, nil
	}
	if currentItem.ExpandReturnType == ContainerAppNode_RevisionReplicaContainer {
		return true, nil
	}
	return false, nil
}

// Expand returns items in the container app
func (e *ContainerAppExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	swaggerResourceType := currentItem.SwaggerResourceType
	if swaggerResourceType != nil && swaggerResourceType.Endpoint.TemplateURL == containerAppRevisionReplicaTemplate {
		return e.expandReplica(ctx, currentItem)
	}
	if currentItem.ExpandReturnType == ContainerAppNode_RevisionReplicaContainers {
		return e.expandReplicaContainers(ctx, currentItem)
	}
	if currentItem.ExpandReturnType == ContainerAppNode_RevisionReplicaContainer {
		return e.expandReplicaContainer(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "ContainerAppExpander request",
	}
}

func (e *ContainerAppExpander) expandReplica(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	matchResult := swaggerResourceType.Endpoint.Match(currentItem.ID)
	if !matchResult.IsMatch {
		return ExpanderResult{
			Err:               fmt.Errorf("Error - failed to match own resource type URL"),
			Response:          ExpanderResponse{Response: "Error!"},
			SourceDescription: "ContainerAppExpander request",
		}
	}

	newItems := []*TreeNode{}
	newItems = append(newItems, &TreeNode{
		Name:                  "Containers",
		Display:               "containers",
		ID:                    currentItem.ID + "/containers",
		Parentid:              currentItem.ID,
		ExpandReturnType:      ContainerAppNode_RevisionReplicaContainers,
		SubscriptionID:        currentItem.SubscriptionID,
		SuppressSwaggerExpand: true,
		SuppressGenericExpand: true,
		Metadata: map[string]string{
			"ReplicaExpandURL":  currentItem.ExpandURL,
			"subscriptionId":    matchResult.Values["subscriptionId"],
			"resourceGroupName": matchResult.Values["resourceGroupName"],
			"containerAppName":  matchResult.Values["containerAppName"],
			"revisionName":      matchResult.Values["revisionName"],
			"replicaName":       matchResult.Values["replicaName"],
		},
	})

	return ExpanderResult{
		IsPrimaryResponse: false,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: "TODO - add content here"},
		SourceDescription: "ContainerAppExpander request",
	}
}

func (e *ContainerAppExpander) expandReplicaContainers(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	replicaExpandURL := currentItem.Metadata["ReplicaExpandURL"]
	data, err := e.client.DoRequest(ctx, "GET", replicaExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: ""},
			SourceDescription: "ContainerAppExpander list",
			IsPrimaryResponse: false,
		}
	}

	var replica ContainerAppReplicaResponse
	err = json.Unmarshal([]byte(data), &replica)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error parsing replica JSON: %s", err),
			IsPrimaryResponse: false,
			SourceDescription: "ContainerAppExpander request",
		}
	}

	containersDisplay, err := json.Marshal(replica.Properties.Containers)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error serializing containers JSON: %s", err),
			IsPrimaryResponse: false,
			SourceDescription: "ContainerAppExpander request",
		}
	}
	_ = containersDisplay

	newItems := []*TreeNode{}
	for _, container := range replica.Properties.Containers {
		newItems = append(newItems, &TreeNode{
			Name:                  "Replica container: " + container.Name,
			Display:               container.Name,
			ID:                    currentItem.ID + "/" + container.Name,
			Parentid:              currentItem.ID,
			ExpandReturnType:      ContainerAppNode_RevisionReplicaContainer,
			SubscriptionID:        currentItem.SubscriptionID,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"ContainerAppNodeType": "containerApp.revision.replica.container",
				"ReplicaExpandURL":     currentItem.ExpandURL,
				"LogStreamEndpoint":    container.LogStreamEndpoint,
				"subscriptionId":       currentItem.Metadata["subscriptionId"],
				"resourceGroupName":    currentItem.Metadata["resourceGroupName"],
				"containerAppName":     currentItem.Metadata["containerAppName"],
				"revisionName":         currentItem.Metadata["revisionName"],
				"replicaName":          currentItem.Metadata["replicaName"],
			},
		})
	}
	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: string(containersDisplay), ResponseType: interfaces.ResponseJSON},
		SourceDescription: "ContainerAppExpander request",
	}
}
func (e *ContainerAppExpander) expandReplicaContainer(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             []*TreeNode{},
		Response:          ExpanderResponse{Response: "Use the command panel to show container logs", ResponseType: interfaces.ResponsePlainText},
		SourceDescription: "ContainerAppExpander request",
	}
}

// GetAuthToken retrieves the authentication token for the container app specifiec in currentItem
func (e *ContainerAppExpander) GetAuthToken(ctx context.Context, currentItem *TreeNode) (string, error) {
	var containerAppNode *TreeNode

	// Walk up the tree looking for the container app node
	for node := currentItem; node != nil; node = node.Parent {
		if node.SwaggerResourceType != nil && node.SwaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}" {
			containerAppNode = node
			break
		}
	}
	if containerAppNode == nil {
		return "", fmt.Errorf("failed to find container app node")
	}

	containerAppEndpoint := containerAppNode.SwaggerResourceType.Endpoint
	authEndpoint, err := endpoints.GetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/getAuthToken", containerAppEndpoint.APIVersion)
	if err != nil {
		return "", fmt.Errorf("failed to get auth endpoint: %s", err)
	}

	match := containerAppEndpoint.Match(containerAppNode.ExpandURL)
	if !match.IsMatch {
		return "", fmt.Errorf("failed to match container app URL")
	}

	authURL, err := authEndpoint.BuildURL(match.Values)
	if err != nil {
		return "", fmt.Errorf("failed to build auth URL: %s", err)
	}

	data, err := e.client.DoRequest(ctx, "POST", authURL)
	if err != nil {
		return "", fmt.Errorf("failed to get auth token: %s", err)
	}

	var authTokenResponse ContainerAppAuthTokenResponse
	err = json.Unmarshal([]byte(data), &authTokenResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse auth token response: %s", err)
	}
	return authTokenResponse.Properties.Token, nil
}

// func (e *ContainerAppExpander) getLogs(ctx context.Context, logStreamEndpoint string, authToken string) (string, error) {
// 	// TODO - actually get the logs!

// 	request, err := http.NewRequestWithContext(ctx, "GET", logStreamEndpoint, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create request: %s", err)
// 	}
// 	request.Header.Set("Authorization", "Bearer "+authToken)

// 	httpClient := http.DefaultClient

// 	response, err := httpClient.Do(request)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to make request: %s", err)
// 	}

// 	defer response.Body.Close() //nolint: errcheck

// }

type ContainerAppReplicaContainer struct {
	ContainerID         string `json:"containerId"`
	ExecEndpoint        string `json:"execEndpoint"`
	LogStreamEndpoint   string `json:"logStreamEndpoint"`
	Name                string `json:"name"`
	Ready               bool   `json:"ready"`
	RestartCount        int    `json:"restartCount"`
	RunningState        string `json:"runningState"`
	RunningStateDetails string `json:"runningStateDetails"`
	Started             bool   `json:"started"`
}

type ContainerAppReplicaResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		Containers []ContainerAppReplicaContainer `json:"containers"`
		// CreatedTime         string                         `json:"createdTime"`
		RunningState string `json:"runningState"`
		// RunningStateDetails string                         `json:"runningStateDetails"`
	} `json:"properties"`
}

type ContainerAppAuthTokenResponse struct {
	Properties struct {
		Token   string `json:"token"`
		Expires string `json:"expires"`
	} `json:"properties"`
}
