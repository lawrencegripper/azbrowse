package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/pkg/endpoints"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// SwaggerResourceExpander expands resource under an AppService
type SwaggerResourceExpander struct {
	initialized   bool
	ResourceTypes []ResourceType
}

// ResourceType holds information about resources that can be displayed
type ResourceType struct {
	Display        string
	Endpoint       endpoints.EndpointInfo
	Verb           string
	SupportsDelete bool
	Children       []ResourceType // Children are auto-loaded (must be able to build the URL => no additional template URL values)
	SubResources   []ResourceType // SubResources are not auto-loaded (these come from the request to the endpoint)
}

// Name returns the name of the expander
func (e *SwaggerResourceExpander) Name() string {
	return "SwaggerResourceExpander"
}

func mustGetEndpointInfoFromURL(url string, apiVersion string) endpoints.EndpointInfo {
	endpoint, err := endpoints.GetEndpointInfoFromURL(url, apiVersion)
	if err != nil {
		panic(err)
	}
	return endpoint
}

func getResourceTypeForURL(ctx context.Context, url string, resourceTypes []ResourceType) *ResourceType {
	span, ctx := tracing.StartSpanFromContext(ctx, "getResourceTypeForURL:"+url)
	defer span.Finish()
	return getResourceTypeForURLInner(url, resourceTypes)
}
func getResourceTypeForURLInner(url string, resourceTypes []ResourceType) *ResourceType {
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
func (e *SwaggerResourceExpander) ensureInitialized() {
	if !e.initialized {
		e.ResourceTypes = e.getResourceTypes()
		e.initialized = true
	}
}

// DoesExpand checks if this is an RG
func (e *SwaggerResourceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	e.ensureInitialized()
	if currentItem.ItemType == resourceType {
		if currentItem.SwaggerResourceType != nil {
			return true, nil
		}
		resourceType := getResourceTypeForURL(ctx, currentItem.ExpandURL, e.ResourceTypes)
		if resourceType != nil {
			currentItem.SwaggerResourceType = resourceType // cache to avoid looking up in Expand
			return true, nil
		}
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

	method := resourceType.Verb
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Nodes:    nil,
			Response: string(data),
			Err:      fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}

	newItems := []*TreeNode{}
	matchResult := resourceType.Endpoint.Match(currentItem.ExpandURL) // TODO - return the matches from getHandledTypeForURL to avoid re-calculating!
	templateValues := matchResult.Values

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the ARM response and add these to newItems
		var resourceResponse armclient.ResourceResponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
			panic(err)
		}
		for _, resource := range resourceResponse.Resources {
			subResourceType := getResourceTypeForURL(ctx, resource.ID, resourceType.SubResources)
			if subResourceType == nil {
				panic(fmt.Errorf("SubResource type not found! %s", resource.ID))
			}
			subResourceTemplateValues := subResourceType.Endpoint.Match(resource.ID).Values
			name := substituteValues(subResourceType.Display, subResourceTemplateValues)
			deleteURL := ""
			if subResourceType.SupportsDelete {
				deleteURL = currentItem.ID
			}
			newItems = append(newItems, &TreeNode{
				Parentid:            currentItem.ID,
				Namespace:           "swagger",
				Name:                name,
				Display:             name,
				ExpandURL:           resource.ID + "?api-version=" + subResourceType.Endpoint.APIVersion,
				ItemType:            "resource",
				DeleteURL:           deleteURL,
				SwaggerResourceType: subResourceType,
			})
		}
	}
	// Add any children to newItems
	for _, child := range resourceType.Children {
		loopChild := child
		url, err := child.Endpoint.BuildURL(templateValues)
		if err != nil {
			err = fmt.Errorf("Error building URL: %s\nURL:%s", child.Display, err)
			panic(err)
		}
		display := substituteValues(child.Display, templateValues)
		deleteURL := ""
		if child.SupportsDelete {
			deleteURL = currentItem.ID
		}
		newItems = append(newItems, &TreeNode{
			Parentid:            currentItem.ID,
			Namespace:           "swagger",
			Name:                display,
			Display:             display,
			ExpandURL:           url,
			ItemType:            "resource",
			DeleteURL:           deleteURL,
			SwaggerResourceType: &loopChild,
		})
	}

	return ExpanderResult{
		Nodes:             newItems,
		Response:          string(data),
		IsPrimaryResponse: true, // only returning items that we are the primary response for
	}
}

// substituteValues applys a value map to strings such as "Name: {name}"
func substituteValues(fmtString string, values map[string]string) string {
	for name, value := range values {
		fmtString = strings.Replace(fmtString, "{"+name+"}", value, -1)
	}
	return fmtString
}
