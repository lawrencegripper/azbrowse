package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

func (e *AzureSearchServiceExpander) loadResourceTypes() []swagger.ResourceType {
	return []swagger.ResourceType{
		{
			Display:  "datasources",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/datasources", "2019-05-06"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{dataSourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/datasources('{dataSourceName}')", "2019-05-06"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/datasources('{dataSourceName}')", "2019-05-06"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/datasources('{dataSourceName}')", "2019-05-06"),
				}},
		},
		{
			Display:  "indexers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexers", "2019-05-06"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{indexerName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/indexers('{indexerName}')", "2019-05-06"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/indexers('{indexerName}')", "2019-05-06"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/indexers('{indexerName}')", "2019-05-06"),
					Children: []swagger.ResourceType{
						{
							Display:  "search.status",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexers('{indexerName}')/search.status", "2019-05-06"),
						}},
				}},
		},
		{
			Display:  "indexes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes", "2019-05-06"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{indexName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')", "2019-05-06"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')", "2019-05-06"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')", "2019-05-06"),
					Children: []swagger.ResourceType{
						{
							Display:  "search.stats",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/search.stats", "2019-05-06"),
						},
						{
							Display:  "docs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs", "2019-05-06"),
							Children: []swagger.ResourceType{
								{
									Display:  "$count",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs/$count", "2019-05-06"),
								},
								{
									Display:  "search.autocomplete",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs/search.autocomplete", "2019-05-06"),
								},
								{
									Display:  "search.suggest",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs/search.suggest", "2019-05-06"),
								}},
							SubResources: []swagger.ResourceType{
								{
									Display:        "{key}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs('{key}')", "2019-05-06"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs/index", "2019-05-06"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/indexes('{indexName}')/docs/index", "2019-05-06"),
								}},
						}},
				}},
		},
		{
			Display:  "servicestats",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/servicestats", "2019-05-06"),
		},
		{
			Display:  "skillsets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/skillsets", "2019-05-06"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{skillsetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/skillsets('{skillsetName}')", "2019-05-06"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/skillsets('{skillsetName}')", "2019-05-06"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/skillsets('{skillsetName}')", "2019-05-06"),
				}},
		},
		{
			Display:  "synonymmaps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/synonymmaps", "2019-05-06"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{synonymMapName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/synonymmaps('{synonymMapName}')", "2019-05-06"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/synonymmaps('{synonymMapName}')", "2019-05-06"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/synonymmaps('{synonymMapName}')", "2019-05-06"),
				}},
		}}

}
