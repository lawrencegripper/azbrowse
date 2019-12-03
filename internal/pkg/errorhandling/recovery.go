package errorhandling

import (
	"context"
	"fmt"
	"os"
	"strings"

	"runtime/debug"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

var history = []string{}
var preNavChannel chan interface{}
var exitFunc func()
var guiClose func()
var ctx context.Context

// RegisterGuiInstance track the gui instance we can use
// to cleanup
func RegisterGuiInstance(ctx context.Context, g *gocui.Gui) {
	history = []string{}

	// exit and guidClose used to allow overriding during testing
	guiClose = func() {
		g.Close()
	}
	exitFunc = func() {
		os.Exit(1)
	}

	// Track current view tree for crash logs
	go func() {
		preNavChannel = eventing.SubscribeToTopic("list.prenavigate")
		for {
			// Stop the routine if the context is cancelled
			select {
			case <-ctx.Done():
				// Unsubscribe from the topic
				eventing.Unsubscribe(preNavChannel)
				// Clear the array
				history = []string{}
				return // returning not to leak the goroutine
			default:
				// Carry on working
			}

			// Subscribe to navigation events
			navigateStateInterface := <-preNavChannel
			navigateState, ok := navigateStateInterface.(string)
			if !ok {
				continue
			}

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
		guiClose()
		fmt.Printf(style.Warning("\n\nSorry a crash occurred\n Error: %s \n"), r)
		fmt.Printf("\n\nPlease visit https://github.com/lawrencegripper/azbrowse/issues to raise a bug.\n")
		fmt.Print("When raising please provide the details below in the issue. \nNote `Navigation Tree` may contain sensitive information, please review before posting.")
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
		exitFunc()
	}
}
