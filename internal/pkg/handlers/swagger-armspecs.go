package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type SwaggerConfigARMResources struct {
	resourceTypes []swagger.SwaggerResourceType
}

func NewSwaggerConfigARMResources() SwaggerConfigARMResources {
	c := SwaggerConfigARMResources{}
	c.resourceTypes = c.loadResourceTypes()
	return c
}

func (c SwaggerConfigARMResources) ID() string {
	return "ARM_RESOURCES_FROM_SPECS"
}
func (c SwaggerConfigARMResources) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerConfigID set

	// handle resource/subresource types
	return node.ItemType == ResourceType || node.ItemType == SubResourceType
}
func (c SwaggerConfigARMResources) GetResourceTypes() []swagger.SwaggerResourceType {
	return c.resourceTypes
}

func (c SwaggerConfigARMResources) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.SwaggerResourceType) (ConfigExpandResponse, error) {

	method := resourceType.Verb
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL)
		return ConfigExpandResponse{Response: data}, err
	}
	subResources := []SubResource{}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the ARM response and add these to newItems

		var resourceResponse armclient.ResourceResponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
			return ConfigExpandResponse{Response: data}, err
		}

		for _, resource := range resourceResponse.Resources {
			subResourceType := getResourceTypeForURL(ctx, resource.ID, resourceType.SubResources)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", resource.ID)
				return ConfigExpandResponse{Response: data}, err
			}
			subResourceTemplateValues := subResourceType.Endpoint.Match(resource.ID).Values
			name := substituteValues(subResourceType.Display, subResourceTemplateValues)

			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return ConfigExpandResponse{Response: data}, err
				}
			}

			subResource := SubResource{
				ID:           resource.ID,
				Name:         name,
				ResourceType: *subResourceType,
				ExpandURL:    resource.ID + "?api-version=" + subResourceType.Endpoint.APIVersion,
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
