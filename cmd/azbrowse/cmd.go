package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const accountCacheKey = "accountCache"
const navigateCacheKey = "navigateCache"
const resumeNodeIDKey = "resumeNode"
const resumeTenantIDKey = "resumeTenant"

func handleCommandAndArgs() {

	rootCmd := createRootCmd()

	rootCmd.AddCommand(createVersionCommand())
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
	var resume bool
	var navigateTo string
	var fuzzerDurationMinutes int
	var tenantID string
	var subscription string
	var mouse bool

	// Start tracking the last node navigated to in storage for the `resume` command
	go func() {
		defer errorhandling.RecoveryWithCleanup()

		navigatedChannel := eventing.SubscribeToTopic("list.navigated")
		for {
			navigateStateInterface := <-navigatedChannel
			navigateState := navigateStateInterface.(views.ListNavigatedEventState)
			storage.PutCache(resumeNodeIDKey, navigateState.NodeID) //nolint: errcheck
			storage.PutCache(resumeTenantIDKey, tenantID)           //nolint: errcheck
		}
	}()

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

			if mouse {
				settings.MouseEnabled = mouse
			}

			if debug {
				settings.EnableTracing = true
				tracing.EnableDebug()
				config.SetDebuggingEnabled(true)
			}

			if navigateTo != "" {
				settings.NavigateToID = navigateTo
				settings.ShouldRender = false
			} else if resume {
				nodeID, err := storage.GetCache(resumeNodeIDKey)
				if err != nil {
					fmt.Println("Cannot resume: " + err.Error())
					os.Exit(1)
				}
				currentTenantID, err := storage.GetCache(resumeTenantIDKey)
				if err != nil {
					fmt.Println("Cannot resume: " + err.Error())
					os.Exit(1)
				}
				settings.TenantID = currentTenantID
				settings.NavigateToID = nodeID
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

			// Hack: To allow resume to track tenant ud easily
			tenantID = settings.TenantID

			run(&settings)
		},
	}
	cmd.Flags().StringVarP(&navigateTo, "navigate", "n", "", "(optional) navigate to resource by resource ID")
	cmd.Flags().StringVar(&tenantID, "tenant-id", "", "(optional) specify the tenant id to get an access token for (see az account list -o json)")
	cmd.Flags().StringVarP(&subscription, "subscription", "s", "", "(optional) specify a subscription to load")
	cmd.Flags().BoolVarP(&resume, "resume", "r", false, "(optional) resume navigating from your last session")
	cmd.Flags().BoolVar(&debug, "debug", false, "run in debug mode")
	cmd.Flags().BoolVar(&demo, "demo", false, "run in demo mode to filter sensitive output")
	cmd.Flags().IntVar(&fuzzerDurationMinutes, "fuzzer", -1, "run fuzzer (optionally specify the duration in minutes)")
	cmd.Flags().BoolVarP(&mouse, "mouse", "m", false, "(optional) enable mouse support. Note this disables normal text selection in the terminal")

	if err := cmd.RegisterFlagCompletionFunc("subscription", subscriptionAutocompletion); err != nil {
		panic(err)
	}

	if err := cmd.RegisterFlagCompletionFunc("navigate", navigateAutocompletion(&subscription)); err != nil {
		panic(err)
	}

	return cmd
}

func getResourceListAndUpdateCache(subscriptions []string, client *armclient.Client) (string, error) {

	query := `resourcecontainers | where type == 'microsoft.resources/subscriptions/resourcegroups'` +
		" | union (resources)" +
		" | project name, id, subscriptionId, tenantId"

	out, err := client.DoResourceGraphQueryReturningObjectArray(context.Background(), subscriptions, query)
	if err != nil {
		cobra.CompErrorln("az graph rest query failed:" + err.Error())
		return "", fmt.Errorf("Failed azGraph when updating cache: %w", err)
	}

	err = storage.PutCacheForTTL(navigateCacheKey, string(out))
	if err != nil {
		cobra.CompErrorln("Failed to save graph response to navigateCache")
		return "", fmt.Errorf("Failed storing azGraph result when updating cache: %w", err)
	}

	return string(out), nil
}

type graphItem struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
}

type graphResponse struct {
	TotalRecords int         `json:"totalRecords"`
	Count        int         `json:"count"`
	Data         []graphItem `json:"data"`
}

func navigateAutocompletion(subscription *string) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Debugging
		cobra.CompDebugln(fmt.Sprintf("Subscription: %+v", *subscription), true)
		cobra.CompDebugln(fmt.Sprintf("toComplete: %+v", toComplete), true)

		accountList, err := getAccountList()
		if err != nil {
			cobra.CompError(err.Error())
		}
		// Graph queries are made against allSubscriptionGUIDs. We build up a list of all subs to query
		// queries are always done over all subs so results can be cached for future autocompletions.
		// If a user is only interested in one sub the filtering to a single subscription is done by this function later.
		var allSubscriptionGUIDs []string
		for _, sub := range accountList {
			allSubscriptionGUIDs = append(allSubscriptionGUIDs, sub.ID)
		}
		// The `subscription` field on azbrowse can be either a subscription name or GUID.
		// The response from Graph only has the `subscriptionID`.
		// This code maps from SubName to GUID.
		var subscriptionGUID string
		if subscription != nil && *subscription != "" {
			for _, sub := range accountList {
				allSubscriptionGUIDs = append(allSubscriptionGUIDs, sub.ID)
				if sub.Name == *subscription || sub.ID == *subscription {
					subscriptionGUID = sub.ID
					break
				}
			}
		}

		// Cache the resource list for x mins
		validCache, value, err := storage.GetCacheWithTTL(navigateCacheKey, time.Minute*5)
		if !validCache || err != nil {
			client := armclient.NewClientFromCLI("")
			if err != nil {
				cobra.CompErrorln("Failed creating armclient:" + err.Error())
				return []string{}, cobra.ShellCompDirectiveError
			}
			resourceList, err := getResourceListAndUpdateCache(allSubscriptionGUIDs, client)
			if err != nil {
				cobra.CompErrorln("Failed getting resource list:" + err.Error())
				return []string{}, cobra.ShellCompDirectiveError
			}
			value = resourceList
		}

		var graphQueryResult graphResponse
		err = json.Unmarshal([]byte(value), &graphQueryResult)
		if err != nil {
			cobra.CompErrorln("Failed to unmarshal graph response:" + err.Error())
			// Clear the cache as it can't be deserialized
			storage.DeleteCache(accountCacheKey) //nolint: errcheck
			return []string{}, cobra.ShellCompDirectiveError
		}

		// Used to limit completion suggestions to only the next segement deep
		toCompleteDepth := len(strings.Split(toComplete, "/"))
		isPartialDepth := false

		values := make([]string, 0, len(graphQueryResult.Data))
		for _, item := range graphQueryResult.Data {
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
		return nil, fmt.Errorf("Failed invoking az to update account list cache: %w", err)
	}

	var accounts []accountItem
	err = json.Unmarshal(out, &accounts)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshalling response from az to update account list cache: %w", err)
	}

	err = storage.PutCacheForTTL(accountCacheKey, string(out))
	if err != nil {
		return nil, fmt.Errorf("Failed to save account list to cache: %w", err)
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
		// Clear the cache as it can't be deserialized
		storage.DeleteCache(accountCacheKey) //nolint: errcheck
		return nil, fmt.Errorf("Failed unmarshalling from cache to get account list: %w", err)
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
		// Add the name and correctly escape spaces and quote the value
		values = append(values, strings.Replace(a.Name, " ", "\\ ", -1))
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

func createCompleteCommand(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion SHELL",
		Short: "Generates shell completion scripts",
		Long: `To load completion run
	
	. <(azbrowse completion SHELL)
	Valid values for SHELL are : bash, fish, powershell, zsh
	
	To configure your bash shell to load completions for each session add to your bashrc:
	
	# ~/.bashrc or ~/.profile
	source <(azbrowse completion bash)

	To configure completion for zsh run the following command:

	$ azbrowse completion zsh > "${fpath[1]}/_azbrowse"
	
	Ensure you have 'autoload -Uz compinit && compinit' present in your '.zshrc' file to load these completions

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
