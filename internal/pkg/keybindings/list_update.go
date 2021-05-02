package keybindings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/lawrencegripper/azbrowse/internal/pkg/editor"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
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
	if !enable || err != nil {
		// TODO - make this more extensible/generic, but for now allow swagger expander to handle items (e.g. from resource group expander)
		enable, err = expanders.GetSwaggerResourceExpander().CanUpdate(h.Context, item)
		if enable && err == nil {
			item.Metadata["UpdateWithSwaggerExpander"] = "true"
		}
	}
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
		err := json.Indent(&formattedBuf, []byte(content), "", "  ")
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

	h.status.Status("Opening content in editor...", false)

	updatedContent, err := editor.OpenForContent(formattedContent, fileExtension)
	if err != nil {
		h.status.Status(fmt.Sprintf("Error opening editor: %s", err), false)
		return err
	}

	if updatedContent == formattedContent {
		h.status.Status("No changes to content - no further action.", false)
		return nil
	}
	if updatedContent == "" {
		h.status.Status("Updated content empty - no further action.", false)
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

		expander := item.Expander
		if _, gotKey := item.Metadata["UpdateWithSwaggerExpander"]; gotKey {
			expander = expanders.GetSwaggerResourceExpander()
		}

		err = expander.Update(h.Context, item, updatedContent)
		if err != nil {
			evt.Failure = true
			evt.Message = evt.Message + " failed: " + err.Error()
			evt.Update()
		}

		done()
	}()

	return nil
}
