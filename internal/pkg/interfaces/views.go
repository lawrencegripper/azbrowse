package interfaces

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
