package swagger

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

func MergeSwaggerDoc(paths []*Path, config *Config, doc *loads.Document, validateCapturedSegments bool) ([]*Path, error) {
	swaggerVersion := doc.Spec().Info.Version
	spec := doc.Analyzer
	allPaths := spec.AllPaths()
	swaggerPaths := getSortedPaths(spec)
	for _, swaggerPath := range swaggerPaths {
		override := config.Overrides[swaggerPath.Path]

		searchPathTemp := override.Path
		var searchPath PathAndCondensedPath
		if searchPathTemp == "" {
			searchPath = PathAndCondensedPath{
				Path:          swaggerPath.Path,
				CondensedPath: swaggerPath.CondensedPath,
			}
		} else {
			searchPath = PathAndCondensedPath{
				Path:          searchPathTemp,
				CondensedPath: condensePath(searchPathTemp), // condense overridden path,
			}
		}
		endpoint, err := endpoints.GetEndpointInfoFromURL(searchPath.Path, swaggerVersion) // logical path
		if err != nil {
			empty := []*Path{}
			return empty, err
		}
		lastSegment := endpoint.URLSegments[len(endpoint.URLSegments)-1]
		name := lastSegment.Match
		if name == "" {
			name = "{" + lastSegment.Name + "}"
		}
		path := Path{
			Endpoint:              &endpoint,
			Name:                  name,
			CondensedEndpointPath: searchPath.CondensedPath,
		}

		getVerb := override.GetVerb
		if getVerb == "" {
			getVerb = "get"
		}
		pathItem := allPaths[swaggerPath.Path]
		getOperation, err := getOperationByVerb(&pathItem, getVerb)
		if err != nil {
			empty := []*Path{}
			return empty, err
		}
		if getOperation != nil {
			path.Operations.Get.Permitted = true
			if getVerb != "get" {
				path.Operations.Get.Verb = getVerb
			}
			if override.Path == "" {
				path.Operations.Get.Endpoint = path.Endpoint
			} else {
				overriddenEndpoint, err := endpoints.GetEndpointInfoFromURL(swaggerPath.Path, swaggerVersion)
				if err != nil {
					empty := []*Path{}
					return empty, err
				}
				path.Operations.Get.Endpoint = &overriddenEndpoint
			}
		}
		if allPaths[swaggerPath.Path].Delete != nil && getVerb != "delete" {
			path.Operations.Delete.Permitted = true
			path.Operations.Delete.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath.Path].Patch != nil && getVerb != "patch" {
			path.Operations.Patch.Permitted = true
			path.Operations.Patch.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath.Path].Post != nil && getVerb != "post" {
			path.Operations.Post.Permitted = true
			path.Operations.Post.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath.Path].Put != nil && getVerb != "put" {
			path.Operations.Put.Permitted = true
			path.Operations.Put.Endpoint = path.Endpoint
		}

		// Add endpoint to paths
		parent := findDeepestPath(paths, searchPath, !validateCapturedSegments)
		if parent == nil {
			paths = append(paths, &path)
		} else {
			if parent.Endpoint.TemplateURL == path.Endpoint.TemplateURL {
				// we have multiple entries with the same path (e.g. when applying a URL override)
				// merge the two entries
				// TODO Consider checking if there is a clash when merging operations
				if path.Operations.Get.Permitted {
					copyOperationFrom(path.Operations.Get, &parent.Operations.Get)
				}
				if path.Operations.Delete.Permitted {
					copyOperationFrom(path.Operations.Delete, &parent.Operations.Delete)
				}
				if path.Operations.Patch.Permitted {
					copyOperationFrom(path.Operations.Patch, &parent.Operations.Patch)
				}
				if path.Operations.Post.Permitted {
					copyOperationFrom(path.Operations.Post, &parent.Operations.Post)
				}
				if path.Operations.Put.Permitted {
					copyOperationFrom(path.Operations.Put, &parent.Operations.Put)
				}
				parent.Children = append(parent.Children, path.Children...)
				parent.SubPaths = append(parent.SubPaths, path.SubPaths...)
			} else {
				if countNameSegments(parent.Endpoint) == countNameSegments(path.Endpoint) {
					// this is a child
					parent.Children = append(parent.Children, &path)
				} else {
					// this is a sub-resource
					parent.SubPaths = append(parent.SubPaths, &path)
				}
			}
		}
	}
	return paths, nil
}

func ConvertToSwaggerResourceTypes(paths []*Path) []SwaggerResourceType {
	resourceTypes := []SwaggerResourceType{}
	for _, path := range paths {
		if path.Operations.Get.Endpoint != nil { // ignore endpoints without a GET
			resourceType := convertToSwaggerResourceType(path)
			resourceTypes = append(resourceTypes, resourceType)
		}
	}
	return resourceTypes
}

func convertToSwaggerResourceType(path *Path) SwaggerResourceType {
	resourceType := SwaggerResourceType{
		Display:      path.Name,
		Endpoint:     endpoints.MustGetEndpointInfoFromURL(path.Operations.Get.Endpoint.TemplateURL, path.Operations.Get.Endpoint.APIVersion),
		Children:     ConvertToSwaggerResourceTypes(path.Children),
		SubResources: ConvertToSwaggerResourceTypes(path.SubPaths),
	}
	if path.Operations.Get.Verb != "" {
		resourceType.Verb = path.Operations.Get.Verb
	}
	if path.Operations.Delete.Permitted {
		resourceType.DeleteEndpoint = endpoints.MustGetEndpointInfoFromURL(path.Operations.Delete.Endpoint.TemplateURL, path.Operations.Delete.Endpoint.APIVersion)
	}
	if path.Operations.Patch.Permitted {
		resourceType.PatchEndpoint = endpoints.MustGetEndpointInfoFromURL(path.Operations.Patch.Endpoint.TemplateURL, path.Operations.Patch.Endpoint.APIVersion)
	}
	if path.Operations.Put.Permitted {
		resourceType.PutEndpoint = endpoints.MustGetEndpointInfoFromURL(path.Operations.Put.Endpoint.TemplateURL, path.Operations.Put.Endpoint.APIVersion)
	}

	return resourceType
}

type PathAndCondensedPath struct {
	Path          string
	CondensedPath string
}
type PathAndCondensedPathList []PathAndCondensedPath

func (a PathAndCondensedPathList) Len() int { return len(a) }
func (a PathAndCondensedPathList) Less(i, j int) bool {
	return strings.Compare(a[i].CondensedPath, a[j].CondensedPath) < 0
}
func (a PathAndCondensedPathList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func getSortedPaths(spec *analysis.Spec) []PathAndCondensedPath {
	// Sort ignoring names of captured sections (e.g. wibble in `/foo/{wibble}/bar`)

	pathPairs := make([]PathAndCondensedPath, len(spec.AllPaths()))
	i := 0
	for key := range spec.AllPaths() {
		pathPairs[i] = PathAndCondensedPath{
			Path:          key,
			CondensedPath: condensePath(key),
		}
		i++
	}
	sort.Sort(PathAndCondensedPathList(pathPairs))

	return pathPairs
}

var regexpCondense *regexp.Regexp

func condensePath(path string) string {
	if regexpCondense == nil {
		regexpCondense, _ = regexp.Compile("\\{[^}]*}") // TODO - handle error
	}
	return regexpCondense.ReplaceAllString(path, "{}")

}

func getOperationByVerb(pathItem *spec.PathItem, verb string) (*spec.Operation, error) {
	switch strings.ToLower(verb) {
	case "get":
		return pathItem.Get, nil
	case "delete":
		return pathItem.Delete, nil
	case "head":
		return pathItem.Head, nil
	case "options":
		return pathItem.Options, nil
	case "patch":
		return pathItem.Patch, nil
	case "post":
		return pathItem.Post, nil
	case "put":
		return pathItem.Put, nil
	default:
		return nil, fmt.Errorf("Unhandled verb: %s", verb)
	}
}

func copyOperationFrom(from PathOperation, to *PathOperation) {
	to.Permitted = from.Permitted
	to.Endpoint = from.Endpoint
	to.Verb = from.Verb
}
func countNameSegments(endpoint *endpoints.EndpointInfo) int {
	count := 0
	for _, segment := range endpoint.URLSegments {
		if segment.Name != "" {
			count++
		}
	}
	return count
}

// findDeepestPath searches the endpoints tree to find the deepest point that the specified path can be nested at (used to build up the endpoint hierarchy)
func findDeepestPath(paths []*Path, pathToFind PathAndCondensedPath, useCondensedPath bool) *Path {
	for _, path := range paths {
		var matchString string
		if useCondensedPath {
			matchString = path.CondensedEndpointPath
		} else {
			matchString = path.Endpoint.TemplateURL
		}

		var pathString string
		if useCondensedPath {
			pathString = pathToFind.CondensedPath
		} else {
			pathString = pathToFind.Path
		}

		if strings.HasPrefix(pathString, matchString) {
			// matches endpoint. Check children
			match := findDeepestPath(path.Children, pathToFind, useCondensedPath)
			if match == nil {
				match = findDeepestPath(path.SubPaths, pathToFind, useCondensedPath)
				if match == nil {
					return path
				}
			}
			return match
		}
	}
	return nil
}
