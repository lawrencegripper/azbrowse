package armclient

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type azCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

func aquireTokenFromAzCLI() (azCLIToken, error) {
	out, err := exec.Command("az", "account", "get-access-token", "--output", "json").CombinedOutput()
	if err != nil {
		return azCLIToken{}, fmt.Errorf("Error authenticating with using the AzCLI: %v %v", string(out), err)
	}

	var r azCLIToken
	err = json.Unmarshal(out, &r)
	if err != nil {
		return azCLIToken{}, err
	}

	return r, nil
}
