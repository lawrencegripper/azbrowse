package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/filesystem"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
)

func handleRunCmd(
	settings *config.Settings,
	demo *bool,
	debug *bool,
	navigateResource *string,
	fuzzerDurationMinutes *int,
	tenantID *string) int {

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
		settings.ShouldRender = false
	}

	if fuzzerDurationMinutes != nil && *fuzzerDurationMinutes > 0 {
		settings.FuzzerEnabled = true
		settings.FuzzerDurationMinutes = *fuzzerDurationMinutes
	}

	if tenantID != nil {
		settings.TenantID = *tenantID
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
	settings := config.Settings{
		ShouldRender: true,
	}

	// Root command
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	// Root flags
	runDemo := runCmd.Bool("demo", false, "run in demo mode to filter sensitive output")
	runDebug := runCmd.Bool("debug", false, "run in debug mode")
	runNavigate := runCmd.String("navigate", "", "navigate to resource")
	runFuzzer := runCmd.Int("fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")
	runTenantID := runCmd.String("tenant-id", "", "(optional) specify the tenant id to get an access token for (see `az")

	// Version command
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	// Version usage
	versionCmd.Usage = runCmd.Usage

	// azfs command
	azfsCmd := flag.NewFlagSet("azfs", flag.ExitOnError)
	azfsAcceptRisk := azfsCmd.Bool("accept-risk", false, "Warning: accept the risk of running this alpha quality filesystem. Do not use on production subscriptions")
	azfsEditEnabled := azfsCmd.Bool("edit", false, "enable editing")
	azfsMount := azfsCmd.String("mount", "/mnt/azfs", "location to mount filesystem")
	azfsSubscription := azfsCmd.String("sub", "", "filter to only show a single subscription, provide the 'name' or 'id' of the subscription")
	azfsDemo := azfsCmd.Bool("demo", false, "run in demo mode to filter sensitive output")

	// Root usage
	runCmd.Usage = func() {
		// Usage
		fmt.Fprintf(os.Stderr, "Usage:  azbrowse [OPTIONS] COMMAND\n")

		// Description
		fmt.Fprintf(os.Stderr, "\nA terminal browser for Microsoft Azure.\n")

		// Flags
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		runCmd.PrintDefaults()

		// Global Flags
		fmt.Fprintf(os.Stderr, "\nGlobal Options:\n")
		fmt.Fprintf(os.Stderr, "  -h, --help    Print help information\n")

		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		fmt.Fprintf(os.Stderr, "  version       Print version information\n")
		fmt.Fprintf(os.Stderr, "  azfs          Mount the Azure ARM API as a fuse filesystem\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		azfsCmd.PrintDefaults()
	}

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
		case "azfs":
			if err := azfsCmd.Parse(os.Args[2:]); err != nil {
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
	if azfsCmd.Parsed() {
		os.Exit(func() int {
			if !*azfsAcceptRisk {
				fmt.Println("This is an alpha quality feature you must accept the risk to your subscription by adding '-accept-risk'. Use '-sub subscriptionname' to only mount a single subscription")
				os.Exit(1)
			}
			closer, err := filesystem.Run(*azfsMount, *azfsSubscription, *azfsEditEnabled, *azfsDemo)
			if err != nil {
				panic(err)
			}
			c := make(chan os.Signal, 2)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			<-c
			fmt.Println("Ctrl+C pressed attempting to unmounting azfs and exit. \n If you see a 'device busy' message exit all processes using the filesystem then unmount will proceed \n Alternatively press CTRL+C again to force exit.")
			go func() {
				<-c
				os.Exit(0)
			}()
			closer()
			os.Exit(0)
			return 1
		}())
	}
	if runCmd.Parsed() {
		os.Exit(handleRunCmd(&settings, runDemo, runDebug, runNavigate, runFuzzer, runTenantID))
	}

	// If no command was parsed, fallback to usage
	usageAndExit()
}
