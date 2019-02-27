package keybindings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/jroimartin/gocui"
)

// KeyMap reprsents the current mappings from Handler -> Key
type KeyMap map[string]gocui.Key

var handlers []KeyHandler
var overrides KeyMap
var usedKeys map[string]string

// Bind sets up key bindings for AzBrowse
func Bind(g *gocui.Gui) error {
	user, _ := user.Current()
	defaultFilePath := user.HomeDir + "/.azbrowse-bindings.json"
	return bindWithFileOverrides(g, defaultFilePath)
}

func bindWithFileOverrides(g *gocui.Gui, filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		err = initializeOverrides(filePath)
		if err != nil {
			return err
		}
	} // ignore file overrides if error
	return bindHandlersToKeys(g)
}

// AddHandler Adds a keybinding handler for use in Bind
func AddHandler(hnd KeyHandler) {
	handlers = append(handlers, hnd)
}

func getKeyBindings() KeyMap {
	keyBindings := KeyMap{}
	for _, handler := range handlers {
		if override, ok := overrides[handler.ID()]; ok {
			keyBindings[handler.ID()] = override
		} else {
			keyBindings[handler.ID()] = handler.DefaultKey()
		}
	}
	return keyBindings
}

// GetKeyBindingsAsStrings provides a map of Handler->Key in string format
func GetKeyBindingsAsStrings() map[string]string {
	keyBindings := map[string]string{}
	keys := getKeyBindings()
	for k, v := range keys {
		keyBindings[k] = KeyToStr[v]
	}
	return keyBindings
}

func bindHandlersToKeys(g *gocui.Gui) error {
	usedKeys = map[string]string{}
	for _, hnd := range handlers {
		if err := bindHandlerToKey(g, hnd); err != nil {
			return err
		}
	}
	// panic(usedKeys)
	return nil
}

func bindHandlerToKey(g *gocui.Gui, hnd KeyHandler) error {
	var key gocui.Key
	if k, ok := overrides[hnd.ID()]; ok {
		key = k
	} else {
		key = hnd.DefaultKey()
	}

	if err := checkKeyNotAlreadyInUse(hnd.Widget(), hnd.ID(), key); err != nil {
		return err
	}

	return g.SetKeybinding(hnd.Widget(), key, gocui.ModNone, hnd.Fn())
}

const reuseKeyError = "Please update your `~/.azbrowse-bindings.json` file to a valid configuration and restart"

func checkKeyNotAlreadyInUse(widget, id string, key gocui.Key) error {
	keyString := KeyToStr[key]
	// Check key isn't already use globally
	if usedBy, alreadyInUse := usedKeys[KeyToStr[key]]; alreadyInUse {
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

func initializeOverrides(filePath string) error {
	rawKeyMap, err := loadBindingsFromFile(filePath)
	if err != nil {
		return err
	}

	overrides, err = parseKeyValues(rawKeyMap)
	if err != nil {
		return err
	}

	return nil
}

func loadBindingsFromFile(filePath string) (map[string]string, error) {
	jsonf, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonf.Close() //nolint: errcheck
	bytes, _ := ioutil.ReadAll(jsonf)
	var rawKeyMap map[string]string
	if err := json.Unmarshal(bytes, &rawKeyMap); err != nil {
		return nil, err
	}
	return rawKeyMap, nil
}

func parseKeyValues(rawKeyMap map[string]string) (KeyMap, error) {
	keyMap := KeyMap{}

	for k, v := range rawKeyMap {
		parsedKey, err := parseKey(k)
		if err != nil {
			continue // ignore invalid keys
		}
		parsedValue, err := parseValue(v)
		if err != nil {
			continue // ignore invalid values
		}
		keyMap[parsedKey] = parsedValue
	}

	return keyMap, nil
}

func parseKey(key string) (string, error) {
	target := cleanKey(key)
	for _, k := range HandlerIds {
		if k == target {
			return target, nil
		}
	}
	return "", fmt.Errorf("%s is an unsupported key", key)
}

func parseValue(value string) (gocui.Key, error) {
	// TODO Parse semantics properly
	target := cleanValue(value)
	if val, ok := StrToKey[target]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("%s is an unsupported value", value)
}

func cleanKey(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}

func cleanValue(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}
