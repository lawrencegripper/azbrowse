package expanders

import (
	"context"

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
	_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		InProgress: true,
		Message:    "Action:" + currentItem.Name + " @ " + currentItem.ID,
	})
	defer done()

	return currentItem.Expander.ExecuteAction(ctx, currentItem)
}

func (e *ActionExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
