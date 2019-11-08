package expanders

import (
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func dummyTokenFunc() func(clearCache bool) (armclient.AzCLIToken, error) {
	return func(clearCache bool) (armclient.AzCLIToken, error) {
		return armclient.AzCLIToken{
			AccessToken:  "bob",
			Subscription: "bill",
			Tenant:       "thing",
			TokenType:    "bearer",
		}, nil
	}
}

type expanderTestCase struct {
	name                string
	statusCode          int
	nodeToExpand        *TreeNode
	urlPath             string
	responseFile        string
	configureGockFunc   *func(t *testing.T)
	treeNodeCheckerFunc func(t *testing.T, r ExpanderResult)
}
