package armclient

import (
	"bytes"
	"errors"
	"fmt"
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

	// Check response error but also return body as it may contain useful information
	// about the error
	var responseErr error
	if response.StatusCode < 200 && response.StatusCode > 299 {
		responseErr = fmt.Errorf("Request returned a non-success status code of %v with a status message of %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", errors.New("Request failed: " + err.Error() + " ResponseErr:" + responseErr.Error())
	}

	return prettyJSON(buf), responseErr
}
