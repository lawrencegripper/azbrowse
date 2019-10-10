package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

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
func (c SwaggerConfigContainerService) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerConfigID set
	// this should never happen for containerService nodes
	return false
}
func (c SwaggerConfigContainerService) GetResourceTypes() []swagger.SwaggerResourceType {
	return c.resourceTypes
}

func (c SwaggerConfigContainerService) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.SwaggerResourceType) (ConfigExpandResponse, error) {

	response, err := c.httpClient.Get(c.serverUrl + currentItem.ExpandURL)
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

	// if len(resourceType.SubResources) > 0 {
	// 	// We have defined subResources - Unmarshal the ARM response and add these to newItems

	// 	var resourceResponse armclient.ResourceResponse
	// 	err = json.Unmarshal([]byte(data), &resourceResponse)
	// 	if err != nil {
	// 		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
	// 		return ConfigExpandResponse{Response: data}, err
	// 	}

	// 	for _, resource := range resourceResponse.Resources {
	// 		subResourceType := getResourceTypeForURL(ctx, resource.ID, resourceType.SubResources)
	// 		if subResourceType == nil {
	// 			err = fmt.Errorf("SubResource type not found! %s", resource.ID)
	// 			return ConfigExpandResponse{Response: data}, err
	// 		}
	// 		subResourceTemplateValues := subResourceType.Endpoint.Match(resource.ID).Values
	// 		name := substituteValues(subResourceType.Display, subResourceTemplateValues)

	// 		deleteURL := ""
	// 		if subResourceType.DeleteEndpoint != nil {
	// 			deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
	// 			if err != nil {
	// 				err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
	// 				return ConfigExpandResponse{Response: data}, err
	// 			}
	// 		}

	// 		subResource := SubResource{
	// 			ID:           resource.ID,
	// 			Name:         name,
	// 			ResourceType: *subResourceType,
	// 			ExpandURL:    resource.ID + "?api-version=" + subResourceType.Endpoint.APIVersion,
	// 			DeleteURL:    deleteURL,
	// 		}
	// 		subResources = append(subResources, subResource)
	// 	}
	// }

	return ConfigExpandResponse{
		Response:     data,
		SubResources: subResources,
	}, nil
}
