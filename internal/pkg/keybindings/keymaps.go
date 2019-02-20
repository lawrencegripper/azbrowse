package keybindings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/jroimartin/gocui"
)

func initializeOverrides(filePath string) error {
	semanticKeyMap, err := loadBindingsFromFile(filePath)
	if err != nil {
		return err
	}

	normalizedKeyMap, err := normalizeSemanticKeyMap(*semanticKeyMap)
	if err != nil {
		return err
	}

	overrides = normalizedKeyMap
	return nil
}

func loadBindingsFromFile(filePath string) (*SemanticKeyMap, error) {
	jsonf, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonf.Close()
	bytes, _ := ioutil.ReadAll(jsonf)
	var semanticKeyMap SemanticKeyMap
	if err := json.Unmarshal(bytes, &semanticKeyMap); err != nil {
		return nil, err
	}
	return &semanticKeyMap, nil
}

func normalizeSemanticKeyMap(semanticKeyMap SemanticKeyMap) (KeyMap, error) {
	keyMap := KeyMap{}

	v := reflect.ValueOf(semanticKeyMap)
	for i := 0; i < v.NumField(); i++ {
		value := fmt.Sprintf("%s", v.Field(i).Interface())
		normalValue, err := normalizeSemanticValue(value)
		if err != nil {
			return nil, err
		}
		key := v.Type().Field(i).Name
		keyMap[key] = normalValue
	}
	return keyMap, nil
}

func normalizeSemanticValue(value string) (gocui.Key, error) {
	// TODO Parse semantics properly
	switch value {
	case "Up":
		return gocui.KeyArrowUp, nil
	case "Down":
		return gocui.KeyArrowDown, nil
	default:
		return 0, fmt.Errorf("%s is an unsupported value", value)
	}
}
