package endpoints

import (
	"fmt"
	"strings"
)

// EndpointInfo represents an ARM endpoint
type EndpointInfo struct {
	TemplateURL string
	APIVersion  string
	URLSegments []EndpointSegment
}

// EndpointSegment reprsesents a segment of a template URL
type EndpointSegment struct {
	// holds a literal to match for fixed segments (e.g. /subscriptions/)
	Match string
	// holds the name of a templated segment (e.g. 'name' for /{name}/)
	Name string
}

// MatchResult holds information about an EndPointInfo match
type MatchResult struct {
	// indicates whether this was a match
	IsMatch bool
	// holds the values pulled from named segments of the URL
	Values map[string]string
}

// GetEndpointInfoFromURL builds an EndpointInfo instance
func GetEndpointInfoFromURL(templateURL string, apiVersion string) (EndpointInfo, error) {
	// This is currently generating at runtime, but would be a build-time task that generated code :-)
	originalTemplateURL := templateURL
	templateURL = strings.TrimPrefix(templateURL, "/")
	templateURLSegments := strings.Split(templateURL, "/")
	urlSegments := make([]EndpointSegment, len(templateURLSegments))
	for i, s := range templateURLSegments {
		if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
			name := strings.TrimPrefix(strings.TrimSuffix(s, "}"), "{")
			if name == "" {
				return EndpointInfo{}, fmt.Errorf("Segment index %d is a named segment but is missing the name", i)
			}
			urlSegments[i] = EndpointSegment{
				Name: name,
			}
		} else {
			urlSegments[i] = EndpointSegment{
				Match: s,
			}
		}
	}

	return EndpointInfo{
		TemplateURL: originalTemplateURL,
		APIVersion:  apiVersion,
		URLSegments: urlSegments,
	}, nil
}

func MustGetEndpointInfoFromURL(url string, apiVersion string) *EndpointInfo {
	endpoint, err := GetEndpointInfoFromURL(url, apiVersion)
	if err != nil {
		panic(err)
	}
	return &endpoint
}

// Match tests whether a URL matches the EndpointInfo (ignoring query string values)
func (ei *EndpointInfo) Match(url string) MatchResult {

	url = strings.TrimPrefix(url, "/")

	// strip off querystring for matching
	if i := strings.Index(url, "?"); i >= 0 {
		url = url[:i]
	}

	urlSegments := strings.Split(url, "/")
	if len(urlSegments) == len(ei.URLSegments) {
		isMatch := true
		matches := make(map[string]string)
		for i, segment := range ei.URLSegments {
			if segment.Name == "" {
				// literal match (ignore case)
				if !strings.EqualFold(segment.Match, urlSegments[i]) {
					isMatch = false
					break
				}
			} else {
				// capture name
				matches[segment.Name] = urlSegments[i]
			}
		}
		if isMatch {
			return MatchResult{
				IsMatch: true,
				Values:  matches,
			}
		}
	}

	return MatchResult{
		IsMatch: false,
	}
}

// BuildURL generates a URL based on the templateURL filling in placeholders with the values map
func (ei *EndpointInfo) BuildURL(values map[string]string) (string, error) {

	url := ""
	for _, segment := range ei.URLSegments {
		if segment.Match == "" {
			value := values[segment.Name]
			if value == "" {
				return "", fmt.Errorf("No value was found with name '%s'", segment.Match)
			}
			url += "/" + value
		} else {
			url += "/" + segment.Match
		}
	}
	if ei.APIVersion != "" {
		url += "?api-version=" + ei.APIVersion
	}
	return url, nil
}
