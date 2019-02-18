package handlers

func (e *AppServiceResourceExpander) getHandledTypes() []handledType {
	return []handledType{
		{
			endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
			children: []handledType{
				{
					display:  "config",
					endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
					children: []handledType{
						{
							display:  "appsettings",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "authsettings",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "connectionstrings",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "logs",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "metadata",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "publishingcredentials",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "pushsettings",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
							verb:     "POST",
						},
						{
							display:  "slotConfigNames",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
						},
						{
							display:  "virtualNetwork",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/virtualNetwork", "2018-02-01"),
						},
						{
							display:  "web",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
						},
					},
				},
				{
					display:  "siteextensions",
					endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
					subResources: []handledType{
						{
							display:  "siteextension: {siteExtensionId}",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
						},
					},
				},
				{
					display:  "slots",
					endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
					subResources: []handledType{
						{
							display:  "slot: {slot}",
							endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
							children: []handledType{
								{
									display:  "config",
									endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
									children: []handledType{
										{
											display:  "appsettings",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/appsettings/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "authsettings",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/authsettings/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "connectionstrings",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/connectionstrings/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "logs",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "metadata",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/metadata/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "publishingcredentials",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/publishingcredentials/list", "2018-02-01"),
											verb:     "POST",
										},
										{
											display:  "pushsettings",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/pushsettings/list", "2018-02-01"),
										},
										{
											display:  "slotConfigNames",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/slotConfigNames", "2018-02-01"),
										},
										{
											display:  "virtualNetwork",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/virtualNetwork", "2018-02-01"),
										},
										{
											display:  "web",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
										},
									},
								},
								{
									display:  "siteextensions",
									endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
									subResources: []handledType{
										{
											display:  "siteextension: {siteExtensionId}",
											endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
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
