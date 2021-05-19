// +build !windows

package filesystem

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"bazil.org/fuse"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func shouldSkipIntTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test as only running short tests")

	}
	subscription := os.Getenv("TESTSUB")
	if subscription == "" {
		t.Skip("Skipping integration test as TESTSUB not set")
	}
}

func TestBrowseToRoot(t *testing.T) {
	gock.Off()
	shouldSkipIntTest(t)

	path, err := ioutil.TempDir("", "azfs")
	if err != nil {
		t.Error(err)
	}

	conn, err := createFS(path, "", false, false)
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup(path, conn)

	// wait for it to be ready
	<-conn.Ready

	time.Sleep(time.Second * 2)

	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(path)

	assert.Equal(t, 5, len(files))
}

func TestEditRG(t *testing.T) {
	gock.Off()
	shouldSkipIntTest(t)

	azfsPath, err := ioutil.TempDir("", "azfs")
	if err != nil {
		t.Error(err)
	}

	conn, err := createFS(azfsPath, "", false, false)
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup(azfsPath, conn)

	fmt.Println(azfsPath)

	// wait for it to be ready
	<-conn.Ready

	// Read root
	_, err = ioutil.ReadDir(azfsPath)
	if err != nil {
		t.Error(err)
		return
	}

	subscription := os.Getenv("TESTSUB")
	if subscription == "" {
		t.Error("must set 'TESTSUB' env to run tests")
		return
	}

	resource := os.Getenv("TESTRESOURCE")
	if resource == "" {
		t.Error("must set 'TESTRESOURCE' env to run tests")
		return
	}

	// Read sub
	subPath := path.Join(azfsPath, subscription)
	_, err = ioutil.ReadDir(subPath)
	if err != nil {
		t.Error(err)
		return
	}

	// Todo: Remove the need to walk the dir to get to this point
	segments := strings.Split(resource, "/")
	builtPath := subPath
	lastSegment := ""
	for _, segment := range segments {
		lastSegment = segment
		builtPath = path.Join(builtPath, segment)
		_, err = ioutil.ReadDir(builtPath)
		if err != nil {
			t.Error(err)
			return
		}

	}

	// Read file
	fullResourcePath := path.Join(builtPath, fmt.Sprintf("resource.%s.json", lastSegment))
	content, err := ioutil.ReadFile(fullResourcePath)
	if err != nil {
		t.Error(err)
		return
	}

	newContent := string(content)
	newContent = strings.Replace(newContent, "replaceme", "replaceme1", -1)

	// Write update
	err = ioutil.WriteFile(fullResourcePath, []byte(newContent), 0600)
	if err != nil {
		log.Println(err)
	}
}

func cleanup(path string, conn *fuse.Conn) {
	Close(path, conn) //nolint: errcheck
}
