package wsl

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// IsWSL attempts to determine if we're running on WSL
func IsWSL() bool {
	// If WSL_DISTRO_NAME env var is set then we're on WSL (unless someone is messing with us)
	// Some early releases of WSL didn't set this IIRC
	return os.Getenv("WSL_DISTRO_NAME") != ""
}

// TranslateToWindowsPath converts a Linux path under WSL to a Windows-accessible path
func TranslateToWindowsPath(localPath string) (string, error) {
	cmd := exec.Command("wslpath", "-w", localPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Error running wslpath: %s", stderr.String())
	}
	windowsPath := strings.TrimSuffix(out.String(), "\n")
	return windowsPath, nil
}

// TrySetClipboard attempts to set the clipboard content if on WSL
func TrySetClipboard(text string) error {
	cmd := exec.Command("clip.exe")
	var out bytes.Buffer
	var stderr bytes.Buffer
	in := bytes.NewBufferString(text)
	cmd.Stdin = in
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running wslpath: %s", stderr.String())
	}
	return nil
}
