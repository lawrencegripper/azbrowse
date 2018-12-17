package armclient

import (
	"encoding/json"
	"os/exec"
)

type azCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

func aquireTokenFromAzCLI() (azCLIToken, error) {
	out, err := exec.Command("az", "account", "get-access-token", "--output", "json").Output()
	if err != nil {
		return azCLIToken{}, err
	}

	var r azCLIToken
	err = json.Unmarshal(out, &r)
	if err != nil {
		return azCLIToken{}, err
	}

	return r, nil
}
