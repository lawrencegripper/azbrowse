package expanders

import (
	"context"
	"testing"

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
	Name() string
	Delete(context context.Context, item *TreeNode) (bool, error)

	// Used for testing the expanders
	testCases() (bool, *[]expanderTestCase)
	setClient(c *armclient.Client)
}

// ExpanderBase provides nil implementations of Expander methods to avoid duplicating code
type ExpanderBase struct{}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *ExpanderBase) Delete(context context.Context, item *TreeNode) (bool, error) {
	return false, nil
}

// ExpanderResponseType is used to indicate the text format of a response
type ExpanderResponseType string

const (
	// ResponsePlainText indicates the response type should not be parsed or colourised
	ResponsePlainText ExpanderResponseType = "Text"
	// ResponseJSON indicates the response type can be parsed and colourised as JSON
	ResponseJSON ExpanderResponseType = "JSON"
	// ResponseYAML indicates the response type can be parsed and colourised as YAML
	ResponseYAML ExpanderResponseType = "YAML"
	// ResponseXML indicates the response type can be parsed and colourised as XML
	ResponseXML ExpanderResponseType = "XML"
)

// ExpanderResponse captures the response text and formt of an expander response
type ExpanderResponse struct {
	Response     string               // the response text
	ResponseType ExpanderResponseType // the response
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

// TreeNode is an item in the ListWidget
type TreeNode struct {
	Parentid            string                // The ID of the parent resource
	ID                  string                // The ID of the resource in ARM
	Name                string                // The name of the object returned by the API
	Display             string                // The Text used to draw the object in the list
	ExpandURL           string                // The URL to call to expand the item
	ItemType            string                // The type of item either subscription, resourcegroup, resource, deployment or action
	ExpandReturnType    string                // The type of the items returned by the expandURL
	DeleteURL           string                // The URL to call to delete the current resource
	Namespace           string                // The ARM Namespace of the item eg StorageAccount
	ArmType             string                // The ARM type of the item eg Microsoft.Storage/StorageAccount
	Metadata            map[string]string     // Metadata is used to pass arbritray data between `Expander`'s
	SubscriptionID      string                // The SubId of this item
	StatusIndicator     string                // Displays the resources status
	SwaggerResourceType *swagger.ResourceType // caches the swagger ResourceType to avoid repeated lookups
	Expander            Expander              // The Expander that created the node (set automatically by the list)
}

const (
	// SubscriptionType defines a sub
	SubscriptionType  = "subscription"
	resourceGroupType = "resourcegroup"
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
	// ActionType defines an action like `listkey` etc
	ActionType = "action"

	// ExpandURLNotSupported is used to identify items which don't support generic expansion
	ExpandURLNotSupported = "notsupported"
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
