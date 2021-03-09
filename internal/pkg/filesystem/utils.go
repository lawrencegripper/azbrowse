// +build !windows

package filesystem

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// Run starts a fuze based filesystem over the Azure ARM API
func Run(mountpoint string, filterToSub string, editMode bool, demoMode bool) (func(), error) {
	c, err := createFS(mountpoint, filterToSub, editMode, demoMode)
	if err != nil {
		log.Println("Failed to create fs")
		return func() {}, err
	}

	// Check if the mount process has an error to report.
	<-c.Ready
	closer := func() {
		Close(mountpoint, c) //nolint: errcheck
	}

	if err := c.MountError; err != nil {
		log.Println("Failed to mount fs")
		return closer, err
	}
	return closer, nil
}

func responseLogge(requestPath string, response *http.Response, responseBody string) {
	log.Println(requestPath)
	log.Println(responseBody)
}

func createFS(mountpoint string, filterToSub string, editMode bool, demoMode bool) (*fuse.Conn, error) {
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
	armclient.LegacyInstance = armClient

	expanders.InitializeExpanders(armClient, nil, nil, nil)
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
				if demoMode {
					message = views.StripSecretVals(status.Message)
				}
				fmt.Printf("%s STATUS: %s IN PROG: %t FAILED: %t \n", status.Icon(), message, status.InProgress, status.Failure)
			case <-timeout:
				// Update the UI
			}
		}
	}()

	srv := fs.New(c, nil)
	filesys := &FS{
		demoMode:             demoMode,
		editMode:             editMode,
		filterToSubscription: filterToSub,
	}
	go func() {
		if err := srv.Serve(filesys); err != nil {
			log.Println(err)
		}
	}()

	return c, nil
}

// Close unmounts the filesystem and waits for fs.Serve to return. Any
// returned error will be stored in Err. It is safe to call Close
// multiple times.
func Close(mount string, conn *fuse.Conn) error {
	prev := ""
	for tries := 0; tries < 1000; tries++ {
		err := fuse.Unmount(mount)
		if err != nil {
			msg := err.Error()
			// hide repeating errors
			if msg != prev {
				// TODO do more than log?

				// silence a very common message we can't do anything
				// about, for the first few tries. it'll still show if
				// the condition persists.
				if !strings.HasSuffix(err.Error(), ": Device or resource busy") || tries > 10 {
					log.Printf("unmount error: %v", err)
					prev = msg
				}
			}
			time.Sleep(100 * time.Millisecond)
			continue
		}
		break
	}
	return conn.Close()
}
