package handlers

import (
	"context"
)

// ActionExpander handles actions
type ActionExpander struct{}

// Name returns the name of the expander
func (e *ActionExpander) Name() string {
	return "ActionExpander"
}

// DoesExpand checks if it is an action
func (e *ActionExpander) DoesExpand(ctx context.Context, currentItem TreeNode) (bool, error) {
	if currentItem.ItemType == ActionType {
		panic("Not implemented")
	}

	return false, nil
}

// Expand performs the action
func (e *ActionExpander) Expand(ctx context.Context, currentItem TreeNode) ExpanderResult {
	panic("Not implemented")
}
