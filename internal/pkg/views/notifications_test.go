package views

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/handlers"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/stuartleeks/gocui"
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

	notView.AddPendingDelete(&handlers.TreeNode{Name: "s1", DeleteURL: "http://delete/s1"})
	notView.AddPendingDelete(&handlers.TreeNode{Name: "s2", DeleteURL: "http://delete/s2"})
	notView.AddPendingDelete(&handlers.TreeNode{Name: "s3", DeleteURL: "http://delete/s3"})

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

	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg1", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg1"})
	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg2", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg2"})
	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg3", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg3"})

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

	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg1", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg1"})
	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg2", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg2"})
	notView.ConfirmDelete()

	// ConfirmDelete returns before it's finished
	WaitForFailureStatusEvent(t, statusEvents, 5)

	if count != 1 {
		t.Error("Expected 1 delete to be sent")
	}
}

func Test_Delete_AddPendingWhileDeleteInProgressRefused(t *testing.T) {
	if testing.Short() {
		t.Log("Skipping integration test")
		return
	}

	// Wait for the last test to clear down
	// Todo: This needs to be fixed. The ARMClient should be moved to a
	// struct and not package level methods.
	time.Sleep(time.Second * 5)

	statusEvents := eventing.SubscribeToStatusEvents()
	defer eventing.Unsubscribe(statusEvents)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg1", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg1"})
	notView.ConfirmDelete()

	notView.AddPendingDelete(&handlers.TreeNode{Name: "rg2", DeleteURL: ts.URL + "/subscriptions/1/resourceGroups/rg2"})

	// ConfirmDelete returns before it's finished
	failureStatus := WaitForFailureStatusEvent(t, statusEvents, 5)
	if failureStatus.Message != "Delete already in progress. Please wait for completion." {
		t.Errorf("Expected message 'Delete already in progress. Please wait for completion.' Got: %s", failureStatus.Message)
	}
}
