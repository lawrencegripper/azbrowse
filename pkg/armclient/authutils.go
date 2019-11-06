package armclient

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// AzCLIToken contains token info from az cli
type AzCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

var currentToken *AzCLIToken

func aquireTokenFromAzCLI(clearCache bool) (AzCLIToken, error) {
	if currentToken == nil || clearCache {
		out, err := exec.Command("az", "account", "get-access-token", "--output", "json").Output()
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
