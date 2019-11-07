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

func Test_SubscriptionExpander_ReturnsContentOnSuccess(t *testing.T) {
	const testServer = "https://management.azure.com"
	const testPath = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups"

	const responseFile = "./testdata/armsamples/resourcegroups/response.json"

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

	expander := SubscriptionExpander{
		client: client,
	}

	ctx := context.Background()

	itemToExpand := &TreeNode{
		Display:        "Thingy1",
		Name:           "Thingy1",
		ID:             "/subscriptions/00000000-0000-0000-0000-000000000000",
		ExpandURL:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01",
		ItemType:       SubscriptionType,
		SubscriptionID: "00000000-0000-0000-0000-000000000000",
	}

	result := expander.Expand(ctx, itemToExpand)

	st.Expect(t, result.Err, nil)
	st.Expect(t, len(result.Nodes), 6)

	// Validate content
	st.Expect(t, result.Nodes[0].Name, "cloudshell")
	st.Expect(t, result.Nodes[0].ExpandURL, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cloudshell/resources?api-version=2017-05-10")

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)

}

func Test_SubscriptionExpander_ReturnsErrorOn500(t *testing.T) {
	const testServer = "https://management.azure.com"
	const testPath = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups"

	defer gock.Off()
	gock.New(testServer).
		Get(testPath).
		Reply(500)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, dummyTokenFunc())

	expander := SubscriptionExpander{
		client: client,
	}

	itemToExpand := &TreeNode{
		Display:        "Thingy1",
		Name:           "Thingy1",
		ID:             "/subscriptions/00000000-0000-0000-0000-000000000000",
		ExpandURL:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups?api-version=2018-05-01",
		ItemType:       SubscriptionType,
		SubscriptionID: "00000000-0000-0000-0000-000000000000",
	}

	ctx := context.Background()

	result := expander.Expand(ctx, itemToExpand)

	if result.Err == nil {
		t.Error("Failed expanding resource. Should have errored and didn't", result)
	}

	t.Log(result.Err.Error())

	st.Expect(t, gock.IsDone(), true)
}
