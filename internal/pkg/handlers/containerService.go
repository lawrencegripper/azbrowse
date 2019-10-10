package handlers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"

	"github.com/go-openapi/loads"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type clusterCredentialsResponse struct {
	KubeConfigs []struct {
		Name  string `json:"name"`
		Value string `json:"Value"`
	} `json:"kubeconfigs"`
}

// kubeConfig is a minimal struct for parsing the parts of the response that we care about
type kubeConfig struct {
	Clusters []struct {
		Name    string `yaml:"name"`
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
	} `yaml: clusters`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Token                 string `yaml:"token"`
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

	kubeConfig, err := e.getClusterConfig(ctx, clusterID)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "Error!",
			SourceDescription: "AzureKubernetesServiceExpander request",
		}
	}

	// NOTE - at the time of writing the AKS API returns a single cluster/user
	// so we're not fully parsing the config, just taking the first user and cluster

	httpClient, err := e.getHttpClientFromConfig(kubeConfig)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "Error!",
			SourceDescription: "AzureKubernetesServiceExpander request",
		}
	}

	serverUrl := kubeConfig.Clusters[0].Cluster.Server

	swaggerResourceTypes, err := e.getSwaggerResourceTypes(*httpClient, serverUrl)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          "Error!",
			SourceDescription: "AzureKubernetesServiceExpander request",
		}
	}

	// Register the swagger config so that the swagger expander can take over
	config := NewSwaggerConfigContainerService(swaggerResourceTypes, *httpClient, currentItem.ID, serverUrl)
	GetSwaggerResourceExpander().AddConfig(config)

	// TODO think about how to avoid re-registering - add something to the current node's metadata?
	newItems := []*TreeNode{}
	for _, child := range swaggerResourceTypes {
		resourceType := child
		display := resourceType.Display
		if display == "{}" {
			display = resourceType.Endpoint.TemplateURL
		}
		newItems = append(newItems, &TreeNode{
			Parentid:            currentItem.ID,
			ID:                  currentItem.ID + "/" + display,
			Namespace:           "swagger",
			Name:                display,
			Display:             display,
			ExpandURL:           resourceType.Endpoint.TemplateURL, // all fixed template URLs
			ItemType:            SubResourceType,
			SwaggerResourceType: &resourceType,
			Metadata: map[string]string{
				"SwaggerConfigID": currentItem.ID,
			},
		})
	}

	return ExpanderResult{
		Err:               nil,
		Response:          "TODO - what should go here?",
		SourceDescription: "AzureKubernetesServiceExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureKubernetesServiceExpander) getSwaggerResourceTypes(httpClient http.Client, serverUrl string) ([]swagger.SwaggerResourceType, error) {

	var swaggerResourceTypes []swagger.SwaggerResourceType

	url := serverUrl + "/openapi/v2"

	response, err := httpClient.Get(url)
	if err != nil {
		return swaggerResourceTypes, err
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return swaggerResourceTypes, err
	}

	spec := json.RawMessage(buf)
	doc, err := loads.Analyzed(spec, "")
	if err != nil {
		return swaggerResourceTypes, err
	}

	config := swagger.Config{}
	var paths []*swagger.Path
	paths, err = swagger.MergeSwaggerDoc(paths, &config, doc)
	if err != nil {
		return swaggerResourceTypes, err
	}

	swaggerResourceTypes = swagger.ConvertToSwaggerResourceTypes(paths)
	return swaggerResourceTypes, nil
}

func (e *AzureKubernetesServiceExpander) getHttpClientFromConfig(kubeConfig kubeConfig) (*http.Client, error) {

	clientCertificate, err := base64.StdEncoding.DecodeString(kubeConfig.Users[0].User.ClientCertificateData)
	if err != nil {
		err = fmt.Errorf("Error decoding client certificate data: %s", err)
		return nil, err
	}
	clientKey, err := base64.StdEncoding.DecodeString(kubeConfig.Users[0].User.ClientKeyData)
	if err != nil {
		err = fmt.Errorf("Error decoding client key data: %s", err)
		return nil, err
	}
	certificateAuthority, err := base64.StdEncoding.DecodeString(kubeConfig.Clusters[0].Cluster.CertificateAuthorityData)
	if err != nil {
		err = fmt.Errorf("Error decoding certificate authority data: %s", err)
		return nil, err
	}
	_ = certificateAuthority

	cert, err := tls.X509KeyPair(clientCertificate, clientKey)
	if err != nil {
		return nil, err
	}

	caCerts, err := x509.SystemCertPool()
	if err != nil {
		err = fmt.Errorf("Error creating certpool: %s", err)
		return nil, err

	}
	caCerts.AppendCertsFromPEM(certificateAuthority)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCerts,
		},
	}

	httpClient := http.Client{
		Transport: transport,
	}

	return &httpClient, nil

}

func (e *AzureKubernetesServiceExpander) getClusterConfig(ctx context.Context, clusterID string) (kubeConfig, error) {

	data, err := armclient.DoRequest(ctx, "POST", clusterID+"/listClusterUserCredential?api-version=2019-08-01")
	if err != nil {
		return kubeConfig{}, fmt.Errorf("Failed to get credentials: " + err.Error() + clusterID)
	}

	var response clusterCredentialsResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, clusterID)
		return kubeConfig{}, err
	}

	if len(response.KubeConfigs) < 1 {
		err = fmt.Errorf("Response has no KubeConfigs\nURL:%s", clusterID)
		return kubeConfig{}, err
	}

	configBase64 := response.KubeConfigs[0].Value

	config, err := base64.StdEncoding.DecodeString(configBase64)
	if err != nil {
		err = fmt.Errorf("Error decoding kubeconfig: %s\nURL:%s", err, clusterID)
		return kubeConfig{}, err
	}

	var kubeConfig kubeConfig
	err = yaml.Unmarshal(config, &kubeConfig)
	if err != nil {
		err = fmt.Errorf("Error parsing kubeconfig: %s\nURL:%s", err, clusterID)
		return kubeConfig, err
	}

	return kubeConfig, nil
}
