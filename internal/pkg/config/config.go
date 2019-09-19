package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

// Config represents the user configuration options
type Config struct {
	KeyBindings map[string]interface{} `json:"keyBindings,omitempty"`
	Editor      EditorConfig      `json:"editor,omitempty"`
}

// EditorConfig represents the user options for external editor
type EditorConfig struct {
	Command                 CommandConfig `json:"command,omitempty"`                 // The command to execute to launch the editor
	TranslateFilePathForWSL bool          `json:"translateFilePathForWSL,omitEmpty"` //nolint:golint,staticcheck // WSL use only. True to translate the path to a Windows path (e.g. when running under WSL but using a Windows editor)
	TempDir                 string        `json:"tempDir,omitempty"`                 // Specify the directory to use for temporary files for editing (defaults to OS temp dir)
}

// CommandConfig respresents the options for launching a command
type CommandConfig struct {
	Executable string   `json:"executable,omitempty"` // The program to run
	Arguments  []string `json:"args,omitempty"`       // The arguments to pass to the executable (filename will automatically be appended)
}

// Load the user configuration settings
func Load() (Config, error) {
	var config Config
	configLocation := "/root/.azbrowse-settings.json"
	user, err := user.Current()
	if err == nil {
		configLocation = user.HomeDir + "/.azbrowse-settings.json"
	}
	_, err = os.Stat(configLocation)
	if err != nil {
		// don't error on no config file
		return config, nil
	}
	configFile, err := os.Open(configLocation)
	if err != nil {
		return config, err
	}
	defer configFile.Close() //nolint: errcheck
	bytes, _ := ioutil.ReadAll(configFile)
	if err := json.Unmarshal(bytes, &config); err != nil {
		return config, err
	}
	return config, nil
}
