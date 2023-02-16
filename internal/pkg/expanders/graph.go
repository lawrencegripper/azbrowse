package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/editor"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// NewGraphExpander creates a new instance of GraphExpander
func NewGraphExpander(armclient *armclient.Client, gui *gocui.Gui, commandPanel interfaces.CommandPanel, contentPanel interfaces.ItemWidget) *GraphExpander {
	return &GraphExpander{
		client:       &http.Client{},
		armClient:    armclient,
		gui:          gui,
		commandPanel: commandPanel,
		contentPanel: contentPanel,
	}
}

// Check interface
var _ Expander = &GraphExpander{}

func (e *GraphExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

const (
	namespace = "graph"

	itemTypeMe      = "graph/me"
	itemTypeAppList = "/applications"
	itemTypeSpList  = "/serviceprincipals"
	itemTypeAppOrSp = "graph/apporsp"

	actionUpdateItem      = "updateitem"
	actionSearch          = "search"
	actionViewOwners      = "owners"
	actionListMyApps      = "listmyapps"
	actionListDeletedApps = "deletedapps"
	actionGetAppByID      = "getappbyid"
	actionNewApp          = "newapp"
)

// MeResponse used for the /me response
type MeResponse struct {
	Mail        string `json:"mail"`
	DisplayName string `json:"displayName"`
}

// AppsResponse used for the Apps / Service Principal list
type AppsResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// AppsListResponse is a wrapper for AppsResponse
type AppsListResponse struct {
	Items []AppsResponse `json:"value"`
}

// GraphExpander expands the Graph resource
type GraphExpander struct {
	ExpanderBase
	client       *http.Client
	armClient    *armclient.Client
	commandPanel interfaces.CommandPanel
	contentPanel interfaces.ItemWidget
	gui          *gocui.Gui
}

// Name returns the name of the expander
func (e *GraphExpander) Name() string {
	return "GraphExpander"
}

// DoesExpand checks if this node can/should expand
func (e *GraphExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.Namespace == namespace {
		return true, nil
	}
	return false, nil
}

// Expand returns items from Graph
func (e *GraphExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	switch currentItem.ItemType {
	case itemTypeMe:
		return e.showMe(ctx, currentItem)
	case itemTypeSpList:
		return e.listAppsOrSps(ctx, currentItem, "")
	case itemTypeAppList:
		return e.listAppsOrSps(ctx, currentItem, "")
	case itemTypeAppOrSp:
		return e.showAppOrSp(ctx, currentItem)
	case "graph":
		return e.listRootMenu(ctx, currentItem)
	case "action":
		return ExpanderResult{
			SourceDescription: "GraphExpander",
			Err:               nil,
		}
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander",
		Err:               fmt.Errorf("Unhandled Graph Expander - " + currentItem.ItemType),
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *GraphExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {

	if currentItem.ItemType == itemTypeAppOrSp {
		_, err := e.armClient.DoRequest(ctx, "DELETE", currentItem.DeleteURL)
		if err == nil {
			return true, nil
		}

		return false, err
	}

	return false, nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *GraphExpander) HasActions(context context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == itemTypeAppOrSp || currentItem.ItemType == itemTypeAppList || currentItem.ItemType == itemTypeSpList {
		return true, nil
	}

	return false, nil
}

// ListActions returns the actions for an app / sp
func (e *GraphExpander) ListActions(context context.Context, currentItem *TreeNode) ListActionsResult {
	nodes := []*TreeNode{}

	if currentItem.ItemType == itemTypeAppList {
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionListMyApps,
				Namespace:             namespace,
				Name:                  "Owned Apps",
				Display:               "Owned Apps",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionListDeletedApps,
				Namespace:             namespace,
				Name:                  "Deleted Apps",
				Display:               "Deleted Apps",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionNewApp,
				Namespace:             namespace,
				Name:                  "New App",
				Display:               "New App",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})
	}

	if currentItem.ItemType == itemTypeAppList || currentItem.ItemType == itemTypeSpList {
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionSearch,
				Namespace:             namespace,
				Name:                  "Search By Name",
				Display:               "Search By Name",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})

		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionGetAppByID,
				Namespace:             namespace,
				Name:                  "Get by Object ID",
				Display:               "Get by Object ID",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})
	}

	if currentItem.ItemType == itemTypeAppOrSp {
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionUpdateItem,
				Namespace:             namespace,
				Name:                  "Update Item",
				Display:               "Update Item",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})

		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionViewOwners,
				Namespace:             namespace,
				Name:                  "View Owners",
				Display:               "View Owners",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})
	}

	return ListActionsResult{
		Nodes:             nodes,
		SourceDescription: "GraphActionsExpander",
		IsPrimaryResponse: true,
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *GraphExpander) ExecuteAction(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	switch currentItem.ID {
	case actionNewApp:
		return e.newApp(ctx, currentItem)
	case actionUpdateItem:
		return e.updateAppOrSp(ctx, currentItem)
	case actionSearch:
		return e.searchAppsOrSps(ctx, currentItem)
	case actionGetAppByID:
		return e.getAppByID(ctx, currentItem)
	case actionViewOwners:
		return e.showOwners(ctx, currentItem)
	case actionListMyApps:
		return e.listFilteredApps(ctx, currentItem, "/myorganization/me/ownedObjects/$/Microsoft.Graph.Application")
	case actionListDeletedApps:
		return e.listFilteredApps(ctx, currentItem, "/directory/deleteditems/Microsoft.Graph.Application")

	}

	return ExpanderResult{
		SourceDescription: "GraphExpander",
		Err:               fmt.Errorf("Unhandled ActionID - " + currentItem.ID),
	}
}

func (e *GraphExpander) listRootMenu(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	baseItems := []*TreeNode{
		{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/me",
			Namespace:             namespace,
			Name:                  "Me",
			Display:               "Me",
			ItemType:              itemTypeMe,
			ExpandURL:             "/me",
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
		{
			Parentid:              currentItem.ID,
			ID:                    "/applications",
			Namespace:             namespace,
			Name:                  "Apps",
			Display:               "Apps",
			ItemType:              itemTypeAppList,
			ExpandURL:             "/applications",
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
		{
			Parentid:              currentItem.ID,
			ID:                    "/serviceprincipals",
			Namespace:             namespace,
			Name:                  "Service Principals",
			Display:               "Service Principals",
			ItemType:              itemTypeSpList,
			ExpandURL:             "/servicePrincipals", //&$search="displayName:damoo"
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
	}

	return ExpanderResult{
		Err:               nil,
		SourceDescription: "GraphExpander request",
		Nodes:             baseItems,
		IsPrimaryResponse: false,
	}
}

func (e *GraphExpander) showMe(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// make the graph call
	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
	if err == nil {
		var meResponse MeResponse
		jsonErr := json.Unmarshal([]byte(data), &meResponse)
		if jsonErr == nil {
			return ExpanderResult{
				Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
				SourceDescription: "GraphMeExpander request",
				Nodes:             []*TreeNode{},
				IsPrimaryResponse: true,
			}
		}
		return ExpanderResult{
			SourceDescription: "GraphExpander failed to map ME result to object",
			Err:               jsonErr,
		}
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to show ME",
		Err:               err,
	}
}

func (e *GraphExpander) listAppsOrSps(ctx context.Context, currentItem *TreeNode, queryText string) ExpanderResult {
	// get the list of apps
	url := currentItem.ExpandURL
	expandURLRoot := currentItem.ID

	if queryText != "" { // are we in search mode?
		expandURLRoot = currentItem.Parent.ID
		url = expandURLRoot + "?$search=\"displayName:" + queryText + "\""
	}

	data, err := e.armClient.DoRequest(ctx, "GET", url)
	if err == nil {
		return makeAppsList(data, expandURLRoot, currentItem)
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to list apps or service principals",
		Err:               err,
	}
}

func (e *GraphExpander) listFilteredApps(ctx context.Context, currentItem *TreeNode, url string) ExpanderResult {
	// get the list of apps
	expandURLRoot := currentItem.Parent.ID

	data, err := e.armClient.DoRequest(ctx, "GET", url)
	if err == nil {
		return makeAppsList(data, expandURLRoot, currentItem)
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to list apps owned by ME",
		Err:               err,
	}
}

func (e *GraphExpander) getAppByID(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	commandChannel := make(chan string, 1)
	commandPanelNotification := func(state interfaces.CommandPanelNotification) {
		if state.EnterPressed {
			commandChannel <- state.CurrentText
			e.commandPanel.Hide()
		}
	}

	e.commandPanel.ShowWithText("App ID:", "", nil, commandPanelNotification)

	appID := <-commandChannel

	// get the list of apps
	expandURLRoot := currentItem.Parent.ID

	data, err := e.armClient.DoRequest(ctx, "GET", expandURLRoot+"/"+appID)
	if err == nil {
		return makeSingleApp(data, expandURLRoot, currentItem)
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to get app by ID",
		Err:               err,
	}
}

func makeSingleApp(data string, urlRoot string, currentItem *TreeNode) ExpanderResult {
	var appsResponse AppsResponse
	err := json.Unmarshal([]byte(data), &appsResponse)
	if err == nil {
		nodes := []*TreeNode{}
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    urlRoot + "/" + appsResponse.ID,
				Namespace:             namespace,
				Name:                  appsResponse.DisplayName,
				Display:               appsResponse.DisplayName,
				ExpandURL:             urlRoot + "/" + appsResponse.ID,
				DeleteURL:             urlRoot + "/" + appsResponse.ID,
				ItemType:              itemTypeAppOrSp,
				SuppressGenericExpand: true,
			})

		return ExpanderResult{
			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
			SourceDescription: "GraphAppsExpander request",
			Nodes:             nodes,
			IsPrimaryResponse: true,
		}
	}
	return ExpanderResult{
		SourceDescription: "GraphExpander failed to deserialize app result",
		Err:               err,
	}
}

func makeAppsList(data string, urlRoot string, currentItem *TreeNode) ExpanderResult {
	var appsListResponse AppsListResponse
	err := json.Unmarshal([]byte(data), &appsListResponse)
	if err == nil {
		nodes := []*TreeNode{}
		for _, app := range appsListResponse.Items {

			nodes = append(nodes,
				&TreeNode{
					Parentid:              currentItem.ID,
					ID:                    urlRoot + "/" + app.ID,
					Namespace:             namespace,
					Name:                  app.DisplayName,
					Display:               app.DisplayName,
					ExpandURL:             urlRoot + "/" + app.ID,
					DeleteURL:             urlRoot + "/" + app.ID,
					ItemType:              itemTypeAppOrSp,
					SuppressGenericExpand: true,
				})
		}

		return ExpanderResult{
			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
			SourceDescription: "GraphAppsExpander request",
			Nodes:             nodes,
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to deserialise apps list response",
		Err:               err,
	}
}

func (e *GraphExpander) showAppOrSp(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	// get the specific app / SP
	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
	if err == nil {
		return ExpanderResult{
			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
			SourceDescription: "GraphAppExpander request",
			Nodes:             []*TreeNode{},
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to show app or SP",
		Err:               err,
	}
}

func (e *GraphExpander) newApp(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// get the data for the editor panel
	data := "{ \"displayName\": \"Replace this\" }"
	content, err := editor.OpenForContent(data, ".json")
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "GraphExpander New App request",
			IsPrimaryResponse: true,
		}
	}
	if content == data || strings.TrimSpace(content) == "" {
		return ExpanderResult{
			Err:               fmt.Errorf("User canceled"),
			SourceDescription: "GraphExpander New App request",
			IsPrimaryResponse: true,
		}
	}

	update, err := e.armClient.DoRequestWithBody(ctx, "POST", currentItem.Parent.ID, content)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "GraphExpander New App request",
			IsPrimaryResponse: true,
		}
	}
	return makeSingleApp(update, currentItem.Parent.ID, currentItem)
}

func (e *GraphExpander) updateAppOrSp(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// get the data for the editor panel
	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.Parent.ExpandURL)
	if err == nil {
		// trim appId and publisherDomain out to enable the update (PATCH)
		inContent := removeString(data, "\"appId\":", ",")
		inContent = removeString(inContent, "\"publisherDomain\":", ",")

		content, err := editor.OpenForContent(inContent, ".json")
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "GraphAppUpdate request",
				IsPrimaryResponse: true,
			}
		}

		if content == data || strings.TrimSpace(content) == "" {
			return ExpanderResult{
				Err:               fmt.Errorf("User canceled"),
				SourceDescription: "GraphAppUpdate request",
				IsPrimaryResponse: true,
			}
		}

		_, err = e.armClient.DoRequestWithBody(ctx, "PATCH", currentItem.Parent.ExpandURL, content)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "GraphAppUpdate request",
				IsPrimaryResponse: true,
			}
		}
		return makeSingleApp(content, itemTypeAppList, currentItem)
	}

	return ExpanderResult{
		Err:               err,
		SourceDescription: "GraphAppUpdate request",
		IsPrimaryResponse: true,
	}
}

func removeString(content string, start string, end string) string {
	beforeIndex := strings.Index(content, start)
	before := content[:beforeIndex]
	afterAppIndex := strings.Index(content[beforeIndex+1:], end)
	after := content[beforeIndex+afterAppIndex+len(end)+1:]
	return before + after
}

func (e *GraphExpander) searchAppsOrSps(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	commandChannel := make(chan string, 1)
	commandPanelNotification := func(state interfaces.CommandPanelNotification) {
		if state.EnterPressed {
			commandChannel <- state.CurrentText
			e.commandPanel.Hide()
		}
	}

	e.commandPanel.ShowWithText("App Name:", "", nil, commandPanelNotification)

	queryText := <-commandChannel

	return e.listAppsOrSps(ctx, currentItem, queryText)
}

func (e *GraphExpander) showOwners(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.Parent.ExpandURL+"/owners")
	if err == nil {
		return ExpanderResult{
			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
			SourceDescription: "GraphAppExpander request",
			Nodes:             []*TreeNode{},
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander failed to show the owners",
		Err:               err,
	}
}
