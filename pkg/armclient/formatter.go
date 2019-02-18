package armclient

import (
	"bytes"
	"encoding/json"
)

func prettyJSON(buffer []byte) string {
	var prettyJSON string
	if len(buffer) > 0 {
		var jsonBuffer bytes.Buffer
		error := json.Indent(&jsonBuffer, buffer, "", "  ")
		if error != nil {
			return string(buffer)
		}
		prettyJSON = jsonBuffer.String()
	} else {
		prettyJSON = ""
	}

	return prettyJSON
}
