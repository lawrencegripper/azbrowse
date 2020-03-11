package filesystem

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"sync/atomic"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fuseutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
)

type File struct {
	fs.NodeRef

	mu sync.Mutex

	// fuse     *fs.Server
	content  atomic.Value
	treeNode *expanders.TreeNode
	// number of write-capable handles currently open
	writers uint
}

var _ fs.Node = (*File)(nil)

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	a.Mode = 0700
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
	f.mu.Lock()
	defer f.mu.Unlock()

	f.writers++

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

var _ fs.HandleWriter = (*File)(nil)

func (f *File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := ioutil.TempFile("", "tempazfs.*.json")

	if err != nil {
		return err
	}

	// Write original content to disk
	file.Write([]byte(f.content.Load().(string)))

	// mutate the result
	file.WriteAt(req.Data, req.Offset)

	// Update content in local system
	reader := bufio.NewReader(file)
	newContent, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	log.Println(string(newContent))

	// Submit to server
	apiSetID := f.treeNode.Metadata["SwaggerAPISetID"]
	apiSetPtr := expanders.GetSwaggerResourceExpander().GetAPISet(apiSetID)
	if apiSetPtr == nil {
		log.Println(err)
		return nil
	}
	apiSet := *apiSetPtr

	err = apiSet.Update(ctx, f.treeNode, string(newContent))
	if err != nil {
		log.Println(err)
		return err
	}

	f.content.Store(string(newContent))
	return nil
}

var _ = fs.NodeSetattrer(&File{})

const maxInt = int(^uint(0) >> 1)

func (f *File) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
	if req.Valid.Size() {
		if req.Size > uint64(maxInt) {
			return fuse.Errno(syscall.EFBIG)
		}

		data := []byte(f.content.Load().(string))

		newLen := int(req.Size)
		switch {
		case newLen > len(data):
			data = append(data, make([]byte, newLen-len(data))...)
		case newLen < len(data):
			data = data[:newLen]
		}

		f.content.Store(string(data))
	}
	return nil
}

var _ = fs.HandleReleaser(&File{})

func (f *File) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	return nil
}

var _ = fs.HandleFlusher(&File{})

func (f *File) Flush(ctx context.Context, req *fuse.FlushRequest) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.writers == 0 {
		// Read-only handles also get flushes. Make sure we don't
		// overwrite valid file contents with a nil buffer.
		return nil
	}

	return nil
}
