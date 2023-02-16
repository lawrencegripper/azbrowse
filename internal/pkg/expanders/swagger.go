package expanders

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

// Check interface
var _ Expander = &SwaggerResourceExpander{}

// SwaggerAPISet represents the configuration for a set of swagger API endpoints that the SwaggerResourceExpander can handle
type SwaggerAPISet interface {
	ID() string
	GetResourceTypes() []swagger.ResourceType
	AppliesToNode(node *TreeNode) bool
	ExpandResource(context context.Context, node *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error)
	MatchChildNodesByName() bool
	Delete(context context.Context, node *TreeNode) (bool, error)
	Update(context context.Context, node *TreeNode, content string) error
}

// SubResource is used to pass sub resource information from SwaggerAPISet to the expander
type SubResource struct {
	ID           string
	Name         string
	ResourceType swagger.ResourceType
	ExpandURL    string
	DeleteURL    string
	Metadata     map[string]string
}

// APISetExpandResponse returns the result of expanding a Resource
type APISetExpandResponse struct {
	Response      string
	ResponseType  interfaces.ExpanderResponseType
	SubResources  []SubResource
	ChildMetadata map[string]string
}

// SwaggerResourceExpander expands resource under an AppService
type SwaggerResourceExpander struct {
	ExpanderBase
	apiSets map[string]*SwaggerAPISet
}

// NewSwaggerResourcesExpander creates a new SwaggerResourceExpander
func NewSwaggerResourcesExpander() *SwaggerResourceExpander {
	return &SwaggerResourceExpander{
		apiSets: map[string]*SwaggerAPISet{},
	}
}

// AddAPISet adds a SwaggerAPISet to the APIs that the expander will handle
func (e *SwaggerResourceExpander) AddAPISet(apiSet SwaggerAPISet) {
	e.apiSets[apiSet.ID()] = &apiSet
}

// GetAPISet returns a SwaggerAPISet by id
func (e *SwaggerResourceExpander) GetAPISet(id string) *SwaggerAPISet {
	return e.apiSets[id]
}

// Name returns the name of the expander
func (e *SwaggerResourceExpander) Name() string {
	return "SwaggerResourceExpander"
}

func (e *SwaggerResourceExpander) getAPISetForItem(currentItem *TreeNode) *SwaggerAPISet {

	if currentItem.Metadata == nil {
		currentItem.Metadata = make(map[string]string)
	}
	if apiSetID := currentItem.Metadata["SwaggerAPISetID"]; apiSetID != "" {
		return e.GetAPISet(apiSetID)
	}
	for _, apiSetPtr := range e.apiSets {
		apiSet := *apiSetPtr
		if apiSet.AppliesToNode(currentItem) {
			currentItem.Metadata["SwaggerAPISetID"] = apiSet.ID()
			return apiSetPtr
		}
	}
	return nil
}

// DoesExpand checks if this is an RG
func (e *SwaggerResourceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.SuppressSwaggerExpand {
		return false, nil
	}
	apiSetPtr := e.getAPISetForItem(currentItem)
	if apiSetPtr == nil {
		return false, nil
	}
	apiSet := *apiSetPtr

	if currentItem.SwaggerResourceType != nil {
		return true, nil
	}
	resourceType := swagger.GetResourceTypeForURL(ctx, currentItem.ExpandURL, apiSet.GetResourceTypes())
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

	apiSetPtr := e.getAPISetForItem(currentItem)
	if apiSetPtr == nil {
		panic(fmt.Errorf("SwaggerAPISet not set"))
	}
	apiSet := *apiSetPtr

	data := ""
	newItems := []*TreeNode{}

	// Get sub resources from config
	expandResult, err := apiSet.ExpandResource(ctx, currentItem, *resourceType)
	if err != nil {
		return ExpanderResult{
			Nodes:             nil,
			Response:          ExpanderResponse{Response: expandResult.Response, ResponseType: expandResult.ResponseType},
			Err:               err,
			SourceDescription: "SwaggerResourceExpander",
		}
	}
	data = expandResult.Response
	dataType := expandResult.ResponseType

	if len(expandResult.SubResources) > 0 {
		for _, subResource := range expandResult.SubResources {
			metadata := map[string]string{
				"SwaggerAPISetID": apiSet.ID(),
			}
			e.copyMetadata(metadata, subResource.Metadata)
			newItems = append(newItems, &TreeNode{
				Parentid:            currentItem.ID,
				Namespace:           "swagger",
				Name:                subResource.Name,
				Display:             subResource.Name,
				ID:                  subResource.ID,
				ExpandURL:           subResource.ExpandURL,
				ItemType:            SubResourceType,
				DeleteURL:           subResource.DeleteURL,
				SwaggerResourceType: &subResource.ResourceType,
				Metadata:            metadata,
			})
		}
	}
	childMetadata := expandResult.ChildMetadata

	// Add any children to newItems
	if len(resourceType.Children) > 0 {
		matchResult := resourceType.Endpoint.Match(currentItem.ExpandURL)
		templateValues := matchResult.Values
		for _, child := range resourceType.Children {
			loopChild := child

			var url string
			var err error
			if apiSet.MatchChildNodesByName() {
				url, err = child.Endpoint.BuildURL(templateValues)
			} else {
				valueArray := resourceType.Endpoint.GenerateValueArrayFromMap(templateValues)
				url, err = child.Endpoint.BuildURLFromArray(valueArray)
			}
			if err != nil {
				err = fmt.Errorf("Error building URL: %s\nURL:%s", child.Display, err)
				return ExpanderResult{
					Nodes:             nil,
					Response:          ExpanderResponse{Response: data, ResponseType: dataType},
					Err:               err,
					SourceDescription: "SwaggerResourceExpander",
				}
			}

			display := substituteValues(child.Display, templateValues)
			deleteURL := ""
			if child.DeleteEndpoint != nil {
				if apiSet.MatchChildNodesByName() {
					deleteURL, err = child.DeleteEndpoint.BuildURL(templateValues)
				} else {
					valueArray := child.DeleteEndpoint.GenerateValueArrayFromMap(templateValues)
					deleteURL, err = child.DeleteEndpoint.BuildURLFromArray(valueArray)
				}
				if err != nil {
					err = fmt.Errorf("Error building child delete url '%s': %s", child.DeleteEndpoint.TemplateURL, err)
					return ExpanderResult{
						Nodes:             nil,
						Response:          ExpanderResponse{Response: data, ResponseType: dataType},
						Err:               err,
						SourceDescription: "SwaggerResourceExpander",
					}
				}
			}
			metadata := map[string]string{
				"SwaggerAPISetID": apiSet.ID(),
			}
			e.copyMetadata(metadata, childMetadata)

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
				Metadata:            metadata,
			})
		}
	}

	return ExpanderResult{
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: data, ResponseType: dataType},
		IsPrimaryResponse: true, // only returning items that we are the primary response for
	}
}

func (e *SwaggerResourceExpander) copyMetadata(target map[string]string, source map[string]string) {
	for key, value := range source {
		target[key] = value
	}
}

// CanUpdate indicates if the item can be updated
func (e SwaggerResourceExpander) CanUpdate(ctx context.Context, item *TreeNode) (bool, error) {
	if item == nil ||
		item.SwaggerResourceType == nil ||
		item.SwaggerResourceType.PutEndpoint == nil ||
		item.Metadata == nil ||
		item.Metadata["SwaggerAPISetID"] == "" {
		return false, nil
	}
	return true, nil
}

// Update updates the item
func (e *SwaggerResourceExpander) Update(ctx context.Context, item *TreeNode, updatedContent string) error {
	apiSetID := item.Metadata["SwaggerAPISetID"]
	apiSetPtr := e.GetAPISet(apiSetID)
	if apiSetPtr == nil {
		return nil
	}
	apiSet := *apiSetPtr

	err := apiSet.Update(ctx, item, updatedContent)
	if err != nil {
		return fmt.Errorf("Error updating: %s", err)
	}
	return nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *SwaggerResourceExpander) Delete(context context.Context, item *TreeNode) (bool, error) {

	apiSetPtr := e.getAPISetForItem(item)
	if apiSetPtr == nil {
		return false, nil // false indicates we didn't try to delete
	}
	apiSet := *apiSetPtr

	return apiSet.Delete(context, item)
}

// substituteValues applies a value map to strings such as "Name: {name}"
func substituteValues(fmtString string, values map[string]string) string {
	for name, value := range values {
		fmtString = strings.Replace(fmtString, "{"+name+"}", value, -1)
	}
	return fmtString
}
