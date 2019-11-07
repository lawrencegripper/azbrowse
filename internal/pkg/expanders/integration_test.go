package expanders

import "testing"

type expanderTestCase struct {
	name string
	expander Expander
	nodeToExpand *TreeNode
	urlPath  string
	responseFile string
	treeNodeCheckerFunc func(t *testing.T, r ExpanderResult)
}

var expanderResponseTests = []expanderTestCase{
	{
		name: "ExpandSubscription->ResourceGroups",
		expander: SubscriptionExpander{},
		nodeToExpand: &TreeNode{
			Display:        "Thingy1",
			Name:           "Thingy1",
			ID:             "/subscriptions/00000000-0000-0000-0000-000000000000",
			ExpandURL:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01",
			ItemType:       SubscriptionType,
			SubscriptionID: "00000000-0000-0000-0000-000000000000",
		},
		urlPath: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups"
		responseFile: "./testdata/armsamples/resourcegroups/response.json"
		treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
			st.Expect(t, result.Err, nil)
			st.Expect(t, len(result.Nodes), 6)
		
			// Validate content
			st.Expect(t, result.Nodes[0].Name, "cloudshell")
			st.Expect(t, result.Nodes[0].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cloudshell/resources?api-version=2017-05-10")			
		}
	}
}

func TestFlagParser(t *testing.T) {
	
	
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
