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

	"github.com/awesome-gocui/gocui"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/lawrencegripper/azbrowse/internal/pkg/wsl"
)

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
	item := h.Content.GetNode()
	if item == nil || item.Expander == nil {
		return false
	}
	enable, err := item.Expander.CanUpdate(h.Context, item)
	return enable && err == nil
}
func (h *ListUpdateHandler) Invoke() error {
	item := h.Content.GetNode()
	if !h.IsEnabled() {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Updating not supported on this item",
			Timeout:    time.Duration(time.Second * 2),
		})
		return nil
	}

	editorConfig, err := h.getEditorConfig()
	if err != nil {
		return err
	}

	fileExtension := ".txt"
	contentType := h.Content.GetContentType()
	content := h.Content.GetContent()
	formattedContent := content

	switch contentType {
	case interfaces.ResponseJSON:
		fileExtension = ".json"

		if !json.Valid([]byte(content)) {
			h.status.Status("Resource content is not valid JSON", false)
			return fmt.Errorf("Resource content is not valid JSON: %s", content)
		}

		var formattedBuf bytes.Buffer
		err = json.Indent(&formattedBuf, []byte(content), "", "  ")
		if err != nil {
			h.status.Status(fmt.Sprintf("Error formatting JSON for editor: %s", err), false)
			return err
		}

		formattedContent = formattedBuf.String()
	case interfaces.ResponseYAML:
		fileExtension = ".yaml"
		formattedContent = content // TODO: add YAML formatter

	case interfaces.ResponseTerraform:
		fileExtension = ".tf"
		formattedContent = content // TODO: add Terraform formatter

	case interfaces.ResponseXML:
		fileExtension = ".xml"
		formattedContent = xmlfmt.FormatXML(content, "", "  ")
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
		gocui.Suspend()
	}

	editorErr := openEditor(editorConfig.Command, editorTmpFile)
	if editorConfig.RevertToStandardBuffer {
		// Init termbox to switch back to alternate buffer and Flush content
		err := gocui.Resume()
		if err != nil {
			panic(err)
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

	// Handle the updating of the item asyncronously to allow tcell to update and redraw UI
	go func() {
		errorhandling.RecoveryWithCleanup()

		evt, done := eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: true,
			Failure:    false,
			Message:    "Updating resource",
			Timeout:    time.Duration(time.Second * 30),
		})

		err = item.Expander.Update(h.Context, item, updatedJSON)
		if err != nil {
			evt.Failure = true
			evt.Message = evt.Message + " failed: " + err.Error()
			evt.Update()
		}

		done()
	}()

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
