package armclient

// MetricNamespaceResponse https://docs.microsoft.com/en-us/rest/api/monitor/metricnamespaces/list
type MetricNamespaceResponse struct {
	Value []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Type           string `json:"type"`
		Classification string `json:"classification"`
		Properties     struct {
			MetricNamespaceName string `json:"metricNamespaceName"`
		} `json:"properties"`
	} `json:"value"`
}

// MetricsListResponse https://docs.microsoft.com/en-us/rest/api/monitor/metricdefinitions/list
type MetricsListResponse struct {
	Value []struct {
		ID         string `json:"id"`
		ResourceID string `json:"resourceId"`
		Namespace  string `json:"namespace"`
		Name       struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"name"`
		IsDimensionRequired       bool     `json:"isDimensionRequired"`
		Unit                      string   `json:"unit"`
		PrimaryAggregationType    string   `json:"primaryAggregationType"`
		SupportedAggregationTypes []string `json:"supportedAggregationTypes"`
		MetricAvailabilities      []struct {
			TimeGrain string `json:"timeGrain"`
			Retention string `json:"retention"`
		} `json:"metricAvailabilities"`
		Dimensions []struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"dimensions"`
	} `json:"value"`
}
