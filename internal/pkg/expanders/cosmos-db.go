package expanders

import (
	// "bytes"

	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"fmt"
	"net/http"
	"net/url"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// NewCosmosDbExpander creates a new instance of CosmosDbExpander
func NewCosmosDbExpander(armclient *armclient.Client) *CosmosDbExpander {
	return &CosmosDbExpander{
		client:    &http.Client{},
		armClient: armclient,
	}
}

// Check interface
var _ Expander = &CosmosDbExpander{}

// CosmosDbListKeyResponse is used to unmarshal a call to listKeys on a storage account
type CosmosDbListKeyResponse struct {
	PrimaryMasterKey string `json:"primaryMasterKey"`
	// other properties not needed
}

// CosmosDbListDocumentResponse is used to unmarshal a ListDocuments response
type CosmosDbListDocumentResponse struct {
	Documents []interface{} `json:"Documents"`
	Count     int           `json:"count"`
}

// CosmosDbContainer is used to unmarshal a GetContainer response
type CosmosDbContainer struct {
	PartitionKey *struct {
		Paths []string `json:"Paths"`
		Kind  string   `json:"kind"`
	} `json:"partitionKey"`
}

const (
	cosmosdbListSQLDocuments = "sql-listdocs"
	cosmosdbSQLDocument      = "sql-document"
)

const (
	cosmosdbActionListKeys = "list-keys"
)

func (e *CosmosDbExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// CosmosDbExpander expands the blob  data-plane aspects of a Storage Account
type CosmosDbExpander struct {
	ExpanderBase
	client    *http.Client
	armClient *armclient.Client
}

// Name returns the name of the expander
func (e *CosmosDbExpander) Name() string {
	return "CosmosDbExpander"
}

// DoesExpand checks if this is a Cosmos DB account
func (e *CosmosDbExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == SubResourceType && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "cosmosdb" {
		return true, nil
	}
	return false, nil
}

// Expand returns items in the Cosmos DB account
func (e *CosmosDbExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	newItems := []*TreeNode{}

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "cosmosdb" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}" {
		matchResult := swaggerResourceType.Endpoint.Match(currentItem.ID)
		if matchResult.IsMatch {
			newItems = append(newItems, &TreeNode{
				Parentid:              currentItem.ID,
				ID:                    currentItem.ID + "/<documents>",
				Namespace:             "cosmosdb",
				Name:                  "Documents",
				Display:               "Documents",
				ItemType:              cosmosdbListSQLDocuments,
				ExpandURL:             ExpandURLNotSupported,
				SuppressSwaggerExpand: true,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"AccountName":   matchResult.Values["accountName"],
					"DatabaseName":  matchResult.Values["databaseName"],
					"ContainerName": matchResult.Values["containerName"],
				},
			})
		}

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "CosmosDbExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	switch currentItem.ItemType {
	case cosmosdbListSQLDocuments:
		return e.expandSQLDocuments(ctx, currentItem)
	case cosmosdbSQLDocument:
		return e.expandSQLDocument(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "CosmosDbExpander request",
	}
}

func (e *CosmosDbExpander) expandSQLDocuments(ctx context.Context, item *TreeNode) ExpanderResult {

	accountName := item.Metadata["AccountName"]
	databaseName := item.Metadata["DatabaseName"]
	containerName := item.Metadata["ContainerName"]

	accountKey, err := e.getAccountKey(ctx, item)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting account key: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	partitionKey, err := e.getCollectionPartitionKey(ctx, accountName, databaseName, containerName, accountKey)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting partition key: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if partitionKey != "" && partitionKey[0] == '/' {
		partitionKey = partitionKey[1:]
	}
	partitionKeySegments := strings.Split(partitionKey, "/")

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s/docs", databaseName, containerName)
	data, statusCode, err := e.doRequest(ctx, "GET", accountName, requestURL, accountKey)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting documents: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if !e.isSuccessCode(statusCode) {
		return ExpanderResult{
			Err:               fmt.Errorf("Error listing documents. StatusCode=%d, Response=%s", statusCode, string(data)),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	// TODO - parse nodes
	var response CosmosDbListDocumentResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error unmarshalling response: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	nodes := []*TreeNode{}
	for _, document := range response.Documents {
		document := document.(map[string]interface{})
		partitionKeyValue := ""
		if partitionKey != "" {
			// get the partitionKey value for the current document
			v, err := getJSONProperty(document, partitionKeySegments...)
			if err != nil {
				return ExpanderResult{
					Err:               fmt.Errorf("Error determining partition key value: %s", err),
					IsPrimaryResponse: true,
					SourceDescription: "CosmosDbExpander request",
				}
			}
			partitionKeyValue = fmt.Sprintf("[\"%v\"]", v)
		}
		id := document["id"].(string)
		node := TreeNode{
			Parentid:              item.ID,
			ID:                    id,
			Namespace:             "cosmosdb",
			Name:                  id,
			Display:               id,
			ItemType:              cosmosdbSQLDocument,
			ExpandURL:             ExpandURLNotSupported,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"AccountName":       accountName,
				"DatabaseName":      databaseName,
				"ContainerName":     containerName,
				"AccountKey":        accountKey,
				"PartitionKeyValue": partitionKeyValue,
			},
		}
		nodes = append(nodes, &node)
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		Nodes:             nodes,
		IsPrimaryResponse: true,
	}
}

func (e *CosmosDbExpander) expandSQLDocument(ctx context.Context, item *TreeNode) ExpanderResult {

	accountName := item.Metadata["AccountName"]
	databaseName := item.Metadata["DatabaseName"]
	containerName := item.Metadata["ContainerName"]
	accountKey := item.Metadata["AccountKey"]
	partitionKeyValue := item.Metadata["PartitionKeyValue"]

	headers := map[string]string{}

	if partitionKeyValue != "" {
		headers["x-ms-documentdb-partitionkey"] = partitionKeyValue
	}

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s/docs/%s", databaseName, containerName, item.ID)
	data, statusCode, err := e.doRequestWithHeaders(ctx, "GET", accountName, requestURL, accountKey, headers)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting documents: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if !e.isSuccessCode(statusCode) {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting document. StatusCode=%d, Response=%s", statusCode, string(data)),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	// TODO - parse nodes
	var response CosmosDbListDocumentResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error unmarshalling response: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		IsPrimaryResponse: true,
	}
}

func (e *CosmosDbExpander) getCollectionPartitionKey(ctx context.Context, accountName string, databaseName string, containerName string, accountKey string) (string, error) {

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s", databaseName, containerName)
	data, statusCode, err := e.doRequest(ctx, "GET", accountName, requestURL, accountKey)
	if err != nil {
		return "", fmt.Errorf("Error getting documents: %s", err)
	}
	if !e.isSuccessCode(statusCode) {
		return "", fmt.Errorf("Error getting collection partition key. StatusCode=%d, Response=%s", statusCode, string(data))
	}

	var response CosmosDbContainer
	if err = json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("Error unmarshalling response: %s", err)
	}

	if response.PartitionKey == nil {
		return "", nil
	}
	return response.PartitionKey.Paths[0], nil
}

// // Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
// func (e *CosmosDbExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {
// 	switch currentItem.ItemType {
// 	case storageBlobNodeBlob, storageBlobNodeBlobMetadata:
// 		return e.deleteBlob(ctx, currentItem)
// 	}
// 	return false, nil
// }

// HasActions is a default implementation returning false to indicate no actions available
func (e *CosmosDbExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	swaggerResourceType := item.SwaggerResourceType
	if item.ItemType == ResourceType && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}" {
			return true, nil
		}
	}

	// switch item.ItemType {
	// case storageBlobNodeBlob,
	// 	storageBlobNodeBlobMetadata:
	// 	return true, nil
	// }
	return false, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *CosmosDbExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {

	nodes := []*TreeNode{}

	swaggerResourceType := item.SwaggerResourceType
	if item.ItemType == ResourceType && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}" {
			nodes = append(nodes,
				&TreeNode{
					Parentid:              item.ID,
					ID:                    item.ID + "?list-keys",
					Namespace:             "cosmos-db",
					Name:                  "List Keys",
					Display:               "List Keys",
					ItemType:              ActionType,
					SuppressGenericExpand: true,
					Metadata: map[string]string{
						"ActionID": cosmosdbActionListKeys,
					},
				})
		}
	}

	return ListActionsResult{
		Nodes:             nodes,
		SourceDescription: "CosmosDbExpander",
		IsPrimaryResponse: true,
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *CosmosDbExpander) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	actionID := item.Metadata["ActionID"]

	switch actionID {
	case cosmosdbActionListKeys:
		return e.cosmosdbActionListKeys(context, item)
	case "":
		return ExpanderResult{
			SourceDescription: "CosmosDbExpander",
			IsPrimaryResponse: true,
			Err:               fmt.Errorf("ActionID metadata not set: %q", item.ID),
		}
	default:
		return ExpanderResult{
			SourceDescription: "CosmosDbExpander",
			IsPrimaryResponse: true,
			Err:               fmt.Errorf("Unhandled ActionID: %q", actionID),
		}
	}
}

func (e *CosmosDbExpander) cosmosdbActionListKeys(ctx context.Context, item *TreeNode) ExpanderResult {

	i := strings.Index(item.ID, "?")
	baseURL := item.ID[0:i]
	listKeysURL := baseURL + "/listKeys?api-version=2020-04-01"

	data, err := e.armClient.DoRequest(ctx, "POST", listKeysURL)

	if err != nil {
		return ExpanderResult{
			Response: ExpanderResponse{
				ResponseType: ResponsePlainText,
				Response:     fmt.Sprintf("Error getting keys: %s", err),
			},
			SourceDescription: "CosmosDbExpander request",
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		Response: ExpanderResponse{
			ResponseType: ResponseJSON,
			Response:     data,
		},
		SourceDescription: "CosmosDbExpander request",
		IsPrimaryResponse: true,
	}
}

func (e *CosmosDbExpander) getAccountKey(ctx context.Context, item *TreeNode) (string, error) {
	parts := strings.Split(item.ID, "/")
	if len(parts) > 9 {
		// Want "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}"
		parts = parts[:9]
	}
	baseURL := strings.Join(parts, "/")
	listKeysURL := baseURL + "/listKeys?api-version=2020-04-01"
	data, err := e.armClient.DoRequest(ctx, "POST", listKeysURL)
	if err != nil {
		return "", err
	}

	var result CosmosDbListKeyResponse
	if err = json.Unmarshal([]byte(data), &result); err != nil {
		return "", err
	}

	return result.PrimaryMasterKey, nil
}

func (e *CosmosDbExpander) isSuccessCode(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func (e *CosmosDbExpander) doRequest(ctx context.Context, verb string, accountName string, requestURL string, accountKey string) ([]byte, int, error) {
	return e.doRequestWithHeaders(ctx, verb, accountName, requestURL, accountKey, map[string]string{})
}
func (e *CosmosDbExpander) doRequestWithHeaders(ctx context.Context, verb string, accountName string, requestURL string, accountKey string, headers map[string]string) ([]byte, int, error) {

	span, _ := tracing.StartSpanFromContext(ctx, "doRequest(cosmosexpander):"+requestURL, tracing.SetTag("url", requestURL))
	defer span.Finish()

	if requestURL[0] == '/' {
		requestURL = requestURL[1:]
	}

	fullURL := fmt.Sprintf("https://%s.documents.azure.com/%s", accountName, requestURL)

	req, err := http.NewRequestWithContext(ctx, verb, fullURL, nil)
	if err != nil {
		return []byte{}, -1, fmt.Errorf("Failed to create request: %s", err)
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	dateString := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Set("x-ms-date", dateString)
	req.Header.Set("x-ms-version", "2018-12-31")

	resourceID, resourceType := parseResource(requestURL)
	masterToken := "master"
	tokenVersion := "1.0"

	// https://docs.microsoft.com/en-us/rest/api/cosmos-db/access-control-on-cosmosdb-resources#constructkeytoken
	parts := []string{strings.ToLower(verb), resourceType, resourceID, strings.ToLower(req.Header.Get("x-ms-date")), "", ""}
	stringToSign := strings.Join(parts, "\n")

	sig, err := signString(stringToSign, accountKey)
	if err != nil {
		return []byte{}, -1, err
	}

	header := url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig)
	req.Header.Add("Authorization", header)

	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, -1, fmt.Errorf("Request failed: %s", err)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, -1, fmt.Errorf("Failed to read body: %s", err)
	}

	return buf, response.StatusCode, nil
}

func parseResource(resourcePath string) (resourceID string, resourceType string) {

	// for dbs/mydb/colls/mycoll/docs return docs, dbs/mydb/colls/mycoll
	// for dbs/mydb/colls/mycoll return colls, dbs/mydb/colls/mycoll

	if resourcePath[len(resourcePath)-1] == '/' {
		resourcePath = resourcePath[:len(resourcePath)-1]
	}
	if resourcePath[0] == '/' {
		resourcePath = resourcePath[1:]
	}

	parts := strings.Split(resourcePath, "/")
	partsLength := len(parts)

	if partsLength%2 == 0 {
		resourceID = resourcePath
		resourceType = parts[partsLength-2]
	} else {
		resourceType = parts[partsLength-1]
		resourceID = strings.Join(parts[0:partsLength-1], "/")
	}

	return resourceID, resourceType
}

func signString(value, key string) (string, error) {
	var ret string
	enc := base64.StdEncoding

	salt, err := enc.DecodeString(key)
	if err != nil {
		return ret, err
	}

	hmac := hmac.New(sha256.New, salt)
	hmac.Write([]byte(value))
	b := hmac.Sum(nil)

	ret = enc.EncodeToString(b)
	return ret, nil
}
