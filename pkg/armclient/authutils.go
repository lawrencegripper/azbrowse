package armclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// AzCLIToken contains token info from az cli
type AzCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

var currentToken *AzCLIToken

func acquireToken(scope string, tenantID string, subscriptionID string) (AzCLIToken, error) {
	if scope == "" {
		return AzCLIToken{}, fmt.Errorf("no scope specified")
	}
	var credOptions *azidentity.DefaultAzureCredentialOptions
	if tenantID != "" {
		credOptions = &azidentity.DefaultAzureCredentialOptions{
			TenantID: tenantID,
		}
	}
	cred, err := azidentity.NewDefaultAzureCredential(credOptions)
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("failed to get credential: %v", err)
	}

	subscriptionClient, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("failed to get subscription client: %v", err)
	}

	ctx := context.Background()

	if subscriptionID == "" {
		subscriptionPager := subscriptionClient.NewListPager(nil)
		for subscriptionPager.More() {
			subscriptionList, err := subscriptionPager.NextPage(ctx)
			if err != nil {
				return AzCLIToken{}, fmt.Errorf("failed to get subscription list: %v", err)
			}
			if subscriptionList.Value != nil && len(subscriptionList.Value) > 0 {
				subscriptionID = *subscriptionList.Value[0].SubscriptionID
				break
			}
		}
	}

	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{
			scope,
		},
	})
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("failed to get access token: %v", err)
	}

	return AzCLIToken{
		AccessToken:  token.Token,
		TokenType:    "Bearer",
		Tenant:       tenantID,
		Subscription: subscriptionID,
	}, nil
}

func acquireTokenFromAzCLI(clearCache bool, tenantID string) (AzCLIToken, error) {
	if currentToken == nil || clearCache {
		t, err := acquireToken("https://management.azure.com/.default", tenantID, "")
		if err != nil {
			return AzCLIToken{}, fmt.Errorf("failed to get token: %v", err)
		}
		currentToken = &t
	}
	return *currentToken, nil
}

// AcquireTokenForResourceFromAzCLI gets a token for the specified resource endpoint
func AcquireTokenForResourceFromAzCLI(subscription string, resource string) (AzCLIToken, error) {
	if !strings.HasSuffix(resource, "/.default") {
		resource += "/.default"
	}
	t, err := acquireToken(resource, "", subscription)
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("failed to get token: %v", err)
	}
	return t, nil
}

// AcquireTokenForGraphFromAzCLI gets a token for MSGraph
func AcquireTokenForGraphFromAzCLI() (AzCLIToken, error) {
	t, err := acquireToken("https://graph.microsoft.com/.default", "", "")
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("failed to get token: %v", err)
	}
	return t, nil
}
