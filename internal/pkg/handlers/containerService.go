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

	result, err := e.test(ctx, kubeConfig)
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
		Response:          result,
		SourceDescription: "AzureKubernetesServiceExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureKubernetesServiceExpander) test(ctx context.Context, kubeConfig kubeConfig) (string, error) {

	clientCertificate, err := base64.StdEncoding.DecodeString(kubeConfig.Users[0].User.ClientCertificateData)
	if err != nil {
		err = fmt.Errorf("Error decoding client certificate data: %s", err)
		return "", err
	}
	clientKey, err := base64.StdEncoding.DecodeString(kubeConfig.Users[0].User.ClientKeyData)
	if err != nil {
		err = fmt.Errorf("Error decoding client key data: %s", err)
		return "", err
	}
	certificateAuthority, err := base64.StdEncoding.DecodeString(kubeConfig.Clusters[0].Cluster.CertificateAuthorityData)
	if err != nil {
		err = fmt.Errorf("Error decoding certificate authority data: %s", err)
		return "", err
	}
	_ = certificateAuthority

	cert, err := tls.X509KeyPair(clientCertificate, clientKey)
	if err != nil {
		return "", err
	}

	caCerts, err := x509.SystemCertPool()
	if err != nil {
		err = fmt.Errorf("Error creating certpool: %s", err)
		return "", err

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

	url := kubeConfig.Clusters[0].Cluster.Server + "/api/v1/nodes"
	response, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
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
