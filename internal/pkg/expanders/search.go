package expanders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

const azureSearchTemplateURL string = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}"

type searchServiceResponse struct {
	Name string `json:"name"`
}

type adminKeysResponse struct {
	PrimaryKey string `json:"primaryKey"`
}

// AzureSearchServiceExpander expands the kubernetes aspects of AKS
type AzureSearchServiceExpander struct {
	client *armclient.Client
}

func (e *AzureSearchServiceExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *AzureSearchServiceExpander) Name() string {
	return "AzureSearchServiceExpander"
}

// DoesExpand checks if this is a storage account
func (e *AzureSearchServiceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == azureSearchTemplateURL {
			return true, nil
		}
	}
	if currentItem.Namespace == "AzureSearchServiceExpander" {
		return true, nil
	}
	return false, nil
}

// Expand returns ManagementPolicies in the StorageAccount
func (e *AzureSearchServiceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "AzureSearchServiceExpander" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == azureSearchTemplateURL {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			ID:        currentItem.ID + "/<service>",
			Parentid:  currentItem.ID,
			Namespace: "AzureSearchServiceExpander",
			Name:      "Search Service",
			Display:   "Search Service",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"SearchID":              currentItem.ID,
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "AzureSearchServiceExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "AzureSearchServiceExpander" && currentItem.ItemType == SubResourceType {
		return e.expandSearchRoot(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "AzureSearchServiceExpander request",
	}
}

func (e *AzureSearchServiceExpander) expandSearchRoot(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	clusterID := currentItem.Metadata["SearchID"]

	// Check for existing config for the cluster
	apiSet := e.getAPISetForCluster(clusterID)
	var err error
	if apiSet == nil {
		apiSet, err = e.createAPISetForCluster(ctx, clusterID)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				Response:          ExpanderResponse{Response: "Error!"},
				SourceDescription: "AzureSearchServiceExpander request",
			}
		}
		GetSwaggerResourceExpander().AddAPISet(*apiSet)
	}

	swaggerResourceTypes := apiSet.GetResourceTypes()

	newItems := []*TreeNode{}
	for _, child := range swaggerResourceTypes {
		resourceType := child
		display := resourceType.Display
		if display == "{}" {
			display = resourceType.Endpoint.TemplateURL
		}
		newItems = append(newItems, &TreeNode{
			Parentid:            currentItem.ID,
			ID:                  currentItem.ID + "/" + display,
			Namespace:           "swagger",
			Name:                display,
			Display:             display,
			ExpandURL:           resourceType.Endpoint.TemplateURL + "?api-version=" + resourceType.Endpoint.APIVersion, // all fixed template URLs
			ItemType:            SubResourceType,
			SwaggerResourceType: &resourceType,
			Metadata: map[string]string{
				"SwaggerAPISetID": currentItem.ID,
			},
		})
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: "TODO - what should go here?"},
		SourceDescription: "AzureSearchServiceExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureSearchServiceExpander) createAPISetForCluster(ctx context.Context, searchID string) (*SwaggerAPISetSearch, error) {

	adminKey, err := e.getAdminKey(ctx, searchID)
	if err != nil {
		return nil, err
	}

	searchEndpoint, err := e.getSearchEndpoint(ctx, searchID)
	if err != nil {
		return nil, err
	}

	swaggerResourceTypes := e.loadResourceTypes()
	if err != nil {
		return nil, err
	}

	// Register the swagger config so that the swagger expander can take over
	apiSet := NewSwaggerAPISetSearch(swaggerResourceTypes, searchID+"/<service>", searchEndpoint, adminKey)
	return &apiSet, nil
}
func (e *AzureSearchServiceExpander) getAPISetForCluster(searchID string) *SwaggerAPISetSearch {

	swaggerAPISet := GetSwaggerResourceExpander().GetAPISet(searchID + "/<service>")
	if swaggerAPISet == nil {
		return nil
	}
	apiSet := (*swaggerAPISet).(SwaggerAPISetSearch)
	return &apiSet
}

func (e *AzureSearchServiceExpander) getAdminKey(ctx context.Context, searchID string) (string, error) {
	data, err := e.client.DoRequest(ctx, "POST", searchID+"/listAdminKeys?api-version=2015-08-19")
	if err != nil {
		return "", fmt.Errorf("Failed to get admin key: " + err.Error() + searchID)
	}

	var response adminKeysResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, searchID)
		return "", err
	}

	adminKey := response.PrimaryKey

	if adminKey == "" {
		return "", fmt.Errorf("Failed to get admin key")
	}

	return adminKey, nil
}

func (e *AzureSearchServiceExpander) getSearchEndpoint(ctx context.Context, searchID string) (string, error) {
	data, err := e.client.DoRequest(ctx, "GET", searchID+"?api-version=2015-08-19")
	if err != nil {
		return "", fmt.Errorf("Failed to get search service data: " + err.Error() + searchID)
	}

	var response searchServiceResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, searchID)
		return "", err
	}

	searchServiceName := response.Name

	if searchServiceName == "" {
		return "", fmt.Errorf("Search service name lookup failed")
	}

	searchServiceEndpoint := fmt.Sprintf("https://%s.search.windows.net", searchServiceName)

	return searchServiceEndpoint, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e AzureSearchServiceExpander) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

func (e *AzureSearchServiceExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
