package filesystem

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

type Folder struct {
	fuse.Dirent
	fs              *FS
	treeNode        *expanders.TreeNode
	items           []*expanders.TreeNode
	subFolders      []*Folder
	indexContent    *expanders.ExpanderResponse
	canDelete       bool
	isParentDeleted func() bool
	isBeingDeleted  bool
	mu              sync.Mutex
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
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.isDeleteInProgress() {
		return nil, syscall.ENOENT
	}

	if d.indexContent == nil {
		d.LoadNodeFromARM()
	}

	if name == indexFileName(d.treeNode, d.indexContent) {
		file := &File{
			treeNode: d.treeNode,
			fs:       d.fs,
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
				isParentDeleted: func() bool {
					d.mu.Lock()
					defer d.mu.Unlock()

					return d.isBeingDeleted
				},
				fs: d.fs,
			}
			d.subFolders = append(d.subFolders, f)
			return f, nil
		}
	}

	log.Printf("Failed to match name: %s on folder %s", name, d.Name)

	return nil, syscall.ENOENT
}

var _ fs.NodeRemover = (*Folder)(nil)

func (d *Folder) isDeleteInProgress() bool {
	if d.isBeingDeleted || d.isParentDeleted() {
		return true
	}
	return false
}

// Todo currently `rm -r thing` gives an error deleting an RG but the delete is processed. Error:
//
// rm: WARNING: Circular directory structure.
// This almost certainly means that you have a corrupted file system.
// NOTIFY YOUR SYSTEM MANAGER.
// The following directory is part of the cycle:
// 30466/Deployments
// rm: cannot remove '30466': Input/output error
func (d *Folder) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	if !d.fs.editMode {
		log.Println("NOOP: Editing disabled")
		return fuse.EPERM
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.isDeleteInProgress() {
		return nil
	}

	for _, treeNode := range d.items {
		if nameFromTreeNode(treeNode) == req.Name && treeNode.ItemType != "subResource" {
			log.Printf("Found matching node '%s' doing delete: %+v", treeNode.Name+treeNode.ItemType, req)

			// Mark folder as being deleted
			for _, subFolder := range d.subFolders {
				if subFolder.Name == req.Name {
					func() {
						subFolder.mu.Lock()
						defer subFolder.mu.Unlock()
						subFolder.isBeingDeleted = true
					}()
				}
			}

			log.Printf("---> Attemping delete on %s", req.Name)
			// Start deletion
			// fallback to ARM request to delete
			fallback := true
			if treeNode.Expander != nil {
				deleted, err := treeNode.Expander.Delete(ctx, treeNode)
				fallback = (err == nil && !deleted)
			}
			if fallback {
				// fallback to ARM request to delete
				if treeNode.DeleteURL == "" {
					log.Printf("Delete not supported skipping node '%s' doing delete: %+v", treeNode.Name+treeNode.ItemType, req)
					return nil
				}
				_, err := armclient.LegacyInstance.DoRequest(ctx, "DELETE", treeNode.DeleteURL)
				if err != nil {
					panic(err)
				}
			}
			log.Printf("Delete complete: %+v", req)
		}
	}

	return nil
}

var _ fs.HandleReadDirAller = (*Folder)(nil)

func (d *Folder) LoadNodeFromARM() {
	rootContent, newItems, err := expanders.ExpandItem(ctx, d.treeNode)
	if err != nil {
		panic(err)
	}

	if d.fs.demoMode {
		rootContent.Response = views.StripSecretVals(rootContent.Response)
	}

	d.items = newItems
	d.indexContent = rootContent
}

func (d *Folder) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.isDeleteInProgress() {
		return nil, nil
	}

	if d.indexContent == nil {
		d.LoadNodeFromARM()
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
