package expanders

import (
	"context"
	"fmt"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

// Expander is used to open/expand items in the left list panel
// a single item can be expanded by 1 or more expanders
// each Expander provides two methods.
// `DoesExpand` should return true if this expander can expand the resource
// `Expand` returns the list of sub items from the resource
type Expander interface {
	DoesExpand(ctx context.Context, currentNode *TreeNode) (bool, error)
	Expand(ctx context.Context, currentNode *TreeNode) ExpanderResult
	Name() string // Returns the name of this expander for logging (ie. TenantExpander)
	Delete(context context.Context, item *TreeNode) (bool, error)

	HasActions(ctx context.Context, currentNode *TreeNode) (bool, error)
	ListActions(ctx context.Context, currentNode *TreeNode) ListActionsResult
	ExecuteAction(ctx context.Context, currentNode *TreeNode) ExpanderResult

	CanUpdate(context context.Context, item *TreeNode) (bool, error)
	Update(context context.Context, item *TreeNode, updatedContent string) error

	// Used for testing the expanders
	testCases() (bool, *[]expanderTestCase)
	setClient(c *armclient.Client)
}

// ExpanderBase provides nil implementations of Expander methods to avoid duplicating code
type ExpanderBase struct{}

func (e *ExpanderBase) setClient(c *armclient.Client) {
}
func (e *ExpanderBase) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
func (e ExpanderBase) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	return false, nil
}
func (e *ExpanderBase) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	return ExpanderResult{
		SourceDescription: "ExpanderBase",
		Err:               fmt.Errorf("ExpanderBase.Expand should not be called"),
	}
}
func (e ExpanderBase) CanUpdate(ctx context.Context, currentItem *TreeNode) (bool, error) {
	return false, nil
}
func (e *ExpanderBase) Update(ctx context.Context, currentItem *TreeNode, updatedContent string) error {
	return fmt.Errorf("ExpanderBase.Update should not be called")
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *ExpanderBase) Delete(context context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

// HasActions is a default implementation returning false to indicate no actions available
func (e *ExpanderBase) HasActions(context context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

// ListActions returns an error as it should not be called as HasActions returns false
func (e *ExpanderBase) ListActions(context context.Context, item *TreeNode) ListActionsResult {
	return ListActionsResult{
		SourceDescription: "ExpanderBase",
		Err:               fmt.Errorf("ExpanderBase.ListActions should not be called"),
	}
}

// ExecuteAction returns an error as it should not be called as HasActions returns false
func (e *ExpanderBase) ExecuteAction(context context.Context, item *TreeNode) ExpanderResult {
	return ExpanderResult{
		SourceDescription: "ExpanderBase",
		Err:               fmt.Errorf("ExpanderBase.ListActions should not be called"),
	}
}

// ExpanderResponse captures the response text and formt of an expander response
type ExpanderResponse struct {
	Response     string                          // the response text
	ResponseType interfaces.ExpanderResponseType // the response
}

// ExpanderResult used to wrap mult-value return for use in channels
type ExpanderResult struct {
	Response          ExpanderResponse
	Nodes             []*TreeNode
	Err               error
	SourceDescription string
	// When set to true this causes the response
	// in the result to be displayed in the content panel
	IsPrimaryResponse bool
}

// ListActionsResult
type ListActionsResult struct {
	Nodes             []*TreeNode // TODO - should this be nodes or metadata that something else renders as nodes?
	SourceDescription string
	Err               error
	IsPrimaryResponse bool // Causes nodes to be added at the top of the list
}

// TreeNode is an item in the ListWidget
type TreeNode struct {
	Parentid               string                // The ID of the parent resource
	Parent                 *TreeNode             // Reference to the parent node
	ID                     string                // The ID of the resource in ARM
	Name                   string                // The name of the object returned by the API
	Display                string                // The Text used to draw the object in the list
	ExpandURL              string                // The URL to call to expand the item
	ItemType               string                // The type of item either subscription, resourcegroup, resource, deployment or action
	ExpandReturnType       string                // The type of the items returned by the expandURL
	DeleteURL              string                // The URL to call to delete the current resource
	Namespace              string                // The ARM Namespace of the item eg StorageAccount
	ArmType                string                // The ARM type of the item eg Microsoft.Storage/StorageAccount
	Metadata               map[string]string     // Metadata is used to pass arbritray data between `Expander`'s
	SubscriptionID         string                // The SubId of this item
	StatusIndicator        string                // Displays the resources status
	SwaggerResourceType    *swagger.ResourceType // caches the swagger ResourceType to avoid repeated lookups
	Expander               Expander              // The Expander that created the node (set automatically by the list)
	SuppressSwaggerExpand  bool                  // Prevent the SwaggerResourceExpander attempting to expand the node
	SuppressGenericExpand  bool                  // Prevent the DefaultExpander (aka GenericExpander) attempting to expand the node
	TimeoutOverrideSeconds *int                  // Override the default expand timeout for a node
	ExpandInPlace          bool                  // Indicates that the node is a "More..." node. Must be the last in the list and will be removed and replaced with the expanded nodes
}

const (
	// SubscriptionType defines a sub
	SubscriptionType = "subscription"
	// ResourceGraphQueryType defines an Azure resource graph query
	ResourceGraphQueryType = "graphquery"
	resourceGroupType      = "resourcegroup"
	// ResourceType defines a top level resource such as a Storage Account or VM
	ResourceType = "resource"
	// MetricsType defines an item which returns a graph
	MetricsType = "metrics"
	// SubResourceType defines a resource under a resource such as a VM Extension under a VM
	SubResourceType = "subResource"
	// deploymentType represents a deployment node
	deploymentType = "deployment"
	// deploymentsType represents the "Deployments" placholder node
	deploymentsType         = "deployments"
	deploymentOperationType = "deploymentOperation"
	activityLogType         = "activityLog"
	subActivityLogType      = "subActivityLog"
	diagnosticSettingsType  = "diagnosticSettings"
	// ActionType defines an action like `listkey` etc
	ActionType = "action"

	// Used to store resourceIds as CVS in TreeItem Metadata
	resourceIdsMeta = "resourceIds"

	// ExpandURLNotSupported is used to identify items which don't support generic expansion
	ExpandURLNotSupported = "notsupported"
	NotSupported          = "notsupported"
)

type expanderTestCase struct {
	name                string
	statusCode          int
	nodeToExpand        *TreeNode
	urlPath             string
	responseFile        string
	configureGockFunc   *func(t *testing.T)
	treeNodeCheckerFunc func(t *testing.T, r ExpanderResult)
}
