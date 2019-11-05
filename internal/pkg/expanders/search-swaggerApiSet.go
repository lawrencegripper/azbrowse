package expanders

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type searchListResponse struct {
	Value []searchListItem `json:"value"`
}
type searchListItem struct {
	Name string `json:"name"`
}

type searchIndexResponse struct {
	Fields []struct {
		Name string `json:"name"`
		Key  bool   `json:"key"`
	} `json:"fields"`
}

type searchIndexDocumentList struct {
	Value []map[string]interface{}
}

var _ SwaggerAPISet = SwaggerAPISetSearch{}

// SwaggerAPISetSearch holds the config for working with an Azure Search Service
type SwaggerAPISetSearch struct {
	resourceTypes  []swagger.ResourceType
	httpClient     http.Client
	searchID       string // ARM resource ID for the search service (/subscriptions/....)
	searchEndpoint string // https://<name>.search.windows.net/
	adminKey       string
}

// NewSwaggerAPISetSearch creates a new SwaggerAPISetSearch
func NewSwaggerAPISetSearch(resourceTypes []swagger.ResourceType, searchID string, searchEndpoint string, adminKey string) SwaggerAPISetSearch {
	c := SwaggerAPISetSearch{}
	c.resourceTypes = resourceTypes
	c.httpClient = http.Client{}
	c.searchID = searchID
	c.searchEndpoint = searchEndpoint
	c.adminKey = adminKey
	return c
}

// ID returns the ID for the APISet
func (c SwaggerAPISetSearch) ID() string {
	return c.searchID
}

// MatchChildNodesByName indicates whether child nodes should be matched by name (or position)
func (c SwaggerAPISetSearch) MatchChildNodesByName() bool {
	return true
}

// AppliesToNode is called by the Swagger exapnder to test whether the node applies to this APISet
func (c SwaggerAPISetSearch) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerAPISetID set
	// this should never happen for search nodes
	return false
}

// GetResourceTypes returns the ResourceTypes for the API Set
func (c SwaggerAPISetSearch) GetResourceTypes() []swagger.ResourceType {
	return c.resourceTypes
}

func (c SwaggerAPISetSearch) doRequest(verb string, url string) (string, error) {
	return c.doRequestWithBody(verb, url, "")
}
func (c SwaggerAPISetSearch) doRequestWithBody(verb string, url string, body string) (string, error) {
	request, err := http.NewRequest(verb, url, bytes.NewReader([]byte(body)))
	if err != nil {
		err = fmt.Errorf("Failed to create request" + err.Error() + url)
		return "", err
	}

	request.Header.Set("api-key", c.adminKey)
	response, err := c.httpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + url)
		return "", err
	}
	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Failed to read body: %s", err)
		return "", err
	}
	data := string(buf)
	if 200 <= response.StatusCode && response.StatusCode < 300 {
		return data, nil
	}
	return "", fmt.Errorf("Response failed with %s (%s): %s", response.Status, url, data)
}

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetSearch) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	subResources := []SubResource{}
	url := c.searchEndpoint + currentItem.ExpandURL
	data, err := c.doRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make request: %s", err)
		return APISetExpandResponse{}, err
	}

	currentItemTemplateURL := currentItem.SwaggerResourceType.Endpoint.TemplateURL

	indexKey := ""
	if currentItemTemplateURL == "/indexes('{indexName}')" {
		indexKey, err = c.getIndexKey(data)
		if err != nil {
			return APISetExpandResponse{Response: data}, err
		}
	} else {
		// propagate indexKey if set in metadata
		indexKey = currentItem.Metadata["IndexKey"]
	}

	// expand if we have subresources. Also don't expand Docs as they are user-defined format
	if len(resourceType.SubResources) > 0 {
		if len(resourceType.SubResources) > 1 {
			return APISetExpandResponse{}, fmt.Errorf("Only expecting a single SubResource type")
		}

		matchResult := resourceType.Endpoint.Match(currentItem.ExpandURL)
		templateValues := matchResult.Values
		subResourceType := resourceType.SubResources[0]

		subResourceEndpoint := subResourceType.Endpoint
		newURLSegment := subResourceEndpoint.URLSegments[len(subResourceEndpoint.URLSegments)-1]
		newTemplateName := newURLSegment.Name

		var extraIDs []string
		var err error
		if currentItemTemplateURL == "/indexes('{indexName}')/docs" {
			extraIDs, err = c.getKeys(data, indexKey)
		} else {
			extraIDs, err = c.getNames(data)
		}
		if err != nil {
			return APISetExpandResponse{Response: data}, err
		}

		for _, item := range extraIDs {

			templateValues[newTemplateName] = item
			subResourceURL, err := subResourceType.Endpoint.BuildURL(templateValues)
			if err != nil {
				return APISetExpandResponse{}, fmt.Errorf("Error building subresource URL: %s", err)
			}

			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				subResourceTemplateValues := subResourceType.Endpoint.Match(subResourceURL).Values
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return APISetExpandResponse{Response: data}, err
				}
			}
			subResource := SubResource{
				ID:           c.searchID + subResourceURL,
				Name:         item,
				ResourceType: subResourceType,
				ExpandURL:    subResourceURL,
				DeleteURL:    deleteURL,
				Metadata: map[string]string{
					"IndexKey": indexKey,
				},
			}
			subResources = append(subResources, subResource)
		}
	}

	return APISetExpandResponse{
		Response:     data,
		ResponseType: ResponseJSON,
		SubResources: subResources,
		ChildMetadata: map[string]string{
			"IndexKey": indexKey,
		},
	}, nil
}

func (c SwaggerAPISetSearch) getIndexKey(response string) (string, error) {

	var indexResponse searchIndexResponse
	err := json.Unmarshal([]byte(response), &indexResponse)
	if err != nil {
		err = fmt.Errorf("Error parsing index response: %s", err)
		return "", err
	}

	for _, field := range indexResponse.Fields {
		if field.Key {
			return field.Name, nil
		}
	}

	return "", fmt.Errorf("No key field found in index")
}

func (c SwaggerAPISetSearch) getNames(response string) ([]string, error) {

	var listResponse searchListResponse
	err := json.Unmarshal([]byte(response), &listResponse)
	if err != nil {
		err = fmt.Errorf("Error parsing response: %s", err)
		return []string{}, err
	}

	names := []string{}
	for _, item := range listResponse.Value {
		name := item.Name
		names = append(names, name)
	}

	return names, nil
}

func (c SwaggerAPISetSearch) getKeys(response string, keyName string) ([]string, error) {

	var listResponse searchIndexDocumentList
	err := json.Unmarshal([]byte(response), &listResponse)
	if err != nil {
		err = fmt.Errorf("Error parsing response: %s", err)
		return []string{}, err
	}

	keys := []string{}
	for _, item := range listResponse.Value {
		key := item[keyName].(string)
		keys = append(keys, key)
	}

	return keys, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetSearch) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	url := c.searchEndpoint + item.DeleteURL
	_, err := c.doRequest("DELETE", url)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s (%s)", err.Error(), item.DeleteURL)
		return false, err
	}
	return true, nil
}

// Update attempts to update the specified item with new content
func (c SwaggerAPISetSearch) Update(ctx context.Context, item *TreeNode, content string) error {
	return fmt.Errorf("Not Implemented")
}
