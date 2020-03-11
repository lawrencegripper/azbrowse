package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func DummyTokenFunc() func(clearCache bool) (armclient.AzCLIToken, error) {
	return func(clearCache bool) (armclient.AzCLIToken, error) {
		return armclient.AzCLIToken{
			AccessToken:  "bob",
			Subscription: "bill",
			Tenant:       "thing",
			TokenType:    "bearer",
		}, nil
	}
}
