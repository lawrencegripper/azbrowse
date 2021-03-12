package errorhandling

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/nbio/st"
)

func TestRecoveryWithCleanup_panicCaptured(t *testing.T) {
	waitForReadyToRegister()

	fmt.Println("NOTE: This test will print the error message generated by the `recovery` func. This is expected.")

	// Clean up and stop `recovery` tracking history when we're done with the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start up gocui and configure some settings
	g := &gocui.Gui{}

	// Don't need to pass a goui gui as not used for history
	RegisterGuiAndStartHistoryTracking(ctx, g)

	// Populate history
	history = []string{
		"item1",
		"item2",
		"item3",
	}

	exitCalled := false
	guiCloseCalled := false

	// Override the exit function of recovery so we can
	// assert it's called
	exitFunc = func() {
		exitCalled = true
	}

	guiClose = func() {
		guiCloseCalled = true
	}

	go func() {
		defer RecoveryWithCleanup()

		panic("Oooooh no an error")
	}()

	<-time.After(300 * time.Millisecond)

	st.Assert(t, exitCalled, true)
	st.Assert(t, guiCloseCalled, true)
}
