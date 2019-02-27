package handlers

import (
	"fmt"
	"testing"
)

// This test ensures that all the `mustGetEndpointInfoFromURL` calls in the swagger generated code succeed.
func TestGeneratedCodeInitialises(t *testing.T) {

	expander := SwaggerResourceExpander{}
	// Ensure that the generated types can be initialized
	resources := expander.getResourceTypes()

	t.Log(fmt.Printf("Generated swagger resources found: %v", len(resources)))
}
