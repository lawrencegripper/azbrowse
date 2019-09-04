package handlers

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// MetricsExpander expands the data-plane aspects of the Microsoft.Insights RP
type MetricsExpander struct {
}

// Name returns the name of the expander
func (e *MetricsExpander) Name() string {
	return "MetricsExpander"
}

// DoesExpand checks if this is a storage account
func (e *MetricsExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == ResourceType {
		return true, nil
	}

	if currentItem.ItemType == "metrics.metricdefinition" {
		return true, nil
	}

	if currentItem.ItemType == "metrics.graph" {
		return true, nil
	}

	return false, nil
}

// Expand adds items for metrics to the list
func (e *MetricsExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	if currentItem.ItemType == "metrics.metricdefinition" {
		data, err := armclient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "MetricsExpander request metricDefinitions",
			}
		}

		var metricsListResponse armclient.MetricsListResponse
		err = json.Unmarshal([]byte(data), &metricsListResponse)
		if err != nil {
			panic(err)
		}

		newItems := []*TreeNode{}

		for _, metric := range metricsListResponse.Value {
			newItems = append(newItems, &TreeNode{
				Name:           metric.Name.Value,
				Display:        metric.Name.Value + "\n  " + style.Subtle("Unit: "+metric.Unit),
				ID:             currentItem.ID,
				Parentid:       currentItem.ID,
				ExpandURL:      currentItem.ID + "/providers/microsoft.Insights/metrics?timespan=" + time.Now().UTC().Add(-3*time.Hour).Format("2006-01-02T15:04:05.000Z") + "/" + time.Now().UTC().Format("2006-01-02T15:04:05.000Z") + "&interval=PT5M&metricnames=" + metric.Name.Value + "&aggregation=" + metric.PrimaryAggregationType + "&metricNamespace=" + metric.Namespace + "&autoadjusttimegrain=true&validatedimensions=false&api-version=2018-01-01",
				ItemType:       "metrics.graph",
				SubscriptionID: currentItem.SubscriptionID,
				Metadata: map[string]string{
					"SuppressSwaggerExpand": "true",
					"SuppressGenericExpand": "true",
					"AggregationType":       strings.ToLower(metric.PrimaryAggregationType),
				},
			})
		}

		return ExpanderResult{
			Response:          data,
			IsPrimaryResponse: true,
			Nodes:             newItems,
			SourceDescription: "MetricsExpander build response metric namespaces",
		}

		// Todo then go and get the metrics and return a list of options to the uers.

	} else if currentItem.ItemType == "metrics.graph" {
		data, err := armclient.DoRequest(ctx, "GET", currentItem.ExpandURL)
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "MetricsExpander request metricDefinitions",
			}
		}

		var metricResponse armclient.MetricResponse
		err = json.Unmarshal([]byte(data), &metricResponse)
		if err != nil {
			panic(err)
		}

		graphData := []float64{}
		for _, datapoint := range metricResponse.Value[0].Timeseries[0].Data {
			value := datapoint[currentItem.Metadata["AggregationType"]].(float64)
			graphData = append(graphData, value)
		}

		graph := asciigraph.Plot(graphData)

		return ExpanderResult{
			Response:          graph,
			IsPrimaryResponse: true,
			SourceDescription: "MetricsExpander build graph",
		}
	} else {

		data, err := armclient.DoRequest(ctx, "GET", currentItem.ID+"/providers/microsoft.insights/metricNamespaces?api-version=2017-12-01-preview")
		if err != nil {
			return ExpanderResult{
				Err:               err,
				SourceDescription: "MetricsExpander request metricNamespaces",
			}
		}

		newItems := []*TreeNode{}

		var metricNamespaceResponse armclient.MetricNamespaceResponse
		err = json.Unmarshal([]byte(data), &metricNamespaceResponse)
		if err != nil {
			panic(err)
		}

		for _, metricNamespace := range metricNamespaceResponse.Value {
			newItems = append(newItems, &TreeNode{
				Name:           metricNamespace.Name,
				Display:        style.Subtle("[Metrics]") + "\n  " + metricNamespace.Name,
				ID:             currentItem.ID,
				Parentid:       currentItem.ID,
				ExpandURL:      currentItem.ID + "/providers/microsoft.insights/metricdefinitions?metricNamespace=" + metricNamespace.Properties.MetricNamespaceName + "&api-version=2018-01-01",
				ItemType:       "metrics.metricdefinition",
				SubscriptionID: currentItem.SubscriptionID,
				Metadata: map[string]string{
					"SuppressSwaggerExpand": "true",
					"SuppressGenericExpand": "true",
				},
			})
		}

		return ExpanderResult{
			IsPrimaryResponse: false,
			Nodes:             newItems,
			SourceDescription: "MetricsExpander build response metric namespaces",
		}
	}
}

/////////////////////////
// Calls

//
// Get metric namespaces relativeUrl: /subscriptions/SUBIDHERE/resourceGroups/lk-scratch/providers/Microsoft.Web/sites/lg-scratch/providers/microsoft.insights/metricNamespaces?api-version=2017-12-01-preview

////////////////////////////

// Docs https://docs.microsoft.com/en-us/rest/api/monitor/metricdefinitions/list
// Get available metrics in namespace relativeUrl: /subscriptions/SUBIDHERE/resourceGroups/lk-scratch/providers/Microsoft.Web/serverFarms/ServicePlan7bdb8347-931a/providers/microsoft.insights/metricdefinitions?metricNamespace=microsoft.web/serverfarms&api-version=2018-01-01

/////////////////////////

// Docs https://docs.microsoft.com/en-us/rest/api/monitor/metrics/list
// Get a metric relativeUrl: /subscriptions/SUBIDHERE/resourceGroups/lk-scratch/providers/Microsoft.Web/sites/lg-scratch/providers/microsoft.Insights/metrics?timespan=2019-09-03T15:25:00.000Z/2019-09-04T15:30:00.000Z&interval=PT5M&metricnames=CpuTime&aggregation=total&metricNamespace=microsoft.web/sites&autoadjusttimegrain=true&validatedimensions=false&api-version=2018-01-01
