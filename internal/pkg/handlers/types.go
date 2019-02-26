package handlers

import (
	"context"
)

// Expander is used to open/expand items in the left list panel
// a single item can be expanded by 1 or more expanders
// each Expander provides two methods.
// `DoesExpand` should return true if this expander can expand the resource
// `Expand` returns the list of sub items from the resource
type Expander interface {
	DoesExpand(ctx context.Context, currentNode *TreeNode) (bool, error)
	Expand(ctx context.Context, currentNode *TreeNode) ExpanderResult
	Name() string
}

// ExpanderResult used to wrap mult-value return for use in channels
type ExpanderResult struct {
	Response          string
	Nodes             []*TreeNode
	Err               error
	SourceDescription string
	// When set to true this causes the response
	// in the result to be displayed in the content panel
	IsPrimaryResponse bool
}

// TreeNode is an item in the ListWidget
type TreeNode struct {
	Parentid            string            // The ID of the parent resource
	ID                  string            // The ID of the resource in ARM
	Name                string            // The name of the object returned by the API
	Display             string            // The Text used to draw the object in the list
	ExpandURL           string            // The URL to call to expand the item
	ItemType            string            // The type of item either subscription, resourcegroup, resource, deployment or action
	ExpandReturnType    string            // The type of the items returned by the expandURL
	DeleteURL           string            // The URL to call to delete the current resource
	Namespace           string            // The ARM Namespace of the item eg StorageAccount
	ArmType             string            // The ARM type of the item eg Microsoft.Storage/StorageAccount
	Metadata            map[string]string // Metadata is used to pass arbritray data between `Expander`'s
	SubscriptionID      string            // The SubId of this item
	StatusIndicator     string            // Displays the resources status
	SwaggerResourceType *ResourceType     // caches the swagger ResourceType to avoid repeated lookups
}

const (
	// SubscriptionType defines a sub
	SubscriptionType  = "subscription"
	resourceGroupType = "resourcegroup"
	resourceType      = "resource"
	deploymentType    = "deployment"
	subDeploymentType = "subDeployment"
	// ActionType defines an action like `listkey` etc
	ActionType = "action"

	// ExpandURLNotSupported is used to identify items which don't support generic expansion
	ExpandURLNotSupported = "notsupported"
)
