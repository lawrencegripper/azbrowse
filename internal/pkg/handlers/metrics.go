package handlers

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// HACK: To draw the graph this handler has to know how big the display area is...
//  to work around this the ItemWidget sets these properties when it's created
//  I don't like it but it works. Will hopefully replace with better mechanism in future
//

// ItemWidgetHeight tracks height of item widget
var ItemWidgetHeight int

// ItemWidgetWidth track width of item widget
var ItemWidgetWidth int

// MetricsExpander expands the data-plane aspects of the Microsoft.Insights RP
type MetricsExpander struct {
}

// Name returns the name of the expander
func (e *MetricsExpander) Name() string {
	return "MetricsExpander"
}

// DoesExpand checks if this is a storage account
func (e *MetricsExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == ResourceType || strings.HasPrefix(currentItem.ItemType, "metrics.") {
		return true, nil
	}

	return false, nil
}

// Expand adds items for metrics to the list
func (e *MetricsExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// We have a metric namespace lets lookup the metric definitions
	if currentItem.ItemType == "metrics.metricdefinition" {
		return expandMetricDefinition(ctx, currentItem)
	}

	// We have a metric definition lets draw the graph
	if currentItem.ItemType == "metrics.graph" {
		return expandGraph(ctx, currentItem)
	}

	// We're looking at a top level resource, lets see if it has a metric namespace
	return expandMetricNamespace(ctx, currentItem)
}

func expandMetricNamespace(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	data, err := armclient.DoRequest(ctx, "GET", currentItem.ID+
		"/providers/microsoft.insights/metricNamespaces?api-version=2017-12-01-preview")
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
		return ExpanderResult{
			Err:               err,
			SourceDescription: "MetricsExpander metricNamespace failed to deserialise",
		}
	}

	for _, metricNamespace := range metricNamespaceResponse.Value {
		newItems = append(newItems, &TreeNode{
			Name:     metricNamespace.Name,
			Display:  style.Subtle("[Metrics]") + "\n  " + metricNamespace.Name,
			ID:       currentItem.ID,
			Parentid: currentItem.ID,
			ExpandURL: currentItem.ID + "/providers/microsoft.insights/metricdefinitions?" +
				"metricNamespace=" + url.QueryEscape(metricNamespace.Properties.MetricNamespaceName) +
				"&api-version=2018-01-01",
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

func expandMetricDefinition(ctx context.Context, currentItem *TreeNode) ExpanderResult {
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
		return ExpanderResult{
			Err:               err,
			SourceDescription: "MetricsExpander metricDefinitions failed to deserialise",
		}
	}

	newItems := []*TreeNode{}

	for _, metric := range metricsListResponse.Value {
		newItems = append(newItems, &TreeNode{
			Name:     metric.Name.Value,
			Display:  metric.Name.Value + "\n  " + style.Subtle("Unit: "+metric.Unit),
			ID:       currentItem.ID,
			Parentid: currentItem.ID,
			ExpandURL: currentItem.ID + "/providers/microsoft.Insights/metrics?timespan=" +
				time.Now().UTC().Add(-4*time.Hour).Format("2006-01-02T15:04:05.000Z") + "/" +
				time.Now().UTC().Format("2006-01-02T15:04:05.000Z") + "&interval=PT1M&metricnames=" +
				url.QueryEscape(metric.Name.Value) + "&aggregation=" +
				url.QueryEscape(metric.PrimaryAggregationType) +
				"&metricNamespace=" + url.QueryEscape(metric.Namespace) +
				"&autoadjusttimegrain=true&validatedimensions=false&api-version=2018-01-01",
			ItemType:       "metrics.graph",
			SubscriptionID: currentItem.SubscriptionID,
			Metadata: map[string]string{
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
				"AggregationType":       strings.ToLower(metric.PrimaryAggregationType),
				"Units":                 strings.ToLower(metric.Unit),
			},
		})
	}

	return ExpanderResult{
		Response:          data,
		IsPrimaryResponse: true,
		Nodes:             newItems,
		SourceDescription: "MetricsExpander build response metric namespaces",
	}
}

func expandGraph(ctx context.Context, currentItem *TreeNode) ExpanderResult {
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
		return ExpanderResult{
			Err:               err,
			SourceDescription: "MetricsExpander graphdata failed to deserialise",
		}
	}

	caption := style.Title(currentItem.Name) +
		style.Subtle(" (Aggregate: '"+currentItem.Metadata["AggregationType"]+"' Unit: '"+
			currentItem.Metadata["Units"]+"')")

	graphData := []float64{}
	for _, datapoint := range metricResponse.Value[0].Timeseries[0].Data {
		value, success := datapoint[currentItem.Metadata["AggregationType"]].(float64)
		if success {
			graphData = append(graphData, value)
		} else {
			graphData = append(graphData, float64(0))
		}
	}

	graph := asciigraph.Plot(graphData,
		asciigraph.Height(ItemWidgetHeight-6),
		asciigraph.Width(ItemWidgetWidth-15),
		asciigraph.Caption("time: 4hrs ago ----> now"))

	return ExpanderResult{
		Response:          "\n\n" + caption + "\n\n" + style.Graph(graph),
		IsPrimaryResponse: true,
		SourceDescription: "MetricsExpander build graph",
	}
}
