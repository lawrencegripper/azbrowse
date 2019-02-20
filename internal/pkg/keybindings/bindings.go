package keybindings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/jroimartin/gocui"
)

type KeyMap map[string]gocui.Key

type SemanticKeyMap struct {
	ListNavigateDown string `json:"list-navigate-down"`
	ListNavigateUp   string `json:"list-navigate-up"`
}

func New() KeyMap {
	defaultFilePath := "bindings.json"
	return NewWithFile(defaultFilePath)
}

func NewWithFile(filePath string) KeyMap {
	keyMap := KeyMap{
		ListNavigateDown: gocui.KeyArrowDown,
		ListNavigateUp:   gocui.KeyArrowUp,
	}
	var err error
	keyMap, err = overrideKeyMapFromFile(filePath, keyMap)
	if err != nil {
		panic(fmt.Sprintf("Couldn't initialize keys, error: %+v", err))
	}
	return keyMap
}

func overrideKeyMapFromFile(filePath string, keyMap KeyMap) (KeyMap, error) {
	// Load string formatted bindings from file
	semanticKeyMap, err := loadBindingsFromFile(filePath)
	if err != nil {
		return nil, err
	}

	// Convert to KeyMap
	normalKeyMap, err := normalizeSemanticKeyMap(*semanticKeyMap)
	if err != nil {
		return nil, err
	}

	// Merge maps with file precedence
	keys := mergeKeyMaps(keyMap, normalKeyMap)
	return keys, nil
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

func normalizeSemanticValue(key string) (gocui.Key, error) {
	// TODO Parse semantics properly
	switch key {
	case "Up":
		return gocui.KeyArrowUp, nil
	case "Down":
		return gocui.KeyArrowDown, nil
	default:
		return 0, fmt.Errorf("%s is an unsupported key", key)
	}
}

func mergeKeyMaps(keymap1, keymap2 KeyMap) KeyMap {
	for k, _ := range keymap1 {
		keymap1[k] = keymap2[k]
	}
	return keymap1
}
