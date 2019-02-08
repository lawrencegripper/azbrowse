package armclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lawrencegripper/azbrowse/storage"
	"github.com/lawrencegripper/azbrowse/tracing"
)

const (
	userAgentStr     = "github.com/lawrencegripper/azbrowse"
	providerCacheKey = "providerCache"
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
func DoRequest(ctx context.Context, method, path string) (string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "request:"+method, tracing.SetTag("path", path))
	defer span.Finish()

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
		span.SetTag("isError", true)
		span.SetTag("errorCode", response.StatusCode)
		span.SetTag("error", response.Status)

		responseErr = fmt.Errorf("Request returned a non-success status code of %v with a status message of %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)

	if err != nil {
		wrappedError := errors.New("Request failed: " + err.Error() + " ResponseErr:" + responseErr.Error())
		span.SetTag("err", wrappedError)
		return "", wrappedError
	}

	prettyOutput := prettyJSON(buf)
	if tracing.IsDebug() {
		span.SetTag("responseBody", truncateString(prettyOutput, 800))
		span.SetTag("requestBody", reqBody)
		span.SetTag("url", url)
	}

	return prettyOutput, responseErr
}

var resourceAPIVersionLookup map[string]string

// GetAPIVersion returns the most recent API version for a resource
func GetAPIVersion(armType string) (string, error) {
	value, exists := resourceAPIVersionLookup[armType]
	if !exists {
		return "MISSING", fmt.Errorf("API not found for the resource: %s", armType)
	}
	return value, nil
}

// PopulateResourceAPILookup is used to build a cache of resourcetypes -> api versions
// this is needed when requesting details from a resource as APIVersion isn't known and is required
func PopulateResourceAPILookup(ctx context.Context) {
	// w.statusView.Status("Getting provider data from cache", true)
	if resourceAPIVersionLookup == nil {
		span, ctx := tracing.StartSpanFromContext(ctx, "populateResCache")
		// Get data from cache
		providerData, err := storage.GetCache(providerCacheKey)

		// w.statusView.Status("Getting provider data from cache: Completed", false)

		if err != nil || providerData == "" {
			// w.statusView.Status("Getting provider data from API", true)
			span.SetTag("error: failed getting cached data", err)
			span.SetTag("cacheData", providerData)

			// Get Subscriptions
			data, err := DoRequest(ctx, "GET", "/providers?api-version=2017-05-10")
			if err != nil {
				panic(err)
			}
			var providerResponse ProvidersResponse
			err = json.Unmarshal([]byte(data), &providerResponse)
			if err != nil {
				panic(err)
			}

			resourceAPIVersionLookup = make(map[string]string)
			for _, provider := range providerResponse.Providers {
				for _, resourceType := range provider.ResourceTypes {
					resourceAPIVersionLookup[provider.Namespace+"/"+resourceType.ResourceType] = resourceType.APIVersions[0]
				}
			}

			bytes, err := json.Marshal(resourceAPIVersionLookup)
			if err != nil {
				panic(err)
			}
			providerData = string(bytes)

			storage.PutCache(providerCacheKey, providerData)
			// w.statusView.Status("Getting provider data from API: Completed", false)

		} else {
			span.LogEvent("Data read from cache")
			var providerCache map[string]string
			err = json.Unmarshal([]byte(providerData), &providerCache)
			if err != nil {
				span.SetTag("error: failed to read data from cache", err)
				span.Finish()
				panic(err)
			}
			resourceAPIVersionLookup = providerCache
			// w.statusView.Status("Getting provider data from cache: Completed", false)

		}
		span.Finish()

	}
}

func truncateString(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i]) + "..."
	}
	return s
}
