package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

var _ SwaggerAPISet = SwaggerAPISetARMResources{}

// SwaggerAPISetARMResources holds the config for working with ARM resources as per the published Swagger specs
type SwaggerAPISetARMResources struct {
	resourceTypes []swagger.ResourceType
	client        *armclient.Client
}

// NewSwaggerAPISetARMResources creates a new SwaggerAPISetARMResources
func NewSwaggerAPISetARMResources(client *armclient.Client) SwaggerAPISetARMResources {
	c := SwaggerAPISetARMResources{}
	c.resourceTypes = c.loadResourceTypes()
	c.client = client
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
	data, err := c.client.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL)
		return APISetExpandResponse{Response: data, ResponseType: interfaces.ResponseJSON}, err
	}
	subResources := []SubResource{}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the ARM response and add these to newItems

		var resourceResponse armclient.ResourceResponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
			return APISetExpandResponse{Response: data, ResponseType: interfaces.ResponseJSON}, err
		}

		for _, resource := range resourceResponse.Resources {
			subResourceURL, err := resourceType.PerformSubPathReplace(resource.ID)
			if err != nil {
				err = fmt.Errorf("Error parsing YAML response: %s", err)
				return APISetExpandResponse{Response: data}, err
			}

			subResourceType := resourceType.GetSubResourceTypeForURL(ctx, subResourceURL)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", subResourceURL)
				return APISetExpandResponse{Response: data, ResponseType: interfaces.ResponseJSON}, err
			}
			subResourceTemplateValues := subResourceType.Endpoint.Match(subResourceURL).Values
			name := substituteValues(subResourceType.Display, subResourceTemplateValues)

			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return APISetExpandResponse{Response: data, ResponseType: interfaces.ResponseJSON}, err
				}
			}

			subResource := SubResource{
				ID:           subResourceURL,
				Name:         name,
				ResourceType: *subResourceType,
				ExpandURL:    subResourceURL + "?api-version=" + subResourceType.Endpoint.APIVersion,
				DeleteURL:    deleteURL,
			}
			subResources = append(subResources, subResource)
		}
	}

	return APISetExpandResponse{
		Response:     data,
		ResponseType: interfaces.ResponseJSON,
		SubResources: subResources,
	}, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetARMResources) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	_, err := c.client.DoRequest(context.Background(), "DELETE", item.DeleteURL)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s (%s)", err.Error(), item.DeleteURL)
		return false, err
	}
	return true, nil
}

// Update attempts to update the specified item with new content
func (c SwaggerAPISetARMResources) Update(ctx context.Context, item *TreeNode, content string) error {

	matchResult := item.SwaggerResourceType.Endpoint.Match(item.ExpandURL)
	if !matchResult.IsMatch {
		return fmt.Errorf("item.ExpandURL didn't match current Endpoint")
	}
	putURL, err := item.SwaggerResourceType.PutEndpoint.BuildURL(matchResult.Values)
	if err != nil {
		return fmt.Errorf("Failed to build PUT URL '%s': %s", item.SwaggerResourceType.PutEndpoint.TemplateURL, err)
	}

	data, err := c.client.DoRequestWithBody(ctx, "PUT", putURL, content)
	if err != nil {
		return fmt.Errorf("Error making PUT request: %s", err)
	}

	errorMessage, err := getAPIErrorMessage(data)
	if err != nil {
		return fmt.Errorf("Error checking for API Error message: %s: %s", data, err)
	}
	if errorMessage != "" {
		return fmt.Errorf("Error: %s", errorMessage)
	}
	return nil
}

func getAPIErrorMessage(responseString string) (string, error) {
	var response map[string]interface{}

	err := json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		err = fmt.Errorf("Error parsing API response: %s: %s", responseString, err)
		return "", err
	}
	if response["error"] != nil {
		serializedError, err := json.Marshal(response["error"])
		if err != nil {
			err = fmt.Errorf("Error serializing error message: %s: %s", responseString, err)
			return "", err
		}
		message := string(serializedError)
		message = strings.Replace(message, "\r", "", -1)
		message = strings.Replace(message, "\n", "", -1)
		return message, nil
		// could dig into the JSON to pull out the error message property
	}
	return "", nil
}
