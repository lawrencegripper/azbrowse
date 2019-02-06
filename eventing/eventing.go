package eventing

import (
	"time"

	"github.com/cskr/pubsub"
	"github.com/satori/go.uuid"
)

// pubSub is the eventbus for the app
// it is then wrapped in methods for use
// to allow future changes
var pubSub = pubsub.New(1)

// StatusEvent is used to show status information
// in the statusbar
type StatusEvent struct {
	Message    string
	Timeout    time.Duration
	createdAt  time.Time
	InProgress bool
	Failure    bool
	id         uuid.UUID
}

// ID returns the status message ID
func (s *StatusEvent) ID() string {
	return s.id.String()
}

// CreatedAt returns the time of the message creation
func (s *StatusEvent) CreatedAt() time.Time {
	return s.createdAt
}

// HasExpired returns true if the message has expired
func (s *StatusEvent) HasExpired() bool {
	return s.createdAt.Add(s.Timeout).After(time.Now())
}

// Update sends and update to the status event
func (s *StatusEvent) Update() {
	SendStatusEvent(*s)
}

// SendStatusEvent sends status events
func SendStatusEvent(s StatusEvent) (StatusEvent, func()) {
	s.id = uuid.NewV4()
	s.createdAt = time.Now()

	// set default timeout
	if s.Timeout == time.Duration(0) {
		s.Timeout = time.Duration(time.Second * 5)
	}

	doneFunc := func() {
		s.InProgress = false
		SendStatusEvent(s)
	}

	pubSub.Pub(s, "statusEvent")
	return s, doneFunc
}

// SubscribeToStatusEvents creates a channel which will receive
// new `StatusEvent` types
func SubscribeToStatusEvents() chan interface{} {
	return pubSub.Sub("statusEvent")
}
