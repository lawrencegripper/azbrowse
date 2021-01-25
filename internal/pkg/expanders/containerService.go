package expanders

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/go-openapi/loads"

	azbrowse_config "github.com/lawrencegripper/azbrowse/internal/pkg/config"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

const clusterTemplateURL string = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}"

type clusterCredentialsResponse struct {
	KubeConfigs []struct {
		Name  string `json:"name"`
		Value string `json:"Value"`
	} `json:"kubeconfigs"`
}

// kubeConfigResponse is a minimal struct for parsing the parts of the response that we care about
type kubeConfigResponse struct {
	Clusters []struct {
		Name    string `yaml:"name"`
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
	} `yaml:"clusters"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Token                 string `yaml:"token"`
		} `yaml:"user"`
	} `yaml:"users"`
}

// Check interface
var _ Expander = &AzureKubernetesServiceExpander{}

// AzureKubernetesServiceExpander expands the kubernetes aspects of AKS
type AzureKubernetesServiceExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *AzureKubernetesServiceExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *AzureKubernetesServiceExpander) Name() string {
	return "AzureKubernetesServiceExpander"
}

// DoesExpand checks if this is a storage account
func (e *AzureKubernetesServiceExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == "resource" && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == clusterTemplateURL {
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
		swaggerResourceType.Endpoint.TemplateURL == clusterTemplateURL {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			ID:                    currentItem.ID + "/<k8sapi>",
			Parentid:              currentItem.ID,
			Namespace:             "AzureKubernetesService",
			Name:                  "Kubernetes API",
			Display:               "Kubernetes API",
			ItemType:              SubResourceType,
			ExpandURL:             ExpandURLNotSupported,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"ClusterID": currentItem.ID, // save full URL to registry
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "AzureKubernetesServiceExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "AzureKubernetesService" && currentItem.ItemType == SubResourceType {
		return e.expandKubernetesAPIRoot(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "AzureKubernetesServiceExpander request",
	}
}

func (e *AzureKubernetesServiceExpander) expandKubernetesAPIRoot(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	clusterID := currentItem.Metadata["ClusterID"]

	// Check for existing config for the cluster
	apiSet := e.getAPISetForCluster(clusterID)
	var err error
	if apiSet == nil {
		apiSet, err = e.createAPISetForCluster(ctx, clusterID)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				Response:          ExpanderResponse{Response: "Error!"},
				SourceDescription: "AzureKubernetesServiceExpander request",
			}
		}
		GetSwaggerResourceExpander().AddAPISet(*apiSet)
	}

	swaggerResourceTypes := apiSet.GetResourceTypes()

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
				"SwaggerAPISetID": currentItem.ID,
			},
		})
	}

	return ExpanderResult{
		Err:               nil,
		Response:          ExpanderResponse{Response: ""},
		SourceDescription: "AzureKubernetesServiceExpander request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

func (e *AzureKubernetesServiceExpander) createAPISetForCluster(ctx context.Context, clusterID string) (*SwaggerAPISetContainerService, error) {
	kubeConfig, err := e.getClusterConfig(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	// NOTE - at the time of writing the AKS API returns a single cluster/user
	// so we're not fully parsing the config, just taking the first user and cluster

	httpClient, err := e.getHTTPClientFromConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	serverURL := kubeConfig.Clusters[0].Cluster.Server

	swaggerResourceTypes, err := e.getSwaggerResourceTypes(*httpClient, serverURL)
	if err != nil {
		return nil, err
	}

	// Register the swagger config so that the swagger expander can take over
	apiSet := NewSwaggerAPISetContainerService(swaggerResourceTypes, *httpClient, clusterID+"/<k8sapi>", serverURL)
	return &apiSet, nil
}
func (e *AzureKubernetesServiceExpander) getAPISetForCluster(clusterID string) *SwaggerAPISetContainerService {

	swaggerAPISet := GetSwaggerResourceExpander().GetAPISet(clusterID + "/<k8sapi>")
	if swaggerAPISet == nil {
		return nil
	}
	apiSet := (*swaggerAPISet).(SwaggerAPISetContainerService)
	return &apiSet
}

func (e *AzureKubernetesServiceExpander) getSwaggerResourceTypes(httpClient http.Client, serverURL string) ([]swagger.ResourceType, error) {

	var swaggerResourceTypes []swagger.ResourceType

	url := serverURL + "/openapi/v2"

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

	config := swagger.Config{
		AdditionalPaths: []swagger.AdditionalPath{
			// add a placeholder for /api/v1/watch/* as otherwise they end up directly under /api/v1 and show as duplicates of /api/v1/nodes etc
			{Name: "watch", Path: "/api/v1/watch", FixedContent: "Select a node to expand"},
			// add a placeholder for /api/v1/watch/* as otherwise they end up directly under /api/v1 and show as duplicates of /api/v1/nodes etc
			{Name: "watch", Path: "/apis/apps/v1/watch", FixedContent: "Select a node to expand"},
			// add as a missing path - and direct to a different endpoint
			{
				Name:    "namespaces",
				Path:    "/apis/apps/v1/namespaces",
				GetPath: "/api/v1/namespaces",
				SubPathRegex: &swagger.RegexReplace{
					Match:   "/api/v1/namespaces/",
					Replace: "/apis/apps/v1/namespaces/",
				},
			},
			// add as a missing path - also overridden to map to the actual endpoint that exists!
			{Name: "{namespace}", Path: "/apis/apps/v1/namespaces/{namespace}", FixedContent: "Select a node to expand"},
		},
		SuppressAPIVersion: true,
	}
	var paths []*swagger.Path
	paths, err = swagger.MergeSwaggerDoc(paths, &config, doc, false, "")
	if err != nil {
		return swaggerResourceTypes, err
	}

	if azbrowse_config.GetDebuggingEnabled() {
		tempBuf, err := yaml.Marshal(paths)
		if err != nil {
			return swaggerResourceTypes, err
		}
		tmpDir := os.Getenv("TEMP")
		if tmpDir == "" {
			tmpDir = "/tmp"
		}
		tmpPath := tmpDir + "/k8s-paths.yml"
		ioutil.WriteFile(tmpPath, tempBuf, 0644) //nolint:errcheck
	}

	swaggerResourceTypes = swagger.ConvertToSwaggerResourceTypes(paths)
	return swaggerResourceTypes, nil
}

func (e *AzureKubernetesServiceExpander) getHTTPClientFromConfig(kubeConfig kubeConfigResponse) (*http.Client, error) {

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

func (e *AzureKubernetesServiceExpander) getClusterConfig(ctx context.Context, clusterID string) (kubeConfigResponse, error) {

	data, err := e.client.DoRequest(ctx, "POST", clusterID+"/listClusterUserCredential?api-version=2019-08-01")
	if err != nil {
		return kubeConfigResponse{}, fmt.Errorf("Failed to get credentials: " + err.Error() + clusterID)
	}

	var response clusterCredentialsResponse
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, clusterID)
		return kubeConfigResponse{}, err
	}

	if len(response.KubeConfigs) < 1 {
		err = fmt.Errorf("Response has no KubeConfigs\nURL:%s", clusterID)
		return kubeConfigResponse{}, err
	}

	configBase64 := response.KubeConfigs[0].Value

	config, err := base64.StdEncoding.DecodeString(configBase64)
	if err != nil {
		err = fmt.Errorf("Error decoding kubeconfig: %s\nURL:%s", err, clusterID)
		return kubeConfigResponse{}, err
	}

	var kubeConfig kubeConfigResponse
	err = yaml.Unmarshal(config, &kubeConfig)
	if err != nil {
		err = fmt.Errorf("Error parsing kubeconfig: %s\nURL:%s", err, clusterID)
		return kubeConfig, err
	}

	return kubeConfig, nil
}
