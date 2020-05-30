package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/filesystem"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/spf13/cobra"
)

func handleCommandAndArgs() {

	rootCmd := createRootCmd()

	rootCmd.AddCommand(createVersionCommand())
	rootCmd.AddCommand(createAzfsCommand())
	rootCmd.AddCommand(createCompleteCommand(rootCmd))

	_ = rootCmd.Execute()
}

func createRootCmd() *cobra.Command {
	var demo bool
	var debug bool
	var navigateTo string
	var fuzzerDurationMinutes int
	var tenantID string
	var tenantIDFromSubscription string

	cmd := &cobra.Command{
		Use: "azbrowse",
		Run: func(cmd *cobra.Command, args []string) {

			settings := config.Settings{
				ShouldRender: true,
			}
			if demo {
				settings.HideGuids = true
			}

			if debug {
				settings.EnableTracing = true
				tracing.EnableDebug()
				config.SetDebuggingEnabled(true)
			}

			if navigateTo != "" {
				settings.NavigateToID = navigateTo
				settings.ShouldRender = false
			}

			if fuzzerDurationMinutes > 0 {
				settings.FuzzerEnabled = true
				settings.FuzzerDurationMinutes = fuzzerDurationMinutes
			}

			if tenantID != "" {
				settings.TenantID = tenantID
			} else if tenantIDFromSubscription != "" {
				// [?name=='mpeck-stuartle' || id== '36ce814f-1b29-4695-9bde-1e2ad14bda0f'].tenantId
				query := fmt.Sprintf("[?name=='%[1]s' || id== '%[1]s'].tenantId", tenantIDFromSubscription)
				out, err := exec.Command("az", "account", "list", "--query", query, "--output", "tsv").Output()
				if err != nil {
					_ = cmd.Usage()
					os.Exit(1)
				}
				settings.TenantID = strings.TrimSuffix(string(out), "\n")
			}
			run(&settings)
		},
	}
	cmd.Flags().BoolVar(&demo, "demo", false, "run in demo mode to filter sensitive output")
	cmd.Flags().BoolVar(&debug, "debug", false, "run in debug mode")
	cmd.Flags().StringVar(&navigateTo, "navigate", "", "navigate to resource")
	cmd.Flags().IntVar(&fuzzerDurationMinutes, "fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")
	cmd.Flags().StringVar(&tenantID, "tenant-id", "", "(optional) specify the tenant id to get an access token for (see az account list -o json)")
	cmd.Flags().StringVar(&tenantIDFromSubscription, "tenant-id-from-sub", "", "(optional) specify a subscription to identify the tenant-id to use")

	if err := cmd.RegisterFlagCompletionFunc("tenant-id-from-sub", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

		out, err := exec.Command("az", "account", "list", "--query", "[].[name, id] | [] | sort(@)", "--output", "tsv").Output()
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		reader := bytes.NewReader(out)
		scanner := bufio.NewScanner(reader)
		values := []string{}
		for scanner.Scan() {
			values = append(values, "\""+strings.ReplaceAll(scanner.Text(), " ", "\\ ")+"\"")
		}

		return values, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		panic(err)
	}

	return cmd
}

func createVersionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
			fmt.Println(commit)
			fmt.Println(date)
			fmt.Println(goversion)
			fmt.Printf("%s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}

	return cmd
}

func createAzfsCommand() *cobra.Command {

	var acceptRisk bool
	var enableEditing bool
	var mount string
	var subscription string
	var demo bool

	cmd := &cobra.Command{
		Use:   "azfs",
		Short: "Mount the Azure ARM API as a fuse filesystem",
		Run: func(cmd *cobra.Command, args []string) {
			if !acceptRisk {
				fmt.Println("This is an alpha quality feature you must accept the risk to your subscription by adding '-accept-risk'. Use '-sub subscriptionname' to only mount a single subscription")
				os.Exit(1)
			}
			closer, err := filesystem.Run(mount, subscription, enableEditing, demo)
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
		},
	}
	cmd.Flags().BoolVar(&acceptRisk, "accept-risk", false, "Warning: accept the risk of running this alpha quality filesystem. Do not use on production subscriptions")
	cmd.Flags().BoolVar(&enableEditing, "edit", false, "enable editing")
	cmd.Flags().StringVar(&mount, "mount", "/mnt/azfs", "location to mount filesystem")
	cmd.Flags().StringVar(&subscription, "sub", "", "filter to only show a single subscription, provide the 'name' or 'id' of the subscription")
	cmd.Flags().BoolVar(&demo, "demo", false, "run in demo mode to filter sensitive output")

	return cmd
}

func createCompleteCommand(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion SHELL",
		Short: "Generates shell completion scripts",
		Long: `To load completion run
	
	. <(azbrowse completion SHELL)
	Valid values for SHELL are : bash, fish, powershell, zsh
	
	For example, to configure your bash shell to load completions for each session add to your bashrc
	
	# ~/.bashrc or ~/.profile
	source <(azbrowse completion bash)
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				_ = cmd.Usage()
				os.Exit(1)
			}
			shell := args[0]
			var err error
			switch strings.ToLower(shell) {
			case "bash":
				err = rootCmd.GenBashCompletion(os.Stdout)
			case "fish":
				err = rootCmd.GenFishCompletion(os.Stdout, true)
			case "powershell":
				err = rootCmd.GenPowerShellCompletion(os.Stdout)
			case "zsh":
				err = rootCmd.GenZshCompletion(os.Stdout)
			default:
				fmt.Printf("Unsupported SHELL value: '%s'\n", shell)
				return cmd.Usage()
			}

			return err
		},
	}
	return cmd
}
