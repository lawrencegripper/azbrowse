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

	"github.com/lawrencegripper/azbrowse/armclient"
	"github.com/lawrencegripper/azbrowse/search"
	"github.com/lawrencegripper/azbrowse/storage"
	"github.com/lawrencegripper/azbrowse/tracing"
	"github.com/lawrencegripper/azbrowse/version"

	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	opentracing "github.com/opentracing/opentracing-go"
	open "github.com/skratchdot/open-golang/open"
)

var enableTracing bool

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
			ToggleHelpView(g)
			storage.PutCache("firstrun", version.BuildDataVersion)
		}()
	}

	if maxX < 72 {
		panic("I can't run in a terminal less than 72 wide ... it's tooooo small!!!")
	}

	leftColumnWidth := 45

	status := NewStatusbarWidget(1, maxY-2, maxX, g)
	content := NewItemWidget(leftColumnWidth+2, 1, maxX-leftColumnWidth-1, maxY-4, "")
	list := NewListWidget(ctx, 1, 1, leftColumnWidth, maxY-4, []string{"Loading..."}, 0, content, status)

	g.SetManager(status, content, list)
	g.SetCurrentView("listWidget")

	if err := g.SetKeybinding("listWidget", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ChangeSelection(list.CurrentSelection() + 1)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ChangeSelection(list.CurrentSelection() - 1)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.ExpandCurrentSelection()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyF5, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.Refresh()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// Handle backspace for modern terminals
	if err := g.SetKeybinding("listWidget", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// When back is pressed in the itemWidget go back and also move
	// focus to the list of resources
	if err := g.SetKeybinding("itemWidget", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView("listWidget")
		g.Cursor = false
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// Handle backspace for out-dated terminals
	// A side-effect is that this key combination clashes with CTRL+H so we can't use that combination for help... oh well.
	// https://superuser.com/questions/375864/ctrlh-causing-backspace-instead-of-help-in-emacs-on-cygwin
	if err := g.SetKeybinding("listWidget", gocui.KeyBackspace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyCtrlA, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return LoadActionsView(ctx, list)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyCtrlO, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		item := list.CurrentItem()
		protalURL := os.Getenv("AZURE_PORTAL_URL")
		if protalURL == "" {
			protalURL = "https://portal.azure.com"
		}
		url := protalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.parentid + "/overview"
		span, _ := tracing.StartSpanFromContext(ctx, "openportal:url")
		open.Run(url)
		span.Finish()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	editModelEnabled := false
	if err := g.SetKeybinding("", gocui.KeyCtrlE, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		editModelEnabled = !editModelEnabled
		if editModelEnabled {
			g.Cursor = true
			g.SetCurrentView("itemWidget")
		} else {
			g.Cursor = false
			g.SetCurrentView("listWidget")
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("listWidget", gocui.KeyArrowRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		editModelEnabled = true
		g.Cursor = true
		g.SetCurrentView("itemWidget")
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("itemWidget", gocui.KeyArrowLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		editModelEnabled = false
		g.Cursor = false
		g.SetCurrentView("listWidget")
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	isFullScreen := false
	if err := g.SetKeybinding("", gocui.KeyCtrlF, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		isFullScreen = !isFullScreen
		if isFullScreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			v.Title = "JSON Response - Fullscreen (CTRL+F to exit)"
			fmt.Fprintf(v, content.GetContent())
			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(content.GetContent())
		status.Status("Current resource's JSON copied to clipboard", false)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	var showHelp bool
	if err := g.SetKeybinding("", gocui.KeyCtrlI, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		showHelp = !showHelp

		// If we're up and running clear and redraw the view
		// if w.g != nil {
		if showHelp {
			v, err := g.SetView("helppopup", 1, 1, 140, 32)
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			DrawHelp(v)
		} else {
			g.DeleteView("helppopup")
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// HACK: To prevent accidental deletes this method requires del to be pressed twice on a resource
	// before it will proceed
	var deleteConfirmItemID string
	var deleteConfirmCount int
	if err := g.SetKeybinding("", gocui.KeyDelete, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		item := list.CurrentItem()
		if deleteConfirmItemID != item.id {
			deleteConfirmItemID = item.id
			deleteConfirmCount = 0
		}
		status.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.deleteURL, true)
		deleteConfirmCount++

		if deleteConfirmCount > 1 {
			status.Status("Delete item? Really? PRESS DEL TO CONFIRM: "+item.deleteURL, true)

			res, err := armclient.DoRequest(ctx, "DELETE", item.deleteURL)
			if err != nil {
				panic(err)
			}
			content.SetContent(res, "Delete response>"+item.name)
			status.Status("Delete request sent successfully: "+item.deleteURL, false)

			deleteConfirmItemID = ""

			list.Refresh()

		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
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

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
