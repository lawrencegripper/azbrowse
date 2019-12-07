package views

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ensures stripSecretValues can successfully replace targeted values
func Test_stripSecretVals_SuccessfullyReplaceSecrets(t *testing.T) {
	for _, test := range successfullyReplaceSecretsTestData {
		// shadow test value to ensure assertion failures
		// print the correct test.expected value
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			actual := stripSecretVals(test.input)
			assert.Equal(t, test.expected, actual)

			// while not a requirement for stripSecretVals, all our tests are valid json
			// so the result should also be valid json
			json.Valid([]byte(actual))
		})
	}
}

var successfullyReplaceSecretsTestData = []struct {
	desc     string
	input    string
	expected string
}{
	{
		desc: "multi-node",
		input: `
		{
			"value": [
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/testname1",
					"name": "testname1",
					"location": "northeurope",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/testname2",
					"name": "testname2",
					"location": "westeurope",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/testname3",
					"name": "testname3",
					"location": "westeurope",
					"managedBy": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourcegroups/testrg/providers/Microsoft.ContainerService/managedClusters/testname3",
					"tags": {},
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/NetworkWatcherRG",
					"name": "NetworkWatcherRG",
					"location": "northeurope",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/stable",
					"name": "stable",
					"location": "westeurope",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/testname4",
					"name": "testname4",
					"location": "eastus",
					"properties": {
						"provisioningState": "Succeeded"
					}
				}
			]
		}`,
		expected: `
		{
			"value": [
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"managedBy": "HIDDEN_MANAGED_BY",
					"tags": {},
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"properties": {
						"provisioningState": "Succeeded"
					}
				},
				{
					"id": "HIDDEN",
					"name": "HIDDEN-NAME",
					"location": "HIDDEN-LOCATION",
					"properties": {
						"provisioningState": "Succeeded"
					}
				}
			]
		}`,
	},
	{
		desc: "subscriptions",
		input: `
		{
			"value": [
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "Legacy, RoleBased",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "Joni's Azure Internal Subscription",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "Legacy",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "CSE UK",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "RoleBased",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "Microsoft Azure Internal Consumption",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "RoleBased",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "Edge DevTools Client",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "RoleBased",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "Strategic Engagements Developers Research",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498",
					"authorizationSource": "RoleBased",
					"subscriptionId": "abcdef12-0751-dead-beef-6150896ac498",
					"displayName": "staskew-OneWeek2019",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				}
			]
		}`,
		expected: `
		{
			"value": [
				{
					"id": "HIDDEN",
					"authorizationSource": "Legacy, RoleBased",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "Joni's Azure Internal Subscription",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "HIDDEN",
					"authorizationSource": "Legacy",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "CSE UK",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "HIDDEN",
					"authorizationSource": "RoleBased",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "Microsoft Azure Internal Consumption",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "HIDDEN",
					"authorizationSource": "RoleBased",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "Edge DevTools Client",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "HIDDEN",
					"authorizationSource": "RoleBased",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "Strategic Engagements Developers Research",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				},
				{
					"id": "HIDDEN",
					"authorizationSource": "RoleBased",
					"subscriptionId": "00000000-0000-0000-0000-HIDDEN000000",
					"displayName": "staskew-OneWeek2019",
					"state": "Enabled",
					"subscriptionPolicies": {
						"locationPlacementId": "Internal_2014-09-01",
						"quotaId": "Internal_2014-09-01",
						"spendingLimit": "Off"
					}
				}
			]
		}`,
	},
	{
		desc: "Microsoft.Network/networkInterfaces",
		input: `{
			"name": "accdev-nic",
			"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Network/networkInterfaces/accdev-nic",
			"etag": "W/\"abcdef12-0751-dead-beef-6150896ac498\"",
			"location": "westeurope",
			"properties": {
				"provisioningState": "Succeeded",
				"resourceGuid": "abcdef12-0751-dead-beef-6150896ac498",
				"ipConfigurations": [{
					"name": "ipConfigNode",
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Network/networkInterfaces/accdev-nic/ipConfigurations/ipConfigNode",
					"etag": "W/\"abcdef12-0751-dead-beef-6150896ac498\"",
					"type": "Microsoft.Network/networkInterfaces/ipConfigurations",
					"properties": {
						"provisioningState": "Succeeded",
						"privateIPAddress": "10.4.0.4",
						"privateIPAllocationMethod": "Dynamic",
						"publicIPAddress": {
							"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Network/publicIPAddresses/accdev-ip"
						},
						"subnet": {
							"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Network/virtualNetworks/VirtualNetwork/subnets/Subnet-1"
						},
						"primary": true,
						"privateIPAddressVersion": "IPv4"
					}
				}],
				"dnsSettings": {
					"dnsServers": [],
					"appliedDnsServers": [],
					"internalDomainNameSuffix": "xidg3h51wk3evancrgy2smiwye.ax.internal.cloudapp.net"
				},
				"macAddress": "00-0D-3A-4A-54-AC",
				"enableAcceleratedNetworking": false,
				"enableIPForwarding": false,
				"networkSecurityGroup": {
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Network/networkSecurityGroups/accdev-nsg"
				},
				"primary": true,
				"virtualMachine": {
					"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/accdev/providers/Microsoft.Compute/virtualMachines/accdev"
				},
				"hostedWorkloads": [],
				"tapConfigurations": []
			},
			"type": "Microsoft.Network/networkInterfaces"
		}`,
		expected: `{
			"name": "HIDDEN-NAME",
			"id": "HIDDEN",
			"etag": "W/\"00000000-0000-0000-0000-HIDDEN000000\"",
			"location": "HIDDEN-LOCATION",
			"properties": {
				"provisioningState": "Succeeded",
				"resourceGuid": "00000000-0000-0000-0000-HIDDEN000000",
				"ipConfigurations": [{
					"name": "HIDDEN-NAME",
					"id": "HIDDEN",
					"etag": "W/\"00000000-0000-0000-0000-HIDDEN000000\"",
					"type": "Microsoft.Network/networkInterfaces/ipConfigurations",
					"properties": {
						"provisioningState": "Succeeded",
						"privateIPAddress": "10.4.0.4",
						"privateIPAllocationMethod": "Dynamic",
						"publicIPAddress": {
							"id": "HIDDEN"
						},
						"subnet": {
							"id": "HIDDEN"
						},
						"primary": true,
						"privateIPAddressVersion": "IPv4"
					}
				}],
				"dnsSettings": {
					"dnsServers": [],
					"appliedDnsServers": [],
					"internalDomainNameSuffix": "xidg3h51wk3evancrgy2smiwye.ax.internal.cloudapp.net"
				},
				"macAddress": "00-0D-3A-4A-54-AC",
				"enableAcceleratedNetworking": false,
				"enableIPForwarding": false,
				"networkSecurityGroup": {
					"id": "HIDDEN"
				},
				"primary": true,
				"virtualMachine": {
					"id": "HIDDEN"
				},
				"hostedWorkloads": [],
				"tapConfigurations": []
			},
			"type": "Microsoft.Network/networkInterfaces"
		}`,
	},
	{
		desc: "webapp/config/authsettings",
		input: `
		{
			"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/my-rg/providers/Microsoft.Web/sites/my-site/config/authsettings",
			"name": "authsettings",
			"type": "Microsoft.Web/sites/config",
			"location": "East US",
			"tags": {
				"hidden-related:/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/my-rg/providers/Microsoft.Web/serverfarms/my-site": "Resource"
			},
			"properties": {
				"enabled": false,
				"unauthenticatedClientAction": null,
				"tokenStoreEnabled": null,
				"allowedExternalRedirectUrls": null,
				"defaultProvider": null,
				"clientId": "test-id",
				"clientSecret": "non-public-secret",
				"clientSecretCertificateThumbprint": null,
				"issuer": null,
				"allowedAudiences": null,
				"additionalLoginParams": null,
				"isAadAutoProvisioned": false,
				"googleClientId": null,
				"googleClientSecret": null,
				"googleOAuthScopes": null,
				"facebookAppId": null,
				"facebookAppSecret": null,
				"facebookOAuthScopes": null,
				"twitterConsumerKey": "twit-key",
				"twitterConsumerSecret": "twit-secret",
				"microsoftAccountClientId": null,
				"microsoftAccountClientSecret": null,
				"microsoftAccountOAuthScopes": null
			}
		}`,
		expected: `
		{
			"id": "HIDDEN",
			"name": "HIDDEN-NAME",
			"type": "Microsoft.Web/sites/config",
			"location": "HIDDEN-LOCATION",
			"tags": {
				"hidden-related:/subscriptions/00000000-0000-0000-0000-HIDDEN000000/resourceGroups/my-rg/providers/Microsoft.Web/serverfarms/my-site": "Resource"
			},
			"properties": {
				"enabled": false,
				"unauthenticatedClientAction": null,
				"tokenStoreEnabled": null,
				"allowedExternalRedirectUrls": null,
				"defaultProvider": null,
				"clientId": "test-id",
				"clientSecret": "HIDDEN-SECRET",
				"clientSecretCertificateThumbprint": null,
				"issuer": null,
				"allowedAudiences": null,
				"additionalLoginParams": null,
				"isAadAutoProvisioned": false,
				"googleClientId": null,
				"googleClientSecret": null,
				"googleOAuthScopes": null,
				"facebookAppId": null,
				"facebookAppSecret": null,
				"facebookOAuthScopes": null,
				"twitterConsumerKey": "HIDDEN-KEY",
				"twitterConsumerSecret": "HIDDEN-SECRET",
				"microsoftAccountClientId": null,
				"microsoftAccountClientSecret": null,
				"microsoftAccountOAuthScopes": null
			}
		}`,
	},
	{
		desc: "webapp/config/connectionstrings",
		input: `
		{
			"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/my-rg/providers/Microsoft.Web/sites/my-site/config/connectionstrings",
			"name": "connectionstrings",
			"type": "Microsoft.Web/sites/config",
			"location": "Central US",
			"properties": {
				"ProductionDatabase": {
					"value": "msql://real-production-system.com:1929",
					"type": "MySql"
				},
				"ProductionDatabase2": {
					"value": "real",
					"type": "SqlAzure"
				}
			}
		}`,
		expected: `
		{
			"id": "HIDDEN",
			"name": "HIDDEN-NAME",
			"type": "Microsoft.Web/sites/config",
			"location": "HIDDEN-LOCATION",
			"properties": {
				"ProductionDatabase": {
					"value": "HIDDEN",
					"type": "MySql"
				},
				"ProductionDatabase2": {
					"value": "HIDDEN",
					"type": "SqlAzure"
				}
			}
		}`,
	},
	{
		desc: "webapp/config/appsettings",
		input: `
		{
			"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/my-rg/providers/Microsoft.Web/sites/my-site/config/appsettings",
			"name": "appsettings",
			"type": "Microsoft.Web/sites/config",
			"location": "Central US",
			"properties": {
				"ProductionSetting1": "PublicValue",
				"ProductionSecret": "HiddenVal",
				"ProductionKey": "HiddenVal",
				"ProductionSetting2": "StillPublicValue"
			}
		}`,
		expected: `
		{
			"id": "HIDDEN",
			"name": "HIDDEN-NAME",
			"type": "Microsoft.Web/sites/config",
			"location": "HIDDEN-LOCATION",
			"properties": {
				"ProductionSetting1": "PublicValue",
				"ProductionSecret": "HIDDEN-SECRET",
				"ProductionKey": "HIDDEN-KEY",
				"ProductionSetting2": "StillPublicValue"
			}
		}`,
	},
	{
		desc: "webapp/config/publishingcredentials",
		input: `
		{
			"id": "/subscriptions/abcdef12-0751-dead-beef-6150896ac498/resourceGroups/my-rg/providers/Microsoft.Web/sites/my-site/publishingcredentials/$unique-name-fj2oifjoijfoji23o",
			"name": "my-site",
			"type": "Microsoft.Web/sites/publishingcredentials",
			"location": "Central US",
			"properties": {
				"name": null,
				"publishingUserName": "$my-site",
				"publishingPassword": "NF0Pdr8sd3yuCqfdpuun9M5ckBPlhXTu7C1jadPGo1jmBExYBsNNzNvu5oY6",
				"publishingPasswordHash": null,
				"publishingPasswordHashSalt": null,
				"metadata": null,
				"isDeleted": false,
				"scmUri": "https://$my-site:NF0Pdr8sd3yuCqfdpuun9M5ckBPlhXTu7C1jadPGo1jmBExYBsNNzNvu5oY6@my-site.scm.azurewebsites.net"
			}
		}`,
		expected: `
		{
			"id": "HIDDEN",
			"name": "HIDDEN-NAME",
			"type": "Microsoft.Web/sites/publishingcredentials",
			"location": "HIDDEN-LOCATION",
			"properties": {
				"name": null,
				"publishingUserName": "$my-site",
				"publishingPassword": "HIDDEN-PASSWORD",
				"publishingPasswordHash": null,
				"publishingPasswordHashSalt": null,
				"metadata": null,
				"isDeleted": false,
				"scmUri": "HIDDEN-URI"
			}
		}`,
	},
	{
		desc:     "webapp/functions/admin/token",
		input:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NzU0NTkzMzUsImV4cCI6MTU3NTQ1OTYzNSwiaWF0IjoxNTc1NDU5MzM1LCJpc3MiOiJodHRwczovL3N1cGVyLXVuaXF1ZS1mdW5jdGlvbi1uYW1lLnNjbS5henVyZXdlYnNpdGVzLm5ldCIsImF1ZCI6Imh0dHBzOi8vc3VwZXItdW5pcXVlLWZ1bmN0aW9uLW5hbWUuYXp1cmV3ZWJzaXRlcy5uZXQvYXp1cmVmdW5jdGlvbnMifQ.ebM87FGp34_FTiv-VUNTZpOzgp6Ie_wPoNQvxv6yw54",
		expected: "HIDDEN-JWT",
	},
	{
		desc: "search/listQueryKeys",
		input: `
		{
			"value": [
				{
					"name": null,
					"key": "9290DE05303CAA205305E92E6284ED91"
				}
			],
			"nextLink": null
		}`,
		expected: `
		{
			"value": [
				{
					"name": null,
					"key": "HIDDEN-KEY"
				}
			],
			"nextLink": null
		}`,
	},
	{
		desc: "search/listAdminKeys",
		input: `
		{
			"primaryKey": "723D027F1AD90B264E49DB239770F20A",
			"secondaryKey": "723D027F1AD90B264E49DB239770F20A",
		}
		`,
		expected: `
		{
			"primaryKey": "HIDDEN-KEY",
			"secondaryKey": "HIDDEN-KEY",
		}
		`,
	},
	{
		desc: "sshkey/4096",
		input: `
		{
			"publicKeyField": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDHUuJ+3W+Vjdu7bW6DQCpr8nUanXR2239Tpt2V5KVIe/uR6E/+U0cpvUrvYKrciYZVlR+tFXz46gEAad6d8JQ5wP7iPhHt8Tdc2zpXXawycOaxVzJEX24xA/YPYQJmZnqtipIH0nzL5y7lLQMkXAuJu0omUYHeTZt0mafs3QrKiFXTqKar1hrUvcCrOyoRlbK0F+1qSQotV/2Gv5pC3/rkMP4ZvCd20U3gmjKUpgWwoB8yqNH9ISCAjjzslwKTXnWJEW8V2FOx8GbqN33wTWXc0+VZ8g59l2vde1CchV2twDtr8CeMatOrSLHCcNnCHW7Et9CuCTwC5Hm65x2eaVajwLez1dFRizhNTfWN4e2s0RexGQW34jvUyd8jTGH2Fny3Vv/apoQiZXNvC8lvLtR8S4doIBD/NEPZnBtwPVA6t7X4EvPIJM4Tty+3PUx0Wx7yvA8HrrknnWzwPcIFjmjcv70Ly3etWo1CcYMw2eK8C7XUmfsYKSbd3qQKQlj1Axb/+muIvAuED/Q87qG4MPnzfTzVpxA5JUtx9Qjhx1nCL2sY47QXwva3HFgH3fdg2sqDoDIcUx6lxkz32g4q0B6PdZxXSCLSooVRjDEAmlDJ1uRdGWwa7D3+Us+RWVh65rKST2/hz610b8CiFh7HTdt9sgs0MEDwP1xxD7sYCR3EAw== ben@BENTOWER",
		}
		`,
		expected: `
		{
			"publicKeyField": "SSH-PUBLIC-KEY-HIDDEN",
		}
		`,
	},
	{
		desc: "sshkey/2048",
		input: `
		{
			"publicKeyField": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAKPZbZejaPKhO5VXesh0+F8QOV6vUIm9VISnoJpHofsysEZiVEhjP82i5SnLCUIT5E94/6GUd2lM9uqW7WP/XTVqNesqufRVfn77etSzpsNmn0odSGu9ESE3/ZILgIcw91wdsD5K49nhjO4rddE+x9Ugn8yGvz+QRRBK8cdZFHOCyGHIh9ottx+hXl8sA4utYB2YExYK2/izRY8N9yxH71aM5hfXFM0tNBpc2TPydUCbubo59f0rAtdetuSmyvn61vVxQlDOp/BG8HvLaUfUX+aaldY6BfXoBAKftpwHaZA1foDnyED+IzuiePc1HrbGapqpszLMLCrHMnQZ1Re2b ben@BENTOWER",
		}
		`,
		expected: `
		{
			"publicKeyField": "SSH-PUBLIC-KEY-HIDDEN",
		}
		`,
	},
}
