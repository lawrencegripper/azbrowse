package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

func (e *AzureDatabricksExpander) loadResourceTypes() []swagger.ResourceType {
	return []swagger.ResourceType{
		{
			Display:  "clusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/clusters/list", ""),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{cluster_id}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/clusters/get", ""),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/clusters/permanent-delete", ""),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/api/2.0/clusters/edit", ""),
				}},
		},
		{
			Display:        "dbfs",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/dbfs/list", ""),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/dbfs/delete", ""),
		},
		{
			Display:  "groups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/groups/list", ""),
		},
		{
			Display:  "instance-pools",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/instance-pools/list", ""),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{instance_pool_id}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/instance-pools/get", ""),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/instance-pools/delete", ""),
				}},
		},
		{
			Display:  "jobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/list", ""),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{job_id}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/get", ""),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/delete", ""),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/reset", ""),
					Children: []swagger.ResourceType{
						{
							Display:  "runs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/list", ""),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{run_id}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/get", ""),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/delete", ""),
								}},
						}},
				}},
		},
		{
			Display:  "runs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/list", ""),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{run_id}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/get", ""),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/jobs/runs/delete", ""),
				}},
		},
		{
			Display:  "secrets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/scopes/list", ""),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{scope}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/{scope}", ""),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/scopes/delete", ""),
					FixedContent:   "Choose a node to expand...",
					Children: []swagger.ResourceType{
						{
							Display:        "acls",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/acls/list", ""),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/acls/delete", ""),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/acls/put", ""),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{principal}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/acls/get", ""),
								}},
						},
						{
							Display:  "secrets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/secrets/list", ""),
						}},
				}},
		},
		{
			Display:  "token",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/token/list", ""),
		},
		{
			Display:        "workspace",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/api/2.0/workspace/list", ""),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/api/2.0/workspace/delete", ""),
		}}

}
