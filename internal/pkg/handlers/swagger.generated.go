package handlers

func (e *SwaggerResourceExpander) getResourceTypes() []ResourceType {
	return []ResourceType{
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.CognitiveServices/operations", "2017-04-18"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus", "2017-04-18"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
			SubResources: []ResourceType{
				{
					Display:  "{accountName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					Children: []ResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus", "2017-04-18"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/usages", "2017-04-18"),
						},
					},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Compute/operations", "2018-10-01"),
		},
		{
			Display:  "availabilitySets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets", "2018-10-01"),
		},
		{
			Display:  "images",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images", "2018-10-01"),
		},
		{
			Display:  "publishers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "types",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "versions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{version}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}", "2018-10-01"),
								},
							},
						},
					},
				},
				{
					Display:  "offers",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "versions",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions", "2018-10-01"),
									SubResources: []ResourceType{
										{
											Display:  "{version}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}", "2018-10-01"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages", "2018-10-01"),
		},
		{
			Display:  "virtualMachines",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/virtualMachines", "2018-10-01"),
		},
		{
			Display:  "vmSizes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes", "2018-10-01"),
		},
		{
			Display:  "virtualMachineScaleSets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachineScaleSets", "2018-10-01"),
		},
		{
			Display:  "virtualMachines",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines", "2018-10-01"),
		},
		{
			Display:  "availabilitySets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{availabilitySetName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "vmSizes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "images",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{imageName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "virtualMachineScaleSets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "virtualMachines",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "publicipaddresses",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", "2017-03-30"),
							SubResources: []ResourceType{
								{
									Display:  "{publicIpAddressName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", "2017-03-30"),
								},
							},
						},
					},
				},
				{
					Display:  "{vmScaleSetName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "extensions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{vmssExtensionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "instanceView",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/instanceView", "2018-10-01"),
						},
						{
							Display:  "osUpgradeHistory",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osUpgradeHistory", "2018-10-01"),
						},
						{
							Display:  "latest",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest", "2018-10-01"),
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/skus", "2018-10-01"),
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "{instanceId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}", "2018-10-01"),
							Children: []ResourceType{
								{
									Display:  "instanceView",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/instanceView", "2018-10-01"),
								},
							},
						},
					},
				},
				{
					Display:  "publicipaddresses",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses", "2017-03-30"),
				},
			},
		},
		{
			Display:  "virtualMachines",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{vmName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "extensions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{vmExtensionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "instanceView",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/instanceView", "2018-10-01"),
						},
						{
							Display:  "vmSizes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/vmSizes", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "runCommands",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{commandId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ContainerInstance/operations", "2018-10-01"),
		},
		{
			Display:  "containerGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
		},
		{
			Display:  "cachedImages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages", "2018-10-01"),
		},
		{
			Display:  "capabilities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities", "2018-10-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages", "2018-10-01"),
		},
		{
			Display:  "containerGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{containerGroupName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					Children: []ResourceType{},
					SubResources: []ResourceType{
						{
							Display:  "logs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "runs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs", "2018-09-01"),
			SubResources: []ResourceType{
				{
					Display:  "{runId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}", "2018-09-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "tasks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks", "2018-09-01"),
			SubResources: []ResourceType{
				{
					Display:  "{taskName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}", "2018-09-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ContainerService/operations", "2018-03-31"),
		},
		{
			Display:  "managedClusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters", "2018-03-31"),
		},
		{
			Display:  "managedClusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters", "2018-03-31"),
			SubResources: []ResourceType{
				{
					Display:  "{resourceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2018-03-31"),
					Children: []ResourceType{
						{
							Display:  "default",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeProfiles/default", "2018-03-31"),
						},
					},
					SubResources: []ResourceType{},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DocumentDB/operations", "2015-04-08"),
		},
		{
			Display:  "databaseAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts", "2015-04-08"),
		},
		{
			Display:  "databaseAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts", "2015-04-08"),
			SubResources: []ResourceType{
				{
					Display:  "{accountName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2015-04-08"),
					Children: []ResourceType{
						{
							Display:  "metricDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metricDefinitions", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/percentile/metrics", "2015-04-08"),
						},
						{
							Display:  "readonlykeys",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/readonlykeys", "2015-04-08"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/usages", "2015-04-08"),
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "metricDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2015-04-08"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages", "2015-04-08"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages", "2015-04-08"),
						},
						{
							Display:  "metricDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metricDefinitions", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metrics", "2015-04-08"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/usages", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sourceRegion/{sourceRegion}/targetRegion/{targetRegion}/percentile/metrics", "2015-04-08"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/targetRegion/{targetRegion}/percentile/metrics", "2015-04-08"),
						},
					},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.EventHub/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
		},
		{
			Display:  "regions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
			SubResources: []ResourceType{
				{
					Display:  "{namespaceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					Children: []ResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []ResourceType{
								{
									Display:  "{authorizationRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "disasterRecoveryConfigs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
							Children: []ResourceType{},
							SubResources: []ResourceType{
								{
									Display:  "{alias}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									Children: []ResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
											SubResources: []ResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []ResourceType{},
												},
											},
										},
									},
								},
							},
						},
						{
							Display:  "eventhubs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
							SubResources: []ResourceType{
								{
									Display:  "{eventHubName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									Children: []ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules", "2017-04-01"),
											SubResources: []ResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []ResourceType{},
												},
											},
										},
										{
											Display:  "consumergroups",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", "2017-04-01"),
											SubResources: []ResourceType{
												{
													Display:  "{consumerGroupName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
												},
											},
										},
									},
								},
							},
						},
						{
							Display:  "messagingplan",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/messagingplan", "2017-04-01"),
						},
					},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Insights/operations", "2015-05-01"),
		},
		{
			Display:  "{scopePath}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}", "2015-05-01"),
			Children: []ResourceType{
				{
					Display:  "item",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
				},
			},
		},
		{
			Display:  "Annotations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{annotationId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}", "2015-05-01"),
				},
			},
		},
		{
			Display:  "{keyId}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}", "2015-05-01"),
		},
		{
			Display:  "ApiKeys",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys", "2015-05-01"),
		},
		{
			Display:  "exportconfiguration",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{exportId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
				},
			},
		},
		{
			Display:  "currentbillingfeatures",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", "2015-05-01"),
		},
		{
			Display:  "featurecapabilities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/featurecapabilities", "2015-05-01"),
		},
		{
			Display:  "getavailablebillingfeatures",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/getavailablebillingfeatures", "2015-05-01"),
		},
		{
			Display:  "quotastatus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/quotastatus", "2015-05-01"),
		},
		{
			Display:  "ProactiveDetectionConfigs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{ConfigurationId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}", "2015-05-01"),
				},
			},
		},
		{
			Display:  "DefaultWorkItemConfig",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/DefaultWorkItemConfig", "2015-05-01"),
		},
		{
			Display:      "WorkItemConfigs",
			Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs", "2015-05-01"),
			SubResources: []ResourceType{},
		},
		{
			Display:  "components",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components", "2015-05-01"),
		},
		{
			Display:  "components",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{resourceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					Children: []ResourceType{
						{
							Display:  "favorites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites", "2015-05-01"),
							SubResources: []ResourceType{
								{
									Display:  "{favoriteId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
								},
							},
						},
						{
							Display:  "syntheticmonitorlocations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations", "2015-05-01"),
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "{purgeId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/operations/{purgeId}", "2015-05-01"),
						},
					},
				},
				{
					Display:  "webtests",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests", "2015-05-01"),
				},
			},
		},
		{
			Display:  "webtests",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests", "2015-05-01"),
		},
		{
			Display:  "webtests",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{webTestName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
				},
			},
		},
		{
			Display:  "workbooks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks", "2015-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{resourceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
				},
			},
		},
		{
			Display:  "default",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default", "2018-10-01"),
			Children: []ResourceType{
				{
					Display:  "predefinedPolicies",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "{predefinedPolicyName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/{predefinedPolicyName}", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "applicationGatewayAvailableWafRuleSets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", "2018-10-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", "2018-10-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{applicationGatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2018-10-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups", "2018-10-01"),
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{applicationSecurityGroupName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "availableDelegations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2018-10-01"),
		},
		{
			Display:  "availableDelegations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2018-10-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls", "2018-10-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{azureFirewallName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "azureFirewallFqdnTags",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags", "2018-10-01"),
		},
		{
			Display:  "CheckDnsNameAvailability",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", "2018-10-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans", "2018-10-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{ddosProtectionPlanName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "virtualNetworkAvailableEndpointServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices", "2018-10-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", "2018-10-01"),
		},
		{
			Display:  "expressRouteServiceProviders",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", "2018-10-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{circuitName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "authorizations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{authorizationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "peerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{peeringName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2018-10-01"),
									Children: []ResourceType{
										{
											Display:  "connections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections", "2018-10-01"),
											SubResources: []ResourceType{
												{
													Display:  "{connectionName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2018-10-01"),
												},
											},
										},
										{
											Display:  "stats",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats", "2018-10-01"),
										},
									},
									SubResources: []ResourceType{},
								},
							},
						},
						{
							Display:  "stats",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections", "2018-10-01"),
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{crossConnectionName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "peerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:      "{peeringName}",
									Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2018-10-01"),
									SubResources: []ResourceType{},
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways", "2018-10-01"),
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{expressRouteGatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "expressRouteConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{connectionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts", "2018-10-01"),
			Children: []ResourceType{
				{
					Display:  "ExpressRoutePortsLocations",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "{locationName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{expressRoutePortName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "links",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{linkName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "interfaceEndpoints",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/interfaceEndpoints", "2018-10-01"),
		},
		{
			Display:  "interfaceEndpoints",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/interfaceEndpoints", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{interfaceEndpointName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/interfaceEndpoints/{interfaceEndpointName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "loadBalancers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers", "2018-10-01"),
		},
		{
			Display:  "loadBalancers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{loadBalancerName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "backendAddressPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{backendAddressPoolName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "frontendIPConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{frontendIPConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "inboundNatRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{inboundNatRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "loadBalancingRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{loadBalancingRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "networkInterfaces",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", "2018-10-01"),
						},
						{
							Display:  "outboundRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{outboundRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "probes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{probeName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", "2018-10-01"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{networkInterfaceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "loadBalancers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers", "2018-10-01"),
						},
						{
							Display:  "tapConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{tapConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "networkProfiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles", "2018-10-01"),
		},
		{
			Display:  "networkProfiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{networkProfileName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups", "2018-10-01"),
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{networkSecurityGroupName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "defaultSecurityRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{defaultSecurityRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "securityRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{securityRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "networkWatchers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkWatchers", "2018-10-01"),
		},
		{
			Display:  "networkWatchers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{networkWatcherName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "connectionMonitors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{connectionMonitorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2018-10-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "packetCaptures",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{packetCaptureName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2018-10-01"),
									Children: []ResourceType{},
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Network/operations", "2018-10-01"),
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", "2018-10-01"),
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{publicIpAddressName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes", "2018-10-01"),
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{publicIpPrefixName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "routeFilters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters", "2018-10-01"),
		},
		{
			Display:  "routeFilters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{routeFilterName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "routeFilterRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{ruleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "routeTables",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables", "2018-10-01"),
		},
		{
			Display:  "routeTables",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{routeTableName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "routes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{routeName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "bgpServiceCommunities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities", "2018-10-01"),
		},
		{
			Display:  "ServiceEndpointPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies", "2018-10-01"),
		},
		{
			Display:  "serviceEndpointPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{serviceEndpointPolicyName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "serviceEndpointPolicyDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{serviceEndpointPolicyDefinitionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages", "2018-10-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", "2018-10-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{virtualNetworkName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "CheckIPAddressAvailability",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", "2018-10-01"),
						},
						{
							Display:  "subnets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{subnetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2018-10-01"),
								},
							},
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", "2018-10-01"),
						},
						{
							Display:  "virtualNetworkPeerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{virtualNetworkPeeringName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "connections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{virtualNetworkGatewayConnectionName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "sharedkey",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2018-10-01"),
							Children: []ResourceType{},
						},
					},
				},
			},
		},
		{
			Display:  "localNetworkGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{localNetworkGatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "virtualNetworkGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{virtualNetworkGatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "connections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/connections", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps", "2018-10-01"),
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{tapName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways", "2018-10-01"),
		},
		{
			Display:  "virtualHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs", "2018-10-01"),
		},
		{
			Display:  "virtualWans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans", "2018-10-01"),
		},
		{
			Display:  "vpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways", "2018-10-01"),
		},
		{
			Display:  "vpnSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites", "2018-10-01"),
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{gatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2018-10-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "virtualHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{virtualHubName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "hubVirtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{connectionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "virtualWans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{VirtualWANName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2018-10-01"),
				},
				{
					Display:  "supportedSecurityProviders",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", "2018-10-01"),
				},
				{
					Display:  "p2sVpnServerConfigurations",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations", "2018-10-01"),
					SubResources: []ResourceType{
						{
							Display:  "{p2SVpnServerConfigurationName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", "2018-10-01"),
						},
					},
				},
			},
		},
		{
			Display:  "vpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{gatewayName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2018-10-01"),
					Children: []ResourceType{
						{
							Display:  "vpnConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", "2018-10-01"),
							SubResources: []ResourceType{
								{
									Display:  "{connectionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2018-10-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "vpnSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{vpnSiteName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2018-10-01"),
				},
			},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", "2017-03-30"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", "2017-03-30"),
			SubResources: []ResourceType{
				{
					Display:  "{networkInterfaceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", "2017-03-30"),
					Children: []ResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2017-03-30"),
							SubResources: []ResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2017-03-30"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "dnszones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones", "2018-05-01"),
		},
		{
			Display:  "dnsZones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones", "2018-05-01"),
			SubResources: []ResourceType{
				{
					Display:  "{zoneName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					Children: []ResourceType{
						{
							Display:  "all",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/all", "2018-05-01"),
						},
						{
							Display:  "recordsets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets", "2018-05-01"),
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "{recordType}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}", "2018-05-01"),
							SubResources: []ResourceType{
								{
									Display:  "{relativeRecordSetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.OperationalInsights/operations", "2015-03-20"),
		},
		{
			Display:  "linkTargets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/linkTargets", "2015-03-20"),
		},
		{
			Display:  "{purgeId}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}", "2015-03-20"),
		},
		{
			Display:  "savedSearches",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches", "2015-03-20"),
			SubResources: []ResourceType{
				{
					Display:  "{savedSearchId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					Children: []ResourceType{
						{
							Display:  "results",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}/results", "2015-03-20"),
						},
					},
				},
			},
		},
		{
			Display:  "storageInsightConfigs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs", "2015-03-20"),
			SubResources: []ResourceType{
				{
					Display:  "{storageInsightName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
				},
			},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.SignalRService/operations", "2018-10-01"),
		},
		{
			Display:  "SignalR",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/locations/{location}/usages", "2018-10-01"),
		},
		{
			Display:  "SignalR",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
			SubResources: []ResourceType{
				{
					Display:  "{resourceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR/{resourceName}", "2018-10-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "containers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers", "2018-07-01"),
			SubResources: []ResourceType{
				{
					Display:  "{containerName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2018-07-01"),
					Children: []ResourceType{},
					SubResources: []ResourceType{
						{
							Display:  "{immutabilityPolicyName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2018-07-01"),
						},
					},
				},
			},
		},
		{
			Display:  "{BlobServicesName}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2018-07-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Storage/operations", "2018-07-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages", "2018-07-01"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/skus", "2018-07-01"),
		},
		{
			Display:  "storageAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts", "2018-07-01"),
		},
		{
			Display:  "storageAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts", "2018-07-01"),
			SubResources: []ResourceType{
				{
					Display:  "{accountName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2018-07-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "hostingEnvironments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
		},
		{
			Display:  "hostingEnvironments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "compute",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/compute", "2018-02-01"),
						},
						{
							Display:  "virtualip",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/virtualip", "2018-02-01"),
						},
						{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{diagnosticsName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics/{diagnosticsName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metricdefinitions", "2018-02-01"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metrics", "2018-02-01"),
						},
						{
							Display:  "multiRolePools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "default",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metrics", "2018-02-01"),
										},
										{
											Display:  "skus",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/skus", "2018-02-01"),
										},
										{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/usages", "2018-02-01"),
										},
									},
									SubResources: []ResourceType{
										{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metrics", "2018-02-01"),
										},
									},
								},
							},
						},
						{
							Display:  "operations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/operations", "2018-02-01"),
						},
						{
							Display:  "serverfarms",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/serverfarms", "2018-02-01"),
						},
						{
							Display:  "sites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/sites", "2018-02-01"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/usages", "2018-02-01"),
						},
						{
							Display:  "workerPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{workerPoolName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metrics", "2018-02-01"),
										},
										{
											Display:  "skus",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/skus", "2018-02-01"),
										},
										{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/usages", "2018-02-01"),
										},
									},
									SubResources: []ResourceType{
										{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metrics", "2018-02-01"),
										},
									},
								},
							},
						},
						{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors/{detectorName}", "2018-02-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "serverfarms",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
		},
		{
			Display:  "serverfarms",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "capabilities",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/capabilities", "2018-02-01"),
						},
						{
							Display:  "limit",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionPlanLimits/limit", "2018-02-01"),
						},
						{
							Display:  "hybridConnectionRelays",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionRelays", "2018-02-01"),
						},
						{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metricdefinitions", "2018-02-01"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metrics", "2018-02-01"),
						},
						{
							Display:  "sites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/sites", "2018-02-01"),
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/skus", "2018-02-01"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/usages", "2018-02-01"),
						},
						{
							Display:  "virtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{vnetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "routes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{routeName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
												},
											},
										},
									},
									SubResources: []ResourceType{
										{
											Display:  "{gatewayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
										},
									},
								},
							},
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "{relayName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "sites",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}/sites", "2018-02-01"),
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "certificates",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates", "2018-02-01"),
		},
		{
			Display:  "certificates",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
				},
			},
		},
		{
			Display:  "deletedSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", "2018-02-01"),
		},
		{
			Display:  "detectors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{detectorName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors/{detectorName}", "2018-02-01"),
				},
			},
		},
		{
			Display:  "diagnostics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{diagnosticCategory}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "analyses",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{analysisName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "detectors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{detectorName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors/{detectorName}", "2018-02-01"),
				},
			},
		},
		{
			Display:  "diagnostics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{diagnosticCategory}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "analyses",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{analysisName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
					},
				},
			},
		},
		{
			Display:  "availableStacks",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/availableStacks", "2018-02-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/operations", "2018-02-01"),
		},
		{
			Display:  "availableStacks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", "2018-02-01"),
		},
		{
			Display:      "recommendations",
			Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", "2018-02-01"),
			Children:     []ResourceType{},
			SubResources: []ResourceType{},
		},
		{
			Display:  "recommendationHistory",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", "2018-02-01"),
		},
		{
			Display:  "recommendations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", "2018-02-01"),
			Children: []ResourceType{},
			SubResources: []ResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", "2018-02-01"),
					Children: []ResourceType{},
				},
			},
		},
		{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
		},
		{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
		},
		{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", "2018-02-01"),
			Children: []ResourceType{
				{
					Display:  "default",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", "2018-02-01"),
				},
			},
		},
		{
			Display:  "resourceHealthMetadata",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", "2018-02-01"),
			Children: []ResourceType{
				{
					Display:  "default",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", "2018-02-01"),
				},
			},
		},
		{
			Display:  "web",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
		},
		{
			Display:  "sourcecontrols",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{sourceControlType}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
				},
			},
		},
		{
			Display:  "billingMeters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/billingMeters", "2018-02-01"),
		},
		{
			Display:  "deploymentLocations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deploymentLocations", "2018-02-01"),
		},
		{
			Display:  "geoRegions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions", "2018-02-01"),
		},
		{
			Display:  "premieraddonoffers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers", "2018-02-01"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/skus", "2018-02-01"),
		},
		{
			Display:  "sites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites", "2018-02-01"),
		},
		{
			Display:  "sites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites", "2018-02-01"),
			SubResources: []ResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
					Children: []ResourceType{
						{
							Display:  "analyzeCustomHostname",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/analyzeCustomHostname", "2018-02-01"),
						},
						{
							Display:  "config",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "appsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "authsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "azurestorageaccounts",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "backup",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "connectionstrings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "logs",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
								},
								{
									Display:  "metadata",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "publishingcredentials",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "pushsettings",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
									Verb:     "POST"},
								{
									Display:  "slotConfigNames",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
								},
								{
									Display:  "web",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "snapshots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{snapshotId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots/{snapshotId}", "2018-02-01"),
													Children: []ResourceType{},
												},
											},
										},
									},
								},
							},
						},
						{
							Display:  "continuouswebjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "deployments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{id}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "log",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}/log", "2018-02-01"),
										},
									},
								},
							},
						},
						{
							Display:  "domainOwnershipIdentifiers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{domainOwnershipIdentifierName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "MSDeploy",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "log",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy/log", "2018-02-01"),
								},
							},
						},
						{
							Display:  "functions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "token",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/admin/token", "2018-02-01"),
								},
							},
							SubResources: []ResourceType{
								{
									Display:  "{functionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
									Children: []ResourceType{},
								},
							},
						},
						{
							Display:  "hostNameBindings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{hostName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "hybridConnectionRelays",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionRelays", "2018-02-01"),
						},
						{
							Display:  "hybridconnection",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{entityName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "instances",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "MSDeploy",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "log",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
										},
									},
								},
								{
									Display:  "processes",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes", "2018-02-01"),
									SubResources: []ResourceType{
										{
											Display:  "{processId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
											Children: []ResourceType{
												{
													Display:  "dump",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
												},
												{
													Display:  "modules",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
													SubResources: []ResourceType{
														{
															Display:  "{baseAddress}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
														},
													},
												},
												{
													Display:  "threads",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
													SubResources: []ResourceType{
														{
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
						{
							Display:  "metricdefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metricdefinitions", "2018-02-01"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metrics", "2018-02-01"),
						},
						{
							Display:  "virtualNetwork",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
						},
						{
							Display:  "perfcounters",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/perfcounters", "2018-02-01"),
						},
						{
							Display:  "phplogging",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/phplogging", "2018-02-01"),
						},
						{
							Display:  "premieraddons",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{premierAddOnName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "virtualNetworks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
						},
						{
							Display:  "processes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{processId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "dump",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/dump", "2018-02-01"),
										},
										{
											Display:  "modules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{baseAddress}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "threads",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{threadId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads/{threadId}", "2018-02-01"),
												},
											},
										},
									},
								},
							},
						},
						{
							Display:  "publicCertificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{publicCertificateName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "siteextensions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{siteExtensionId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
								},
							},
						},
						{
							Display:  "slots",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
							Children: []ResourceType{},
							SubResources: []ResourceType{
								{
									Display:  "{slot}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "analyzeCustomHostname",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/analyzeCustomHostname", "2018-02-01"),
										},
										{
											Display:  "config",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
											Children: []ResourceType{
												{
													Display:  "logs",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
												},
												{
													Display:  "web",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
													Children: []ResourceType{
														{
															Display:  "snapshots",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots", "2018-02-01"),
															SubResources: []ResourceType{
																{
																	Display:  "{snapshotId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																	Children: []ResourceType{},
																},
															},
														},
													},
												},
											},
										},
										{
											Display:  "continuouswebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
													Children: []ResourceType{},
												},
											},
										},
										{
											Display:  "deployments",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{id}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
													Children: []ResourceType{
														{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}/log", "2018-02-01"),
														},
													},
												},
											},
										},
										{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{domainOwnershipIdentifierName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "MSDeploy",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
											Children: []ResourceType{
												{
													Display:  "log",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy/log", "2018-02-01"),
												},
											},
										},
										{
											Display:  "functions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions", "2018-02-01"),
											Children: []ResourceType{
												{
													Display:  "token",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/admin/token", "2018-02-01"),
												},
											},
											SubResources: []ResourceType{
												{
													Display:  "{functionName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
													Children: []ResourceType{},
												},
											},
										},
										{
											Display:  "hostNameBindings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{hostName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "hybridConnectionRelays",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionRelays", "2018-02-01"),
										},
										{
											Display:  "hybridconnection",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{entityName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "instances",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "MSDeploy",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													Children: []ResourceType{
														{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
														},
													},
												},
												{
													Display:  "processes",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes", "2018-02-01"),
													SubResources: []ResourceType{
														{
															Display:  "{processId}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															Children: []ResourceType{
																{
																	Display:  "dump",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																},
																{
																	Display:  "modules",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																	SubResources: []ResourceType{
																		{
																			Display:  "{baseAddress}",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																		},
																	},
																},
																{
																	Display:  "threads",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																	SubResources: []ResourceType{
																		{
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
										{
											Display:  "metricdefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metrics", "2018-02-01"),
										},
										{
											Display:  "status",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/migratemysql/status", "2018-02-01"),
										},
										{
											Display:  "virtualNetwork",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
										},
										{
											Display:  "perfcounters",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/perfcounters", "2018-02-01"),
										},
										{
											Display:  "phplogging",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/phplogging", "2018-02-01"),
										},
										{
											Display:  "premieraddons",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{premierAddOnName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "virtualNetworks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
										},
										{
											Display:  "processes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{processId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
													Children: []ResourceType{
														{
															Display:  "dump",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/dump", "2018-02-01"),
														},
														{
															Display:  "modules",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules", "2018-02-01"),
															SubResources: []ResourceType{
																{
																	Display:  "{baseAddress}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																},
															},
														},
														{
															Display:  "threads",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads", "2018-02-01"),
															SubResources: []ResourceType{
																{
																	Display:  "{threadId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																},
															},
														},
													},
												},
											},
										},
										{
											Display:  "publicCertificates",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{publicCertificateName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "siteextensions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{siteExtensionId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
												},
											},
										},
										{
											Display:  "snapshots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshots", "2018-02-01"),
											Children: []ResourceType{
												{
													Display:  "snapshotsdr",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshotsdr", "2018-02-01"),
												},
											},
										},
										{
											Display:  "web",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
										},
										{
											Display:  "triggeredwebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													Children: []ResourceType{
														{
															Display:  "history",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
															SubResources: []ResourceType{
																{
																	Display:  "{id}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																},
															},
														},
													},
												},
											},
										},
										{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/usages", "2018-02-01"),
										},
										{
											Display:  "virtualNetworkConnections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{vnetName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													SubResources: []ResourceType{
														{
															Display:  "{gatewayName}",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														},
													},
												},
											},
										},
										{
											Display:  "webjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs/{webJobName}", "2018-02-01"),
												},
											},
										},
									},
									SubResources: []ResourceType{
										{
											Display:  "{relayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children: []ResourceType{},
										},
										{
											Display:  "{view}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkFeatures/{view}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/operationresults/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/{operationId}", "2018-02-01"),
										},
									},
								},
							},
						},
						{
							Display:  "snapshots",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshots", "2018-02-01"),
							Children: []ResourceType{
								{
									Display:  "snapshotsdr",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshotsdr", "2018-02-01"),
								},
							},
						},
						{
							Display:  "web",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
						},
						{
							Display:  "triggeredwebjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
									Children: []ResourceType{
										{
											Display:  "history",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
											SubResources: []ResourceType{
												{
													Display:  "{id}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
												},
											},
										},
									},
								},
							},
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/usages", "2018-02-01"),
						},
						{
							Display:  "virtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{vnetName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
									SubResources: []ResourceType{
										{
											Display:  "{gatewayName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
										},
									},
								},
							},
						},
						{
							Display:  "webjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs", "2018-02-01"),
							SubResources: []ResourceType{
								{
									Display:  "{webJobName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs/{webJobName}", "2018-02-01"),
								},
							},
						},
					},
					SubResources: []ResourceType{
						{
							Display:  "{relayName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
							Children: []ResourceType{},
						},
						{
							Display:  "{view}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkFeatures/{view}", "2018-02-01"),
						},
						{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/operationresults/{operationId}", "2018-02-01"),
						},
						{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/{operationId}", "2018-02-01"),
						},
						{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
						},
						{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/{operationId}", "2018-02-01"),
						},
					},
				},
			},
		},
	}
}
