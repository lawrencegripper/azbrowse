package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

// SwaggerAPISetARMResources holds the config for working with ARM resources as per the published Swagger specs
type SwaggerAPISetARMResources struct {
	resourceTypes []swagger.ResourceType
}

// NewSwaggerAPISetARMResources creates a new SwaggerAPISetARMResources
func NewSwaggerAPISetARMResources() SwaggerAPISetARMResources {
	c := SwaggerAPISetARMResources{}
	c.resourceTypes = c.loadResourceTypes()
	return c
}

// ID returns the ID for the APISet
func (c SwaggerAPISetARMResources) ID() string {
	return "ARM_RESOURCES_FROM_SPECS"
}

// MatchChildNodesByName indicates whether child nodes should be matched by name (or position)
func (c SwaggerAPISetARMResources) MatchChildNodesByName() bool {
	return true
}

// AppliesToNode is called by the Swagger exapnder to test whether the node applies to this APISet
func (c SwaggerAPISetARMResources) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerAPISetID set

	// handle resource/subresource types
	return node.ItemType == ResourceType || node.ItemType == SubResourceType
}

// GetResourceTypes returns the ResourceTypes for the API Set
func (c SwaggerAPISetARMResources) GetResourceTypes() []swagger.ResourceType {
	return c.resourceTypes
}

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetARMResources) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	method := resourceType.Verb
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL)
		return APISetExpandResponse{Response: data, ResponseType: ResponseJSON}, err
	}
	subResources := []SubResource{}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the ARM response and add these to newItems

		var resourceResponse armclient.ResourceResponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
			return APISetExpandResponse{Response: data, ResponseType: ResponseJSON}, err
		}

		for _, resource := range resourceResponse.Resources {
			subResourceType := getResourceTypeForURL(ctx, resource.ID, resourceType.SubResources)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", resource.ID)
				return APISetExpandResponse{Response: data, ResponseType: ResponseJSON}, err
			}
			subResourceTemplateValues := subResourceType.Endpoint.Match(resource.ID).Values
			name := substituteValues(subResourceType.Display, subResourceTemplateValues)

			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return APISetExpandResponse{Response: data, ResponseType: ResponseJSON}, err
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

	return APISetExpandResponse{
		Response:     data,
		ResponseType: ResponseJSON,
		SubResources: subResources,
	}, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetARMResources) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	_, err := armclient.DoRequest(context.Background(), "DELETE", item.DeleteURL)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s (%s)", err.Error(), item.DeleteURL)
		return false, err
	}
	return true, nil
}
