// +build !windows

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
	printVerboseLogs = false
)

func createResponseLogger(t *testing.T) armclient.ResponseProcessor {
	return func(requestPath string, response *http.Response, responseBody string) {
		if printVerboseLogs {
			t.Log(requestPath)
			t.Log(responseBody)
		}
	}
}

var rgPath = fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", rgNameMock)
var resourcePath = fmt.Sprintf("%s/providers/Microsoft.Storage/storageAccounts/%s", rgPath, resourceNameMock)

func createExactGetMatcher(exactPath string) func(req *http.Request, gockReq *gock.Request) (bool, error) {
	return func(req *http.Request, gockReq *gock.Request) (bool, error) {
		if req.URL.Path == exactPath && req.Method == "GET" {
			return true, nil
		}
		return false, nil
	}
}

func addMockSub(t *testing.T) {
	gock.New(testServer).
		AddMatcher(createExactGetMatcher("/subscriptions")).
		Reply(200).
		File("../expanders/testdata/armsamples/subscriptions/response.json")
}

func addMockGroups(t *testing.T) {
	gock.New(testServer).
		AddMatcher(createExactGetMatcher("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups")).
		Reply(200).
		File("../expanders/testdata/armsamples/resourcegroups/response.json")
}

func addMockGroupResources(t *testing.T) {
	gock.New(testServer).
		AddMatcher(createExactGetMatcher(rgPath + "/resources")).
		Reply(200).
		File("../expanders/testdata/armsamples/resourcegroups/resourcelist.json")
}

func addMockResource(t *testing.T) {
	gock.New(testServer).
		AddMatcher(createExactGetMatcher(resourcePath)).
		Reply(200).
		File("../expanders/testdata/armsamples/resource/response.json")
}

func checkPendingMocks(t *testing.T) {
	pendingMocks := gock.Pending()
	assert.Equal(t, 0, len(pendingMocks), "Expect all mocks APIs to be called")
	for _, m := range pendingMocks {
		t.Logf("Pending mock not called: %s %+v", m.Request().URLStruct, m.Request())
	}
}

func configureExpandersAndGock(t *testing.T) {
	// Load the db
	storage.LoadDB()

	// Reset gock
	gock.Off()
	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromConfig(httpClient, expanders.DummyTokenFunc(), 5000, createResponseLogger(t))
	armclient.LegacyInstance = client

	expanders.InitializeExpanders(client, nil, nil, nil)
	providerData, err := storage.GetCache("providerCache")
	if err != nil || providerData == "" {
		gock.New(testServer).
			AddMatcher(createExactGetMatcher("/providers")).
			Reply(200).
			File("../expanders/testdata/armsamples/providers/response.json")

		client.PopulateResourceAPILookup(ctx)
	}

	// print status messages
	if printVerboseLogs {
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

}

func setupMount(t *testing.T, filesys fs.FS) (mnt *fstestutil.Mount) {
	mnt, err := fstestutil.MountedT(t, filesys, nil)
	if err != nil {
		t.Fatal(err)
	}
	return mnt
}

func Test_Get_Subs(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)

	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

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
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)

	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

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
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)

	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

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

	assert.Equal(t, 14, len(files), "Expected 13 resources + index file from mock response")
	checkPendingMocks(t)
}

func Test_Get_Resource_DirectNavigation(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)
	addMockResource(t)

	filesystem := &FS{}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	files, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 11, len(files), "Expected 4 sub resources + index file from mock response")
	checkPendingMocks(t)
}

func Test_Edit_Resource_DirectNavigation(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)
	addMockResource(t)

	// Allow an edit
	gock.New(testServer).
		Put(resourcePath).
		Reply(200).
		File("../expanders/testdata/armsamples/resource/response.json")

	filesystem := &FS{
		editMode: true,
	}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	_, err := ioutil.ReadDir(builtPath)
	if err != nil {
		t.Fatal(err)
	}

	//Todo: Do some editing here....
	resourceFilePath := path.Join(builtPath, fmt.Sprintf("resource.%s.json", resourceNameMock))
	err = ioutil.WriteFile(resourceFilePath, []byte("{ 'somejson': 'here' }"), 0777)

	assert.NoError(t, err, "Expect write to succeed")

	checkPendingMocks(t)
}

func Test_Delete_Resource_EditMode_Off(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)

	filesystem := &FS{
		editMode: false,
	}

	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	err := os.RemoveAll(builtPath)

	assert.Error(t, err, "Expected error when deleting resource with edit mode disabled: %s", builtPath)
	assert.Equal(t, 1, len(gock.Pending()), "Expect delete mock to not have been called")
}

func Test_Delete_Resource_DirectNavigation(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)

	filesystem := &FS{
		editMode: true,
	}

	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock, resourceNameMock)
	err := os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	checkPendingMocks(t)
}

func Test_Delete_RG_DirectNavigation(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)

	filesystem := &FS{
		editMode: true,
	}
	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)
	err := os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	checkPendingMocks(t)
}

// This test use "RemoveAll" which looks like it is different to "rm -r"
func Test_Delete_RG_AfterBrowse(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)

	// Allow delete call on rg
	gock.New(testServer).
		Delete(rgPath).
		Reply(200).
		JSON("{}")

	filesystem := &FS{
		editMode: true,
	}
	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)

	_, err := ioutil.ReadDir(builtPath)
	assert.NoError(t, err, "Expect to be able to list rg")

	err = os.RemoveAll(builtPath)

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	checkPendingMocks(t)
}

// This test uses "rm -r" as it behaves differently from "RemoveAll"
func Test_Delete_RG_AfterBrowse_Withrm(t *testing.T) {
	configureExpandersAndGock(t)
	addMockSub(t)
	addMockGroups(t)
	addMockGroupResources(t)

	// Allow gets and deletes on RG and providers under
	providers := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers", rgNameMock)
	rgPath := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", rgNameMock)
	gock.New(testServer).
		AddMatcher(func(req *http.Request, gockReq *gock.Request) (bool, error) {
			if (strings.HasPrefix(req.URL.Path, providers) || req.URL.Path == rgPath) && (req.Method == "DELETE" || req.Method == "GET") {
				return true, nil
			}
			return false, nil
		}).
		Persist().
		Reply(200).
		JSON("{}")

	filesystem := &FS{
		editMode: true,
	}

	mnt := setupMount(t, filesystem)

	defer mnt.Close()

	builtPath := path.Join(mnt.Dir, subNameMock, rgNameMock)

	_, err := ioutil.ReadDir(builtPath)
	assert.NoError(t, err, "Expect to be able to list rg")

	cmd := exec.Command("sh", "-c", "rm -r "+builtPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(string(output))
	}

	assert.NoError(t, err, "Expected no error when deleting resource: %s", builtPath)
	pendingMocks := gock.Pending()
	assert.Equal(t, 1, len(pendingMocks), "Expect nearly all mocks APIs to be called (pending catch all mock doesn't track calls to it)")
}
