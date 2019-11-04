package endpoints

import "testing"

func TestNonMatchOnFixedString(t *testing.T) {

	// shouldn't match because "Microsoft.Web" != "Random"
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Random/sites/azbrowsetest/config")

	if matchResult.IsMatch {
		t.Error("Shouldn't match")
	}
}
func TestNonMatchOnMissingSegment(t *testing.T) {

	// Shouldn't match because url is missing final "config" segment
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Random/sites/azbrowsetest")

	if matchResult.IsMatch {
		t.Error("Shouldn't match")
	}
}
func TestNonMatchOnExtraSegment(t *testing.T) {

	//shouldn't match because url has extra "config" segment
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Random/sites/azbrowsetest/config")

	if matchResult.IsMatch {
		t.Error("Shouldn't match")
	}
}

func TestMatch(t *testing.T) {

	//should match
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Microsoft.Web/sites/azbrowsetest/config")

	if !matchResult.IsMatch {
		t.Error("Expected IsMatch to be true")
	} else {
		t.Log("verifying values")
		verifyMap(
			t,
			map[string]string{
				"subscriptionId":    "random",
				"resourceGroupName": "aaaaa",
				"name":              "azbrowsetest",
			},
			matchResult.Values)
	}
}
func TestMatchEndingWithNamedSegment(t *testing.T) {

	//should match
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Microsoft.Web/sites/azbrowsetest")

	if !matchResult.IsMatch {
		t.Error("Expected IsMatch to be true")
	} else {
		t.Log("verifying values")
		verifyMap(
			t,
			map[string]string{
				"subscriptionId":    "random",
				"resourceGroupName": "aaaaa",
				"name":              "azbrowsetest",
			},
			matchResult.Values)
	}
}
func TestMatchDifferentCase(t *testing.T) {

	//should match even though case differs on literal segments
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"/subscriptions/random/resourcegroups/aaaaa/providers/microsoft.web/SITES/azbrowsetest/config")

	if !matchResult.IsMatch {
		t.Error("Expected IsMatch to be true")
	} else {
		t.Log("verifying values")
		verifyMap(
			t,
			map[string]string{
				"subscriptionId":    "random",
				"resourceGroupName": "aaaaa",
				"name":              "azbrowsetest",
			},
			matchResult.Values)
	}
}
func TestMatchWithQueryString(t *testing.T) {

	//shouldn't match because url has extra "config" segment
	matchResult := getMatchResult(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"/subscriptions/random/resourceGroups/aaaaa/providers/Microsoft.Web/sites/azbrowsetest/config?api-version=2020-02-02")

	if !matchResult.IsMatch {
		t.Error("Expected IsMatch to be true")
	} else {
		t.Log("verifying values")
		verifyMap(
			t,
			map[string]string{
				"subscriptionId":    "random",
				"resourceGroupName": "aaaaa",
				"name":              "azbrowsetest",
			},
			matchResult.Values)
	}
}

func TestBuildWithParameterisedName(t *testing.T) {
	endpoint, err := GetEndpointInfoFromURL(
		"/datasources('{name}')",
		"")
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	url, err := endpoint.BuildURL(map[string]string{
		"name": "wibble",
	})
	expectedURL := "/datasources('wibble')"
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	if url != expectedURL {
		t.Errorf("Expected URL '%s' but got '%s", url, expectedURL)
	}
}
func TestMatchWithParameterisedNameMatch(t *testing.T) {

	matchResult := getMatchResult(
		"/datasources('{name}')",
		"/datasources('wibble')")

	if !matchResult.IsMatch {
		t.Error("Expected IsMatch to be true")
	} else {
		t.Log("verifying values")
		verifyMap(
			t,
			map[string]string{
				"name": "wibble",
			},
			matchResult.Values)
	}
}

func TestBuildWithoutQueryString(t *testing.T) {
	endpoint, err := GetEndpointInfoFromURL(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"")
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	url, err := endpoint.BuildURL(map[string]string{
		"subscriptionId":    "123456789",
		"resourceGroupName": "mygroup",
		"name":              "mysite",
	})
	expectedURL := "/subscriptions/123456789/resourceGroups/mygroup/providers/Microsoft.Web/sites/mysite/config"
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	if url != expectedURL {
		t.Errorf("Expected URL '%s' but got '%s", url, expectedURL)
	}
}
func TestBuildWithQueryString(t *testing.T) {
	endpoint, err := GetEndpointInfoFromURL(
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config",
		"2020-02-02")
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	url, err := endpoint.BuildURL(map[string]string{
		"subscriptionId":    "123456789",
		"resourceGroupName": "mygroup",
		"name":              "mysite",
	})
	expectedURL := "/subscriptions/123456789/resourceGroups/mygroup/providers/Microsoft.Web/sites/mysite/config?api-version=2020-02-02"
	if err != nil {
		t.Errorf("Expected success but got error: %s", err)
		return
	}
	if url != expectedURL {
		t.Errorf("Expected URL '%s' but got '%s", url, expectedURL)
	}
}

func verifyMap(t *testing.T, expected map[string]string, actual map[string]string) {

	for key, value := range expected {
		if actual[key] != value {
			t.Errorf("Expected key '%s' to have value '%s' but got '%s'", key, value, actual[key])
		}
	}

	if len(actual) != len(expected) {
		for key := range actual {
			if expected[key] == "" {
				t.Errorf("Actual has key '%s' (with value '%s') which was not expected ", key, actual[key])
			}
		}
	}
}

func getMatchResult(templateURL string, url string) MatchResult {
	endpoint, _ := GetEndpointInfoFromURL(templateURL, "")
	// TODO test err
	matchResult := endpoint.Match(url)
	return matchResult
}
