package expanders

import (
	"context"
	"net/http"
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func Test_DefaultExpander_ReturnsErrorOn500(t *testing.T) {
	mockServer := new500ARMServer(t, func(r *http.Request) bool {
		if r.URL.Path == "/subscriptions/thing" {
			return true
		} else {
			return false
		}
	})
	defer mockServer.TestServer.Close()

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(mockServer.TestServer.Client(), dummyTokenFunc())

	defaultExpander := DefaultExpander{
		client: client,
	}

	ctx := context.Background()

	result := defaultExpander.Expand(ctx, &TreeNode{
		ExpandURL: mockServer.TestServer.URL + "/subscriptions/thing",
	})

	if result.Err == nil {
		t.Error("Failed expanding resource. Should have errored and didn't", result)
	}

	if mockServer.MatchedCallCount != 1 {
		t.Error("Expected 1 match didn't get that")
	}

	if mockServer.TotalCallCount != 1 {
		t.Error("Expected 1 call didn't get that")
	}
}
