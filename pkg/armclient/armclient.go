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
	"time"

	"golang.org/x/time/rate"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
)

const (
	userAgentStr     = "github.com/lawrencegripper/azbrowse"
	providerCacheKey = "providerCache"
)

// TokenFunc is the interface to meet for functions which retrieve tokens for the ARMClient
type TokenFunc func(clearCache bool) (AzCLIToken, error)

// ResponseProcessor can be used to handle additional actions once a response is received
type ResponseProcessor func(requestPath string, response *http.Response, responseBody string)

// Client is used to talk to the ARM API's in Azure
type Client struct {
	client             *http.Client
	tenantID           string
	responseProcessors []ResponseProcessor
	limiter            *rate.Limiter

	acquireToken TokenFunc
}

// LegacyInstance is a singleton ARMClient used while migrating to the
// injected client
var LegacyInstance *Client

const requestPerSecLimit = 10
const requestPerSecBurst = 5

// NewClientFromCLI creates a new client using the auth details on disk used by the azurecli
func NewClientFromCLI(tenantID string, responseProcessors ...ResponseProcessor) *Client {
	aquireToken := func(clearCache bool) (AzCLIToken, error) {
		return acquireTokenFromAzCLI(clearCache, tenantID)
	}
	return &Client{
		responseProcessors: responseProcessors,
		limiter:            rate.NewLimiter(requestPerSecLimit, requestPerSecBurst),
		acquireToken:       aquireToken,
		client:             &http.Client{},
	}
}

// NewClientFromConfig create a client for testing using custom token func and httpclient
func NewClientFromConfig(client *http.Client, tokenFunc TokenFunc, reqPerSecLimit float64, responseProcessors ...ResponseProcessor) *Client {
	return &Client{
		responseProcessors: responseProcessors,
		acquireToken:       tokenFunc,
		limiter:            rate.NewLimiter(rate.Limit(reqPerSecLimit), 10), // Keep the rate limitter but set high values for tests to complete quickly
		client:             client,
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

	var span opentracing.Span
	reservation := c.limiter.Reserve()
	if !reservation.OK() {
		panic("Ratelimitter prevented request which should never happen.")
	}
	delay := reservation.Delay()
	if delay.Seconds() > 0 {
		span, _ = tracing.StartSpanFromContext(ctx, "ratelimitted")
		eventing.SendFailureStatus("Request rate limitted due to high call volume")
	}
	time.Sleep(reservation.Delay())
	if span != nil {
		span.Finish()
	}
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
		responseProcessor(path, response, string(buf))
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
	queryBody := QueryBody{
		Subscriptions: []string{subscription},
		Query:         query,
		Options: QueryOptions{
			Top:  1000,
			Skip: 0,
		},
	}
	messageBody, err := json.Marshal(queryBody) //nolint: errcheck
	if err != nil {
		return "", err
	}
	tracing.SetTagOnCtx(ctx, "query", messageBody)
	return c.DoRequestWithBody(ctx, "POST", "/providers/Microsoft.ResourceGraph/resources?api-version=2018-09-01-preview", string(messageBody))
}

// DoResourceGraphQueryReturningObjectArray performs an azure graph query on all subs you have access too
func (c *Client) DoResourceGraphQueryReturningObjectArray(ctx context.Context, subscriptionGUIDs []string, query string) (string, error) {
	resultFmt := "objectArray"
	queryBody := QueryBody{
		Subscriptions: subscriptionGUIDs,
		Query:         query,
		Options: QueryOptions{
			Top:          1000,
			Skip:         0,
			Resultformat: resultFmt,
		},
	}
	messageBody, err := json.Marshal(queryBody) //nolint: errcheck
	if err != nil {
		return "", err
	}

	tracing.SetTagOnCtx(ctx, "query", messageBody)
	return c.DoRequestWithBody(ctx, "POST", "/providers/Microsoft.ResourceGraph/resources?api-version=2018-09-01-preview", string(messageBody))
}

var resourceAPIVersionLookup map[string]string
var resourceAPIVersionPreviewLookup map[string]string

// GetAPIVersion returns the most recent API version for a resource
func GetAPIVersion(armType string) (string, error) {
	armTypeKey := strings.ToLower(armType)
	value, exists := resourceAPIVersionLookup[armTypeKey]
	if exists {
		return value, nil
	}
	value, exists = resourceAPIVersionPreviewLookup[armTypeKey]
	if exists {
		return value, nil
	}
	return "MISSING", fmt.Errorf("API not found for the resource: %s", armType)
}

// PopulateResourceAPILookup is used to build a cache of resourcetypes -> api versions
// this is needed when requesting details from a resource as APIVersion isn't known and is required
func (c *Client) PopulateResourceAPILookup(ctx context.Context, msg *eventing.StatusEvent) {
	// w.statusView.Status("Getting provider data from cache", true)
	if resourceAPIVersionLookup == nil {
		span, ctx := tracing.StartSpanFromContext(ctx, "populateResCache")
		// Get data from cache
		valid, providerData, err := storage.GetCacheWithTTL(providerCacheKey, time.Hour*24)

		// w.statusView.Status("Getting provider data from cache: Completed", false)

		if err != nil || providerData == "" || !valid {
			msg.Message = "Getting provider data from Azure API"
			msg.Update()

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
			resourceAPIVersionPreviewLookup = make(map[string]string)
			for _, provider := range providerResponse.Providers {
				for _, resourceType := range provider.ResourceTypes {
					for _, apiVersion := range resourceType.APIVersions {
						armType := provider.Namespace + "/" + resourceType.ResourceType
						armTypeKey := strings.ToLower(armType)
						if strings.Contains(apiVersion, "preview") {
							// don't break here as we want to allow non-preview version to be set, instead check to avoid overwriting
							if resourceAPIVersionPreviewLookup[armTypeKey] == "" {
								resourceAPIVersionPreviewLookup[armTypeKey] = apiVersion
							}
						} else {
							resourceAPIVersionLookup[armTypeKey] = apiVersion
							break
						}
					}
				}
			}

			bytes, err := json.Marshal(resourceAPIVersionLookup)
			if err != nil {
				panic(err)
			}
			providerData = string(bytes)

			err = storage.PutCacheForTTL(providerCacheKey, providerData)
			if err != nil {
				msg.Failure = true
				msg.Message = "Failed to save provider data to cache"
				msg.Update()
			}

		} else {
			msg.Message = "Got provider data from cache"
			msg.Update()
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

// QueryBody is the structured body require for an Azure Resource Graph submission
type QueryBody struct {
	Subscriptions []string     `json:"subscriptions"`
	Query         string       `json:"query"`
	Options       QueryOptions `json:"options"`
}

// QueryOptions is for use with QueryBocy to set max returned items and structure
type QueryOptions struct {
	Top          int    `json:"$top"`
	Skip         int    `json:"$skip"`
	Resultformat string `json:"resultFormat,omitempty"`
}
