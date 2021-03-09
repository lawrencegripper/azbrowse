package expanders

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type containerRegistryResponse struct {
	Properties struct {
		LoginServer string `json:"loginServer"`
	} `json:"properties"`
}

// NewContainerRegistryExpander creates a new instance of ContainerRegistryExpander
func NewContainerRegistryExpander(armclient *armclient.Client) *ContainerRegistryExpander {
	return &ContainerRegistryExpander{
		client:    &http.Client{},
		armClient: armclient,
	}
}

// Check interface
var _ Expander = &ContainerRegistryExpander{}

func (e *ContainerRegistryExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// ContainerRegistryExpander expands Tthe data-plane aspects of a Container Registry
type ContainerRegistryExpander struct {
	ExpanderBase
	client    *http.Client
	armClient *armclient.Client
}

// Name returns the name of the expander
func (e *ContainerRegistryExpander) Name() string {
	return "ContainerRegistryExpander"
}

// DoesExpand checks if this is a storage account
func (e *ContainerRegistryExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "containerRegistry" {
		return true, nil
	}
	return false, nil
}

// Expand returns ManagementPolicies in the StorageAccount
func (e *ContainerRegistryExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "ContainerRegistry" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}" {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/<repositories>",
			Namespace:             "containerRegistry",
			Name:                  "Repositories",
			Display:               "Repositories",
			ItemType:              SubResourceType,
			ExpandURL:             ExpandURLNotSupported,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"RegistryID": currentItem.ExpandURL, // save full URL to registry
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "ContainerRegistryExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "containerRegistry" && currentItem.ItemType == SubResourceType {
		return e.expandRepositories(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository" {
		return e.expandRepository(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.tags" {
		return e.expandRepositoryTags(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.tag" {
		return e.expandRepositoryTag(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.manifests" {
		return e.expandRepositoryManifests(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.manifest" {
		return e.expandRepositoryManifest(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "ContainerRegistryExpander request",
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *ContainerRegistryExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {
	switch currentItem.ItemType {
	case "containerRegistry.repository":
		return e.deleteRepository(ctx, currentItem)
	case "containerRegistry.repository.tag":
		return e.deleteRepositoryTag(ctx, currentItem)
	case "containerRegistry.repository.manifest":
		return e.deleteRepositoryManifest(ctx, currentItem)
	}
	return false, nil
}

func (e *ContainerRegistryExpander) expandRepositories(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	registryID := currentItem.Metadata["RegistryID"]

	loginServer, err := e.getLoginServer(ctx, registryID)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}
	currentItem.Metadata["loginServer"] = loginServer // expandNode requires this is in the Metadata!

	return e.expandNode(
		ctx,
		currentItem,
		fmt.Sprintf("https://%s/v2/_catalog", loginServer),
		"registry:catalog:*",
		"repositories",
		"",
		func(currentItem *TreeNode, item string) *TreeNode {
			return &TreeNode{
				Parentid:  currentItem.ID,
				ID:        currentItem.ID + "/" + item,
				Namespace: "containerRegistry",
				Name:      item,
				Display:   item,
				ItemType:  "containerRegistry.repository",
				ExpandURL: ExpandURLNotSupported,
				DeleteURL: currentItem.ID + "/" + item,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  item,
				},
			}
		},
		func(currentItem *TreeNode, lastItem string) *TreeNode {
			return &TreeNode{
				Parentid:              currentItem.ID,
				ID:                    currentItem.ID + "/<more>",
				Namespace:             "containerRegistry",
				Name:                  "more...",
				Display:               "more...",
				ItemType:              SubResourceType,
				ExpandURL:             ExpandURLNotSupported,
				SuppressSwaggerExpand: true,
				SuppressGenericExpand: true,
				ExpandInPlace:         true,
				Metadata: map[string]string{
					"RegistryID": registryID,
					"lastItem":   lastItem,
				},
			}
		})
}

func (e *ContainerRegistryExpander) expandRepository(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	accessToken, err := e.getRegistryToken(ctx, loginServer, fmt.Sprintf("repository:%s:pull", repository))
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	responseBuf, err := e.doRequest(ctx, "GET", fmt.Sprintf("https://%s/acr/v1/%s", loginServer, repository), accessToken)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}
	response := string(responseBuf)

	newItems := []*TreeNode{
		e.getCreateTagsNodeFunc(loginServer, repository, "Tags")(currentItem, ""),
		e.getCreateManifestsNodeFunc(loginServer, repository, "Manifests")(currentItem, ""),
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: response, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}
func (e *ContainerRegistryExpander) deleteRepository(ctx context.Context, currentItem *TreeNode) (bool, error) {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	accessToken, err := e.getRegistryToken(ctx, loginServer, fmt.Sprintf("repository:%s:delete", repository))
	if err != nil {
		return false, err
	}

	_, err = e.doRequest(ctx, "DELETE", fmt.Sprintf("https://%s/acr/v1/%s", loginServer, repository), accessToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *ContainerRegistryExpander) expandRepositoryTags(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	return e.expandNode(
		ctx,
		currentItem,
		fmt.Sprintf("https://%s/acr/v1/%s/_tags", loginServer, repository),
		fmt.Sprintf("repository:%s:pull", repository),
		"tags",
		"name",
		e.getCreateTagNodeFunc(loginServer, repository),
		e.getCreateTagsNodeFunc(loginServer, repository, "more..."))
}

func (e *ContainerRegistryExpander) expandRepositoryTag(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	tag := currentItem.Metadata["tag"]

	accessToken, err := e.getRegistryToken(ctx, loginServer, fmt.Sprintf("repository:%s:metadata_read", repository))
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	responseBuf, err := e.doRequest(ctx, "GET", fmt.Sprintf("https://%s/acr/v1/%s/_tags/%s", loginServer, repository, tag), accessToken)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}
	response := string(responseBuf)

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(responseBuf, &jsonResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling repositories response: %s, %s", err, response)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	tagElement := jsonResponse["tag"].(map[string]interface{})
	digest := tagElement["digest"].(string)
	newItems := []*TreeNode{e.getCreateManifestNodeFunc(loginServer, repository)(currentItem, digest)}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: response, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}
func (e *ContainerRegistryExpander) deleteRepositoryTag(ctx context.Context, currentItem *TreeNode) (bool, error) {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	tag := currentItem.Metadata["tag"]

	accessToken, err := e.getRegistryToken(ctx, loginServer, fmt.Sprintf("repository:%s:delete", repository))
	if err != nil {
		return false, err
	}

	_, err = e.doRequest(ctx, "DELETE", fmt.Sprintf("https://%s/acr/v1/%s/_tags/%s", loginServer, repository, tag), accessToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *ContainerRegistryExpander) expandRepositoryManifests(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	return e.expandNode(
		ctx,
		currentItem,
		fmt.Sprintf("https://%s/acr/v1/%s/_manifests", loginServer, repository),
		fmt.Sprintf("repository:%s:pull", repository),
		"manifests",
		"digest",
		e.getCreateManifestNodeFunc(loginServer, repository),
		e.getCreateManifestsNodeFunc(loginServer, repository, "more..."))
}

func (e *ContainerRegistryExpander) expandRepositoryManifest(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	digest := currentItem.Metadata["digest"]

	return e.expandNode(
		ctx,
		currentItem,
		fmt.Sprintf("https://%s/v2/%s/manifests/%s", loginServer, repository, digest),
		fmt.Sprintf("repository:%s:metadata_read", repository),
		"manifest.tags",
		"",
		e.getCreateTagNodeFunc(loginServer, repository),
		nil)
}
func (e *ContainerRegistryExpander) deleteRepositoryManifest(ctx context.Context, currentItem *TreeNode) (bool, error) {

	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	digest := currentItem.Metadata["digest"]

	accessToken, err := e.getRegistryToken(ctx, loginServer, fmt.Sprintf("repository:%s:delete", repository))
	if err != nil {
		return false, err
	}

	_, err = e.doRequest(ctx, "DELETE", fmt.Sprintf("https://%s/v2/%s/manifests/%s", loginServer, repository, digest), accessToken)
	if err != nil {
		return false, err
	}
	return true, nil

}

type createItemNode func(currentItem *TreeNode, item string) *TreeNode

func (e *ContainerRegistryExpander) getCreateManifestNodeFunc(loginServer string, repository string) createItemNode {
	return func(currentItem *TreeNode, item string) *TreeNode {
		return &TreeNode{
			Parentid:  currentItem.ID,
			ID:        currentItem.ID + "/" + item,
			Namespace: "containerRegistry",
			Name:      item,
			Display:   item,
			ItemType:  "containerRegistry.repository.manifest",
			ExpandURL: ExpandURLNotSupported,
			DeleteURL: currentItem.ID + "/" + item,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"repository":  repository,
				"digest":      item,
			},
		}
	}
}

func (e *ContainerRegistryExpander) getCreateTagNodeFunc(loginServer string, repository string) createItemNode {
	return func(currentItem *TreeNode, item string) *TreeNode {
		return &TreeNode{
			Parentid:  currentItem.ID,
			ID:        currentItem.ID + "/" + item,
			Namespace: "containerRegistry",
			Name:      item,
			Display:   item,
			ItemType:  "containerRegistry.repository.tag",
			ExpandURL: ExpandURLNotSupported,
			DeleteURL: currentItem.ID + "/" + item,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"repository":  repository,
				"tag":         item,
			},
		}
	}
}

func (e *ContainerRegistryExpander) getCreateManifestsNodeFunc(loginServer string, repository string, title string) createItemNode {
	return func(currentItem *TreeNode, lastItem string) *TreeNode {
		return &TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			ID:        currentItem.ID + "/" + title,
			Name:      title,
			Display:   title,
			ItemType:  "containerRegistry.repository.manifests",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"repository":  repository,
				"lastItem":    lastItem,
			},
		}
	}
}

func (e *ContainerRegistryExpander) getCreateTagsNodeFunc(loginServer string, repository string, title string) createItemNode {
	return func(currentItem *TreeNode, lastItem string) *TreeNode {
		return &TreeNode{
			Parentid:  currentItem.ID,
			ID:        currentItem.ID + "/" + title,
			Namespace: "containerRegistry",
			Name:      title,
			Display:   title,
			ItemType:  "containerRegistry.repository.tags",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"repository":  repository,
				"lastItem":    lastItem,
			},
		}
	}
}

func (e *ContainerRegistryExpander) expandNode(
	ctx context.Context,
	currentItem *TreeNode, // requires loginServer and repository Metadata
	url string,
	accessTokenScope string,
	collectionPath string,
	itemPath string,
	createItemNodeFunc createItemNode,
	createContinuationFunc createItemNode) ExpanderResult {

	span, ctx := tracing.StartSpanFromContext(ctx, "expand(containerregistry):"+currentItem.ItemType+":"+currentItem.Name+":"+currentItem.ID, tracing.SetTag("item", currentItem))
	defer span.Finish()

	// TODO - add context around errors

	loginServer := currentItem.Metadata["loginServer"]
	lastItem := currentItem.Metadata["lastItem"]

	// get token
	accessToken, err := e.getRegistryToken(ctx, loginServer, accessTokenScope)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	// query
	continuation := ""
	if lastItem != "" && createContinuationFunc != nil {
		continuation = fmt.Sprintf("?last=%s", lastItem)
	}
	urlTemp := fmt.Sprintf("%s%s", url, continuation)

	response, items, err := e.getItemsForURL(ctx, urlTemp, accessToken, collectionPath, itemPath)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	newItems := []*TreeNode{}
	for _, item := range items {
		newItems = append(newItems, createItemNodeFunc(currentItem, item))
	}

	if len(newItems) > 0 {
		if createContinuationFunc != nil {
			newLastItem := items[len(items)-1]
			continuation = fmt.Sprintf("?last=%s", newLastItem)
			urlTemp = fmt.Sprintf("%s%s", url, continuation)
			_, nextItems, _ := e.getItemsForURL(ctx, urlTemp, accessToken, collectionPath, itemPath)
			if len(nextItems) > 0 {
				newItems = append(newItems, createContinuationFunc(currentItem, newLastItem))
			}
		}
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: response, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}
func (e *ContainerRegistryExpander) getItemsForURL(ctx context.Context, url string, accessToken string, collectionPath string, itemPath string) (string, []string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "getItemsForURL(containerregistry):"+url, tracing.SetTag("url", url))
	defer span.Finish()

	responseBuf, err := e.doRequest(ctx, "GET", url, accessToken)
	if err != nil {
		return "", []string{}, err
	}

	// project nodes
	response := string(responseBuf)

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(responseBuf, &jsonResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling repositories response: %s, %s", err, response)
		return "", []string{}, err
	}

	itemsResult := []string{}
	var itemsTemp interface{} //nolint:gosimple
	itemsTemp = jsonResponse
	pathSegments := strings.Split(collectionPath, ".")
	for _, pathSegment := range pathSegments {
		itemsTemp = itemsTemp.(map[string]interface{})[pathSegment]
		if itemsTemp == nil {
			break
		}
	}
	if itemsTemp != nil {
		items := itemsTemp.([]interface{})
		for _, itemTemp := range items {
			var itemName string
			if itemPath == "" {
				itemName = itemTemp.(string)
			} else {
				item := itemTemp.(map[string]interface{})
				itemName = item[itemPath].(string)
			}
			itemsResult = append(itemsResult, itemName)
		}
	}

	return response, itemsResult, nil
}

func (e *ContainerRegistryExpander) doRequest(ctx context.Context, verb string, url string, accessToken string) ([]byte, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "doRequest(containerregistry):"+url, tracing.SetTag("url", url))
	defer span.Finish()

	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to create request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, fmt.Errorf("Request failed: %s", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return []byte{}, fmt.Errorf("DoRequest failed %v for '%s'", response.StatusCode, url)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read body: %s", err)
	}

	return buf, nil
}

func (e *ContainerRegistryExpander) getLoginServer(ctx context.Context, registryID string) (string, error) {
	data, err := e.armClient.DoRequest(ctx, "GET", registryID)
	if err != nil {
		return "", fmt.Errorf("Failed to get registry: " + err.Error() + registryID)
	}
	var response containerRegistryResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, registryID)
		return "", err
	}

	// TODO also capture SKU to ensure it is a managed SKU
	loginServer := response.Properties.LoginServer
	return loginServer, nil
}

func (e *ContainerRegistryExpander) getRegistryToken(ctx context.Context, loginServer string, scope string) (string, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "getRegistryToken(containerregistry):"+loginServer+":"+scope, tracing.SetTag("loginServer", loginServer))
	defer span.Finish()
	// This logic is based around https://github.com/Azure/azure-cli/blob/c83710e4176cf598fccd57180263a4c5b0fc561e/src/azure-cli/azure/cli/command_modules/acr/_docker_utils.py#L110

	// TODO - add support for admin credentials if enabled and the AAD approach fails

	// Verify the loginServer/v2 endpoint returns a 401 with WWW-Authenticate header on a raw GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/v2", loginServer), bytes.NewReader([]byte("")))
	if err != nil {
		return "", fmt.Errorf("Failed to create request to validate loginserver/v2: %s", err)
	}
	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("Error validating loginserver/v2 endpoint: %s", err)
	}
	if response.StatusCode != 401 || response.Header.Get("WWW-Authenticate") == "" {
		return "", fmt.Errorf("Expected a 401 with WWW-Authenticate from loginserver/v2")
	}
	// TODO - currently using the loginServer to generate the realm etc but should consider pulling from the WWW-Authenticate header value

	// Make an accesstoken request
	tenantID := e.armClient.GetTenantID()
	armCLIToken, err := e.armClient.GetToken()
	if err != nil {
		return "", fmt.Errorf("Failed to get CLI token: %s", err)
	}
	body := fmt.Sprintf("grant_type=access_token&service=%s&tenant=%s&access_token=%s", loginServer, tenantID, armCLIToken.AccessToken)
	req, err = http.NewRequest("POST", fmt.Sprintf("https://%s/oauth2/exchange", loginServer), bytes.NewReader([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("Failed to create accesstoken request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err = e.client.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("Error making accesstoken request: %s", err)
	}
	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read body: %s", err)
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("Accesstoken request failed: %v: %s", response.StatusCode, string(buf))
	}
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(buf, &jsonResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling accesstoken response: %s, %s", err, buf)
		return "", err
	}

	refreshToken := jsonResponse["refresh_token"].(string)

	// Make a refreshtoken request
	body = fmt.Sprintf("grant_type=refresh_token&service=%s&scope=%s&refresh_token=%s", loginServer, scope, refreshToken)
	req, err = http.NewRequest("POST", fmt.Sprintf("https://%s/oauth2/token", loginServer), bytes.NewReader([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("Failed to create refreshtoken request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err = e.client.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("Error making refreshtoken request: %s", err)
	}
	defer response.Body.Close() //nolint: errcheck
	buf, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read body: %s", err)
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("Refreshtoken request failed: %v: %s", response.StatusCode, string(buf))
	}
	err = json.Unmarshal(buf, &jsonResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling refreshtoken response: %s, %s", err, buf)
		return "", err
	}

	accessToken := jsonResponse["access_token"].(string)

	return accessToken, nil
}
