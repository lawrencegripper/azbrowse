package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type ContainerRegistryResponse struct {
	Properties struct {
		LoginServer string `json:"loginServer"`
	} `json:"properties`
}

func NewContainerRegistryExpander() *ContainerRegistryExpander {
	return &ContainerRegistryExpander{
		client: &http.Client{},
	}
}

// ContainerRegistryExpander expands Tthe data-plane aspects of a Container Registry
type ContainerRegistryExpander struct {
	client *http.Client
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
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "Repositories",
			Display:   "Repositories",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"RegistryID":            currentItem.ExpandURL, // save full URL to registry
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          "TODO",
			SourceDescription: "ContainerRegistryExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "containerRegistry" && currentItem.ItemType == SubResourceType {
		return e.ExpandRepositories(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository" {
		return e.ExpandRepository(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.tags" {
		return e.ExpandRepositoryTags(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.tag" {
		return e.ExpandRepositoryTag(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.manifests" {
		return e.ExpandRepositoryManifests(ctx, currentItem)
	} else if currentItem.ItemType == "containerRegistry.repository.manifest" {
		return e.ExpandRepositoryManifest(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          "Error!",
		SourceDescription: "ContainerRegistryExpander request",
	}
}

func (e *ContainerRegistryExpander) GetLoginServer(ctx context.Context, registryID string) (string, error) {
	data, err := armclient.DoRequest(ctx, "GET", registryID)
	if err != nil {
		return "", fmt.Errorf("Failed to get registry: " + err.Error() + registryID)
	}
	var containerRegistryResponse ContainerRegistryResponse
	err = json.Unmarshal([]byte(data), &containerRegistryResponse)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, registryID)
		return "", err
	}

	// TODO also capture SKU to ensure it is a managed SKU
	loginServer := containerRegistryResponse.Properties.LoginServer
	return loginServer, nil
}
func (e *ContainerRegistryExpander) ExpandRepositories(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	registryID := currentItem.Metadata["RegistryID"]
	lastRepository := currentItem.Metadata["lastRepository"]

	loginServer, err := e.GetLoginServer(ctx, registryID)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	token, err := e.GetRegistryToken(loginServer, "registry:catalog:*")
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	continuation := ""
	if lastRepository != "" {
		continuation = fmt.Sprintf("&last=%s", lastRepository)
	}
	responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/v2/_catalog?n=3%s", loginServer, continuation), token)
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

	newItems := []*TreeNode{}
	repositoriesTemp := jsonResponse["repositories"]
	if repositoriesTemp != nil {
		repositories := repositoriesTemp.([]interface{})
		lastRepository = ""

		for _, repositoryTemp := range repositories {
			repository := repositoryTemp.(string)
			lastRepository = repository
			newItems = append(newItems, &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      repository,
				Display:   repository,
				ItemType:  "containerRegistry.repository",
				ExpandURL: ExpandURLNotSupported,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  repository,
				},
			})
		}
		newItems = append(newItems, &TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "more...",
			Display:   "more...",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"RegistryID":            registryID,
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
				"lastRepository":        lastRepository,
			},
		})
	}
	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *ContainerRegistryExpander) ExpandRepository(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	accessToken, err := e.GetRegistryToken(loginServer, fmt.Sprintf("repository:%s:pull", repository))
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/acr/v1/%s", loginServer, repository), accessToken)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}
	response := string(responseBuf)

	newItems := []*TreeNode{
		&TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "Tags",
			Display:   "Tags",
			ItemType:  "containerRegistry.repository.tags",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"accessToken": accessToken,
				"repository":  repository,
			},
		},
		&TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "Manifests",
			Display:   "Manifests",
			ItemType:  "containerRegistry.repository.manifests",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"accessToken": accessToken,
				"repository":  repository,
			},
		},
	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

type createItemNode func(currentItem *TreeNode, item string) *TreeNode

// TODO - mark functions as private!

func (e *ContainerRegistryExpander) ExpandNode(
	ctx context.Context,
	currentItem *TreeNode, // requires loginServer and repository Metadata
	url string,
	accessTokenScope string,
	collectionPath string,
	itemPath string,
	createItemNodeFunc createItemNode,
	createContinuationFunc createItemNode) ExpanderResult {

	// TODO - add context around errors

	loginServer := currentItem.Metadata["loginServer"]
	lastItem := currentItem.Metadata["lastItem"]

	// get token
	accessToken, err := e.GetRegistryToken(loginServer, accessTokenScope)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	// query
	continuation := ""
	if lastItem != "" {
		continuation = fmt.Sprintf("?last=%s", lastItem)
	}
	responseBuf, err := e.DoRequest(fmt.Sprintf("%s%s", url, continuation), accessToken)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	// project nodes
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

	newItems := []*TreeNode{}
	itemsTemp := jsonResponse[collectionPath]
	if itemsTemp != nil {
		items := itemsTemp.([]interface{})
		lastItem := ""
		for _, itemTemp := range items {
			item := itemTemp.(map[string]interface{})
			itemName := item[itemPath].(string)
			lastItem = itemName

			newItems = append(newItems, createItemNodeFunc(currentItem, itemName))
		}
		newItems = append(newItems, createContinuationFunc(currentItem, lastItem))
	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}

}

func (e *ContainerRegistryExpander) ExpandRepositoryTags(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]

	return e.ExpandNode(
		ctx,
		currentItem,
		fmt.Sprintf("https://%s/acr/v1/%s/_tags", loginServer, repository),
		fmt.Sprintf("repository:%s:pull", repository),
		"tags",
		"name",
		func(currentItem *TreeNode, item string) *TreeNode {
			return &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      item,
				Display:   item,
				ItemType:  "containerRegistry.repository.tag",
				ExpandURL: ExpandURLNotSupported,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  repository,
					"tag":         item,
				},
			}
		},
		func(currentItem *TreeNode, lastItem string) *TreeNode {
			return &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      "more...",
				Display:   "more...",
				ItemType:  "containerRegistry.repository.tags",
				ExpandURL: ExpandURLNotSupported,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  repository,
					"lastItem":     lastItem,
				},
			}
		})
}

func (e *ContainerRegistryExpander) ExpandRepositoryTag(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	tag := currentItem.Metadata["tag"]

	accessToken, err := e.GetRegistryToken(loginServer, fmt.Sprintf("repository:%s:metadata_read", repository))
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	// TODO - need to handle continuation calls for long lists
	responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/acr/v1/%s/_tags/%s", loginServer, repository, tag), accessToken)
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
	newItems := []*TreeNode{
		&TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      digest,
			Display:   digest,
			ItemType:  "containerRegistry.repository.manifest",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer": loginServer,
				"repository":  repository,
				"digest":      digest,
			},
		},
	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *ContainerRegistryExpander) ExpandRepositoryManifests(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	accessToken := currentItem.Metadata["accessToken"]
	lastManifest := currentItem.Metadata["lastManifest"]

	continuation := ""
	if lastManifest != "" {
		continuation = fmt.Sprintf("?last=%s", lastManifest)
	}
	responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/acr/v1/%s/_manifests%s", loginServer, repository, continuation), accessToken)
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

	newItems := []*TreeNode{}
	manifestsTemp := jsonResponse["manifests"]
	if manifestsTemp != nil {
		manifests := manifestsTemp.([]interface{})
		lastManifest := ""
		for _, manifestTemp := range manifests {
			manifest := manifestTemp.(map[string]interface{})
			digest := manifest["digest"].(string)
			lastManifest = digest
			newItems = append(newItems, &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      digest,
				Display:   digest,
				ItemType:  "containerRegistry.repository.manifest",
				ExpandURL: ExpandURLNotSupported,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  repository,
					"digest":      digest,
				},
			})
		}
		newItems = append(newItems, &TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "more...",
			Display:   "more...",
			ItemType:  "containerRegistry.repository.manifests",
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"loginServer":  loginServer,
				"accessToken":  accessToken,
				"repository":   repository,
				"lastManifest": lastManifest,
			},
		})
	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *ContainerRegistryExpander) ExpandRepositoryManifest(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	loginServer := currentItem.Metadata["loginServer"]
	repository := currentItem.Metadata["repository"]
	digest := currentItem.Metadata["digest"]

	accessToken, err := e.GetRegistryToken(loginServer, fmt.Sprintf("repository:%s:metadata_read", repository))
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "ContainerRegistryExpander request",
		}
	}

	// TODO - need to handle continuation calls for long lists
	responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/acr/v1/%s/_manifests/%s", loginServer, repository, digest), accessToken)
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

	newItems := []*TreeNode{}
	manifestElement := jsonResponse["manifest"].(map[string]interface{})
	tagsTemp := manifestElement["tags"]
	if tagsTemp != nil {
		tags := tagsTemp.([]interface{})

		for _, tagTemp := range tags {
			tag := tagTemp.(string)
			newItems = append(newItems, &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      tag,
				Display:   tag,
				ItemType:  "containerRegistry.repository.tag",
				ExpandURL: ExpandURLNotSupported,
				Metadata: map[string]string{
					"loginServer": loginServer,
					"repository":  repository,
					"tag":         tag,
				},
			})
		}
	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *ContainerRegistryExpander) DoRequest(url string, accessToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to create request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	response, err := e.client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Request failed: %s", err)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read body: %s", err)
	}

	return buf, nil
}

func (e *ContainerRegistryExpander) GetRegistryToken(loginServer string, scope string) (string, error) {
	// This logic is based around https://github.com/Azure/azure-cli/blob/c83710e4176cf598fccd57180263a4c5b0fc561e/src/azure-cli/azure/cli/command_modules/acr/_docker_utils.py#L110

	// TODO - add support for admin credentials if enabled and the AAD approach fails

	// Verify the loginServer/v2 endpoint returns a 401 with WWW-Authenticate header on a raw GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/v2", loginServer), bytes.NewReader([]byte("")))
	if err != nil {
		return "", fmt.Errorf("Failed to create request to validate loginserver/v2: %s", err)
	}
	response, err := e.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error validating loginserver/v2 endpoint: %s", err)
	}
	if response.StatusCode != 401 || response.Header.Get("WWW-Authenticate") == "" {
		return "", fmt.Errorf("Expected a 401 with WWW-Authenticate from loginserver/v2")
	}
	// TODO - currently using the loginServer to generate the realm etc but should consider pulling from the WWW-Authenticate header value

	// Make an accesstoken request
	tenantID := armclient.GetTenantID()
	armCLIToken, err := armclient.GetToken()
	if err != nil {
		return "", fmt.Errorf("Failed to get CLI token: %s", err)
	}
	body := fmt.Sprintf("grant_type=access_token&service=%s&tenant=%s&access_token=%s", loginServer, tenantID, armCLIToken.AccessToken)
	req, err = http.NewRequest("POST", fmt.Sprintf("https://%s/oauth2/exchange", loginServer), bytes.NewReader([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("Failed to create accesstoken request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err = e.client.Do(req)
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

	response, err = e.client.Do(req)
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
