package main

import (
	"regexp"
)

func stripSecretVals(s string) string {
	guidRegex := regexp.MustCompile(`[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?`)
	s = guidRegex.ReplaceAllString(s, "00000000-0000-0000-0000-HIDDEN000000")

	idRegex := regexp.MustCompile(`"id":.*`)
	s = idRegex.ReplaceAllString(s, `"id": "HIDDEN",`)

	locationRegex := regexp.MustCompile(`"location":.*`)
	s = locationRegex.ReplaceAllString(s, `"location": "HIDDEN-LOCATION",`)

	fqdnRegex := regexp.MustCompile(`"fqdn":.*`)
	s = fqdnRegex.ReplaceAllString(s, `"fqdn": "HIDDEN-FQDN",`)

	nameRegex := regexp.MustCompile(`"name":.*`)
	s = nameRegex.ReplaceAllString(s, `"name": "HIDDEN-NAME",`)

	storageKeyRegex := regexp.MustCompile(`([A-Za-z/+0-9]{86})(==)`)
	s = storageKeyRegex.ReplaceAllString(s, "HIDDEN-KEY")

	sshRegex := regexp.MustCompile(`ssh-rsa AAAA[0-9A-Za-z+/]+[=]{0,3} ([^@]+@[^@]+)`)
	s = sshRegex.ReplaceAllString(s, "SSH-PUBLIC-KEY-HIDDEN")
	return s
}
