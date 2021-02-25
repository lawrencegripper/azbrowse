package expanders

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// NewStorageBlobExpander creates a new instance of StorageBlobExpander
func NewStorageBlobExpander(armclient *armclient.Client) *StorageBlobExpander {
	return &StorageBlobExpander{
		client:    &http.Client{},
		armClient: armclient,
	}
}

// Check interface
var _ Expander = &StorageBlobExpander{}

// StorageListKeyResponse is used to unmarshal a call to listKeys on a storage account
type StorageListKeyResponse struct {
	Keys []struct {
		KeyName     string `json:"keyName"`
		Value       string `json:"value"`
		Permissions string `json:"permissions"`
	} `json:"keys"`
}

// StorageAccountResponse is a partial representation of the storage account response
type StorageAccountResponse struct {
	Properties struct {
		PrimaryEndpoints struct {
			Blob string `json:"blob"`
			Dfs  string `json:"dfs"`
		} `json:"primaryEndpoints"`
	} `json:"properties"`
}

// ContainerListResponse is a partial representation of the List container response
type ContainerListResponse struct {
	XMLName    xml.Name `xml:"EnumerationResults"`
	Blobs      []Blob   `xml:"Blobs>Blob"`
	NextMarker string   `xml:"NextMarker"`
}

type Blob struct {
	Name       string `xml:"Name"`
	Properties struct {
		CreationTime       string `xml:"Creation-Time"`
		LastModified       string `xml:"Last-Modified"`
		Etag               string `xml:"Etag"`
		ContentLength      int    `xml:"Content-Length"`
		ContentEncoding    string `xml:"Content-Encoding"`
		ContentLanguage    string `xml:"Content-Language"`
		ContentMD5         string `xml:"Content-MD5"`
		CacheControl       string `xml:"Cache-Control"`
		ContentDisposition string `xml:"Content-Disposition"`
		BlobType           string `xml:"BlobType"`
		AccessTier         string `xml:"AccessTier"`
		AccessTierInferred string `xml:"AccessTierInferred"`
		LeaseStatus        string `xml:"LeaseStatus"`
		LeaseState         string `xml:"LeaseState"`
		ServerEncrypted    string `xml:"ServerEncrypted"`
	} `xml:"Properties"`
}

const (
	storageBlobNodeBlobMetadata     = "blob-metadata"
	storageBlobNodeListBlobMetadata = "blob-metadata-list"
	storageBlobNodeBlob             = "blob"
	storageBlobNodeListBlob         = "blob-list"
)

const (
	storageBlobActionLeaseAcquire = "lease-acquire"
	storageBlobActionLeaseBreak   = "lease-break"
)

func (e *StorageBlobExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// StorageBlobExpander expands the blob  data-plane aspects of a Storage Account
type StorageBlobExpander struct {
	ExpanderBase
	client    *http.Client
	armClient *armclient.Client
}

// Name returns the name of the expander
func (e *StorageBlobExpander) Name() string {
	return "StorageBlobExpander"
}

// DoesExpand checks if this is a storage account
func (e *StorageBlobExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == SubResourceType && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "storageBlob" {
		return true, nil
	}
	return false, nil
}

// Expand returns blobs in the StorageAccount congtainer
func (e *StorageBlobExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "storageBlob" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}" {
		newItems := []*TreeNode{
			{
				Parentid:              currentItem.ID,
				ID:                    currentItem.ID + "/<blobs>",
				Namespace:             "storageBlob",
				Name:                  "Blob Metadata",
				Display:               "Blob Metadata",
				ItemType:              storageBlobNodeListBlobMetadata,
				ExpandURL:             ExpandURLNotSupported,
				SuppressSwaggerExpand: true,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"ContainerID": currentItem.ExpandURL, // save resourceID of blob
				},
			},
			{
				Parentid:              currentItem.ID,
				ID:                    currentItem.ID + "/<blobs>",
				Namespace:             "storageBlob",
				Name:                  "Blobs",
				Display:               "Blobs",
				ItemType:              storageBlobNodeListBlob,
				ExpandURL:             ExpandURLNotSupported,
				SuppressSwaggerExpand: true,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"ContainerID": currentItem.ExpandURL, // save resourceID of blob
				},
			},
		}

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "StorageBlobExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	switch currentItem.ItemType {
	case storageBlobNodeListBlobMetadata:
		return e.expandMetadataList(ctx, currentItem)
	case storageBlobNodeBlobMetadata:
		return e.expandMetadata(ctx, currentItem)
	case storageBlobNodeListBlob:
		return e.expandBlobList(ctx, currentItem)
	case storageBlobNodeBlob:
		return e.expandBlob(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "StorageBlobExpander request",
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *StorageBlobExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {
	switch currentItem.ItemType {
	case storageBlobNodeBlob, storageBlobNodeBlobMetadata:
		return e.deleteBlob(ctx, currentItem)
	}
	return false, nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *StorageBlobExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	switch item.ItemType {
	case storageBlobNodeBlob,
		storageBlobNodeBlobMetadata:
		return true, nil
	}
	return false, nil

}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *StorageBlobExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {
	nodes := []*TreeNode{}
	switch item.ItemType {
	case storageBlobNodeBlob,
		storageBlobNodeBlobMetadata:
		nodes = append(nodes,
			&TreeNode{
				Parentid:              item.ID,
				ID:                    item.ID + "?lease-acquire",
				Namespace:             "storageBlob",
				Name:                  "Acquire Lease",
				Display:               "Acquire Lease",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"ActionID":      storageBlobActionLeaseAcquire,
					"BlobName":      item.Metadata["BlobName"],
					"ContainerID":   item.Metadata["ContainerID"],
					"ContainerName": item.Metadata["ContainerName"],
					"AccountName":   item.Metadata["AccountName"],
					"AccountKey":    item.Metadata["AccountKey"],
					"BlobEndpoint":  item.Metadata["BlobEndpoint"],
				},
			},
			&TreeNode{
				Parentid:              item.ID,
				ID:                    item.ID + "?lease-break",
				Namespace:             "storageBlob",
				Name:                  "Break Lease",
				Display:               "Break Lease",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"ActionID":      storageBlobActionLeaseBreak,
					"BlobName":      item.Metadata["BlobName"],
					"ContainerID":   item.Metadata["ContainerID"],
					"ContainerName": item.Metadata["ContainerName"],
					"AccountName":   item.Metadata["AccountName"],
					"AccountKey":    item.Metadata["AccountKey"],
					"BlobEndpoint":  item.Metadata["BlobEndpoint"],
				},
			},
		)
	default:
		return ListActionsResult{
			SourceDescription: "StorageBlobExpander",
			Err:               fmt.Errorf("ListActions not supported for ItemType %q", item.ItemType),
		}
	}
	return ListActionsResult{
		Nodes:             nodes,
		SourceDescription: "StorageBlobExpander",
		IsPrimaryResponse: true,
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *StorageBlobExpander) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	actionID := item.Metadata["ActionID"]

	switch actionID {
	case storageBlobActionLeaseAcquire:
		return e.storageBlobLeaseAcquire(context, item)
	case storageBlobActionLeaseBreak:
		return e.storageBlobLeaseBreak(context, item)
	case "":
		return ExpanderResult{
			SourceDescription: "StorageBlobExpander",
			Err:               fmt.Errorf("ActionID metadata not set: %q", item.ID),
		}
	default:
		return ExpanderResult{
			SourceDescription: "StorageBlobExpander",
			Err:               fmt.Errorf("Unhandled ActionID: %q", actionID),
		}
	}
}

func (e *StorageBlobExpander) storageBlobLeaseBreak(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	containerName := currentItem.Metadata["ContainerName"]
	accountName := currentItem.Metadata["AccountName"]
	accountKey := currentItem.Metadata["AccountKey"]
	blobName := currentItem.Metadata["BlobName"]
	blobEndpoint := currentItem.Metadata["BlobEndpoint"]

	// Lease Blob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/lease-blob
	url := blobEndpoint + containerName + "/" + blobName + "?comp=lease"
	headers := map[string]string{
		"x-ms-lease-action":       "break",
		"x-ms-lease-break-period": "0",
	}
	_, err := e.doRequestWithHeaders(ctx, "PUT", url, accountName, accountKey, "/"+accountName+"/"+containerName, headers)

	if err != nil {
		return ExpanderResult{
			Response: ExpanderResponse{
				ResponseType: interfaces.ResponsePlainText,
				Response:     fmt.Sprintf("Error breaking blob lease: %s", err),
			},
			SourceDescription: "StorageBlobExpander request",
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		Response: ExpanderResponse{
			ResponseType: interfaces.ResponsePlainText,
			Response:     "Success",
		},
		SourceDescription: "StorageBlobExpander request",
		IsPrimaryResponse: true,
	}
}
func (e *StorageBlobExpander) storageBlobLeaseAcquire(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	containerName := currentItem.Metadata["ContainerName"]
	accountName := currentItem.Metadata["AccountName"]
	accountKey := currentItem.Metadata["AccountKey"]
	blobName := currentItem.Metadata["BlobName"]
	blobEndpoint := currentItem.Metadata["BlobEndpoint"]

	// Lease Blob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/lease-blob
	url := blobEndpoint + containerName + "/" + blobName + "?comp=lease"
	headers := map[string]string{
		"x-ms-lease-action":   "acquire",
		"x-ms-lease-duration": "-1",
	}
	_, err := e.doRequestWithHeaders(ctx, "PUT", url, accountName, accountKey, "/"+accountName+"/"+containerName, headers)

	if err != nil {
		return ExpanderResult{
			Response: ExpanderResponse{
				ResponseType: interfaces.ResponsePlainText,
				Response:     fmt.Sprintf("Error acquiring blob lease: %s", err),
			},
			SourceDescription: "StorageBlobExpander request",
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		Response: ExpanderResponse{
			ResponseType: interfaces.ResponsePlainText,
			Response:     "Success",
		},
		SourceDescription: "StorageBlobExpander request",
		IsPrimaryResponse: true,
	}
}

func (e *StorageBlobExpander) expandMetadataList(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	return e.expandList(
		ctx,
		currentItem,
		func(currentItem *TreeNode, blob Blob) (*TreeNode, error) {
			// Had issues with these APIs so just breaking out the XML from the list for the item content :-)
			// https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-metadata
			// https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-properties
			content, err := xml.MarshalIndent(blob, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("Error marshaling blob: %s", err)
			}

			node := TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "storageBlob",
				ID:        currentItem.ID + "/" + blob.Name,
				Name:      blob.Name,
				Display:   blob.Name,
				ItemType:  storageBlobNodeBlobMetadata,
				ExpandURL: ExpandURLNotSupported,
				DeleteURL: currentItem.ID + "/" + blob.Name,
				Metadata: map[string]string{
					"BlobName": blob.Name,
					"Content":  string(content),
				},
			}

			return &node, nil
		},
		storageBlobNodeListBlobMetadata)
}

func (e *StorageBlobExpander) expandBlobList(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	return e.expandList(
		ctx,
		currentItem,
		func(currentItem *TreeNode, blob Blob) (*TreeNode, error) {
			node := TreeNode{
				Parentid:  currentItem.ID,
				Namespace: "storageBlob",
				ID:        currentItem.ID + "/" + blob.Name,
				Name:      blob.Name,
				Display:   blob.Name,
				ItemType:  storageBlobNodeBlob,
				ExpandURL: ExpandURLNotSupported,
				DeleteURL: currentItem.ID + "/" + blob.Name,
				Metadata: map[string]string{
					"BlobName": blob.Name,
				},
			}

			return &node, nil
		},
		storageBlobNodeListBlob)
}

func (e *StorageBlobExpander) expandList(ctx context.Context, currentItem *TreeNode, createNodeFunc func(currentItem *TreeNode, blob Blob) (*TreeNode, error), continuationItemType string) ExpanderResult {

	// https://docs.microsoft.com/en-us/rest/api/storageservices/enumerating-blob-resources#Subheading5

	containerID := currentItem.Metadata["ContainerID"]
	containerName := e.getContainerName(containerID)
	accountName := e.getAccountName(containerID)
	accountKey, err := e.getAccountKey(ctx, containerID)
	marker := currentItem.Metadata["Marker"]
	if err != nil {
		err = fmt.Errorf("Error getting account key: %s", err)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "StorageBlobExpander request",
		}
	}
	blobEndpoint, err := e.getStorageBlobEndpoint(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("Error getting blob endpoint: %s", err)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "StorageBlobExpander request",
		}
	}

	// ListBlob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/list-blobs
	url := blobEndpoint + containerName + "?restype=container&comp=list&maxresults=50"
	if marker != "" {
		url += "&marker=" + marker
	}
	buf, err := e.doRequest(ctx, "GET", url, accountName, accountKey, "/"+accountName+"/"+containerName)

	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error listing blobs: %s", err),
			SourceDescription: "StorageBlobExpander request",
		}
	}

	response := &ContainerListResponse{}
	err = xml.Unmarshal(buf, response)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error Unmarshalling ContainerListResponse: %s", err),
			SourceDescription: "StorageBlobExpander request",
		}
	}
	nodes := []*TreeNode{}

	for _, blob := range response.Blobs {
		node, err := createNodeFunc(currentItem, blob)

		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "StorageBlobExpander request",
			}
		}
		node.Metadata["ContainerID"] = containerID
		node.Metadata["ContainerName"] = containerName
		node.Metadata["AccountName"] = accountName
		node.Metadata["AccountKey"] = accountKey
		node.Metadata["BlobEndpoint"] = blobEndpoint

		nodes = append(nodes, node)
	}
	if response.NextMarker != "" {
		node := TreeNode{
			Parentid:      currentItem.ID,
			Namespace:     "storageBlob",
			ID:            currentItem.ID + "/" + "...more",
			Name:          "more...",
			Display:       "more...",
			ItemType:      continuationItemType,
			ExpandURL:     ExpandURLNotSupported,
			ExpandInPlace: true,
			Metadata: map[string]string{
				"ContainerID":   containerID, // save resourceID of blob
				"ContainerName": containerName,
				"AccountName":   accountName,
				"AccountKey":    accountKey,
				"BlobEndpoint":  blobEndpoint,
				"Marker":        response.NextMarker,
			},
		}

		nodes = append(nodes, &node)
	}

	result := string(buf)
	return ExpanderResult{
		Response:          ExpanderResponse{Response: result, ResponseType: interfaces.ResponseXML},
		SourceDescription: "StorageBlobExpander request",
		Nodes:             nodes,
		IsPrimaryResponse: true,
	}
}

func (e *StorageBlobExpander) expandMetadata(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	containerName := currentItem.Metadata["ContainerName"]
	accountName := currentItem.Metadata["AccountName"]
	accountKey := currentItem.Metadata["AccountKey"]
	blobName := currentItem.Metadata["BlobName"]
	blobEndpoint := currentItem.Metadata["BlobEndpoint"]

	// Blob Properties: https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-properties
	url := blobEndpoint + containerName + "/" + blobName
	_, headers, err := e.doRequestWithHeadersIncludeResponseHeaders(ctx, "HEAD", url, accountName, accountKey, "/"+accountName+"/"+containerName, map[string]string{})

	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting blob metadata: %s", err),
			SourceDescription: "StorageBlobExpander request",
			IsPrimaryResponse: true,
		}
	}

	simpleHeaders := map[string]string{}
	for k := range headers {
		simpleHeaders[k] = headers.Get(k)
	}
	buf, err := json.Marshal(simpleHeaders)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error marshaling blob metadata to JSON: %s", err),
			SourceDescription: "StorageBlobExpander request",
			IsPrimaryResponse: true,
		}
	}
	content := string(buf)
	return ExpanderResult{
		Response:          ExpanderResponse{Response: content, ResponseType: interfaces.ResponseJSON},
		SourceDescription: "StorageBlobExpander request",
		Nodes:             []*TreeNode{},
		IsPrimaryResponse: true,
	}
}

func (e *StorageBlobExpander) expandBlob(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	containerID := currentItem.Metadata["ContainerID"]
	containerName := e.getContainerName(containerID)
	accountName := currentItem.Metadata["AccountName"]
	accountKey := currentItem.Metadata["AccountKey"]
	blobName := currentItem.Metadata["BlobName"]

	blobEndpoint, err := e.getStorageBlobEndpoint(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("Error getting blob endpoint: %s", err)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "StorageBlobExpander request",
		}
	}

	// GetBlob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob
	url := blobEndpoint + containerName + "/" + blobName
	buf, err := e.doRequest(ctx, "GET", url, accountName, accountKey, "/"+accountName+"/"+containerName+"/"+blobName)

	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting blob: %s", err),
			SourceDescription: "StorageBlobExpander request",
		}
	}

	result := string(buf)
	return ExpanderResult{
		Response:          ExpanderResponse{Response: result, ResponseType: interfaces.ResponsePlainText},
		SourceDescription: "StorageBlobExpander request",
		Nodes:             []*TreeNode{},
		IsPrimaryResponse: true,
	}

}

func (e *StorageBlobExpander) deleteBlob(ctx context.Context, currentItem *TreeNode) (bool, error) {

	containerID := currentItem.Metadata["ContainerID"]
	containerName := e.getContainerName(containerID)
	accountName := currentItem.Metadata["AccountName"]
	accountKey := currentItem.Metadata["AccountKey"]
	blobName := currentItem.Metadata["BlobName"]

	blobEndpoint, err := e.getStorageBlobEndpoint(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("Error getting blob endpoint: %s", err)
		return false, err
	}

	// DeleteBlob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/delete-blob
	url := blobEndpoint + containerName + "/" + blobName
	_, err = e.doRequest(ctx, "DELETE", url, accountName, accountKey, "/"+accountName+"/"+containerName+"/"+blobName)

	if err != nil {
		return false, fmt.Errorf("Error deleting blob: %s", err)
	}

	return true, nil

}

func (e *StorageBlobExpander) getAccountName(containerID string) string {
	i := strings.Index(containerID, "/blobServices")
	accountID := containerID[0:i]
	i = strings.LastIndex(accountID, "/")
	return accountID[i+1:]
}
func (e *StorageBlobExpander) getContainerName(containerID string) string {
	i := strings.LastIndex(containerID, "/")
	containerName := containerID[i+1:]
	i = strings.Index(containerName, "?")
	if i >= 0 {
		containerName = containerName[:i] // strip query string
	}
	return containerName
}

func (e *StorageBlobExpander) getAccountKey(ctx context.Context, containerID string) (string, error) {

	i := strings.Index(containerID, "/blobServices")
	rootURL := containerID[0:i]
	listKeysURL := rootURL + "/listKeys?api-version=2019-06-01"

	data, err := e.armClient.DoRequest(ctx, "POST", listKeysURL)
	if err != nil {
		return "", fmt.Errorf("Error calling listKeys: %s", err)
	}
	response := StorageListKeyResponse{}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, listKeysURL)
		return "", err
	}
	if len(response.Keys) == 0 {
		err = fmt.Errorf("No keys in response: %s", err)
		return "", err
	}

	return response.Keys[0].Value, nil
}
func (e *StorageBlobExpander) getStorageBlobEndpoint(ctx context.Context, containerID string) (string, error) {

	i := strings.Index(containerID, "/blobServices")
	rootURL := containerID[0:i]
	storageAccountURL := rootURL + "?api-version=2019-06-01"

	data, err := e.armClient.DoRequest(ctx, "GET", storageAccountURL)
	if err != nil {
		return "", fmt.Errorf("Error getting storage account: %s", err)
	}
	response := StorageAccountResponse{}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, storageAccountURL)
		return "", err
	}

	return response.Properties.PrimaryEndpoints.Blob, nil
}

func (e *StorageBlobExpander) doRequest(ctx context.Context, verb string, url string, accountName string, accountKey string, accountAndPath string) ([]byte, error) {
	return e.doRequestWithHeaders(ctx, verb, url, accountName, accountKey, accountAndPath, map[string]string{})
}
func (e *StorageBlobExpander) doRequestWithHeaders(ctx context.Context, verb string, url string, accountName string, accountKey string, accountAndPath string, headers map[string]string) ([]byte, error) {
	buf, _, err := e.doRequestWithHeadersIncludeResponseHeaders(ctx, verb, url, accountName, accountKey, accountAndPath, headers)
	return buf, err
}
func (e *StorageBlobExpander) doRequestWithHeadersIncludeResponseHeaders(ctx context.Context, verb string, url string, accountName string, accountKey string, accountAndPath string, headers map[string]string) ([]byte, http.Header, error) {

	span, _ := tracing.StartSpanFromContext(ctx, "doRequest(blobexpnder):"+url, tracing.SetTag("url", url))
	defer span.Finish()

	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("Failed to create request: %s", err)
	}
	if req.Header.Get("x-ms-version") == "" {
		req.Header.Set("x-ms-version", "2018-03-28")
	}
	dateString := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Set("x-ms-date", dateString)

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	err = e.addAuthHeader(req, accountName, accountKey)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("Failed to add auth header: %s", err)
	}

	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, nil, fmt.Errorf("Request failed: %s", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		defer response.Body.Close() //nolint: errcheck
		return []byte{}, nil, fmt.Errorf("DoRequest failed %v for '%s'", response.Status, url)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("Failed to read body: %s", err)
	}

	buf = e.stripBOM(buf)

	return buf, response.Header, nil
}

func (e *StorageBlobExpander) stripBOM(buf []byte) []byte {
	if len(buf) < 3 {
		return buf
	}
	if buf[0] == 0xEF && buf[1] == 0xBB && buf[2] == 0xBF {
		return buf[3:]
	}
	return buf
}

// Auth helper code based on https://github.com/Azure/azure-storage-blob-go
// (https://github.com/Azure/azure-storage-blob-go/blob/3efca72bd11c050222deab57e25ea90df03b9692/azblob/zc_credential_shared_key.go#L55)
func (e *StorageBlobExpander) addAuthHeader(request *http.Request, accountName string, accountKey string) error {

	// Add a x-ms-date header if it doesn't already exist
	if d := request.Header.Get(headerXmsDate); d == "" {
		request.Header[headerXmsDate] = []string{time.Now().UTC().Format(http.TimeFormat)}
	}
	stringToSign, err := e.buildStringToSign(request, accountName)
	if err != nil {
		return fmt.Errorf("Failed to build string to sign: %s", err)
	}
	signature, err := e.ComputeHMACSHA256(stringToSign, accountKey)
	if err != nil {
		return fmt.Errorf("Failed to compute signature: %s", err)
	}
	authHeader := strings.Join([]string{"SharedKey ", accountName, ":", signature}, "")
	request.Header[headerAuthorization] = []string{authHeader}
	return nil
}

// Constants ensuring that header names are correctly spelled and consistently cased.
const (
	headerAuthorization     = "Authorization"
	headerContentEncoding   = "Content-Encoding"
	headerContentLanguage   = "Content-Language"
	headerContentLength     = "Content-Length"
	headerContentMD5        = "Content-MD5"
	headerContentType       = "Content-Type"
	headerIfMatch           = "If-Match"
	headerIfModifiedSince   = "If-Modified-Since"
	headerIfNoneMatch       = "If-None-Match"
	headerIfUnmodifiedSince = "If-Unmodified-Since"
	headerRange             = "Range"
	headerXmsDate           = "x-ms-date"
)

// ComputeHMACSHA256 generates a hash signature for an HTTP request or for a SAS.
func (e *StorageBlobExpander) ComputeHMACSHA256(message string, accountKey string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(accountKey)
	if err != nil {
		return "", fmt.Errorf("Failed to decode storage account key: %s", err)
	}
	h := hmac.New(sha256.New, bytes)
	_, err = h.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("Failed to write bytes: %s", err)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func (e *StorageBlobExpander) buildStringToSign(request *http.Request, accountName string) (string, error) {
	// https://docs.microsoft.com/en-us/rest/api/storageservices/authentication-for-the-azure-storage-services
	headers := request.Header
	contentLength := headers.Get(headerContentLength)
	if contentLength == "0" {
		contentLength = ""
	}

	canonicalizedResource, err := e.buildCanonicalizedResource(request.URL, accountName)
	if err != nil {
		return "", err
	}

	stringToSign := strings.Join([]string{
		request.Method,
		headers.Get(headerContentEncoding),
		headers.Get(headerContentLanguage),
		contentLength,
		headers.Get(headerContentMD5),
		headers.Get(headerContentType),
		"", // Empty date because x-ms-date is expected (as per web page above)
		headers.Get(headerIfModifiedSince),
		headers.Get(headerIfMatch),
		headers.Get(headerIfNoneMatch),
		headers.Get(headerIfUnmodifiedSince),
		headers.Get(headerRange),
		buildCanonicalizedHeader(headers),
		canonicalizedResource,
	}, "\n")
	return stringToSign, nil
}

func buildCanonicalizedHeader(headers http.Header) string {
	cm := map[string][]string{}
	for k, v := range headers {
		headerName := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(headerName, "x-ms-") {
			cm[headerName] = v // NOTE: the value must not have any whitespace around it.
		}
	}
	if len(cm) == 0 {
		return ""
	}

	keys := make([]string, 0, len(cm))
	for key := range cm {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	ch := bytes.NewBufferString("")
	for i, key := range keys {
		if i > 0 {
			ch.WriteRune('\n')
		}
		ch.WriteString(key)
		ch.WriteRune(':')
		ch.WriteString(strings.Join(cm[key], ","))
	}
	return ch.String()
}

func (e *StorageBlobExpander) buildCanonicalizedResource(u *url.URL, accountName string) (string, error) {
	// https://docs.microsoft.com/en-us/rest/api/storageservices/authentication-for-the-azure-storage-services
	cr := bytes.NewBufferString("/")
	cr.WriteString(accountName)

	if len(u.Path) > 0 {
		// Any portion of the CanonicalizedResource string that is derived from
		// the resource's URI should be encoded exactly as it is in the URI.
		// -- https://msdn.microsoft.com/en-gb/library/azure/dd179428.aspx
		cr.WriteString(u.EscapedPath())
	} else {
		// a slash is required to indicate the root path
		cr.WriteString("/")
	}

	// params is a map[string][]string; param name is key; params values is []string
	params, err := url.ParseQuery(u.RawQuery) // Returns URL decoded values
	if err != nil {
		return "", errors.New("parsing query parameters must succeed, otherwise there might be serious problems in the SDK/generated code")
	}

	if len(params) > 0 { // There is at least 1 query parameter
		paramNames := []string{} // We use this to sort the parameter key names
		for paramName := range params {
			paramNames = append(paramNames, paramName) // paramNames must be lowercase
		}
		sort.Strings(paramNames)

		for _, paramName := range paramNames {
			paramValues := params[paramName]
			sort.Strings(paramValues)

			// Join the sorted key values separated by ','
			// Then prepend "keyName:"; then add this string to the buffer
			cr.WriteString("\n" + paramName + ":" + strings.Join(paramValues, ","))
		}
	}
	return cr.String(), nil
}
