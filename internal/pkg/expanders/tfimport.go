package expanders

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"sort"
	"strings"

	"github.com/go-logr/logr"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tfprovider"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

// lookup mapping of resource type to regexp expression to match resource IDs
var tfimportBaseConfig = map[string]*endpoints.EndpointInfo{
	"azurerm_resource_group":   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}", ""),
	"azurerm_app_service_plan": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{farmName}", ""),
	"azurerm_app_service":      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}", ""),
	"azurerm_storage_account":  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", ""),
	"azurerm_mssql_server":     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", ""),
	"azurerm_private_endpoint": endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{endpointName}", ""),
}

// TODO
//  - expand the list of mapped resource types
//  - handle VMs - need more than the ID to determine Windows/Linux VM
//  - handle non ARM resources (e.g. storage containers)

// TfImportResourceType represents a mapping of resource name to resource ID regexp
type TfImportResourceType struct {
	ResourceName string
	IDRegexp     *regexp.Regexp
}

const (
	tfimportActionGetTerraform = "GetTerraform"
)

// NewTerraformImportExpander creates a new instance of TerraformImportExpander
func NewTerraformImportExpander() *TerraformImportExpander {
	return &TerraformImportExpander{
		nullLogger: NewNullLogger(),
	}
}

// Check interface
var _ Expander = &TerraformImportExpander{}

// TerraformImportExpander provides actions
type TerraformImportExpander struct {
	ExpanderBase
	tfProvider *tfprovider.TerraformProvider
	nullLogger logr.Logger
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

func (e *TerraformImportExpander) getResourceTypeNameFromResourceID(resourceID string) (string, error) {
	for resourceTypeName, resourceEndpoint := range tfimportBaseConfig {
		result := resourceEndpoint.Match(resourceID)
		if result.IsMatch {
			return resourceTypeName, nil
		}
	}
	return "", nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *TerraformImportExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	resourceTypeName, err := e.getResourceTypeNameFromResourceID(item.ID)
	if err != nil {
		return false, err
	}
	if resourceTypeName == "" {
		return false, nil
	}
	return true, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *TerraformImportExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {

	resourceTypeName, err := e.getResourceTypeNameFromResourceID(item.ID)
	if err != nil {
		return ListActionsResult{
			SourceDescription: "TerraformImportExpander",
			Err:               err,
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

	terraform, err := e.getTerraformFor(id, resourceTypeName)
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

func (e *TerraformImportExpander) getTerraformFor(id string, resourceTypeName string) (string, error) {
	terraformProviderSchema := e.tfProvider.Plugin.GetSchema()
	importRequest := providers.ImportResourceStateRequest{
		TypeName: resourceTypeName,
		ID:       id,
	}
	importResponse := e.tfProvider.Plugin.ImportResourceState(importRequest)

	result := ""
	for _, resource := range importResponse.ImportedResources {
		readRequest := providers.ReadResourceRequest{
			TypeName:   resource.TypeName,
			PriorState: resource.State,
		}
		readResponse := e.tfProvider.Plugin.ReadResource(readRequest)

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
