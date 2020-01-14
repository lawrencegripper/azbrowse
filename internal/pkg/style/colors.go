package style

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/stuartleeks/colorjson"
)

// ColorJSON formats the json with colors for the terminal
func ColorJSON(content string) string {
	d := json.NewDecoder(strings.NewReader(content))
	d.UseNumber()
	var obj interface{}
	err := d.Decode(&obj)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to display as JSON: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
		return content
	}
	jsonFormatter := colorjson.NewFormatter()
	jsonFormatter.Indent = 2
	s, err := jsonFormatter.Marshal(obj)
	if err != nil {
		return content
	}
	return string(s)
}

// Subtle use magenta and faint to format the text
func Subtle(s string) string {
	return color.New(color.FgMagenta, color.Faint).Sprint(s)
}

// Separator use magenta and faint to format the text
func Separator(s string) string {
	return color.New(color.FgBlack, color.Faint, color.Concealed).Sprint(s)
}

// Title make the text bold
func Title(s string) string {
	return color.New(color.Bold).Sprint(s)
}

// Loading make the text red and blink
func Loading(s string) string {
	return color.New(color.BlinkSlow, color.FgRed).Sprint(s)
}

// Completed make the text green
func Completed(s string) string {
	return color.New(color.FgGreen).Sprint(s)
}

// Header make the background blue and text white
func Header(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}

// Warning make the background red and text white
func Warning(s string) string {
	return color.New(color.BgRed, color.FgWhite).Sprint(s)
}

// Highlight make the background green and text white
func Highlight(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}

// Graph make the background green and text white
func Graph(s string) string {
	return color.New(color.FgBlue).Sprint(s)
}
