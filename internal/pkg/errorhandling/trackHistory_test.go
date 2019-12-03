package errorhandling

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/nbio/st"
)

func waitForEvents(number int) {
	// Wait for the events to come through
	timeout := time.After(5 * time.Second)
Loop:
	for len(history) < number {
		//wait
		select {
		case <-timeout:
			fmt.Println("Timed out waiting")
			break Loop
		default:
			// Wait a bit
			fmt.Println("Waiting for events to be processed")

			<-time.After(50 * time.Millisecond)
		}
	}
}

func waitForReadyToRegister() {
	for started {
		fmt.Println("Waiting for previous instance to stop")
		<-time.After(50 * time.Millisecond)
	}
}

func TestRecoveryWithCleanup_tracksHistory_startupAndCancel(t *testing.T) {
	waitForReadyToRegister()

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, nil)

	eventing.Publish("list.prenavigate", "item1")
	<-time.After(5 * time.Second)

}

func TestRecoveryWithCleanup_tracksHistory_basic(t *testing.T) {
	waitForReadyToRegister()

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, nil)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "item3")

	waitForEvents(3)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}

func TestRecoveryWithCleanup_tracksHistory_withBack(t *testing.T) {
	waitForReadyToRegister()

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "shouldntseeme")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "shouldntseeme")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "item3")

	waitForEvents(3)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}

func TestRecoveryWithCleanup_tracksHistory_tooManyBacks(t *testing.T) {
	waitForReadyToRegister()

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, nil)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "GOBACK")
	eventing.Publish("list.prenavigate", "GOBACK")

	<-time.After(800 * time.Millisecond)

	st.Assert(t, history, []string{})
}

func TestRecoveryWithCleanup_tracksHistory_invalidType(t *testing.T) {
	waitForReadyToRegister()

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, nil)
	defer eventing.Unsubscribe(preNavChannel)

	<-time.After(800 * time.Millisecond)

	eventing.Publish("list.prenavigate", struct{ bob string }{})
	eventing.Publish("list.prenavigate", "item1")
	eventing.Publish("list.prenavigate", "item2")
	eventing.Publish("list.prenavigate", "item3")

	waitForEvents(3)

	st.Assert(t, history, []string{"item1", "item2", "item3"})
}
