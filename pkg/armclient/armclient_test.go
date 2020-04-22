package armclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_ArmClient_AzCliToken_Refresh(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}))
	defer ts.Close()

	time.Sleep(time.Second * 5)

	cacheCleared := false
	tokenFunc := func(clearCache bool) (AzCLIToken, error) {
		if clearCache {
			cacheCleared = true
		}
		return AzCLIToken{}, nil
	}
	client := NewClientFromConfig(ts.Client(), tokenFunc, 5000)

	client.DoRequest(context.Background(), "GET", ts.URL+"/subscriptions/1/resourceGroups/rg1") //nolint: errcheck

	if !cacheCleared {
		t.Error("Expected cache to be cleared for azcli token")
	}
}

func Test_ArmClient_AzCliToken_DontRefresh(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		http.Error(w, "Unauthorized", http.StatusForbidden)

	}))
	defer ts.Close()

	time.Sleep(time.Second * 5)

	cacheCleared := false
	tokenFunc := func(clearCache bool) (AzCLIToken, error) {
		if clearCache {
			cacheCleared = true
		}
		return AzCLIToken{}, nil
	}
	// Set the ARM client to use out test server
	client := NewClientFromConfig(ts.Client(), tokenFunc, 5000)

	client.DoRequest(context.Background(), "GET", ts.URL+"/subscriptions/1/resourceGroups/rg1") //nolint: errcheck

	if cacheCleared {
		t.Error("Expected cache not to be cleared for azcli token")
	}
}
