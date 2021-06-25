package expanders

import (
	"context"
	"fmt"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

// dummy change in go file
type expanderAndListActionsResponse struct {
	Expander          Expander
	ListActionsResult ListActionsResult
}

// GetActions returns the available actions for a TreeNode
func GetActions(ctx context.Context, item *TreeNode) ([]*TreeNode, error) {
	newItems := []*TreeNode{}
	statusEvent, _ := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Getting available actions",
		InProgress: true,
	})
	defer statusEvent.Done()

	span, ctx := tracing.StartSpanFromContext(ctx, "actions:"+item.ItemType+":"+item.Name, tracing.SetTag("item", item))
	defer span.Finish()

	// New handler approach
	handlerExpanding := 0

	completedExpands := make(chan expanderAndListActionsResponse)

	// Check which expanders are interested and kick them off
	spanQuery, _ := tracing.StartSpanFromContext(ctx, "querexpanders", tracing.SetTag("item", item))
	expanders := append(getRegisteredExpanders(), GetDefaultExpander())
	for _, h := range expanders {
		doesExpand, err := h.HasActions(ctx, item)
		spanQuery.SetTag(h.Name(), doesExpand)
		if err != nil {
			return []*TreeNode{}, err
		}
		if !doesExpand {
			continue
		}

		// Fire each handler in parallel
		hCurrent := h // capture current iteration variable
		go func() {
			// recover from panic, if one occurrs, and leave terminal usable
			defer errorhandling.RecoveryWithCleanup()

			completedExpands <- expanderAndListActionsResponse{
				Expander:          hCurrent,
				ListActionsResult: hCurrent.ListActions(ctx, item),
			}
		}()

		handlerExpanding = handlerExpanding + 1
	}
	spanQuery.Finish()

	// Lets give all the expanders 45secs to completed (unless debugging)
	hasPrimaryResponse := false
	timeout := time.After(time.Second * 45)

	for index := 0; index < handlerExpanding; index++ {
		select {
		case done := <-completedExpands:
			result := done.ListActionsResult
			span, _ := tracing.StartSpanFromContext(ctx, "subexpand:"+result.SourceDescription, tracing.SetTag("result", done))
			// Did it fail?
			if result.Err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					Failure: true,
					Message: "Expander '" + result.SourceDescription + "' failed on resource: " + item.ID + "Err: " + result.Err.Error(),
					Timeout: time.Duration(time.Second * 15),
				})
			}
			if result.IsPrimaryResponse {
				if hasPrimaryResponse {
					panic(fmt.Sprintf("Two handlers returned a primary response for this item... failing. ID: %s EXPANDER: %s", item.ID, result.SourceDescription))
				}
				// Log that we have a primary response
				hasPrimaryResponse = true
			}
			if result.Nodes == nil {
				continue
			}
			for _, node := range result.Nodes {
				node.Expander = done.Expander
				node.ItemType = ActionType
				node.Parent = item
			}
			// Add the items it found
			if result.IsPrimaryResponse {
				newItems = append(result.Nodes, newItems...)
			} else {
				newItems = append(newItems, result.Nodes...)
			}
			span.Finish()
		case <-timeout:
			eventing.SendStatusEvent(&eventing.StatusEvent{
				Failure: true,
				Message: "Timed out opening:" + item.ID,
				Timeout: time.Duration(time.Second * 10),
			})
			return nil, fmt.Errorf("Timed out opening: %s", item.ID)
		}
	}
	return newItems, nil
}
