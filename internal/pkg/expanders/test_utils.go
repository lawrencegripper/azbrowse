package expanders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func dummyTokenFunc() func(clearCache bool) (armclient.AzCLIToken, error) {
	return func(clearCache bool) (armclient.AzCLIToken, error) {
		return armclient.AzCLIToken{
			AccessToken:  "bob",
			Subscription: "bill",
			Tenant:       "thing",
			TokenType:    "bearer",
		}, nil
	}
}

type mockMatchingFunc func(r *http.Request) bool

type mockARMServer struct {
	TotalCallCount     int
	MatchedCallCount   int
	UnMatchedCallCount int
	TestServer         *httptest.Server

	testServerCreate func() *httptest.Server
	matchFunc        mockMatchingFunc
}

func new500ARMServer(t *testing.T, matcher mockMatchingFunc) *mockARMServer {
	m := &mockARMServer{
		TotalCallCount:     0,
		MatchedCallCount:   0,
		UnMatchedCallCount: 0,
		matchFunc:          matcher,
	}

	m.testServerCreate = func() *httptest.Server {
		return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
			m.TotalCallCount = m.TotalCallCount + 1

			if m.matchFunc(r) {
				m.MatchedCallCount = m.MatchedCallCount + 1
			} else {
				m.UnMatchedCallCount = m.UnMatchedCallCount + 1
			}

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}))
	}

	m.TestServer = m.testServerCreate()

	return m
}
