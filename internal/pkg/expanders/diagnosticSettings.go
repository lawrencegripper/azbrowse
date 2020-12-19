package expanders

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Check interface
var _ Expander = &DiagnosticSettingsExpander{}

// DiagnosticSettingsExpander expands diagnostic settings under an RG
type DiagnosticSettingsExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *DiagnosticSettingsExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *DiagnosticSettingsExpander) Name() string {
	return "ResourceGroupExpander"
}

// DoesExpand checks if this is an RG
func (e *DiagnosticSettingsExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == diagnosticSettingsType {
		return true, nil
	}

	return false, nil
}

// Expand returns Resources in the RG
func (e *DiagnosticSettingsExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	resourceIds := strings.Split(currentItem.Metadata[resourceIdsMeta], ",")

	diagnosticSettingsItems := []*TreeNode{}

	for _, resourceId := range resourceIds {
		diagSettingListUrl := resourceId + "/providers/microsoft.insights/diagnosticSettings?api-version=2017-05-01-preview"
		result, err := e.client.DoRequest(ctx, "GET", diagSettingListUrl)
		if err != nil {
			// Expected some things won't have items
			fmt.Printf("Missing diagnostic setting err %q", err)
		}

		json, err := fastJSONParser.Parse(result)
		if err != nil {
			return ExpanderResult{
				Err:               fmt.Errorf("Error - Failed to parse JSON response"),
				Response:          ExpanderResponse{Response: result},
				SourceDescription: "DiagnosticSettingsExpander request",
			}
		}

		itemArray := json.GetArray("value")
		if itemArray == nil {
			return ExpanderResult{
				Err:               fmt.Errorf("Error - JSON response not array"),
				Response:          ExpanderResponse{Response: result},
				SourceDescription: "DiagnosticSettingsExpander request",
			}
		}

		for _, diagSetting := range itemArray {
			expandUrl := strings.Replace(diagSetting.Get("id").String(), "\"", "", -1) + "?api-version=2017-05-01-preview"
			diagnosticSettingsItems = append(diagnosticSettingsItems, &TreeNode{
				Name:      "diagSetting",
				Display:   style.Subtle("[microsoft.insights] \n  ") + diagSetting.Get("name").String(),
				ExpandURL: expandUrl,
				DeleteURL: expandUrl,
			})
		}
	}

	return ExpanderResult{
		IsPrimaryResponse: true,
		Nodes:             diagnosticSettingsItems,
		SourceDescription: "DiagnosticSettingsExpander request",
	}
}

func (e *DiagnosticSettingsExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}
