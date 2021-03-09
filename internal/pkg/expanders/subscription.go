package expanders

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/nbio/st"
)

// Check interface
var _ Expander = &SubscriptionExpander{}

// SubscriptionExpander expands RGs under a subscription
type SubscriptionExpander struct {
	ExpanderBase
	client *armclient.Client
}

// Name returns the name of the expander
func (e *SubscriptionExpander) Name() string {
	return "SubscriptionExpander"
}

// DoesExpand checks if this is an RG
func (e *SubscriptionExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == SubscriptionType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *SubscriptionExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"

	data, err := e.client.DoRequest(ctx, method, currentItem.ExpandURL)
	newItems := []*TreeNode{}

	//    \/ It's not the usual ... look out
	if err == nil {
		var rgResponse ResourceGroupResponse
		err = json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		for _, rg := range rgResponse.Groups {
			newItems = append(newItems, &TreeNode{
				Name:             rg.Name,
				Display:          rg.Name,
				ID:               rg.ID,
				Parentid:         currentItem.ID,
				ExpandURL:        rg.ID + "/resources?api-version=2017-05-10",
				ExpandReturnType: ResourceType,
				ItemType:         resourceGroupType,
				DeleteURL:        rg.ID + "?api-version=2017-05-10",
				SubscriptionID:   currentItem.SubscriptionID,
				StatusIndicator:  DrawStatus(rg.Properties.ProvisioningState),
			})
		}
	}

	return ExpanderResult{
		Err:               err,
		Nodes:             newItems,
		Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
		SourceDescription: "Resource Group Request",
		IsPrimaryResponse: true,
	}
}

// ResourceGroupResponse ResourceGroup rest type
type ResourceGroupResponse struct {
	Groups []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Location   string `json:"location"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
		} `json:"properties"`
	} `json:"value"`
}

func (e *SubscriptionExpander) setClient(c *armclient.Client) {
	e.client = c
}

func (e *SubscriptionExpander) testCases() (bool, *[]expanderTestCase) {
	return true, &[]expanderTestCase{
		{
			name: "Subscription->ResourceGroups",
			nodeToExpand: &TreeNode{
				Display:        "Thingy1",
				Name:           "Thingy1",
				ID:             "/subscriptions/00000000-0000-0000-0000-000000000000",
				ExpandURL:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01",
				ItemType:       SubscriptionType,
				SubscriptionID: "00000000-0000-0000-0000-000000000000",
			},
			urlPath:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups",
			responseFile: "./testdata/armsamples/resourcegroups/response.json",
			statusCode:   200,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)
				st.Expect(t, len(r.Nodes), 6)

				// Validate content
				st.Expect(t, r.Nodes[0].Name, "1testrg")
				st.Expect(t, r.Nodes[0].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/1testrg/resources?api-version=2017-05-10")
			},
		},
		{
			name: "Subscription->500StatusCode",
			nodeToExpand: &TreeNode{
				Display:        "Thingy1",
				Name:           "Thingy1",
				ID:             "/subscriptions/00000000-0000-0000-0000-000000000000",
				ExpandURL:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01",
				ItemType:       SubscriptionType,
				SubscriptionID: "00000000-0000-0000-0000-000000000000",
			},
			urlPath:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups",
			statusCode: 500,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				if r.Err == nil {
					t.Error("Failed expanding resource. Should have errored and didn't", r)
				}
			},
		},
	}
}
