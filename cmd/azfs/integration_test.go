package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const expectedSubs = 4

func TestBrowseToRoot(t *testing.T) {
	// horrible hack
	boolTrue := true
	demoMode = &boolTrue

	path, err := ioutil.TempDir("", "azfs")
	if err != nil {
		t.Error(err)
	}

	conn, err := createFS(path)

	// wait for it to be ready
	<-conn.Ready

	time.Sleep(time.Second * 2)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Error(err)
		return
	}

	if len(files) != expectedSubs {
		t.Error("Expected 4 subscriptions")
	}

	// close the mount
	defer func() {
		conn.Close()
		err := os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}()
}
