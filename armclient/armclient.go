package armclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	userAgentStr = "github.com/lawrencegripper/azbrowse"
)

// func isWriteVerb(verb string) bool {
// 	v := strings.ToUpper(verb)
// 	return v == "PUT" || v == "POST" || v == "PATCH"
// }

var tenantID string

// GetTenantID gets the current tenandid from AzCli
func GetTenantID() string {
	return tenantID
}

// DoRequest makes an ARM rest request
func DoRequest(method, path string) (string, error) {
	url, err := getRequestURL(path)
	if err != nil {
		return "", err
	}

	var reqBody string
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(reqBody)))

	cliToken, err := aquireTokenFromAzCLI()
	if err != nil {
		return "", errors.New("Failed to acquire auth token: " + err.Error())
	}
	tenantID = cliToken.Tenant

	req.Header.Set("Authorization", cliToken.TokenType+" "+cliToken.AccessToken)
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-ms-client-request-id", newUUID())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return "", errors.New("Request failed: " + err.Error())
	}

	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", errors.New("Request failed: " + err.Error())
	}

	return prettyJSON(buf), nil
}
