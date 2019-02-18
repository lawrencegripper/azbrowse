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

// AppServiceResourceExpander expands resource under an AppService
type AppServiceResourceExpander struct {
	initialized  bool
	handledTypes []handledType
}

type handledType struct {
	display  string
	endpoint endpoints.EndpointInfo
	// children are auto-loaded (must be able to build the URL => no additional template URL values)
	children []handledType
	// subResources are not auto-loaded (these come from the request to the endpoint)
	subResources []handledType
}

// Name returns the name of the expander
func (e *AppServiceResourceExpander) Name() string {
	return "AppServiceResourceExpander"
}

func mustGetEndpointInfoFromURL(url string, apiVersion string) endpoints.EndpointInfo {
	return mustGetEndpointInfoFromURLWithVerb(url, apiVersion, "GET")
}
func mustGetEndpointInfoFromURLWithVerb(url string, apiVersion string, verb string) endpoints.EndpointInfo {
	endpoint, err := endpoints.GetEndpointInfoFromURL(url, apiVersion)
	if err != nil {
		panic(err)
	}
	endpoint.Verb = verb
	return endpoint
}

func getHandledTypeForURL(url string, handledTypes []handledType) *handledType {
	for _, handledType := range handledTypes {
		matchResult := handledType.endpoint.Match(url)
		if matchResult.IsMatch {
			return &handledType
		}
		if result := getHandledTypeForURL(url, handledType.subResources); result != nil {
			return result
		}
		if result := getHandledTypeForURL(url, handledType.children); result != nil {
			return result
		}
	}
	return nil
}
func (e *AppServiceResourceExpander) ensureInitialized() {
	if !e.initialized {
		e.handledTypes = e.getHandledTypes()
		e.initialized = true
	}
}

// DoesExpand checks if this is an RG
func (e *AppServiceResourceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	e.ensureInitialized()
	if currentItem.ItemType == resourceType {
		item := getHandledTypeForURL(currentItem.ExpandURL, e.handledTypes)
		if item != nil {
			return true, nil
		}
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *AppServiceResourceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	span, ctx := tracing.StartSpanFromContext(ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name+":"+currentItem.ID, tracing.SetTag("item", currentItem))
	defer span.Finish()

	handledType := getHandledTypeForURL(currentItem.ExpandURL, e.handledTypes)
	if handledType == nil {
		panic(fmt.Errorf("Node Item not found"))
	}

	method := handledType.endpoint.Verb
	data, err := armclient.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Nodes:    nil,
			Response: string(data),
			Err:      fmt.Errorf("Failed" + err.Error() + currentItem.ExpandURL),
		}
	}

	newItems := []*TreeNode{}
	matchResult := handledType.endpoint.Match(currentItem.ExpandURL) // TODO - return the matches from getHandledTypeForURL to avoid re-calculating!
	templateValues := matchResult.Values

	if len(handledType.subResources) > 0 {
		// We have defined subResources - Unmarshal the ARM response and add these to newItems
		var resourceResponse armclient.ResourceResponse
		err = json.Unmarshal([]byte(data), &resourceResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
			panic(err)
		}
		for _, resource := range resourceResponse.Resources {
			subResourceEndpoint := getHandledTypeForURL(resource.ID, handledType.subResources)
			subResourceTemplateValues := subResourceEndpoint.endpoint.Match(resource.ID).Values
			name := substituteValues(subResourceEndpoint.display, subResourceTemplateValues)
			newItems = append(newItems, &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "appservice",
				Name:      name,
				Display:   name,
				ExpandURL: resource.ID + "?api-version=" + subResourceEndpoint.endpoint.APIVersion,
				ItemType:  resourceType,
				DeleteURL: "TODO",
			})
		}
	}
	// Add any children to newItems
	for _, child := range handledType.children {

		url, err := child.endpoint.BuildURL(templateValues)
		if err != nil {
			err = fmt.Errorf("Error building URL: %s\nURL:%s", child.display, err)
			panic(err)
		}
		display := substituteValues(child.display, templateValues)
		newItems = append(newItems, &TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "appservice",
			Name:      display,
			Display:   display,
			ExpandURL: url,
			ItemType:  resourceType,
			DeleteURL: "NotSupported",
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
