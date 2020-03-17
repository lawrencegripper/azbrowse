package filesystem

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
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

const (
	testServer       = "https://management.azure.com"
	subNameMock      = "1testsub"
	rgNameMock       = "1testrg"
	resourceNameMock = "1teststorageaccount"
)

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

var rgPath string
var resourcePath string

func addMockSub(t *testing.T) {
	gock.New(testServer).
		Get("/subscriptions").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/subscriptions/response.json"))

	gock.New(testServer).
		Get("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resourcegroups/response.json"))
}

func addMockRG(t *testing.T) {
	gock.New(testServer).
		Get(rgPath + "/resources").
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resourcegroups/resourcelist.json"))

}

func addMockResource(t *testing.T) {
	gock.New(testServer).
		Get(resourcePath).
		Reply(200).
		JSON(getJSONFromFile(t, "../expanders/testdata/armsamples/resource/response.json"))
}

func checkPendingMocks(t *testing.T) {
	pendingMocks := gock.Pending()
	assert.Equal(t, 0, len(pendingMocks), "Expect all mocks APIs to be called")
	for _, m := range pendingMocks {
		t.Logf("Pending mock not called: %s %+v", m.Request().URLStruct, m.Request())
	}
}

func configureExpanders(t *testing.T) {
	// ctx := context.Background()
	boolFalsePtr := false
	boolTruePtr := true
	DemoMode = &boolFalsePtr
	EditMode = &boolTruePtr

	// Load the db
	storage.LoadDB()

	rgPath = fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", rgNameMock)
	resourcePath = fmt.Sprintf("%s/providers/Microsoft.Storage/storageAccounts/%s", rgPath, resourceNameMock)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, expanders.DummyTokenFunc(), createResponseLogger(t))
	armclient.LegacyInstance = client

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
				t.Logf("%s STATUS: %s IN PROG: %t FAILED: %t \n", status.Icon(), message, status.InProgress, status.Failure)
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

func Test_Get_Subs(t *testing.T) {
	addMockSub(t)
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	files, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		t.Log(f.Name())
	}

	assert.Equal(t, 4, len(files), "Expected 3 subscriptions + index file from mock response")
	checkPendingMocks(t)
}

func Test_Get_Sub_TreeWalk(t *testing.T) {
	addMockSub(t)
	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	// Walk to Sub level
	subFiles, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}

	files, err := ioutil.ReadDir(path.Join(mnt.Dir, subFiles[0].Name()))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 7, len(files), "Expected 6 RGs + index file from mock response")
	checkPendingMocks(t)
}

func Test_Get_RG_WalkTree(t *testing.T) {
	addMockSub(t)
	addMockRG(t)

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

	assert.Equal(t, 13, len(files), "Expected 12 resources + index file from mock response")
	checkPendingMocks(t)
}

func Test_Get_Resource_DirectNavigation(t *testing.T) {
	addMockSub(t)
	addMockRG(t)

	configureExpanders(t)
	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	files, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 5, len(files), "Expected 4 sub resources + index file from mock response")
	checkPendingMocks(t)
}

// func Test_Edit_Resource_DirectNavigation(t *testing.T) {
// 	configureExpanders(t)
// 	filesystem := &FS{}

// 	mnt := setupMount(t, filesystem)

// 	defer mnt.Close()
// 	defer storage.CloseDB()

// 	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
// 	files, err := ioutil.ReadDir(builtPath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	//Todo: Do some editing here....

// 	assert.Equal(t, 5, len(files), "Expected 4 sub resources + index file from mock response")
// }

func Test_Delete_Resource_DirectNavigation(t *testing.T) {
	addMockSub(t)
	addMockRG(t)

	configureExpanders(t)
	filesystem := &FS{}

	// Allow delete call on rg
	rgPath := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", rgNameMock)
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	err := os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)

	checkPendingMocks(t)
}

func Test_Delete_RG_DirectNavigation(t *testing.T) {
	addMockSub(t)

	configureExpanders(t)
	filesystem := &FS{}

	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)
	err := os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	checkPendingMocks(t)
}

// This test use "RemoveAll" which looks like it is different to "rm -r"
func Test_Delete_RG_AfterBrowse(t *testing.T) {
	addMockSub(t)
	addMockRG(t)

	configureExpanders(t)
	filesystem := &FS{}

	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)

	_, err := ioutil.ReadDir(builtPath)
	assert.NoError(t, err, "Expect to be able to list rg")

	err = os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	checkPendingMocks(t)
}

// This test uses "rm -r" as it behaves differently from "RemoveAll"
func Test_Delete_RG_AfterBrowse_Withrm(t *testing.T) {
	addMockSub(t)
	addMockRG(t)

	configureExpanders(t)

	// Allow to delete this stuff
	providers := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers", rgNameMock)
	rgPath := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", rgNameMock)

	gock.New(testServer).
		AddMatcher(func(req *http.Request, gockReq *gock.Request) (bool, error) {
			if (strings.HasPrefix(req.URL.Path, providers) || req.URL.Path == rgPath) && req.Method == "DELETE" {
				return true, nil
			}
			return false, nil
		}).
		Persist().
		Reply(200).
		JSON("{}")

	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()
	defer storage.CloseDB()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)

	_, err := ioutil.ReadDir(builtPath)
	assert.NoError(t, err, "Expect to be able to list rg")

	cmd := exec.Command("sh", "-c", "rm -r "+builtPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(string(output))
	}

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	assert.Equal(t, 1, len(gock.Pending()), "Expect nearly all mocks APIs to be called, the persistent mock doesn't count")
}
