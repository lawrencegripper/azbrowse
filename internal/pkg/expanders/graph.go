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
	namespace           = "graph"
	itemTypeMe          = "graph/me"
	itemTypeAppOrSpList = "graph/apporsplist"
	itemTypeAppOrSp     = "graph/apporsp"

	searchAppsOrSps = "apporsp/search"
)

type MeResponse struct {
	Mail        string `json:"mail"`
	DisplayName string `json:"displayName"`
}

type AppsResponse struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

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
	if currentItem.Namespace == "graph" {
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
	}

	return e.listRootMenu(ctx, currentItem)

	// if currentItem.Name == "Service Principals" {
	// 	// get the list of sp's
	// 	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
	// 	if err == nil {
	// 		var spListResponse AppsListResponse
	// 		err = json.Unmarshal([]byte(data), &spListResponse)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		nodes := []*TreeNode{}
	// 		for _, servicePrincipal := range spListResponse.Items {

	// 			nodes = append(nodes,
	// 				&TreeNode{
	// 					Parentid:              currentItem.ID,
	// 					ID:                    servicePrincipal.Id,
	// 					Namespace:             "graph",
	// 					Name:                  servicePrincipal.DisplayName,
	// 					Display:               servicePrincipal.DisplayName,
	// 					ExpandURL:             "/servicePrincipals/" + servicePrincipal.Id,
	// 					ItemType:              "app",
	// 					SuppressGenericExpand: true,
	// 				})
	// 		}

	// 		return ExpanderResult{
	// 			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
	// 			SourceDescription: "GraphSPListExpander request",
	// 			Nodes:             nodes,
	// 			IsPrimaryResponse: true,
	// 		}
	// 	}
	// }

	// if currentItem.Parentid == "graph/servicePrincipals" {
	// 	// get the list of sps
	// 	data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
	// 	if err == nil {
	// 		return ExpanderResult{
	// 			Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
	// 			SourceDescription: "GraphAppExpander request",
	// 			Nodes:             []*TreeNode{},
	// 			IsPrimaryResponse: true,
	// 		}
	// 	}
	//}
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
	if currentItem.ItemType == itemTypeAppOrSp {
		return true, nil
	}

	return false, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *GraphExpander) ListActions(context context.Context, currentItem *TreeNode) ListActionsResult {
	nodes := []*TreeNode{}

	if currentItem.ItemType == itemTypeAppOrSp {
		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    searchAppsOrSps,
				Namespace:             namespace,
				Name:                  "Search",
				Display:               "Search",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
			})

		nodes = append(nodes,
			&TreeNode{
				Parentid:              currentItem.ID,
				ID:                    currentItem.ItemType + "?update",
				Namespace:             namespace,
				Name:                  "Update Item",
				Display:               "Update Item",
				ItemType:              ActionType,
				SuppressGenericExpand: true,
				Metadata: map[string]string{
					"ActionID": "UpdateAppOrSp",
				},
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

	if currentItem.Name == "Update Item" {
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
	}

	if currentItem.Name == "Search" {
		commandChannel := make(chan string, 1)
		commandPanelNotification := func(state interfaces.CommandPanelNotification) {
			if state.EnterPressed {
				commandChannel <- state.CurrentText
				e.commandPanel.Hide()
			}
		}

		e.commandPanel.ShowWithText("App Name:", "", nil, commandPanelNotification)

		queryText := <-commandChannel

		// Force UI to re-render to pickup
		e.gui.Update(func(g *gocui.Gui) error {
			return nil
		})

		return e.listAppsOrSps(ctx, currentItem, queryText)
	}

	return ExpanderResult{
		SourceDescription: "GraphExpander",
		Err:               fmt.Errorf("Unhandled ActionID"),
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

func (e *GraphExpander) listAppsOrSps(ctx context.Context, currentItem *TreeNode, queryText string) ExpanderResult {
	// get the list of apps
	url := currentItem.ExpandURL
	expandURLRoot := currentItem.ID

	if queryText != "" { // are we in search mode - only available in the app / sp list view
		expandURLRoot = currentItem.Parent.Parent.ID
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
					ID:                    app.Id,
					Namespace:             "graph",
					Name:                  app.DisplayName,
					Display:               app.DisplayName,
					ExpandURL:             expandURLRoot + "/" + app.Id,
					DeleteURL:             expandURLRoot + "/" + app.Id,
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
