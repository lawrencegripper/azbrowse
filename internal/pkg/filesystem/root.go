package filesystem

import (
	"context"
	"os"
	"strings"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
)

type FS struct {
}

var ctx = context.Background()
var DemoMode *bool
var EditMode *bool

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
