package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

func handleRunCmd(settings *Settings, demo *bool, debug *bool, navigateResource *string, fuzzerDurationMinutes *int) int {
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

func handleVersionCmd(settings *Settings) int {
	fmt.Println(version)
	fmt.Println(commit)
	fmt.Println(date)
	fmt.Println(goversion)
	fmt.Println(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	return 0
}

func usage() int {
	flag.PrintDefaults()
	return 1
}

func handleCommandAndArgs() {
	settings := Settings{}

	// Root command
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runDemo := runCmd.Bool("demo", false, "run in demo mode to filter sensitive output")
	runDebug := runCmd.Bool("debug", false, "run in debug mode")
	runNavigate := runCmd.String("navigate", "", "navigate to resource")
	runFuzzer := runCmd.Int("fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")

	if len(os.Args) < 2 {
		if err := runCmd.Parse(os.Args[1:]); err != nil {
			os.Exit(usage())
		}
	}

	// Subcommands and flags
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "version":
			if err := versionCmd.Parse(os.Args[2:]); err != nil {
				os.Exit(usage())
			}
		default:
			if err := runCmd.Parse(os.Args[1:]); err != nil {
				os.Exit(usage())
			}
		}
	}

	if versionCmd.Parsed() {
		os.Exit(handleVersionCmd(&settings))
	}
	if runCmd.Parsed() {
		os.Exit(handleRunCmd(&settings, runDemo, runDebug, runNavigate, runFuzzer))
	}

	// Default
	os.Exit(usage())
}
