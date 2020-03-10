package views

import (
	"regexp"
)

// NameAndNodeType represents fields `name` and `type` from a node structure
type NameAndNodeType struct {
	name     string
	nodeType string
}

// Matcher for connectionstrings config
var globalConnectionStringsConfig = NameAndNodeType{"connectionstrings", "Microsoft.Web/sites/config"}

// getNameAndType of a json object, if possible.
func getNameAndType(s string) (NameAndNodeType, bool) {
	typRe := regexp.MustCompile(`"type":\s*"(.+?(?:\\"|[^"])*)"`)
	possibleTypes := typRe.FindStringSubmatch(s)
	nameRe := regexp.MustCompile(`"name":\s*"(.+?(?:\\"|[^"])*)"`)
	possibleNames := nameRe.FindStringSubmatch(s)

	// if there is not sufficient data for name and type, fail
	if len(possibleTypes) < 1 || len(possibleNames) < 1 {
		return NameAndNodeType{"", ""}, false
	}

	// note: we'll always select the first value, if multiple matches are present
	return NameAndNodeType{possibleNames[1], possibleTypes[1]}, true
}

// StripSecretVals removes secret values
func StripSecretVals(s string) string {
	guidRegex := regexp.MustCompile(`[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?`)
	s = guidRegex.ReplaceAllString(s, "00000000-0000-0000-0000-HIDDEN000000")

	idRegex := regexp.MustCompile(`"id":\s*".+?(?:\\"|[^"])*"`)
	s = idRegex.ReplaceAllString(s, `"id": "HIDDEN"`)

	managedByRegex := regexp.MustCompile(`"managedBy":\s*".+?(?:\\"|[^"])*"`)
	s = managedByRegex.ReplaceAllString(s, `"managedBy": "HIDDEN_MANAGED_BY"`)

	locationRegex := regexp.MustCompile(`"location":\s*".+?(?:\\"|[^"])*"`)
	s = locationRegex.ReplaceAllString(s, `"location": "HIDDEN-LOCATION"`)

	fqdnRegex := regexp.MustCompile(`"fqdn":\s*".+?(?:\\"|[^"])*"`)
	s = fqdnRegex.ReplaceAllString(s, `"fqdn": "HIDDEN-FQDN"`)

	scmURIRegex := regexp.MustCompile(`"scmUri":\s*".+?(?:\\"|[^"])*"`)
	s = scmURIRegex.ReplaceAllString(s, `"scmUri": "HIDDEN-URI"`)

	storageKeyRegex := regexp.MustCompile(`([A-Za-z/+0-9]{86})(==)`)
	s = storageKeyRegex.ReplaceAllString(s, "HIDDEN-KEY")

	// this jwt regex only matches if the full string contains only a token (e.g. ^...$)
	jwtTokenRegex := regexp.MustCompile(`^(["]?)[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*(["]?)$`)
	s = jwtTokenRegex.ReplaceAllString(s, "${1}HIDDEN-JWT${2}")

	anyKeyRegex := regexp.MustCompile(`(".*[kK]ey":\s*").+?(?:\\"|[^"])*(")`)
	s = anyKeyRegex.ReplaceAllString(s, "${1}HIDDEN-KEY${2}")

	anySecretRegex := regexp.MustCompile(`(".*[sS]ecret":\s*").+?(?:\\"|[^"])*(")`)
	s = anySecretRegex.ReplaceAllString(s, "${1}HIDDEN-SECRET${2}")

	anyPasswordRegex := regexp.MustCompile(`(".*[pP]assword":\s*").+?(?:\\"|[^"])*(")`)
	s = anyPasswordRegex.ReplaceAllString(s, "${1}HIDDEN-PASSWORD${2}")

	sshRegex := regexp.MustCompile(`ssh-rsa AAAA[0-9A-Za-z+/]+[=]{0,3}[ ]?(?:[^@]+@[^@"]+)?`)
	s = sshRegex.ReplaceAllString(s, "SSH-PUBLIC-KEY-HIDDEN")

	// begin type-specific implementations
	// note: this requires that the name and type have not been obfuscated yet
	if nt, ok := getNameAndType(s); ok {
		switch nt {
		case globalConnectionStringsConfig:
			valueRegex := regexp.MustCompile(`"value":\s*".+?(?:\\"|[^"])*"`)
			s = valueRegex.ReplaceAllString(s, `"value": "HIDDEN"`)
		}
	}

	// name must run after type-specific, since type-specific parses the name node
	nameRegex := regexp.MustCompile(`"name":\s*".+?(?:\\"|[^"])*"`)
	s = nameRegex.ReplaceAllString(s, `"name": "HIDDEN-NAME"`)

	return s
}
