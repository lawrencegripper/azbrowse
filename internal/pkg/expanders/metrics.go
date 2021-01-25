package expanders

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

// Check interface
var _ Expander = &MetricsExpander{}

// MetricsExpander expands the data-plane aspects of the Microsoft.Insights RP
type MetricsExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *MetricsExpander) setClient(c *armclient.Client) {
	e.client = c
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
		return e.expandMetricDefinition(ctx, currentItem)
	}

	// We have a metric definition lets draw the graph
	if currentItem.ItemType == "metrics.graph" {
		return e.expandGraph(ctx, currentItem)
	}

	// We're looking at a top level resource, lets see if it has a metric namespace
	return e.expandMetricNamespace(ctx, currentItem)
}

func (e *MetricsExpander) expandMetricNamespace(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	data, err := e.client.DoRequest(ctx, "GET", currentItem.ID+
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
			ID:       currentItem.ID + "/providers/microsoft.insights/metricdefinitions",
			Parentid: currentItem.ID,
			ExpandURL: currentItem.ID + "/providers/microsoft.insights/metricdefinitions?" +
				"metricNamespace=" + url.QueryEscape(metricNamespace.Properties.MetricNamespaceName) +
				"&api-version=2018-01-01",
			ItemType:              "metrics.metricdefinition",
			SubscriptionID:        currentItem.SubscriptionID,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"ResourceID": currentItem.ID,
			},
		})
	}

	return ExpanderResult{
		IsPrimaryResponse: false,
		Nodes:             newItems,
		SourceDescription: "MetricsExpander build response metric namespaces",
	}
}

func (e *MetricsExpander) expandMetricDefinition(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	data, err := e.client.DoRequest(ctx, "GET", currentItem.ExpandURL)
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
			ID:       currentItem.Metadata["ResourceID"] + "/providers/microsoft.Insights/metrics",
			Parentid: currentItem.ID,
			ExpandURL: currentItem.Metadata["ResourceID"] + "/providers/microsoft.Insights/metrics?timespan=" +
				time.Now().UTC().Add(-4*time.Hour).Format("2006-01-02T15:04:05.000Z") + "/" +
				time.Now().UTC().Format("2006-01-02T15:04:05.000Z") + "&interval=PT1M&metricnames=" +
				url.QueryEscape(metric.Name.Value) + "&aggregation=" +
				url.QueryEscape(metric.PrimaryAggregationType) +
				"&metricNamespace=" + url.QueryEscape(metric.Namespace) +
				"&autoadjusttimegrain=true&validatedimensions=false&api-version=2018-01-01",
			ItemType:              "metrics.graph",
			SubscriptionID:        currentItem.SubscriptionID,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"AggregationType": strings.ToLower(metric.PrimaryAggregationType),
				"Units":           strings.ToLower(metric.Unit),
			},
		})
	}

	return ExpanderResult{
		Response:          ExpanderResponse{Response: data},
		IsPrimaryResponse: true,
		Nodes:             newItems,
		SourceDescription: "MetricsExpander build response metric namespaces",
	}
}

func (e *MetricsExpander) expandGraph(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	data, err := e.client.DoRequest(ctx, "GET", currentItem.ExpandURL)
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

	// handle empty response
	if len(metricResponse.Value) < 1 || len(metricResponse.Value[0].Timeseries) < 1 {
		return ExpanderResult{
			Err:               err,
			SourceDescription: "MetricsExpander graphdata failed to deserialise",
		}
	}

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
		Response:          ExpanderResponse{Response: "\n\n" + caption + "\n\n" + style.Graph(graph)},
		IsPrimaryResponse: true,
		SourceDescription: "MetricsExpander build graph",
	}
}
