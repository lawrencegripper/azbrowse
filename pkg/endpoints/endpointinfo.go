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
	Name   string
	Prefix string
	Suffix string
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
	templateURL = strings.TrimSuffix(templateURL, "/")

	templateURLSegments := strings.Split(templateURL, "/")
	urlSegments := []EndpointSegment{}
	for i, s := range templateURLSegments {
		if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
			name := strings.TrimPrefix(strings.TrimSuffix(s, "}"), "{")
			if name == "" {
				return EndpointInfo{}, fmt.Errorf("Segment index %d is a named segment but is missing the name", i)
			}
			urlSegments = append(urlSegments, EndpointSegment{
				Prefix: "/",
				Name:   name,
			})
		} else {
			// check for a `docs('{name}')` style match
			parameterIndex := strings.Index(s, "('{")
			if parameterIndex < 0 {
				// Fixed value segment
				urlSegments = append(urlSegments, EndpointSegment{
					Prefix: "/",
					Match:  s,
				})
			} else {
				// If we have a segment such as `docs('{name}')` we expect the string to end with `}')`
				if !strings.HasSuffix(s, "}')") {
					return EndpointInfo{}, fmt.Errorf("Found parameterised segment but didn't find expected suffix")
				}
				fixedValue := s[:parameterIndex]
				name := s[parameterIndex+3 : len(s)-3]
				urlSegments = append(urlSegments, EndpointSegment{
					Prefix: "/",
					Match:  fixedValue,
				}, EndpointSegment{
					Prefix: "('",
					Suffix: "')",
					Name:   name,
				})
			}
		}
	}

	return EndpointInfo{
		TemplateURL: originalTemplateURL,
		APIVersion:  apiVersion,
		URLSegments: urlSegments,
	}, nil
}

// MustGetEndpointInfoFromURL creates an endpoint or panics
func MustGetEndpointInfoFromURL(url string, apiVersion string) *EndpointInfo {
	endpoint, err := GetEndpointInfoFromURL(url, apiVersion)
	if err != nil {
		panic(err)
	}
	return &endpoint
}

// Match tests whether a URL matches the EndpointInfo (ignoring query string values)
func (ei *EndpointInfo) Match(url string) MatchResult {

	// url = strings.TrimPrefix(url, "/")

	// strip off querystring for matching
	if i := strings.Index(url, "?"); i >= 0 {
		url = url[:i]
	}

	remainingURLToMatch := url

	matches := make(map[string]string)
	for i, segment := range ei.URLSegments {
		if !strings.HasPrefix(remainingURLToMatch, segment.Prefix) {
			return MatchResult{IsMatch: false}
		}
		remainingURLToMatch = remainingURLToMatch[len(segment.Prefix):]
		if segment.Match != "" {
			// literal match - check (case-insensitively) that the remainingURL starts with segment.Match
			matchTest := remainingURLToMatch[:len(segment.Match)]
			if strings.EqualFold(segment.Match, matchTest) {
				remainingURLToMatch = remainingURLToMatch[len(segment.Match):]
			} else {
				return MatchResult{IsMatch: false}
			}
		} else {
			// name match - match up to:
			//  * segment.Suffix (if set)
			//  * next Segment.Prefix (if there is a next segment)
			//  * end of the string failing that!
			matchTerminator := ""
			additionalSkipAmount := 0
			if segment.Suffix != "" {
				matchTerminator = segment.Suffix
				additionalSkipAmount = len(matchTerminator) // skip past the suffix if we match
			} else {
				if i+1 < len(ei.URLSegments) {
					matchTerminator = ei.URLSegments[i+1].Prefix // don't skip the prefix as that will be handled on the next loop iteration
				} else {
					// match is the rest of the URL
					matches[segment.Name] = remainingURLToMatch
					continue
				}
			}
			terminatorIndex := strings.Index(remainingURLToMatch, matchTerminator)
			if terminatorIndex < 0 {
				return MatchResult{IsMatch: false}
			}
			matches[segment.Name] = remainingURLToMatch[:terminatorIndex]
			remainingURLToMatch = remainingURLToMatch[terminatorIndex+additionalSkipAmount:]
		}
	}

	if remainingURLToMatch == "" {
		return MatchResult{
			IsMatch: true,
			Values:  matches,
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
		segmentValue := ""
		if segment.Match == "" {
			value := values[segment.Name]
			if value == "" {
				return "", fmt.Errorf("No value was found with name '%s'", segment.Match)
			}
			segmentValue = value
		} else {
			segmentValue = segment.Match
		}
		url += segment.Prefix + segmentValue + segment.Suffix
	}
	if ei.APIVersion != "" {
		url += "?api-version=" + ei.APIVersion
	}
	return url, nil
}

// GenerateValueArrayFromMap builds an ordered array of template match values from a templateValues map for use with BuildURLFromArray
func (ei *EndpointInfo) GenerateValueArrayFromMap(templateValues map[string]string) []string {
	valueArray := make([]string, len(templateValues))
	valueArrayIndex := 0
	for _, segment := range ei.URLSegments {
		if segment.Name != "" {
			valueArray[valueArrayIndex] = templateValues[segment.Name]
			valueArrayIndex++
		}
	}
	return valueArray
}

// BuildURLFromArray generates a URL based on the templateURL filling in placeholders with the values array by index
func (ei *EndpointInfo) BuildURLFromArray(values []string) (string, error) {

	url := ""
	valueIndex := 0
	for _, segment := range ei.URLSegments {
		if segment.Match == "" {
			if valueIndex >= len(values) {
				return "", fmt.Errorf("Too few values")
			}
			value := values[valueIndex] // TODO - check array bounds!
			url += "/" + value
			valueIndex++
		} else {
			url += "/" + segment.Match
		}
	}
	if ei.APIVersion != "" {
		url += "?api-version=" + ei.APIVersion
	}
	return url, nil
}
