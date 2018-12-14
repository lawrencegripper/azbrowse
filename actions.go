package main

import (
	"encoding/json"
	// "fmt"
	"github.com/lawrencegripper/azbrowse/armclient"
	// "strings"
)

// LoadActionsView Shows available actions for the current resource
func LoadActionsView(list *ListWidget) error {
	data, err := armclient.DoRequest("GET", "/providers/Microsoft.Authorization/providerOperations/"+list.CurrentItem().namespace+"?api-version=2018-01-01-preview&$expand=resourceTypes")
	if err != nil {
		panic(err)
	}
	var opsRequest armclient.OperationsRequest
	err = json.Unmarshal([]byte(data), &opsRequest)
	if err != nil {
		panic(err)
	}

	list.contentView.Content = data

	// items := []TreeNode
	// for _, resOps := range opsRequest.ResourceTypes {
	// 	if resOps.Name == strings.Split(list.CurrentItem().armType, "/")[1] {
	// 		for _, op := range resOps.Operations {
	// 			// items = append(items, TreeNode{
	// 			// 	name: op.DisplayName,
	// 			// 	expandURL: op.Properties.
	// 			// })
	// 		}
	// 	}
	// }
	return nil
}
