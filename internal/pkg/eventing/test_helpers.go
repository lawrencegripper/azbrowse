package eventing

import (
	"testing"
	"time"
)

func WaitForCompletedStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int) StatusEvent {
	return WaitForStatusEvent(t, statusEvents, waitForSec, false)
}

func WaitForFailureStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int) StatusEvent {
	return WaitForStatusEvent(t, statusEvents, waitForSec, true)
}

func WaitForStatusEvent(t *testing.T, statusEvents chan interface{}, waitForSec int, expectError bool) StatusEvent {
	for index := 0; index < waitForSec; index++ {
		select {
		case <-time.After(time.Second):
			t.Log("Waited 1 sec...")
		case statusRaw := <-statusEvents:
			statusEvent := statusRaw.(StatusEvent)
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
		}
	}

	t.Error("Waited for event which never occurred")
	return StatusEvent{}
}
