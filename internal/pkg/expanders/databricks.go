package expanders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

const azureDatabricksTemplateURL string = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}"

// Auth docs: https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token

// This is the Azure AD application ID of the global databricks application
// this is constant for all tentants/subscriptions as owned by databricks team
const azureDatabricksGlobalApplicationID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"

// This is the azure management endpoint used. Would need updating for sovereign clouds etc
const azureManagementEndpoint string = "https://management.core.windows.net/"

type workspaceResponse struct {
	Properties struct {
		WorkspaceURL string `json:"workspaceUrl"`
	} `json:"properties"`
}

// AzureDatabricksExpander expands the kubernetes aspects of AKS
type AzureDatabricksExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *AzureDatabricksExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *AzureDatabricksExpander) Name() string {
	return "AzureDatabricksExpander"
}

// DoesExpand checks if this is a storage account
func (e *AzureDatabricksExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == azureDatabricksTemplateURL {
			return true, nil
		}
	}
	if currentItem.Namespace == "AzureDatabricksExpander" {
		return true, nil
	}
	return false, nil
}

// Expand returns the SwaggerExpander set up for this Databricks workspace
func (e *AzureDatabricksExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "AzureDatabricksExpander" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == azureDatabricksTemplateURL {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			ID:        currentItem.ID + "/<workspace>",
			Parentid:  currentItem.ID,
			Namespace: "AzureDatabricksExpander",
			Name:      "Connect to Databricks workspace",
			Display:   "Connect to Databricks workspace",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"WorkspaceID":           currentItem.ID,
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "AzureDatabricksExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "AzureDatabricksExpander" && currentItem.ItemType == SubResourceType {
		return e.expandWorkspaceRoot(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "AzureDatabricksExpander request",
	}
}

func (e *AzureDatabricksExpander) expandWorkspaceRoot(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	workspaceID := currentItem.Metadata["WorkspaceID"]

	// Check for existing config for the cluster
	apiSet := e.getAPISetForWorkspace(workspaceID)
	var err error
	if apiSet == nil {
		apiSet, err = e.createAPISetForWorkspace(ctx, workspaceID)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				Response:          ExpanderResponse{Response: "Error!"},
				SourceDescription: "AzureDatabricksExpander request",
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
		queryString := ""
		if resourceType.Endpoint.TemplateURL == "/api/2.0/workspace/list" {
			queryString = "?path=/"
		}
		newItems = append(newItems, &TreeNode{
			Parentid:            currentItem.ID,
			ID:                  currentItem.ID + "/" + display,
			Namespace:           "swagger",
			Name:                display,
			Display:             display,
			ExpandURL:           resourceType.Endpoint.TemplateURL + queryString, // all fixed template URLs or have a starting query string
			ItemType:            SubResourceType,
			SwaggerResourceType: &resourceType,
			Metadata: map[string]string{
				"SwaggerAPISetID": currentItem.ID,
			},
		})
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: ""},
		SourceDescription: "AzureDatabricksExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureDatabricksExpander) createAPISetForWorkspace(ctx context.Context, workspaceID string) (*SwaggerAPISetDatabricks, error) {

	managementToken, err := armclient.AcquireTokenForResourceFromAzCLI(azureManagementEndpoint)
	if err != nil {
		return nil, err
	}
	databricksToken, err := armclient.AcquireTokenForResourceFromAzCLI(azureDatabricksGlobalApplicationID)
	if err != nil {
		return nil, err
	}

	workspaceURL, err := e.getWorkspaceUrl(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	swaggerResourceTypes := e.loadResourceTypes()
	if err != nil {
		return nil, err
	}

	// Register the swagger config so that the swagger expander can take over
	apiSet := NewSwaggerAPISetDatabricks(swaggerResourceTypes, workspaceID, workspaceID+"/<workspace>", workspaceURL, managementToken.AccessToken, databricksToken.AccessToken)
	return &apiSet, nil
}
func (e *AzureDatabricksExpander) getAPISetForWorkspace(workspaceID string) *SwaggerAPISetDatabricks {

	swaggerAPISet := GetSwaggerResourceExpander().GetAPISet(workspaceID + "/<workspace>")
	if swaggerAPISet == nil {
		return nil
	}
	apiSet := (*swaggerAPISet).(SwaggerAPISetDatabricks)
	return &apiSet
}

func (e *AzureDatabricksExpander) getWorkspaceUrl(ctx context.Context, workspaceID string) (string, error) {
	data, err := e.client.DoRequest(ctx, "GET", workspaceID+"?api-version=2018-04-01")
	if err != nil {
		return "", fmt.Errorf("Failed to get workspace data: " + err.Error() + workspaceID)
	}

	var response workspaceResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, workspaceID)
		return "", err
	}

	workspaceURL := response.Properties.WorkspaceURL

	if workspaceURL == "" {
		return "", fmt.Errorf("Workspace URL lookup failed")
	}

	return workspaceURL, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e AzureDatabricksExpander) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

func (e *AzureDatabricksExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
