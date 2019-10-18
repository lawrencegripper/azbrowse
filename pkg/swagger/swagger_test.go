package swagger

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"

	"github.com/go-openapi/loads"
)

func Test_Simple_PreOrderedSpec(t *testing.T) {

	// Test a simple hierarchy that is preordered

	specJson := `{  "swagger": "2.0",
  "info": {
    "title": "DnsManagementClient",
    "description": "The DNS Management Client.",
    "version": "2018-05-01"
  },
  "host": "management.azure.com",
  "schemes": [
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
	],
	"paths": {
		"/test/{testname}": { "get": {} },
		"/test/{testname}/child1": {  "get": {} },
		"/test/{testname}/child1/{name2}": { "get": {} }
	}
}
`
	spec := json.RawMessage(specJson)
	doc, err := loads.Analyzed(spec, "")
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	config := Config{}

	var paths []*Path
	paths, err = MergeSwaggerDoc(paths, &config, doc, false)
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	assert.Assert(t, is.Len(paths, 1))

	// /test/{testname}
	path := paths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}")
	assert.Assert(t, is.Len(path.Children, 1))
	assert.Assert(t, is.Len(path.SubPaths, 0))

	// /test/{testname}/child1
	path = path.Children[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 1))

	// /test/{testname}/child1/{name2}
	path = path.SubPaths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1/{name2}")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 0))
}

func Test_Simple_NonOrderedSpec(t *testing.T) {

	// Test a simple hierarchy that is not in the order needed for hierarchy matching to automatically work

	specJson := `{  "swagger": "2.0",
  "info": {
    "title": "DnsManagementClient",
    "description": "The DNS Management Client.",
    "version": "2018-05-01"
  },
  "host": "management.azure.com",
  "schemes": [
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
	],
	"paths": {
		"/test/{testname}/child1/{name2}": { "get": {} },
		"/test/{testname}/child1": {  "get": {} },
		"/test/{testname}": { "get": {} }
	}
}
`
	spec := json.RawMessage(specJson)
	doc, err := loads.Analyzed(spec, "")
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	config := Config{}

	var paths []*Path
	paths, err = MergeSwaggerDoc(paths, &config, doc, false)
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	assert.Assert(t, is.Len(paths, 1))

	// /test/{testname}
	path := paths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}")
	assert.Assert(t, is.Len(path.Children, 1))
	assert.Assert(t, is.Len(path.SubPaths, 0))

	// /test/{testname}/child1
	path = path.Children[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 1))

	// /test/{testname}/child1/{name2}
	path = path.SubPaths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1/{name2}")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 0))
}

func Test_PathOverride(t *testing.T) {

	// Test a simple hierarchy that is not in the order needed for hierarchy matching to automatically work

	specJson := `{  "swagger": "2.0",
  "info": {
    "title": "DnsManagementClient",
    "description": "The DNS Management Client.",
    "version": "2018-05-01"
  },
  "host": "management.azure.com",
  "schemes": [
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
	],
	"paths": {
		"/testx/{testname}/child1/{name2}": { "get": {} },
		"/testx/{testname}/child1": {  "get": {} },
		"/test/{testname}": { "get": {} }
	}
}
`
	spec := json.RawMessage(specJson)
	doc, err := loads.Analyzed(spec, "")
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	config := Config{
		Overrides: map[string]PathOverride{
			// These paths would not be aggregated under /test/... without overrides
			"/testx/{testname}/child1": {Path: "/test/{testname}/child1"},
			"/testx/{testname}/child1/{name2}": {Path: "/test/{testname}/child1/{name2}"},
		},
	}

	var paths []*Path
	paths, err = MergeSwaggerDoc(paths, &config, doc, false)
	if err != nil {
		t.Logf("Failed to load spec: %v", err)
		t.Fail()
	}

	assert.Assert(t, is.Len(paths, 1))

	// /test/{testname}
	path := paths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}")
	assert.Assert(t, is.Len(path.Children, 1))
	assert.Assert(t, is.Len(path.SubPaths, 0))

	// /test/{testname}/child1
	path = path.Children[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1")
	assert.Equal(t, path.Operations.Get.Endpoint.TemplateURL, "/testx/{testname}/child1")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 1))

	// /test/{testname}/child1/{name2}
	path = path.SubPaths[0]
	assert.Equal(t, path.Endpoint.TemplateURL, "/test/{testname}/child1/{name2}")
	assert.Equal(t, path.Operations.Get.Endpoint.TemplateURL, "/testx/{testname}/child1/{name2}")
	assert.Assert(t, is.Len(path.Children, 0))
	assert.Assert(t, is.Len(path.SubPaths, 0))
}
