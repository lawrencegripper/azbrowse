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
	isPrimaryResponse := false
	response := "TODO"
	newItems := []*TreeNode{}

	if currentItem.Namespace == "containerRegistry" {
		isPrimaryResponse = true
		registryID := currentItem.Metadata["RegistryID"]
		// TODO - factor this out to a separate function!

		loginServer, err := e.GetLoginServer(ctx, registryID)
		if err != nil {
			return ExpanderResult{
				Err: err,
				SourceDescription: "ContainerRegistryExpander request",
			}
		}

		token, err := e.GetRegistryToken(loginServer)
		if err != nil {
			return ExpanderResult{
				Err: err,
				SourceDescription: "ContainerRegistryExpander request",
			}
		}

		responseBuf, err := e.DoRequest(fmt.Sprintf("https://%s/acr/v1/_catalog", loginServer), token)
		if err != nil {
			return ExpanderResult{
				Err: err,
				SourceDescription: "ContainerRegistryExpander request",
			}
		}
		response = string(responseBuf)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(responseBuf, &jsonResponse)
		if err != nil {
			err = fmt.Errorf("Error unmarshalling repositories response: %s, %s", err, response)
			return ExpanderResult{
				Err: err,
				SourceDescription: "ContainerRegistryExpander request",
			}
		}

		repositories := jsonResponse["repositories"].([]interface{})

		for _, repositoryTemp := range repositories {
			repository := repositoryTemp.(string)
			newItems = append(newItems, &TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "containerRegistry",
				Name:      repository,
				Display:   repository,
				ItemType:  "containerRegistry.repository",
				ExpandURL: ExpandURLNotSupported,
			})
		}
	}

	swaggerResourceType := currentItem.SwaggerResourceType
	if swaggerResourceType != nil && swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}" {

		newItems = append(newItems, &TreeNode{
			Parentid:  currentItem.ID,
			Namespace: "containerRegistry",
			Name:      "Repositories",
			Display:   "Repositories",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			// SwaggerResourceType: &resourceType,
			Metadata: map[string]string{
				"RegistryID": currentItem.ExpandURL, // save full URL to registry
				"SuppressSwaggerExpand": "true",
			},
		})

	}

	return ExpanderResult{
		Err:               nil,
		Response:          response,
		SourceDescription: "ContainerRegistryExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: isPrimaryResponse,
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

func (e *ContainerRegistryExpander) GetRegistryToken(loginServer string) (string, error) {
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
	body = fmt.Sprintf("grant_type=refresh_token&service=%s&scope=registry:catalog:*&refresh_token=%s", loginServer, refreshToken)
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
