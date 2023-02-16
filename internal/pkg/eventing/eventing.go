package eventing

import (
	"time"

	"github.com/cskr/pubsub"
	uuid "github.com/satori/go.uuid"
)

const pubSubCapcityPerTopic = 30

// pubSub is the eventbus for the app
// it is then wrapped in methods for use
// to allow future changes.
// NOTE: When capacity is reached events are silently dropped
var pubSub = pubsub.New(pubSubCapcityPerTopic)

// StatusEvent is used to show status information
// in the statusbar
type StatusEvent struct {
	Message    string
	Timeout    time.Duration
	createdAt  time.Time
	InProgress bool
	IsToast    bool
	Failure    bool
	id         uuid.UUID
}

// Icon returns an icon representing the event
func (s *StatusEvent) Icon() string {
	if s.InProgress {
		return "⏳"
	} else if s.Failure {
		return "☠"
	} else {
		return "✓"
	}
}

// ID returns the status message ID
func (s *StatusEvent) ID() string {
	return s.id.String()
}

// SetTimeout sets the timeout for a message, starting from now
func (s *StatusEvent) SetTimeout(t time.Duration) {
	s.createdAt = time.Now()
	s.Timeout = t
}

// CreatedAt returns the time of the message creation
func (s *StatusEvent) CreatedAt() time.Time {
	return s.createdAt
}

// HasExpired returns true if the message has expired
func (s *StatusEvent) HasExpired() bool {
	return s.createdAt.Add(s.Timeout).Before(time.Now())
}

// Update sends and update to the status event
func (s *StatusEvent) Update() {
	// lets assume this should bump the timeout by 5secs
	s.SetTimeout(time.Second * 5)
	SendStatusEvent(s)
}

// Done marks the status event as done
func (s *StatusEvent) Done() {
	s.InProgress = false
	if s.IsToast {
		// Hide completed toast after a few secs
		s.Timeout = time.Duration(time.Second * 5)
	}
	s.Update()
}

// SendFailureStatusFromError sends a status event representing the error mesage and reason
func SendFailureStatusFromError(reason string, err error) *StatusEvent {
	event, _ := SendStatusEvent(&StatusEvent{
		Failure: true,
		Message: reason + ": " + err.Error(),
	})
	return event
}

// SendFailureStatus sends a status event representing the reason
func SendFailureStatus(reason string) *StatusEvent {
	event, _ := SendStatusEvent(&StatusEvent{
		Failure: true,
		Message: reason,
	})
	return event
}

// SendStatusEvent sends status events
func SendStatusEvent(s *StatusEvent) (*StatusEvent, func()) {
	if s.id == [16]byte{} {
		s.id = uuid.NewV4()
	}
	if s.createdAt.IsZero() {
		s.createdAt = time.Now()
	}

	// set default timeout
	if s.Timeout == time.Duration(0) {
		s.Timeout = time.Duration(time.Second * 5)
	}

	doneFunc := func() {
		s.Done()
	}

	Publish("statusEvent", s)
	return s, doneFunc
}

// SubscribeToStatusEvents creates a channel which will receive
// new `StatusEvent` types
func SubscribeToStatusEvents() chan interface{} {
	return SubscribeToTopic("statusEvent")
}

// Unsubscribe from events
func Unsubscribe(ch chan interface{}) {
	pubSub.Unsub(ch)
}

// Publish publishes any event
func Publish(topic string, event interface{}) {
	pubSub.TryPub(event, topic)
}

// SubscribeToTopic creates a channel which will receive event in that topic
// WARNING: Subscribers MUST process events QUICKLY when received or the event sender will BLOCK
// this results in the UI locking up.
func SubscribeToTopic(topic string) chan interface{} {
	return pubSub.Sub(topic)
}
