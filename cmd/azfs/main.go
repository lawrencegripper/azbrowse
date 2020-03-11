package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
	"bazil.org/fuse/fuseutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

var ctx = context.Background()
var demoMode *bool
var editMode *bool
var mountLocation *string

func usage() {
	flag.PrintDefaults()
}

func run(mountpoint string) error {
	c, err := fuse.Mount(
		mountpoint,
		fuse.FSName("azfs"),
		fuse.Subtype("azfsfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("Azure ARM filesystem"),
	)
	if err != nil {
		return err
	}
	defer c.Close() //nolint: errcheck

	if p := c.Protocol(); !p.HasInvalidate() {
		return fmt.Errorf("kernel FUSE support is too old to have invalidations: version %v", p)
	}

	// Load the db
	storage.LoadDB()

	// Start tracking async responses from ARM
	responseProcessor, err := views.StartWatchingAsyncARMRequests(ctx)
	if err != nil {
		log.Panicln(err)
	}

	// Create an ARMClient instance for us to use
	armClient := armclient.NewClientFromCLI("", responseProcessor)
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
	filesys := &FS{}

	if err := srv.Serve(filesys); err != nil {
		return err
	}

	// Check if the mount process has an error to report.
	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Usage = usage
	mountLocation = flag.String("mount", "/mnt/azfs", "defualt: /mnt/azfs location for mounting the filesystem")
	demoMode = flag.Bool("demo", false, "enable demo mode")
	editMode = flag.Bool("enableEdit", false, "enable delete/edit")

	flag.Parse()

	if err := run(*mountLocation); err != nil {
		log.Fatal(err)
	}
}

// FS Structure
type FS struct {
}

var _ fs.FS = (*FS)(nil)

func (f *FS) Root() (fs.Node, error) {
	// Create an empty tentant TreeNode. This by default expands
	// to show the current tenants subscriptions
	_, newItems, err := expanders.ExpandItem(ctx, &expanders.TreeNode{
		ItemType:  expanders.TentantItemType,
		ID:        "AvailableSubscriptions",
		ExpandURL: expanders.ExpandURLNotSupported,
	})
	if err != nil {
		panic(err)
	}

	return &RootDir{
		fs:    f,
		items: newItems,
	}, nil
}

// RootDir implements both Node and Handle for the root directory.
type RootDir struct {
	fs    *FS
	items []*expanders.TreeNode
}

var _ fs.Node = (*RootDir)(nil)

func (d *RootDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555
	return nil
}

var _ fs.NodeStringLookuper = (*RootDir)(nil)

func nameFromTreeNode(treeNode *expanders.TreeNode) string {
	// fmt.Printf("%+v", *treeNode)
	return treeNode.Name
}

func canEdit(item *expanders.TreeNode) bool {
	if item == nil ||
		item.SwaggerResourceType == nil ||
		item.SwaggerResourceType.PutEndpoint == nil ||
		item.Metadata == nil ||
		item.Metadata["SwaggerAPISetID"] == "" {
		return false
	}
	return true
}

func indexFileName(treeNode *expanders.TreeNode, response *expanders.ExpanderResponse) string {
	if treeNode == nil || response == nil {
		return "index.json"
	}
	return treeNode.ItemType + "." + treeNode.Name + "." + strings.ToLower(string(response.ResponseType))
}

func (d *RootDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	for i, treeNode := range d.items {
		if nameFromTreeNode(treeNode) == name {
			f := &Folder{
				treeNode: treeNode,
				Dirent: fuse.Dirent{
					Inode: uint64(i),
					Type:  fuse.DT_Dir,
					Name:  nameFromTreeNode(treeNode),
				},
			}
			return f, nil
		}
	}
	return nil, syscall.ENOENT
}

var _ fs.HandleReadDirAller = (*RootDir)(nil)

func (d *RootDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	dirItems := make([]fuse.Dirent, len(d.items))
	for i, treeNode := range d.items {
		dirItems[i] = fuse.Dirent{
			Inode: uint64(i),
			Type:  fuse.DT_Dir,
			Name:  nameFromTreeNode(treeNode),
		}
	}
	return dirItems, nil
}

type Folder struct {
	fuse.Dirent
	treeNode     *expanders.TreeNode
	items        []*expanders.TreeNode
	indexContent *expanders.ExpanderResponse
	canDelete    bool
}

var _ fs.Node = (*Folder)(nil)

func (d *Folder) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	if d.canDelete {
		a.Mode = os.ModeDir | 0700
	} else {
		a.Mode = os.ModeDir | 0555
	}
	return nil
}

var _ fs.NodeStringLookuper = (*Folder)(nil)

func (d *Folder) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == indexFileName(d.treeNode, d.indexContent) {
		file := &File{}
		content := d.indexContent.Response

		if d.indexContent.ResponseType == "JSON" {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(d.indexContent.Response), "", "   ")
			if err != nil {
				panic(err)
			}
			content = prettyJSON.String()
		}

		file.content.Store(content)
		return file, nil
	}
	for _, treeNode := range d.items {
		if nameFromTreeNode(treeNode) == name {
			f := &Folder{
				treeNode:  treeNode,
				canDelete: treeNode.DeleteURL != "",
				Dirent: fuse.Dirent{
					// Inode: uint64(i),
					Type: fuse.DT_Dir,
					Name: nameFromTreeNode(treeNode),
				},
			}
			return f, nil
		}
	}
	return nil, syscall.ENOENT
}

var _ fs.NodeRemover = (*Folder)(nil)

// Todo currently `rm -r thing` gives an error deleting an RG but the delete is processed. Error:
//
// rm: WARNING: Circular directory structure.
// This almost certainly means that you have a corrupted file system.
// NOTIFY YOUR SYSTEM MANAGER.
// The following directory is part of the cycle:
// 30466/Deployments
// rm: cannot remove '30466': Input/output error
func (d *Folder) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	for _, treeNode := range d.items {
		if nameFromTreeNode(treeNode) == req.Name {
			// if d.treeNode.ItemType == "subscription" {
			// 	return fmt.Errorf("Can't delete subs, noop: %+v", req)
			// }

			log.Printf("Found matching node '%s' doing delete: %+v", treeNode.Name+treeNode.ItemType, req)
			fallback := true
			if treeNode.Expander != nil {
				deleted, err := treeNode.Expander.Delete(ctx, treeNode)
				fallback = (err == nil && !deleted)
			}
			if fallback {
				// fallback to ARM request to delete
				_, err := armclient.LegacyInstance.DoRequest(ctx, "DELETE", treeNode.DeleteURL)
				panic(err)
			}
			log.Printf("Delete complete: %+v", req)
		}
	}

	return nil
}

var _ fs.HandleReadDirAller = (*Folder)(nil)

func (d *Folder) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	rootContent, newItems, err := expanders.ExpandItem(ctx, d.treeNode)
	if err != nil {
		panic(err)
	}
	d.items = newItems
	d.indexContent = rootContent

	if *demoMode {
		rootContent.Response = views.StripSecretVals(rootContent.Response)
	}

	dirItems := make([]fuse.Dirent, len(d.items)+1)
	for i, treeNode := range d.items {
		dirItems[i] = fuse.Dirent{
			// Inode: uint64(i),
			Type: fuse.DT_Dir,
			Name: nameFromTreeNode(treeNode),
		}
	}
	lastIndex := len(d.items)
	dirItems[lastIndex] = fuse.Dirent{
		// Inode: uint64(lastIndex),
		Type: fuse.DT_File,
		Name: indexFileName(d.treeNode, d.indexContent),
	}
	return dirItems, nil
}

type File struct {
	fs.NodeRef

	// fuse     *fs.Server
	content atomic.Value
	// treeNode *expanders.TreeNode
}

var _ fs.Node = (*File)(nil)

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 2
	a.Mode = 0444
	t, hasContent := f.content.Load().(string)
	if hasContent {
		a.Size = uint64(len(t))
	} else {
		a.Size = uint64(0)
	}
	return nil
}

var _ fs.NodeOpener = (*File)(nil)

func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	if !req.Flags.IsReadOnly() {
		return nil, fuse.Errno(syscall.EACCES)
	}
	resp.Flags |= fuse.OpenKeepCache
	return f, nil
}

var _ fs.Handle = (*File)(nil)

var _ fs.HandleReader = (*File)(nil)

func (f *File) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	t := f.content.Load().(string)
	fuseutil.HandleRead(req, resp, []byte(t))
	return nil
}

var _ fs.NodeRemover = (*File)(nil)

func (d *File) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	log.Println(fmt.Errorf("Can't delete files as they're imaginary, noop: %+v", req))
	return nil
}
