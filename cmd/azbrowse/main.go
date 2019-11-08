package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/keybindings"
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

var enableTracing bool
var hideGuids bool
var navigateToID string

func main() {
	readCommandLineArgs()

	confirmAndSelfUpdate()

	// Setup the root context and span for open tracing
	ctx, span := configureTracing()

	// Create an ARMClient instance for us to use
	armClient := armclient.NewClientFromCLI()

	// Initialize the expanders which will let the user walk the tree of
	// resources in Azure
	expanders.InitializeExpanders(armClient)

	// Start up gocui and configure some settings
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan
	g.InputEsc = true

	// Create the views we'll use to display information and
	// bind up all the keys use to interact with the views
	list := setupViewsAndKeybindings(ctx, g, armClient)

	// Start a go routine to populate the list with root of the nodes
	startPopulatingList(ctx, g, list, armClient)

	// Start a go routine to handling automated naviging to an item via the
	// `--navigate` command
	handleNavigateTo(list)

	// Close the span used to track startup times
	span.Finish()

	automatedFuzzer(list)

	// Start the main loop of gocui to draw the UI
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func readCommandLineArgs() {
	args := os.Args[1:] // skip first arg

	for len(args) >= 1 {
		arg := args[0]
		handled := false

		if strings.Contains(arg, "version") {
			fmt.Println(version)
			fmt.Println(commit)
			fmt.Println(date)
			fmt.Println(goversion)
			fmt.Println(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
			os.Exit(0)
		}

		if strings.Contains(arg, "debug") {
			enableTracing = true
			tracing.EnableDebug()
			config.SetDebuggingEnabled(true)
			handled = true
		}
		if strings.Contains(arg, "demo") {
			hideGuids = true
			handled = true
		}

		if strings.Contains(arg, "navigate") {
			if len(args) >= 2 {
				navigateToID = args[1] // capture the next arg
				args = args[1:]        // move past the the captured arg
				handled = true
			}
		}

		if !handled {
			// unhandled arg
			fmt.Println("Usage:")
			fmt.Println("\tazbrowse [debug] [demo] [navigate <id to navigate to>]")
			fmt.Println()
			os.Exit(-1)
		}

		args = args[1:] // move to the next arg
	}
}

func configureTracing() (context.Context, opentracing.Span) {
	var ctx context.Context
	var span opentracing.Span

	if enableTracing {
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

func setupViewsAndKeybindings(ctx context.Context, g *gocui.Gui, client *armclient.Client) *views.ListWidget {
	maxX, maxY := g.Size()
	// Padding
	maxX = maxX - 2
	maxY = maxY - 2

	if maxX < 72 {
		panic("I can't run in a terminal less than 72 wide ... it's tooooo small!!!")
	}

	leftColumnWidth := 45

	// Create the views used
	status := views.NewStatusbarWidget(1, maxY-2, maxX, hideGuids, g)
	content := views.NewItemWidget(leftColumnWidth+2, 1, maxX-leftColumnWidth-1, maxY-4, hideGuids, "")
	list := views.NewListWidget(ctx, 1, 1, leftColumnWidth, maxY-4, []string{"Loading..."}, 0, content, status, enableTracing, "Subscriptions", g)
	notifications := views.NewNotificationWidget(maxX-45, 1, 45, hideGuids, g, client)
	commandPanel := views.NewCommandPanelWidget(leftColumnWidth+8, 5, maxX-leftColumnWidth-20, g)

	g.SetManager(status, content, list, notifications, commandPanel)
	g.SetCurrentView("listWidget")

	var editModeEnabled bool
	var isFullscreen bool
	var showHelp bool

	// Global handlers
	// NOTE> Global handlers must be registered first to
	//       ensure double key registration is prevented
	keybindings.AddHandler(keybindings.NewFullscreenHandler(list, content, &isFullscreen))
	keybindings.AddHandler(keybindings.NewCopyHandler(content, status))
	keybindings.AddHandler(keybindings.NewHelpHandler(&showHelp))
	keybindings.AddHandler(keybindings.NewQuitHandler())
	keybindings.AddHandler(keybindings.NewConfirmDeleteHandler(notifications))
	keybindings.AddHandler(keybindings.NewClearPendingDeleteHandler(notifications))
	keybindings.AddHandler(keybindings.NewOpenCommandPanelHandler(commandPanel))
	keybindings.AddHandler(keybindings.NewCommandPanelFilterHandler(commandPanel))
	keybindings.AddHandler(keybindings.NewCloseCommandPanelHandler(commandPanel))

	// List handlers
	keybindings.AddHandler(keybindings.NewListDownHandler(list))
	keybindings.AddHandler(keybindings.NewListUpHandler(list))
	keybindings.AddHandler(keybindings.NewListExpandHandler(list))
	keybindings.AddHandler(keybindings.NewListRefreshHandler(list))
	keybindings.AddHandler(keybindings.NewListBackHandler(list))
	keybindings.AddHandler(keybindings.NewListBackLegacyHandler(list))
	keybindings.AddHandler(keybindings.NewListActionsHandler(list, ctx))
	keybindings.AddHandler(keybindings.NewListRightHandler(list, &editModeEnabled))
	keybindings.AddHandler(keybindings.NewListEditHandler(list, &editModeEnabled))
	keybindings.AddHandler(keybindings.NewListOpenHandler(list, ctx))
	keybindings.AddHandler(keybindings.NewListDeleteHandler(list, notifications))
	keybindings.AddHandler(keybindings.NewListUpdateHandler(list, status, ctx, content, g))
	keybindings.AddHandler(keybindings.NewListPageDownHandler(list))
	keybindings.AddHandler(keybindings.NewListPageUpHandler(list))
	keybindings.AddHandler(keybindings.NewListEndHandler(list))
	keybindings.AddHandler(keybindings.NewListHomeHandler(list))
	keybindings.AddHandler(keybindings.NewListClearFilterHandler(list))

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

func automatedFuzzer(list *views.ListWidget) {
	go func() {
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")
		depthCount := 0
		for {
			nodeListInterface := <-navigatedChannel
			nodeList := nodeListInterface.([]*expanders.TreeNode)
			if len(nodeList) < 1 {
				for index := 0; index < depthCount; index++ {
					list.GoBack()
				}
				continue
			}

			if len(nodeList) > 0 {
				list.ChangeSelection(rand.Intn(len(nodeList)))
				list.ExpandCurrentSelection()
				depthCount++
			}
		}
	}()
}

func handleNavigateTo(list *views.ListWidget) {
	if navigateToID != "" {
		navigateToIDLower := strings.ToLower(navigateToID)
		go func() {
			navigatedChannel := eventing.SubscribeToTopic("list.navigated")
			var lastNavigatedNode *expanders.TreeNode

			processNavigations := true

			for {
				nodeListInterface := <-navigatedChannel

				if processNavigations {
					nodeList := nodeListInterface.([]*expanders.TreeNode)

					if lastNavigatedNode != nil && lastNavigatedNode != list.CurrentExpandedItem() {
						processNavigations = false
					} else {

						gotNode := false
						for nodeIndex, node := range nodeList {
							// use prefix matching
							// but need additional checks as target of /foo/bar would be matched by  /foo/bar  and /foo/ba
							// additional check is that the lengths match, or the next char in target is a '/'
							nodeIDLower := strings.ToLower(node.ID)
							if strings.HasPrefix(navigateToIDLower, nodeIDLower) && (len(navigateToID) == len(nodeIDLower) || navigateToIDLower[len(nodeIDLower)] == '/') {
								list.ChangeSelection(nodeIndex)
								lastNavigatedNode = node
								list.ExpandCurrentSelection()
								gotNode = true
								break
							}
						}

						if !gotNode {
							// we got as far as we could - now stop!
							processNavigations = false
						}
					}
				}
			}
		}()
	}
}
