package armclient

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// AzCLIToken contains token info from az cli
type AzCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

var currentToken *AzCLIToken

func acquireTokenFromAzCLI(clearCache bool, tenantID string) (AzCLIToken, error) {
	if currentToken == nil || clearCache {
		args := []string{"account", "get-access-token", "--output", "json"}

		if tenantID != "" {
			query := fmt.Sprintf("[?tenantId=='%s'].id| [0] ", tenantID)
			out, err := exec.Command("az", "account", "list", "--output", "tsv", "--query", query).Output()
			if err != nil {
				return AzCLIToken{}, fmt.Errorf("Error looking up subscription from tenant: %s", err)
			}
			subscription := strings.TrimSpace(string(out))
			args = append(args, "--subscription", subscription)
		}

		out, err := exec.Command("az", args...).Output()
		if err != nil {
			return AzCLIToken{}, fmt.Errorf("%s (try running 'az account get-access-token' to get more details)", err)
		}

		var r AzCLIToken
		err = json.Unmarshal(out, &r)
		if err != nil {
			return AzCLIToken{}, err
		}
		currentToken = &r
		return r, nil
	}

	return *currentToken, nil
}

// AcquireTokenForResourceFromAzCLI gets a token for the specified resource endpoint
func AcquireTokenForResourceFromAzCLI(resource string) (AzCLIToken, error) {
	args := []string{"account", "get-access-token", "--output", "json", "--resource", resource}

	out, err := exec.Command("az", args...).Output()
	if err != nil {
		return AzCLIToken{}, fmt.Errorf("%s (try running 'az account get-access-token' to get more details)", err)
	}

	var r AzCLIToken
	err = json.Unmarshal(out, &r)
	if err != nil {
		return AzCLIToken{}, err
	}
	return r, nil
}
