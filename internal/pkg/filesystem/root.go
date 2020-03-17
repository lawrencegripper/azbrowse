package filesystem

import (
	"context"
	"strings"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
)

type FS struct {
	demoMode bool
	editMode bool
}

var ctx = context.Background()

var _ fs.FS = (*FS)(nil)

func (f *FS) Root() (fs.Node, error) {
	// Create an empty tentant TreeNode. This by default expands
	// to show the current tenants subscriptions
	content, newItems, err := expanders.ExpandItem(ctx, &expanders.TreeNode{
		ItemType:  expanders.TentantItemType,
		ID:        "AvailableSubscriptions",
		ExpandURL: expanders.ExpandURLNotSupported,
	})
	if err != nil {
		panic(err)
	}

	return &Folder{
		Dirent: fuse.Dirent{
			// Inode: uint64(i),
			Type: fuse.DT_Dir,
			Name: "root",
		},
		items: newItems,
		treeNode: &expanders.TreeNode{
			Name:     "root",
			ItemType: "subscription",
		},
		indexContent:    content,
		isParentDeleted: func() bool { return false },
		fs:              f,
	}, nil
}

func nameFromTreeNode(treeNode *expanders.TreeNode) string {
	// fmt.Printf("%+v", *treeNode)
	return treeNode.Name
}

func indexFileName(treeNode *expanders.TreeNode, response *expanders.ExpanderResponse) string {
	if treeNode == nil || response == nil {
		return "index.json"
	}
	return treeNode.ItemType + "." + treeNode.Name + "." + strings.ToLower(string(response.ResponseType))
}
