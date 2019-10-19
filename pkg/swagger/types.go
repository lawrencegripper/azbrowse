package swagger

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

// ResourceType holds information about resources that can be displayed
type ResourceType struct {
	Display        string
	Endpoint       *endpoints.EndpointInfo
	Verb           string
	DeleteEndpoint *endpoints.EndpointInfo
	PatchEndpoint  *endpoints.EndpointInfo
	PutEndpoint    *endpoints.EndpointInfo
	Children       []ResourceType // Children are auto-loaded (must be able to build the URL => no additional template URL values)
	SubResources   []ResourceType // SubResources are not auto-loaded (these come from the request to the endpoint)
	FixedContent   string
	SubPathRegex   *RegexReplace
}

/////////////////////////////////////////////////////////////////////////////
// Path models

// Path represents a path that we want to consider emitting in code-gen. It is derived from
type Path struct {
	Name                  string
	CondensedEndpointPath string
	FixedContent          string
	Endpoint              *endpoints.EndpointInfo // The logical endpoint. May be overridden for an operation
	Operations            PathOperations
	Children              []*Path
	SubPaths              []*Path
	SubPathRegex          *RegexReplace
}

// RegexReplace holds match and replacement info
type RegexReplace struct {
	Match   string
	Replace string
}

// PathOperations gives details on the operations for a resource
type PathOperations struct {
	Get    PathOperation
	Delete PathOperation
	Patch  PathOperation
	Post   PathOperation
	Put    PathOperation
}

// PathOperation represents an operation on the path (GET, PUT, ...)
type PathOperation struct {
	Permitted bool                    // true if the operation is permitted for the path
	Verb      string                  // Empty unless the Verb is overridden for the operation
	Endpoint  *endpoints.EndpointInfo // nil unless the endpoint is overridden for the operation
}

/*
	Path
		Endpoint (logical)
		Operations
			GET
				Accepted
				Verb (overridden)
				Endpoint (overridden)
			PUT
			POST
			DELETE
*/

/////////////////////////////////////////////////////////////////////////////
// Config

// Config handles configuration of url handling
type Config struct {
	// Overrides is keyed on url
	Overrides map[string]PathOverride
	// AdditionalPaths contains extra paths to include in the generated hierarchy
	AdditionalPaths []AdditionalPath
	// SuppressAPIVersion true to prevent the api version querystring
	SuppressAPIVersion bool
}

// PathOverride captures Path and/or Verb overrides
type PathOverride struct {
	Path    string // actual url to use
	GetVerb string // Verb to use for logical GET requests
}

// AdditionalPath provides metadata for additional paths to inject into the generated hierarchy
type AdditionalPath struct {
	// Name is the name to use for the generated path
	Name string
	// Path is the path to inject in the hierarchy
	Path string
	// GetPath allows the actual path used at runtime to be overridden
	GetPath string
	// FixedContent provides static content to render in place of making an API call
	FixedContent string
	// SubPathRegesx holds regex info for modifying subpath URLs
	SubPathRegex *RegexReplace
}
