package filesystem

import (
	"io/ioutil"
	"net/http"
	"testing"

	"bazil.org/fuse/fs"
	"bazil.org/fuse/fs/fstestutil"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const testServer = "https://management.azure.com"

func configureExpanders(t *testing.T) {
	// ctx := context.Background()

	// Load the db
	storage.LoadDB()

	dat, err := ioutil.ReadFile("../expanders/testdata/armsamples/resourcegroups/response.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expectedJSONResponse := string(dat)

	gock.New(testServer).
		Get("subscriptions").
		Reply(200).
		JSON(expectedJSONResponse)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, expanders.DummyTokenFunc())

	expanders.InitializeExpanders(client)
	// client.PopulateResourceAPILookup(ctx)

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
		t.Error(err)
	}

	assert.Equal(t, 3, len(files), "Expected 3 subscriptions from mock response")

}
