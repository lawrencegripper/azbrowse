package expanders

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func Test_TenantExpander_ReturnsContentOnSuccess(t *testing.T) {
	const testServer = "https://management.azure.com"
	const testPath = "subscriptions"

	const responseFile = "./testdata/armsamples/subscriptions/response.json"
	dat, err := ioutil.ReadFile(responseFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expectedJSONResponse := string(dat)

	defer gock.Off()
	gock.New(testServer).
		Get(testPath).
		Reply(200).
		JSON(expectedJSONResponse)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, dummyTokenFunc())

	expander := TenantExpander{
		client: client,
	}

	ctx := context.Background()

	itemToExpand := &TreeNode{
		ItemType:  TentantItemType,
		ID:        "AvailableSubscriptions",
		ExpandURL: ExpandURLNotSupported,
	}

	result := expander.Expand(ctx, itemToExpand)

	st.Expect(t, result.Err, nil)
	st.Expect(t, len(result.Nodes), 3)

	// Validate content
	st.Expect(t, result.Nodes[0].Display, "Thingy1")
	st.Expect(t, result.Nodes[0].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01")

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)

}

func Test_TenantExpander_ReturnsErrorOn500(t *testing.T) {
	const testServer = "https://management.azure.com"
	const testPath = "subscriptions"

	defer gock.Off()
	gock.New(testServer).
		Get(testPath).
		Reply(500)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, dummyTokenFunc())

	expander := TenantExpander{
		client: client,
	}

	itemToExpand := &TreeNode{
		ItemType:  TentantItemType,
		ID:        "AvailableSubscriptions",
		ExpandURL: ExpandURLNotSupported,
	}

	ctx := context.Background()

	result := expander.Expand(ctx, itemToExpand)

	if result.Err == nil {
		t.Error("Failed expanding resource. Should have errored and didn't", result)
	}

	t.Log(result.Err.Error())

	st.Expect(t, gock.IsDone(), true)
}
