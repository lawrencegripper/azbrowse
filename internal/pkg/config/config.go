package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

// Settings to enable different behavior on startup
type Settings struct {
	EnableTracing         bool
	HideGuids             bool
	NavigateToID          string
	FuzzerEnabled         bool
	FuzzerDurationMinutes int
	TenantID              string // the tenant ID to get an access token for from `az account get-access-token`
}

// Config represents the user configuration options
type Config struct {
	KeyBindings map[string]interface{} `json:"keyBindings,omitempty"`
	Editor      EditorConfig           `json:"editor,omitempty"`
}

// EditorConfig represents the user options for external editor
type EditorConfig struct {
	Command                 CommandConfig `json:"command,omitempty"`                 // The command to execute to launch the editor
	TranslateFilePathForWSL bool          `json:"translateFilePathForWSL,omitEmpty"` //nolint:golint,staticcheck // WSL use only. True to translate the path to a Windows path (e.g. when running under WSL but using a Windows editor)
	TempDir                 string        `json:"tempDir,omitempty"`                 // Specify the directory to use for temporary files for editing (defaults to OS temp dir)
	RevertToStandardBuffer  bool          `json:"revertToStandardBuffer,omitempty"`  // Set to true to revert to standard buffer while editing (e.g. for terminal-based editors)
}

// CommandConfig respresents the options for launching a command
type CommandConfig struct {
	Executable string   `json:"executable,omitempty"` // The program to run
	Arguments  []string `json:"args,omitempty"`       // The arguments to pass to the executable (filename will automatically be appended)
}

// Load the user configuration settings
func Load() (Config, error) {
	var config Config

	configLocation := os.Getenv("AZBROWSE_SETTINGS_PATH")
	if configLocation == "" {
		configLocation = "/root/.azbrowse-settings.json"
		user, err := user.Current()
		if err == nil {
			configLocation = user.HomeDir + "/.azbrowse-settings.json"
		}
	}
	_, err := os.Stat(configLocation)
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

var (
	debuggingEnabled = false
)

// SetDebuggingEnabled sets whether debugging is enabled
func SetDebuggingEnabled(value bool) {
	debuggingEnabled = value
}

// GetDebuggingEnabled gets whether debugging is enabled
func GetDebuggingEnabled() bool {
	return debuggingEnabled
}
