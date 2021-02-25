package interfaces

// CommandPanel provides an interface for the command panel widget to prevent circular references between views and expanders
type CommandPanel interface {
	Hide()
	ShowWithText(title string, s string, options *[]CommandPanelListOption, handler CommandPanelNotificationHandler)
}

// CommandPanelListOption represents a list item that can be displayed in the CommandPanel
type CommandPanelListOption struct {
	ID          string
	DisplayText string
}

// CommandPanelNotification holds the event state for a panel changed notification
type CommandPanelNotification struct {
	CurrentText  string
	SelectedID   string
	EnterPressed bool
}

// CommandPanelNotificationHandler is the function signature for a panel changed notification handler
type CommandPanelNotificationHandler func(state CommandPanelNotification)

// ExpanderResponseType is used to indicate the text format of a response
type ExpanderResponseType string

const (
	// interfaces.ResponsePlainText indicates the response type should not be parsed or colourised
	ResponsePlainText ExpanderResponseType = "Text"
	// interfaces.ResponseJSON indicates the response type can be parsed and colourised as JSON
	ResponseJSON ExpanderResponseType = "JSON"
	// ResponseYAML indicates the response type can be parsed and colourised as YAML
	ResponseYAML ExpanderResponseType = "YAML"
	// ResponseXML indicates the response type can be parsed and colourised as XML
	ResponseXML ExpanderResponseType = "XML"
	// ResponseTerraform indicates the response type can be parsed and colourised as Terraform
	ResponseTerraform ExpanderResponseType = "Terraform"
)

// ItemWidget provides an interface for the command panel widget to prevent circular references between views and expanders
type ItemWidget interface {
	SetContent(content string, contentType ExpanderResponseType, title string)
}
