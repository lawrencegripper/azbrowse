package armclient

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

// ResourceReseponse Resources list rest type
type ResourceReseponse struct {
	Resources []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Sku  struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku"`
		Kind     string `json:"kind"`
		Location string `json:"location"`
		Tags     struct {
			MsResourceUsage string `json:"ms-resource-usage"`
		} `json:"tags"`
	} `json:"value"`
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
