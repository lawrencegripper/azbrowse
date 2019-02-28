package handlers

import (
	"github.com/valyala/fastjson"

	"strings"
)

var fastJSONParser fastjson.Parser

func init() {
	fastJSONParser = fastjson.Parser{}
}

// DrawStatus converts a status string to an icon
func DrawStatus(s string) string {
	switch s {
	case "Deleting":
		return "☠"
	case "Failed":
		return "⛈"
	case "Updating":
		return "⟳"
	case "Running":
		return "⟳"
	case "Resuming":
		return "⛅"
	case "Starting":
		return "⛅"
	case "Provisioning":
		return "⌛"
	case "Creating":
		return "🏗"
	case "Preparing":
		return "🏗"
	case "Scaling":
		return "⚖"
	case "Suspended":
		return "⛔"
	case "Suspending":
		return "⛔"
	case "Succeeded":
		return "☼"
	}
	return ""
}

func getNamespaceFromARMType(s string) string {
	return strings.Split(s, "/")[0]
}
