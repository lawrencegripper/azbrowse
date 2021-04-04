package keybindings

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

type CopyHandler struct {
	GlobalHandler
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

var _ Command = &CopyHandler{}

func NewCopyHandler(content *views.ItemWidget, statusbar *views.StatusbarWidget) *CopyHandler {
	handler := &CopyHandler{
		Content:   content,
		StatusBar: statusbar,
	}
	handler.id = HandlerIDCopy
	return handler
}

func (h CopyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return h.Invoke()
	}
}

func (h *CopyHandler) DisplayText() string {
	return "Copy content"
}

func (h *CopyHandler) IsEnabled() bool {
	return true
}

func (h *CopyHandler) Invoke() error {
	var err error
	contentType := h.Content.GetContentType()
	content := h.Content.GetContent()

	var formattedContent string
	switch contentType {
	case interfaces.ResponseJSON:
		if !json.Valid([]byte(content)) {
			h.StatusBar.Status("Resource content is not valid JSON", false)
			return fmt.Errorf("Resource content is not valid JSON: %s", content)
		}

		var formattedBuf bytes.Buffer
		err = json.Indent(&formattedBuf, []byte(content), "", "  ")
		if err != nil {
			h.StatusBar.Status(fmt.Sprintf("Error formatting JSON for editor: %s", err), false)
			return err
		}

		formattedContent = formattedBuf.String()
	case interfaces.ResponseYAML:
		formattedContent = content // TODO: add YAML formatter

	default:
		formattedContent = content
	}

	if err := copyToClipboard(formattedContent); err != nil {
		h.StatusBar.Status(fmt.Sprintf("Failed to copy to clipboard: %s", err.Error()), false)
		return nil
	}
	h.StatusBar.Status("Current resource's content copied to clipboard", false)
	return nil
}
