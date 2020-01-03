package armclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

const (
	userAgentStr     = "github.com/lawrencegripper/azbrowse"
	providerCacheKey = "providerCache"
)

// TokenFunc is the interface to meet for functions which retrieve tokens for the ARMClient
type TokenFunc func(clearCache bool) (AzCLIToken, error)

// ResponseProcessor can be used to handle additional actions once a response is received
type ResponseProcessor func(response *http.Response, responseBody string)

// Client is used to talk to the ARM API's in Azure
type Client struct {
	client             *http.Client
	tenantID           string
	responseProcessors []ResponseProcessor

	acquireToken TokenFunc
}

// LegacyInstance is a singleton ARMClient used while migrating to the
// injected client
var LegacyInstance Client

// NewClientFromCLI creates a new client using the auth details on disk used by the azurecli
func NewClientFromCLI(tenantID string, responseProcessors ...ResponseProcessor) *Client {
	aquireToken := func(clearCache bool) (AzCLIToken, error) {
		return aquireTokenFromAzCLI(clearCache, tenantID)
	}
	return &Client{
		responseProcessors: responseProcessors,
		acquireToken:       aquireToken,
		client:             &http.Client{},
	}
}

// NewClientFromClientAndTokenFunc create a client for testing using custom token func and httpclient
func NewClientFromClientAndTokenFunc(client *http.Client, tokenFunc TokenFunc) *Client {
	return &Client{
		acquireToken: tokenFunc,
		client:       client,
	}
}

// SetClient is used to override the HTTP Client used.
// This is useful when testing
func (c *Client) SetClient(newClient *http.Client) {
	c.client = newClient
}

// SetAquireToken lets you override the token func for testing
// or other purposes
func (c *Client) SetAquireToken(aquireFunc func(clearCache bool) (AzCLIToken, error)) {
	c.acquireToken = aquireFunc
}

// GetTenantID gets the current tenandid from AzCli
func (c *Client) GetTenantID() string {
	return c.tenantID
}

// GetToken gets the cached cli token
func (c *Client) GetToken() (AzCLIToken, error) {
	return c.acquireToken(false)
}

// RequestResult used with async channel
type RequestResult struct {
	Result string
	Error  error
}

// DoRequestAsync makes an ARM rest request
func (c *Client) DoRequestAsync(ctx context.Context, method, path string) chan RequestResult {
	requestResultChan := make(chan RequestResult)
	go func() {
		// recover from panic, if one occurrs, and leave terminal usable
		defer errorhandling.RecoveryWithCleanup()

		data, err := c.DoRequestWithBody(ctx, method, path, "")
		requestResultChan <- RequestResult{
			Error:  err,
			Result: data,
		}
	}()
	return requestResultChan
}

// DoRequest makes an ARM rest request
func (c *Client) DoRequest(ctx context.Context, method, path string) (string, error) {
	return c.DoRequestWithBody(ctx, method, path, "")
}

// DoRawRequest makes a raw request with ARM authentication headers set
func (c *Client) DoRawRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	cliToken, err := c.acquireToken(false)
	if err != nil {
		return nil, errors.New("Failed to acquire auth token: " + err.Error())
	}
	c.tenantID = cliToken.Tenant

	req.Header.Set("Authorization", cliToken.TokenType+" "+cliToken.AccessToken)
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-ms-client-request-id", newUUID())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req.WithContext(ctx))
}

// DoRequestWithBody makes an ARM rest request
func (c *Client) DoRequestWithBody(ctx context.Context, method, path, body string) (string, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "request:"+method, tracing.SetTag("path", path))
	defer span.Finish()

	url, err := getRequestURL(path)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return "", errors.New("Failed to create request for body: " + err.Error())
	}

	response, err := c.DoRawRequest(ctx, req)

	if response != nil && response.StatusCode == 401 {
		// This might be because the token we've cached has expired.
		// Get a new token forcing it to clear cache
		cliToken, err := c.acquireToken(true)
		if err != nil {
			return "", errors.New("Failed to acquire auth token: " + err.Error())
		}
		c.tenantID = cliToken.Tenant

		// Retry the request now we have a valid token
		response, err = c.client.Do(req.WithContext(ctx)) //nolint:staticcheck
	}
	if err != nil {
		return "", errors.New("Request failed: " + err.Error())
	}

	// Check response error but also return body as it may contain useful information
	// about the error
	var responseErr error
	if response.StatusCode < 200 || response.StatusCode > 299 {
		span.SetTag("isError", true)
		span.SetTag("errorCode", response.StatusCode)
		span.SetTag("error", response.Status)

		responseErr = fmt.Errorf("Request returned a non-success status code of %v with a status message of %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)

	// Call the response Processors
	for _, responseProcessor := range c.responseProcessors {
		responseProcessor(response, string(buf))
	}

	if err != nil {
		wrappedError := errors.New("Request failed: " + err.Error() + " ResponseErr:" + responseErr.Error())
		span.SetTag("err", wrappedError)
		return "", wrappedError
	}

	if tracing.IsDebug() {
		span.SetTag("responseBody", truncateString(string(buf), 1500))
		span.SetTag("requestBody", body)
		span.SetTag("url", url)
	}

	return string(buf), responseErr
}

// DoResourceGraphQuery performs an azure graph query
func (c *Client) DoResourceGraphQuery(ctx context.Context, subscription, query string) (string, error) {
	messageBody := `{"subscriptions": ["SUB_HERE"], "query": "QUERY_HERE", "options": {"$top": 1000, "$skip": 0}}`
	messageBody = strings.Replace(messageBody, "SUB_HERE", subscription, -1)
	messageBody = strings.Replace(messageBody, "QUERY_HERE", query, -1)
	tracing.SetTagOnCtx(ctx, "query", messageBody)
	return c.DoRequestWithBody(ctx, "POST", "/providers/Microsoft.ResourceGraph/resources?api-version=2018-09-01-preview", messageBody)
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
func (c *Client) PopulateResourceAPILookup(ctx context.Context) {
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
			data, err := c.DoRequest(ctx, "GET", "/providers?api-version=2017-05-10")
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

			storage.PutCache(providerCacheKey, providerData) //nolint: errcheck

		} else {
			span.SetTag("Data read from cache", true)
			var providerCache map[string]string
			err = json.Unmarshal([]byte(providerData), &providerCache)
			if err != nil {
				span.SetTag("error: failed to read data from cache", err)
				span.Finish()
				panic(err)
			}
			resourceAPIVersionLookup = providerCache
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
