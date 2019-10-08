package handlers

// Register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var Register = []Expander{
	&ResourceGroupResourceExpander{},
	&SubscriptionExpander{},
	&ActionExpander{},
	&MetricsExpander{},
	&SwaggerResourceExpander{},
	&DeploymentsExpander{},
	&ActivityLogExpander{},
	&JSONExpander{},
	&StorageManagementPoliciesExpander{}, // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
	NewContainerRegistryExpander(),       // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
	&ContainerInstanceExpander{},
	&AppInsightsExpander{},
	&AzureKubernetesServiceExpander{},
}
