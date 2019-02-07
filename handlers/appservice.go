package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/endpoints"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/tracing"
)

// AppServiceResourceExpander expands resource under an AppService
type AppServiceResourceExpander struct {
	initialized  bool
	handledTypes []handledType
}

type handledType struct {
	name     string
	display  string
	endpoint endpoints.EndpointInfo
	children []handledType
}

// Name returns the name of the expander
func (e *AppServiceResourceExpander) Name() string {
	return "AppServiceResourceExpander"
}

func (e *AppServiceResourceExpander) ensureInitialized() {
	if !e.initialized {
		e.handledTypes = []handledType{
			{
				name:     "site",
				endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
				children: []handledType{
					{
						display:  "config",
						endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
						children: []handledType{
							{
								display:  "appsettings",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01", "POST"),
							},
							{
								display:  "authsettings",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01", "POST"),
							},
							{
								display:  "connectionstrings",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01", "POST"),
							},
							{
								display:  "logs",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs/list", "2018-02-01", "POST"),
							},
							{
								display:  "metadata",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01", "POST"),
							},
							{
								display:  "publishingcredentials",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01", "POST"),
							},
							{
								display:  "pushsettings",
								endpoint: getEndpointInfoFromURLWithVerbAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01", "POST"),
							},
							{
								display:  "slotConfigNames",
								endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
							},
							{
								display:  "virtualNetwork",
								endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/virtualNetwork", "2018-02-01"),
							},
							{
								display:  "web",
								endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
							},
						},
					},
					{
						display:  "siteextensions",
						endpoint: getEndpointInfoFromURLAndPanicOnError("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
					},
				},
			},
		}
		e.initialized = true
	}
}

func getEndpointInfoFromURLAndPanicOnError(url string, apiVersion string) endpoints.EndpointInfo {
	return getEndpointInfoFromURLWithVerbAndPanicOnError(url, apiVersion, "GET")
}
func getEndpointInfoFromURLWithVerbAndPanicOnError(url string, apiVersion string, verb string) endpoints.EndpointInfo {
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
		result := getHandledTypeForURL(url, handledType.children)
		if result != nil {
			return result
		}
	}
	return nil
}

// DoesExpand checks if this is an RG
func (e *AppServiceResourceExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
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
func (e *AppServiceResourceExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {

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

	var resourceResponse armclient.ResourceReseponse
	err = json.Unmarshal([]byte(data), &resourceResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, currentItem.ExpandURL)
		panic(err)
	}

	newItems := []TreeNode{}
	matchResult := handledType.endpoint.Match(currentItem.ExpandURL) // TODO - return the matches from getHandledTypeForURL to avoid re-calculating!
	templateValues := matchResult.Values
	for _, child := range handledType.children {

		url, err := child.endpoint.BuildURL(templateValues)
		if err != nil {
			err = fmt.Errorf("Error building URL: %s\nURL:%s", child.display, err)
			panic(err)
		}
		newItems = append(newItems, TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "appservice",
			Name:      child.display,
			Display:   child.display,
			ExpandURL: url,
			ItemType:  resourceType,
			DeleteURL: "NotSupported",
		})
	}

	return ExpanderResult{
		Nodes:    &newItems,
		Response: string(data),
		IsPrimaryResponse: true, // only returning items that we are the primary response for
	}
}
