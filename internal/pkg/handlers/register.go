package handlers

var swaggerResourceExpander *SwaggerResourceExpander

func GetSwaggerResourceExpander() *SwaggerResourceExpander {
	if swaggerResourceExpander == nil {
		swaggerResourceExpander = NewSwaggerResourcesExpander()
		swaggerResourceExpander.AddConfig(NewSwaggerConfigARMResources())
	}
	return swaggerResourceExpander
}

// Register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var Register = []Expander{
	&ResourceGroupResourceExpander{},
	&SubscriptionExpander{},
	&ActionExpander{},
	&MetricsExpander{},
	GetSwaggerResourceExpander(),
	&DeploymentsExpander{},
	&ActivityLogExpander{},
	&JSONExpander{},
	&StorageManagementPoliciesExpander{}, // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
	NewContainerRegistryExpander(),       // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
	&ContainerInstanceExpander{},
	&AppInsightsExpander{},
	&AzureKubernetesServiceExpander{},
}
