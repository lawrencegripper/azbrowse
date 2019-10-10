package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

// ResourceResponse Resources list rest type
type KubernetesListResponse struct {
	Items []KubernetesItem `json:"items"`
}
type KubernetesItem struct {
	Metadata struct {
		Name     string `yaml:"name"`
		SelfLink string `yaml:"selfLink"`
	} `yaml:"metadata`
}

type SwaggerConfigContainerService struct {
	resourceTypes []swagger.SwaggerResourceType
	httpClient    http.Client
	clusterID     string
	serverUrl     string
}

func NewSwaggerConfigContainerService(resourceTypes []swagger.SwaggerResourceType, httpClient http.Client, clusterID string, serverUrl string) SwaggerConfigContainerService {
	c := SwaggerConfigContainerService{}
	c.resourceTypes = resourceTypes
	c.httpClient = httpClient
	c.clusterID = clusterID
	c.serverUrl = serverUrl
	return c
}

func (c SwaggerConfigContainerService) ID() string {
	return c.clusterID
}
func (c SwaggerConfigContainerService) MatchChildNodesByName() bool {
	return false
}
func (c SwaggerConfigContainerService) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerConfigID set
	// this should never happen for containerService nodes
	return false
}
func (c SwaggerConfigContainerService) GetResourceTypes() []swagger.SwaggerResourceType {
	return c.resourceTypes
}

func (c SwaggerConfigContainerService) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.SwaggerResourceType) (ConfigExpandResponse, error) {

	url := c.serverUrl + currentItem.ExpandURL
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte("")))

	request.Header.Set("Accept", "application/yaml")

	response, err := c.httpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL)
		return ConfigExpandResponse{}, err
	}
	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Failed to read body: %s", err)
		return ConfigExpandResponse{}, err
	}
	data := string(buf)

	subResources := []SubResource{}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the response and add these to newItems

		var listResponse KubernetesListResponse
		err = yaml.Unmarshal([]byte(data), &listResponse)
		if err != nil {
			err = fmt.Errorf("Error parsing YAML response: %s", err)
			return ConfigExpandResponse{Response: data}, err
		}

		for _, item := range listResponse.Items {
			subResourceType := getResourceTypeForURL(ctx, item.Metadata.SelfLink, resourceType.SubResources)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", item.Metadata.SelfLink)
				return ConfigExpandResponse{Response: data}, err
			}
			name := item.Metadata.Name
			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				subResourceTemplateValues := subResourceType.Endpoint.Match(item.Metadata.SelfLink).Values
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return ConfigExpandResponse{Response: data}, err
				}
			}
			subResource := SubResource{
				ID:           c.clusterID + item.Metadata.SelfLink,
				Name:         name,
				ResourceType: *subResourceType,
				ExpandURL:    item.Metadata.SelfLink,
				DeleteURL:    deleteURL,
			}
			subResources = append(subResources, subResource)
		}
	}

	return ConfigExpandResponse{
		Response:     data,
		SubResources: subResources,
	}, nil
}
