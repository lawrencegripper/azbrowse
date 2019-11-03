package expanders

import (
	"context"
	"net/http"
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// func FAILING_DefaultExpander_ReturnsContentOnSuccess(t *testing.T) {
// 	const testPath = "/subscriptions/thing"
// 	const responseJSON = `{ "json" : "value" }`
// 	mockServer := newARMServer(t, func(r *http.Request) bool {
// 		return r.URL.Path == testPath
// 	}, func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("{ 'json' : 'value' }"))
// 	})
// 	defer mockServer.TestServer.Close()

// 	// Set the ARM client to use out test server
// 	client := armclient.NewClientFromClientAndTokenFunc(mockServer.TestServer.Client(), dummyTokenFunc())

// 	defaultExpander := DefaultExpander{
// 		client: client,
// 	}

// 	ctx := context.Background()

// 	result := defaultExpander.Expand(ctx, &TreeNode{
// 		ExpandURL: mockServer.TestServer.URL + testPath,
// 	})

// 	if result.Response.Response != responseJSON {
// 		t.Errorf("Expected '%s' Go: '%s'", responseJSON, result.Response.Response)
// 	}

// 	if result.Err != nil {
// 		t.Error("Got: error wanted: no error", result)
// 	}

// 	if mockServer.MatchedCallCount != 1 {
// 		t.Error("Expected 1 match didn't get that")
// 	}

// 	if mockServer.TotalCallCount != 1 {
// 		t.Error("Expected 1 call didn't get that")
// 	}

// }

func Test_DefaultExpander_ReturnsErrorOn500(t *testing.T) {
	const testPath = "/subscriptions/thing"
	mockServer := new500ARMServer(t, func(r *http.Request) bool {
		return r.URL.Path == testPath
	})
	defer mockServer.TestServer.Close()

	// Set the ARM client to use out test server
	client := armclient.NewClientFromClientAndTokenFunc(mockServer.TestServer.Client(), dummyTokenFunc())

	defaultExpander := DefaultExpander{
		client: client,
	}

	ctx := context.Background()

	result := defaultExpander.Expand(ctx, &TreeNode{
		ExpandURL: mockServer.TestServer.URL + testPath,
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
