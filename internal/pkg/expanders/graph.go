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

	itemTypeMe          = "graph/me"
	itemTypeAppOrSpList = "graph/apporsplist"
	itemTypeAppOrSp     = "graph/apporsp"

	actionUpdateItem = "updateitem"
	actionSearch     = "search"
	actionViewOwners = "owners"
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
	case itemTypeAppOrSpList:
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
	if currentItem.ItemType == itemTypeAppOrSp || currentItem.ItemType == itemTypeAppOrSpList {
		return true, nil
	}

	return false, nil
}

// ListActions returns the actions for an app / sp
func (e *GraphExpander) ListActions(context context.Context, currentItem *TreeNode) ListActionsResult {
	nodes := []*TreeNode{}

	if currentItem.ItemType == itemTypeAppOrSpList {
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    actionSearch,
				Namespace:             namespace,
				Name:                  "Search",
				Display:               "Search",
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
	case actionUpdateItem:
		return e.updateAppOrSp(ctx, currentItem)
	case actionSearch:
		return e.searchAppsOrSps(ctx, currentItem)
	case actionViewOwners:
		return e.showOwners(ctx, currentItem)
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
			ItemType:              itemTypeAppOrSpList,
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
			ItemType:              itemTypeAppOrSpList,
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
		var appsListResponse AppsListResponse
		err = json.Unmarshal([]byte(data), &appsListResponse)
		if err != nil {
			panic(err)
		}

		nodes := []*TreeNode{}
		for _, app := range appsListResponse.Items {

			nodes = append(nodes,
				&TreeNode{
					Parentid:              currentItem.ID,
					ID:                    expandURLRoot + "/" + app.ID,
					Namespace:             namespace,
					Name:                  app.DisplayName,
					Display:               app.DisplayName,
					ExpandURL:             expandURLRoot + "/" + app.ID,
					DeleteURL:             expandURLRoot + "/" + app.ID,
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
		SourceDescription: "GraphExpander failed to list apps or service principals",
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

func (e *GraphExpander) updateAppOrSp(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// get the data for the editor panel
	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.Parent.ExpandURL)
	if err == nil {
		content, err := editor.OpenForContent(data, ".json")
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

		update, err := e.armClient.DoRequest(ctx, "PATCH", currentItem.Parent.ExpandURL)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "GraphAppUpdate request",
				IsPrimaryResponse: true,
			}
		}
		return ExpanderResult{
			Response:          ExpanderResponse{Response: update, ResponseType: interfaces.ResponseJSON},
			SourceDescription: "GraphAppUpdate request",
			Nodes:             []*TreeNode{},
			IsPrimaryResponse: true,
		}
	}

	return ExpanderResult{
		Err:               err,
		SourceDescription: "GraphAppUpdate request",
		IsPrimaryResponse: true,
	}
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
