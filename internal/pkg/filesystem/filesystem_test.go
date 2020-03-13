package filesystem

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"testing"
	"time"

	"bazil.org/fuse/fs"
	"bazil.org/fuse/fs/fstestutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const testServer = "https://management.azure.com"

func createResponseLogger(t *testing.T) armclient.ResponseProcessor {
	return func(requestPath string, response *http.Response, responseBody string) {
		t.Log(requestPath)
		t.Log(responseBody)
	}
}

func getJSONFromFile(t *testing.T, path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(dat)
}

func configureExpanders(t *testing.T) {
	// ctx := context.Background()
	boolFalsePtr := false
	DemoMode = &boolFalsePtr
	EditMode = &boolFalsePtr

	// Load the db
	storage.LoadDB()

	// Create mock ARM API
	gock.New(testServer).
		Get("providers/microsoft.insights/metricNamespaces").
		Reply(200).
		JSON("{}")

	gock.New(testServer).
		Get("subscriptions").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/subscriptions/response.json"))

	gock.New(testServer).
		Get("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/1thing/providers/Microsoft.Storage/storageAccounts/thingsthings123").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resource/response.json"))

	gock.New(testServer).
		Get("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/1thing/resources").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resourcegroups/resourcelist.json"))

	gock.New(testServer).
		Get("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resourcegroups/response.json"))

	gock.New(testServer).
		Get("providers").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/providers/response.json"))

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, expanders.DummyTokenFunc(), createResponseLogger(t))

	expanders.InitializeExpanders(client)
	client.PopulateResourceAPILookup(ctx)

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
				fmt.Printf("%s STATUS: %s IN PROG: %t FAILED: %t \n", status.Icon(), message, status.InProgress, status.Failure)
			case <-timeout:
				// Update the UI
			}
		}
	}()

}

func setupMount(t *testing.T, filesys fs.FS) (mnt *fstestutil.Mount) {
	mnt, err := fstestutil.MountedT(t, filesys, nil)
	if err != nil {
		t.Fatal(err)
	}
	return mnt
}

func TestSubscriptions(t *testing.T) {
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	files, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(files), "Expected 3 subscriptions from mock response")

}

func TestRG(t *testing.T) {
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	//TODO: Remove
	// Walk to Sub level
	subFiles, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	files, err := ioutil.ReadDir(path.Join(mnt.Dir, subFiles[0].Name()))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 7, len(files), "Expected 6 RGs from mock response")
}

func TestRGResourceList(t *testing.T) {
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	//TODO: Remove
	// Walk to Sub level
	subFiles, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	builtPath := path.Join(mnt.Dir, subFiles[0].Name())
	//TODO: Remove
	// Walk to RG level
	rgFiles, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	builtPath = path.Join(builtPath, rgFiles[0].Name())
	files, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 13, len(files), "Expected 13 resource files from mock response")
}

func TestGetResource(t *testing.T) {
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	//TODO: Remove
	// Walk to Sub level
	subFiles, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	builtPath := path.Join(mnt.Dir, subFiles[0].Name())
	//TODO: Remove
	// Walk to RG level
	rgFiles, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	builtPath = path.Join(builtPath, rgFiles[0].Name())
	resourceFiles, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range resourceFiles {
		t.Log(f.Name())

	}

	builtPath = path.Join(builtPath, "thingsthings123")
	files, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 5, len(files), "Expected 5 subresources / files from mock response")
}
