package armclient

import (
	"testing"

	"gotest.tools/assert"
)

func Test_GetSubscriptionIDFromResourceID_WithValidResourceID(t *testing.T) {
	result := GetSubscriptionIDFromResourceID("/subscriptions/247bb195-ea6f-415c-b5b9-719c92d3cdab/resourceGroups")
	assert.Equal(t, result, "247bb195-ea6f-415c-b5b9-719c92d3cdab")
}

func Test_GetSubscriptionIDFromResourceID_WithInvalidResourceID(t *testing.T) {
	result := GetSubscriptionIDFromResourceID("/blah/")
	assert.Equal(t, result, "")
}
