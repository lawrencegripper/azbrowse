package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/keybindings"
	"github.com/lawrencegripper/azbrowse/internal/pkg/search"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/jroimartin/gocui"
	opentracing "github.com/opentracing/opentracing-go"
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

func main() {

	if len(os.Args) >= 2 {
		arg := os.Args[1]
		if strings.Contains(arg, "version") {
			fmt.Println(version)
			fmt.Println(commit)
			fmt.Println(date)
			fmt.Println(goversion)
			fmt.Println(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
			os.Exit(0)
		}

		if strings.Contains(arg, "search") {
			fmt.Print("Getting resources \n")
			subRequest, _, err := getSubscriptions(context.Background())
			if err != nil {
				panic(err)
			}
			err = search.CrawlResources(context.Background(), subRequest)
			if err != nil {
				panic(err)
			}
			fmt.Print("Build suggester \n")

			suggester, _ := search.NewSuggester()
			fmt.Print("Get suggestions \n")

			suggestions := suggester.Autocomplete(os.Args[2])
			fmt.Printf("%v \n", suggestions)
			os.Exit(0)
		}

		if strings.Contains(arg, "debug") {
			enableTracing = true
			tracing.EnableDebug()
		}

		if strings.Contains(arg, "demo") {
			hideGuids = true
		}
	}

	confirmAndSelfUpdate()
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

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan

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
	list := views.NewListWidget(ctx, 1, 1, leftColumnWidth, maxY-4, []string{"Loading..."}, 0, content, status, enableTracing)
	notifications := views.NewNotificationWidget(maxX-45, 1, 45, hideGuids, g)
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
	keybindings.AddHandler(keybindings.NewListUpdateHandler(list, status, ctx, content))
	keybindings.AddHandler(keybindings.NewListPageDownHandler(list))
	keybindings.AddHandler(keybindings.NewListPageUpHandler(list))
	keybindings.AddHandler(keybindings.NewListEndHandler(list))
	keybindings.AddHandler(keybindings.NewListHomeHandler(list))
	keybindings.AddHandler(keybindings.NewItemViewPageDownHandler(content))
	keybindings.AddHandler(keybindings.NewItemViewPageUpHandler(content))

	// Item handlers
	keybindings.AddHandler(keybindings.NewItemBackHandler(list))
	keybindings.AddHandler(keybindings.NewItemLeftHandler(&editModeEnabled))

	if err = keybindings.Bind(g); err != nil { // apply late binding for keys
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

	go func() {
		time.Sleep(time.Second * 1)

		status.Status("Fetching Subscriptions", true)
		subRequest, data, err := getSubscriptions(ctx)
		if err != nil {
			g.Close()
			log.Panicln(err)
		}

		g.Update(func(gui *gocui.Gui) error {
			g.SetCurrentView("listWidget")

			status.Status("Getting provider data", true)

			armclient.PopulateResourceAPILookup(ctx)
			status.Status("Done getting provider data", false)

			list.SetSubscriptions(subRequest)

			if err != nil {
				content.SetContent(err.Error(), "Error")
				return nil
			}
			content.SetContent(data, "Subscriptions response")
			return nil
		})

		status.Status("Fetching Subscriptions: Completed", false)

	}()

	span.Finish()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func getSubscriptions(ctx context.Context) (armclient.SubResponse, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "expand:subs")
	defer span.Finish()

	// Get Subscriptions
	data, err := armclient.DoRequest(ctx, "GET", "/subscriptions?api-version=2018-01-01")
	if err != nil {
		return armclient.SubResponse{}, "", fmt.Errorf("Failed to load subscriptions: %s", err)
	}

	var subRequest armclient.SubResponse
	err = json.Unmarshal([]byte(data), &subRequest)
	if err != nil {
		return armclient.SubResponse{}, "", fmt.Errorf("Failed to load subscriptions: %s", err)
	}
	return subRequest, data, nil
}
