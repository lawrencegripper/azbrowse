package views

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type response struct {
	httpResponse *http.Response
	body         string
	requestPath  string
}

type pollItem struct {
	pollURI string
	title   string
	status  string
	event   *eventing.StatusEvent
}

// StartWatchingAsyncARMRequests is used to start watching for any async response
// from ARM and trigger notifications to watch their progress
func StartWatchingAsyncARMRequests(ctx context.Context) (armclient.ResponseProcessor, error) {
	requestChan := make(chan response)

	inflightRequests := map[string]pollItem{}

	go func() {
		for {
			defer errorhandling.RecoveryWithCleanup()

			// Stop the routine if the context is canceled
			select {
			case <-ctx.Done():
				// returning not to leak the goroutine
				return
			case request := <-requestChan:
				if !isAsyncResponse(request.httpResponse) {
					// Not an async action, move on.
					continue
				}

				var pollLocation []string
				var exists bool

				pollLocation, exists = request.httpResponse.Header["Azure-AsyncOperation"] //nolint:staticcheck Azure header isn't canonical
				if !exists {
					pollLocation, exists = request.httpResponse.Header["Location"]
					if !exists {
						eventing.SendFailureStatusFromError("Failed to find async poll header", fmt.Errorf("Missing header in %+v", request.httpResponse.Header))
						continue
					}
				}

				parsedURL, err := url.Parse(request.requestPath)
				if err != nil {
					eventing.SendFailureStatusFromError("Failed to parse url while making async request", err)
					continue
				}

				resource := request.requestPath
				pathSegments := strings.Split(parsedURL.Path, "/")
				if len(pathSegments) >= 2 {
					resource = strings.Join(pathSegments[len(pathSegments)-2:], "/")
				}

				item := pollItem{
					pollURI: strings.Join(pollLocation, ""),
					title:   request.httpResponse.Request.Method + " " + resource,
					status:  "unknown",
					event: &eventing.StatusEvent{
						Message:    "Tracing async event to completion",
						Timeout:    time.Minute * 15,
						InProgress: true,
						IsToast:    true,
					},
				}

				eventing.SendStatusEvent(item.event)
				inflightRequests[request.requestPath] = item
			}
		}
	}()

	go func() {
		for {
			defer errorhandling.RecoveryWithCleanup()

			select {
			case <-ctx.Done():
				// returning not to leak the goroutine
				return
			case <-time.After(time.Second * 1):
				// Continue
			}

			for ID, pollItem := range inflightRequests {
				_, done := eventing.SendStatusEvent(&eventing.StatusEvent{
					Message:    "Polling async completion " + pollItem.pollURI,
					Timeout:    time.Second * 5,
					InProgress: true,
				})
				req, err := http.NewRequest("GET", pollItem.pollURI, nil)
				if err != nil {
					eventing.SendFailureStatusFromError("Failed create async poll request", err)
					delete(inflightRequests, ID)
					continue
				}
				response, err := armclient.LegacyInstance.DoRawRequest(ctx, req)
				if err != nil {
					eventing.SendFailureStatusFromError("Failed making async poll request", err)
					delete(inflightRequests, ID)
					continue
				}

				if response.StatusCode == 200 {
					// completed
					pollItem.event.InProgress = false
					pollItem.event.Message = pollItem.title + " COMPLETED"
					pollItem.event.SetTimeout(time.Second * 5)
					eventing.SendStatusEvent(pollItem.event)

					delete(inflightRequests, ID)
				}

				if isAsyncResponse(response) {
					// continue processing
					pollItem.event.Message = pollItem.title
					eventing.SendStatusEvent(pollItem.event)
				}

				// Pause between each poll item so as not to make huge volume of requests.
				<-time.After(time.Second * 2)
				done()
			}
		}
	}()

	return func(requestPath string, httpResponse *http.Response, responseBody string) {
		requestChan <- response{
			requestPath:  requestPath,
			httpResponse: httpResponse,
			body:         responseBody,
		}
	}, nil
}

func isAsyncResponse(response *http.Response) bool {
	return response.StatusCode == 201 || response.StatusCode == 202
}
