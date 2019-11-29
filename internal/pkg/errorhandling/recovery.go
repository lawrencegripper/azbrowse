package errorhandling

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

var gui *gocui.Gui
var lastNavigationEvent []byte

// RegisterGuiInstance track the gui instance we can use
// to cleanup
func RegisterGuiInstance(g *gocui.Gui) {
	gui = g

	go func() {
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")

		for {
			navEvent := <-navigatedChannel
			if navEvent == nil {
				continue
			}
			lastNavigationEvent, _ = json.Marshal(navEvent)

		}
	}()
}

// RecoveryWithCleainup cleans up a go routine panic
// ensuring the terminal is left usable
// Example: (required on all go routines)
//  `defer errorhandling.RecoveryWithCleainup(recover())`
func RecoveryWithCleainup() {
	if r := recover(); r != nil {
		gui.Close()
		fmt.Printf(style.Warning("\n\nSorry a crash occurred\n Error: %s \n"), r)
		fmt.Printf("\n\nPlease visit https://github.com/lawrencegripper/azbrowse/issues to raise a bug.\n")
		fmt.Print("When raising please provide the details below in the issue.")
		fmt.Printf(style.Subtle("\n\nStack Trace: \n%s\n"), debug.Stack())
		fmt.Printf(style.Subtle("\n\nLast Navigation Event: \n%s\n"), lastNavigationEvent)
		fmt.Println()
		// debug.PrintStack()
		os.Exit(1)
	}
}
