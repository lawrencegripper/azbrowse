package expanders

import (
	"context"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// The storage API currently doesn't provide a way to list management policies for a storage account
// I.e. no GET /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies
// This means that there is no way to navigate via swagger from the account to the management policies
// Any policy that exists also currently has to have the name "default" (i.e. can only actually have 0 or 1)
// This expander adds the link from storage account to management policy

// Check interface
var _ Expander = &StorageManagementPoliciesExpander{}

// StorageManagementPoliciesExpander expands The default management policy under a storage account
type StorageManagementPoliciesExpander struct {
	ExpanderBase
}

func (e *StorageManagementPoliciesExpander) setClient(c *armclient.Client) {
	// noop
}

// Name returns the name of the expander
func (e *StorageManagementPoliciesExpander) Name() string {
	return "StorageManagementPoliciesExpander"
}

// DoesExpand checks if this is a storage account
func (e *StorageManagementPoliciesExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}" {
			return true, nil
		}
	}
	return false, nil
}

// Expand returns ManagementPolicies in the StorageAccount
func (e *StorageManagementPoliciesExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	isPrimaryResponse := false

	newItems := []*TreeNode{}

	swaggerResourceType := currentItem.SwaggerResourceType
	if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}" {

		matchResult := swaggerResourceType.Endpoint.Match(currentItem.ExpandURL) // TODO - return the matches from getHandledTypeForURL to avoid re-calculating!
		templateValues := matchResult.Values

		if swaggerResourceType.SubResources != nil {
			for _, resourceType := range swaggerResourceType.SubResources {
				if resourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}" {
					// got the management policies endpoint
					// now create the expand URL etc etc

					defaultPolicyTemplateValues := make(map[string]string)
					for k, v := range templateValues {
						defaultPolicyTemplateValues[k] = v
					}
					defaultPolicyTemplateValues["managementPolicyName"] = "default"

					url, err := resourceType.Endpoint.BuildURL(defaultPolicyTemplateValues)
					if err != nil {
						return ExpanderResult{
							Err: err,
						}
					}
					resourceTypeRef := resourceType
					newItems = append(newItems, &TreeNode{
						Parentid:            currentItem.ID,
						Namespace:           "storageManagementPolicies",
						Name:                "ManagementPolicy",
						Display:             "Management Policy",
						ItemType:            SubResourceType,
						ExpandURL:           url,
						SwaggerResourceType: &resourceTypeRef,
					})

					break
				}
			}
		}
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: ""},
		SourceDescription: "StorageManagementPoliciesExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: isPrimaryResponse,
	}
}

func (e *StorageManagementPoliciesExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
