package keybindings

import (
	"strings"
)

// Command represents a command that can be listed in the Command Palette
type Command interface {
	ID() string
	DisplayText() string
	IsEnabled() bool
	Invoke() error
}

// SortByDisplayText allows sorting an array of Commands by DisplayText
type SortByDisplayText []Command

func (s SortByDisplayText) Len() int {
	return len(s)
}
func (s SortByDisplayText) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortByDisplayText) Less(i, j int) bool {
	return strings.Compare(s[i].DisplayText(), s[j].DisplayText()) < 0
}
