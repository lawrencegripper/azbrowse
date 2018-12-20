package search

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/storage"
)

// StartCrawler grabs all the resources in all subs and stores their name/id in the boltdb for searching over
func StartCrawler(subs armclient.SubResponse) error {
	wait := &sync.WaitGroup{}

	for _, sub := range subs.Subs {
		wait.Add(1)
		subID := sub.ID
		go func() {
			fmt.Println("Starting with sub " + subID)
			rgListURL := subID + "/resources?api-version=2014-04-01"
			err := fetchAndStoreURL(subID, rgListURL)
			if err != nil {
				panic(err)
			}
			wait.Done()
		}()
	}

	wait.Wait()

	res, err := storage.GetAllResources()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Got %v resources", len(*res))
	return nil
}

func fetchAndStoreURL(subID, url string) error {
	fmt.Println("-- Fetching for sub " + subID)
	data, err := armclient.DoRequest("GET", url)
	if err != nil {
		panic(err)
	}
	var resourceResponse armclient.SubResourcesResponse
	err = json.Unmarshal([]byte(data), &resourceResponse)
	if err != nil {
		return err
	}
	resourcesToStore := make([]storage.Resource, 0, len(resourceResponse.Resources))
	for _, r := range resourceResponse.Resources {
		resourcesToStore = append(resourcesToStore, storage.Resource{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	storage.PutResourceBatch(subID, resourcesToStore)

	if resourceResponse.NextLink != "" {
		return fetchAndStoreURL(subID, resourceResponse.NextLink)
	}
	return nil
}
