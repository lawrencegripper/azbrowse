package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

var swaggerResourceExpander *SwaggerResourceExpander

// GetSwaggerResourceExpander returns the (singleton) instance of SwaggerResourceExpander
func GetSwaggerResourceExpander(client *armclient.Client) *SwaggerResourceExpander {
	if swaggerResourceExpander == nil {
		swaggerResourceExpander = NewSwaggerResourcesExpander()
		swaggerResourceExpander.AddAPISet(NewSwaggerAPISetARMResources(client))
	}
	return swaggerResourceExpander
}

// register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var register []Expander

// InitializeExpanders create instances of all the expanders
// needed by the app
func InitializeExpanders(client *armclient.Client) {
	register = []Expander{
		&TenantExpander{
			client: client,
		},
		&ResourceGroupResourceExpander{
			client: client,
		},
		&SubscriptionExpander{
			client: client,
		},
		&ActionExpander{
			client: client,
		},
		&MetricsExpander{
			client: client,
		},
		GetSwaggerResourceExpander(client),
		&DeploymentsExpander{
			client: client,
		},
		&ActivityLogExpander{
			client: client,
		},
		&JSONExpander{},
		&StorageManagementPoliciesExpander{}, // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		NewContainerRegistryExpander(client), // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		&ContainerInstanceExpander{
			client: client,
		},
		&AppInsightsExpander{
			client: client,
		},
		&AzureKubernetesServiceExpander{
			client: client,
		},
	}
}

func getRegisteredExpanders() []Expander {
	return register
}
