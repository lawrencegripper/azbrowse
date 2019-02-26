package armclient

import (
	"context"
	"encoding/json"
	"os/exec"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

type azCLIToken struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	Tenant       string `json:"tenant"`
	Subscription string `json:"subscription"`
	ExpiresOn    string `json:"expiresOn"`
}

var cachedToken azCLIToken

func checkTokenIsValid(token azCLIToken) (azCLIToken, bool) {
	if token.ExpiresOn == "" {
		return azCLIToken{}, false
	}

	return token, true
}

func aquireAccessToken(ctx context.Context) (azCLIToken, error) {

	span, ctx := tracing.StartSpanFromContext(ctx, "aquireAccessToken")
	defer span.Finish()

	token, valid := checkTokenIsValid(cachedToken)
	if valid {
		return token, nil
	}

	token, err := aquireTokenFromAzCLI()
	if err != nil {
		return azCLIToken{}, err
	}

	cachedToken = token
	return token, nil
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
