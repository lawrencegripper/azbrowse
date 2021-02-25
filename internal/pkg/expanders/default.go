package expanders

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/nbio/st"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Check interface
var _ Expander = &DefaultExpander{}

// DefaultExpander expands RGs under a subscription
type DefaultExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *DefaultExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *DefaultExpander) Name() string {
	return "GenericExpander"
}

// DoesExpand checks if this handler can expand this item
func (e *DefaultExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ExpandURL == ExpandURLNotSupported || currentItem.SuppressGenericExpand {
		return false, nil
	}
	return true, nil
}

// Expand returns resource
func (e *DefaultExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"

	data, err := e.client.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
			SourceDescription: "Default Expander Request",
		}
	}

	var resource armclient.Resource
	err = json.Unmarshal([]byte(data), &resource)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
			SourceDescription: "Default Expander Request",
		}
	}

	// Update the existing state as we have more up-to-date info
	newStatus := DrawStatus(resource.Properties.ProvisioningState)
	if newStatus != currentItem.StatusIndicator {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: true,
			Message:    "Updated resource status -> " + DrawStatus(resource.Properties.ProvisioningState),
			Timeout:    time.Duration(time.Second * 3),
		})
		currentItem.StatusIndicator = newStatus
	}

	return ExpanderResult{
		Err:               err,
		Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
		SourceDescription: "Default Expander Request",
	}
}

// HasActions indicates whether we can attempt to retrieve actions for the current node
func (e *DefaultExpander) HasActions(context context.Context, item *TreeNode) (bool, error) {
	if item.ItemType != ResourceType ||
		item.Namespace == "" ||
		item.ArmType == "" {
		return false, nil
	}
	return true, nil
}

// ListActions returns the actions available from querying ARM
func (e *DefaultExpander) ListActions(ctx context.Context, item *TreeNode) ListActionsResult {
	var namespace string
	var armType string

	if item.ItemType == ResourceType {
		namespace = item.Namespace
		armType = item.ArmType
	}

	if namespace == "" || armType == "" {
		return ListActionsResult{
			Nodes: []*TreeNode{},
		}
	}

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+item.Name, tracing.SetTag("item", item))
	defer span.Finish()

	data, err := armclient.LegacyInstance.DoRequest(ctx, "GET", "/providers/Microsoft.Authorization/providerOperations/"+namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		return ListActionsResult{
			Err: fmt.Errorf("Failed to get actions: %s", err),
		}
	}
	var opsRequest OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	items := []*TreeNode{}
	for _, resOps := range opsRequest.ResourceTypes {
		if resOps.Name == strings.Split(armType, "/")[1] {
			for _, op := range resOps.Operations {
				resourceAPIVersion, err := armclient.GetAPIVersion(item.ArmType)
				if err != nil {
					return ListActionsResult{
						Err: fmt.Errorf("Failed to find an api version: %s", err),
					}
				}
				stripArmType := strings.Replace(op.Name, item.ArmType, "", -1)
				name := op.DisplayName
				if name == "" {
					name = strings.Replace(stripArmType, "/action", "", -1)
				}
				actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + resourceAPIVersion
				items = append(items, &TreeNode{
					Name:             name,
					Display:          name,
					ExpandURL:        item.ID + "/" + actionURL,
					ExpandReturnType: ActionType,
					ItemType:         "action",
					ID:               item.ID + "/" + actionURL,
				})
			}
		}
	}
	return ListActionsResult{
		Nodes: items,
	}
}

// ExecuteAction executes an action returned from ListActions
func (e *DefaultExpander) ExecuteAction(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "POST"
	data, err := e.client.DoRequest(ctx, method, currentItem.ExpandURL)

	return ExpanderResult{
		Err:               err,
		Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
		SourceDescription: "Resource Group Request",
		IsPrimaryResponse: true,
	}
}

func (e *DefaultExpander) testCases() (bool, *[]expanderTestCase) {
	const testPath = "subscriptions/thing"
	itemToExpand := &TreeNode{
		ExpandURL: "https://management.azure.com/" + testPath,
	}
	const testResponseFile = "./testdata/armsamples/resource/failingResource.json"

	return true, &[]expanderTestCase{
		{
			name:         "Default->Resource",
			nodeToExpand: itemToExpand,
			urlPath:      testPath,
			responseFile: testResponseFile,
			statusCode:   200,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)
				st.Expect(t, len(r.Nodes), 0)

				dat, err := ioutil.ReadFile(testResponseFile)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				st.Expect(t, strings.TrimSpace(r.Response.Response), string(dat))
				st.Expect(t, itemToExpand.StatusIndicator, "⛈")
			},
		},
		{
			name:         "Default->500StatusCode",
			nodeToExpand: itemToExpand,
			urlPath:      testPath,
			responseFile: testResponseFile,
			statusCode:   500,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				if r.Err == nil {
					t.Error("Failed expanding resource. Should have errored and didn't", r)
				}
			},
		},
	}
}

// OperationsRequest list the actions that can be performed
type OperationsRequest struct {
	DisplayName string `json:"displayName"`
	Operations  []struct {
		Name         string      `json:"name"`
		DisplayName  string      `json:"displayName"`
		Description  string      `json:"description"`
		Origin       interface{} `json:"origin"`
		Properties   interface{} `json:"properties"`
		IsDataAction bool        `json:"isDataAction"`
	} `json:"operations"`
	ResourceTypes []struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Operations  []struct {
			Name         string      `json:"name"`
			DisplayName  string      `json:"displayName"`
			Description  string      `json:"description"`
			Origin       interface{} `json:"origin"`
			Properties   interface{} `json:"properties"`
			IsDataAction bool        `json:"isDataAction"`
		} `json:"operations"`
	} `json:"resourceTypes"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
