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
}

type pollItem struct {
	requestURI string
	pollURI    string
	title      string
	status     string
	event      eventing.StatusEvent
	retryCount int
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
						panic("202 returned but no location header found")
					}
				}

				url, err := url.Parse(request.httpResponse.Request.RequestURI)
				if err != nil {
					panic(err)
				}

				item := pollItem{
					pollURI:    strings.Join(pollLocation, ""),
					requestURI: request.httpResponse.Request.RequestURI,
					title:      request.httpResponse.Request.Method + " " + url.Path,
					status:     "unknown",
					event: eventing.StatusEvent{
						Message:    "Tracing async event to completion",
						Timeout:    time.Minute * 15,
						InProgress: true,
						IsToast:    true,
					},
				}

				inflightRequests[request.httpResponse.Request.RequestURI] = item
				eventing.SendStatusEvent(item.event)
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
			case <-time.After(time.Second * 15):
				// Update items to see if any have changed
			}

			for _, pollItem := range inflightRequests {
				pollItem.retryCount++

				req, err := http.NewRequest("GET", pollItem.pollURI, nil)
				if err != nil {
					panic(err)
				}
				response, err := armclient.LegacyInstance.DoRawRequest(ctx, req)
				if err != nil {
					panic(err)
				}

				if response.StatusCode == 200 {
					// completed
					pollItem.event.InProgress = false
					eventing.SendStatusEvent(pollItem.event)
				}

				if isAsyncResponse(response) {
					// continue processing
					pollItem.event.Message = pollItem.title + fmt.Sprintf(" checking:%v", pollItem.retryCount)
					eventing.SendStatusEvent(pollItem.event)
				}

				// Get status from the body here...

			}
		}
	}()

	return func(httpResponse *http.Response, responseBody string) {
		requestChan <- response{
			httpResponse: httpResponse,
			body:         responseBody,
		}
	}, nil
}

func isAsyncResponse(response *http.Response) bool {
	return response.StatusCode == 201 || response.StatusCode == 200
}
