package armclient

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	armEndpoint       string = "https://management.azure.com"
	armEndpointSuffix string = "management.azure.com"
)

func isArmURLPath(urlPath string) bool {
	urlPath = strings.ToLower(urlPath)
	return strings.HasPrefix(urlPath, "/subscriptions") ||
		strings.HasPrefix(urlPath, "/tenants") ||
		strings.HasPrefix(urlPath, "/providers")
}

func getRequestURL(path string) (string, error) {
	u, err := url.ParseRequestURI(path)

	if err != nil || !u.IsAbs() {
		if !isArmURLPath(path) {
			return "", errors.New("Url path specified is invalid")
		}

		return armEndpoint + path, nil
	}

	if u.Scheme != "https" {
		return "", errors.New("Scheme must be https")
	}

	// 127.0.0.1 is to allow integration testing with locally mocked server
	if !strings.HasSuffix(u.Hostname(), armEndpointSuffix) && u.Hostname() != "127.0.0.1" {
		return "", fmt.Errorf("'%s' is not an ARM endpoint", u.Hostname())
	}

	if !isArmURLPath(u.Path) {
		return "", fmt.Errorf("Url path '%s' is invalid", u.Path)
	}

	return path, nil
}
