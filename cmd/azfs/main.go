package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
	"bazil.org/fuse/fuseutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

var ctx = context.Background()

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s MOUNTPOINT\n", os.Args[0])
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
	defer c.Close()

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
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	mountpoint := flag.Arg(0)

	if err := run(mountpoint); err != nil {
		log.Fatal(err)
	}
}

// FS Structure
type FS struct {
	stateMap map[string]*expanders.TreeNode
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
	indexContent string
}

var _ fs.Node = (*Folder)(nil)

func (d *Folder) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555
	return nil
}

var _ fs.NodeStringLookuper = (*Folder)(nil)

func (d *Folder) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == "index.json" {
		file := &File{}
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(views.StripSecretVals(d.indexContent)), "", "   ")
		if err != nil {
			panic(err)
		}
		file.content.Store(string(prettyJSON.Bytes()))
		return file, nil
	}
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

var _ fs.HandleReadDirAller = (*Folder)(nil)

func (d *Folder) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	rootContent, newItems, err := expanders.ExpandItem(ctx, d.treeNode)
	if err != nil {
		panic(err)
	}
	d.items = newItems
	d.indexContent = rootContent.Response

	dirItems := make([]fuse.Dirent, len(d.items)+1)
	for i, treeNode := range d.items {
		dirItems[i] = fuse.Dirent{
			Inode: uint64(i),
			Type:  fuse.DT_Dir,
			Name:  nameFromTreeNode(treeNode),
		}
	}
	lastIndex := len(d.items)
	dirItems[lastIndex] = fuse.Dirent{
		Inode: uint64(lastIndex),
		Type:  fuse.DT_File,
		Name:  "index.json",
	}
	return dirItems, nil
}

type File struct {
	fs.NodeRef
	// fuse     *fs.Server
	content  atomic.Value
	treeNode *expanders.TreeNode
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

// func (f *File) tick() {
// 	// Intentionally a variable-length format, to demonstrate size changes.
// 	f.count++
// 	s := fmt.Sprintf("%d\t%s\n", f.count, time.Now())
// 	f.content.Store(s)

// 	// For simplicity, this example tries to send invalidate
// 	// notifications even when the kernel does not hold a reference to
// 	// the node, so be extra sure to ignore ErrNotCached.
// 	if err := f.fuse.InvalidateNodeData(f); err != nil && err != fuse.ErrNotCached {
// 		log.Printf("invalidate error: %v", err)
// 	}
// }

// func (f *File) update() {
// 	tick := time.NewTicker(1 * time.Second)
// 	defer tick.Stop()
// 	for range tick.C {
// 		f.tick()
// 	}
// }
