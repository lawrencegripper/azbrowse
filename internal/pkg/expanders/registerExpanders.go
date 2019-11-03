package expanders

var swaggerResourceExpander *SwaggerResourceExpander

// GetSwaggerResourceExpander returns the (singleton) instance of SwaggerResourceExpander
func GetSwaggerResourceExpander() *SwaggerResourceExpander {
	if swaggerResourceExpander == nil {
		swaggerResourceExpander = NewSwaggerResourcesExpander()
		swaggerResourceExpander.AddAPISet(NewSwaggerAPISetARMResources())
	}
	return swaggerResourceExpander
}

// Register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var Register = []Expander{
	&TenantExpander{},
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
