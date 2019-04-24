package views

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

func Test_Delete_AddPendingDelete(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer g.Close()

	notView := NewNotificationWidget(0, 0, 47, false, g)
	g.SetManager(notView)

	notView.AddPendingDelete("s1", "http://delete/s1")
	notView.AddPendingDelete("s2", "http://delete/s2")
	notView.AddPendingDelete("s3", "http://delete/s3")

	builder := &strings.Builder{}
	err = notView.layoutInternal(builder)
	if err != nil {
		t.Error(err)
	}

	viewResult := builder.String()
	if !strings.Contains(viewResult, "s1") {
		t.Error("Missing s1 item")
	}
	if !strings.Contains(viewResult, "s2") {
		t.Error("Missing s1 item")
	}
	if !strings.Contains(viewResult, "s3") {
		t.Error("Missing s1 item")
	}

	if !strings.Contains(viewResult, "Do you want to delete these items?") {
		t.Error("Missing delete message")
	}
}

func Test_Delete_MessageSent(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}
	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		count = count + 1
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method. Got: %s", r.Method)
		}
	}))
	defer ts.Close()

	time.Sleep(time.Second * 5)

	// Set the ARM client to use out test server
	armclient.SetClient(ts.Client())
	armclient.SetAquireToken(dummyTokenFunc())

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer g.Close()
	notView := NewNotificationWidget(0, 0, 45, false, g)

	notView.AddPendingDelete("rg1", ts.URL+"/subscriptions/1/resourceGroups/rg1")
	notView.AddPendingDelete("rg2", ts.URL+"/subscriptions/1/resourceGroups/rg2")
	notView.AddPendingDelete("rg3", ts.URL+"/subscriptions/1/resourceGroups/rg3")
	notView.ConfirmDelete()

	// ConfirmDelete returns before it's finished
	WaitForCompletedStatusEvent(t, statusEvents, 5)

	if count != 3 {
		t.Error("Expected 3 delete's to be sent")
	}
}

func Test_Delete_StopAfterFailure(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}
	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	count := 0
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		count = count + 1

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}))
	defer ts.Close()

	// Set the ARM client to use out test server
	armclient.SetClient(ts.Client())
	armclient.SetAquireToken(dummyTokenFunc())

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer g.Close()
	notView := NewNotificationWidget(0, 0, 45, false, g)

	notView.AddPendingDelete("rg1", ts.URL+"/subscriptions/1/resourceGroups/rg1")
	notView.AddPendingDelete("rg2", ts.URL+"/subscriptions/1/resourceGroups/rg2")
	notView.ConfirmDelete()

	// ConfirmDelete returns before it's finished
	WaitForFailureStatusEvent(t, statusEvents, 5)

	if count != 1 {
		t.Error("Expected 1 delete to be sent")
	}
}

func Test_Delete_RefusedDeleteWhileInprogress(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}
	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	count := 0
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		count = count + 1
		time.Sleep(time.Second * 5) // Make the ConfirmDelete take a while
	}))
	defer ts.Close()

	// Set the ARM client to use out test server
	armclient.SetClient(ts.Client())
	armclient.SetAquireToken(dummyTokenFunc())

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer g.Close()
	notView := NewNotificationWidget(0, 0, 45, false, g)

	notView.AddPendingDelete("rg1", ts.URL+"/subscriptions/1/resourceGroups/rg1")
	notView.ConfirmDelete()

	// Simulate double tap of delet key
	notView.ConfirmDelete()

	// ConfirmDelete returns before it's finished
	failureStatus := WaitForFailureStatusEvent(t, statusEvents, 5)
	if failureStatus.Message != "Delete already in progress. Please wait for completion." {
		t.Errorf("Expected message 'Delete already in progress. Please wait for completion.' Got: %s", failureStatus.Message)
	}
}

func Test_Delete_RefusedAddPendingWhileInprogress(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}
	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("SEVER MESSAGE: received: %s method: %s", r.URL.String(), r.Method)
		time.Sleep(time.Second * 5) // Make the ConfirmDelete take a while
	}))
	defer ts.Close()

	// Set the ARM client to use out test server
	armclient.SetClient(ts.Client())
	armclient.SetAquireToken(dummyTokenFunc())

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer g.Close()
	notView := NewNotificationWidget(0, 0, 45, false, g)

	notView.AddPendingDelete("rg1", ts.URL+"/subscriptions/1/resourceGroups/rg1")
	notView.ConfirmDelete()

	// Wait for delete to be submitted
	time.Sleep(time.Second * 1)

	// Simulate double tap of delet key
	notView.AddPendingDelete("rg2", ts.URL+"/subscriptions/1/resourceGroups/rg2")

	// ConfirmDelete returns before it's finished
	failureStatus := WaitForFailureStatusEvent(t, statusEvents, 5)
	if failureStatus.Message != "Delete already in progress. Please wait for completion." {
		t.Errorf("Expected message 'Delete already in progress. Please wait for completion.' Got: %s", failureStatus.Message)
	}
}
