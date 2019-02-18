package handlers

func (e *AppServiceResourceExpander) getResourceTypes() []ResourceType {
	return []ResourceType{
		{
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
			Children: []ResourceType{
				{
					Display:  "config",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "appsettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "authsettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "connectionstrings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "logs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "metadata",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "publishingcredentials",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "pushsettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
							Verb:     "POST",
						},
						{
							Display:  "slotConfigNames",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
						},
						{
							Display:  "virtualNetwork",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/virtualNetwork", "2018-02-01"),
						},
						{
							Display:  "web",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
						},
					},
				},
				{
					Display:  "siteextensions",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
					SubResources: []ResourceType{
						{
							Display:  "siteextension: {siteExtensionId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
						},
					},
				},
				{
					Display:  "slots",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
					SubResources: []ResourceType{
						{
							Display:  "slot: {slot}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "config",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "appsettings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/appsettings/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "authsettings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/authsettings/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "connectionstrings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/connectionstrings/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "logs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "metadata",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/metadata/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "publishingcredentials",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/publishingcredentials/list", "2018-02-01"),
											Verb:     "POST",
										},
										{
											Display:  "pushsettings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/pushsettings/list", "2018-02-01"),
										},
										{
											Display:  "slotConfigNames",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/slotConfigNames", "2018-02-01"),
										},
										{
											Display:  "virtualNetwork",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/virtualNetwork", "2018-02-01"),
										},
										{
											Display:  "web",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
										},
									},
								},
								{
									Display:  "siteextensions",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
									SubResources: []ResourceType{
										{
											Display:  "siteextension: {siteExtensionId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
