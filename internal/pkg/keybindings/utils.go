package keybindings

import (
	"github.com/atotto/clipboard"
	"github.com/lawrencegripper/azbrowse/internal/pkg/wsl"
)

func toggle(b bool) bool {
	return !b
}

func copyToClipboard(content string) error {
	var err error
	if wsl.IsWSL() {
		err = wsl.TrySetClipboard(content)
	} else {
		err = clipboard.WriteAll(content)
	}
	return err
}
