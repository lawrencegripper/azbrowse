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
	"github.com/lawrencegripper/azbrowse/style"
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
				fmt.Printf("To see the trace details for the session visit: %s. \n Press any key to exit when you are done. \n", startTraceDashboardForSpan(span))
				bufio.NewReader(os.Stdin).ReadString('\n')
			}
		}()
	} else {

		rootCtx := context.Background()
		span, ctx = tracing.StartSpanFromContext(rootCtx, "azbrowseStart")
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue

	// help := NewHelpWidget("help", 1, 1, helpText)
	maxX, maxY := g.Size()
	// Padding
	maxX = maxX - 2
	maxY = maxY - 2

	if maxX < 72 {
		panic("I can't run in a terminal less than 72 wide ... it's tooooo small!!!")
	}

	status := NewStatusbarWidget(1, maxY-2, maxX, g)
	header := NewHeaderWidget(1, 1, 70, 9)
	content := NewItemWidget(70+2, 1, maxX-70-1, maxY-4, "")
	list := NewListWidget(ctx, 1, 11, 70, maxY-14, []string{"Loading..."}, 0, content, status)

	g.SetManager(status, content, list, header)
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

	if err := g.SetKeybinding("listWidget", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		list.GoBack()
		return nil
	}); err != nil {
		log.Panicln(err)
	}

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
		open.Run(protalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.parentid + "/overview")
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
			content.SetContent(style.Title("Delete response for item:"+item.deleteURL+"\n ------------------------------- \n") + res)
			status.Status("Delete request sent successfully: "+item.deleteURL, false)

			deleteConfirmItemID = ""

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
				content.SetContent(err.Error())
				return nil
			}
			content.SetContent(data)
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
