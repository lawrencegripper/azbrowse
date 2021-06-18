package expanders

import (
	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

var swaggerResourceExpander *SwaggerResourceExpander
var defaultExpander *DefaultExpander

// GetSwaggerResourceExpander returns the (singleton) instance of SwaggerResourceExpander
func GetSwaggerResourceExpander() *SwaggerResourceExpander {
	return swaggerResourceExpander
}

// GetDefaultExpander returns the (singleton) instance of DefaultExpander
func GetDefaultExpander() *DefaultExpander {
	return defaultExpander
}

// register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var register []Expander

// InitializeExpanders create instances of all the expanders
// needed by the app
func InitializeExpanders(client *armclient.Client, gui *gocui.Gui, commandPanel interfaces.CommandPanel, contentPanel interfaces.ItemWidget) {
	swaggerResourceExpander = NewSwaggerResourcesExpander()
	swaggerResourceExpander.AddAPISet(NewSwaggerAPISetARMResources(client))

	defaultExpander = &DefaultExpander{
		client: client,
	}

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
		&ResourceGraphQueryExpander{
			client: client,
		},
		&ActionExpander{
			client: client,
			gui:    gui,
		},
		&MetricsExpander{
			client: client,
		},
		swaggerResourceExpander,
		&DeploymentsExpander{
			client: client,
		},
		&ActivityLogExpander{
			client: client,
		},
		&JSONExpander{},
		&StorageManagementPoliciesExpander{},                         // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		NewContainerRegistryExpander(client),                         // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		NewStorageBlobExpander(client),                               // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		NewCosmosDbExpander(client, gui, commandPanel, contentPanel), // Needs to be registered after SwaggerResourceExpander as it depends on SwaggerResourceType being set
		&ContainerInstanceExpander{
			client: client,
		},
		&AppInsightsExpander{
			client: client,
		},
		&AzureKubernetesServiceExpander{
			client: client,
		},
		&AzureSearchServiceExpander{
			client: client,
		},
		&AzureDatabricksExpander{
			client: client,
		},
		&DiagnosticSettingsExpander{
			client: client,
		},
		NewTerraformImportExpander(client),
	}
}

func getRegisteredExpanders() []Expander {
	return register
}
