package handlers

// HandlerFunc is used to expand a node in the `TreeView`
// It takes current node and returns a list of new nodes
// to be displayed
type HandlerFunc func(currentNode TreeNode) (matched bool, newNodes []TreeNode, responseContent string, err error)

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
