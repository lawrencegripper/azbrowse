package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type clusterCredentialsResponse struct {
	KubeConfigs []struct {
		Name  string `json:"name"`
		Value string `json:"Value"`
	} `json:"kubeconfigs"`
}

// kubeConfig is a minimal struct for parsing the parts of the response that we care about
type kubeConfig struct {
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			Token string `yaml:"token"`
		} `yaml:"user"`
	} `yaml:"users"`
}

// AzureKubernetesServiceExpander expands the kubernetes aspects of AKS
type AzureKubernetesServiceExpander struct {
	client *http.Client
}

// Name returns the name of the expander
func (e *AzureKubernetesServiceExpander) Name() string {
	return "AzureKubernetesServiceExpander"
}

// DoesExpand checks if this is a storage account
func (e *AzureKubernetesServiceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "AzureKubernetesService" {
		return true, nil
	}
	return false, nil
}

// Expand returns ManagementPolicies in the StorageAccount
func (e *AzureKubernetesServiceExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "AzureKubernetesService" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}" {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			ID:        currentItem.ID + "/<k8sapi>",
			Parentid:  currentItem.ID,
			Namespace: "AzureKubernetesService",
			Name:      "Kubernetes API",
			Display:   "Kubernetes API",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"ClusterID":             currentItem.ID, // save full URL to registry
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          "", // Swagger expander will supply the response
			SourceDescription: "AzureKubernetesServiceExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "AzureKubernetesService" && currentItem.ItemType == SubResourceType {
		return e.expandKubernetesApiRoot(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          "Error!",
		SourceDescription: "AzureKubernetesServiceExpander request",
	}
}

func (e *AzureKubernetesServiceExpander) expandKubernetesApiRoot(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	clusterID := currentItem.Metadata["ClusterID"]

	token, err := e.getClusterToken(ctx, clusterID)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "Error!",
			SourceDescription: "AzureKubernetesServiceExpander request",
		}
	}

	newItems := []*TreeNode{}
	return ExpanderResult{
		Err:               nil,
		Response:          token,
		SourceDescription: "AzureKubernetesServiceExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureKubernetesServiceExpander) getClusterToken(ctx context.Context, clusterID string) (string, error) {

	data, err := armclient.DoRequest(ctx, "POST", clusterID+"/listClusterAdminCredential?api-version=2019-08-01")
	if err != nil {
		return "", fmt.Errorf("Failed to get credentials: " + err.Error() + clusterID)
	}

	var response clusterCredentialsResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, clusterID)
		return "", err
	}

	if len(response.KubeConfigs) < 1 {
		err = fmt.Errorf("Response has no KubeConfigs\nURL:%s", clusterID)
		return "", err
	}

	configBase64 := response.KubeConfigs[0].Value

	config, err := base64.StdEncoding.DecodeString(configBase64)
	if err != nil {
		err = fmt.Errorf("Error decoding kubeconfig: %s\nURL:%s", err, clusterID)
		return "", err
	}

	var kubeConfig kubeConfig
	err = yaml.Unmarshal(config, &kubeConfig)
	if err != nil {
		err = fmt.Errorf("Error parsing kubeconfig: %s\nURL:%s", err, clusterID)
		return "", err
	}

	return kubeConfig.Users[0].User.Token, nil
}
