package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/automation"
	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/keybindings"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/stuartleeks/gocui"
)

// Overridden via ldflags
var (
	version   = "99.0.1-devbuild"
	commit    = "unknown"
	date      = "unknown"
	goversion = "unknown"
)

func main() {
	handleCommandAndArgs()
}

func run(settings *config.Settings) {
	confirmAndSelfUpdate()

	// Setup the root context and span for open tracing
	ctx, span := configureTracing(settings)

	// Load the db
	storage.LoadDB()

	// Start tracking async responses from ARM
	responseProcessor, err := views.StartWatchingAsyncARMRequests(ctx)
	if err != nil {
		log.Panicln(err)
	}

	// Create an ARMClient instance for us to use
	armClient := armclient.NewClientFromCLI(settings.TenantID, responseProcessor)
	armclient.LegacyInstance = *armClient

	// Initialize the expanders which will let the user walk the tree of
	// resources in Azure
	expanders.InitializeExpanders(armClient)

	// Start up gocui and configure some settings
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	// Give error handling the Gui instance so it can cleanup
	// when a panic occurs
	errorhandling.RegisterGuiAndStartHistoryTracking(ctx, g)

	// recover from normal exit of the program
	defer g.Close()

	// recover from panic, if one occurrs, and leave terminal usable
	defer errorhandling.RecoveryWithCleanup()

	// Configure the gui instance
	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan
	g.InputEsc = true

	// Create the views we'll use to display information and
	// bind up all the keys use to interact with the views
	list := setupViewsAndKeybindings(ctx, g, settings, armClient)

	// Start a go routine to populate the list with root of the nodes
	startPopulatingList(ctx, g, list, armClient)

	// Start a go routine to handling automated naviging to an item via the
	// `--navigate` command
	if settings.NavigateToID != "" {
		automation.NavigateTo(list, settings.NavigateToID)
	}

	if settings.FuzzerEnabled {
		automation.StartAutomatedFuzzer(list, settings, g)
	}

	// Close the span used to track startup times
	span.Finish()

	// Start the main loop of gocui to draw the UI
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func configureTracing(settings *config.Settings) (context.Context, opentracing.Span) {
	var ctx context.Context
	var span opentracing.Span

	if settings.EnableTracing {
		startTraceDashboardForSpan := tracing.StartTracing()

		rootCtx := context.Background()
		span, ctx = tracing.StartSpanFromContext(rootCtx, "azbrowseStart")

		startTraceDashboardForSpan(span)

		defer func() {
			// recover from panic if one occurred and show user the trace URL for debugging.
			if r := recover(); r != nil {
				fmt.Printf("A crash occurred: %s", r)
				debug.PrintStack()
				fmt.Printf("To see the trace details for the session visit: %s. \n Visit https://github.com/lawrencegripper/azbrowse/issues to raise a bug. \n Press any key to exit when you are done. \n", startTraceDashboardForSpan(span))

				bufio.NewReader(os.Stdin).ReadString('\n') //nolint:golint,errcheck
			}
		}()
	} else {

		rootCtx := context.Background()
		span, ctx = tracing.StartSpanFromContext(rootCtx, "azbrowseStart")

		defer func() {
			// recover from panic if one occurred and explain to the user how to proceed
			if r := recover(); r != nil {
				fmt.Printf("A crash occurred: %s", r)
				debug.PrintStack()
				fmt.Printf("To capture a more detailed train run 'azbrowse --debug' and reproduce the issue. Visit https://github.com/lawrencegripper/azbrowse/issues to raise a bug.")
			}
		}()
	}

	return ctx, span
}

func startPopulatingList(ctx context.Context, g *gocui.Gui, list *views.ListWidget, armClient *armclient.Client) {
	go func() {
		defer errorhandling.RecoveryWithCleanup()

		time.Sleep(time.Second * 1)

		_, done := eventing.SendStatusEvent(eventing.StatusEvent{
			Message:    "Updating API Version details",
			InProgress: true,
		})

		armClient.PopulateResourceAPILookup(ctx)

		done()

		g.Update(func(gui *gocui.Gui) error {
			g.SetCurrentView("listWidget")

			// Create an empty tentant TreeNode. This by default expands
			// to show the current tenants subscriptions
			newContent, newItems, err := expanders.ExpandItem(ctx, &expanders.TreeNode{
				ItemType:  expanders.TentantItemType,
				ID:        "AvailableSubscriptions",
				ExpandURL: expanders.ExpandURLNotSupported,
			})

			if err != nil {
				panic(err)
			}

			list.Navigate(newItems, newContent, "Subscriptions")

			return nil
		})
	}()
}

func setupViewsAndKeybindings(ctx context.Context, g *gocui.Gui, settings *config.Settings, client *armclient.Client) *views.ListWidget {
	maxX, maxY := g.Size()
	// Padding
	maxX = maxX - 2
	maxY = maxY - 2

	if maxX < 72 {
		panic("I can't run in a terminal less than 72 wide ... it's tooooo small!!!")
	}

	leftColumnWidth := 45

	// Create the views used
	status := views.NewStatusbarWidget(1, maxY-2, maxX, settings.HideGuids, g)
	content := views.NewItemWidget(leftColumnWidth+2, 1, maxX-leftColumnWidth-1, maxY-4, settings.HideGuids, settings.ShouldRender, "")
	list := views.NewListWidget(ctx, 1, 1, leftColumnWidth, maxY-4, []string{"Loading..."}, 0, content, status, settings.EnableTracing, "Subscriptions", settings.ShouldRender, g)
	notifications := views.NewNotificationWidget(maxX-45, 1, 45, g, client)

	commandPanel := views.NewCommandPanelWidget(leftColumnWidth+3, 0, maxX-leftColumnWidth-20, g)

	copyCommand := keybindings.NewCopyHandler(content, status)
	toggleDemoModeCommand := keybindings.NewToggleDemoModeHandler(settings, list, status, content)

	commandPanelFilterCommand := keybindings.NewCommandPanelFilterHandler(commandPanel, list)
	commandPanelAzureSearchQueryCommand := keybindings.NewCommandPanelAzureSearchQueryHandler(commandPanel, content, list)

	listActionsCommand := keybindings.NewListActionsHandler(list, ctx)
	listOpenCommand := keybindings.NewListOpenHandler(list, ctx)
	listUpdateCommand := keybindings.NewListUpdateHandler(list, status, ctx, content, g)
	listCopyItemIDCommand := keybindings.NewListCopyItemIDHandler(list, status)
	listDebugCopyItemDataCommand := keybindings.NewListDebugCopyItemDataHandler(list, status)

	commands := []keybindings.Command{
		commandPanelFilterCommand,
		copyCommand,
		commandPanelAzureSearchQueryCommand,
		listActionsCommand,
		listOpenCommand,
		listUpdateCommand,
		listCopyItemIDCommand,
		toggleDemoModeCommand,
	}
	if settings.EnableTracing {
		commands = append(commands, listDebugCopyItemDataCommand)
	}
	sort.Sort(keybindings.SortByDisplayText(commands))

	g.SetManager(status, content, list, notifications, commandPanel)
	g.SetCurrentView("listWidget")

	var editModeEnabled bool
	var isFullscreen bool
	var showHelp bool

	// Global handlers
	// NOTE> Global handlers must be registered first to
	//       ensure double key registration is prevented
	keybindings.AddHandler(keybindings.NewFullscreenHandler(list, content, &isFullscreen))
	keybindings.AddHandler(copyCommand)
	keybindings.AddHandler(keybindings.NewHelpHandler(&showHelp))
	keybindings.AddHandler(keybindings.NewQuitHandler())
	keybindings.AddHandler(keybindings.NewConfirmDeleteHandler(notifications))
	keybindings.AddHandler(keybindings.NewClearPendingDeleteHandler(notifications))
	keybindings.AddHandler(keybindings.NewOpenCommandPanelHandler(g, commandPanel, commands))
	keybindings.AddHandler(commandPanelFilterCommand)
	keybindings.AddHandler(keybindings.NewCloseCommandPanelHandler(commandPanel))
	keybindings.AddHandler(keybindings.NewCommandPanelDownHandler(commandPanel))
	keybindings.AddHandler(keybindings.NewCommandPanelUpHandler(commandPanel))
	keybindings.AddHandler(keybindings.NewCommandPanelEnterHandler(commandPanel))
	keybindings.AddHandler(toggleDemoModeCommand)

	// List handlers
	keybindings.AddHandler(keybindings.NewListDownHandler(list))
	keybindings.AddHandler(keybindings.NewListUpHandler(list))
	keybindings.AddHandler(keybindings.NewListExpandHandler(list))
	keybindings.AddHandler(keybindings.NewListRefreshHandler(list))
	keybindings.AddHandler(keybindings.NewListBackHandler(list))
	keybindings.AddHandler(keybindings.NewListBackLegacyHandler(list))
	keybindings.AddHandler(listActionsCommand)
	keybindings.AddHandler(keybindings.NewListRightHandler(list, &editModeEnabled))
	keybindings.AddHandler(keybindings.NewListEditHandler(list, &editModeEnabled))
	keybindings.AddHandler(listOpenCommand)
	keybindings.AddHandler(keybindings.NewListDeleteHandler(list, notifications))
	keybindings.AddHandler(listUpdateCommand)
	keybindings.AddHandler(keybindings.NewListPageDownHandler(list))
	keybindings.AddHandler(keybindings.NewListPageUpHandler(list))
	keybindings.AddHandler(keybindings.NewListEndHandler(list))
	keybindings.AddHandler(keybindings.NewListHomeHandler(list))
	keybindings.AddHandler(keybindings.NewListClearFilterHandler(list))
	keybindings.AddHandler(commandPanelAzureSearchQueryCommand)
	keybindings.AddHandler(listCopyItemIDCommand)
	if settings.EnableTracing {
		keybindings.AddHandler(listDebugCopyItemDataCommand)
	}

	// ItemView handlers
	keybindings.AddHandler(keybindings.NewItemViewPageDownHandler(content))
	keybindings.AddHandler(keybindings.NewItemViewPageUpHandler(content))

	// Item handlers
	keybindings.AddHandler(keybindings.NewItemBackHandler(list))
	keybindings.AddHandler(keybindings.NewItemLeftHandler(&editModeEnabled))

	if err := keybindings.Bind(g); err != nil { // apply late binding for keys
		g.Close()

		fmt.Println("\n \n" + err.Error())
		os.Exit(1)
	}

	// Update views with keybindings
	keyBindings := keybindings.GetKeyBindingsAsStrings()
	status.HelpKeyBinding = strings.Join(keyBindings["help"], ",")
	list.ActionKeyBinding = strings.Join(keyBindings["listactions"], ",")
	list.FullscreenKeyBinding = strings.Join(keyBindings["fullscreen"], ",")
	notifications.ConfirmDeleteKeyBinding = strings.Join(keyBindings["confirmdelete"], ",")
	notifications.ClearPendingDeletesKeyBinding = strings.Join(keyBindings["clearpendingdeletes"], ",")

	return list
}
