package handlers

func (e *AppServiceResourceExpander) getResourceTypes() []ResourceType {
	return []ResourceType{
		ResourceType{
			Display:  "hostingEnvironments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
		},
		ResourceType{
			Display:  "hostingEnvironments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
					Children: []ResourceType {
						ResourceType{
							Display:  "compute",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/compute", "2018-02-01"),
						},
						ResourceType{
							Display:  "virtualip",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/virtualip", "2018-02-01"),
						},
						ResourceType{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{diagnosticsName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics/{diagnosticsName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metricdefinitions", "2018-02-01"),
						},
						ResourceType{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metrics", "2018-02-01"),
						},
						ResourceType{
							Display:  "multiRolePools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "default",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metricdefinitions", "2018-02-01"),
										},
										ResourceType{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metrics", "2018-02-01"),
										},
										ResourceType{
											Display:  "skus",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/skus", "2018-02-01"),
										},
										ResourceType{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/usages", "2018-02-01"),
										},
									},
									SubResources: []ResourceType {
										ResourceType{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metricdefinitions", "2018-02-01"),
										},
										ResourceType{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metrics", "2018-02-01"),
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "operations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/operations", "2018-02-01"),
						},
						ResourceType{
							Display:  "serverfarms",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/serverfarms", "2018-02-01"),
						},
						ResourceType{
							Display:  "sites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/sites", "2018-02-01"),
						},
						ResourceType{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/usages", "2018-02-01"),
						},
						ResourceType{
							Display:  "workerPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{workerPoolName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metricdefinitions", "2018-02-01"),
										},
										ResourceType{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metrics", "2018-02-01"),
										},
										ResourceType{
											Display:  "skus",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/skus", "2018-02-01"),
										},
										ResourceType{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/usages", "2018-02-01"),
										},
									},
									SubResources: []ResourceType {
										ResourceType{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metricdefinitions", "2018-02-01"),
										},
										ResourceType{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metrics", "2018-02-01"),
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors/{detectorName}", "2018-02-01"),
								},
							},
						},
					},
				},
			},
		},
		ResourceType{
			Display:  "serverfarms",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
		},
		ResourceType{
			Display:  "serverfarms",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
					Children: []ResourceType {
						ResourceType{
							Display:  "capabilities",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/capabilities", "2018-02-01"),
						},
						ResourceType{
							Display:  "limit",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionPlanLimits/limit", "2018-02-01"),
						},
						ResourceType{
							Display:  "hybridConnectionRelays",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionRelays", "2018-02-01"),
						},
						ResourceType{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metricdefinitions", "2018-02-01"),
						},
						ResourceType{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metrics", "2018-02-01"),
						},
						ResourceType{
							Display:  "sites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/sites", "2018-02-01"),
						},
						ResourceType{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/skus", "2018-02-01"),
						},
						ResourceType{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/usages", "2018-02-01"),
						},
						ResourceType{
							Display:  "virtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{vnetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "routes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{routeName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
												},
											},
										},
									},
									SubResources: []ResourceType {
										ResourceType{
											Display:  "{gatewayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
										},
									},
								},
							},
						},
					},
					SubResources: []ResourceType {
						ResourceType{
							Display:  "{relayName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "sites",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}/sites", "2018-02-01"),
								},
							},
						},
					},
				},
			},
		},
		ResourceType{
			Display:  "certificates",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates", "2018-02-01"),
		},
		ResourceType{
			Display:  "certificates",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "deletedSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", "2018-02-01"),
		},
		ResourceType{
			Display:  "detectors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{detectorName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors/{detectorName}", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "diagnostics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{diagnosticCategory}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}", "2018-02-01"),
					Children: []ResourceType {
						ResourceType{
							Display:  "analyses",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{analysisName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
						ResourceType{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
					},
				},
			},
		},
		ResourceType{
			Display:  "detectors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{detectorName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors/{detectorName}", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "diagnostics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{diagnosticCategory}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}", "2018-02-01"),
					Children: []ResourceType {
						ResourceType{
							Display:  "analyses",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{analysisName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
						ResourceType{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
					},
				},
			},
		},
		ResourceType{
			Display:  "availableStacks",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/availableStacks", "2018-02-01"),
		},
		ResourceType{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/operations", "2018-02-01"),
		},
		ResourceType{
			Display:  "availableStacks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", "2018-02-01"),
		},
		ResourceType{
			Display:  "recommendations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", "2018-02-01"),
			Children: []ResourceType {
			},
			SubResources: []ResourceType {
			},
		},
		ResourceType{
			Display:  "recommendationHistory",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", "2018-02-01"),
		},
		ResourceType{
			Display:  "recommendations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", "2018-02-01"),
			Children: []ResourceType {
			},
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", "2018-02-01"),
					Children: []ResourceType {
					},
				},
			},
		},
		ResourceType{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
		},
		ResourceType{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
		},
		ResourceType{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", "2018-02-01"),
			Children: []ResourceType {
				ResourceType{
					Display:  "default",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", "2018-02-01"),
			Children: []ResourceType {
				ResourceType{
					Display:  "default",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "web",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
		},
		ResourceType{
			Display:  "sourcecontrols",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{sourceControlType}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
				},
			},
		},
		ResourceType{
			Display:  "billingMeters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/billingMeters", "2018-02-01"),
		},
		ResourceType{
			Display:  "deploymentLocations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deploymentLocations", "2018-02-01"),
		},
		ResourceType{
			Display:  "geoRegions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions", "2018-02-01"),
		},
		ResourceType{
			Display:  "premieraddonoffers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers", "2018-02-01"),
		},
		ResourceType{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/skus", "2018-02-01"),
		},
		ResourceType{
			Display:  "sites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites", "2018-02-01"),
		},
		ResourceType{
			Display:  "sites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites", "2018-02-01"),
			SubResources: []ResourceType {
				ResourceType{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
					Children: []ResourceType {
						ResourceType{
							Display:  "analyzeCustomHostname",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/analyzeCustomHostname", "2018-02-01"),
						},
						ResourceType{
							Display:  "config",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "appsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "authsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "azurestorageaccounts",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "backup",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "connectionstrings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "logs",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
								},
								ResourceType{
									Display:  "metadata",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "publishingcredentials",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "pushsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
									Verb:     "POST"								},
								ResourceType{
									Display:  "slotConfigNames",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
								},
								ResourceType{
									Display:  "web",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "snapshots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{snapshotId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots/{snapshotId}", "2018-02-01"),
													Children: []ResourceType {
													},
												},
											},
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "continuouswebjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
						ResourceType{
							Display:  "deployments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{id}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "log",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}/log", "2018-02-01"),
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "domainOwnershipIdentifiers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{domainOwnershipIdentifierName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "MSDeploy",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "log",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy/log", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "functions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "token",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/admin/token", "2018-02-01"),
								},
							},
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{functionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
									Children: []ResourceType {
									},
								},
							},
						},
						ResourceType{
							Display:  "hostNameBindings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{hostName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "hybridConnectionRelays",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionRelays", "2018-02-01"),
						},
						ResourceType{
							Display:  "hybridconnection",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{entityName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "instances",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "MSDeploy",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "log",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
										},
									},
								},
								ResourceType{
									Display:  "processes",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes", "2018-02-01"),
									SubResources: []ResourceType {
										ResourceType{
											Display:  "{processId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
											Children: []ResourceType {
												ResourceType{
													Display:  "dump",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
												},
												ResourceType{
													Display:  "modules",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
													SubResources: []ResourceType {
														ResourceType{
															Display:  "{baseAddress}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
														},
													},
												},
												ResourceType{
													Display:  "threads",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
													SubResources: []ResourceType {
														ResourceType{
															Display:  "{threadId}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
														},
													},
												},
											},
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metricdefinitions", "2018-02-01"),
						},
						ResourceType{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metrics", "2018-02-01"),
						},
						ResourceType{
							Display:  "virtualNetwork",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
						},
						ResourceType{
							Display:  "perfcounters",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/perfcounters", "2018-02-01"),
						},
						ResourceType{
							Display:  "phplogging",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/phplogging", "2018-02-01"),
						},
						ResourceType{
							Display:  "premieraddons",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{premierAddOnName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "virtualNetworks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
						},
						ResourceType{
							Display:  "processes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{processId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "dump",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/dump", "2018-02-01"),
										},
										ResourceType{
											Display:  "modules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{baseAddress}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "threads",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{threadId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads/{threadId}", "2018-02-01"),
												},
											},
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "publicCertificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{publicCertificateName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "siteextensions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{siteExtensionId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "slots",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
							Children: []ResourceType {
							},
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{slot}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "analyzeCustomHostname",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/analyzeCustomHostname", "2018-02-01"),
										},
										ResourceType{
											Display:  "config",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
											Children: []ResourceType {
												ResourceType{
													Display:  "logs",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
												},
												ResourceType{
													Display:  "web",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
													Children: []ResourceType {
														ResourceType{
															Display:  "snapshots",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots", "2018-02-01"),
															SubResources: []ResourceType {
																ResourceType{
																	Display:  "{snapshotId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																	Children: []ResourceType {
																	},
																},
															},
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "continuouswebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
													Children: []ResourceType {
													},
												},
											},
										},
										ResourceType{
											Display:  "deployments",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{id}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
													Children: []ResourceType {
														ResourceType{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}/log", "2018-02-01"),
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{domainOwnershipIdentifierName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "MSDeploy",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
											Children: []ResourceType {
												ResourceType{
													Display:  "log",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy/log", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "functions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions", "2018-02-01"),
											Children: []ResourceType {
												ResourceType{
													Display:  "token",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/admin/token", "2018-02-01"),
												},
											},
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{functionName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
													Children: []ResourceType {
													},
												},
											},
										},
										ResourceType{
											Display:  "hostNameBindings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{hostName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "hybridConnectionRelays",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionRelays", "2018-02-01"),
										},
										ResourceType{
											Display:  "hybridconnection",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{entityName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "instances",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "MSDeploy",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													Children: []ResourceType {
														ResourceType{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
														},
													},
												},
												ResourceType{
													Display:  "processes",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes", "2018-02-01"),
													SubResources: []ResourceType {
														ResourceType{
															Display:  "{processId}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															Children: []ResourceType {
																ResourceType{
																	Display:  "dump",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																},
																ResourceType{
																	Display:  "modules",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																	SubResources: []ResourceType {
																		ResourceType{
																			Display:  "{baseAddress}",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																		},
																	},
																},
																ResourceType{
																	Display:  "threads",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																	SubResources: []ResourceType {
																		ResourceType{
																			Display:  "{threadId}",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metricdefinitions", "2018-02-01"),
										},
										ResourceType{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metrics", "2018-02-01"),
										},
										ResourceType{
											Display:  "status",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/migratemysql/status", "2018-02-01"),
										},
										ResourceType{
											Display:  "virtualNetwork",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
										},
										ResourceType{
											Display:  "perfcounters",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/perfcounters", "2018-02-01"),
										},
										ResourceType{
											Display:  "phplogging",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/phplogging", "2018-02-01"),
										},
										ResourceType{
											Display:  "premieraddons",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{premierAddOnName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "virtualNetworks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
										},
										ResourceType{
											Display:  "processes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{processId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
													Children: []ResourceType {
														ResourceType{
															Display:  "dump",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/dump", "2018-02-01"),
														},
														ResourceType{
															Display:  "modules",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules", "2018-02-01"),
															SubResources: []ResourceType {
																ResourceType{
																	Display:  "{baseAddress}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																},
															},
														},
														ResourceType{
															Display:  "threads",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads", "2018-02-01"),
															SubResources: []ResourceType {
																ResourceType{
																	Display:  "{threadId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																},
															},
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "publicCertificates",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{publicCertificateName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "siteextensions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{siteExtensionId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "snapshots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshots", "2018-02-01"),
											Children: []ResourceType {
												ResourceType{
													Display:  "snapshotsdr",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshotsdr", "2018-02-01"),
												},
											},
										},
										ResourceType{
											Display:  "web",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
										},
										ResourceType{
											Display:  "triggeredwebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													Children: []ResourceType {
														ResourceType{
															Display:  "history",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
															SubResources: []ResourceType {
																ResourceType{
																	Display:  "{id}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																},
															},
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/usages", "2018-02-01"),
										},
										ResourceType{
											Display:  "virtualNetworkConnections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{vnetName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													SubResources: []ResourceType {
														ResourceType{
															Display:  "{gatewayName}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														},
													},
												},
											},
										},
										ResourceType{
											Display:  "webjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs/{webJobName}", "2018-02-01"),
												},
											},
										},
									},
									SubResources: []ResourceType {
										ResourceType{
											Display:  "{relayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children: []ResourceType {
											},
										},
										ResourceType{
											Display:  "{view}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkFeatures/{view}", "2018-02-01"),
										},
										ResourceType{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/operationresults/{operationId}", "2018-02-01"),
										},
										ResourceType{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/{operationId}", "2018-02-01"),
										},
										ResourceType{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
										},
										ResourceType{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/{operationId}", "2018-02-01"),
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "snapshots",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshots", "2018-02-01"),
							Children: []ResourceType {
								ResourceType{
									Display:  "snapshotsdr",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshotsdr", "2018-02-01"),
								},
							},
						},
						ResourceType{
							Display:  "web",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
						},
						ResourceType{
							Display:  "triggeredwebjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
									Children: []ResourceType {
										ResourceType{
											Display:  "history",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
											SubResources: []ResourceType {
												ResourceType{
													Display:  "{id}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
												},
											},
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/usages", "2018-02-01"),
						},
						ResourceType{
							Display:  "virtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{vnetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
									SubResources: []ResourceType {
										ResourceType{
											Display:  "{gatewayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
										},
									},
								},
							},
						},
						ResourceType{
							Display:  "webjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs", "2018-02-01"),
							SubResources: []ResourceType {
								ResourceType{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs/{webJobName}", "2018-02-01"),
								},
							},
						},
					},
					SubResources: []ResourceType {
						ResourceType{
							Display:  "{relayName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
							Children: []ResourceType {
							},
						},
						ResourceType{
							Display:  "{view}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkFeatures/{view}", "2018-02-01"),
						},
						ResourceType{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/operationresults/{operationId}", "2018-02-01"),
						},
						ResourceType{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/{operationId}", "2018-02-01"),
						},
						ResourceType{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
						},
						ResourceType{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/{operationId}", "2018-02-01"),
						},
					},
				},
			},
		},
	}
}
