package expanders

import (
	"encoding/json"

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

func getJSONProperty(jsonData interface{}, properties ...string) (interface{}, error) {
	switch jsonData := jsonData.(type) {
	case map[string]interface{}:
		jsonMap := jsonData
		name := properties[0]
		jsonSubtree, ok := jsonMap[name]
		if ok {
			if len(properties) == 1 {
				return jsonSubtree, nil
			}
			return getJSONProperty(jsonSubtree, properties[1:]...)
		} else {
			return nil, nil // TODO - error if not found?
		}
	default:
		return nil, nil // TODO - error if not able to walk the tree?
	}
}

func getJSONPropertyFromString(jsonString string, properties ...string) (interface{}, error) {
	var jsonData map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		return nil, err
	}

	return getJSONProperty(jsonData, properties...)
}
