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
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/version"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/jroimartin/gocui"
	opentracing "github.com/opentracing/opentracing-go"
)

var enableTracing bool
var hideGuids bool

func main() {

	if len(os.Args) >= 2 {
		arg := os.Args[1]
		if strings.Contains(arg, "version") {
			fmt.Println(version.BuildDataVersion)
			fmt.Println(version.BuildDataGitCommit)
			fmt.Println(version.BuildDataGoVersion)
			fmt.Println(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
			fmt.Println(version.BuildDataBuildDate)
			os.Exit(0)
		}

		if strings.Contains(arg, "search") {
			fmt.Print("Getting resources \n")
			subRequest, _ := getSubscriptions(context.Background())
			search.CrawlResources(context.Background(), subRequest)
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
				bufio.NewReader(os.Stdin).ReadString('\n')
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

	// Show help if this is the first time the app has run
	firstRun, err := storage.GetCache("firstrun")
	if firstRun == "" || err != nil {
		go func() {
			time.Sleep(time.Second * 1)
			views.ToggleHelpView(g)
			storage.PutCache("firstrun", version.BuildDataVersion)
		}()
	}

	if maxX < 72 {
		panic("I can't run in a terminal less than 72 wide ... it's tooooo small!!!")
	}

	leftColumnWidth := 45

	status := views.NewStatusbarWidget(1, maxY-2, maxX, hideGuids, g)
	content := views.NewItemWidget(leftColumnWidth+2, 1, maxX-leftColumnWidth-1, maxY-4, hideGuids, "")
	list := views.NewListWidget(ctx, 1, 1, leftColumnWidth, maxY-4, []string{"Loading..."}, 0, content, status, enableTracing)

	g.SetManager(status, content, list)
	g.SetCurrentView("listWidget")

	var editModeEnabled bool
	var isFullscreen bool
	var showHelp bool
	var deleteConfirmItemID string
	var deleteConfirmCount int
	// List handlers
	keybindings.AddHandler(keybindings.ListDownHandler{List: list})
	keybindings.AddHandler(keybindings.ListUpHandler{List: list})
	keybindings.AddHandler(keybindings.ListExpandHandler{List: list})
	keybindings.AddHandler(keybindings.ListRefreshHandler{List: list})
	keybindings.AddHandler(keybindings.ListBackHandler{List: list})
	keybindings.AddHandler(keybindings.ListBackLegacyHandler{List: list})
	keybindings.AddHandler(keybindings.ListActionsHandler{List: list, Context: ctx})
	keybindings.AddHandler(keybindings.ListRightHandler{List: list, EditModeEnabled: &editModeEnabled})
	keybindings.AddHandler(keybindings.ListEditHandler{List: list, EditModeEnabled: &editModeEnabled})
	keybindings.AddHandler(keybindings.ListOpenHandler{List: list, Context: ctx})

	// Item handlers
	keybindings.AddHandler(keybindings.ItemBackHandler{List: list})
	keybindings.AddHandler(keybindings.ItemLeftHandler{EditModeEnabled: &editModeEnabled})

	// Global handlers
	keybindings.AddHandler(keybindings.FullscreenHandler{List: list, IsFullscreen: &isFullscreen, Content: content})
	keybindings.AddHandler(keybindings.CopyHandler{StatusBar: status, Content: content})
	keybindings.AddHandler(keybindings.HelpHandler{ShowHelp: &showHelp})
	keybindings.AddHandler(keybindings.DeleteHandler{List: list, StatusBar: status, Content: content, Context: ctx, DeleteConfirmItemID: deleteConfirmItemID, DeleteConfirmCount: deleteConfirmCount})
	keybindings.AddHandler(keybindings.SaveHandler{StatusBar: status, Content: content})
	keybindings.AddHandler(keybindings.QuitHandler{})

	if err := keybindings.Bind(g); err != nil { // apply late binding for keys
		log.Panicln(err)
	}

	// Update views with keybindings
	keyBindings := keybindings.GetKeyBindingsAsStrings()
	status.HelpKeyBinding = keyBindings["help"]
	list.ActionKeyBinding = keyBindings["listactions"]
	list.FullscreenKeyBinding = keyBindings["fullscreen"]

	go func() {
		time.Sleep(time.Second * 1)

		status.Status("Fetching Subscriptions", true)
		subRequest, data := getSubscriptions(ctx)

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

func getSubscriptions(ctx context.Context) (armclient.SubResponse, string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "expand:subs")
	defer span.Finish()

	// Get Subscriptions
	data, err := armclient.DoRequest(ctx, "GET", "/subscriptions?api-version=2018-01-01")
	if err != nil {
		panic(err)
	}

	var subRequest armclient.SubResponse
	err = json.Unmarshal([]byte(data), &subRequest)
	if err != nil {
		panic(err)
	}
	return subRequest, data
}
