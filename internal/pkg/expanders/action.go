package expanders

import (
	"context"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// ActionExpander handles actions
type ActionExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *ActionExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *ActionExpander) Name() string {
	return "ActionExpander"
}

// DoesExpand checks if it is an action
func (e *ActionExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == ActionType {
		return true, nil
	}

	return false, nil
}

// Expand performs the action
func (e *ActionExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "POST"

	_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		InProgress: true,
		Message:    "Action:" + currentItem.Name + " @ " + currentItem.ID,
		Timeout:    time.Duration(time.Second * 45),
	})
	defer done()

	data, err := e.client.DoRequest(ctx, method, currentItem.ExpandURL)

	return ExpanderResult{
		Err:               err,
		Response:          ExpanderResponse{Response: string(data), ResponseType: ResponseJSON},
		SourceDescription: "Resource Group Request",
		IsPrimaryResponse: true,
	}
}

func (e *ActionExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
