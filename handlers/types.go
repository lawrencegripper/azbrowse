package handlers

import (
	"context"
)

// HandlerFunc is used to expand a node in the `TreeView`
// It takes current node and returns a list of new nodes
// to be displayed
type HandlerFunc func(ctx context.Context, currentNode TreeNode) (
	matched bool, // Does the handler work on this type of node?
	newNodes *[]TreeNode, // What new items does it want to add to the list?
	responseContent string, // What response did it get from the API
	err error)

// Register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var Register = []HandlerFunc{
	ResourceGroupHandler,
}

// TreeNode is an item in the ListWidget
type TreeNode struct {
	Parentid         string // The ID of the parent resource
	ID               string // The ID of the resource in ARM
	Name             string // The name of the object returned by the API
	Display          string // The Text used to draw the object in the list
	ExpandURL        string // The URL to call to expand the item
	ItemType         string // The type of item either subscription, resourcegroup, resource, deployment or action
	ExpandReturnType string // The type of the items returned by the expandURL
	DeleteURL        string // The URL to call to delete the current resource
	Namespace        string // The ARM Namespace of the item eg StorageAccount
	ArmType          string // The ARM type of the item eg Microsoft.Storage/StorageAccount
}
