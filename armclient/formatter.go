package armclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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

func responseDetail(response *http.Response, duration time.Duration, reqheaders []string) string {
	var buffer bytes.Buffer
	fmt.Fprint(&buffer, "---------- Request -----------------------\n")
	fmt.Fprintln(&buffer)
	fmt.Fprintf(&buffer, "%s %s\n", response.Request.Method, response.Request.URL.String())
	fmt.Fprintf(&buffer, "Host: %s\n", response.Request.URL.Host)
	fmt.Fprintf(&buffer, "Authorization: %s...\n", response.Request.Header.Get("Authorization")[0:15])
	fmt.Fprintf(&buffer, "User-Agent: %s\n", response.Request.UserAgent())
	fmt.Fprintf(&buffer, "Accept: %s\n", response.Request.Header.Get("Accept"))
	fmt.Fprintf(&buffer, "x-ms-client-request-id: %s\n", response.Request.Header.Get("x-ms-client-request-id"))

	if reqheaders != nil {
		for _, h := range reqheaders {
			fmt.Fprintf(&buffer, "%s: %s\n", h, response.Request.Header.Get(h))
		}
	}

	fmt.Fprintln(&buffer)
	fmt.Fprintf(&buffer, "---------- Response (%s) ------------\n", duration.Truncate(time.Millisecond).String())
	fmt.Fprintln(&buffer)
	fmt.Fprintf(&buffer, "%s: %s\n", response.Proto, response.Status)

	for name, headers := range response.Header {
		for _, h := range headers {
			name = strings.ToLower(name)
			fmt.Fprintf(&buffer, "%v: %v\n", name, h)
		}
	}

	return buffer.String()
}
