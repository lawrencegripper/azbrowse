package errorhandling

import (
	"testing"
	"time"

	"github.com/nbio/st"
	"github.com/stuartleeks/gocui"
)

func TestRecoveryWithCleanup_panicCaptured(t *testing.T) {

	// Start up gocui and configure some settings
	g := &gocui.Gui{}

	// Don't need to pass a goui gui as not used for history
	RegisterGuiInstance(g)

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
