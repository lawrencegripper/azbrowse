package automation

import (
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

var navigateToInProgress = true

// NavigateTo will navigate through the tree to a node with
// a matching ItemID or as far as it can get
func NavigateTo(list *views.ListWidget, itemID string) {

	navigateToIDLower := strings.ToLower(itemID)
	go func() {
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")
		var lastNavigatedNode *expanders.TreeNode

		for {
			navigateStateInterface := <-navigatedChannel

			if navigateToInProgress {
				navigateState := navigateStateInterface.(views.ListNavigatedEventState)
				if !navigateState.Success {
					// we got as far as we could - now stop!
					navigateToInProgress = false
					list.SetShouldRender(true)
					continue
				}
				nodeList := navigateState.NewNodes

				if lastNavigatedNode != nil && lastNavigatedNode != list.CurrentExpandedItem() {
					navigateToInProgress = false
					list.SetShouldRender(true)
				} else {

					gotNode := false
					for nodeIndex, node := range nodeList {
						// use prefix matching
						// but need additional checks as target of /foo/bar would be matched by  /foo/bar  and /foo/ba
						// additional check is that the lengths match, or the next char in target is a '/'
						nodeIDLower := strings.ToLower(node.ID)
						if strings.HasPrefix(navigateToIDLower, nodeIDLower) && (len(itemID) == len(nodeIDLower) || navigateToIDLower[len(nodeIDLower)] == '/') {
							list.ChangeSelection(nodeIndex)
							lastNavigatedNode = node
							list.ExpandCurrentSelection()
							gotNode = true
							break
						}
					}

					if !gotNode {
						// we got as far as we could - now stop!
						navigateToInProgress = false
						list.SetShouldRender(true)
					}
				}
			}
		}
	}()
}
