package armclient

import (
	"encoding/json"
	"os/exec"
)

const (
	msiEndpoint             = "http://localhost:50342/oauth2/token"
	activeDirectoryEndpoint = "https://login.microsoftonline.com/"
	armResource             = "https://management.core.windows.net/"
	clientAppID             = "04b07795-8ddb-461a-bbee-02f9e1bf7b46"
	commonTenant            = "common"
)

type azCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
}

func aquireTokenFromAzCLI() (azCLIToken, error) {
	out, err := exec.Command("az", "account", "get-access-token").Output()
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
