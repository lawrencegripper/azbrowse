package filesystem

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type Folder struct {
	fuse.Dirent
	treeNode     *expanders.TreeNode
	items        []*expanders.TreeNode
	indexContent *expanders.ExpanderResponse
	canDelete    bool
}

var _ fs.Node = (*Folder)(nil)

func (d *Folder) Attr(ctx context.Context, a *fuse.Attr) error {
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
		file := &File{
			treeNode: d.treeNode,
		}
		content := d.indexContent.Response

		if d.indexContent.ResponseType == "JSON" {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(d.indexContent.Response), "", "   ")
			if err != nil {
				// todo: log failure to json format
			} else {
				content = prettyJSON.String()
			}
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
	if d.indexContent == nil {
		rootContent, newItems, err := expanders.ExpandItem(ctx, d.treeNode)
		if err != nil {
			panic(err)
		}

		if *DemoMode {
			rootContent.Response = views.StripSecretVals(rootContent.Response)
		}

		d.items = newItems
		d.indexContent = rootContent
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
