package swagger

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

// MergeSwaggerDoc merges api endpoints from the specified swagger doc into the Paths array
func MergeSwaggerDoc(currentPaths []*Path, config *Config, doc *loads.Document, validateCapturedSegments bool, pathPrefix string) ([]*Path, error) {
	allPaths, err := GetPathsFromSwagger(doc, config, pathPrefix)
	if err != nil {
		empty := []*Path{}
		return empty, err // TODO add context to errors!
	}
	return MergeSwaggerPaths(currentPaths, config, allPaths, validateCapturedSegments, pathPrefix)
}

// MergeSwaggerPaths merges api endpoints into the currentPaths array
func MergeSwaggerPaths(currentPaths []*Path, config *Config, newPaths []Path, validateCapturedSegments bool, pathPrefix string) ([]*Path, error) {
	allPaths, err := addConfigPaths(newPaths, config)
	if err != nil {
		empty := []*Path{}
		return empty, err
	}

	allPaths = getSortedPaths(allPaths)

	resultPaths := currentPaths

	for _, path := range allPaths {
		loopPath := path

		// Add endpoint to paths
		parent := findDeepestPath(resultPaths, loopPath, !validateCapturedSegments)
		if parent == nil {
			resultPaths = append(resultPaths, &loopPath)
		} else {
			if parent.Endpoint.TemplateURL == loopPath.Endpoint.TemplateURL {
				// we have multiple entries with the same path (e.g. when applying a URL override)
				// merge the two entries
				// TODO Consider checking if there is a clash when merging operations
				if loopPath.Operations.Get.Permitted {
					copyOperationFrom(loopPath.Operations.Get, &parent.Operations.Get)
				}
				if loopPath.Operations.Delete.Permitted {
					copyOperationFrom(loopPath.Operations.Delete, &parent.Operations.Delete)
				}
				if loopPath.Operations.Patch.Permitted {
					copyOperationFrom(loopPath.Operations.Patch, &parent.Operations.Patch)
				}
				if loopPath.Operations.Post.Permitted {
					copyOperationFrom(loopPath.Operations.Post, &parent.Operations.Post)
				}
				if loopPath.Operations.Put.Permitted {
					copyOperationFrom(loopPath.Operations.Put, &parent.Operations.Put)
				}
				parent.Children = append(parent.Children, path.Children...)
				parent.SubPaths = append(parent.SubPaths, path.SubPaths...)
			} else {
				if countNameSegments(parent.Endpoint) == countNameSegments(path.Endpoint) {
					// this is a child
					parent.Children = append(parent.Children, &loopPath)
				} else {
					// this is a sub-resource
					parent.SubPaths = append(parent.SubPaths, &loopPath)
				}
			}
		}
	}
	return resultPaths, nil
}

// ConvertToSwaggerResourceTypes converts the Path array to an array of SwaggerResourceTypes for use with the Swagger expander
func ConvertToSwaggerResourceTypes(paths []*Path) []ResourceType {
	resourceTypes := []ResourceType{}
	for _, path := range paths {
		if path.Operations.Get.Endpoint != nil { // ignore endpoints without a GET
			resourceType := convertToSwaggerResourceType(path)
			resourceTypes = append(resourceTypes, resourceType)
		}
	}
	return resourceTypes
}

func convertToSwaggerResourceType(path *Path) ResourceType {
	resourceType := ResourceType{
		Display:      path.Name,
		Endpoint:     endpoints.MustGetEndpointInfoFromURL(path.Operations.Get.Endpoint.TemplateURL, path.Operations.Get.Endpoint.APIVersion),
		Children:     ConvertToSwaggerResourceTypes(path.Children),
		SubResources: ConvertToSwaggerResourceTypes(path.SubPaths),
		FixedContent: path.FixedContent,
		SubPathRegex: path.SubPathRegex,
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

type pathArraySortByCondensedPath []Path

func (a pathArraySortByCondensedPath) Len() int { return len(a) }
func (a pathArraySortByCondensedPath) Less(i, j int) bool {
	return strings.Compare(a[i].Endpoint.TemplateURL, a[j].Endpoint.TemplateURL) < 0
}
func (a pathArraySortByCondensedPath) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func getSortedPaths(paths []Path) []Path {
	// Sort ignoring names of captured sections (e.g. wibble in `/foo/{wibble}/bar`)

	sort.Sort(pathArraySortByCondensedPath(paths))
	return paths
}

func addConfigPaths(paths []Path, config *Config) ([]Path, error) {
	if config.AdditionalPaths != nil {
		for _, additionalPath := range config.AdditionalPaths {
			path := strings.TrimRight(additionalPath.Path, "/")
			endpoint, err := endpoints.GetEndpointInfoFromURL(path, "")
			if err != nil {
				return []Path{}, err
			}
			newPath := Path{
				Name:                  additionalPath.Name,
				Endpoint:              &endpoint,
				CondensedEndpointPath: stripPathNames(path),
				SubPathRegex:          additionalPath.SubPathRegex,
			}
			if additionalPath.FixedContent != "" {
				newPath.FixedContent = additionalPath.FixedContent
			}
			getOperation := PathOperation{
				Permitted: true,
			}
			if additionalPath.GetPath == "" {
				getOperation.Endpoint = &endpoint
			} else {
				getEndpoint, err := endpoints.GetEndpointInfoFromURL(additionalPath.GetPath, "")
				if err != nil {
					return []Path{}, err
				}
				getOperation.Endpoint = &getEndpoint
			}
			newPath.Operations.Get = getOperation
			paths = append(paths, newPath)
		}
	}
	return paths, nil
}

// GetPathsFromSwagger returns the mapped Paths from the document
func GetPathsFromSwagger(doc *loads.Document, config *Config, pathPrefix string) ([]Path, error) {

	swaggerVersion := doc.Spec().Info.Version
	if config.SuppressAPIVersion {
		swaggerVersion = ""
	}

	spec := doc.Analyzer

	swaggerPaths := spec.AllPaths()
	paths := make([]Path, len(swaggerPaths))

	pathIndex := 0
	for swaggerPath, swaggerPathItem := range swaggerPaths {

		swaggerPath = pathPrefix + swaggerPath
		override := config.Overrides[swaggerPath]

		searchPathTemp := override.Path
		var searchPath string
		if searchPathTemp == "" {
			searchPath = swaggerPath
		} else {
			searchPath = searchPathTemp
		}
		searchPath = strings.TrimRight(searchPath, "/")
		endpoint, err := endpoints.GetEndpointInfoFromURL(searchPath, swaggerVersion) // logical path
		if err != nil {
			return []Path{}, err
		}
		lastSegment := endpoint.URLSegments[len(endpoint.URLSegments)-1]
		name := lastSegment.Match
		if name == "" {
			name = "{" + lastSegment.Name + "}"
		}
		path := Path{
			Endpoint:              &endpoint,
			Name:                  name,
			CondensedEndpointPath: stripPathNames(searchPath),
		}

		getVerb := override.GetVerb
		if getVerb == "" {
			getVerb = "get"
		}

		getOperation, err := getOperationByVerb(&swaggerPathItem, getVerb)
		if err != nil {
			return []Path{}, err
		}
		if getOperation != nil {
			path.Operations.Get.Permitted = true
			if getVerb != "get" {
				path.Operations.Get.Verb = getVerb
			}
			if override.Path == "" || override.RewritePath {
				path.Operations.Get.Endpoint = path.Endpoint
			} else {
				overriddenEndpoint, err := endpoints.GetEndpointInfoFromURL(swaggerPath, swaggerVersion)
				if err != nil {
					return []Path{}, err
				}
				path.Operations.Get.Endpoint = &overriddenEndpoint
			}
		}
		if swaggerPathItem.Delete != nil && getVerb != "delete" {
			path.Operations.Delete.Permitted = true
			path.Operations.Delete.Endpoint = path.Endpoint
		}
		if override.DeletePath != "" {
			path.Operations.Delete.Permitted = true
			path.Operations.Delete.Endpoint = endpoints.MustGetEndpointInfoFromURL(override.DeletePath, path.Endpoint.APIVersion)
		}
		if swaggerPathItem.Patch != nil && getVerb != "patch" {
			path.Operations.Patch.Permitted = true
			path.Operations.Patch.Endpoint = path.Endpoint
		}
		if swaggerPathItem.Post != nil && getVerb != "post" {
			path.Operations.Post.Permitted = true
			path.Operations.Post.Endpoint = path.Endpoint
		}
		if swaggerPathItem.Put != nil && getVerb != "put" {
			path.Operations.Put.Permitted = true
			path.Operations.Put.Endpoint = path.Endpoint
		}
		if override.PutPath != "" {
			path.Operations.Put.Permitted = true
			path.Operations.Put.Endpoint = endpoints.MustGetEndpointInfoFromURL(override.PutPath, path.Endpoint.APIVersion)
		}

		paths[pathIndex] = path
		pathIndex++
	}

	return paths, nil
}

var regexpStripNames *regexp.Regexp

func stripPathNames(path string) string {
	if regexpStripNames == nil {
		var err error
		regexpStripNames, err = regexp.Compile(`\{[^}]*}`)
		if err != nil {
			panic(err)
		}
	}
	return regexpStripNames.ReplaceAllString(path, "{}")

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
func findDeepestPath(paths []*Path, pathToFind Path, useStrippedNamesPath bool) *Path {
	for _, path := range paths {
		var matchString string
		var pathToFindString string
		if useStrippedNamesPath {
			matchString = path.CondensedEndpointPath
			pathToFindString = pathToFind.CondensedEndpointPath
		} else {
			matchString = path.Endpoint.TemplateURL
			pathToFindString = pathToFind.Endpoint.TemplateURL
		}

		// Test if matchString is a prefix match on pathToFindString
		// But need to verify that it hasn't matched /example against /examples but does against /example/test
		// ok if strings are equal or substring match with a slash on
		if pathToFindString == matchString ||
			(strings.HasPrefix(pathToFindString, matchString) &&
				(pathToFindString[len(matchString)] == '/' ||
					pathToFindString[len(matchString)] == '(')) {

			// matches endpoint. Check children
			match := findDeepestPath(path.Children, pathToFind, useStrippedNamesPath)
			if match == nil {
				match = findDeepestPath(path.SubPaths, pathToFind, useStrippedNamesPath)
				if match == nil {
					return path
				}
			}
			return match
		}
	}
	return nil
}
