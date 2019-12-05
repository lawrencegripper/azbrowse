package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

func handleRunCmd(settings *config.Settings, demo *bool, debug *bool, navigateResource *string, fuzzerDurationMinutes *int) int {
	if demo != nil && *demo {
		settings.HideGuids = true
	}

	if debug != nil && *debug {
		settings.EnableTracing = true
		tracing.EnableDebug()
		config.SetDebuggingEnabled(true)
	}

	if navigateResource != nil && len(*navigateResource) > 0 {
		settings.NavigateToID = *navigateResource
	}

	if fuzzerDurationMinutes != nil && *fuzzerDurationMinutes > 0 {
		settings.FuzzerEnabled = true
		settings.FuzzerDurationMinutes = *fuzzerDurationMinutes
	}

	run(settings)
	return 0
}

func handleVersionCmd(settings *config.Settings) int {
	fmt.Println(version)
	fmt.Println(commit)
	fmt.Println(date)
	fmt.Println(goversion)
	fmt.Println(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	return 0
}

func usageAndExit() {
	flag.Usage()
	os.Exit(1)
}

func handleCommandAndArgs() {
	settings := config.Settings{}

	// Root command
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	// Root flags
	runDemo := runCmd.Bool("demo", false, "run in demo mode to filter sensitive output")
	runDebug := runCmd.Bool("debug", false, "run in debug mode")
	runNavigate := runCmd.String("navigate", "", "navigate to resource")
	runFuzzer := runCmd.Int("fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")
	// Root usage
	runCmd.Usage = func() {
		// Usage
		fmt.Fprintf(os.Stderr, "Usage:  azbrowse [OPTIONS] COMMAND\n")

		// Description
		fmt.Fprintf(os.Stderr, "\nA terminal browzer for Microsoft Azure.\n")

		// Flags
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		runCmd.PrintDefaults()

		// Global Flags
		fmt.Fprintf(os.Stderr, "\nGlobal Options:\n")
		fmt.Fprintf(os.Stderr, "  -h, --help    Print help information\n")

		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		fmt.Fprintf(os.Stderr, "  version       Print version information\n")
	}

	// Version command
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	// Version usage
	versionCmd.Usage = runCmd.Usage

	// Handle root command
	if len(os.Args) < 2 {
		if err := runCmd.Parse(os.Args[1:]); err != nil {
			usageAndExit()
		}
	}

	// Handle subcommands
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "version":
			if err := versionCmd.Parse(os.Args[2:]); err != nil {
				usageAndExit()
			}
		default: // default to root command
			if err := runCmd.Parse(os.Args[1:]); err != nil {
				usageAndExit()
			}
		}
	}

	// Detect which command was parsed and invoke  handler
	if versionCmd.Parsed() {
		os.Exit(handleVersionCmd(&settings))
	}
	if runCmd.Parsed() {
		os.Exit(handleRunCmd(&settings, runDemo, runDebug, runNavigate, runFuzzer))
	}

	// If no command was parsed, fallback to usage
	usageAndExit()
}
