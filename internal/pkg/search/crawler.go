package search

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/schollz/closestmatch"
)

// CrawlResources grabs all the resources in all subs and stores their name/id in the boltdb for searching over
func CrawlResources(ctx context.Context, subs armclient.SubResponse) error {
	wait := &sync.WaitGroup{}

	err := storage.ClearResources()
	if err != nil {
		return err
	}

	for _, sub := range subs.Subs {
		wait.Add(1)
		subID := sub.ID
		go func() {
			ctx, cancel := context.WithTimeout(ctx, time.Duration(30*time.Minute))
			defer cancel()
			rgList := subID + "/resourceGroups?api-version=2014-04-01"
			err := fetchAndStoreGroups(ctx, rgList)
			if err != nil {
				panic(err)
			}

			resourcesListURL := subID + "/resources?api-version=2014-04-01"
			err = fetchAndStoreResourceURL(ctx, resourcesListURL)
			if err != nil {
				panic(err)
			}
			wait.Done()
		}()
	}

	wait.Wait()
	return nil
}

func fetchAndStoreGroups(ctx context.Context, url string) error {
	data, err := armclient.DoRequest(ctx, "GET", url)
	if err != nil {
		return fmt.Errorf("Failed requesting %s: %v", url, err)
	}

	var rgResponse armclient.ResourceGroupResponse
	err = json.Unmarshal([]byte(data), &rgResponse)
	if err != nil {
		panic(err)
	}

	rgToStore := make([]storage.Resource, 0, len(rgResponse.Groups))
	for _, r := range rgResponse.Groups {
		rgToStore = append(rgToStore, storage.Resource{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	storage.PutResourceBatch(url, rgToStore)

	return nil
}

// fetchAndStoreURL takes a link to a page of '/resources'
// and stores the resulting batch into a bucket in boltdb
func fetchAndStoreResourceURL(ctx context.Context, url string) error {
	data, err := armclient.DoRequest(ctx, "GET", url)
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
	if len(resourcesToStore) == 0 {
		return nil
	}
	storage.PutResourceBatch(url, resourcesToStore)

	if resourceResponse.NextLink != "" {
		// Next links are fully formed, including 'https://management.azure.com/'
		// we don't want this so we strip that out
		return fetchAndStoreResourceURL(ctx, strings.Replace(resourceResponse.NextLink, "https://management.azure.com", "", 1))
	}
	return nil
}

// Suggester provides search functionality over all resources in all subs
type Suggester struct {
	matcher *closestmatch.ClosestMatch
}

// NewSuggester builds a suggester based on your resources
func NewSuggester() (*Suggester, error) {
	resources, err := storage.GetAllResources()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(resources))
	for _, r := range resources {
		names = append(names, r.Name)
	}

	// Choose a set of bag sizes, more is more accurate but slower
	bagSizes := []int{3}

	// Create a closestmatch object
	cm := closestmatch.New(names, bagSizes)
	return &Suggester{matcher: cm}, nil
}

// Autocomplete finds items based on the query string
func (s *Suggester) Autocomplete(query string) []string {
	return s.matcher.ClosestN(query, 8)
}
