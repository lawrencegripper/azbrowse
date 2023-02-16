package expanders

import (
	"context"
	"fmt"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

type expanderAndResponse struct {
	Expander       Expander
	ExpanderResult ExpanderResult
}

// ExpandItem finds child nodes of the item and their content
func ExpandItem(ctx context.Context, currentItem *TreeNode) (*ExpanderResponse, []*TreeNode, error) {
	return ExpandItemAllowDefaultExpander(ctx, currentItem, true)
}

// ExpandItemAllowDefaultExpander finds child nodes of the item and their content and allows the default expander to be suppressed
func ExpandItemAllowDefaultExpander(ctx context.Context, currentItem *TreeNode, allowDefaultExpander bool) (*ExpanderResponse, []*TreeNode, error) {

	newItems := []*TreeNode{}

	_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Opening: " + currentItem.ID,
		InProgress: true,
	})
	defer done()

	span, ctx := tracing.StartSpanFromContext(ctx, "expand:"+currentItem.ItemType+":"+currentItem.Name, tracing.SetTag("item", currentItem))
	defer span.Finish()

	// New handler approach
	handlerExpanding := 0

	completedExpands := make(chan expanderAndResponse)

	// Check which expanders are interested and kick them off
	spanQuery, _ := tracing.StartSpanFromContext(ctx, "querexpanders", tracing.SetTag("item", currentItem))
	for _, h := range getRegisteredExpanders() {
		doesExpand, err := h.DoesExpand(ctx, currentItem)
		spanQuery.SetTag(h.Name(), doesExpand)
		if err != nil {
			panic(err)
		}
		if !doesExpand {
			continue
		}

		// Fire each handler in parallel
		hCurrent := h // capture current iteration variable
		go func() {
			// recover from panic, if one occurrs, and leave terminal usable
			defer errorhandling.RecoveryWithCleanup()

			completedExpands <- expanderAndResponse{
				Expander:       hCurrent,
				ExpanderResult: hCurrent.Expand(ctx, currentItem),
			}
		}()

		handlerExpanding = handlerExpanding + 1
	}
	spanQuery.Finish()

	// Lets give all the expanders 45secs to completed (unless debugging)
	hasPrimaryResponse := false
	timeoutSeconds := 45
	if currentItem.TimeoutOverrideSeconds != nil {
		timeoutSeconds = *currentItem.TimeoutOverrideSeconds
	}

	timeout := time.After(time.Second * time.Duration(timeoutSeconds))
	var newContent ExpanderResponse

	for index := 0; index < handlerExpanding; index++ {
		select {
		case done := <-completedExpands:
			result := done.ExpanderResult
			span, _ := tracing.StartSpanFromContext(ctx, "subexpand:"+result.SourceDescription, tracing.SetTag("result", done))
			// Did it fail?
			if result.Err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					Failure: true,
					Message: "Expander '" + result.SourceDescription + "' failed on resource: " + currentItem.ID + "Err: " + result.Err.Error(),
					Timeout: time.Duration(time.Second * 15),
				})
			}
			if result.IsPrimaryResponse {
				if hasPrimaryResponse {
					panic(fmt.Sprintf("Two handlers returned a primary response for this item... failing. ID: %s EXPANDER: %s", currentItem.ID, result.SourceDescription))
				}
				// Log that we have a primary response
				hasPrimaryResponse = true
				if result.Response.Response != "" || result.Err == nil {
					newContent = result.Response
				} else {
					newContent = ExpanderResponse{Response: result.Err.Error(), ResponseType: interfaces.ResponsePlainText}
				}
			}
			if result.Nodes == nil {
				continue
			}
			for _, node := range result.Nodes {
				if node.Expander == nil { // action expander sets this to the underlying expander - don't overwrite
					node.Expander = done.Expander
				}
				node.Parent = currentItem
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
				Message: "Timed out opening:" + currentItem.ID,
				Timeout: time.Duration(time.Second * 10),
			})
			return nil, nil, fmt.Errorf("Timed out opening: %s", currentItem.ID)
		}
	}

	if allowDefaultExpander {
		// Use the default handler to get the resource JSON for display
		defaultExpanderWorksOnThisItem, _ := GetDefaultExpander().DoesExpand(ctx, currentItem)
		if !hasPrimaryResponse && defaultExpanderWorksOnThisItem {
			result := GetDefaultExpander().Expand(ctx, currentItem)
			if result.Err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					Failure: true,
					Message: "Failed to expand resource: " + result.Err.Error(),
					Timeout: time.Duration(time.Second * 15),
				})
			}
			newContent = result.Response
		}
	}

	return &newContent, newItems, nil
}
