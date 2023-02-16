package armclient

// ResourceResponse Resources list rest type
type ResourceResponse struct {
	Resources []Resource `json:"value"`
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
