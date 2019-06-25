package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

// Config represents the user configuration options
type Config struct {
	KeyBindings map[string]string `json:"keyBindings,omitempty"`
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
