package expanders

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

// ResponsePropertyMapping stores the mapping from response property name to Metadata item name
type ResponsePropertyMapping struct {
	// ResponsePropertyName is the name of the property used to look up a value in the response object
	ResponsePropertyName string
	// MetadataName is the name of the key to use to store the value in the item metadata
	MetadataName string
}

// DatabricksAPIResponseMetadata describes how to process responses from the Databricks API
type DatabricksAPIResponseMetadata struct {
	// ResponseArrayPath is the path to the array of items in the JSON response
	ResponseArrayPath string
	// ResponseArrayPath is name of the identified property for items in the JSON response
	ResponseIDPropertyName string
	// SubResourceQueryStringName is the name of the query string parameter to identify an item in sub resource requests ()
	SubResourceQueryStringName string
	// ResponsePropertyMappings holds details for properties to extra from the response object into item metadata
	ResponsePropertyMappings []ResponsePropertyMapping
	// SubResourceQueryStringValues is an array of query string parameters to populate from item metadata (this includes the SubResourceQueryStringName)
	SubResourceQueryStringValues []string
}

type DatabricksWorkspaceGetStatusResponse struct {
	Path       string `json:"path"`
	Language   string `json:"language"`
	ObjectType string `json:"object_type"`
	ObjectID   string `jsong:"object_id"`
}

// NewDatabricksAPIResponseMetadata creates a DatabricksAPIResponseMetadata instance
func NewDatabricksAPIResponseMetadata(responseArrayPath string, responseIDPropertyName string, subResourceQueryStringName string, subResourceAdditionalMetadata []string) DatabricksAPIResponseMetadata {
	metadataNames := append(subResourceAdditionalMetadata, subResourceQueryStringName)
	return DatabricksAPIResponseMetadata{
		ResponseArrayPath:            responseArrayPath,
		ResponseIDPropertyName:       responseIDPropertyName,
		SubResourceQueryStringName:   subResourceQueryStringName,
		ResponsePropertyMappings:     []ResponsePropertyMapping{{ResponsePropertyName: responseIDPropertyName, MetadataName: subResourceQueryStringName}},
		SubResourceQueryStringValues: metadataNames,
	}
}

var _ SwaggerAPISet = SwaggerAPISetDatabricks{}

// SwaggerAPISetDatabricks holds the config for working with an Azure Search Service
type SwaggerAPISetDatabricks struct {
	resourceTypes   []swagger.ResourceType
	httpClient      http.Client
	workspaceID     string // ARM resource ID for the search service (/subscriptions/....)
	nodeID          string
	workspaceURL    string
	managementToken string
	databricksToken string
}

// NewSwaggerAPISetDatabricks creates a new SwaggerAPISetDatabricks
func NewSwaggerAPISetDatabricks(resourceTypes []swagger.ResourceType, workspaceID string, nodeID string, workspaceURL string, managementToken string, databricksToken string) SwaggerAPISetDatabricks {
	c := SwaggerAPISetDatabricks{}
	c.resourceTypes = resourceTypes
	c.httpClient = http.Client{}
	c.workspaceID = workspaceID
	c.nodeID = nodeID
	c.workspaceURL = workspaceURL
	c.managementToken = managementToken
	c.databricksToken = databricksToken
	return c
}

// ID returns the ID for the APISet
func (c SwaggerAPISetDatabricks) ID() string {
	return c.nodeID
}

// MatchChildNodesByName indicates whether child nodes should be matched by name (or position)
func (c SwaggerAPISetDatabricks) MatchChildNodesByName() bool {
	return true
}

// AppliesToNode is called by the Swagger exapnder to test whether the node applies to this APISet
func (c SwaggerAPISetDatabricks) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerAPISetID set
	// this should never happen for search nodes
	return false
}

// GetResourceTypes returns the ResourceTypes for the API Set
func (c SwaggerAPISetDatabricks) GetResourceTypes() []swagger.ResourceType {
	return c.resourceTypes
}

// DoRequest makes a request against the search endpoint
func (c SwaggerAPISetDatabricks) DoRequest(verb string, url string) (string, error) {
	return c.DoRequestWithBody(verb, url, "")
}

// DoRequestWithBody makes a request against the search endpoint
func (c SwaggerAPISetDatabricks) DoRequestWithBody(verb string, url string, body string) (string, error) {
	return c.DoRequestWithBodyAndHeaders(verb, url, body, map[string]string{})
}

// DoRequestWithBodyAndHeaders makes a request against the search endpoint
func (c SwaggerAPISetDatabricks) DoRequestWithBodyAndHeaders(verb string, url string, body string, headers map[string]string) (string, error) {

	request, err := http.NewRequest(verb, url, bytes.NewReader([]byte(body)))
	if err != nil {
		err = fmt.Errorf("Failed to create request" + err.Error() + url)
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+c.databricksToken)
	request.Header.Set("X-Databricks-Azure-SP-Management-Token", c.managementToken)
	request.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", c.workspaceID)
	for name, value := range headers {
		request.Header.Set(name, value)
	}
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

func (c SwaggerAPISetDatabricks) unpackFileContents(data string, contentPath string) (string, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
		err = fmt.Errorf("Failed to json decode dbfs content: %v", err)
		return "", err
	}
	data = jsonData[contentPath].(string)
	buf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		err = fmt.Errorf("Failed to base64 decode dbfs content : %v", err)
		return "", err
	}
	data = string(buf)
	return data, nil
}

func (c SwaggerAPISetDatabricks) addWorkspaceContentNode(subResources *[]SubResource, currentItem *TreeNode, expandURL string, itemID string, metadata map[string]string) (string, error) {

	// Replace content with get_status response
	path := currentItem.Metadata["path"]
	url := "https://" + c.workspaceURL + "/api/2.0/workspace/get-status?path=" + path
	data, err := c.DoRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make get_status request: %v", err)
		return "", err
	}

	var getStatusResponse DatabricksWorkspaceGetStatusResponse
	if err = json.Unmarshal([]byte(data), &getStatusResponse); err != nil {
		err = fmt.Errorf("Failed to unmarshal get_status response: %v", err)
		return "", err
	}
	metadata["object_type"] = getStatusResponse.ObjectType
	metadata["language"] = getStatusResponse.Language

	// Add a child item to show the contents
	subResource := SubResource{
		ID: c.nodeID + expandURL + "/content",
		ResourceType: swagger.ResourceType{
			Children:       []swagger.ResourceType{},
			SubResources:   []swagger.ResourceType{},
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/workspace/export", ""),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/api/2.0/workspace/import", ""),
			DeleteEndpoint: currentItem.SwaggerResourceType.DeleteEndpoint,
		},
		ExpandURL: "/api/2.0/workspace/export?format=SOURCE&path=" + itemID,
		Name:      "Content: " + itemID,
		Metadata:  metadata,
	}
	// TODO - consider adding a PutEndpoint and handling this in Update
	// see https://docs.databricks.com/dev-tools/api/latest/examples.html#import-a-notebook-or-directory
	*subResources = append(*subResources, subResource)
	return data, nil
}
func (c SwaggerAPISetDatabricks) addDbfsContentNode(subResources *[]SubResource, currentItem *TreeNode, item map[string]interface{}, expandURL string, itemID string, metadata map[string]string) (string, error) {
	fileSize := item["file_size"].(float64)

	// Replace content with get_status response
	path := currentItem.Metadata["path"]
	url := "https://" + c.workspaceURL + "/api/2.0/dbfs/get-status?path=" + path
	data, err := c.DoRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make get_status request: %v", err)
		return "", err
	}
	if fileSize == 0 /* 0 => directory */ {
		return data, nil
	}
	var subResource SubResource
	if fileSize > 8192 /* pick a max size to handle */ {
		subResource = SubResource{
			ID: c.nodeID + expandURL + "/content",
			ResourceType: swagger.ResourceType{
				Children:     []swagger.ResourceType{},
				SubResources: []swagger.ResourceType{},
				FixedContent: "File too large to load",
			},
			Name: "Content: " + itemID,
		}
	} else {
		// Add a child item to show the contents
		subResource = SubResource{
			ID: c.nodeID + expandURL + "/content",
			ResourceType: swagger.ResourceType{
				Children:       []swagger.ResourceType{},
				SubResources:   []swagger.ResourceType{},
				Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/dbfs/read", ""),
				PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/api/2.0/dbfs/put", ""),
				DeleteEndpoint: currentItem.SwaggerResourceType.DeleteEndpoint,
			},
			ExpandURL: "/api/2.0/dbfs/read?offset=0&path=" + itemID,
			Name:      "Content: " + itemID,
			Metadata:  metadata,
		}
		// see https://docs.databricks.com/dev-tools/api/latest/dbfs.html#put
	}
	*subResources = append(*subResources, subResource)
	return data, nil
}

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetDatabricks) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	currentItemTemplateURL := currentItem.SwaggerResourceType.Endpoint.TemplateURL
	if currentItem.SwaggerResourceType.FixedContent != "" {
		return APISetExpandResponse{
			Response:      currentItem.SwaggerResourceType.FixedContent,
			ResponseType:  interfaces.ResponsePlainText,
			ChildMetadata: currentItem.Metadata,
		}, nil
	}

	// handle query string values for items that were added as Child nodes by the swagger expander
	expandQueryString := ""
	switch currentItem.ExpandURL {
	case "/api/2.0/secrets/list", "/api/2.0/secrets/acls/list":
		expandQueryString = "?scope=" + currentItem.Metadata["scope"]
	case "/api/2.0/jobs/runs/list":
		expandQueryString = "?limit=0" // TODO add paging. "limit=0" sets the maximum number allowed - see https://docs.databricks.com/dev-tools/api/latest/jobs.html#runs-list
		jobID := currentItem.Metadata["job_id"]
		if jobID != "" {
			expandQueryString = expandQueryString + "&job_id=" + jobID
		}
	}

	// Perform the request
	url := "https://" + c.workspaceURL + currentItem.ExpandURL + expandQueryString
	data, err := c.DoRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make request: %v", err)
		return APISetExpandResponse{}, err
	}

	// unpack content responses that are wrapped in JSON/Base64 encoded
	responseType := interfaces.ResponseJSON
	if currentItemTemplateURL == "/api/2.0/dbfs/read" {
		responseType = interfaces.ResponsePlainText
		if data, err = c.unpackFileContents(data, "data"); err != nil {
			return APISetExpandResponse{}, err
		}
	}
	if currentItemTemplateURL == "/api/2.0/workspace/export" {
		responseType = interfaces.ResponsePlainText
		if data, err = c.unpackFileContents(data, "content"); err != nil {
			return APISetExpandResponse{}, err
		}
	}

	// Process subresources
	subResources := []SubResource{}
	if len(resourceType.SubResources) > 0 ||
		currentItemTemplateURL == "/api/2.0/dbfs/list" ||
		currentItemTemplateURL == "/api/2.0/workspace/list" {

		var subResourceType swagger.ResourceType
		switch len(resourceType.SubResources) {
		case 0:
			subResourceType = resourceType // e.g. "/api/2.0/workspaces/list" where we reuse the current resource type as we drill down
		case 1:
			subResourceType = resourceType.SubResources[0]
		default:
			return APISetExpandResponse{}, fmt.Errorf("Only expecting a single SubResource type")
		}

		var jsonData map[string]interface{}
		if err = json.Unmarshal([]byte(data), &jsonData); err != nil {
			return APISetExpandResponse{}, err
		}

		expandParameters := c.getExpandParameters(currentItemTemplateURL)
		itemArrayTemp := jsonData[expandParameters.ResponseArrayPath]
		if itemArrayTemp != nil {
			itemArray := itemArrayTemp.([]interface{})

			for _, item := range itemArray {
				item := item.(map[string]interface{})

				metadata := map[string]string{}
				for k, v := range currentItem.Metadata {
					metadata[k] = v
				}

				// save metadata items from response
				for _, responseItemToStore := range expandParameters.ResponsePropertyMappings {
					valueTemp := item[responseItemToStore.ResponsePropertyName]
					value := fmt.Sprintf("%v", valueTemp)
					metadata[responseItemToStore.MetadataName] = value
				}

				// build query string from metadata
				queryString := ""
				queryStringSeparator := ""
				for _, queryStringName := range expandParameters.SubResourceQueryStringValues {
					queryStringValue := metadata[queryStringName]
					if queryStringValue != "" {
						queryString += fmt.Sprintf("%s%s=%s", queryStringSeparator, queryStringName, queryStringValue)
						queryStringSeparator = "&"
					}
				}
				expandURL := fmt.Sprintf("%s?%s", subResourceType.Endpoint.TemplateURL, queryString)

				itemID := metadata[expandParameters.SubResourceQueryStringName]
				if itemID == currentItem.Metadata[expandParameters.SubResourceQueryStringName] {
					// skip adding the item (e.g. workspace list returns existing item when on a file)
					// and determinte whether to add a "Contents" child node
					switch currentItemTemplateURL {
					case "/api/2.0/workspace/list":
						newData, err := c.addWorkspaceContentNode(&subResources, currentItem, expandURL, itemID, metadata)
						if err != nil {
							return APISetExpandResponse{}, err
						}
						if newData != "" {
							data = newData
						}
					case "/api/2.0/dbfs/list":
						newData, err := c.addDbfsContentNode(&subResources, currentItem, item, expandURL, itemID, metadata)
						if err != nil {
							return APISetExpandResponse{}, err
						}
						if newData != "" {
							data = newData
						}
					}
				} else {
					subResource := SubResource{
						ID:           c.nodeID + expandURL,
						ResourceType: subResourceType,
						ExpandURL:    expandURL,
						Name:         itemID,
						Metadata:     metadata,
					}
					if subResourceType.DeleteEndpoint != nil && subResourceType.DeleteEndpoint.TemplateURL != "" {
						subResource.DeleteURL = subResourceType.DeleteEndpoint.TemplateURL // delete urls are all fixed values
					}
					subResources = append(subResources, subResource)
				}
			}
		}
	}

	return APISetExpandResponse{
		Response:      data,
		ResponseType:  responseType,
		SubResources:  subResources,
		ChildMetadata: currentItem.Metadata, // propagate metadata (e.g. job_id) down the tree
	}, nil
}

// get returns arrayPath, idPropertyName, queryStringName and an array of additional query strings to pass
func (c SwaggerAPISetDatabricks) getExpandParameters(templateURL string) DatabricksAPIResponseMetadata {

	switch templateURL {
	case "/api/2.0/clusters/list":
		return NewDatabricksAPIResponseMetadata("clusters", "cluster_id", "cluster_id", []string{})
	case "/api/2.0/instance-pools/list":
		return NewDatabricksAPIResponseMetadata("instance_pools", "instance_pool_id", "instance_pool_id", []string{})
	case "/api/2.0/jobs/list":
		return NewDatabricksAPIResponseMetadata("jobs", "job_id", "job_id", []string{})
	case "/api/2.0/jobs/runs/list":
		return NewDatabricksAPIResponseMetadata("runs", "run_id", "run_id", []string{})
	case "/api/2.0/secrets/scopes/list":
		return NewDatabricksAPIResponseMetadata("scopes", "name", "scope", []string{})
	case "/api/2.0/secrets/list":
		return NewDatabricksAPIResponseMetadata("secrets", "key", "key", []string{"scope"})
	case "/api/2.0/secrets/acls/list":
		return NewDatabricksAPIResponseMetadata("items", "principal", "principal", []string{"scope"})
	case "/api/2.0/token/list":
		return NewDatabricksAPIResponseMetadata("token_infos", "token_id", "token_id", []string{})
	case "/api/2.0/dbfs/list":
		return NewDatabricksAPIResponseMetadata("files", "path", "path", []string{})
	case "/api/2.0/workspace/list":
		return NewDatabricksAPIResponseMetadata("objects", "path", "path", []string{})
	}
	return NewDatabricksAPIResponseMetadata("", "", "", []string{})
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetDatabricks) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	metadata := item.Metadata
	bodyValue := map[string]interface{}{}

	switch item.SwaggerResourceType.Endpoint.TemplateURL {
	case "/api/2.0/clusters/get":
		bodyValue["cluster_id"] = metadata["cluster_id"]
	case "/api/2.0/instance-pools/get":
		bodyValue["instance_pool_id"] = metadata["instance_pool_id"]
	case "/api/2.0/jobs/get":
		bodyValue["job_id"] = metadata["job_id"]
	case "/api/2.0/jobs/runs/get":
		bodyValue["run_id"] = metadata["run_id"]
	case "/api/2.0/secrets/{scope}":
		bodyValue["scope"] = metadata["scope"]
	case "/api/2.0/secrets/get":
		bodyValue["scope"] = metadata["scope"]
		bodyValue["key"] = metadata["key"]
	case "/api/2.0/secrets/acls/get":
		bodyValue["scope"] = metadata["scope"]
		bodyValue["principal"] = metadata["principal"]
	case "/api/2.0/dbfs/list":
		bodyValue["path"] = metadata["path"]
		bodyValue["recursive"] = true
	case "/api/2.0/workspace/list":
		bodyValue["path"] = metadata["path"]
		bodyValue["recursive"] = true
	}

	if len(bodyValue) > 0 {
		bodyBuf, err := json.Marshal(bodyValue)
		if err != nil {
			return false, fmt.Errorf("Failed to serialize DELETE body: %v", err)
		}
		body := string(bodyBuf)
		url := "https://" + c.workspaceURL + item.DeleteURL
		_, err = c.DoRequestWithBody("POST", url, body)
		if err != nil {
			return false, fmt.Errorf("Failed to serialize DELETE body: %v", err)
		}
		return true, nil
	}

	return false, fmt.Errorf("Not implemented")
}

// Update attempts to update the specified item with new content
func (c SwaggerAPISetDatabricks) Update(ctx context.Context, item *TreeNode, content string) error {

	// Assumptions:
	//  - All updates are POST operations
	//  - All updates use fixed URLS (i.e. the ID is in the body, not the URL)
	// Exceptions:
	//  - /api/2.0/dbfs/put - body needs wrapping
	body := content

	switch item.SwaggerResourceType.PutEndpoint.TemplateURL {
	case "/api/2.0/dbfs/put":
		path := item.Metadata["path"]
		base64Content := base64.StdEncoding.EncodeToString([]byte(content))
		body = fmt.Sprintf("{\"path\":\"%s\", \"overwrite\":true,\"contents\":\"%s\"}", path, base64Content)
	case "/api/2.0/workspace/import":
		path := item.Metadata["path"]
		language := item.Metadata["language"]
		base64Content := base64.StdEncoding.EncodeToString([]byte(content))
		body = fmt.Sprintf("{\"path\":\"%s\", \"overwrite\":true, \"language\":\"%s\",\"content\":\"%s\"}", path, language, base64Content)
	}

	url := "https://" + c.workspaceURL + item.SwaggerResourceType.PutEndpoint.TemplateURL
	_, err := c.DoRequestWithBody("POST", url, body)
	return err
}
