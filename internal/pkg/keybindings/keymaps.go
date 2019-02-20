package keybindings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

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
			// Ignore invalid values
			continue
		}
		key := v.Type().Field(i).Name
		if key == "" {
			// Ignore empty keys
			continue
		}
		keyMap[key] = normalValue
	}
	return keyMap, nil
}

func normalizeSemanticValue(value string) (gocui.Key, error) {
	// TODO Parse semantics properly
	target := removeWhitespace(strings.ToLower(value))
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

func removeWhitespace(str string) string {
	return strings.Replace(str, " ", "", -1)
}
