package main

import (
	"encoding/json"
	// "fmt"
	"github.com/lawrencegripper/azbrowse/armclient"
	"strings"
)

// LoadActionsView Shows available actions for the current resource
func LoadActionsView(list *ListWidget) error {
	list.statusView.Status("Getting available Actions", true)
	currentItem := list.CurrentItem()
	data, err := armclient.DoRequest("GET", "/providers/Microsoft.Authorization/providerOperations/"+list.CurrentItem().namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		list.statusView.Status("Failed to get actions: "+err.Error(), false)
	}
	var opsRequest armclient.OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	items := []TreeNode{}
	for _, resOps := range opsRequest.ResourceTypes {
		if resOps.Name == strings.Split(list.CurrentItem().armType, "/")[1] {
			for _, op := range resOps.Operations {
				stripArmType := strings.Replace(op.Name, currentItem.armType, "", -1)
				actionURL := strings.Replace(stripArmType, "/action", "", -1) + "?api-version=" + list.resourceAPIVersionLookup[currentItem.armType]
				items = append(items, TreeNode{
					name:             op.DisplayName,
					expandURL:        currentItem.id + "/" + actionURL,
					expandReturnType: actionType,
				})
			}
		}
	}
	if len(items) > 1 {
		list.SetNodes(items)
	}
	list.statusView.Status("Fetched available Actions", false)

	return nil
}
