package expanders

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"sort"
	"strings"

	"github.com/go-logr/logr"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tfprovider"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

var vmEndpoint = endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2020-06-01")

// lookup mapping of resource type to regexp expression to match resource IDs
var tfimportBaseConfig = map[string]*endpoints.EndpointInfo{
	"azurerm_resource_group": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}", ""),

	// App Service
	"azurerm_app_service_plan": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{farmName}", ""),
	"azurerm_app_service":      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}", ""),

	// Storage
	"azurerm_storage_account": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", ""),

	// SQL Database
	"azurerm_mssql_server":   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", ""),
	"azurerm_mssql_database": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", ""),

	// Networking
	"azurerm_private_endpoint":                      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{endpointName}", ""),
	"azurerm_network_interface":                     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", ""),
	"azurerm_network_security_group":                endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", ""),
	"azurerm_network_security_rule":                 endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", ""),
	"azurerm_private_dns_zone":                      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", ""),
	"azurerm_private_dns_zone_virtual_network_link": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", ""),
	"azurerm_public_ip":                             endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", ""),
	"azurerm_virtual_network_gateway":               endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", ""),
	"azurerm_virtual_network":                       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", ""),
	"azurerm_subnet":                                endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", ""),

	// KeyVault
	"azurerm_key_vault": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", ""),

	// Virtual machines
	"azurerm_virtual_machine_extension": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", ""),
	"azurerm_managed_disk":              endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", ""),
	"azurerm_windows_virtual_machine":   vmEndpoint,
	"azurerm_linux_virtual_machine":     vmEndpoint,

	// "":              endpoints.MustGetEndpointInfoFromURL("", ""),

}

// TODO
//  - expand the list of mapped resource types
//  - handle Azure data-plane resources (e.g. storage containers)
//  - handle non-Azure resources? (e.g. databricks cluster)

const (
	tfimportActionGetTerraform = "GetTerraform"
)

// NewTerraformImportExpander creates a new instance of TerraformImportExpander
func NewTerraformImportExpander(armclient *armclient.Client) *TerraformImportExpander {
	return &TerraformImportExpander{
		nullLogger: NewNullLogger(),
		client:     armclient,
	}
}

// Check interface
var _ Expander = &TerraformImportExpander{}

// TerraformImportExpander provides actions
type TerraformImportExpander struct {
	ExpanderBase
	tfProvider *tfprovider.TerraformProvider
	nullLogger logr.Logger
	client     *armclient.Client
}

func (e *TerraformImportExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *TerraformImportExpander) Name() string {
	return "TerraformImportExpander"
}

func (e *TerraformImportExpander) ensureTfProviderInitialized(context context.Context) error {
	if e.tfProvider != nil {
		return nil
	}

	// Get a provider instance by installing or using existing binary
	azbPath := "/root/.azbrowse/terraform/"
	user, err := user.Current()
	if err == nil {
		azbPath = user.HomeDir + "/.azbrowse/terraform/"
	}
	err = os.MkdirAll(azbPath, 0777)
	if err != nil {
		return err
	}

	config := tfprovider.TerraformProviderConfig{
		ProviderName:      "azurerm",
		ProviderVersion:   "2.38.0",
		ProviderConfigHCL: "features {}",
		ProviderPath:      azbPath,
	}
	provider, err := tfprovider.SetupProvider(context, e.nullLogger, config) // TODO - update to use azbrowse profile folder as cache
	if err != nil {
		return err
	}
	e.tfProvider = provider
	return nil
}

func (e *TerraformImportExpander) getResourceTypeNameFromResourceID(context context.Context, resourceID string) (string, error) {
	result := vmEndpoint.Match(resourceID)
	if result.IsMatch {
		body, err := e.client.DoRequest(context, "GET", resourceID+"?api-version="+vmEndpoint.APIVersion)
		if err != nil {
			return "", err
		}
		value, err := getJSONPropertyFromString(body, "properties", "storageProfile", "osDisk", "osType")
		if err != nil {
			return "", err
		}
		osType, ok := value.(string)
		if !ok {
			return "", err
		}
		switch osType {
		case "Windows":
			return "azurerm_windows_virtual_machine", nil
		case "Linux":
			return "azurerm_linux_virtual_machine", nil
		}
	}

	for resourceTypeName, resourceEndpoint := range tfimportBaseConfig {
		result := resourceEndpoint.Match(resourceID)
		if result.IsMatch {
			return resourceTypeName, nil
		}
	}
	return "", nil
}

func getJSONPropertyFromString(jsonString string, properties ...string) (interface{}, error) {
	var jsonData map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		return nil, err
	}

	return getJSONProperty(jsonData, properties...)
}
func getJSONProperty(jsonData interface{}, properties ...string) (interface{}, error) {
	switch jsonData.(type) {
	case map[string]interface{}:
		jsonMap := jsonData.(map[string]interface{})
		name := properties[0]
		jsonSubtree, ok := jsonMap[name]
		if ok {
			if len(properties) == 1 {
				return jsonSubtree, nil
			}
			return getJSONProperty(jsonSubtree, properties[1:]...)
		} else {
			return nil, nil // TODO - error if not found?
		}
	default:
		return nil, nil // TODO - error if not able to walk the tree?
	}

}

// HasActions is a default implementation returning false to indicate no actions available
func (e *TerraformImportExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	resourceTypeName, err := e.getResourceTypeNameFromResourceID(context, item.ID)
	if err != nil {
		return false, err
	}
	if resourceTypeName == "" {
		return false, nil
	}
	if item.Metadata == nil {
		item.Metadata = map[string]string{}
	}
	item.Metadata["TerraformImportExpander_ResourceTypeName"] = resourceTypeName // cache to avoid repeating lookup (avoids ARM call in VM case)
	return true, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *TerraformImportExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {

	resourceTypeName := item.Metadata["TerraformImportExpander_ResourceTypeName"]
	if resourceTypeName == "" {
		return ListActionsResult{
			SourceDescription: "TerraformImportExpander",
			Err:               fmt.Errorf("ResourceTypeName not set"),
		}
	}

	nodes := []*TreeNode{
		{
			Parentid:              item.ID,
			ID:                    item.ID + "?terraform-import",
			Namespace:             "tfimport",
			Name:                  "Get Terraform",
			Display:               "Get Terraform",
			ItemType:              ActionType,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"ActionID":         tfimportActionGetTerraform,
				"ResourceTypeName": resourceTypeName,
			},
		},
	}
	return ListActionsResult{
		Nodes:             nodes,
		SourceDescription: "TerraformImportExpander",
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *TerraformImportExpander) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	actionID := item.Metadata["ActionID"]

	var err error
	switch actionID {
	case tfimportActionGetTerraform:
		return e.getTerraformForNode(context, item)
	case "":
		err = fmt.Errorf("ActionID metadata not set: %q", item.ID)
	default:
		err = fmt.Errorf("Unhandled ActionID: %q", actionID)
	}
	return ExpanderResult{
		SourceDescription: "TerraformImportExpander",
		Err:               err,
	}
}

func (e *TerraformImportExpander) getTerraformForNode(context context.Context, item *TreeNode) ExpanderResult {
	span, context := tracing.StartSpanFromContext(context, "terraform:get-for-node:"+item.ItemType+":"+item.Name, tracing.SetTag("item", item))
	defer span.Finish()
	err := e.ensureTfProviderInitialized(context)
	if err != nil {
		return ExpanderResult{
			SourceDescription: "TerraformImportExpander",
			Err:               err,
		}
	}

	resourceTypeName := item.Metadata["ResourceTypeName"]
	if resourceTypeName == "" {
		return ExpanderResult{
			SourceDescription: "TerraformImportExpander",
			Err:               fmt.Errorf("ResourceTypeName not set"),
		}
	}

	id := item.ID
	endpoint := tfimportBaseConfig[resourceTypeName]
	endpointMatch := endpoint.Match(id)
	if !endpointMatch.IsMatch {
		return ExpanderResult{
			SourceDescription: "TerraformImportExpander",
			Err:               fmt.Errorf("Failed to match resource type name"),
		}
	}
	// use the endpoint to rebuild the ID URL to ensure the case matches what the azurerm provider expects
	id, err = endpoint.BuildURL(endpointMatch.Values)
	if !endpointMatch.IsMatch {
		return ExpanderResult{
			SourceDescription: "TerraformImportExpander",
			Err:               err,
		}
	}

	terraform, err := e.getTerraformFor(context, id, resourceTypeName)
	if err != nil {
		return ExpanderResult{
			SourceDescription: "TerraformImportExpander",
			Err:               err,
		}
	}
	return ExpanderResult{
		SourceDescription: "TerraformImportExpander",
		Response: ExpanderResponse{
			ResponseType: ResponseTerraform,
			Response:     terraform,
		},
		IsPrimaryResponse: true,
	}
}

func (e *TerraformImportExpander) getTerraformFor(context context.Context, id string, resourceTypeName string) (string, error) {
	span, context := tracing.StartSpanFromContext(context, "terraform:get-schema:"+resourceTypeName)
	defer span.Finish()
	terraformProviderSchema := e.tfProvider.Plugin.GetSchema()
	importRequest := providers.ImportResourceStateRequest{
		TypeName: resourceTypeName,
		ID:       id,
	}
	spanImport, context := tracing.StartSpanFromContext(context, "terraform:import:"+resourceTypeName)
	defer spanImport.Finish()
	importResponse := e.tfProvider.Plugin.ImportResourceState(importRequest)

	result := ""
	for _, resource := range importResponse.ImportedResources {
		span, _ := tracing.StartSpanFromContext(context, "terraform:read:"+resource.TypeName)
		readRequest := providers.ReadResourceRequest{
			TypeName:   resource.TypeName,
			PriorState: resource.State,
		}
		readResponse := e.tfProvider.Plugin.ReadResource(readRequest)
		span.Finish()

		resourceSchema := terraformProviderSchema.ResourceTypes[resource.TypeName]

		hclString, err := e.printState(readResponse.NewState, resource.TypeName, resourceSchema)
		if err != nil {
			return "", err
		}
		result += hclString
		result += "\n"
	}

	return result, nil
}

func (e *TerraformImportExpander) writeBlock(outputBlock *hclwrite.Block, terraformBlock *configschema.Block, state cty.Value) {
	// Sort attribute names:
	//    id
	//    name
	//    location
	//    resource_group_name
	//    <names ending in `_id`/`_ids`>
	//    <other names>
	keys := make([]string, 0, len(terraformBlock.Attributes))
	for k := range terraformBlock.Attributes {
		prefix := "z"
		switch {
		case k == "id":
			prefix = "a"
		case k == "name":
			prefix = "b"
		case k == "location":
			prefix = "c"
		case k == "resource_group_name":
			prefix = "d"
		case strings.HasSuffix(k, "_id") || strings.HasSuffix(k, "_ids"):
			prefix = "e"
		}
		keys = append(keys, prefix+k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		attributeName := k[1:]
		attributeSchema := terraformBlock.Attributes[attributeName]
		if !attributeSchema.Computed {
			attributeValue := state.GetAttr(attributeName)
			outputBlock.Body().SetAttributeValue(attributeName, attributeValue)
		}
	}

	for blockName, blockSchema := range terraformBlock.BlockTypes {
		// TODO - might need to look at blockSchema.Nesting and handle accordingly
		newState := state.GetAttr(blockName)

		if newState.Type().IsObjectType() {
			newBlock := outputBlock.Body().AppendNewBlock(blockName, []string{})
			e.writeBlock(newBlock, &blockSchema.Block, newState)
		}
		if newState.CanIterateElements() {
			iterator := newState.ElementIterator()
			for iterator.Next() {
				_, value := iterator.Element()
				if value.Type().IsObjectType() {
					newBlock := outputBlock.Body().AppendNewBlock(blockName, []string{})
					e.writeBlock(newBlock, &blockSchema.Block, value)
				}
			}
		}
	}
}
func (e *TerraformImportExpander) printState(state cty.Value, resourceTypeName string, schema providers.Schema) (string, error) {
	file := hclwrite.NewEmptyFile()
	terraformBlock := schema.Block
	name := "todo_resource_name"
	if state.Type().HasAttribute("name") {
		attribute := state.GetAttr("name")
		name = strings.ReplaceAll(attribute.AsString(), "-", "_")
	}
	block := file.Body().AppendNewBlock("resource", []string{resourceTypeName, name})
	e.writeBlock(block, terraformBlock, state)

	var buf bytes.Buffer
	_, err := file.WriteTo(&buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ************************************************************************************** TODO put in a separate file
type NullLogger struct {
}

var _ logr.Logger = &NullLogger{}

// Info implements Logger.Info
func (l *NullLogger) Info(msg string, kvs ...interface{}) {
}

// Enabled implements Logger.Enabled
func (*NullLogger) Enabled() bool {
	return false
}

// Error implements Logger.Error
func (l *NullLogger) Error(err error, msg string, kvs ...interface{}) {
	kvs = append(kvs, "error", err)
	l.Info(msg, kvs...)
}

// V implements Logger.V
func (l *NullLogger) V(_ int) logr.Logger {
	return l
}

// WithName implements Logger.WithName
func (l *NullLogger) WithName(name string) logr.Logger {
	return l
}

// WithValues implements Logger.WithValues
func (l *NullLogger) WithValues(kvs ...interface{}) logr.Logger {
	return l
}

// NewNullLogger creates a new NullLogger
func NewNullLogger() logr.Logger {
	return &NullLogger{}
}
