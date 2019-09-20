package handlers

func (e *SwaggerResourceExpander) getResourceTypes() []SwaggerResourceType {
	return []SwaggerResourceType{
		{
			Display:  "addsservices",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices", "2014-01-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "premiumCheck",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/premiumCheck", "2014-01-01"),
				}},
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "addomainservicemembers",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addomainservicemembers", "2014-01-01"),
						},
						{
							Display:  "addsservicemembers",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addsservicemembers", "2014-01-01"),
						},
						{
							Display:  "alerts",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/alerts", "2014-01-01"),
						},
						{
							Display:  "configuration",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/configuration", "2014-01-01"),
						},
						{
							Display:  "forestsummary",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/forestsummary", "2014-01-01"),
						},
						{
							Display:  "metricmetadata",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata", "2014-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{metricName}",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{groupName}",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
										}},
								}},
						},
						{
							Display:  "replicationdetails",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationdetails", "2014-01-01"),
						},
						{
							Display:  "replicationstatus",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationstatus", "2014-01-01"),
						},
						{
							Display:  "replicationsummary",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationsummary", "2014-01-01"),
						},
						{
							Display:  "servicemembers",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers", "2014-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{serviceMemberId}",
									Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "alerts",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
										},
										{
											Display:  "credentials",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
										}},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{dimension}",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/dimensions/{dimension}", "2014-01-01"),
						},
						{
							Display:        "userpreference",
							Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
						},
						{
							Display:  "{groupName}",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "average",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
								},
								{
									Display:  "sum",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
								}},
						}},
				}},
		},
		{
			Display:       "configuration",
			Endpoint:      mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
			PatchEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/operations", "2014-01-01"),
		},
		{
			Display:  "IsDevOps",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/reports/DevOps/IsDevOps", "2014-01-01"),
		},
		{
			Display:  "connectors",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/service/{serviceName}/servicemembers/{serviceMemberId}/connectors", "2014-01-01"),
		},
		{
			Display:  "services",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services", "2014-01-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "premiumCheck",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/premiumCheck", "2014-01-01"),
				}},
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "alerts",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/alerts", "2014-01-01"),
						},
						{
							Display:  "counts",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/counts", "2014-01-01"),
						},
						{
							Display:  "listV2",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/listV2", "2014-01-01"),
						},
						{
							Display:  "exportstatus",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exportstatus", "2014-01-01"),
						},
						{
							Display:  "metricmetadata",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata", "2014-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{metricName}",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{groupName}",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
										}},
								}},
						},
						{
							Display:  "user",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/badpassword/details/user", "2014-01-01"),
						},
						{
							Display:  "blobUris",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/riskyIp/blobUris", "2014-01-01"),
						},
						{
							Display:  "servicemembers",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers", "2014-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{serviceMemberId}",
									Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "alerts",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
										},
										{
											Display:  "credentials",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
										},
										{
											Display:  "exportstatus",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/exportstatus", "2014-01-01"),
										},
										{
											Display:  "globalconfiguration",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/globalconfiguration", "2014-01-01"),
										},
										{
											Display:  "serviceconfiguration",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/serviceconfiguration", "2014-01-01"),
										}},
									SubResources: []SwaggerResourceType{
										{
											Display:  "{metricName}",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}", "2014-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{groupName}",
													Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
												}},
										}},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{featureName}",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/TenantWhitelisting/{featureName}", "2014-01-01"),
						},
						{
							Display:  "{featureName}",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/checkServiceFeatureAvailibility/{featureName}", "2014-01-01"),
						},
						{
							Display:  "alertfeedback",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/feedbacktype/alerts/{shortName}/alertfeedback", "2014-01-01"),
						},
						{
							Display:  "{groupName}",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "average",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
								},
								{
									Display:  "sum",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
								}},
						}},
				}},
		},
		{
			Display:  "metadata",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata", "2017-04-19"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{name}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata/{name}", "2017-04-19"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/operations", "2017-04-19"),
		},
		{
			Display:     "configurations",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
		},
		{
			Display:  "recommendations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations", "2017-04-19"),
		},
		{
			Display:  "suppressions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions", "2017-04-19"),
		},
		{
			Display:     "configurations",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
		},
		{
			Display:  "{recommendationId}",
			Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}", "2017-04-19"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
				}},
		},
		{
			Display:  "smartDetectorAlertRules",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/microsoft.alertsManagement/smartDetectorAlertRules", "2019-06-01"),
		},
		{
			Display:  "smartDetectorAlertRules",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{alertRuleName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.AnalysisServices/operations", "2017-08-01"),
		},
		{
			Display:  "{operationId}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationresults/{operationId}", "2017-08-01"),
		},
		{
			Display:  "{operationId}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationstatuses/{operationId}", "2017-08-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/servers", "2017-08-01"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/skus", "2017-08-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers", "2017-08-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/skus", "2017-08-01"),
						}},
				}},
		},
		{
			Display:  "apis",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis", "2019-01-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "apisByTags",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apisByTags", "2019-01-01"),
				}},
			SubResources: []SwaggerResourceType{
				{
					Display:        "{apiId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{diagnosticId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
								}},
						},
						{
							Display:  "issues",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{issueId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "attachments",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{attachmentId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
												}},
										},
										{
											Display:  "comments",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{commentId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "operations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations", "2019-01-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "operationsByTags",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags", "2019-01-01"),
								}},
							SubResources: []SwaggerResourceType{
								{
									Display:        "{operationId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "policies",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{policyId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
												}},
										},
										{
											Display:  "tags",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{tagId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "policies",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{policyId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
								}},
						},
						{
							Display:  "products",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products", "2019-01-01"),
						},
						{
							Display:  "releases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{releaseId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
								}},
						},
						{
							Display:  "revisions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/revisions", "2019-01-01"),
						},
						{
							Display:  "schemas",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{schemaId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
								}},
						},
						{
							Display:  "tagDescriptions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
								}},
						},
						{
							Display:  "tags",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
								}},
						}},
				}},
		},
		{
			Display:  "apiVersionSets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{versionSetId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
				}},
		},
		{
			Display:  "authorizationServers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{authsid}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
				}},
		},
		{
			Display:  "backends",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{backendId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "caches",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{cacheId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
				}},
		},
		{
			Display:  "certificates",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{certificateId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ApiManagement/operations", "2019-01-01"),
		},
		{
			Display:  "service",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/service", "2019-01-01"),
		},
		{
			Display:  "service",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus", "2019-01-01"),
						},
						{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{diagnosticId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
								}},
						},
						{
							Display:  "templates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{templateName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
								}},
						},
						{
							Display:  "groups",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{groupId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:      "users",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", "2019-01-01"),
											SubResources: []SwaggerResourceType{},
										}},
								}},
						},
						{
							Display:  "identityProviders",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{identityProviderName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
								}},
						},
						{
							Display:  "issues",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{issueId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues/{issueId}", "2019-01-01"),
								}},
						},
						{
							Display:  "loggers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{loggerId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
								}},
						},
						{
							Display:  "networkstatus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus", "2019-01-01"),
						},
						{
							Display:  "notifications",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{notificationName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-01-01"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:      "recipientEmails",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails", "2019-01-01"),
											SubResources: []SwaggerResourceType{},
										},
										{
											Display:      "recipientUsers",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers", "2019-01-01"),
											SubResources: []SwaggerResourceType{},
										}},
								}},
						},
						{
							Display:  "openidConnectProviders",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{opid}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
								}},
						},
						{
							Display:  "policies",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{policyId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
								}},
						},
						{
							Display:  "policySnippets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policySnippets", "2019-01-01"),
						},
						{
							Display:       "delegation",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
							PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
						},
						{
							Display:       "signin",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
							PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
						},
						{
							Display:       "signup",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
							PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
						},
						{
							Display:  "products",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products", "2019-01-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "productsByTags",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/productsByTags", "2019-01-01"),
								}},
							SubResources: []SwaggerResourceType{
								{
									Display:        "{productId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:      "apis",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis", "2019-01-01"),
											SubResources: []SwaggerResourceType{},
										},
										{
											Display:      "groups",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups", "2019-01-01"),
											SubResources: []SwaggerResourceType{},
										},
										{
											Display:  "policies",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{policyId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
												}},
										},
										{
											Display:  "subscriptions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/subscriptions", "2019-01-01"),
										},
										{
											Display:  "tags",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags", "2019-01-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{tagId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "properties",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{propId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
								}},
						},
						{
							Display:  "regions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/regions", "2019-01-01"),
						},
						{
							Display:  "byApi",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byApi", "2019-01-01"),
						},
						{
							Display:  "byGeo",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byGeo", "2019-01-01"),
						},
						{
							Display:  "byOperation",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byOperation", "2019-01-01"),
						},
						{
							Display:  "byProduct",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byProduct", "2019-01-01"),
						},
						{
							Display:  "byRequest",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byRequest", "2019-01-01"),
						},
						{
							Display:  "bySubscription",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/bySubscription", "2019-01-01"),
						},
						{
							Display:  "byTime",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byTime", "2019-01-01"),
						},
						{
							Display:  "byUser",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byUser", "2019-01-01"),
						},
						{
							Display:  "subscriptions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{sid}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "tagResources",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tagResources", "2019-01-01"),
						},
						{
							Display:  "tags",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
								}},
						},
						{
							Display:  "users",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{userId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "groups",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/groups", "2019-01-01"),
										},
										{
											Display:  "identities",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/identities", "2019-01-01"),
										},
										{
											Display:  "subscriptions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/subscriptions", "2019-01-01"),
										}},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "networkstatus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus", "2019-01-01"),
						},
						{
							Display:       "{quotaCounterKey}",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-01-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-01-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:       "{quotaPeriodKey}",
									Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-01-01"),
									PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-01-01"),
								}},
						},
						{
							Display:       "{accessName}",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-01-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-01-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "git",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git", "2019-01-01"),
									Children: []SwaggerResourceType{},
								}},
						},
						{
							Display:  "syncState",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/syncState", "2019-01-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Insights/operations", "2015-05-01"),
		},
		{
			Display:  "{scopePath}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}", "2015-05-01"),
			Children: []SwaggerResourceType{
				{
					Display:        "item",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
				}},
		},
		{
			Display:     "Annotations",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations", "2015-05-01"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{annotationId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}", "2015-05-01"),
				}},
		},
		{
			Display:        "{keyId}",
			Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}", "2015-05-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}", "2015-05-01"),
		},
		{
			Display:  "ApiKeys",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys", "2015-05-01"),
		},
		{
			Display:  "exportconfiguration",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{exportId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
				}},
		},
		{
			Display:     "currentbillingfeatures",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", "2015-05-01"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", "2015-05-01"),
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
			SubResources: []SwaggerResourceType{
				{
					Display:     "{ConfigurationId}",
					Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}", "2015-05-01"),
					PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}", "2015-05-01"),
				}},
		},
		{
			Display:  "DefaultWorkItemConfig",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/DefaultWorkItemConfig", "2015-05-01"),
		},
		{
			Display:  "WorkItemConfigs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workItemConfigId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
				}},
		},
		{
			Display:  "components",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components", "2015-05-01"),
		},
		{
			Display:  "components",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "favorites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites", "2015-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{favoriteId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
								}},
						},
						{
							Display:  "syntheticmonitorlocations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations", "2015-05-01"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{purgeId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/operations/{purgeId}", "2015-05-01"),
						}},
				},
				{
					Display:  "webtests",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests", "2015-05-01"),
				}},
		},
		{
			Display:  "webtests",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests", "2015-05-01"),
		},
		{
			Display:  "webtests",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{webTestName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
				}},
		},
		{
			Display:  "workbooks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks", "2015-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
				}},
		},
		{
			Display:  "classicAdministrators",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/classicAdministrators", "2015-07-01"),
		},
		{
			Display:  "providerOperations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations", "2015-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{resourceProviderNamespace}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}", "2015-07-01"),
				}},
		},
		{
			Display:  "roleAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:        "{roleAssignmentId}",
			Endpoint:       mustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{roleAssignmentName}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
				}},
		},
		{
			Display:  "permissions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
		},
		{
			Display:  "permissions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
		},
		{
			Display:  "{roleDefinitionId}",
			Endpoint: mustGetEndpointInfoFromURL("/{roleDefinitionId}", "2015-07-01"),
		},
		{
			Display:  "roleDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions", "2015-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{roleDefinitionId}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
				}},
		},
		{
			Display:  "python2Packages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages", "2018-06-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{packageName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
				}},
		},
		{
			Display:  "runbooks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks", "2018-06-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{runbookName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					Children: []SwaggerResourceType{
						{
							Display:  "content",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/content", "2018-06-30"),
						},
						{
							Display:  "draft",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", "2018-06-30"),
							Children: []SwaggerResourceType{
								{
									Display:     "content",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2018-06-30"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2018-06-30"),
								},
								{
									Display:     "testJob",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2018-06-30"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2018-06-30"),
									Children: []SwaggerResourceType{
										{
											Display:  "streams",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", "2018-06-30"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{jobStreamId}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", "2018-06-30"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Kusto/operations", "2019-01-21"),
		},
		{
			Display:  "clusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters", "2019-01-21"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/skus", "2019-01-21"),
		},
		{
			Display:  "clusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters", "2019-01-21"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{clusterName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-01-21"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-01-21"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-01-21"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-01-21"),
					Children: []SwaggerResourceType{
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases", "2019-01-21"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-01-21"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-01-21"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-01-21"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-01-21"),
									Children: []SwaggerResourceType{
										{
											Display:  "dataConnections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections", "2019-01-21"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{dataConnectionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-01-21"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-01-21"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-01-21"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-01-21"),
												}},
										}},
								}},
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/skus", "2019-01-21"),
						}},
				}},
		},
		{
			Display:  "diagnosticSettings",
			Endpoint: mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings", "2017-04-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "diagnosticSettingsCategories",
					Endpoint: mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettingsCategories", "2017-04-01"),
				}},
			SubResources: []SwaggerResourceType{
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/microsoft.aadiam/operations", "2017-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.AzureStack/operations", "2017-06-01"),
		},
		{
			Display:  "customerSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions", "2017-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{customerSubscriptionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
				}},
		},
		{
			Display:  "products",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products", "2017-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{productName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}", "2017-06-01"),
					Children: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "registrations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations", "2017-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{registrationName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Batch/operations", "2019-04-01"),
		},
		{
			Display:  "batchAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts", "2019-04-01"),
		},
		{
			Display:  "quotas",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/quotas", "2019-04-01"),
		},
		{
			Display:  "batchAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "applications",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{applicationName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "versions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions", "2019-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{versionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "certificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "pools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{poolName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.BatchAI/operations", "2018-05-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/locations/{location}/usages", "2018-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "clusters",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters", "2018-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{clusterName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "experiments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments", "2018-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{experimentName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "jobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs", "2018-05-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{jobName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "fileServers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers", "2018-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{fileServerName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
								}},
						}},
				}},
		},
		{
			Display:  "edgenodes",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/edgenodes", "2019-04-15"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/operations", "2019-04-15"),
		},
		{
			Display:  "profiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles", "2019-04-15"),
		},
		{
			Display:  "profiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles", "2019-04-15"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{profileName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					Children: []SwaggerResourceType{
						{
							Display:  "endpoints",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints", "2019-04-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{endpointName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									Children: []SwaggerResourceType{
										{
											Display:  "customDomains",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains", "2019-04-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{customDomainName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "origins",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins", "2019-04-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:       "{originName}",
													Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-04-15"),
													PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-04-15"),
												}},
										}},
								}},
						}},
				}},
		},
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
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					Children: []SwaggerResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus", "2017-04-18"),
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/usages", "2017-04-18"),
						}},
				}},
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/skus", "2019-04-01"),
		},
		{
			Display:  "containerServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices", "2017-01-31"),
		},
		{
			Display:  "containerServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices", "2017-01-31"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{containerServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
				}},
		},
		{
			Display:  "balances",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/balances", "2019-01-01"),
		},
		{
			Display:  "balances",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/balances", "2019-01-01"),
		},
		{
			Display:  "reservationDetails",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationDetails", "2019-01-01"),
		},
		{
			Display:  "reservationSummaries",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationSummaries", "2019-01-01"),
		},
		{
			Display:  "reservationDetails",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationDetails", "2019-01-01"),
		},
		{
			Display:  "reservationSummaries",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationSummaries", "2019-01-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Consumption/operations", "2019-01-01"),
		},
		{
			Display:  "aggregatedcost",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/Microsoft.Consumption/aggregatedcost", "2019-01-01"),
		},
		{
			Display:  "aggregatedcost",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Consumption/aggregatedcost", "2019-01-01"),
		},
		{
			Display:  "default",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default", "2019-01-01"),
		},
		{
			Display:  "forecasts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/forecasts", "2019-01-01"),
		},
		{
			Display:  "default",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default", "2019-01-01"),
		},
		{
			Display:  "reservationRecommendations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/reservationRecommendations", "2019-01-01"),
		},
		{
			Display:  "budgets",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{budgetName}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-01-01"),
				}},
		},
		{
			Display:  "charges",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/charges", "2019-01-01"),
		},
		{
			Display:  "marketplaces",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/marketplaces", "2019-01-01"),
		},
		{
			Display:  "tags",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/tags", "2019-01-01"),
		},
		{
			Display:  "usageDetails",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/usageDetails", "2019-01-01"),
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
			SubResources: []SwaggerResourceType{
				{
					Display:        "{containerGroupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					Children:       []SwaggerResourceType{},
					SubResources: []SwaggerResourceType{
						{
							Display:  "logs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", "2018-10-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ContainerRegistry/operations", "2019-05-01"),
		},
		{
			Display:  "registries",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
		},
		{
			Display:  "registries",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{registryName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "listUsages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listUsages", "2019-05-01"),
						},
						{
							Display:  "replications",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications", "2019-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{replicationName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
								}},
						},
						{
							Display:  "webhooks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks", "2019-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{webhookName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "orchestrators",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators", "2019-06-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ContainerService/operations", "2019-06-01"),
		},
		{
			Display:  "managedClusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters", "2019-06-01"),
		},
		{
			Display:  "managedClusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "agentPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools", "2019-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{agentPoolName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "availableAgentPoolVersions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/availableAgentPoolVersions", "2019-06-01"),
										},
										{
											Display:  "default",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default", "2019-06-01"),
										}},
								}},
						},
						{
							Display:  "default",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeProfiles/default", "2019-06-01"),
						}},
					SubResources: []SwaggerResourceType{},
				}},
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
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2015-04-08"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2015-04-08"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2015-04-08"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2015-04-08"),
					Children: []SwaggerResourceType{
						{
							Display:  "keyspaces",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces", "2015-04-08"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{keyspaceName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}", "2015-04-08"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}", "2015-04-08"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}", "2015-04-08"),
									Children: []SwaggerResourceType{
										{
											Display:     "throughput",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/settings/throughput", "2015-04-08"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/settings/throughput", "2015-04-08"),
										},
										{
											Display:  "tables",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables", "2015-04-08"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{tableName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}", "2015-04-08"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}", "2015-04-08"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}", "2015-04-08"),
													Children: []SwaggerResourceType{
														{
															Display:     "throughput",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}/settings/throughput", "2015-04-08"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}/settings/throughput", "2015-04-08"),
														}},
												}},
										}},
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases", "2015-04-08"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}", "2015-04-08"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}", "2015-04-08"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}", "2015-04-08"),
									Children: []SwaggerResourceType{
										{
											Display:  "graphs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs", "2015-04-08"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{graphName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}", "2015-04-08"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}", "2015-04-08"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}", "2015-04-08"),
													Children: []SwaggerResourceType{
														{
															Display:     "throughput",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}/settings/throughput", "2015-04-08"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}/settings/throughput", "2015-04-08"),
														}},
												}},
										},
										{
											Display:     "throughput",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/settings/throughput", "2015-04-08"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/settings/throughput", "2015-04-08"),
										}},
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases", "2015-04-08"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}", "2015-04-08"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}", "2015-04-08"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}", "2015-04-08"),
									Children: []SwaggerResourceType{
										{
											Display:  "collections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections", "2015-04-08"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{collectionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}", "2015-04-08"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}", "2015-04-08"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}", "2015-04-08"),
													Children: []SwaggerResourceType{
														{
															Display:     "throughput",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}/settings/throughput", "2015-04-08"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}/settings/throughput", "2015-04-08"),
														}},
												}},
										},
										{
											Display:     "throughput",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/settings/throughput", "2015-04-08"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/settings/throughput", "2015-04-08"),
										}},
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases", "2015-04-08"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}", "2015-04-08"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}", "2015-04-08"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}", "2015-04-08"),
									Children: []SwaggerResourceType{
										{
											Display:  "containers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers", "2015-04-08"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{containerName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}", "2015-04-08"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}", "2015-04-08"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}", "2015-04-08"),
													Children: []SwaggerResourceType{
														{
															Display:     "throughput",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}/settings/throughput", "2015-04-08"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}/settings/throughput", "2015-04-08"),
														}},
												}},
										},
										{
											Display:     "throughput",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/settings/throughput", "2015-04-08"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/settings/throughput", "2015-04-08"),
										}},
								}},
						},
						{
							Display:  "tables",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables", "2015-04-08"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{tableName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}", "2015-04-08"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}", "2015-04-08"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}", "2015-04-08"),
									Children: []SwaggerResourceType{
										{
											Display:     "throughput",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}/settings/throughput", "2015-04-08"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}/settings/throughput", "2015-04-08"),
										}},
								}},
						},
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
						}},
					SubResources: []SwaggerResourceType{
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
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/operations", "2019-01-01"),
		},
		{
			Display:  "dimensions",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/dimensions", "2019-01-01"),
		},
		{
			Display:  "exports",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{exportName}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-01-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.CustomerInsights/operations", "2017-04-26"),
		},
		{
			Display:  "hubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
		},
		{
			Display:  "hubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{hubName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					Children: []SwaggerResourceType{
						{
							Display:  "authorizationPolicies",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{authorizationPolicyName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
									Children:    []SwaggerResourceType{},
								}},
						},
						{
							Display:  "connectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{connectorName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									Children: []SwaggerResourceType{
										{
											Display:  "mappings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings", "2017-04-26"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{mappingName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
												}},
										}},
								}},
						},
						{
							Display:  "interactions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{interactionName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
									Children:    []SwaggerResourceType{},
								}},
						},
						{
							Display:  "kpi",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{kpiName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "links",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{linkName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
								}},
						},
						{
							Display:  "predictions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{predictionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "profiles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{profileName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "relationshipLinks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{relationshipLinkName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
								}},
						},
						{
							Display:  "relationships",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{relationshipName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
								}},
						},
						{
							Display:  "roleAssignments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{assignmentName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
								}},
						},
						{
							Display:  "roles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roles", "2017-04-26"),
						},
						{
							Display:  "views",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{viewName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
								}},
						},
						{
							Display:  "widgetTypes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes", "2017-04-26"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{widgetTypeName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes/{widgetTypeName}", "2017-04-26"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataBox/operations", "2018-01-01"),
		},
		{
			Display:  "jobs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs", "2018-01-01"),
		},
		{
			Display:  "jobs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs", "2018-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{jobName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2018-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2018-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2018-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2018-01-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Databricks/operations", "2018-04-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataCatalog/operations", "2016-03-30"),
		},
		{
			Display:  "catalogs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs", "2016-03-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{catalogName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataFactory/operations", "2018-06-01"),
		},
		{
			Display:  "factories",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
		},
		{
			Display:  "factories",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{factoryName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "datasets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{datasetName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
								}},
						},
						{
							Display:  "integrationRuntimes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{integrationRuntimeName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									Children:       []SwaggerResourceType{},
									SubResources: []SwaggerResourceType{
										{
											Display:        "{nodeName}",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											Children:       []SwaggerResourceType{},
										}},
								}},
						},
						{
							Display:  "linkedservices",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{linkedServiceName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
								}},
						},
						{
							Display:  "pipelines",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{pipelineName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "triggers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{triggerName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:      "rerunTriggers",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers", "2018-06-01"),
											SubResources: []SwaggerResourceType{},
										}},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{runId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}", "2018-06-01"),
							Children: []SwaggerResourceType{},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeAnalytics/operations", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
		},
		{
			Display:  "capability",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/locations/{location}/capability", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "computePolicies",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{computePolicyName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
								}},
						},
						{
							Display:  "dataLakeStoreAccounts",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{dataLakeStoreAccountName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
								}},
						},
						{
							Display:  "storageAccounts",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{storageAccountName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "containers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers", "2016-11-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{containerName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2016-11-01"),
													Children: []SwaggerResourceType{},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeStore/operations", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
		},
		{
			Display:  "capability",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability", "2016-11-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/usages", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "firewallRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
								}},
						},
						{
							Display:  "trustedIdProviders",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{trustedIdProviderName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
								}},
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules", "2016-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataMigration/operations", "2018-04-19"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/locations/{location}/usages", "2018-04-19"),
		},
		{
			Display:  "services",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/services", "2018-04-19"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus", "2018-04-19"),
		},
		{
			Display:  "services",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services", "2018-04-19"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					Children: []SwaggerResourceType{
						{
							Display:  "projects",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects", "2018-04-19"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{projectName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									Children: []SwaggerResourceType{
										{
											Display:  "tasks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks", "2018-04-19"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{taskName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/skus", "2018-04-19"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Devices/operations", "2018-04-01"),
		},
		{
			Display:  "provisioningServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
		},
		{
			Display:  "provisioningServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{provisioningServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					Children: []SwaggerResourceType{
						{
							Display:  "certificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates", "2018-01-22"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/skus", "2018-01-22"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{operationId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/operationresults/{operationId}", "2018-01-22"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DevSpaces/operations", "2019-04-01"),
		},
		{
			Display:  "controllers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
		},
		{
			Display:  "controllers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DevTestLab/operations", "2018-09-15"),
		},
		{
			Display:  "labs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
		},
		{
			Display:  "{name}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/locations/{locationName}/operations/{name}", "2018-09-15"),
		},
		{
			Display:  "schedules",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
		},
		{
			Display:  "labs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "artifactsources",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "armtemplates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{name}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "artifacts",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{name}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}", "2018-09-15"),
									Children: []SwaggerResourceType{},
								}},
						},
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
						}},
				},
				{
					Display:     "{name}",
					Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
					PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
				},
				{
					Display:  "customimages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "formulas",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "galleryimages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/galleryimages", "2018-09-15"),
				},
				{
					Display:  "notificationchannels",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							Children:       []SwaggerResourceType{},
						}},
				},
				{
					Display:  "policies",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "schedules",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							Children:       []SwaggerResourceType{},
						}},
				},
				{
					Display:  "servicerunners",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "users",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
						},
						{
							Display:  "disks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "environments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "secrets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "servicefabrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									Children:       []SwaggerResourceType{},
								},
								{
									Display:  "schedules",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules", "2018-09-15"),
									SubResources: []SwaggerResourceType{
										{
											Display:        "{name}",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											Children:       []SwaggerResourceType{},
										}},
								}},
						}},
				},
				{
					Display:  "virtualmachines",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							Children:       []SwaggerResourceType{},
						},
						{
							Display:  "schedules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules", "2018-09-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				},
				{
					Display:  "virtualnetworks",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks", "2018-09-15"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{name}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
						}},
				},
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "schedules",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "dnszones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones", "2018-05-01"),
		},
		{
			Display:  "dnsZones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones", "2018-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{zoneName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "all",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/all", "2018-05-01"),
						},
						{
							Display:  "recordsets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets", "2018-05-01"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{recordType}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}", "2018-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{relativeRecordSetName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.AAD/operations", "2017-06-01"),
		},
		{
			Display:  "domainServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices", "2017-06-01"),
		},
		{
			Display:  "domainServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices", "2017-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{domainServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "replicaSets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets", "2017-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{replicaSetName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DataBoxEdge/operations", "2019-03-01"),
		},
		{
			Display:  "dataBoxEdgeDevices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-03-01"),
		},
		{
			Display:  "dataBoxEdgeDevices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-03-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{deviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-03-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-03-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-03-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-03-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "alerts",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{name}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts/{name}", "2019-03-01"),
								}},
						},
						{
							Display:  "bandwidthSchedules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-03-01"),
								}},
						},
						{
							Display:  "default",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/networkSettings/default", "2019-03-01"),
						},
						{
							Display:  "orders",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders", "2019-03-01"),
							Children: []SwaggerResourceType{
								{
									Display:        "default",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-03-01"),
								}},
						},
						{
							Display:  "roles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-03-01"),
								}},
						},
						{
							Display:  "shares",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-03-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "storageAccountCredentials",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-03-01"),
								}},
						},
						{
							Display:  "triggers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-03-01"),
								}},
						},
						{
							Display:  "default",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/updateSummary/default", "2019-03-01"),
						},
						{
							Display:  "users",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-03-01"),
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{name}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/jobs/{name}", "2019-03-01"),
						},
						{
							Display:  "{name}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/operationsStatus/{name}", "2019-03-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/operations", "2019-06-01"),
		},
		{
			Display:  "topicTypes",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{topicTypeName}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}", "2019-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "eventTypes",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes", "2019-06-01"),
						}},
				}},
		},
		{
			Display:  "domains",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "topics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
		},
		{
			Display:  "domains",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{domainName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "topics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics", "2019-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{domainTopicName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
								},
								{
									Display:  "eventSubscriptions",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{topicName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "topics",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{topicName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventTypes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventTypes", "2019-06-01"),
		},
		{
			Display:        "{eventSubscriptionName}",
			Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			Children:       []SwaggerResourceType{},
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
			SubResources: []SwaggerResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "disasterRecoveryConfigs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
							Children: []SwaggerResourceType{},
							SubResources: []SwaggerResourceType{
								{
									Display:        "{alias}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "eventhubs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{eventHubName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "consumergroups",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{consumerGroupName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
												}},
										}},
								}},
						},
						{
							Display:  "messagingplan",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/messagingplan", "2017-04-01"),
						},
						{
							Display:     "default",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
						}},
				}},
		},
		{
			Display:  "frontDoors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors", "2019-04-01"),
		},
		{
			Display:  "frontDoors",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{frontDoorName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "backendPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{backendPoolName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", "2019-04-01"),
								}},
						},
						{
							Display:  "frontendEndpoints",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{frontendEndpointName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "healthProbeSettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/healthProbeSettings", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{healthProbeSettingsName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/healthProbeSettings/{healthProbeSettingsName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/healthProbeSettings/{healthProbeSettingsName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/healthProbeSettings/{healthProbeSettingsName}", "2019-04-01"),
								}},
						},
						{
							Display:  "loadBalancingSettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/loadBalancingSettings", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{loadBalancingSettingsName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/loadBalancingSettings/{loadBalancingSettingsName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/loadBalancingSettings/{loadBalancingSettingsName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/loadBalancingSettings/{loadBalancingSettingsName}", "2019-04-01"),
								}},
						},
						{
							Display:  "routingRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/routingRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{routingRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/routingRules/{routingRuleName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/routingRules/{routingRuleName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/routingRules/{routingRuleName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.GuestConfiguration/operations", "2018-11-20"),
		},
		{
			Display:  "guestConfigurationAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments", "2018-11-20"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{guestConfigurationAssignmentName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
					Children: []SwaggerResourceType{
						{
							Display:  "reports",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports", "2018-11-20"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{reportId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports/{reportId}", "2018-11-20"),
								}},
						}},
				}},
		},
		{
			Display:  "applications",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications", "2018-06-01-preview"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{applicationName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
				}},
		},
		{
			Display:  "clusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
		},
		{
			Display:  "clusters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{clusterName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					Children: []SwaggerResourceType{
						{
							Display:        "clustermonitoring",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
						},
						{
							Display:      "scriptActions",
							Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions", "2018-06-01-preview"),
							SubResources: []SwaggerResourceType{},
						},
						{
							Display:  "scriptExecutionHistory",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory", "2018-06-01-preview"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{scriptExecutionId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}", "2018-06-01-preview"),
									Children: []SwaggerResourceType{},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:        "{extensionName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
						}},
				}},
		},
		{
			Display:  "billingSpecs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/billingSpecs", "2018-06-01-preview"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/usages", "2018-06-01-preview"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.HDInsight/operations", "2018-06-01-preview"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.HybridData/operations", "2016-06-01"),
		},
		{
			Display:  "dataManagers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
		},
		{
			Display:  "dataManagers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{dataManagerName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "dataServices",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{dataServiceName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}", "2016-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "jobDefinitions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions", "2016-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{jobDefinitionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "jobs",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs", "2016-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{jobId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}", "2016-06-01"),
																	Children: []SwaggerResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "jobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs", "2016-06-01"),
										}},
								}},
						},
						{
							Display:  "dataStoreTypes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{dataStoreTypeName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes/{dataStoreTypeName}", "2016-06-01"),
								}},
						},
						{
							Display:  "dataStores",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{dataStoreName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
								}},
						},
						{
							Display:  "jobDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobDefinitions", "2016-06-01"),
						},
						{
							Display:  "jobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs", "2016-06-01"),
						},
						{
							Display:  "publicKeys",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{publicKeyName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys/{publicKeyName}", "2016-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.IoTCentral/operations", "2018-09-01"),
		},
		{
			Display:  "IoTApps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
		},
		{
			Display:  "IoTApps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
				}},
		},
		{
			Display:  "IotHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/IotHubs", "2018-04-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/usages", "2018-04-01"),
		},
		{
			Display:  "IotHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs", "2018-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "routingEndpointsHealth",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/routingEndpointsHealth", "2018-04-01"),
				},
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2018-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2018-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2018-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2018-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "IotHubStats",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/IotHubStats", "2018-04-01"),
						},
						{
							Display:  "certificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates", "2018-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2018-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2018-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2018-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "jobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs", "2018-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{jobId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs/{jobId}", "2018-04-01"),
								}},
						},
						{
							Display:  "quotaMetrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/quotaMetrics", "2018-04-01"),
						},
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/skus", "2018-04-01"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "ConsumerGroups",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups", "2018-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2018-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2018-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2018-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "deletedVaults",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedVaults", "2018-02-14"),
		},
		{
			Display:  "{vaultName}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}", "2018-02-14"),
			Children: []SwaggerResourceType{},
		},
		{
			Display:  "vaults",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/vaults", "2018-02-14"),
		},
		{
			Display:  "vaults",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults", "2018-02-14"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{vaultName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					Children: []SwaggerResourceType{
						{
							Display:  "secrets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets", "2018-02-14"),
							SubResources: []SwaggerResourceType{
								{
									Display:       "{secretName}",
									Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
									PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
									PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
								}},
						}},
					SubResources: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "resources",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resources", "2019-05-10"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.KeyVault/operations", "2018-02-14"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.LabServices/operations", "2018-10-15"),
		},
		{
			Display:  "labaccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
		},
		{
			Display:  "{operationName}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/locations/{locationName}/operations/{operationName}", "2018-10-15"),
		},
		{
			Display:  "labaccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{labAccountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					Children: []SwaggerResourceType{
						{
							Display:  "galleryimages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages", "2018-10-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{galleryImageName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
								}},
						},
						{
							Display:  "labs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs", "2018-10-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{labName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									Children: []SwaggerResourceType{
										{
											Display:  "environmentsettings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings", "2018-10-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{environmentSettingName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													Children: []SwaggerResourceType{
														{
															Display:  "environments",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments", "2018-10-15"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{environmentName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	Children:       []SwaggerResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "users",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users", "2018-10-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{userName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Logic/operations", "2016-06-01"),
		},
		{
			Display:  "integrationAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts", "2016-06-01"),
		},
		{
			Display:  "workflows",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows", "2016-06-01"),
		},
		{
			Display:  "integrationAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts", "2016-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{integrationAccountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2016-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2016-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2016-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2016-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "agreements",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{agreementName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2016-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "assemblies",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{assemblyArtifactName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2016-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "batchConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{batchConfigurationName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2016-06-01"),
								}},
						},
						{
							Display:  "certificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2016-06-01"),
								}},
						},
						{
							Display:  "maps",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{mapName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2016-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "partners",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{partnerName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2016-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "schemas",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{schemaName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2016-06-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "sessions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{sessionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2016-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2016-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2016-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "workflows",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows", "2016-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workflowName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2016-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2016-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2016-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2016-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "runs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{runName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}", "2016-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "actions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions", "2016-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{actionName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}", "2016-06-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "repetitions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions", "2016-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{repetitionName}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}", "2016-06-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "requestHistories",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories", "2016-06-01"),
																			SubResources: []SwaggerResourceType{
																				{
																					Display:  "{requestHistoryName}",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories/{requestHistoryName}", "2016-06-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "requestHistories",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories", "2016-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{requestHistoryName}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories/{requestHistoryName}", "2016-06-01"),
																}},
														},
														{
															Display:  "scopeRepetitions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions", "2016-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{repetitionName}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions/{repetitionName}", "2016-06-01"),
																}},
														}},
												}},
										}},
									SubResources: []SwaggerResourceType{
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/operations/{operationId}", "2016-06-01"),
										}},
								}},
						},
						{
							Display:  "triggers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{triggerName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}", "2016-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "histories",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories", "2016-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{historyName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}", "2016-06-01"),
													Children: []SwaggerResourceType{},
												}},
										},
										{
											Display:  "json",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/schemas/json", "2016-06-01"),
										}},
								}},
						},
						{
							Display:  "versions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions", "2016-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:      "{versionId}",
									Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}", "2016-06-01"),
									SubResources: []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearning/operations", "2017-01-01"),
		},
		{
			Display:  "webServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/webServices", "2017-01-01"),
		},
		{
			Display:  "webServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices", "2017-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{webServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "listKeys",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}/listKeys", "2017-01-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearningServices/operations", "2019-05-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usages", "2019-05-01"),
		},
		{
			Display:  "vmSizes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/vmSizes", "2019-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces", "2019-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces", "2019-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-05-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "computes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes", "2019-05-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{computeName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-05-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-05-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-05-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-05-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ManagedServices/operations", "2019-06-01"),
		},
		{
			Display:  "registrationAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{registrationAssignmentId}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
				}},
		},
		{
			Display:  "registrationDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{registrationDefinitionId}",
					Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Maps/operations", "2018-05-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Maps/accounts", "2018-05-01"),
		},
		{
			Display:  "accounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts", "2018-05-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DBforMariaDB/operations", "2018-06-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/locations/{locationName}/performanceTiers", "2018-06-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/servers", "2018-06-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers", "2018-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "configurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations/{configurationName}", "2018-06-01"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations/{configurationName}", "2018-06-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/logFiles", "2018-06-01"),
						},
						{
							Display:  "replicas",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/replicas", "2018-06-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules", "2018-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2018-06-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2018-06-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.MarketplaceOrdering/operations", "2015-06-01"),
		},
		{
			Display:  "agreements",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements", "2015-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{planId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}", "2015-06-01"),
					Children: []SwaggerResourceType{},
				}},
		},
		{
			Display:     "current",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
		},
		{
			Display:  "accountFilters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{filterName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Media/operations", "2018-07-01"),
		},
		{
			Display:  "mediaservices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{accountName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
				}},
		},
		{
			Display:  "mediaservices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "liveEvents",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents", "2018-07-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{liveEventName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "liveOutputs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs", "2018-07-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{liveOutputName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
												}},
										}},
								}},
						},
						{
							Display:  "streamingEndpoints",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints", "2018-07-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{streamingEndpointName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "assets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{assetName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "assetFilters",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters", "2018-07-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{filterName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
								}},
						}},
				}},
		},
		{
			Display:  "contentKeyPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{contentKeyPolicyName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "transforms",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{transformName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "jobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs", "2018-07-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "streamingLocators",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{streamingLocatorName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "streamingPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies", "2018-07-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{streamingPolicyName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Migrate/operations", "2018-02-02"),
		},
		{
			Display:  "assessmentOptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/locations/{locationName}/assessmentOptions", "2018-02-02"),
		},
		{
			Display:  "projects",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/projects", "2018-02-02"),
		},
		{
			Display:  "assessments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/assessments", "2018-02-02"),
		},
		{
			Display:  "groups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups", "2018-02-02"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{groupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					Children: []SwaggerResourceType{
						{
							Display:  "assessments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments", "2018-02-02"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{assessmentName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									Children: []SwaggerResourceType{
										{
											Display:  "assessedMachines",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines", "2018-02-02"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{assessedMachineName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}", "2018-02-02"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "machines",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines", "2018-02-02"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{machineName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines/{machineName}", "2018-02-02"),
				}},
		},
		{
			Display:  "projects",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects", "2018-02-02"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{projectName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "actionGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/microsoft.insights/actionGroups", "2019-06-01"),
		},
		{
			Display:  "actionGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{actionGroupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ManagedIdentity/operations", "2018-11-30"),
		},
		{
			Display:  "userAssignedIdentities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
		},
		{
			Display:  "userAssignedIdentities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DBforMySQL/operations", "2017-12-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/performanceTiers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/servers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers", "2017-12-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "configurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/logFiles", "2017-12-01"),
						},
						{
							Display:  "replicas",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/replicas", "2017-12-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.NetApp/operations", "2019-06-01"),
		},
		{
			Display:  "netAppAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts", "2019-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "capacityPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools", "2019-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{poolName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-06-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "volumes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes", "2019-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{volumeName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-06-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-06-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-06-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-06-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "mountTargets",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/mountTargets", "2019-06-01"),
														},
														{
															Display:  "snapshots",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots", "2019-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{snapshotName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-06-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-06-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-06-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-06-01"),
																}},
														}},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "applicationGatewayAvailableRequestHeaders",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableRequestHeaders", "2019-04-01"),
		},
		{
			Display:  "applicationGatewayAvailableResponseHeaders",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableResponseHeaders", "2019-04-01"),
		},
		{
			Display:  "applicationGatewayAvailableServerVariables",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableServerVariables", "2019-04-01"),
		},
		{
			Display:  "default",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default", "2019-04-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "predefinedPolicies",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies", "2019-04-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{predefinedPolicyName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/{predefinedPolicyName}", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "applicationGatewayAvailableWafRuleSets",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", "2019-04-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", "2019-04-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{applicationGatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-04-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups", "2019-04-01"),
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{applicationSecurityGroupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-04-01"),
				}},
		},
		{
			Display:  "availableDelegations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-04-01"),
		},
		{
			Display:  "availableDelegations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-04-01"),
		},
		{
			Display:  "availablePrivateEndpointTypes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-04-01"),
		},
		{
			Display:  "availablePrivateEndpointTypes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-04-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls", "2019-04-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{azureFirewallName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-04-01"),
				}},
		},
		{
			Display:  "azureFirewallFqdnTags",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags", "2019-04-01"),
		},
		{
			Display:  "bastionHosts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts", "2019-04-01"),
		},
		{
			Display:  "bastionHosts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{bastionHostName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-04-01"),
				}},
		},
		{
			Display:  "CheckDnsNameAvailability",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", "2019-04-01"),
		},
		{
			Display:        "{ddosCustomPolicyName}",
			Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-04-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-04-01"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-04-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-04-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans", "2019-04-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{ddosProtectionPlanName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-04-01"),
				}},
		},
		{
			Display:  "virtualNetworkAvailableEndpointServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices", "2019-04-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", "2019-04-01"),
		},
		{
			Display:  "expressRouteServiceProviders",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", "2019-04-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{circuitName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "authorizations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{authorizationName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-04-01"),
								}},
						},
						{
							Display:  "peerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{peeringName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "connections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections", "2019-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{connectionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-04-01"),
												}},
										},
										{
											Display:  "peerConnections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections", "2019-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{connectionName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}", "2019-04-01"),
												}},
										},
										{
											Display:  "stats",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats", "2019-04-01"),
										}},
									SubResources: []SwaggerResourceType{},
								}},
						},
						{
							Display:  "stats",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-04-01"),
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:       "{crossConnectionName}",
					Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-04-01"),
					PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-04-01"),
					PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "peerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{peeringName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-04-01"),
									SubResources:   []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways", "2019-04-01"),
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{expressRouteGatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "expressRouteConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{connectionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts", "2019-04-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "ExpressRoutePortsLocations",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations", "2019-04-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{locationName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{expressRoutePortName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "links",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{linkName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "loadBalancers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers", "2019-04-01"),
		},
		{
			Display:  "loadBalancers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{loadBalancerName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "backendAddressPools",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{backendAddressPoolName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", "2019-04-01"),
								}},
						},
						{
							Display:  "frontendIPConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{frontendIPConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}", "2019-04-01"),
								}},
						},
						{
							Display:  "inboundNatRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{inboundNatRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-04-01"),
								}},
						},
						{
							Display:  "loadBalancingRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{loadBalancingRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}", "2019-04-01"),
								}},
						},
						{
							Display:  "networkInterfaces",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", "2019-04-01"),
						},
						{
							Display:  "outboundRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{outboundRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}", "2019-04-01"),
								}},
						},
						{
							Display:  "probes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{probeName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "natGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways", "2019-04-01"),
		},
		{
			Display:  "natGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{natGatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-04-01"),
				}},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", "2019-04-01"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{networkInterfaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2019-04-01"),
								}},
						},
						{
							Display:  "loadBalancers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers", "2019-04-01"),
						},
						{
							Display:  "tapConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{tapConfigurationName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "networkProfiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles", "2019-04-01"),
		},
		{
			Display:  "networkProfiles",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{networkProfileName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-04-01"),
				}},
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups", "2019-04-01"),
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{networkSecurityGroupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "defaultSecurityRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{defaultSecurityRuleName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}", "2019-04-01"),
								}},
						},
						{
							Display:  "securityRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{securityRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "networkWatchers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkWatchers", "2019-04-01"),
		},
		{
			Display:  "networkWatchers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{networkWatcherName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "connectionMonitors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{connectionMonitorName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "packetCaptures",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{packetCaptureName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Network/operations", "2019-04-01"),
		},
		{
			Display:  "privateEndpoints",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints", "2019-04-01"),
		},
		{
			Display:  "privateEndpoints",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{privateEndpointName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-04-01"),
				}},
		},
		{
			Display:  "privateLinkServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateLinkServices", "2019-04-01"),
		},
		{
			Display:  "privateLinkServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-04-01"),
					SubResources:   []SwaggerResourceType{},
				}},
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", "2019-04-01"),
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{publicIpAddressName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-04-01"),
				}},
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes", "2019-04-01"),
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{publicIpPrefixName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-04-01"),
				}},
		},
		{
			Display:  "routeFilters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters", "2019-04-01"),
		},
		{
			Display:  "routeFilters",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{routeFilterName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "routeFilterRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{ruleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "routeTables",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables", "2019-04-01"),
		},
		{
			Display:  "routeTables",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{routeTableName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "routes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{routeName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "bgpServiceCommunities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities", "2019-04-01"),
		},
		{
			Display:  "ServiceEndpointPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies", "2019-04-01"),
		},
		{
			Display:  "serviceEndpointPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serviceEndpointPolicyName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "serviceEndpointPolicyDefinitions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{serviceEndpointPolicyDefinitionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "serviceTags",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags", "2019-04-01"),
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages", "2019-04-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", "2019-04-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{virtualNetworkName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "CheckIPAddressAvailability",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", "2019-04-01"),
						},
						{
							Display:  "subnets",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{subnetName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "ResourceNavigationLinks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks", "2019-04-01"),
										},
										{
											Display:  "ServiceAssociationLinks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks", "2019-04-01"),
										}},
								}},
						},
						{
							Display:  "usages",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", "2019-04-01"),
						},
						{
							Display:  "virtualNetworkPeerings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkPeeringName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "connections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{virtualNetworkGatewayConnectionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:     "sharedkey",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-04-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-04-01"),
							Children:    []SwaggerResourceType{},
						}},
				}},
		},
		{
			Display:  "localNetworkGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{localNetworkGatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-04-01"),
				}},
		},
		{
			Display:  "virtualNetworkGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{virtualNetworkGatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "connections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/connections", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps", "2019-04-01"),
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{tapName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-04-01"),
				}},
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways", "2019-04-01"),
		},
		{
			Display:  "virtualHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs", "2019-04-01"),
		},
		{
			Display:  "virtualWans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans", "2019-04-01"),
		},
		{
			Display:  "vpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways", "2019-04-01"),
		},
		{
			Display:  "vpnSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites", "2019-04-01"),
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{gatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-04-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "virtualHubs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{virtualHubName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "hubVirtualNetworkConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{connectionName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "virtualWans",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{VirtualWANName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-04-01"),
				},
				{
					Display:  "supportedSecurityProviders",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", "2019-04-01"),
				},
				{
					Display:  "p2sVpnServerConfigurations",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations", "2019-04-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{p2SVpnServerConfigurationName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", "2019-04-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", "2019-04-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "vpnGateways",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{gatewayName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "vpnConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{connectionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "vpnSites",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{vpnSiteName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-04-01"),
				}},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", "2017-03-30"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", "2017-03-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{networkInterfaceName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", "2017-03-30"),
					Children: []SwaggerResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2017-03-30"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2017-03-30"),
								}},
						}},
				}},
		},
		{
			Display:  "publicipaddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses", "2017-03-30"),
		},
		{
			Display:  "publicipaddresses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", "2017-03-30"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{publicIpAddressName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", "2017-03-30"),
				}},
		},
		{
			Display:  "ApplicationGatewayWebApplicationFirewallPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-04-01"),
		},
		{
			Display:  "ApplicationGatewayWebApplicationFirewallPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policyName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-04-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.NotificationHubs/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "notificationHubs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{notificationHubName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						}},
				}},
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
			SubResources: []SwaggerResourceType{
				{
					Display:        "{savedSearchId}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					Children: []SwaggerResourceType{
						{
							Display:  "results",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}/results", "2015-03-20"),
						}},
				}},
		},
		{
			Display:  "storageInsightConfigs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs", "2015-03-20"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{storageInsightName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
				}},
		},
		{
			Display:  "$metadata",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.PolicyInsights/policyEvents/$metadata", "2018-04-04"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.PolicyInsights/operations", "2018-04-04"),
		},
		{
			Display:  "$metadata",
			Endpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.PolicyInsights/policyStates/$metadata", "2018-04-04"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DBforPostgreSQL/operations", "2017-12-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/performanceTiers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers", "2017-12-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "Replicas",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/Replicas", "2017-12-01"),
						},
						{
							Display:  "configurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
									PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/logFiles", "2017-12-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules", "2017-12-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.PowerBIDedicated/operations", "2017-10-01"),
		},
		{
			Display:  "capacities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
		},
		{
			Display:  "skus",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/skus", "2017-10-01"),
		},
		{
			Display:  "capacities",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{dedicatedCapacityName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "skus",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/skus", "2017-10-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.PowerBI/operations", "2016-01-29"),
		},
		{
			Display:  "workspaceCollections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
		},
		{
			Display:  "workspaceCollections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{workspaceCollectionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					Children: []SwaggerResourceType{
						{
							Display:  "workspaces",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/workspaces", "2016-01-29"),
						}},
				}},
		},
		{
			Display:  "privateDnsZones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
		},
		{
			Display:  "privateDnsZones",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{privateZoneName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "ALL",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL", "2018-09-01"),
						},
						{
							Display:  "virtualNetworkLinks",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks", "2018-09-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualNetworkLinkName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "{recordType}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}", "2018-09-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{relativeRecordSetName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
								}},
						}},
				}},
		},
		{
			Display:  "replicationUsages",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages", "2016-06-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.RecoveryServices/operations", "2016-06-01"),
		},
		{
			Display:  "vaults",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
		},
		{
			Display:  "vaults",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{vaultName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:       "vaultExtendedInfo",
							Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
							PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
							PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
						}},
				}},
		},
		{
			Display:  "usages",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/usages", "2016-06-01"),
		},
		{
			Display:        "{intentObjectName}",
			Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
		},
		{
			Display:  "backupJobs",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs", "2017-07-01"),
			Children: []SwaggerResourceType{},
			SubResources: []SwaggerResourceType{
				{
					Display:  "{operationId}",
					Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/operationResults/{operationId}", "2017-07-01"),
				},
				{
					Display:  "{jobName}",
					Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}", "2017-07-01"),
				}},
		},
		{
			Display:  "backupPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies", "2017-07-01"),
		},
		{
			Display:  "backupProtectedItems",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectedItems", "2017-07-01"),
		},
		{
			Display:  "backupProtectionIntents",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionIntents", "2017-07-01"),
		},
		{
			Display:  "backupUsageSummaries",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupUsageSummaries", "2017-07-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/operations", "2018-07-10"),
		},
		{
			Display:  "replicationAlertSettings",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:     "{alertSettingName}",
					Endpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
					PutEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationEvents",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{eventName}",
					Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationFabrics",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{fabricName}",
					Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					Children: []SwaggerResourceType{
						{
							Display:  "replicationLogicalNetworks",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{logicalNetworkName}",
									Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks/{logicalNetworkName}", "2018-07-10"),
								}},
						},
						{
							Display:  "replicationNetworks",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{networkName}",
									Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}", "2018-07-10"),
									Children: []SwaggerResourceType{
										{
											Display:  "replicationNetworkMappings",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{networkMappingName}",
													Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
												}},
										}},
								}},
						},
						{
							Display:  "replicationProtectionContainers",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:     "{protectionContainerName}",
									Endpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
									PutEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
									Children: []SwaggerResourceType{
										{
											Display:  "replicationMigrationItems",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{migrationItemName}",
													Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													Children: []SwaggerResourceType{
														{
															Display:  "migrationRecoveryPoints",
															Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints", "2018-07-10"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{migrationRecoveryPointName}",
																	Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints/{migrationRecoveryPointName}", "2018-07-10"),
																}},
														}},
												}},
										},
										{
											Display:  "replicationProtectableItems",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{protectableItemName}",
													Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}", "2018-07-10"),
												}},
										},
										{
											Display:  "replicationProtectedItems",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{replicatedProtectedItemName}",
													Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													Children: []SwaggerResourceType{
														{
															Display:  "recoveryPoints",
															Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints", "2018-07-10"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{recoveryPointName}",
																	Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}", "2018-07-10"),
																}},
														},
														{
															Display:  "targetComputeSizes",
															Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/targetComputeSizes", "2018-07-10"),
														}},
												}},
										},
										{
											Display:  "replicationProtectionContainerMappings",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{mappingName}",
													Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "replicationRecoveryServicesProviders",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{providerName}",
									Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "replicationStorageClassifications",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{storageClassificationName}",
									Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}", "2018-07-10"),
									Children: []SwaggerResourceType{
										{
											Display:  "replicationStorageClassificationMappings",
											Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings", "2018-07-10"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{storageClassificationMappingName}",
													Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
												}},
										}},
								}},
						},
						{
							Display:  "replicationvCenters",
							Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters", "2018-07-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{vCenterName}",
									Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
								}},
						}},
				}},
		},
		{
			Display:  "replicationJobs",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs", "2018-07-10"),
			Children: []SwaggerResourceType{},
			SubResources: []SwaggerResourceType{
				{
					Display:  "{jobName}",
					Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}", "2018-07-10"),
					Children: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "replicationMigrationItems",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationMigrationItems", "2018-07-10"),
		},
		{
			Display:  "replicationNetworkMappings",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings", "2018-07-10"),
		},
		{
			Display:  "replicationNetworks",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworks", "2018-07-10"),
		},
		{
			Display:  "replicationPolicies",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policyName}",
					Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationProtectedItems",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectedItems", "2018-07-10"),
		},
		{
			Display:  "replicationProtectionContainerMappings",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings", "2018-07-10"),
		},
		{
			Display:  "replicationProtectionContainers",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers", "2018-07-10"),
		},
		{
			Display:  "replicationRecoveryPlans",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{recoveryPlanName}",
					Endpoint:       mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "replicationRecoveryServicesProviders",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders", "2018-07-10"),
		},
		{
			Display:  "replicationStorageClassificationMappings",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings", "2018-07-10"),
		},
		{
			Display:  "replicationStorageClassifications",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications", "2018-07-10"),
		},
		{
			Display:  "replicationSupportedOperatingSystems",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationSupportedOperatingSystems", "2018-07-10"),
		},
		{
			Display:  "replicationVaultHealth",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth", "2018-07-10"),
			Children: []SwaggerResourceType{},
		},
		{
			Display:  "replicationVaultSettings",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings", "2018-07-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:     "{vaultSettingName}",
					Endpoint:    mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
					PutEndpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationvCenters",
			Endpoint: mustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters", "2018-07-10"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Cache/operations", "2018-03-01"),
		},
		{
			Display:  "Redis",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/Redis", "2018-03-01"),
		},
		{
			Display:  "Redis",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis", "2018-03-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "firewallRules",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules", "2018-03-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{ruleName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
						}},
				},
				{
					Display:  "patchSchedules",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/patchSchedules", "2018-03-01"),
				},
				{
					Display:        "{name}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "linkedServers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", "2018-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{linkedServerName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
								}},
						},
						{
							Display:  "listUpgradeNotifications",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/listUpgradeNotifications", "2018-03-01"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:        "{default}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Relay/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "authorizationRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "hybridConnections",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{hybridConnectionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "wcfRelays",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{relayName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/operations", "2017-11-01"),
		},
		{
			Display:  "reservationOrders",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders", "2017-11-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{reservationOrderId}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}", "2017-11-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "reservations",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations", "2017-11-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:       "{reservationId}",
									Endpoint:      mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2017-11-01"),
									PatchEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2017-11-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "revisions",
											Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions", "2017-11-01"),
										}},
								}},
						}},
				}},
		},
		{
			Display:  "appliedReservations",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/appliedReservations", "2017-11-01"),
		},
		{
			Display:  "catalogs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/catalogs", "2017-11-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ResourceGraph/operations", "2019-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ResourceHealth/operations", "2017-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2017-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2017-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2017-07-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "current",
					Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current", "2017-07-01"),
				}},
		},
		{
			Display:  "childAvailabilityStatuses",
			Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/childAvailabilityStatuses", "2017-07-01"),
			Children: []SwaggerResourceType{
				{
					Display:  "current",
					Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/childAvailabilityStatuses/current", "2017-07-01"),
				}},
		},
		{
			Display:  "childResources",
			Endpoint: mustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/childResources", "2017-07-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments", "2019-01-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments", "2019-01-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyAssignments", "2019-01-01"),
		},
		{
			Display:        "{policyAssignmentId}",
			Endpoint:       mustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-01-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-01-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-01-01"),
		},
		{
			Display:        "{policyAssignmentName}",
			Endpoint:       mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-01-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-01-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-01-01"),
		},
		{
			Display:  "policyDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{policyDefinitionName}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "policyDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policyDefinitionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "policyDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policyDefinitionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policySetDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{policySetDefinitionName}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policySetDefinitionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{policySetDefinitionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-01-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Features/operations", "2015-12-01"),
		},
		{
			Display:  "features",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/features", "2015-12-01"),
		},
		{
			Display:  "features",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features", "2015-12-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{featureName}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}", "2015-12-01"),
					Children: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "{}",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/", "2019-05-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{deploymentName}",
					Endpoint:       mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
					Children: []SwaggerResourceType{
						{
							Display:  "operations",
							Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-05-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{operationId}",
									Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-05-10"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Resources/operations", "2019-05-10"),
		},
		{
			Display:  "providers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers", "2019-05-10"),
			Children: []SwaggerResourceType{
				{
					Display:  "{}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/", "2019-05-10"),
					SubResources: []SwaggerResourceType{
						{
							Display:        "{deploymentName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
							Children: []SwaggerResourceType{
								{
									Display:  "operations",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-05-10"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{operationId}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-05-10"),
										}},
								}},
						}},
				},
				{
					Display:  "applications",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applications", "2018-06-01"),
				},
				{
					Display:  "jobCollections",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
				},
				{
					Display:  "searchServices",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices", "2015-08-19"),
				},
				{
					Display:  "alerts",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/alerts", "2019-01-01"),
				},
				{
					Display:  "settings",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings", "2019-01-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:     "{settingName}",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings/{settingName}", "2019-01-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings/{settingName}", "2019-01-01"),
						}},
				},
				{
					Display:  "operations",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/operations", "2018-05-01"),
				},
				{
					Display:  "namespaces",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
				},
				{
					Display:  "premiumMessagingRegions",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/premiumMessagingRegions", "2017-04-01"),
				},
				{
					Display:  "clusters",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/clusters", "2018-02-01"),
				},
				{
					Display:  "SignalR",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
				},
				{
					Display:  "managers",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
				},
				{
					Display:  "skus",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/skus", "2019-04-01"),
				},
				{
					Display:  "storageAccounts",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts", "2019-04-01"),
				},
				{
					Display:  "jobs",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
				},
				{
					Display:  "storageSyncServices",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/storageSyncServices", "2019-03-01"),
				},
				{
					Display:  "streamingjobs",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
				},
				{
					Display:  "environments",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
				},
				{
					Display:        "default",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
				},
				{
					Display:  "trafficmanagerprofiles",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
				},
				{
					Display:  "dedicatedCloudNodes",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
				},
				{
					Display:  "dedicatedCloudServices",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
				},
				{
					Display:  "virtualMachines",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
				},
				{
					Display:  "certificateOrders",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders", "2018-02-01"),
				},
				{
					Display:  "domains",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/domains", "2018-02-01"),
				},
				{
					Display:  "topLevelDomains",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains", "2018-02-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{name}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}", "2018-02-01"),
							Children: []SwaggerResourceType{},
						}},
				},
				{
					Display:  "hostingEnvironments",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
				},
				{
					Display:  "serverfarms",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
				},
				{
					Display:  "certificates",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates", "2018-02-01"),
				},
				{
					Display:  "deletedSites",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", "2018-02-01"),
				},
				{
					Display:  "availableStacks",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", "2018-02-01"),
				},
				{
					Display:      "recommendations",
					Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", "2018-02-01"),
					Children:     []SwaggerResourceType{},
					SubResources: []SwaggerResourceType{},
				},
				{
					Display:  "resourceHealthMetadata",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
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
				}},
			SubResources: []SwaggerResourceType{
				{
					Display:  "{resourceProviderNamespace}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}", "2019-05-10"),
					Children: []SwaggerResourceType{},
				},
				{
					Display:  "alerts",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/alerts", "2019-01-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:      "{alertName}",
							Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/alerts/{alertName}", "2019-01-01"),
							SubResources: []SwaggerResourceType{},
						}},
				},
				{
					Display:  "{default}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}", "2018-05-01"),
					Children: []SwaggerResourceType{},
				},
				{
					Display:  "regions",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/sku/{sku}/regions", "2017-04-01"),
				},
				{
					Display:  "clusterVersions",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions", "2018-02-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{clusterVersion}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}", "2018-02-01"),
						}},
				},
				{
					Display:  "clusterVersions",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions", "2018-02-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{clusterVersion}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}", "2018-02-01"),
						}},
				},
				{
					Display:  "usages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/locations/{location}/usages", "2018-10-01"),
				},
				{
					Display:  "capabilities",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/capabilities", "2015-05-01"),
				},
				{
					Display:  "usages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages", "2015-05-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{usageName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages/{usageName}", "2015-05-01"),
						}},
				},
				{
					Display:  "usages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages", "2019-04-01"),
				},
				{
					Display:  "quotas",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas", "2016-03-01"),
				},
				{
					Display:  "availabilities",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/availabilities", "2019-04-01"),
				},
				{
					Display:  "{operationId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/operationResults/{operationId}", "2019-04-01"),
				},
				{
					Display:  "privateClouds",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds", "2019-04-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{pcName}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}", "2019-04-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "resourcePools",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools", "2019-04-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{resourcePoolName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}", "2019-04-01"),
										}},
								},
								{
									Display:  "virtualMachineTemplates",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates", "2019-04-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{virtualMachineTemplateName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}", "2019-04-01"),
										}},
								},
								{
									Display:  "virtualNetworks",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks", "2019-04-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{virtualNetworkName}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
										}},
								}},
						}},
				},
				{
					Display:  "usages",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/usages", "2019-04-01"),
				},
				{
					Display:  "deletedSites",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites", "2018-02-01"),
					SubResources: []SwaggerResourceType{
						{
							Display:  "{deletedSiteId}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites/{deletedSiteId}", "2018-02-01"),
						}},
				}},
		},
		{
			Display:  "resources",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/resources", "2019-05-10"),
		},
		{
			Display:  "resourcegroups",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups", "2019-05-10"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{resourceGroupName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-05-10"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-05-10"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-05-10"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-05-10"),
					Children: []SwaggerResourceType{
						{
							Display:  "{}",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/", "2019-05-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{deploymentName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-05-10"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "clusters",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters", "2018-02-01"),
						},
						{
							Display:  "streamingjobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									Children:       []SwaggerResourceType{},
									SubResources: []SwaggerResourceType{
										{
											Display:       "{transformationName}",
											Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
											PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
											PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
										}},
								}},
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "operations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations", "2019-05-10"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{operationId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}", "2019-05-10"),
								}},
						},
						{
							Display:        "{resourceName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-05-10"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-05-10"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-05-10"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-05-10"),
						},
						{
							Display:  "functions",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", "2016-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{functionName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "inputs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs", "2016-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{inputName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "outputs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs", "2016-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{outputName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									Children:       []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:      "tagNames",
			Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/tagNames", "2019-05-10"),
			SubResources: []SwaggerResourceType{},
		},
		{
			Display:        "{resourceId}",
			Endpoint:       mustGetEndpointInfoFromURL("/{resourceId}", "2019-05-10"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{resourceId}", "2019-05-10"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/{resourceId}", "2019-05-10"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{resourceId}", "2019-05-10"),
		},
		{
			Display:  "applicationDefinitions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions", "2018-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{applicationDefinitionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
				}},
		},
		{
			Display:  "applications",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications", "2018-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{applicationName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
				}},
		},
		{
			Display:        "{applicationId}",
			Endpoint:       mustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
		},
		{
			Display:  "jobCollections",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{jobCollectionName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "jobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs", "2016-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "history",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history", "2016-03-01"),
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Search/operations", "2015-08-19"),
		},
		{
			Display:  "searchServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices", "2015-08-19"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{searchServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					Children: []SwaggerResourceType{
						{
							Display:  "listQueryKeys",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys", "2015-08-19"),
						}},
					SubResources: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "alerts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/alerts", "2019-01-01"),
		},
		{
			Display:  "alerts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/alerts", "2019-01-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:      "{alertName}",
					Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/alerts/{alertName}", "2019-01-01"),
					SubResources: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ServiceBus/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "disasterRecoveryConfigs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
							Children: []SwaggerResourceType{},
							SubResources: []SwaggerResourceType{
								{
									Display:        "{alias}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "eventhubs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
						},
						{
							Display:  "migrationConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{configName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:     "default",
							Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
							PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
						},
						{
							Display:  "queues",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{queueName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "topics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics", "2017-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{topicName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "subscriptions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions", "2017-04-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{subscriptionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "rules",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules", "2017-04-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{ruleName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																}},
														}},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ServiceFabric/operations", "2018-02-01"),
		},
		{
			Display:        "{clusterName}",
			Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2018-02-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2018-02-01"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2018-02-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2018-02-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.SignalRService/operations", "2018-10-01"),
		},
		{
			Display:      "SignalR",
			Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
			SubResources: []SwaggerResourceType{},
		},
		{
			Display:        "{resourceName}",
			Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			Children:       []SwaggerResourceType{},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.StorSimple/operations", "2017-06-01"),
		},
		{
			Display:  "managers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{managerName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "accessControlRecords",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords", "2017-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{accessControlRecordName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
								}},
						},
						{
							Display:  "alerts",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts", "2017-06-01"),
						},
						{
							Display:  "backups",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/backups", "2016-10-01"),
						},
						{
							Display:  "devices",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices", "2017-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{deviceName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									Children: []SwaggerResourceType{
										{
											Display:     "default",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
										},
										{
											Display:  "backupScheduleGroups",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups", "2016-10-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{scheduleGroupName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
												}},
										},
										{
											Display:      "backups",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", "2017-06-01"),
											SubResources: []SwaggerResourceType{},
										},
										{
											Display:  "chapSettings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings", "2016-10-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{chapUserName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
												}},
										},
										{
											Display:  "disks",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/disks", "2016-10-01"),
										},
										{
											Display:  "fileservers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers", "2016-10-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{fileServerName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "metrics",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metrics", "2016-10-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "metricsDefinitions",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metricsDefinitions", "2016-10-01"),
																}},
														},
														{
															Display:  "shares",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares", "2016-10-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{shareName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metrics", "2016-10-01"),
																			Children: []SwaggerResourceType{
																				{
																					Display:  "metricsDefinitions",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metricsDefinitions", "2016-10-01"),
																				}},
																		}},
																}},
														}},
												}},
										},
										{
											Display:  "iscsiservers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers", "2016-10-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{iscsiServerName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "disks",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks", "2016-10-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{diskName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metrics", "2016-10-01"),
																			Children: []SwaggerResourceType{
																				{
																					Display:  "metricsDefinitions",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metricsDefinitions", "2016-10-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "metrics",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metrics", "2016-10-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "metricsDefinitions",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metricsDefinitions", "2016-10-01"),
																}},
														}},
												}},
										},
										{
											Display:  "jobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs", "2017-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{jobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs/{jobName}", "2017-06-01"),
													Children: []SwaggerResourceType{},
												}},
										},
										{
											Display:  "metrics",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metrics", "2017-06-01"),
											Children: []SwaggerResourceType{
												{
													Display:  "metricsDefinitions",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metricsDefinitions", "2017-06-01"),
												}},
										},
										{
											Display:       "default",
											Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
											PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
										},
										{
											Display:  "shares",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/shares", "2016-10-01"),
										},
										{
											Display:     "default",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
										},
										{
											Display:  "default",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/updateSummary/default", "2017-06-01"),
										},
										{
											Display:  "backupPolicies",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies", "2017-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{backupPolicyName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "schedules",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules", "2017-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{backupScheduleName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																}},
														}},
												}},
										},
										{
											Display:      "hardwareComponentGroups",
											Endpoint:     mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups", "2017-06-01"),
											SubResources: []SwaggerResourceType{},
										},
										{
											Display:       "default",
											Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
											PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
											Children:      []SwaggerResourceType{},
										},
										{
											Display:  "volumeContainers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers", "2017-06-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{volumeContainerName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "metrics",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metrics", "2017-06-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "metricsDefinitions",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metricsDefinitions", "2017-06-01"),
																}},
														},
														{
															Display:  "volumes",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes", "2017-06-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{volumeName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metrics", "2017-06-01"),
																			Children: []SwaggerResourceType{
																				{
																					Display:  "metricsDefinitions",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metricsDefinitions", "2017-06-01"),
																				}},
																		}},
																}},
														}},
												}},
										},
										{
											Display:  "volumes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumes", "2017-06-01"),
										}},
								}},
						},
						{
							Display:  "default",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/encryptionSettings/default", "2017-06-01"),
						},
						{
							Display:        "vaultExtendedInfo",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
						},
						{
							Display:  "fileservers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/fileservers", "2016-10-01"),
						},
						{
							Display:  "iscsiservers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/iscsiservers", "2016-10-01"),
						},
						{
							Display:  "jobs",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/jobs", "2017-06-01"),
						},
						{
							Display:  "metrics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metrics", "2017-06-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "metricsDefinitions",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metricsDefinitions", "2017-06-01"),
								}},
						},
						{
							Display:  "storageAccountCredentials",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials", "2017-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{credentialName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
								},
								{
									Display:        "{storageAccountCredentialName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
								}},
						},
						{
							Display:  "storageDomains",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains", "2016-10-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{storageDomainName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
								}},
						},
						{
							Display:  "bandwidthSettings",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings", "2017-06-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{bandwidthSettingName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
								}},
						},
						{
							Display:  "cloudApplianceConfigurations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/cloudApplianceConfigurations", "2017-06-01"),
						},
						{
							Display:  "features",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/features", "2017-06-01"),
						}},
					SubResources: []SwaggerResourceType{},
				}},
		},
		{
			Display:  "containers",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{containerName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-04-01"),
					Children:       []SwaggerResourceType{},
					SubResources: []SwaggerResourceType{
						{
							Display:        "{immutabilityPolicyName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-04-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-04-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-04-01"),
						}},
				}},
		},
		{
			Display:     "{BlobServicesName}",
			Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-04-01"),
			PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Storage/operations", "2019-04-01"),
		},
		{
			Display:  "storageAccounts",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts", "2019-04-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-04-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-04-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-04-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-04-01"),
					Children:       []SwaggerResourceType{},
					SubResources: []SwaggerResourceType{
						{
							Display:        "{managementPolicyName}",
							Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-04-01"),
							DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-04-01"),
							PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-04-01"),
						}},
				}},
		},
		{
			Display:  "locations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations", "2016-11-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{locationName}",
					Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations/{locationName}", "2016-11-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/operations", "2016-11-01"),
		},
		{
			Display:  "jobs",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{jobName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					Children:       []SwaggerResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.StorageSync/operations", "2019-03-01"),
		},
		{
			Display:  "{operationId}",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/locations/{locationName}/workflows/{workflowId}/operations/{operationId}", "2019-03-01"),
		},
		{
			Display:  "storageSyncServices",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices", "2019-03-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:        "{storageSyncServiceName}",
					Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-03-01"),
					DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-03-01"),
					PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-03-01"),
					PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-03-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "registeredServers",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{serverId}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-03-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "syncGroups",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{syncGroupName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-03-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-03-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-03-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "cloudEndpoints",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints", "2019-03-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{cloudEndpointName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-03-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-03-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-03-01"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "serverEndpoints",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints", "2019-03-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{serverEndpointName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-03-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-03-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-03-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-03-01"),
													Children:       []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "workflows",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows", "2019-03-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{workflowId}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows/{workflowId}", "2019-03-01"),
									Children: []SwaggerResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.StreamAnalytics/operations", "2016-03-01"),
		},
		{
			Display:  "subscriptions",
			Endpoint: mustGetEndpointInfoFromURL("/subscriptions", "2016-06-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:  "{subscriptionId}",
					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}", "2016-06-01"),
					Children: []SwaggerResourceType{
						{
							Display:  "locations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/locations", "2016-06-01"),
						}},
					SubResources: []SwaggerResourceType{
						{
							Display:  "environments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{environmentName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									Children: []SwaggerResourceType{
										{
											Display:  "accessPolicies",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies", "2017-11-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{accessPolicyName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
												}},
										},
										{
											Display:  "eventSources",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources", "2017-11-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{eventSourceName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
												}},
										},
										{
											Display:  "referenceDataSets",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets", "2017-11-15"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{referenceDataSetName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
												}},
										}},
								}},
						},
						{
							Display:  "trafficmanagerprofiles",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{profileName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									SubResources: []SwaggerResourceType{
										{
											Display:  "{heatMapType}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/heatMaps/{heatMapType}", "2018-04-01"),
										},
										{
											Display:        "{endpointName}",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
										}},
								}},
						},
						{
							Display:  "dedicatedCloudNodes",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{dedicatedCloudNodeName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
								}},
						},
						{
							Display:  "dedicatedCloudServices",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{dedicatedCloudServiceName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
								}},
						},
						{
							Display:  "virtualMachines",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{virtualMachineName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									Children:       []SwaggerResourceType{},
								}},
						},
						{
							Display:  "certificateOrders",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{certificateOrderName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "certificates",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{name}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "domains",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{domainName}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{name}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "hostingEnvironments",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									Children: []SwaggerResourceType{
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
											SubResources: []SwaggerResourceType{
												{
													Display:  "{diagnosticsName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics/{diagnosticsName}", "2018-02-01"),
												}},
										},
										{
											Display:  "inboundNetworkDependenciesEndpoints",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/inboundNetworkDependenciesEndpoints", "2018-02-01"),
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
											Children: []SwaggerResourceType{
												{
													Display:       "default",
													Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													Children: []SwaggerResourceType{
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
														}},
													SubResources: []SwaggerResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metrics", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "operations",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/operations", "2018-02-01"),
										},
										{
											Display:  "outboundNetworkDependenciesEndpoints",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/outboundNetworkDependenciesEndpoints", "2018-02-01"),
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
											SubResources: []SwaggerResourceType{
												{
													Display:       "{workerPoolName}",
													Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													Children: []SwaggerResourceType{
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
														}},
													SubResources: []SwaggerResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metrics", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors/{detectorName}", "2018-02-01"),
												}},
										}},
								},
								{
									Display:  "recommendationHistory",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendationHistory", "2018-02-01"),
								},
								{
									Display:  "recommendations",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations", "2018-02-01"),
									Children: []SwaggerResourceType{},
									SubResources: []SwaggerResourceType{
										{
											Display:  "{name}",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations/{name}", "2018-02-01"),
											Children: []SwaggerResourceType{},
										}},
								}},
						},
						{
							Display:  "serverfarms",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									Children: []SwaggerResourceType{
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
											SubResources: []SwaggerResourceType{
												{
													Display:  "{vnetName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "routes",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{routeName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																}},
														}},
													SubResources: []SwaggerResourceType{
														{
															Display:     "{gatewayName}",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														}},
												}},
										}},
									SubResources: []SwaggerResourceType{
										{
											Display:        "{relayName}",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children: []SwaggerResourceType{
												{
													Display:  "sites",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}/sites", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "certificates",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
								}},
						},
						{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors/{detectorName}", "2018-02-01"),
								}},
						},
						{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{diagnosticCategory}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}", "2018-02-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "analyses",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{analysisName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
													Children: []SwaggerResourceType{},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
													Children: []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "detectors",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors/{detectorName}", "2018-02-01"),
								}},
						},
						{
							Display:  "diagnostics",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:  "{diagnosticCategory}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}", "2018-02-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "analyses",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{analysisName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
													Children: []SwaggerResourceType{},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
													Children: []SwaggerResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "recommendationHistory",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", "2018-02-01"),
						},
						{
							Display:  "recommendations",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", "2018-02-01"),
							Children: []SwaggerResourceType{},
							SubResources: []SwaggerResourceType{
								{
									Display:  "{name}",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", "2018-02-01"),
									Children: []SwaggerResourceType{},
								}},
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", "2018-02-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "default",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", "2018-02-01"),
								}},
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", "2018-02-01"),
							Children: []SwaggerResourceType{
								{
									Display:  "default",
									Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", "2018-02-01"),
								}},
						},
						{
							Display:  "sites",
							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites", "2018-02-01"),
							SubResources: []SwaggerResourceType{
								{
									Display:        "{name}",
									Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									Children: []SwaggerResourceType{
										{
											Display:  "analyzeCustomHostname",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/analyzeCustomHostname", "2018-02-01"),
										},
										{
											Display:  "config",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
											Children: []SwaggerResourceType{
												{
													Display:     "appsettings",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings", "2018-02-01"),
												},
												{
													Display:     "authsettings",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings", "2018-02-01"),
												},
												{
													Display:     "azurestorageaccounts",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts", "2018-02-01"),
												},
												{
													Display:        "backup",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list", "2018-02-01"),
													Verb:           "POST",
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
												},
												{
													Display:     "connectionstrings",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings", "2018-02-01"),
												},
												{
													Display:     "logs",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
												},
												{
													Display:     "metadata",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata", "2018-02-01"),
												},
												{
													Display:  "publishingcredentials",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
													Verb:     "POST",
												},
												{
													Display:     "pushsettings",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings", "2018-02-01"),
												},
												{
													Display:     "slotConfigNames",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
												},
												{
													Display:       "web",
													Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "snapshots",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{snapshotId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																	Children: []SwaggerResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "continuouswebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{webJobName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "deployments",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{id}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}/log", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{domainOwnershipIdentifierName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
												}},
										},
										{
											Display:     "MSDeploy",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
											Children: []SwaggerResourceType{
												{
													Display:  "log",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy/log", "2018-02-01"),
												}},
										},
										{
											Display:  "functions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions", "2018-02-01"),
											Children: []SwaggerResourceType{
												{
													Display:  "token",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/admin/token", "2018-02-01"),
												}},
											SubResources: []SwaggerResourceType{
												{
													Display:        "{functionName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													Children:       []SwaggerResourceType{},
												}},
										},
										{
											Display:  "hostNameBindings",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{hostName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
												}},
										},
										{
											Display:  "hybridConnectionRelays",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionRelays", "2018-02-01"),
										},
										{
											Display:  "hybridconnection",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{entityName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
												}},
										},
										{
											Display:  "instances",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:     "MSDeploy",
													Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "log",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
														}},
												},
												{
													Display:  "processes",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes", "2018-02-01"),
													SubResources: []SwaggerResourceType{
														{
															Display:        "{processId}",
															Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "dump",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																},
																{
																	Display:  "modules",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																	SubResources: []SwaggerResourceType{
																		{
																			Display:  "{baseAddress}",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																		}},
																},
																{
																	Display:  "threads",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																	SubResources: []SwaggerResourceType{
																		{
																			Display:  "{threadId}",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																		}},
																}},
														}},
												}},
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
											Display:        "virtualNetwork",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
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
											SubResources: []SwaggerResourceType{
												{
													Display:        "{premierAddOnName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
												}},
										},
										{
											Display:     "virtualNetworks",
											Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
											PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
										},
										{
											Display:  "processes",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{processId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "dump",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/dump", "2018-02-01"),
														},
														{
															Display:  "modules",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{baseAddress}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																}},
														},
														{
															Display:  "threads",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{threadId}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																}},
														}},
												}},
										},
										{
											Display:  "publicCertificates",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{publicCertificateName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
												}},
										},
										{
											Display:  "siteextensions",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{siteExtensionId}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
												}},
										},
										{
											Display:  "slots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
											Children: []SwaggerResourceType{},
											SubResources: []SwaggerResourceType{
												{
													Display:        "{slot}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "analyzeCustomHostname",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/analyzeCustomHostname", "2018-02-01"),
														},
														{
															Display:  "config",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
															Children: []SwaggerResourceType{
																{
																	Display:     "logs",
																	Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
																	PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
																},
																{
																	Display:       "web",
																	Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "snapshots",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots", "2018-02-01"),
																			SubResources: []SwaggerResourceType{
																				{
																					Display:  "{snapshotId}",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																					Children: []SwaggerResourceType{},
																				}},
																		}},
																}},
														},
														{
															Display:  "continuouswebjobs",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{webJobName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
																	Children:       []SwaggerResourceType{},
																}},
														},
														{
															Display:  "deployments",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{id}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "log",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}/log", "2018-02-01"),
																		}},
																}},
														},
														{
															Display:  "domainOwnershipIdentifiers",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{domainOwnershipIdentifierName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																}},
														},
														{
															Display:     "MSDeploy",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "log",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy/log", "2018-02-01"),
																}},
														},
														{
															Display:  "functions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions", "2018-02-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "token",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/admin/token", "2018-02-01"),
																}},
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{functionName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	Children:       []SwaggerResourceType{},
																}},
														},
														{
															Display:  "hostNameBindings",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{hostName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																}},
														},
														{
															Display:  "hybridConnectionRelays",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionRelays", "2018-02-01"),
														},
														{
															Display:  "hybridconnection",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{entityName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																}},
														},
														{
															Display:  "instances",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:     "MSDeploy",
																	Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
																	PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "log",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
																		}},
																},
																{
																	Display:  "processes",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes", "2018-02-01"),
																	SubResources: []SwaggerResourceType{
																		{
																			Display:        "{processId}",
																			Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
																			DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
																			Children: []SwaggerResourceType{
																				{
																					Display:  "dump",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																				},
																				{
																					Display:  "modules",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																					SubResources: []SwaggerResourceType{
																						{
																							Display:  "{baseAddress}",
																							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																						}},
																				},
																				{
																					Display:  "threads",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																					SubResources: []SwaggerResourceType{
																						{
																							Display:  "{threadId}",
																							Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																						}},
																				}},
																		}},
																}},
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
															Display:        "virtualNetwork",
															Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
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
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{premierAddOnName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																}},
														},
														{
															Display:     "virtualNetworks",
															Endpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
															PutEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
														},
														{
															Display:  "processes",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{processId}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "dump",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/dump", "2018-02-01"),
																		},
																		{
																			Display:  "modules",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules", "2018-02-01"),
																			SubResources: []SwaggerResourceType{
																				{
																					Display:  "{baseAddress}",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																				}},
																		},
																		{
																			Display:  "threads",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads", "2018-02-01"),
																			SubResources: []SwaggerResourceType{
																				{
																					Display:  "{threadId}",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "publicCertificates",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{publicCertificateName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																}},
														},
														{
															Display:  "siteextensions",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{siteExtensionId}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																}},
														},
														{
															Display:  "snapshots",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshots", "2018-02-01"),
															Children: []SwaggerResourceType{
																{
																	Display:  "snapshotsdr",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshotsdr", "2018-02-01"),
																}},
														},
														{
															Display:        "web",
															Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
														},
														{
															Display:  "triggeredwebjobs",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{webJobName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
																	Children: []SwaggerResourceType{
																		{
																			Display:  "history",
																			Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
																			SubResources: []SwaggerResourceType{
																				{
																					Display:  "{id}",
																					Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "usages",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/usages", "2018-02-01"),
														},
														{
															Display:  "virtualNetworkConnections",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:        "{vnetName}",
																	Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	SubResources: []SwaggerResourceType{
																		{
																			Display:       "{gatewayName}",
																			Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																			PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																			PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																		}},
																}},
														},
														{
															Display:  "webjobs",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{webJobName}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs/{webJobName}", "2018-02-01"),
																}},
														}},
													SubResources: []SwaggerResourceType{
														{
															Display:        "{relayName}",
															Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															Children:       []SwaggerResourceType{},
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
														}},
												}},
										},
										{
											Display:  "snapshots",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshots", "2018-02-01"),
											Children: []SwaggerResourceType{
												{
													Display:  "snapshotsdr",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshotsdr", "2018-02-01"),
												}},
										},
										{
											Display:        "web",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
										},
										{
											Display:  "triggeredwebjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{webJobName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													Children: []SwaggerResourceType{
														{
															Display:  "history",
															Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
															SubResources: []SwaggerResourceType{
																{
																	Display:  "{id}",
																	Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																}},
														}},
												}},
										},
										{
											Display:  "usages",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/usages", "2018-02-01"),
										},
										{
											Display:  "virtualNetworkConnections",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:        "{vnetName}",
													Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													SubResources: []SwaggerResourceType{
														{
															Display:       "{gatewayName}",
															Endpoint:      mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PatchEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PutEndpoint:   mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "webjobs",
											Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs", "2018-02-01"),
											SubResources: []SwaggerResourceType{
												{
													Display:  "{webJobName}",
													Endpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs/{webJobName}", "2018-02-01"),
												}},
										}},
									SubResources: []SwaggerResourceType{
										{
											Display:        "{relayName}",
											Endpoint:       mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											DeleteEndpoint: mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											PatchEndpoint:  mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											PutEndpoint:    mustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children:       []SwaggerResourceType{},
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
										}},
								}},
						}},
				}},
		},
		{
			Display:  "tenants",
			Endpoint: mustGetEndpointInfoFromURL("/tenants", "2016-06-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.TimeSeriesInsights/operations", "2017-11-15"),
		},
		{
			Display:  "default",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default", "2018-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.VMwareCloudSimple/operations", "2019-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.CertificateRegistration/operations", "2018-02-01"),
		},
		{
			Display:  "operations",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.DomainRegistration/operations", "2018-02-01"),
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
			Display:     "web",
			Endpoint:    mustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
			PutEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
		},
		{
			Display:  "sourcecontrols",
			Endpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols", "2018-02-01"),
			SubResources: []SwaggerResourceType{
				{
					Display:     "{sourceControlType}",
					Endpoint:    mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
					PutEndpoint: mustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
				}},
		}}

}
