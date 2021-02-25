package expanders

import (
	// "bytes"

	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"fmt"
	"net/http"
	"net/url"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/stuartleeks/gocui"
)

// NewCosmosDbExpander creates a new instance of CosmosDbExpander
func NewCosmosDbExpander(armclient *armclient.Client, commandPanel interfaces.CommandPanel, gui *gocui.Gui) *CosmosDbExpander {
	return &CosmosDbExpander{
		client:       &http.Client{},
		armClient:    armclient,
		commandPanel: commandPanel,
		gui:          gui,
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

type doRequestResponse struct {
	Data       []byte
	StatusCode int
	Headers    http.Header
}

type CosmosDbQuery struct {
	Query string `json:"query"`
}

const (
	cosmosdbListSQLDocuments             = "sql-listdocs"
	cosmosdbListSQLDocumentsContinuation = "sql-listdocs-continue"
	cosmosdbListSQLQuery                 = "sql-query"
	cosmosdbSQLDocument                  = "sql-document"
)

const (
	cosmosdbActionListKeys    = "list-keys"
	cosmosdbActionGetDocument = "get-document"
	cosmosdbActionSQLQuery    = "sql-query"
)

func (e *CosmosDbExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// CosmosDbExpander expands the blob  data-plane aspects of a Storage Account
type CosmosDbExpander struct {
	ExpanderBase
	client       *http.Client
	armClient    *armclient.Client
	commandPanel interfaces.CommandPanel
	gui          *gocui.Gui
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
	case cosmosdbListSQLDocumentsContinuation:
		return e.expandSQLDocumentsContinuation(ctx, currentItem)
	case cosmosdbSQLDocument:
		return e.expandSQLDocumentNode(ctx, currentItem)
	}

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "CosmosDbExpander request",
	}
}

// CanUpdate indicates if the item can be updated
func (e CosmosDbExpander) CanUpdate(ctx context.Context, item *TreeNode) (bool, error) {
	switch item.ItemType {
	case cosmosdbSQLDocument:
		return true, nil
	default:
		return false, nil
	}
}

// Update updates the item
func (e *CosmosDbExpander) Update(ctx context.Context, item *TreeNode, updatedContent string) error {
	switch item.ItemType {
	case cosmosdbSQLDocument:
		return e.updateSQLDocument(ctx, item, updatedContent)
	default:
		return fmt.Errorf("Unsupported item type: %s", item.ItemType)
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *CosmosDbExpander) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	switch item.ItemType {
	case cosmosdbSQLDocument:
		return e.deleteSQLDocument(ctx, item)
	}
	return false, nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *CosmosDbExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	swaggerResourceType := item.SwaggerResourceType

	for tempItem := item; swaggerResourceType == nil && tempItem.Parent != nil; {
		tempItem = tempItem.Parent
		swaggerResourceType = tempItem.SwaggerResourceType
	}
	if swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}" {
			return true, nil
		}
		if strings.HasPrefix(swaggerResourceType.Endpoint.TemplateURL, "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}") {
			return true, nil
		}
	}
	return false, nil
}

// ListActions returns an action for listing keys on the cosmos db
func (e *CosmosDbExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {

	nodes := []*TreeNode{}

	swaggerResourceType := item.SwaggerResourceType
	for tempItem := item; swaggerResourceType == nil && tempItem.Parent != nil; {
		tempItem = tempItem.Parent
		swaggerResourceType = tempItem.SwaggerResourceType
	}
	if swaggerResourceType != nil {
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
		if strings.HasPrefix(swaggerResourceType.Endpoint.TemplateURL, "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}") {
			nodes = append(nodes,
				&TreeNode{
					Parentid:              item.ID,
					ID:                    item.ID + "?get-document",
					Namespace:             "cosmos-db",
					Name:                  "Get Document",
					Display:               "Get Document",
					ItemType:              ActionType,
					SuppressGenericExpand: true,
					Metadata: map[string]string{
						"ActionID": cosmosdbActionGetDocument,
					},
				},
				&TreeNode{
					Parentid:              item.ID,
					ID:                    item.ID + "?sql-query",
					Namespace:             "cosmos-db",
					Name:                  "Execute Query",
					Display:               "ExecuteQuery",
					ItemType:              ActionType,
					SuppressGenericExpand: true,
					Metadata: map[string]string{
						"ActionID": cosmosdbActionSQLQuery,
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

// ExecuteAction implements action for listing keys
func (e *CosmosDbExpander) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	actionID := item.Metadata["ActionID"]

	switch actionID {
	case cosmosdbActionListKeys:
		return e.cosmosdbActionListKeys(context, item)
	case cosmosdbActionGetDocument:
		return e.cosmosdbActionGetDocument(context, item)
	case cosmosdbActionSQLQuery:
		return e.cosmosdbActionExecuteQuery(context, item)
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
	return e.expandSQLDocumentsCommon(ctx, item, accountName, databaseName, containerName, accountKey, partitionKey, "")
}

func (e *CosmosDbExpander) expandSQLDocumentsContinuation(ctx context.Context, item *TreeNode) ExpanderResult {
	accountName := item.Metadata["AccountName"]
	databaseName := item.Metadata["DatabaseName"]
	containerName := item.Metadata["ContainerName"]
	accountKey := item.Metadata["AccountKey"]
	partitionKey := item.Metadata["PartitionKey"]
	continuationToken := item.Metadata["ContinuationToken"]

	return e.expandSQLDocumentsCommon(ctx, item, accountName, databaseName, containerName, accountKey, partitionKey, continuationToken)
}

func (e *CosmosDbExpander) expandSQLDocumentsCommon(ctx context.Context, item *TreeNode, accountName string, databaseName string, containerName string, accountKey string, partitionKey string, continuationToken string) ExpanderResult {
	partitionKeySegments := strings.Split(partitionKey, "/")

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s/docs", databaseName, containerName)
	headers := map[string]string{}
	if continuationToken != "" {
		headers["x-ms-continuation"] = continuationToken
	}
	response, err := e.doRequestWithHeaders(ctx, "GET", accountName, requestURL, accountKey, headers)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting documents: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return ExpanderResult{
			Err:               fmt.Errorf("Error listing documents. StatusCode=%d, Response=%s", response.StatusCode, data),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	// TODO - parse nodes
	var list CosmosDbListDocumentResponse
	if err = json.Unmarshal(response.Data, &list); err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error unmarshalling response: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	nodes := []*TreeNode{}
	for _, document := range list.Documents {
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
			DeleteURL:             "placeholder",
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

	if continuationToken := response.Headers.Get("x-ms-continuation"); continuationToken != "" {
		node := TreeNode{
			Parentid:      item.ID,
			Namespace:     "cosmosdb",
			ID:            item.ID + "/" + "...more",
			Name:          "more...",
			Display:       "more...",
			ItemType:      cosmosdbListSQLDocumentsContinuation,
			ExpandURL:     ExpandURLNotSupported,
			ExpandInPlace: true,
			Metadata: map[string]string{
				"AccountName":       accountName,
				"DatabaseName":      databaseName,
				"ContainerName":     containerName,
				"AccountKey":        accountKey,
				"PartitionKey":      partitionKey,
				"ContinuationToken": continuationToken,
			},
		}

		nodes = append(nodes, &node)
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: string(response.Data), ResponseType: ResponseJSON},
		Nodes:             nodes,
		IsPrimaryResponse: true,
	}
}

func (e *CosmosDbExpander) expandSQLDocumentNode(ctx context.Context, item *TreeNode) ExpanderResult {
	accountName := item.Metadata["AccountName"]
	databaseName := item.Metadata["DatabaseName"]
	containerName := item.Metadata["ContainerName"]
	accountKey := item.Metadata["AccountKey"]
	partitionKeyValue := item.Metadata["PartitionKeyValue"]

	return e.expandSQLDocument(ctx, accountName, databaseName, containerName, accountKey, partitionKeyValue, item.ID)
}

func (e *CosmosDbExpander) expandSQLDocument(ctx context.Context, accountName string, databaseName string, containerName string, accountKey string, partitionKeyValue string, id string) ExpanderResult {

	headers := map[string]string{}

	if partitionKeyValue != "" {
		headers["x-ms-documentdb-partitionkey"] = partitionKeyValue
	}

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s/docs/%s", databaseName, containerName, id)
	response, err := e.doRequestWithHeaders(ctx, "GET", accountName, requestURL, accountKey, headers)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting documents: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting document. StatusCode=%d, Response=%s", response.StatusCode, data),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: string(response.Data), ResponseType: ResponseJSON},
		IsPrimaryResponse: true,
	}
}

func (e *CosmosDbExpander) updateSQLDocument(ctx context.Context, item *TreeNode, updatedContent string) error {

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
	response, err := e.doRequestWithHeadersAndBody(ctx, "PUT", accountName, requestURL, accountKey, headers, strings.NewReader(updatedContent))
	if err != nil {
		return fmt.Errorf("Error updating documents: %s", err)
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return fmt.Errorf("Error updating document. StatusCode=%d, Response=%s", response.StatusCode, data)
	}

	return nil
}

func (e *CosmosDbExpander) deleteSQLDocument(ctx context.Context, item *TreeNode) (bool, error) {

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
	response, err := e.doRequestWithHeaders(ctx, "DELETE", accountName, requestURL, accountKey, headers)
	if err != nil {
		return false, fmt.Errorf("Error getting documents: %s", err)
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return false, fmt.Errorf("Error getting document. StatusCode=%d, Response=%s", response.StatusCode, data)
	}

	return true, nil
}

func (e *CosmosDbExpander) getCollectionPartitionKey(ctx context.Context, accountName string, databaseName string, containerName string, accountKey string) (string, error) {

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s", databaseName, containerName)
	response, err := e.doRequest(ctx, "GET", accountName, requestURL, accountKey)
	if err != nil {
		return "", fmt.Errorf("Error getting documents: %s", err)
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return "", fmt.Errorf("Error getting collection partition key. StatusCode=%d, Response=%s", response.StatusCode, data)
	}

	var container CosmosDbContainer
	if err = json.Unmarshal(response.Data, &container); err != nil {
		return "", fmt.Errorf("Error unmarshalling response: %s", err)
	}

	if container.PartitionKey == nil {
		return "", nil
	}
	return container.PartitionKey.Paths[0], nil
}

func (e *CosmosDbExpander) cosmosdbActionListKeys(ctx context.Context, item *TreeNode) ExpanderResult {

	i := strings.Index(item.ID, "?")
	baseURL := item.ID[0:i]
	listKeysURL := baseURL + "/listKeys?api-version=2020-04-01"

	data, err := e.armClient.DoRequest(ctx, "POST", listKeysURL)

	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting keys: %s", err),
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

func (e *CosmosDbExpander) cosmosdbActionGetDocument(ctx context.Context, item *TreeNode) ExpanderResult {

	swaggerResourceType := item.Parent.SwaggerResourceType // item is action node, item.Parent is the node the action relates to
	tempItem := item
	for swaggerResourceType == nil && tempItem.Parent != nil {
		tempItem = tempItem.Parent
		swaggerResourceType = tempItem.SwaggerResourceType
	}
	matchResult := swaggerResourceType.Endpoint.Match(tempItem.ID)
	if !matchResult.IsMatch {
		return ExpanderResult{
			Err:               fmt.Errorf("Endpoint should match"),
			SourceDescription: "CosmosDbExpander request",
			IsPrimaryResponse: true,
		}
	}

	accountName := matchResult.Values["accountName"]
	databaseName := matchResult.Values["databaseName"]
	containerName := matchResult.Values["containerName"]

	accountKey, err := e.getAccountKey(ctx, tempItem)
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

	commandChannel := make(chan string, 1)
	commandPanelNotification := func(state interfaces.CommandPanelNotification) {
		if state.EnterPressed {
			commandChannel <- state.CurrentText
			e.commandPanel.Hide()
		}
	}
	partitionKeyValue := ""
	if partitionKey != "" {
		e.commandPanel.ShowWithText("partition key:", "", nil, commandPanelNotification)
		// Force UI to re-render to pickup
		e.gui.Update(func(g *gocui.Gui) error {
			return nil
		})
		partitionKeyValue = <-commandChannel
		partitionKeyValue = fmt.Sprintf("[\"%s\"]", partitionKeyValue)
	}
	e.commandPanel.ShowWithText("id:", "", nil, commandPanelNotification)
	id := <-commandChannel
	_, _ = e.gui.SetCurrentView("listWidget")
	// Force UI to re-render to pickup
	e.gui.Update(func(g *gocui.Gui) error {
		return nil
	})

	response := e.expandSQLDocument(ctx, accountName, databaseName, containerName, accountKey, partitionKeyValue, id)

	if response.Err != nil {
		return response
	}

	node := TreeNode{
		Parentid:              item.ID,
		ID:                    id,
		Namespace:             "cosmosdb",
		Name:                  id,
		Display:               id,
		ItemType:              cosmosdbSQLDocument,
		ExpandURL:             ExpandURLNotSupported,
		DeleteURL:             "placeholder",
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

	return ExpanderResult{
		Response:          ExpanderResponse{Response: "<-- Open the node to view the document :-)", ResponseType: ResponsePlainText},
		Nodes:             []*TreeNode{&node},
		IsPrimaryResponse: true,
		SourceDescription: "CosmosDbExpander request",
	}
}

func (e *CosmosDbExpander) cosmosdbActionExecuteQuery(ctx context.Context, item *TreeNode) ExpanderResult {

	swaggerResourceType := item.Parent.SwaggerResourceType // item is action node, item.Parent is the node the action relates to
	tempItem := item
	for swaggerResourceType == nil && tempItem.Parent != nil {
		tempItem = tempItem.Parent
		swaggerResourceType = tempItem.SwaggerResourceType
	}
	matchResult := swaggerResourceType.Endpoint.Match(tempItem.ID)
	if !matchResult.IsMatch {
		return ExpanderResult{
			Err:               fmt.Errorf("Endpoint should match"),
			SourceDescription: "CosmosDbExpander request",
			IsPrimaryResponse: true,
		}
	}

	accountName := matchResult.Values["accountName"]
	databaseName := matchResult.Values["databaseName"]
	containerName := matchResult.Values["containerName"]

	accountKey, err := e.getAccountKey(ctx, tempItem)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting account key: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	commandChannel := make(chan string, 1)
	commandPanelNotification := func(state interfaces.CommandPanelNotification) {
		if state.EnterPressed {
			commandChannel <- state.CurrentText
			e.commandPanel.Hide()
		}
	}
	queryText := ""
	e.commandPanel.ShowWithText("query:", "SELECT * FROM c", nil, commandPanelNotification)
	// Force UI to re-render to pickup
	e.gui.Update(func(g *gocui.Gui) error {
		return nil
	})
	queryText = <-commandChannel
	_, _ = e.gui.SetCurrentView("listWidget")
	// Force UI to re-render to pickup
	e.gui.Update(func(g *gocui.Gui) error {
		return nil
	})

	headers := map[string]string{}

	headers["x-ms-documentdb-isquery"] = "true"
	headers["Content-Type"] = "application/query+json"
	headers["x-ms-documentdb-query-enablecrosspartition"] = "true"

	query := CosmosDbQuery{
		Query: queryText,
	}
	buf, err := json.Marshal(query)
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error marshalling query as JSON: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	requestURL := fmt.Sprintf("/dbs/%s/colls/%s/docs", databaseName, containerName)
	response, err := e.doRequestWithHeadersAndBody(ctx, "POST", accountName, requestURL, accountKey, headers, bytes.NewBuffer(buf))
	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting documents: %s", err),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}
	if !e.isSuccessCode(response.StatusCode) {
		data := ""
		if response != nil {
			data = string(response.Data)
		}
		return ExpanderResult{
			Err:               fmt.Errorf("Error getting document. StatusCode=%d, Response=%s", response.StatusCode, data),
			IsPrimaryResponse: true,
			SourceDescription: "CosmosDbExpander request",
		}
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: string(response.Data), ResponseType: ResponseJSON},
		Nodes:             []*TreeNode{},
		IsPrimaryResponse: true,
		SourceDescription: "CosmosDbExpander request",
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

func (e *CosmosDbExpander) doRequest(ctx context.Context, verb string, accountName string, requestURL string, accountKey string) (*doRequestResponse, error) {
	return e.doRequestWithHeaders(ctx, verb, accountName, requestURL, accountKey, map[string]string{})
}
func (e *CosmosDbExpander) doRequestWithHeaders(ctx context.Context, verb string, accountName string, requestURL string, accountKey string, headers map[string]string) (*doRequestResponse, error) {
	return e.doRequestWithHeadersAndBody(ctx, verb, accountName, requestURL, accountKey, headers, nil)
}
func (e *CosmosDbExpander) doRequestWithHeadersAndBody(ctx context.Context, verb string, accountName string, requestURL string, accountKey string, headers map[string]string, body io.Reader) (*doRequestResponse, error) {

	span, _ := tracing.StartSpanFromContext(ctx, "doRequest(cosmosexpander):"+requestURL, tracing.SetTag("url", requestURL))
	defer span.Finish()

	if requestURL[0] == '/' {
		requestURL = requestURL[1:]
	}

	fullURL := fmt.Sprintf("https://%s.documents.azure.com/%s", accountName, requestURL)

	req, err := http.NewRequestWithContext(ctx, verb, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
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
		return nil, err
	}

	header := url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig)
	req.Header.Add("Authorization", header)

	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("Request failed: %s", err)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read body: %s", err)
	}

	return &doRequestResponse{
		Data:       buf,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
	}, nil
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
		return "", err
	}

	hmac := hmac.New(sha256.New, salt)
	_, err = hmac.Write([]byte(value))
	if err != nil {
		return "", err
	}
	b := hmac.Sum(nil)

	ret = enc.EncodeToString(b)
	return ret, nil
}
