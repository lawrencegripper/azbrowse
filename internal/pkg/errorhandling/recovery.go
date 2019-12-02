package errorhandling

import (
	"fmt"
	"strings"

	// "strings"
	"os"
	"runtime/debug"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

var gui *gocui.Gui
var history = []string{}

// RegisterGuiInstance track the gui instance we can use
// to cleanup
func RegisterGuiInstance(g *gocui.Gui) {
	gui = g

	// Track current view tree for crash logs
	go func() {
		for {
			// Subscribe to navigation events
			navigatedChannel := eventing.SubscribeToTopic("list.prenavigate")
			navigateStateInterface := <-navigatedChannel
			navigateState := navigateStateInterface.(string)

			if navigateState == "GOBACK" && len(history) > 0 {
				history = history[:len(history)-1]
			} else {
				history = append(history, navigateState)
			}
		}
	}()
}

// RecoveryWithCleanup cleans up a go routine panic
// ensuring the terminal is left usable
// Example: (required on all go routines)
//  `defer errorhandling.RecoveryWithCleanup(recover())`
func RecoveryWithCleanup() {
	if r := recover(); r != nil {
		gui.Close()
		fmt.Printf(style.Warning("\n\nSorry a crash occurred\n Error: %s \n"), r)
		fmt.Printf("\n\nPlease visit https://github.com/lawrencegripper/azbrowse/issues to raise a bug.\n")
		fmt.Print("When raising please provide the details below in the issue.")
		fmt.Printf(style.Subtle("\n\nStack Trace: \n%s\n"), debug.Stack())

		navTreeBuilder := strings.Builder{}
		for i, id := range history {
			navTreeBuilder.WriteString("|" + strings.Repeat("-", i+1))
			navTreeBuilder.WriteString("> " + id + "\n")
			if i < len(history)-1 {
				navTreeBuilder.WriteString("| \n")
			}
		}

		fmt.Printf(style.Subtle("Navigation Tree: \n%s\n"), navTreeBuilder.String())
		if len(history) > 0 {
			fmt.Printf(style.Subtle("Retry: \nazbrowse -navigate '%s'\n"), history[len(history)-1])
		}

		fmt.Println()
		// debug.PrintStack()
		os.Exit(1)
	}
}
