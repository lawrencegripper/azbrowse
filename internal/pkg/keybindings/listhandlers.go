package keybindings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/skratchdot/open-golang/open"
)

////////////////////////////////////////////////////////////////////
type ListActionsHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

func NewListActionsHandler(list *views.ListWidget, context context.Context) *ListActionsHandler {
	handler := &ListActionsHandler{
		Context: context,
		List:    list,
	}
	handler.Index = 7
	return handler
}

func (h ListActionsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return views.LoadActionsView(h.Context, h.List)
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListBackHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListBackHandler(list *views.ListWidget) *ListBackHandler {
	handler := &ListBackHandler{
		List: list,
	}
	handler.Index = 8
	return handler
}

func (h ListBackHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListBackLegacyHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListBackLegacyHandler(list *views.ListWidget) *ListBackLegacyHandler {
	handler := &ListBackLegacyHandler{
		List: list,
	}
	handler.Index = 9
	return handler
}

func (h ListBackLegacyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.GoBack()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListDownHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListDownHandler(list *views.ListWidget) *ListDownHandler {
	handler := &ListDownHandler{
		List: list,
	}
	handler.Index = 10
	return handler
}

func (h ListDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveDown()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListUpHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListUpHandler(list *views.ListWidget) *ListUpHandler {
	handler := &ListUpHandler{
		List: list,
	}
	handler.Index = 11
	return handler
}

func (h ListUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveUp()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListRightHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListRightHandler(list *views.ListWidget, editModeEnabled *bool) *ListRightHandler {
	handler := &ListRightHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.Index = 12
	return handler
}

func (h ListRightHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = true
		g.SetCurrentView("itemWidget")
		return nil
	}
}

////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////
type ListPageDownHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListPageDownHandler(list *views.ListWidget) *ListPageDownHandler {
	handler := &ListPageDownHandler{
		List: list,
	}
	handler.Index = 18
	return handler
}

func (h ListPageDownHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MovePageDown()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////
type ListPageUpHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListPageUpHandler(list *views.ListWidget) *ListPageUpHandler {
	handler := &ListPageUpHandler{
		List: list,
	}
	handler.Index = 19
	return handler
}

func (h ListPageUpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MovePageUp()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////
type ListEndHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListEndHandler(list *views.ListWidget) *ListEndHandler {
	handler := &ListEndHandler{
		List: list,
	}
	handler.Index = 20
	return handler
}

func (h ListEndHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveEnd()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////
type ListHomeHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListHomeHandler(list *views.ListWidget) *ListHomeHandler {
	handler := &ListHomeHandler{
		List: list,
	}
	handler.Index = 21
	return handler
}

func (h ListHomeHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.MoveHome()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListEditHandler struct {
	ListHandler
	List            *views.ListWidget
	EditModeEnabled *bool
}

func NewListEditHandler(list *views.ListWidget, editModeEnabled *bool) *ListEditHandler {
	handler := &ListEditHandler{
		List:            list,
		EditModeEnabled: editModeEnabled,
	}
	handler.Index = 13
	return handler
}

func (h ListEditHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.EditModeEnabled)
		h.EditModeEnabled = &tmp // memory leak?
		if *h.EditModeEnabled {
			g.Cursor = true
			g.SetCurrentView("itemWidget")
		} else {
			g.Cursor = false
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListExpandHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListExpandHandler(list *views.ListWidget) *ListExpandHandler {
	handler := &ListExpandHandler{
		List: list,
	}
	handler.Index = 14
	return handler
}

func (h ListExpandHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ExpandCurrentSelection()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListOpenHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

func NewListOpenHandler(list *views.ListWidget, context context.Context) *ListOpenHandler {
	handler := &ListOpenHandler{
		List:    list,
		Context: context,
	}
	handler.Index = 15
	return handler
}

func (h ListOpenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		portalURL := os.Getenv("AZURE_PORTAL_URL")
		if portalURL == "" {
			portalURL = "https://portal.azure.com"
		}
		url := portalURL + "/#@" + armclient.GetTenantID() + "/resource/" + item.ID
		span, _ := tracing.StartSpanFromContext(h.Context, "openportal:url")
		err := open.Run(url)
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: false,
				Failure:    true,
				Message:    "Failed opening resources in browser: " + err.Error(),
				Timeout:    time.Duration(time.Second * 4),
			})
			return nil
		}
		span.Finish()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListRefreshHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListRefreshHandler(list *views.ListWidget) *ListRefreshHandler {
	handler := &ListRefreshHandler{
		List: list,
	}
	handler.Index = 16
	return handler
}

func (h ListRefreshHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.Refresh()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListDeleteHandler struct {
	ListHandler
	List               *views.ListWidget
	NotificationWidget *views.NotificationWidget
}

func NewListDeleteHandler(list *views.ListWidget, notificationWidget *views.NotificationWidget) *ListDeleteHandler {
	handler := &ListDeleteHandler{
		List:               list,
		NotificationWidget: notificationWidget,
	}
	handler.Index = 2
	return handler
}

func (h ListDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		h.NotificationWidget.AddPendingDelete(item.Name, item.DeleteURL)
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListUpdateHandler struct {
	ListHandler
	List    *views.ListWidget
	status  *views.StatusbarWidget
	Context context.Context
	Content *views.ItemWidget
}

func NewListUpdateHandler(list *views.ListWidget, statusbar *views.StatusbarWidget, ctx context.Context, content *views.ItemWidget) *ListUpdateHandler {
	handler := &ListUpdateHandler{
		List:    list,
		status:  statusbar,
		Context: ctx,
		Content: content,
	}
	handler.Index = 17
	return handler
}

func (h ListUpdateHandler) getEditorConfig() (config.EditorConfig, error) {
	userConfig, err := config.Load()
	if err != nil {
		return config.EditorConfig{}, err
	}
	if userConfig.Editor.Command.Executable != "" {
		return userConfig.Editor, nil
	}
	// generate default config
	translateFilePathForWSL := os.Getenv("WSL_DISTRO_NAME") != "" // If WSL_DISTRO_NAME env var is set then translate path so that it is valid when loaded by VS code in Windows
	return config.EditorConfig{
		Command: config.CommandConfig{
			Executable: "code",
			Arguments:  []string{"--wait"},
		},
		TranslateFilePathForWSL: translateFilePathForWSL,
	}, nil
}

func (h ListUpdateHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentExpandedItem()
		if item == nil ||
			item.SwaggerResourceType == nil ||
			item.SwaggerResourceType.PutEndpoint == nil {
			return nil
		}

		editorConfig, err := h.getEditorConfig()
		if err != nil {
			return err
		}

		tempDir := editorConfig.TempDir
		if tempDir == "" {
			tempDir = os.TempDir() // fall back to Temp dir as default
		}
		tmpFile, err := ioutil.TempFile(tempDir, "azbrowse-*.json")
		if err != nil {
			h.status.Status(fmt.Sprintf("Cannot create temporary file: %s", err), false)
			return err
		}

		// Remember to clean up the file afterwards
		defer os.Remove(tmpFile.Name()) //nolint: errcheck

		originalJSON := h.Content.GetContent()

		_, err = tmpFile.WriteString(originalJSON)
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: false,
				Failure:    true,
				Message:    "Failed saving file for editing: " + err.Error(),
				Timeout:    time.Duration(time.Second * 4),
			})
			return nil
		}
		err = tmpFile.Close()
		if err != nil {
			eventing.SendStatusEvent(eventing.StatusEvent{
				InProgress: false,
				Failure:    true,
				Message:    "Failed closing file: " + err.Error(),
				Timeout:    time.Duration(time.Second * 4),
			})
			return nil
		}

		h.status.Status("Opening JSON in editor...", false)
		editorTmpFile := tmpFile.Name()
		// check if we should perform path translation for WSL (Windows Subsytem for Linux)
		if editorConfig.TranslateFilePathForWSL {
			cmd := exec.Command("wslpath", "-w", editorTmpFile)
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				return fmt.Errorf("Error running wslpath: %s", stderr.String())
			}
			editorTmpFile = strings.TrimSuffix(out.String(), "\n")
		}
		err = openEditor(editorConfig.Command, editorTmpFile)
		if err != nil {
			h.status.Status(fmt.Sprintf("Cannot open editor (ensure https://code.visualstudio.com is installed): %s", err), false)
			return nil
		}

		updatedJSONBytes, err := ioutil.ReadFile(tmpFile.Name())
		if err != nil {
			h.status.Status(fmt.Sprintf("Cannot open edited file: %s", err), false)
			return nil
		}

		updatedJSON := string(updatedJSONBytes)
		if updatedJSON == originalJSON {
			h.status.Status("No changes to JSON - no further action.", false)
			return nil
		}
		if updatedJSON == "" {
			h.status.Status("Updated JSON empty - no further action.", false)
			return nil
		}

		matchResult := item.SwaggerResourceType.Endpoint.Match(item.ExpandURL)
		if !matchResult.IsMatch {
			h.status.Status(fmt.Sprintf("item.ExpandURL didn't match current Endpoint"), false)
			return err
		}
		putURL, err := item.SwaggerResourceType.PutEndpoint.BuildURL(matchResult.Values)
		if err != nil {
			h.status.Status(fmt.Sprintf("Failed to build PUT URL '%s': %s", item.SwaggerResourceType.PutEndpoint.TemplateURL, err), false)
			return nil
		}

		done := h.status.Status(fmt.Sprintf("Making PUT request: %s", putURL), true)
		data, err := armclient.DoRequestWithBody(h.Context, "PUT", putURL, string(updatedJSON))
		done()
		if err != nil {
			h.status.Status(fmt.Sprintf("Error making PUT request: %s", err), false)
			return nil
		}

		errorMessage, err := getAPIErrorMessage(data)
		if err != nil {
			h.status.Status(fmt.Sprintf("Error checking for API Error message: %s: %s", data, err), false)
			return nil
		}
		if errorMessage != "" {
			h.status.Status(fmt.Sprintf("Error: %s", errorMessage), false)
			return nil
		}
		h.status.Status("Done", false)
		return nil
	}
}

func openEditor(command config.CommandConfig, filename string) error {
	// TODO - handle no Executable configured
	args := command.Arguments
	args = append(args, filename)
	cmd := exec.Command(command.Executable, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func getAPIErrorMessage(responseString string) (string, error) {
	var response map[string]interface{}

	err := json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		err = fmt.Errorf("Error parsing API response: %s: %s", responseString, err)
		return "", err
	}
	if response["error"] != nil {
		serializedError, err := json.Marshal(response["error"])
		if err != nil {
			err = fmt.Errorf("Error serializing error message: %s: %s", responseString, err)
			return "", err
		}
		message := string(serializedError)
		message = strings.Replace(message, "\r", "", -1)
		message = strings.Replace(message, "\n", "", -1)
		return message, nil
		// could dig into the JSON to pull out the error message property
	}
	return "", nil
}

////////////////////////////////////////////////////////////////////
