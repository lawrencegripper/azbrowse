package eventing

import (
	"time"

	"github.com/cskr/pubsub"
	"github.com/satori/go.uuid"
)

// pubSub is the eventbus for the app
var pubSub = pubsub.New(1)

// StatusEvent is used to show status information
// in the statusbar
type StatusEvent struct {
	Message    string
	Timeout    time.Duration
	InProgress bool
	id         uuid.UUID
}

// SendStatusEvent sends status events
func SendStatusEvent(s StatusEvent) StatusEvent {
	s.id = uuid.NewV4()
	pubSub.Pub(s, "statusEvent")
	return s
}

// SubscribeToStatusEvents creates a channel which will recieve
// new `StatusEvent` types
func SubscribeToStatusEvents() chan interface{} {
	return pubSub.Sub("statusEvent")
}
