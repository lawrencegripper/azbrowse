package expanders

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

const (
	exampleResourceTypeContainerRepo = "Microsoft.ContainerRegistry/registries"
	exampleResourceIDContainerRepo   = "/subscriptions/5774ad8f-0000-0000-0000-0447910568d3/resourceGroups/stable/providers/Microsoft.ContainerRegistry/registries/test"
)

func Test_DefaultExpander_ReturnsContentOnSuccessAndUpdatesStatus(t *testing.T) {
	const testServer = "http://127.0.0.1"
	const testPath = "/subscriptions/thing"

	const expectedJSONResponse = `{"id":"/subscriptions/5774ad8f-0000-0000-0000-0447910568d3/resourceGroups/stable/providers/Microsoft.ContainerRegistry/registries/test","name":"test","type":"/subscriptions/5774ad8f-0000-0000-0000-0447910568d3/resourceGroups/stable/providers/Microsoft.ContainerRegistry/registries/test","sku":{"name":"","tier":""},"kind":"","location":"WestEurope","properties":{"provisioningState":"Failed"}}`
	defer gock.Off()
	gock.New(testServer).
		Get(testPath).
		Reply(200).
		JSON(expectedJSONResponse)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, dummyTokenFunc())

	defaultExpander := DefaultExpander{
		client: client,
	}

	ctx := context.Background()

	itemToExpand := &TreeNode{
		ExpandURL: testServer + testPath,
	}
	result := defaultExpander.Expand(ctx, itemToExpand)

	st.Expect(t, result.Err, nil)
	st.Expect(t, strings.TrimSpace(result.Response.Response), expectedJSONResponse)
	st.Expect(t, itemToExpand.StatusIndicator, "⛈")

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)

}

func Test_DefaultExpander_ReturnsErrorOn500(t *testing.T) {
	const testServer = "http://127.0.0.1"
	const testPath = "/subscriptions/thing"

	defer gock.Off()
	gock.New(testServer).
		Get(testPath).
		Reply(500)

	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	httpClient := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(httpClient)

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(httpClient, dummyTokenFunc())

	defaultExpander := DefaultExpander{
		client: client,
	}

	ctx := context.Background()

	result := defaultExpander.Expand(ctx, &TreeNode{
		ExpandURL: testServer + testPath,
	})

	if result.Err == nil {
		t.Error("Failed expanding resource. Should have errored and didn't", result)
	}

	st.Expect(t, gock.IsDone(), true)
}
