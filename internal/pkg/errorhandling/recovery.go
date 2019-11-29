package errorhandling

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

var gui *gocui.Gui

// RegisterGuiInstance track the gui instance we can use
// to cleanup
func RegisterGuiInstance(g *gocui.Gui) {
	gui = g
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
		fmt.Println()
		// debug.PrintStack()
		os.Exit(1)
	}
}
