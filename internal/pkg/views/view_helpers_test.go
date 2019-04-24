package views

import (
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func dummyTokenFunc() func(clearCache bool) (armclient.AzCLIToken, error) {
	return func(clearCache bool) (armclient.AzCLIToken, error) {
		return armclient.AzCLIToken{
			AccessToken:  "bob",
			Subscription: "bill",
			Tenant:       "thing",
			TokenType:    "bearer",
		}, nil
	}
}

func WaitForCompletedStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int) eventing.StatusEvent {
	return WaitForStatusEvent(t, statusEvents, waitForSec, false)
}

func WaitForFailureStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int) eventing.StatusEvent {
	return WaitForStatusEvent(t, statusEvents, waitForSec, true)
}

func WaitForStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int, expectError bool) eventing.StatusEvent {
	for index := 0; index < waitForSec; index++ {
		select {
		case statusRaw := <-statusEvents:
			statusEvent := statusRaw.(eventing.StatusEvent)
			t.Logf("EVENT STATUS MESSAGE: %s Failure: %v InProgress: %v", statusEvent.Message, statusEvent.Failure, statusEvent.InProgress)
			// Wait for things to finish
			if statusEvent.Failure && !expectError {
				t.Error(statusEvent.Message)
				t.FailNow()
			}
			if expectError && statusEvent.Failure {
				return statusEvent
			}
			if statusEvent.InProgress == false {
				return statusEvent
			}
		case <-time.After(1 * time.Second):
			t.Log("Waited 1 sec...")
		}
	}

	t.Error("Waited for event which never occurred")
	return eventing.StatusEvent{}
}
