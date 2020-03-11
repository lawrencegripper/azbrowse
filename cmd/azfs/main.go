package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/filesystem"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func usage() {
	flag.PrintDefaults()
}

var demoMode *bool
var editMode *bool
var mountLocation *string

func run(mountpoint string) error {

	c, err := createFS(mountpoint)
	if err != nil {
		return err
	}

	// Check if the mount process has an error to report.
	<-c.Ready

	if err := c.MountError; err != nil {
		return err
	}
	return nil
}

func responseLogge(requestPath string, response *http.Response, responseBody string) {
	log.Println(requestPath)
	log.Println(responseBody)
}

func createFS(mountpoint string) (*fuse.Conn, error) {
	c, err := fuse.Mount(
		mountpoint,
		fuse.FSName("azfs"),
		fuse.Subtype("azfsfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("Azure ARM filesystem"),
	)

	if err != nil {
		return nil, err
	}

	if p := c.Protocol(); !p.HasInvalidate() {
		return nil, fmt.Errorf("kernel FUSE support is too old to have invalidations: version %v", p)
	}

	ctx := context.Background()

	// Load the db
	storage.LoadDB()

	// Create an ARMClient instance for us to use
	armClient := armclient.NewClientFromCLI("", responseLogge)
	armclient.LegacyInstance = *armClient

	expanders.InitializeExpanders(armClient)
	armClient.PopulateResourceAPILookup(ctx)

	// print status messages
	go func() {
		newEvents := eventing.SubscribeToStatusEvents()
		for {
			// Wait for a second to see if we have any new messages
			timeout := time.After(time.Second)
			select {
			case eventObj := <-newEvents:
				status := eventObj.(*eventing.StatusEvent)
				message := status.Message
				if *demoMode {
					message = views.StripSecretVals(status.Message)
				}
				fmt.Printf("%s STATUS: %s IN PROG: %t FAILED: %t \n", status.Icon(), message, status.InProgress, status.Failure)
			case <-timeout:
				// Update the UI
			}
		}
	}()

	srv := fs.New(c, nil)
	filesys := &filesystem.FS{}
	go func() {
		if err := srv.Serve(filesys); err != nil {
			log.Println(err)
		}
	}()

	return c, nil
}

func main() {
	flag.Usage = usage
	mountLocation = flag.String("mount", "/mnt/azfs", "defualt: /mnt/azfs location for mounting the filesystem")
	demoMode = flag.Bool("demo", false, "enable demo mode")
	editMode = flag.Bool("enableEdit", false, "enable delete/edit")

	flag.Parse()

	filesystem.EditMode = editMode
	filesystem.DemoMode = demoMode

	if err := run(*mountLocation); err != nil {
		log.Fatal(err)
	}

	stop := make(chan bool)
	<-stop
}
