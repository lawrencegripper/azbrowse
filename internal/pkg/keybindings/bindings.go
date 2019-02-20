package keybindings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/keyhandlers"
)

type KeyMap map[string]gocui.Key

var handlers []keyhandlers.KeyHandler
var overrides KeyMap

func Bind(g *gocui.Gui) error {
	defaultFilePath := "bindings.json"
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

func AddHandler(hnd keyhandlers.KeyHandler) {
	handlers = append(handlers, hnd)
}

func bindHandlersToKeys(g *gocui.Gui) error {
	for _, hnd := range handlers {
		if err := bindHandlerToKey(g, hnd); err != nil {
			return err
		}
	}
	return nil
}

func bindHandlerToKey(g *gocui.Gui, hnd keyhandlers.KeyHandler) error {
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
	for _, k := range keyhandlers.HandlerIds {
		if k == target {
			return target, nil
		}
	}
	return "", fmt.Errorf("%s is an unsupported key", key)
}

func parseValue(value string) (gocui.Key, error) {
	// TODO Parse semantics properly
	target := cleanValue(value)
	switch target {
	case "up":
		return gocui.KeyArrowUp, nil
	case "down":
		return gocui.KeyArrowDown, nil
	case "left":
		return gocui.KeyArrowLeft, nil
	case "right":
		return gocui.KeyArrowRight, nil
	case "backspace":
		return gocui.KeyBackspace, nil
	case "backspace2":
		return gocui.KeyBackspace2, nil
	case "delete":
		return gocui.KeyDelete, nil
	case "home":
		return gocui.KeyHome, nil
	case "end":
		return gocui.KeyEnd, nil
	case "pageup":
		return gocui.KeyPgup, nil
	case "pagedown":
		return gocui.KeyPgdn, nil
	case "insert":
		return gocui.KeyInsert, nil
	case "tab":
		return gocui.KeyTab, nil
	case "space":
		return gocui.KeySpace, nil
	case "ctrl+2":
		return gocui.KeyCtrl2, nil
	case "ctrl+3":
		return gocui.KeyCtrl3, nil
	case "ctrl+4":
		return gocui.KeyCtrl4, nil
	case "ctrl+5":
		return gocui.KeyCtrl5, nil
	case "ctrl+6":
		return gocui.KeyCtrl6, nil
	case "ctrl+7":
		return gocui.KeyCtrl7, nil
	case "ctrl+8":
		return gocui.KeyCtrl8, nil
	case "ctrl+[":
		return gocui.KeyCtrlLsqBracket, nil
	case "ctrl+]":
		return gocui.KeyCtrlRsqBracket, nil
	case "ctrl+space":
		return gocui.KeyCtrlSpace, nil
	case "ctrl+_":
		return gocui.KeyCtrlUnderscore, nil
	case "ctrl+~":
		return gocui.KeyCtrlTilde, nil
	case "ctrl+a":
		return gocui.KeyCtrlA, nil
	case "ctrl+b":
		return gocui.KeyCtrlB, nil
	case "ctrl+c":
		return gocui.KeyCtrlC, nil
	case "ctrl+d":
		return gocui.KeyCtrlD, nil
	case "ctrl+e":
		return gocui.KeyCtrlE, nil
	case "ctrl+f":
		return gocui.KeyCtrlF, nil
	case "ctrl+g":
		return gocui.KeyCtrlG, nil
	case "ctrl+h":
		return gocui.KeyCtrlH, nil
	case "ctrl+i":
		return gocui.KeyCtrlI, nil
	case "ctrl+j":
		return gocui.KeyCtrlJ, nil
	case "ctrl+k":
		return gocui.KeyCtrlK, nil
	case "ctrl+l":
		return gocui.KeyCtrlL, nil
	case "ctrl+m":
		return gocui.KeyCtrlM, nil
	case "ctrl+n":
		return gocui.KeyCtrlN, nil
	case "ctrl+o":
		return gocui.KeyCtrlO, nil
	case "ctrl+p":
		return gocui.KeyCtrlP, nil
	case "ctrl+q":
		return gocui.KeyCtrlQ, nil
	case "ctrl+r":
		return gocui.KeyCtrlR, nil
	case "ctrl+s":
		return gocui.KeyCtrlS, nil
	case "ctrl+t":
		return gocui.KeyCtrlT, nil
	case "ctrl+u":
		return gocui.KeyCtrlU, nil
	case "ctrl+v":
		return gocui.KeyCtrlV, nil
	case "ctrl+w":
		return gocui.KeyCtrlW, nil
	case "ctrl+x":
		return gocui.KeyCtrlX, nil
	case "ctrl+y":
		return gocui.KeyCtrlY, nil
	case "ctrl+z":
		return gocui.KeyCtrlZ, nil
	case "esc":
		return gocui.KeyEsc, nil
	case "f1":
		return gocui.KeyF1, nil
	case "f2":
		return gocui.KeyF2, nil
	case "f3":
		return gocui.KeyF3, nil
	case "f4":
		return gocui.KeyF4, nil
	case "f5":
		return gocui.KeyF5, nil
	case "f6":
		return gocui.KeyF6, nil
	case "f7":
		return gocui.KeyF7, nil
	case "f8":
		return gocui.KeyF8, nil
	case "f9":
		return gocui.KeyF9, nil
	case "f10":
		return gocui.KeyF10, nil
	case "f11":
		return gocui.KeyF11, nil
	case "f12":
		return gocui.KeyF12, nil
	default:
		return 0, fmt.Errorf("%s is an unsupported value", value)
	}
}

func cleanKey(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}

func cleanValue(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}
