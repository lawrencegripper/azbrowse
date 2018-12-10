package armclient

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
