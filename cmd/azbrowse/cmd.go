package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/filesystem"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const accountCacheKey = "accountCache"
const navigateCacheKey = "navigateCache"

func handleCommandAndArgs() {

	rootCmd := createRootCmd()

	rootCmd.AddCommand(createVersionCommand())
	rootCmd.AddCommand(createAzfsCommand())
	rootCmd.AddCommand(createCompleteCommand(rootCmd))

	// Special case used to generate markdown docs for the commands
	if os.Getenv("AZB_GEN_COMMAND_MARKDOWN") == "TRUE" {
		rootCmd.DisableAutoGenTag = true
		err := doc.GenMarkdownTree(rootCmd, "./docs/commandline/")
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	_ = rootCmd.Execute()
}

func createRootCmd() *cobra.Command {
	var demo bool
	var debug bool
	var navigateTo string
	var fuzzerDurationMinutes int
	var tenantID string
	var subscription string

	cmd := &cobra.Command{
		Use:   "azbrowse",
		Short: "An interactive CLI for browsing Azure",
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
			} else if subscription != "" {
				// [?name=='subscriptionName' || id== 'ab99572b-a482-4a02-acf3-96ba46e90f76'].tenantId
				// get tenant id from subscription id/name
				query := fmt.Sprintf("[?name=='%[1]s' || id== '%[1]s'].tenantId", subscription)
				out, err := exec.Command("az", "account", "list", "--query", query, "--output", "tsv").Output()
				if err != nil {
					_ = cmd.Usage()
					os.Exit(1)
				}
				settings.TenantID = strings.TrimSuffix(string(out), "\n")

				if settings.NavigateToID == "" {
					// Set to navigate to the subscription only if --navigate hasn't also been set
					// get subscription id from id/name
					query := fmt.Sprintf("[?name=='%[1]s' || id== '%[1]s'].id", subscription)
					out, err := exec.Command("az", "account", "list", "--query", query, "--output", "tsv").Output()
					if err != nil {
						_ = cmd.Usage()
						os.Exit(1)
					}
					settings.NavigateToID = "/subscriptions/" + strings.TrimSuffix(string(out), "\n")
				}
			}
			run(&settings)
		},
	}
	cmd.Flags().StringVarP(&navigateTo, "navigate", "n", "", "(optional) navigate to resource by resource ID")
	cmd.Flags().StringVar(&tenantID, "tenant-id", "", "(optional) specify the tenant id to get an access token for (see az account list -o json)")
	cmd.Flags().StringVarP(&subscription, "subscription", "s", "", "(optional) specify a subscription to load")
	cmd.Flags().BoolVar(&debug, "debug", false, "run in debug mode")
	cmd.Flags().BoolVar(&demo, "demo", false, "run in demo mode to filter sensitive output")
	cmd.Flags().IntVar(&fuzzerDurationMinutes, "fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")

	if err := cmd.RegisterFlagCompletionFunc("subscription", subscriptionAutocompletion); err != nil {
		panic(err)
	}

	if err := cmd.RegisterFlagCompletionFunc("navigate", navigateAutocompletion(&subscription)); err != nil {
		panic(err)
	}

	return cmd
}

func getResourceListAndUpdateCache() (string, error) {
	query := `resourcecontainers | where type == "microsoft.resources/subscriptions/resourcegroups"` +
		" | union (resources)" +
		" | project name, id, subscriptionId, tenantId"

	graphArgs := []string{"graph", "query", "--graph-query", query, "--output", "json"}
	cobra.CompDebugln(fmt.Sprintf("command: %+v", graphArgs), true)

	out, err := exec.Command("az", graphArgs...).Output()
	if err != nil {
		cobra.CompErrorln("az graph command failed")
		return "", errors.Wrap(err, "Failed azGraph when updating cache")
	}

	err = storage.PutCacheForTTL(navigateCacheKey, string(out))
	if err != nil {
		cobra.CompErrorln("Failed to save graph response to navigateCache")
		return "", errors.Wrap(err, "Failed storing azGraph result when updating cache")
	}

	return string(out), nil
}

type graphItem struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
}

func navigateAutocompletion(subscription *string) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Debugging
		cobra.CompDebugln(fmt.Sprintf("Subscription: %+v", *subscription), true)
		cobra.CompDebugln(fmt.Sprintf("toComplete: %+v", toComplete), true)
		// Cache the resource list for x mins
		validCache, value, err := storage.GetCacheWithTTL(navigateCacheKey, time.Minute*5)
		if !validCache || err != nil {
			resourceList, err := getResourceListAndUpdateCache()
			if err != nil {
				return []string{}, cobra.ShellCompDirectiveError
			}
			value = resourceList
		}

		// The `subscription` field on azbrowse can be either a subscription name or GUID.
		// The response from Graph only has the `subscriptionID`.
		// This code maps from SubName to GUID.
		var subscriptionGUID string
		if subscription != nil && *subscription != "" {
			accountList, err := getAccountList()
			if err != nil {
				cobra.CompError(err.Error())
			} else {
				for _, sub := range accountList {
					if sub.Name == *subscription || sub.ID == *subscription {
						subscriptionGUID = sub.ID
						break
					}
				}
			}
		}

		var graphResponse []graphItem
		err = json.Unmarshal([]byte(value), &graphResponse)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		// Used to limit completion suggestions to only the next segement deep
		toCompleteDepth := len(strings.Split(toComplete, "/"))
		isPartialDepth := false

		values := make([]string, 0, len(graphResponse))
		for _, item := range graphResponse {
			// Filter to only subs the user is interested in
			if subscriptionGUID != "" && subscriptionGUID != item.SubscriptionID {
				// Skip as not in the subscription we're interested in
				continue
			}

			// Provide completion only to the next segment of the resource
			resourceDepth := len(strings.Split(item.ID, "/"))
			if toCompleteDepth+1 < resourceDepth {
				values = append(values, strings.Join(strings.Split(item.ID, "/")[:toCompleteDepth], "/")+"/")
				isPartialDepth = true
				continue
			}

			values = append(values, item.ID)
		}

		// If this is a completion limted by segement don't put a space
		if isPartialDepth {
			return values, cobra.ShellCompDirectiveNoSpace
		}

		return values, cobra.ShellCompDirectiveNoFileComp
	}
}

type accountItem struct {
	CloudName string `json:"cloudName"`
	ID        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
	Name      string `json:"name"`
	State     string `json:"state"`
	TenantID  string `json:"tenantId"`
	User      struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}

// This allows azbrowse to update the account cache used for autocompletion
// due to it's use in completion func errors are suppressed
func getAccountListAndUpdateCache() ([]accountItem, error) {
	out, err := exec.Command("az", "account", "list", "--output", "json").Output()
	if err != nil {
		return nil, errors.Wrap(err, "Failed invoking az to update account list cache")
	}

	var accounts []accountItem
	err = json.Unmarshal(out, &accounts)
	if err != nil {
		return nil, errors.Wrap(err, "Failed unmarshalling response from az to update account list cache")
	}

	err = storage.PutCacheForTTL(accountCacheKey, string(out))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to save account list to cache")
	}
	return accounts, nil
}

func getAccountList() ([]accountItem, error) {
	validCache, value, err := storage.GetCacheWithTTL(accountCacheKey, time.Hour*6)
	if !validCache || err != nil {
		azAccountOutput, err := getAccountListAndUpdateCache()
		if err != nil {
			return nil, err
		}
		return azAccountOutput, nil
	}

	var accountList []accountItem
	err = json.Unmarshal([]byte(value), &accountList)
	if err != nil {
		return nil, errors.Wrap(err, "Failed unmarshalling from cache to get account list")
	}

	return accountList, nil
}

// Provide support for autocompleting subscriptions
func subscriptionAutocompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	accountList, err := getAccountList()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	values := make([]string, len(accountList)*2)
	for _, a := range accountList {
		values = append(values, a.Name)
		values = append(values, a.ID)
	}

	return values, cobra.ShellCompDirectiveNoFileComp
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
