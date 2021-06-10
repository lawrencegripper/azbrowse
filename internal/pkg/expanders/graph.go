package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// GraphExpander creates a new instance of GraphExpander
func NewGraphExpander(armclient *armclient.Client) *GraphExpander {
	return &GraphExpander{
		client:    &http.Client{},
		armClient: armclient,
	}
}

// Check interface
var _ Expander = &GraphExpander{}

func (e *GraphExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// GraphExpander expands the Graph resource
type GraphExpander struct {
	ExpanderBase
	client    *http.Client
	armClient *armclient.Client
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

	baseItems := []*TreeNode{
		{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/me",
			Namespace:             "graph",
			Name:                  "Me",
			Display:               "Me",
			ItemType:              "graphMe",
			ExpandURL:             "/me",
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
		{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/apps",
			Namespace:             "graph",
			Name:                  "Apps",
			Display:               "Apps",
			ItemType:              "apps",
			ExpandURL:             "/applications?$count=true", //&$search="displayName:damoo"
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
		{
			Parentid:              currentItem.ID,
			ID:                    currentItem.ID + "/servicePrincipals",
			Namespace:             "graph",
			Name:                  "Service Principals",
			Display:               "Service Principals",
			ItemType:              "servicePrincipals",
			ExpandURL:             "/servicePrincipals?$count=true", //&$search="displayName:damoo"
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
		},
	}

	if currentItem.Name == "Me" {
		// make the graph call
		data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err == nil {
			var meResponse MeResponse
			err = json.Unmarshal([]byte(data), &meResponse)
			if err != nil {
				panic(err)
			}

			return ExpanderResult{
				Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
				SourceDescription: "GraphMeExpander request",
				Nodes:             []*TreeNode{},
				IsPrimaryResponse: true,
			}
		}
	}

	if currentItem.Name == "Apps" {
		// get the list of apps
		data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
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
						ExpandURL:             "/applications/" + app.Id,
						ItemType:              "app",
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
	}

	if currentItem.Parentid == "graph/apps" {
		// get the list of apps
		data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err == nil {
			return ExpanderResult{
				Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
				SourceDescription: "GraphAppExpander request",
				Nodes:             []*TreeNode{},
				IsPrimaryResponse: true,
			}
		}
	}

	if currentItem.Name == "Service Principals" {
		// get the list of sp's
		data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err == nil {
			var spListResponse AppsListResponse
			err = json.Unmarshal([]byte(data), &spListResponse)
			if err != nil {
				panic(err)
			}

			nodes := []*TreeNode{}
			for _, servicePrincipal := range spListResponse.Items {

				nodes = append(nodes,
					&TreeNode{
						Parentid:              currentItem.ID,
						ID:                    servicePrincipal.Id,
						Namespace:             "graph",
						Name:                  servicePrincipal.DisplayName,
						Display:               servicePrincipal.DisplayName,
						ExpandURL:             "/servicePrincipals/" + servicePrincipal.Id,
						ItemType:              "app",
						SuppressGenericExpand: true,
					})
			}

			return ExpanderResult{
				Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
				SourceDescription: "GraphSPListExpander request",
				Nodes:             nodes,
				IsPrimaryResponse: true,
			}
		}
	}

	if currentItem.Parentid == "graph/servicePrincipals" {
		// get the list of sps
		data, err := e.armClient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err == nil {
			return ExpanderResult{
				Response:          ExpanderResponse{Response: data, ResponseType: interfaces.ResponseJSON},
				SourceDescription: "GraphAppExpander request",
				Nodes:             []*TreeNode{},
				IsPrimaryResponse: true,
			}
		}
	}

	return ExpanderResult{
		Err:               nil,
		SourceDescription: "GraphExpander request",
		Nodes:             baseItems,
		IsPrimaryResponse: false,
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *GraphExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {
	return false, nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *GraphExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *GraphExpander) ListActions(context context.Context, item *TreeNode) ListActionsResult {
	return ListActionsResult{
		Nodes:             nil,
		SourceDescription: "GraphExpander",
		IsPrimaryResponse: true,
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *GraphExpander) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	return ExpanderResult{
		SourceDescription: "GraphExpander",
		Err:               fmt.Errorf("Unhandled ActionID: %q", nil),
	}
}

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
