package errorhandling

import (
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/nbio/st"
)

func TestRecoveryWithCleanup_tracksHistory_basic(t *testing.T) {
	// Don't need to pass a goui gui as not used for history
	RegisterGuiInstance(nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "item3")

	<-time.After(800 * time.Millisecond)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}

func TestRecoveryWithCleanup_tracksHistory_withBack(t *testing.T) {
	// Don't need to pass a goui gui as not used for history
	RegisterGuiInstance(nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "shouldntseeme")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "shouldntseeme")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "item3")

	<-time.After(800 * time.Millisecond)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}

func TestRecoveryWithCleanup_tracksHistory_tooManyBacks(t *testing.T) {
	// Don't need to pass a goui gui as not used for history
	RegisterGuiInstance(nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "GOBACK")

	<-time.After(800 * time.Millisecond)

	st.Assert(t, history, []string{})
}

func TestRecoveryWithCleanup_tracksHistory_invalidType(t *testing.T) {
	// Don't need to pass a goui gui as not used for history
	RegisterGuiInstance(nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", struct{ bob string }{})
	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "item3")

	<-time.After(800 * time.Millisecond)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}
