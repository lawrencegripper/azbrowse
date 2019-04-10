package views

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_tests(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		count = count + 1
	}))
	defer ts.Close()

	_, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("Expected 1 call to server")
	}
}
