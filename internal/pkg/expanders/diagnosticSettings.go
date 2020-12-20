package expanders

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/nbio/st"
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
			return ExpanderResult{
				Err:               fmt.Errorf("Error - Failed retrieving diagnostic settings"),
				Response:          ExpanderResponse{Response: result},
				SourceDescription: "DiagnosticSettingsExpander request",
			}
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

		for _, diagSetting := range itemArray {
			itemId := string(diagSetting.GetStringBytes("id"))
			if itemId == "" {
				continue
			}
			expandUrl := itemId + "?api-version=2017-05-01-preview"
			name := string(diagSetting.GetStringBytes("name"))
			if name == "" {
				continue
			}
			diagnosticSettingsItems = append(diagnosticSettingsItems, &TreeNode{
				Name:      name,
				Display:   style.Subtle("[microsoft.insights] \n  ") + name,
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
	return true, &[]expanderTestCase{
		{
			name:         "diagnosticSettingsFound",
			responseFile: "./testdata/armsamples/diagSettings/responseNormal.json",
			statusCode:   200,
			urlPath:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/stable/providers/microsoft.containerregistry/registries/aregistry/providers/microsoft.insights/diagnosticSettings",
			nodeToExpand: &TreeNode{
				Parentid:  "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/bob",
				Namespace: "None",
				Display:   style.Subtle("[Microsoft.Insights]") + "\n  Diagnostic Settings",
				Name:      "Diagnostic Settings",
				ID:        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/bob/<diagsettings>",
				ExpandURL: ExpandURLNotSupported,
				ItemType:  diagnosticSettingsType,
				Metadata: map[string]string{
					// Diagnostic settings hang off resources so a list is passed for the
					// expander to use
					resourceIdsMeta: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/stable/providers/microsoft.containerregistry/registries/aregistry",
				},
			},
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)

				st.Expect(t, len(r.Nodes), 1)

				// Validate content
				st.Expect(t, r.Nodes[0].Name, "d1")
			},
		},
		{
			name:         "diagnosticSettingsMissing",
			responseFile: "./testdata/armsamples/diagSettings/responseMalformed.json",
			statusCode:   200,
			urlPath:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/stable/providers/microsoft.containerregistry/registries/aregistry/providers/microsoft.insights/diagnosticSettings",
			nodeToExpand: &TreeNode{
				Parentid:  "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/bob",
				Namespace: "None",
				Display:   style.Subtle("[Microsoft.Insights]") + "\n  Diagnostic Settings",
				Name:      "Diagnostic Settings",
				ID:        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/bob/<diagsettings>",
				ExpandURL: ExpandURLNotSupported,
				ItemType:  diagnosticSettingsType,
				Metadata: map[string]string{
					// Diagnostic settings hang off resources so a list is passed for the
					// expander to use
					resourceIdsMeta: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/stable/providers/microsoft.containerregistry/registries/aregistry",
				},
			},
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)

				st.Expect(t, len(r.Nodes), 0)
			},
		},
	}
}
