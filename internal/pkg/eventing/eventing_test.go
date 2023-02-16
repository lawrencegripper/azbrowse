package eventing

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestStatusEvent_HasExpired(t *testing.T) {

	tests := []struct {
		name       string
		fields     StatusEvent
		hasExpired bool
	}{
		{
			name: "expired event",
			fields: StatusEvent{
				Message:   "bob",
				createdAt: time.Now().Add(time.Second * 45 * -1), // 45 secs ago
				Timeout:   time.Second * 15,
			},
			hasExpired: true,
		},
		{
			name: "non-expired event",
			fields: StatusEvent{
				Message:   "bob",
				createdAt: time.Now().Add(time.Second * 10 * -1), // 10 secs ago
				Timeout:   time.Second * 15,
			},
			hasExpired: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tt.fields
			if got := s.HasExpired(); got != tt.hasExpired {
				t.Errorf("StatusEvent.HasExpired() = %v, want %v", got, tt.hasExpired)
			}
		})
	}
}

func TestStatusEvent_Update_shouldntSetCreatedAtOrID(t *testing.T) {
	createdAt := time.Now().Add(time.Second * 5000)
	id := uuid.NewV4()

	// Create already initialized event to simulate an update scenario
	s := &StatusEvent{
		createdAt: createdAt,
		id:        id,
	}

	// Call the update
	s, _ = SendStatusEvent(s)

	if !s.createdAt.Equal(createdAt) {
		t.Error("SendStatusEvent shouldn't update an already set 'createdAt' field")
	}

	if s.id != id {
		t.Error("SendStatusEvent shouldn't update an already set 'id' field")
	}

}

func TestStatusEvent_End2End_Expired(t *testing.T) {
	newEvents := SubscribeToStatusEvents()
	defer Unsubscribe(newEvents)

	SendStatusEvent(&StatusEvent{
		Message: "bob",
		Timeout: time.Second * 5,
	})

	<-time.After(time.Second * 6)

	eventRaw := <-newEvents
	eventObj := eventRaw.(*StatusEvent)

	if eventObj.HasExpired() {
		t.Log("Event successfully expired")
	} else {
		t.Error("Event didn't expire")
	}
}

// Todo: Maybe squish into above as table driven test
func TestStatusEvent_End2End_Valid(t *testing.T) {
	newEvents := SubscribeToStatusEvents()
	defer Unsubscribe(newEvents)

	SendStatusEvent(&StatusEvent{
		Message: "bob",
		Timeout: time.Second * 10,
	})

	<-time.After(time.Second * 5)

	eventRaw := <-newEvents
	eventObj := eventRaw.(*StatusEvent)

	if eventObj.HasExpired() {
		t.Error("Event shouldn't have expired")
	}
}
