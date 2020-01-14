package expanders

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
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
	if currentItem.ExpandURL == ExpandURLNotSupported || currentItem.Metadata["SuppressGenericExpand"] == "true" {
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
			Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
			SourceDescription: "Default Expander Request",
		}
	}

	var resource armclient.Resource
	err = json.Unmarshal([]byte(data), &resource)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
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
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		SourceDescription: "Default Expander Request",
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
