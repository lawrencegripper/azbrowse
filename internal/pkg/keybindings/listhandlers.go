package keybindings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/internal/pkg/wsl"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"

	"github.com/nsf/termbox-go"
	"github.com/skratchdot/open-golang/open"
	"github.com/stuartleeks/gocui"
)

////////////////////////////////////////////////////////////////////
type ListActionsHandler struct {
	ListHandler
	List    *views.ListWidget
	Context context.Context
}

var _ Command = &ListActionsHandler{}

func NewListActionsHandler(list *views.ListWidget, context context.Context) *ListActionsHandler {
	handler := &ListActionsHandler{
		Context: context,
		List:    list,
	}
	handler.id = HandlerIDListActions
	return handler
}

func (h ListActionsHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}
func (h *ListActionsHandler) DisplayText() string {
	return "Show Actions"
}
func (h *ListActionsHandler) IsEnabled() bool {
	return h.List.CurrentExpandedItem() != nil
}
func (h *ListActionsHandler) Invoke() error {
	return views.LoadActionsView(h.Context, h.List)
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
	handler.id = HandlerIDListBack
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
	handler.id = HandlerIDListBackLegacy
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
	handler.id = HandlerIDListDown
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
	handler.id = HandlerIDListUp
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
	handler.id = HandlerIDListRight
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
	handler.id = HandlerIDListPageDown
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
	handler.id = HandlerIDListPageUp
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
	handler.id = HandlerIDListEnd
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
	handler.id = HandlerIDListHome
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
	handler.id = HandlerIDListEdit
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
	handler.id = HandlerIDListExpand
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

var _ Command = &ListOpenHandler{}

func NewListOpenHandler(list *views.ListWidget, context context.Context) *ListOpenHandler {
	handler := &ListOpenHandler{
		List:    list,
		Context: context,
	}
	handler.id = HandlerIDListOpen
	return handler
}

func (h ListOpenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListOpenHandler) DisplayText() string {
	return "Open in Portal"
}
func (h *ListOpenHandler) IsEnabled() bool {
	return true // TODO - filter to Azure resource nodes
}
func (h *ListOpenHandler) Invoke() error {
	item := h.List.CurrentItem()
	portalURL := os.Getenv("AZURE_PORTAL_URL")
	if portalURL == "" {
		portalURL = "https://portal.azure.com"
	}
	url := portalURL + "/#@" + armclient.LegacyInstance.GetTenantID() + "/resource/" + item.ID
	span, _ := tracing.StartSpanFromContext(h.Context, "openportal:url")
	var err error
	if wsl.IsWSL() {
		err = wsl.TryLaunchBrowser(url)
	} else {
		err = open.Run(url)
	}
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
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
	handler.id = HandlerIDListRefresh
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
	handler.id = HandlerIDListDelete
	return handler
}

func (h ListDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		item := h.List.CurrentItem()
		h.NotificationWidget.AddPendingDelete(item)
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
	Gui     *gocui.Gui
}

var _ Command = &ListUpdateHandler{}

func NewListUpdateHandler(list *views.ListWidget, statusbar *views.StatusbarWidget, ctx context.Context, content *views.ItemWidget, gui *gocui.Gui) *ListUpdateHandler {
	handler := &ListUpdateHandler{
		List:    list,
		status:  statusbar,
		Context: ctx,
		Content: content,
		Gui:     gui,
	}
	handler.id = HandlerIDListUpdate
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
	return config.EditorConfig{
		Command: config.CommandConfig{
			Executable: "code",
			Arguments:  []string{"--wait"},
		},
		TranslateFilePathForWSL: false, // previously used wsl.IsWSL to determine whether to translate path, but VSCode  now performs translation from WSL (so we get a bad path if we have translated it)
	}, nil
}

func (h ListUpdateHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}
func (h *ListUpdateHandler) DisplayText() string {
	return "Update current item"
}
func (h *ListUpdateHandler) IsEnabled() bool {
	item := h.List.CurrentExpandedItem()
	if item == nil ||
		item.SwaggerResourceType == nil ||
		item.SwaggerResourceType.PutEndpoint == nil ||
		item.Metadata == nil ||
		item.Metadata["SwaggerAPISetID"] == "" {
		return false
	}
	return true
}
func (h *ListUpdateHandler) Invoke() error {
	item := h.List.CurrentExpandedItem()
	if !h.IsEnabled() {
		return nil
	}

	editorConfig, err := h.getEditorConfig()
	if err != nil {
		return err
	}

	var formattedContent string
	fileExtension := ".txt"
	contentType := h.Content.GetContentType()
	content := h.Content.GetContent()

	switch contentType {
	case expanders.ResponseJSON:
		fileExtension = ".json"

		if !json.Valid([]byte(content)) {
			h.status.Status(fmt.Sprintf("Resource content is not valid JSON"), false)
			return fmt.Errorf("Resource content is not valid JSON: %s", content)
		}

		var formattedBuf bytes.Buffer
		err = json.Indent(&formattedBuf, []byte(content), "", "  ")
		if err != nil {
			h.status.Status(fmt.Sprintf("Error formatting JSON for editor: %s", err), false)
			return err
		}

		formattedContent = formattedBuf.String()
	case expanders.ResponseYAML:
		fileExtension = ".yaml"

		formattedContent = content // TODO: add YAML formatter
	}

	tempDir := editorConfig.TempDir
	if tempDir == "" {
		tempDir = os.TempDir() // fall back to Temp dir as default
	}
	tmpFile, err := ioutil.TempFile(tempDir, "azbrowse-*"+fileExtension)
	if err != nil {
		h.status.Status(fmt.Sprintf("Cannot create temporary file: %s", err), false)
		return err
	}

	// Remember to clean up the file afterwards
	defer os.Remove(tmpFile.Name()) //nolint: errcheck

	_, err = tmpFile.WriteString(formattedContent)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed saving file for editing: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
		return nil
	}
	err = tmpFile.Close()
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
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
		editorTmpFile, err = wsl.TranslateToWindowsPath(editorTmpFile)
		if err != nil {
			return err
		}
	}

	if editorConfig.RevertToStandardBuffer {
		// Close termbox to revert to normal buffer
		termbox.Close()
	}

	editorErr := openEditor(editorConfig.Command, editorTmpFile)
	if editorConfig.RevertToStandardBuffer {
		// Init termbox to switch back to alternate buffer and Flush content
		err = termbox.Init()
		if err != nil {
			return fmt.Errorf("Failed to reinitialise termbox: %v", err)
		}
		err = h.Gui.Flush()
		if err != nil {
			return fmt.Errorf("Failed to reinitialise termbox: %v", err)
		}
	}
	if editorErr != nil {
		h.status.Status(fmt.Sprintf("Cannot open editor (ensure https://code.visualstudio.com is installed): %s", editorErr), false)
		return nil
	}

	updatedJSONBytes, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		h.status.Status(fmt.Sprintf("Cannot open edited file: %s", err), false)
		return nil
	}

	updatedJSON := string(updatedJSONBytes)
	if updatedJSON == formattedContent {
		h.status.Status("No changes to JSON - no further action.", false)
		return nil
	}
	if updatedJSON == "" {
		h.status.Status("Updated JSON empty - no further action.", false)
		return nil
	}

	apiSetID := item.Metadata["SwaggerAPISetID"]
	apiSetPtr := expanders.GetSwaggerResourceExpander().GetAPISet(apiSetID)
	if apiSetPtr == nil {
		return nil
	}
	apiSet := *apiSetPtr

	err = apiSet.Update(h.Context, item, updatedJSON)
	if err != nil {
		h.status.Status(fmt.Sprintf("Error updating: %s", err), false)
		return nil
	}

	h.status.Status("Done", false)
	return nil

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

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListClearFilterHandler struct {
	ListHandler
	List *views.ListWidget
}

func NewListClearFilterHandler(list *views.ListWidget) *ListClearFilterHandler {
	handler := &ListClearFilterHandler{
		List: list,
	}
	handler.id = HandlerIDListClearFilter
	return handler
}

func (h ListClearFilterHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.List.ClearFilter()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type CommandPanelAzureSearchQueryHandler struct {
	ListHandler
	commandPanelWidget *views.CommandPanelWidget
	list               *views.ListWidget
	content            *views.ItemWidget
}

var _ Command = &CommandPanelAzureSearchQueryHandler{}

func NewCommandPanelAzureSearchQueryHandler(commandPanelWidget *views.CommandPanelWidget, content *views.ItemWidget, list *views.ListWidget) *CommandPanelAzureSearchQueryHandler {
	handler := &CommandPanelAzureSearchQueryHandler{
		commandPanelWidget: commandPanelWidget,
		content:            content,
		list:               list,
	}
	handler.id = HandlerIDAzureSearchQuery

	return handler
}

func (h *CommandPanelAzureSearchQueryHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if h.IsEnabled() {
			return h.Invoke()
		}
		return nil
	}
}

func (h *CommandPanelAzureSearchQueryHandler) DisplayText() string {
	return "Azure search query"
}

func (h *CommandPanelAzureSearchQueryHandler) IsEnabled() bool {
	currentItem := h.list.CurrentExpandedItem()
	if currentItem != nil && currentItem.SwaggerResourceType != nil && currentItem.SwaggerResourceType.Endpoint.TemplateURL == "/indexes('{indexName}')/docs" {
		return true
	}
	return false
}

func (h *CommandPanelAzureSearchQueryHandler) Invoke() error {
	h.commandPanelWidget.ShowWithText("search query:", "search=", nil, h.CommandPanelNotification)
	return nil
}

func (h *CommandPanelAzureSearchQueryHandler) CommandPanelNotification(state views.CommandPanelNotification) {

	if state.EnterPressed {
		queryString := state.CurrentText
		currentItem := h.list.CurrentExpandedItem()

		apiSetID := currentItem.Metadata["SwaggerAPISetID"]
		apiSetPtr := expanders.GetSwaggerResourceExpander().GetAPISet(apiSetID)
		if apiSetPtr == nil {
			return
		}
		apiSet := *apiSetPtr
		searchApiSet := apiSet.(expanders.SwaggerAPISetSearch)

		data, err := searchApiSet.DoRequest("GET", currentItem.ExpandURL+"&"+queryString)
		if err != nil {
			h.content.SetContent(fmt.Sprintf("%s", err), expanders.ResponseJSON, queryString)
		} else {
			h.content.SetContent(data, expanders.ResponseJSON, queryString)
		}
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListCopyItemIDHandler struct {
	ListHandler
	List      *views.ListWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &ListCopyItemIDHandler{}

func NewListCopyItemIDHandler(list *views.ListWidget, statusBar *views.StatusbarWidget) *ListCopyItemIDHandler {
	handler := &ListCopyItemIDHandler{
		List:      list,
		StatusBar: statusBar,
	}
	handler.id = HandlerIDListCopyItemID
	return handler
}

func (h ListCopyItemIDHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListCopyItemIDHandler) DisplayText() string {
	return "Copy current resource ID"
}
func (h *ListCopyItemIDHandler) IsEnabled() bool {
	return h.List.CurrentExpandedItem() != nil
}
func (h *ListCopyItemIDHandler) Invoke() error {
	item := h.List.CurrentExpandedItem()
	if item != nil {
		if err := copyToClipboard(item.ID); err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to copy resource ID to clipboard: %s", err.Error()), false)
			return nil
		}
		h.StatusBar.Status("Current resource ID copied to clipboard", false)
		return nil
	}
	return nil
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ListDebugCopyItemDataHandler struct {
	ListHandler
	List      *views.ListWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &ListDebugCopyItemDataHandler{}

func NewListDebugCopyItemDataHandler(list *views.ListWidget, statusBar *views.StatusbarWidget) *ListDebugCopyItemDataHandler {
	handler := &ListDebugCopyItemDataHandler{
		List:      list,
		StatusBar: statusBar,
	}
	handler.id = HandlerIDListDebugCopyItemData
	return handler
}

func (h ListDebugCopyItemDataHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *ListDebugCopyItemDataHandler) DisplayText() string {
	return "DEBUG: Copy current node data"
}
func (h *ListDebugCopyItemDataHandler) IsEnabled() bool {
	return h.List.CurrentExpandedItem() != nil
}
func (h *ListDebugCopyItemDataHandler) Invoke() error {
	item := h.List.CurrentExpandedItem()
	if item != nil {
		buf, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to generate node summary: %s", err.Error()), false)
			return nil
		}
		if err := copyToClipboard(string(buf)); err != nil {
			h.StatusBar.Status(fmt.Sprintf("Failed to copy node data to clipboard: %s", err.Error()), false)
			return nil
		}
		h.StatusBar.Status("Current node data copied to clipboard", false)
		return nil
	}
	return nil
}

////////////////////////////////////////////////////////////////////
