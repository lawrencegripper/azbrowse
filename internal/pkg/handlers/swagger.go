package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

// SwaggerConfig represents the configuration for a set of swagger resources that the SwaggerResourceExpander can handle
type SwaggerConfig interface {
	ID() string
	GetResourceTypes() []swagger.SwaggerResourceType
	AppliesToNode(node *TreeNode) bool
	ExpandResource(context context.Context, node *TreeNode, resourceType swagger.SwaggerResourceType) (ConfigExpandResponse, error)
}

// SubResource is used to pass sub resource information from SwaggerConfig to the expander
type SubResource struct {
	ID           string
	Name         string
	ResourceType swagger.SwaggerResourceType
	ExpandURL    string
	DeleteURL    string
}

// ConfigExpandResource returns the result of expanding a Resource
type ConfigExpandResponse struct {
	Response     string
	SubResources []SubResource
}

// SwaggerResourceExpander expands resource under an AppService
type SwaggerResourceExpander struct {
	configs map[string]*SwaggerConfig
}

func NewSwaggerResourcesExpander() *SwaggerResourceExpander {
	return &SwaggerResourceExpander{
		configs: map[string]*SwaggerConfig{},
	}
}

func (e *SwaggerResourceExpander) AddConfig(config SwaggerConfig) {
	e.configs[config.ID()] = &config
}
func (e *SwaggerResourceExpander) GetConfig(id string) *SwaggerConfig {
	return e.configs[id]
}

// Name returns the name of the expander
func (e *SwaggerResourceExpander) Name() string {
	return "SwaggerResourceExpander"
}

func getResourceTypeForURL(ctx context.Context, url string, resourceTypes []swagger.SwaggerResourceType) *swagger.SwaggerResourceType {
	span, _ := tracing.StartSpanFromContext(ctx, "getResourceTypeForURL:"+url)
	defer span.Finish()
	return getResourceTypeForURLInner(url, resourceTypes)
}
func getResourceTypeForURLInner(url string, resourceTypes []swagger.SwaggerResourceType) *swagger.SwaggerResourceType {
	for _, resourceType := range resourceTypes {
		matchResult := resourceType.Endpoint.Match(url)
		if matchResult.IsMatch {
			return &resourceType
		}
		if result := getResourceTypeForURLInner(url, resourceType.SubResources); result != nil {
			return result
		}
		if result := getResourceTypeForURLInner(url, resourceType.Children); result != nil {
			return result
		}
	}
	return nil
}

func (e *SwaggerResourceExpander) getConfigForItem(currentItem *TreeNode) *SwaggerConfig {

	if currentItem.Metadata == nil {
		currentItem.Metadata = make(map[string]string)
	}
	if configID := currentItem.Metadata["SwaggerConfigID"]; configID != "" {
		return e.GetConfig(configID)
	}
	for _, configPtr := range e.configs {
		config := *configPtr
		if config.AppliesToNode(currentItem) {
			currentItem.Metadata["SwaggerConfigID"] = config.ID()
			return configPtr
		}
	}
	return nil
}

// DoesExpand checks if this is an RG
func (e *SwaggerResourceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.Metadata["SuppressSwaggerExpand"] == "true" {
		return false, nil
	}
	configPtr := e.getConfigForItem(currentItem)
	if configPtr == nil {
		return false, nil
	}
	config := *configPtr

	if currentItem.SwaggerResourceType != nil {
		return true, nil
	}
	resourceType := getResourceTypeForURL(ctx, currentItem.ExpandURL, config.GetResourceTypes())
	if resourceType != nil {
		currentItem.SwaggerResourceType = resourceType // cache to avoid looking up in Expand
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *SwaggerResourceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	span, ctx := tracing.StartSpanFromContext(ctx, "expand(swagger):"+currentItem.ItemType+":"+currentItem.Name+":"+currentItem.ID, tracing.SetTag("item", currentItem))
	defer span.Finish()

	resourceType := currentItem.SwaggerResourceType
	if resourceType == nil {
		panic(fmt.Errorf("SwaggerResourceType not set"))
	}

	configPtr := e.getConfigForItem(currentItem)
	if configPtr == nil {
		panic(fmt.Errorf("SwaggerConfig not set"))
	}
	config := *configPtr

	data := ""

	// Get sub resources from config
	expandResult, err := config.ExpandResource(ctx, currentItem, *resourceType)
	if err != nil {
		return ExpanderResult{
			Nodes:             nil,
			Response:          expandResult.Response,
			Err:               err,
			SourceDescription: "SwaggerResourceExpander",
		}
	}
	data = expandResult.Response

	newItems := []*TreeNode{}
	if len(expandResult.SubResources) > 0 {
		for _, subResource := range expandResult.SubResources {
			newItems = append(newItems, &TreeNode{
				Parentid:            currentItem.ID,
				Namespace:           "swagger",
				Name:                subResource.Name,
				Display:             subResource.Name,
				ID:                  subResource.ID,
				ExpandURL:           subResource.ID + "?api-version=" + subResource.ResourceType.Endpoint.APIVersion,
				ItemType:            SubResourceType,
				DeleteURL:           subResource.DeleteURL,
				SwaggerResourceType: &subResource.ResourceType,
				Metadata: map[string]string{
					"SwaggerConfigID": config.ID(),
				},
			})
		}
	}

	// Add any children to newItems
	matchResult := resourceType.Endpoint.Match(currentItem.ExpandURL)
	templateValues := matchResult.Values
	for _, child := range resourceType.Children {
		loopChild := child
		url, err := child.Endpoint.BuildURL(templateValues)
		if err != nil {
			err = fmt.Errorf("Error building URL: %s\nURL:%s", child.Display, err)
			panic(err)
		}
		display := substituteValues(child.Display, templateValues)
		deleteURL := ""
		if child.DeleteEndpoint != nil {
			deleteURL, err = child.DeleteEndpoint.BuildURL(templateValues)
			if err != nil {
				panic(fmt.Errorf("Error building child delete url '%s': %s", child.DeleteEndpoint.TemplateURL, err))
			}
		}
		newItems = append(newItems, &TreeNode{
			Parentid:            currentItem.ID,
			ID:                  currentItem.ID + "/" + display,
			Namespace:           "swagger",
			Name:                display,
			Display:             display,
			ExpandURL:           url,
			ItemType:            SubResourceType,
			DeleteURL:           deleteURL,
			SwaggerResourceType: &loopChild,
			Metadata: map[string]string{
				"SwaggerConfigID": config.ID(),
			},
		})
	}

	return ExpanderResult{
		Nodes:             newItems,
		Response:          data,
		IsPrimaryResponse: true, // only returning items that we are the primary response for
	}
}

// substituteValues applies a value map to strings such as "Name: {name}"
func substituteValues(fmtString string, values map[string]string) string {
	for name, value := range values {
		fmtString = strings.Replace(fmtString, "{"+name+"}", value, -1)
	}
	return fmtString
}
