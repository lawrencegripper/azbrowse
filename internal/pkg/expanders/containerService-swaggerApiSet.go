package expanders

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type kubernetesListResponse struct {
	Items []kubernetesItem `json:"items"`
}
type kubernetesItem struct {
	Metadata struct {
		Name     string `yaml:"name"`
		SelfLink string `yaml:"selfLink"`
	} `yaml:"metadata"`
}
type podResponse struct {
	Spec struct {
		Containers []struct {
			Name string `yaml:"name"`
		} `yaml:"containers"`
	} `yaml:"spec"`
}

var _ SwaggerAPISet = SwaggerAPISetContainerService{}
var maxTailLines = 100

// SwaggerAPISetContainerService holds the config for working with an AKS cluster API
type SwaggerAPISetContainerService struct {
	resourceTypes []swagger.ResourceType
	httpClient    http.Client
	clusterID     string
	serverURL     string
}

// NewSwaggerAPISetContainerService creates a new SwaggerAPISetContainerService
func NewSwaggerAPISetContainerService(resourceTypes []swagger.ResourceType, httpClient http.Client, clusterID string, serverURL string) SwaggerAPISetContainerService {
	c := SwaggerAPISetContainerService{}
	c.resourceTypes = resourceTypes
	c.httpClient = httpClient
	c.clusterID = clusterID
	c.serverURL = serverURL
	return c
}

// ID returns the ID for the APISet
func (c SwaggerAPISetContainerService) ID() string {
	return c.clusterID
}

// MatchChildNodesByName indicates whether child nodes should be matched by name (or position)
func (c SwaggerAPISetContainerService) MatchChildNodesByName() bool {
	return false
}

// AppliesToNode is called by the Swagger exapnder to test whether the node applies to this APISet
func (c SwaggerAPISetContainerService) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerAPISetID set
	// this should never happen for containerService nodes
	return false
}

// GetResourceTypes returns the ResourceTypes for the API Set
func (c SwaggerAPISetContainerService) GetResourceTypes() []swagger.ResourceType {
	return c.resourceTypes
}

func (c SwaggerAPISetContainerService) doRequest(ctx context.Context, verb string, url string) (string, error) {
	return c.doRequestWithBody(ctx, verb, url, "")
}

func (c SwaggerAPISetContainerService) doRequestWithBody(ctx context.Context, verb string, url string, body string) (string, error) {
	request, err := http.NewRequest(verb, url, bytes.NewReader([]byte(body)))
	if err != nil {
		err = fmt.Errorf("Failed to create request" + err.Error() + url)
		return "", err
	}

	request.Header.Set("Content-Type", "application/yaml")
	request.Header.Set("Accept", "application/yaml")
	response, err := c.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + url)
		return "", err
	}
	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Failed to read body: %s", err)
		return "", err
	}
	data := string(buf)
	if 200 <= response.StatusCode && response.StatusCode < 300 {
		return data, nil
	}
	return "", fmt.Errorf("Response failed with %s (%s): %s", response.Status, url, data)
}

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetContainerService) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	if resourceType.Endpoint.TemplateURL == "/api/v1/namespaces/{namespace}/pods/{name}/log" {
		if !strings.Contains(currentItem.ExpandURL, "?") { // we haven't already set the container name/tailLines!

			logURL := c.serverURL + currentItem.ExpandURL
			containerURL := logURL[:len(logURL)-3]
			data, err := c.doRequest(ctx, "GET", containerURL)
			if err != nil {
				err = fmt.Errorf("Failed to make request: %s", err)
				return APISetExpandResponse{}, err
			}

			var podInfo podResponse
			err = yaml.Unmarshal([]byte(data), &podInfo)
			if err != nil {
				err = fmt.Errorf("Error parsing YAML response: %s", err)
				return APISetExpandResponse{Response: data}, err
			}
			if podInfo.Spec.Containers == nil || len(podInfo.Spec.Containers) == 0 {
				err = fmt.Errorf("No containers in response: %s", err)
				return APISetExpandResponse{}, err
			}

			if len(podInfo.Spec.Containers) == 1 {
				// if only a single resopnse then set the tailLines param and  fall through to just return logs for the single container
				currentItem.ExpandURL += "?tailLines=" + strconv.Itoa(maxTailLines)
			} else {
				if len(podInfo.Spec.Containers) > 1 {
					subResources := []SubResource{}
					for _, container := range podInfo.Spec.Containers {
						subResource := SubResource{
							ID:           currentItem.ID + "/" + container.Name,
							Name:         container.Name,
							ResourceType: resourceType,
							ExpandURL:    currentItem.ExpandURL + "?container=" + container.Name + "&tailLines=" + strconv.Itoa(maxTailLines),
						}
						subResources = append(subResources, subResource)
					}
					return APISetExpandResponse{
						Response:     "Pick a container to view logs",
						SubResources: subResources,
					}, nil
				}
			}
		}
	}

	subResources := []SubResource{}
	url := c.serverURL + currentItem.ExpandURL
	data, err := c.doRequest(ctx, "GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make request: %s", err)
		return APISetExpandResponse{}, err
	}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the response and add these to newItems

		var listResponse kubernetesListResponse
		err = yaml.Unmarshal([]byte(data), &listResponse)
		if err != nil {
			err = fmt.Errorf("Error parsing YAML response: %s", err)
			return APISetExpandResponse{Response: data}, err
		}

		for _, item := range listResponse.Items {
			subResourceURL, err := resourceType.PerformSubPathReplace(item.Metadata.SelfLink)
			if err != nil {
				err = fmt.Errorf("Error parsing YAML response: %s", err)
				return APISetExpandResponse{Response: data}, err
			}

			subResourceType := resourceType.GetSubResourceTypeForURL(ctx, subResourceURL)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", subResourceURL)
				return APISetExpandResponse{Response: data}, err
			}
			name := item.Metadata.Name
			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				subResourceTemplateValues := subResourceType.Endpoint.Match(subResourceURL).Values
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return APISetExpandResponse{Response: data}, err
				}
			}
			subResource := SubResource{
				ID:           c.clusterID + subResourceURL,
				Name:         name,
				ResourceType: *subResourceType,
				ExpandURL:    subResourceURL,
				DeleteURL:    deleteURL,
			}
			subResources = append(subResources, subResource)
		}
	}

	return APISetExpandResponse{
		Response:     data,
		ResponseType: interfaces.ResponseYAML,
		SubResources: subResources,
	}, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetContainerService) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	url := c.serverURL + item.DeleteURL
	_, err := c.doRequest(ctx, "DELETE", url)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s (%s)", err.Error(), item.DeleteURL)
		return false, err
	}
	return true, nil
}

// Update attempts to update the specified item with new content
func (c SwaggerAPISetContainerService) Update(ctx context.Context, item *TreeNode, content string) error {
	matchResult := item.SwaggerResourceType.Endpoint.Match(item.ExpandURL)
	if !matchResult.IsMatch {
		return fmt.Errorf("item.ExpandURL didn't match current Endpoint")
	}
	putURL, err := item.SwaggerResourceType.PutEndpoint.BuildURL(matchResult.Values)
	putURL = c.serverURL + putURL
	if err != nil {
		return fmt.Errorf("Failed to build PUT URL '%s': %s", item.SwaggerResourceType.PutEndpoint.TemplateURL, err)
	}
	_, err = c.doRequestWithBody(ctx, "PUT", putURL, content)
	if err != nil {
		return fmt.Errorf("Error making PUT request: %s", err)
	}

	return nil
}
