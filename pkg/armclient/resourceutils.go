package armclient

import "regexp"

var subscriptionIDRegex *regexp.Regexp

// GetSubscriptionIDFromResourceID returns the subscription ID from a resource ID
// Returns empty string on no match
func GetSubscriptionIDFromResourceID(resourceID string) string {
	if subscriptionIDRegex == nil {
		subscriptionIDRegex = regexp.MustCompile("/subscriptions/(.*)/resourceGroups")
	}
	matches := subscriptionIDRegex.FindStringSubmatch(resourceID)
	if matches == nil {
		return ""
	}
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
