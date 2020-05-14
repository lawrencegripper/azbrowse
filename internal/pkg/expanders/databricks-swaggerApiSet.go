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

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetDatabricks) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	if currentItem.SwaggerResourceType.Endpoint.TemplateURL == "/api/2.0/secrets/{scope}" {
		// fake node added to tree structure
		// no resources to add, but pass child metadata
		return APISetExpandResponse{
			Response:      "Choose a node to expand...",
			ResponseType:  ResponsePlainText,
			ChildMetadata: currentItem.Metadata,
		}, nil
	}

	if currentItem.ExpandURL == "/api/2.0/secrets/list" || currentItem.ExpandURL == "/api/2.0/secrets/acls/list" {
		// handle query string values for items that were added as Child nodes by the swagger expander

		// TODO Update DeleteURL  when handling deletion!
		currentItem.ExpandURL = currentItem.ExpandURL + "?scope=" + currentItem.Metadata["scope"]
	}

	subResources := []SubResource{}
	url := "https://" + c.workspaceURL + currentItem.ExpandURL
	if currentItem.SwaggerResourceType.Endpoint.TemplateURL == "/api/2.0/jobs/runs/list" {
		url = url + "?limit=0" // TODO add paging. "limit=0" sets the maximum number allowed - see https://docs.databricks.com/dev-tools/api/latest/jobs.html#runs-list
		jobID := currentItem.Metadata["job_id"]
		if jobID != "" {
			url = url + "&job_id=" + jobID
		}
	}

	data, err := c.DoRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make request: %s", err)
		return APISetExpandResponse{}, err
	}
	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the response and add these to newItems
		// TODO!

		if len(resourceType.SubResources) > 1 {
			return APISetExpandResponse{}, fmt.Errorf("Only expecting a single SubResource type")
		}
		subResourceType := resourceType.SubResources[0]

		arrayPath, idPropertyName, queryStringName, additionalQueryStrings := c.getExpandParameters(currentItem.SwaggerResourceType.Endpoint.TemplateURL)

		var jsonData map[string]interface{}
		if err = json.Unmarshal([]byte(data), &jsonData); err != nil {
			return APISetExpandResponse{}, err
		}
		itemArrayTemp := jsonData[arrayPath]
		if itemArrayTemp != nil {
			itemArray := itemArrayTemp.([]interface{})

			for _, item := range itemArray {
				item := item.(map[string]interface{})
				itemIDTemp := item[idPropertyName]
				itemID := fmt.Sprintf("%v", itemIDTemp)
				expandURL := fmt.Sprintf("%s?%s=%s", subResourceType.Endpoint.TemplateURL, queryStringName, itemID)

				for _, queryStringName := range additionalQueryStrings {
					queryStringValue := currentItem.Metadata[queryStringName]
					expandURL += fmt.Sprintf("&%s=%s", queryStringName, queryStringValue)
				}

				metadata := map[string]string{}
				for k, v := range currentItem.Metadata {
					metadata[k] = v
				}
				metadata[queryStringName] = itemID

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

	return APISetExpandResponse{
		Response:      data,
		ResponseType:  ResponseJSON,
		SubResources:  subResources,
		ChildMetadata: currentItem.Metadata, // propagate metadata (e.g. job_id) down the tree
	}, nil
}

// get returns arrayPath, idPropertyName, queryStringName and an array of additional query strings to pass
func (c SwaggerAPISetDatabricks) getExpandParameters(templateURL string) (string, string, string, []string) {
	switch templateURL {
	case "/api/2.0/clusters/list":
		return "clusters", "cluster_id", "cluster_id", []string{}
	case "/api/2.0/instance-pools/list":
		return "instance_pools", "instance_pool_id", "instance_pool_id", []string{}
	case "/api/2.0/jobs/list":
		return "jobs", "job_id", "job_id", []string{}
	case "/api/2.0/jobs/runs/list":
		return "runs", "run_id", "run_id", []string{}
	case "/api/2.0/secrets/scopes/list":
		return "scopes", "name", "scope", []string{}
	case "/api/2.0/secrets/list":
		return "secrets", "key", "key", []string{}
	case "/api/2.0/secrets/acls/list":
		return "items", "principal", "principal", []string{"scope"}
	case "/api/2.0/token/list":
		return "token_infos", "token_id", "token_id", []string{"scope"}
	}
	return "", "", "", []string{}
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
	url := "https://" + c.workspaceURL + item.SwaggerResourceType.PutEndpoint.TemplateURL
	_, err := c.DoRequestWithBody("POST", url, content)
	return err
}
