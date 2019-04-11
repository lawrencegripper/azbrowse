package armclient

import (
	"encoding/json"
	"os/exec"
)

// AzCLIToken contains token info from az cli
type AzCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

func aquireTokenFromAzCLI() (AzCLIToken, error) {
	out, err := exec.Command("az", "account", "get-access-token", "--output", "json").Output()
	if err != nil {
		return AzCLIToken{}, err
	}

	var r AzCLIToken
	err = json.Unmarshal(out, &r)
	if err != nil {
		return AzCLIToken{}, err
	}

	return r, nil
}
