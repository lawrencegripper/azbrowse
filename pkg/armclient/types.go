package armclient

import (
	"time"
)

// SubResponse Subscriptions REST type
type SubResponse struct {
	Subs []struct {
		ID                   string `json:"id"`
		SubscriptionID       string `json:"subscriptionId"`
		DisplayName          string `json:"displayName"`
		State                string `json:"state"`
		SubscriptionPolicies struct {
			LocationPlacementID string `json:"locationPlacementId"`
			QuotaID             string `json:"quotaId"`
			SpendingLimit       string `json:"spendingLimit"`
		} `json:"subscriptionPolicies"`
	} `json:"value"`
}

// ResourceGroupResponse ResourceGroup rest type
type ResourceGroupResponse struct {
	Groups []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Location   string `json:"location"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
		} `json:"properties"`
	} `json:"value"`
}

// ResourceResponse Resources list rest type
type ResourceResponse struct {
	Resources []Resource `json:"value"`
}

// ResourceQueryResponse list query response
type ResourceQueryResponse struct {
	TotalRecords int `json:"totalRecords"`
	Count        int `json:"count"`
	Data         struct {
		Columns []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Rows [][]interface{} `json:"rows"`
	} `json:"data"`
	Facets          []interface{} `json:"facets"`
	ResultTruncated string        `json:"resultTruncated"`
}

// Resource is a resource in azure
type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Sku  struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	} `json:"sku"`
	Kind       string `json:"kind"`
	Location   string `json:"location"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
	} `json:"properties"`
}

// ProvidersResponse providers list rest type
type ProvidersResponse struct {
	Providers []struct {
		ID            string `json:"id"`
		Namespace     string `json:"namespace"`
		Authorization struct {
			ApplicationID    string `json:"applicationId"`
			RoleDefinitionID string `json:"roleDefinitionId"`
		} `json:"authorization,omitempty"`
		ResourceTypes []struct {
			ResourceType string        `json:"resourceType"`
			Locations    []interface{} `json:"locations"`
			APIVersions  []string      `json:"apiVersions"`
		} `json:"resourceTypes"`
		RegistrationState string `json:"registrationState"`
		Authorizations    []struct {
			ApplicationID    string `json:"applicationId"`
			RoleDefinitionID string `json:"roleDefinitionId"`
		} `json:"authorizations,omitempty"`
	} `json:"value"`
}

// OperationsRequest list the actions that can be performed
type OperationsRequest struct {
	DisplayName string `json:"displayName"`
	Operations  []struct {
		Name         string      `json:"name"`
		DisplayName  string      `json:"displayName"`
		Description  string      `json:"description"`
		Origin       interface{} `json:"origin"`
		Properties   interface{} `json:"properties"`
		IsDataAction bool        `json:"isDataAction"`
	} `json:"operations"`
	ResourceTypes []struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Operations  []struct {
			Name         string      `json:"name"`
			DisplayName  string      `json:"displayName"`
			Description  string      `json:"description"`
			Origin       interface{} `json:"origin"`
			Properties   interface{} `json:"properties"`
			IsDataAction bool        `json:"isDataAction"`
		} `json:"operations"`
	} `json:"resourceTypes"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// SubResourcesResponse is the response from the /resources call on a sub
type SubResourcesResponse struct {
	Resources []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Location string `json:"location"`
		Tags     struct {
			Test string `json:"test"`
		} `json:"tags,omitempty"`
		Sku struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku,omitempty"`
		Kind string `json:"kind,omitempty"`
		Plan struct {
			Name          string `json:"name"`
			PromotionCode string `json:"promotionCode"`
			Product       string `json:"product"`
			Publisher     string `json:"publisher"`
		} `json:"plan,omitempty"`
	} `json:"value"`
	NextLink string `json:"nextLink"`
}

// DeploymentsResponse is returned by a request for deployments in an RG
type DeploymentsResponse struct {
	Value []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Properties struct {
			CorrelationID string `json:"correlationId"`
			Dependencies  []struct {
				DependsOn []struct {
					ID           string `json:"id"`
					ResourceName string `json:"resourceName"`
					ResourceType string `json:"resourceType"`
				} `json:"dependsOn"`
				ID           string `json:"id"`
				ResourceName string `json:"resourceName"`
				ResourceType string `json:"resourceType"`
			} `json:"dependencies"`
			Duration        string `json:"duration"`
			Mode            string `json:"mode"`
			OutputResources []struct {
				ID string `json:"id"`
			} `json:"outputResources"`
			Outputs    map[string]interface{} `json:"outputs"`
			Parameters map[string]interface{} `json:"parameters"`
			Providers  []struct {
				Namespace     string `json:"namespace"`
				ResourceTypes []struct {
					Locations    []string `json:"locations"`
					ResourceType string   `json:"resourceType"`
				} `json:"resourceTypes"`
			} `json:"providers"`
			ProvisioningState string                 `json:"provisioningState"`
			TemplateHash      string                 `json:"templateHash"`
			Template          map[string]interface{} `json:"template"`
			TemplateLink      struct {
				ContentVersion string `json:"contentVersion"`
				URI            string `json:"uri"`
			} `json:"templateLink"`
			Timestamp string `json:"timestamp"`
		} `json:"properties"`
	} `json:"value"`
}

// DeploymentOperationsResponse is a struct to enable splitting out json value array
type DeploymentOperationsResponse struct {
	Value []struct {
		ID          string `json:"id"`
		OperationID string `json:"operationId"`
		Properties  struct {
			StatusCode            string      `json:"statusCode"`
			StatusMessage         interface{} `json:"statusMessage"`
			Timestamp             string      `json:"timestamp"`
			Duration              string      `json:"duration"`
			ProvisioningOperation string      `json:"provisioningOperation"`
			ProvisioningState     string      `json:"provisioningState"`
			TrackingID            string      `json:"trackingId"`
			TargetResource        struct {
				ID           string `json:"id"`
				ResourceType string `json:"resourceType"`
				ResourceName string `json:"resourceName"`
			} `json:"targetResource"`
		} `json:"properties"`
	} `json:"value"`
}

// ActivityLogResource is returned when requesting activity logs for an RG
type ActivityLogResource struct {
	Value []struct {
		Authorization struct {
			Action string `json:"action"`
			Scope  string `json:"scope"`
		} `json:"authorization"`
		Caller   string `json:"caller"`
		Channels string `json:"channels"`
		Claims   struct {
			Aud                                                       string `json:"aud"`
			Iss                                                       string `json:"iss"`
			Iat                                                       string `json:"iat"`
			Nbf                                                       string `json:"nbf"`
			Exp                                                       string `json:"exp"`
			Aio                                                       string `json:"aio"`
			Appid                                                     string `json:"appid"`
			Appidacr                                                  string `json:"appidacr"`
			HTTPSchemasMicrosoftComIdentityClaimsIdentityprovider     string `json:"http://schemas.microsoft.com/identity/claims/identityprovider"`
			HTTPSchemasMicrosoftComIdentityClaimsObjectidentifier     string `json:"http://schemas.microsoft.com/identity/claims/objectidentifier"`
			HTTPSchemasXmlsoapOrgWs200505IdentityClaimsNameidentifier string `json:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier"`
			HTTPSchemasMicrosoftComIdentityClaimsTenantid             string `json:"http://schemas.microsoft.com/identity/claims/tenantid"`
			Uti                                                       string `json:"uti"`
			Ver                                                       string `json:"ver"`
		} `json:"claims"`
		CorrelationID string `json:"correlationId"`
		Description   string `json:"description"`
		EventDataID   string `json:"eventDataId"`
		EventName     struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"eventName"`
		Category struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"category"`
		ID                   string `json:"id"`
		Level                string `json:"level"`
		ResourceGroupName    string `json:"resourceGroupName"`
		ResourceProviderName struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"resourceProviderName"`
		ResourceID   string `json:"resourceId"`
		ResourceType struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"resourceType"`
		OperationID   string `json:"operationId"`
		OperationName struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"operationName"`
		Properties struct {
			IsComplianceCheck string `json:"isComplianceCheck"`
			ResourceLocation  string `json:"resourceLocation"`
			Ancestors         string `json:"ancestors"`
			Policies          string `json:"policies"`
		} `json:"properties,omitempty"`
		Status struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"status"`
		SubStatus struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"subStatus"`
		EventTimestamp      time.Time `json:"eventTimestamp"`
		SubmissionTimestamp time.Time `json:"submissionTimestamp"`
		SubscriptionID      string    `json:"subscriptionId"`
		TenantID            string    `json:"tenantId"`
		HTTPRequest         struct {
			ClientRequestID string `json:"clientRequestId"`
			ClientIPAddress string `json:"clientIpAddress"`
			Method          string `json:"method"`
		} `json:"httpRequest,omitempty"`
	} `json:"value"`
}

// ContainerGroupResponse is the response to a get request on a container group
type ContainerGroupResponse struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Containers []struct {
			Name       string `json:"name"`
			Properties struct {
				Command              []interface{} `json:"command"`
				EnvironmentVariables []interface{} `json:"environmentVariables"`
				Image                string        `json:"image"`
				Ports                []struct {
					Port int `json:"port"`
				} `json:"ports"`
				InstanceView struct {
					RestartCount int `json:"restartCount"`
					CurrentState struct {
						State        string    `json:"state"`
						StartTime    time.Time `json:"startTime"`
						DetailStatus string    `json:"detailStatus"`
					} `json:"currentState"`
					Events []struct {
						Count          int       `json:"count"`
						FirstTimestamp time.Time `json:"firstTimestamp"`
						LastTimestamp  time.Time `json:"lastTimestamp"`
						Name           string    `json:"name"`
						Message        string    `json:"message"`
						Type           string    `json:"type"`
					} `json:"events"`
				} `json:"instanceView"`
				Resources struct {
					Requests struct {
						CPU        float64 `json:"cpu"`
						MemoryInGB float64 `json:"memoryInGB"`
					} `json:"requests"`
				} `json:"resources"`
				VolumeMounts []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly"`
				} `json:"volumeMounts"`
			} `json:"properties"`
		} `json:"containers"`
		ImageRegistryCredentials []struct {
			Server   string `json:"server"`
			Username string `json:"username"`
		} `json:"imageRegistryCredentials"`
		IPAddress struct {
			IP    string `json:"ip"`
			Ports []struct {
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
			Type string `json:"type"`
		} `json:"ipAddress"`
		OsType            string `json:"osType"`
		ProvisioningState string `json:"provisioningState"`
		Volumes           []struct {
			AzureFile struct {
				ReadOnly           bool   `json:"readOnly"`
				ShareName          string `json:"shareName"`
				StorageAccountName string `json:"storageAccountName"`
			} `json:"azureFile"`
			Name string `json:"name"`
		} `json:"volumes"`
	} `json:"properties"`
	Type string `json:"type"`
}
