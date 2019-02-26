package keybindings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/jroimartin/gocui"
)

type KeyMap map[string]gocui.Key

var handlers []KeyHandler
var overrides KeyMap

func Bind(g *gocui.Gui) error {
	user, _ := user.Current()
	defaultFilePath := user.HomeDir + "/.azbrowse-bindings.json"
	return BindWithFileOverrides(g, defaultFilePath)
}

func BindWithFileOverrides(g *gocui.Gui, filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		err = initializeOverrides(filePath)
		if err != nil {
			return err
		}
	} // ignore file overrides if error
	if err := bindHandlersToKeys(g); err != nil {
		return err
	}
	return nil
}

func AddHandler(hnd KeyHandler) {
	handlers = append(handlers, hnd)
}

func GetKeyBindings() KeyMap {
	keyBindings := KeyMap{}
	for _, handler := range handlers {
		if override, ok := overrides[handler.Id()]; ok {
			keyBindings[handler.Id()] = override
		} else {
			keyBindings[handler.Id()] = handler.DefaultKey()
		}
	}
	return keyBindings
}

func GetKeyBindingsAsStrings() map[string]string {
	keyBindings := map[string]string{}
	keys := GetKeyBindings()
	for k, v := range keys {
		keyBindings[k] = KeyToStr[v]
	}
	return keyBindings
}

func bindHandlersToKeys(g *gocui.Gui) error {
	for _, hnd := range handlers {
		if err := bindHandlerToKey(g, hnd); err != nil {
			return err
		}
	}
	return nil
}

func bindHandlerToKey(g *gocui.Gui, hnd KeyHandler) error {
	var key gocui.Key
	if k, ok := overrides[hnd.Id()]; ok {
		key = k
	} else {
		key = hnd.DefaultKey()
	}
	return g.SetKeybinding(hnd.Widget(), key, gocui.ModNone, hnd.Fn())
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
	defer jsonf.Close()
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
