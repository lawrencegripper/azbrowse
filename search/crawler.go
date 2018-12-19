package search

import (
	"encoding/json"
	"fmt"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/storage"
)

// StartCrawler grabs all the resources in all subs and stores their name/id in the boltdb for searching over
func StartCrawler(subs armclient.SubResponse) error {
	for _, sub := range subs.Subs {
		fmt.Println("Starting with sub " + sub.DisplayName)
		rgListURL := sub.ID + "/resourceGroups?api-version=2014-04-01"
		data, err := armclient.DoRequest("GET", rgListURL)
		if err != nil {
			panic(err)
		}

		var rgResponse armclient.ResourceGroupResponse
		err = json.Unmarshal([]byte(data), &rgResponse)
		if err != nil {
			panic(err)
		}

		for _, rg := range rgResponse.Groups {
			fmt.Println("--Starting with rg " + rg.Name)

			data, err := armclient.DoRequest("GET", rg.ID+"/resources?api-version=2017-05-10")

			var resourcesResponse armclient.ResourceReseponse
			err = json.Unmarshal([]byte(data), &resourcesResponse)
			if err != nil {
				panic(err)
			}

			for _, resource := range resourcesResponse.Resources {
				fmt.Println("---- Saving resource " + resource.Name)

				storage.PutResource(resource.ID, storage.Resource{
					ID:                resource.ID,
					Name:              resource.Name,
					ResourceGroupID:   rg.ID,
					ResourceGroupName: rg.Name,
				})
			}
		}

	}
	return nil
}
