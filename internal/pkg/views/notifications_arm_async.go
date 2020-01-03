package views

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
	event      *eventing.StatusEvent
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
						log.Panicf("Failed to find header on response: %+v", request.httpResponse)
					}
				}

				item := pollItem{
					pollURI:    strings.Join(pollLocation, ""),
					requestURI: request.httpResponse.Request.RequestURI,
					title:      request.httpResponse.Request.Method + " " + request.httpResponse.Request.RequestURI,
					status:     "unknown",
					event: &eventing.StatusEvent{
						Message:    "Tracing async event to completion",
						Timeout:    time.Minute * 15,
						InProgress: true,
						IsToast:    true,
					},
				}

				eventing.SendStatusEvent(item.event)
				inflightRequests[request.httpResponse.Request.RequestURI] = item
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
			case <-time.After(time.Second * 5):
				// Update items to see if any have changed
			}

			for ID, pollItem := range inflightRequests {
				pollItem.retryCount = pollItem.retryCount + 1

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
					pollItem.event.Message = pollItem.title + "COMPLETED"
					eventing.SendStatusEvent(pollItem.event)
					delete(inflightRequests, ID)
				}

				if isAsyncResponse(response) {
					// continue processing
					pollItem.event.Message = pollItem.title + fmt.Sprintf(" poll#:%v", pollItem.retryCount)
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
	return response.StatusCode == 201 || response.StatusCode == 202
}
