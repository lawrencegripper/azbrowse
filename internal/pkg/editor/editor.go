package editor

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/awesome-gocui/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/wsl"
)

// OpenForContent creates a temporary file with the specified content and opens the editor for it. Returned value is the contents after editing
func OpenForContent(content string, fileExtension string) (string, error) {

	editorConfig, err := getEditorConfig()
	if err != nil {
		return "", err
	}

	tempDir := editorConfig.TempDir
	if tempDir == "" {
		tempDir = os.TempDir() // fall back to Temp dir as default
	}
	tmpFile, err := os.CreateTemp(tempDir, "azbrowse-*"+fileExtension)
	if err != nil {
		return "", fmt.Errorf("Cannot create temporary file: %s", err)
	}

	// Remember to clean up the file afterwards
	defer os.Remove(tmpFile.Name()) //nolint: errcheck

	_, err = tmpFile.WriteString(content)
	if err != nil {
		return "", fmt.Errorf("Failed to write file for editing: %s", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return "", fmt.Errorf("Failed to close file for editing: %s", err)
	}

	editorTmpFile := tmpFile.Name()
	// check if we should perform path translation for WSL (Windows Subsytem for Linux)
	if editorConfig.TranslateFilePathForWSL {
		editorTmpFile, err = wsl.TranslateToWindowsPath(editorTmpFile)
		if err != nil {
			return "", fmt.Errorf("Failed on WSL path translation: %s", err)
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
		return "", fmt.Errorf("Cannot open editor (ensure https://code.visualstudio.com is installed): %s", editorErr)
	}

	updatedContentBuf, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("Cannot open edited file: %s", err)
	}

	updatedContent := string(updatedContentBuf)

	return updatedContent, nil

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

func getEditorConfig() (config.EditorConfig, error) {
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
