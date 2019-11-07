package views

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "Test 1",
			input:    `{"value":[{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/testname1","name":"testname1","location":"northeurope","properties":{"provisioningState":"Succeeded"}},{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/testname2","name":"testname2","location":"westeurope","properties":{"provisioningState":"Succeeded"}},{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/testname3","name":"testname3","location":"westeurope","managedBy":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourcegroups/testrg/providers/Microsoft.ContainerService/managedClusters/testname3","tags":{},"properties":{"provisioningState":"Succeeded"}},{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/NetworkWatcherRG","name":"NetworkWatcherRG","location":"northeurope","properties":{"provisioningState":"Succeeded"}},{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/stable","name":"stable","location":"westeurope","properties":{"provisioningState":"Succeeded"}},{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/testname4","name":"testname4","location":"eastus","properties":{"provisioningState":"Succeeded"}}]}`,
			expected: `{"value":[{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","properties":{"provisioningState":"Succeeded"}},{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","properties":{"provisioningState":"Succeeded"}},{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","id": "HIDDEN_MANAGED_BY",","tags":{},"properties":{"provisioningState":"Succeeded"}},{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","properties":{"provisioningState":"Succeeded"}},{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","properties":{"provisioningState":"Succeeded"}},{"id": "HIDDEN",","name": "HIDDEN-NAME",","location": "HIDDEN-LOCATION",","properties":{"provisioningState":"Succeeded"}}]}`,
		},
		{
			desc:     "Test 2",
			input:    `{"value":[{"id":"/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3","authorizationSource":"Legacy, RoleBased","subscriptionId":"1724ed8f-d6ne-4552-l73e-0567920518f3","displayName":"Joni's Azure Internal Subscription","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id":"/subscriptions/69e18cb1-0da9-4b34-bd6e-53c4bd4e91f0","authorizationSource":"Legacy","subscriptionId":"69e18cb1-0da9-4b34-bd6e-53c4bd4e91f0","displayName":"CSE UK","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id":"/subscriptions/0d769062-023f-4c80-9b9a-f19d80b97bf6","authorizationSource":"RoleBased","subscriptionId":"0d769062-023f-4c80-9b9a-f19d80b97bf6","displayName":"Microsoft Azure Internal Consumption","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id":"/subscriptions/d21a0e9f-5e29-4b39-8ba5-0e189bc5fe2d","authorizationSource":"RoleBased","subscriptionId":"d21a0e9f-5e29-4b39-8ba5-0e189bc5fe2d","displayName":"Edge DevTools Client","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id":"/subscriptions/04f7ec88-8e28-41ed-8537-5e17766001f5","authorizationSource":"RoleBased","subscriptionId":"04f7ec88-8e28-41ed-8537-5e17766001f5","displayName":"Strategic Engagements Developers Research","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id":"/subscriptions/b9a40aba-781c-4814-84c3-c5efdba10b86","authorizationSource":"RoleBased","subscriptionId":"b9a40aba-781c-4814-84c3-c5efdba10b86","displayName":"staskew-OneWeek2019","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}}]}`,
			expected: `{"value":[{"id": "HIDDEN",","authorizationSource":"Legacy, RoleBased","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"Joni's Azure Internal Subscription","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id": "HIDDEN",","authorizationSource":"Legacy","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"CSE UK","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id": "HIDDEN",","authorizationSource":"RoleBased","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"Microsoft Azure Internal Consumption","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id": "HIDDEN",","authorizationSource":"RoleBased","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"Edge DevTools Client","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id": "HIDDEN",","authorizationSource":"RoleBased","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"Strategic Engagements Developers Research","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}},{"id": "HIDDEN",","authorizationSource":"RoleBased","subscriptionId":"00000000-0000-0000-0000-HIDDEN000000","displayName":"staskew-OneWeek2019","state":"Enabled","subscriptionPolicies":{"locationPlacementId":"Internal_2014-09-01","quotaId":"Internal_2014-09-01","spendingLimit":"Off"}}]}`,
		},
		{
			desc: "Test 3",
			input: `{
				"name": "accdev-nic",
				"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Network/networkInterfaces/accdev-nic",
				"etag": "W/\"1724ed8f-d6ne-4552-l73e-0567920518f3\"",
				"location": "westeurope",
				"properties": {
					"provisioningState": "Succeeded",
					"resourceGuid": "1724ed8f-d6ne-4552-l73e-0567920518f3",
					"ipConfigurations": [{
						"name": "ipConfigNode",
						"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Network/networkInterfaces/accdev-nic/ipConfigurations/ipConfigNode",
						"etag": "W/\"1724ed8f-d6ne-4552-l73e-0567920518f3\"",
						"type": "Microsoft.Network/networkInterfaces/ipConfigurations",
						"properties": {
							"provisioningState": "Succeeded",
							"privateIPAddress": "10.4.0.4",
							"privateIPAllocationMethod": "Dynamic",
							"publicIPAddress": {
								"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Network/publicIPAddresses/accdev-ip"
							},
							"subnet": {
								"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Network/virtualNetworks/VirtualNetwork/subnets/Subnet-1"
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
						"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Network/networkSecurityGroups/accdev-nsg"
					},
					"primary": true,
					"virtualMachine": {
						"id": "/subscriptions/1724ed8f-d6ne-4552-l73e-0567920518f3/resourceGroups/accdev/providers/Microsoft.Compute/virtualMachines/accdev"
					},
					"hostedWorkloads": [],
					"tapConfigurations": []
				},
				"type": "Microsoft.Network/networkInterfaces"
			}`,
			expected: `{
				"name": "HIDDEN-NAME",
				"id": "HIDDEN",
				"etag": "W/\"1724ed8f-d6ne-4552-l73e-0567920518f3\"",
				"location": "HIDDEN-LOCATION",
				"properties": {
					"provisioningState": "Succeeded",
					"resourceGuid": "1724ed8f-d6ne-4552-l73e-0567920518f3",
					"ipConfigurations": [{
						"name": "HIDDEN-NAME",
						"id": "HIDDEN",
						"etag": "W/\"1724ed8f-d6ne-4552-l73e-0567920518f3\"",
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
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			output := stripSecretVals(test.input)
			assert.Equal(t, output, test.expected)
		})
	}
}
