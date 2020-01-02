package keybindings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/stuartleeks/gocui"
)

// KeyMap reprsents the current mappings from Handler -> Key
type KeyMap map[string][]interface{}

var keyHandlers []KeyHandler
var overrides KeyMap
var usedKeys map[string]string

// Bind sets up key bindings for AzBrowse
func Bind(g *gocui.Gui) error {
	config, err := config.Load()
	if err != nil {
		return err
	}
	return bindWithConfigOverrides(g, config.KeyBindings)
}

func bindWithConfigOverrides(g *gocui.Gui, keyOverrideSettings map[string]interface{}) error {
	err := initializeOverrides(keyOverrideSettings)
	if err != nil {
		return err
	}
	return bindHandlersToKeys(g)
}

// AddHandler Adds a keybinding handler for use in Bind
func AddHandler(hnd KeyHandler) {
	keyHandlers = append(keyHandlers, hnd)
}

func getKeyBindings() KeyMap {
	keyBindings := KeyMap{}
	for _, handler := range keyHandlers {
		if override, ok := overrides[handler.ID()]; ok {
			keyBindings[handler.ID()] = override
		} else {
			keyBindings[handler.ID()] = []interface{}{handler.DefaultKey()}
		}
	}
	return keyBindings
}

// GetKeyBindingsAsStrings provides a map of Handler->Key in string format
func GetKeyBindingsAsStrings() map[string][]string {
	keyBindings := map[string][]string{}
	keys := getKeyBindings()
	for k, values := range keys {
		stringValues := []string{}
		for _, v := range values {
			if v != nil {
				stringValues = append(stringValues, keyToString(v))
			}
		}
		keyBindings[k] = stringValues
	}
	return keyBindings
}

func bindHandlersToKeys(g *gocui.Gui) error {
	usedKeys = map[string]string{}
	for _, hnd := range keyHandlers {
		if err := bindHandlerToKey(g, hnd); err != nil {
			return err
		}
	}
	// panic(usedKeys)
	return nil
}

func bindHandlerToKey(g *gocui.Gui, hnd KeyHandler) error {
	var keys []interface{}
	if k, ok := overrides[hnd.ID()]; ok {
		keys = k
	} else {
		defaultKey := hnd.DefaultKey()
		switch defaultKey := defaultKey.(type) {
		case []interface{}:
			keys = defaultKey // already an array
		default:
			keys = []interface{}{defaultKey} // not an array so wrap
		}
	}

	for _, key := range keys {
		if key == nil {
			continue
		}
		if err := checkKeyNotAlreadyInUse(hnd.Widget(), hnd.ID(), key); err != nil {
			return err
		}

		err := g.SetKeybinding(hnd.Widget(), key, gocui.ModNone, hnd.Fn())
		if err != nil {
			return err
		}
	}
	return nil
}

const reuseKeyError = "Please update your `~/.azbrowse-settings.json` file to a valid configuration and restart"

func checkKeyNotAlreadyInUse(widget, id string, key interface{}) error {

	keyString := keyToString(key)

	// Check key isn't already use globally
	if usedBy, alreadyInUse := usedKeys[keyString]; alreadyInUse {
		return errors.New("Failed when configurig `" + id + "`. The key `" + keyString + "` is already in use by `" + usedBy + "`(Global binding). " + reuseKeyError)
	}
	// Check key isn't already in use by a widget
	if usedBy, alreadyInUse := usedKeys[widget+keyString]; alreadyInUse {
		return errors.New("Failed when configurig `" + id + "`. The key `" + keyString + "` is already in use by `" + usedBy + "`. " + reuseKeyError)
	}

	// Track which keys are in use using a compound key of "WidgetKeyName"
	// this allows a key to be used by multiple widgets but not multiple times within a widget
	usedKeys[widget+keyString] = id

	return nil
}

func initializeOverrides(keyOverrideSettings map[string]interface{}) error {
	var err error
	overrides, err = parseKeyValues(keyOverrideSettings)
	if err != nil {
		return err
	}

	return nil
}

func parseKeyValues(keyOverrideSettings map[string]interface{}) (KeyMap, error) {
	keyMap := KeyMap{}

	for k, v := range keyOverrideSettings {
		parsedKey, err := parseKey(k)
		if err != nil {
			continue // ignore invalid keys
		}

		var values []string
		switch v.(type) { //nolint: gosimple
		case string:
			values = []string{v.(string)}
		case []interface{}:
			values = []string{}
			for _, item := range v.([]interface{}) {
				values = append(values, item.(string))
			}
		}

		parsedValues := []interface{}{}
		for _, value := range values {
			parsedValue, err := parseValue(value)
			if err != nil {
				continue // ignore invalid values
			}
			parsedValues = append(parsedValues, parsedValue)
		}
		keyMap[parsedKey] = parsedValues
	}

	return keyMap, nil
}

func parseKey(key string) (string, error) {
	target := cleanKey(key)
	for _, k := range keyHandlers {
		if k.ID() == target {
			return target, nil
		}
	}
	return "", fmt.Errorf("%s is an unsupported key", key)
}

func parseValue(value string) (interface{}, error) {
	// TODO Parse semantics properly
	target := cleanValue(value)
	if val, ok := StrToGocuiKey[target]; ok {
		return val, nil
	}

	if len(target) == 1 {
		// attempt as rune
		return rune(target[0]), nil
	}

	return 0, fmt.Errorf("%s is an unsupported value", value)
}

func cleanKey(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}

func cleanValue(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}

func keyToString(key interface{}) string {
	switch t := key.(type) { //nolint: gosimple
	case []interface{}:
		keyStrings := []string{}
		keys := key.([]interface{})
		for _, k := range keys {
			keyStrings = append(keyStrings, keyToString(k))
		}
		return strings.Join(keyStrings, ", ")
	case gocui.Key:
		return GocuiKeyToStr[key.(gocui.Key)]
	case rune:
		return strings.ToUpper(string(key.(rune)))
	default:
		panic(fmt.Sprintf("Unhandled key type: %v\n", t))
	}
}
