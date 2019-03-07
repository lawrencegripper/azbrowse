package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

/////////////////////////////////////////////////////////////////////////////
// Path models

// Path represents a path that we want to consider emitting in code-gen. It is derived from
type Path struct {
	Name       string
	Endpoint   *endpoints.EndpointInfo // The logical endpoint. May be overridden for an operation
	Operations PathOperations
	Children   []*Path
	SubPaths   []*Path
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
	// Overrides is keyed on url and
	Overrides map[string]SwaggerPathOverride
}

// SwaggerPathOverride captures Path and/or Verb overrides
type SwaggerPathOverride struct {
	Path    string // actual url to use
	GetVerb string // Verb to use for logical GET requests
}

func showUsage() {
	fmt.Println("swagger-codegen")
	fmt.Println("===============")
	fmt.Println("")
	flag.Usage()
}
func main() {
	outputFileFlag := flag.String("output-file", "", "path to the file to output the generated code to")
	flag.Parse()
	if *outputFileFlag == "" {
		showUsage()
		return
	}

	config := getConfig()
	var paths []*Path

	serviceFileInfos, err := ioutil.ReadDir("swagger-specs/top-level")
	if err != nil {
		panic(err)
	}
	for _, serviceFileInfo := range serviceFileInfos {
		if serviceFileInfo.IsDir() {
			fmt.Printf("Processing service folder: %s\n", serviceFileInfo.Name())
			resourceTypeFileInfos, err := ioutil.ReadDir("swagger-specs/top-level/" + serviceFileInfo.Name())
			if err != nil {
				panic(err)
			}
			for _, resourceTypeFileInfo := range resourceTypeFileInfos {
				if resourceTypeFileInfo.IsDir() && resourceTypeFileInfo.Name() != "common" {
					swaggerPath := getFirstNonCommonPath(getFirstNonCommonPath("swagger-specs/top-level/" + serviceFileInfo.Name() + "/" + resourceTypeFileInfo.Name()))
					swaggerFileInfos, err := ioutil.ReadDir(swaggerPath)
					if err != nil {
						panic(err)
					}
					for _, swaggerFileInfo := range swaggerFileInfos {
						if !swaggerFileInfo.IsDir() && strings.HasSuffix(swaggerFileInfo.Name(), ".json") {
							fmt.Printf("\tprocessing %s/%s\n", swaggerPath, swaggerFileInfo.Name())
							doc := loadDoc(swaggerPath + "/" + swaggerFileInfo.Name())
							paths = mergeSwaggerDoc(paths, &config, doc)
						}
					}
				}
			}
		}
	}

	writer, err := os.Create(*outputFileFlag)
	if err != nil {
		panic(fmt.Errorf("Error opening file: %s", err))
	}
	defer func() {
		err := writer.Close()
		if err != nil {
			panic(fmt.Errorf("Failed to close output file: %s", err))
		}
	}()

	writeTemplate(writer, paths, &config)
}

func getFirstNonCommonPath(path string) string {
	// get the first non `common` path

	subfolders, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, subpath := range subfolders {
		if subpath.IsDir() && subpath.Name() != "common" {
			return path + "/" + subpath.Name()
		}
	}
	panic(fmt.Errorf("No suitable path found"))
}

func getConfig() Config {
	config := Config{
		Overrides: map[string]SwaggerPathOverride{
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings",
				GetVerb: "post",
			},
		},
	}
	return config
}
func mergeSwaggerDoc(paths []*Path, config *Config, doc *loads.Document) []*Path {
	swaggerVersion := doc.Spec().Info.Version
	spec := doc.Analyzer
	allPaths := spec.AllPaths()
	swaggerPaths := getSortedPaths(spec)
	for _, swaggerPath := range swaggerPaths {
		override := config.Overrides[swaggerPath]

		searchPath := override.Path
		if searchPath == "" {
			searchPath = swaggerPath
		}
		endpoint, err := endpoints.GetEndpointInfoFromURL(searchPath, swaggerVersion) // logical path
		if err != nil {
			panic(err)
		}
		lastSegment := endpoint.URLSegments[len(endpoint.URLSegments)-1]
		name := lastSegment.Match
		if name == "" {
			name = "{" + lastSegment.Name + "}"
		}
		path := Path{
			Endpoint: &endpoint,
			Name:     name,
		}

		getVerb := override.GetVerb
		if getVerb == "" {
			getVerb = "get"
		}
		pathItem := allPaths[swaggerPath]
		getOperation := getOperationByVerb(&pathItem, getVerb)
		if getOperation != nil {
			path.Operations.Get.Permitted = true
			if getVerb != "get" {
				path.Operations.Get.Verb = getVerb
			}
			if override.Path == "" {
				path.Operations.Get.Endpoint = path.Endpoint
			} else {
				overriddenEndpoint, err := endpoints.GetEndpointInfoFromURL(swaggerPath, swaggerVersion)
				if err != nil {
					panic(err)
				}
				path.Operations.Get.Endpoint = &overriddenEndpoint
			}
		}
		if allPaths[swaggerPath].Delete != nil && getVerb != "delete" {
			path.Operations.Delete.Permitted = true
			path.Operations.Delete.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath].Patch != nil && getVerb != "patch" {
			path.Operations.Patch.Permitted = true
			path.Operations.Patch.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath].Post != nil && getVerb != "post" {
			path.Operations.Post.Permitted = true
			path.Operations.Post.Endpoint = path.Endpoint
		}
		if allPaths[swaggerPath].Put != nil && getVerb != "put" {
			path.Operations.Put.Permitted = true
			path.Operations.Put.Endpoint = path.Endpoint
		}

		// Add endpoint to paths
		parent := findDeepestPath(paths, searchPath)
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
	return paths
}
func getOperationByVerb(pathItem *spec.PathItem, verb string) *spec.Operation {
	switch strings.ToLower(verb) {
	case "get":
		return pathItem.Get
	case "delete":
		return pathItem.Delete
	case "head":
		return pathItem.Head
	case "options":
		return pathItem.Options
	case "patch":
		return pathItem.Patch
	case "post":
		return pathItem.Post
	case "put":
		return pathItem.Put
	default:
		panic(fmt.Errorf("Unhandled verb: %s", verb))
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

func writeTemplate(w io.Writer, paths []*Path, config *Config) {

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}
	t := template.Must(template.New("code-gen").Funcs(funcMap).Parse(tmpl))

	type Context struct {
		Paths []*Path
	}

	context := Context{
		Paths: paths,
	}

	err := t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

// findDeepestPath searches the endpoints tree to find the deepest point that the specified path can be nested at (used to build up the endpoint hierarchy)
func findDeepestPath(paths []*Path, pathString string) *Path {
	for _, path := range paths {
		if strings.HasPrefix(pathString, path.Endpoint.TemplateURL) {
			// matches endpoint. Check children
			match := findDeepestPath(path.Children, pathString)
			if match == nil {
				match = findDeepestPath(path.SubPaths, pathString)
				if match == nil {
					return path
				}
			}
			return match
		}
	}
	return nil
}
func loadDoc(path string) *loads.Document {

	document, err := loads.Spec(path)
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	document, err = document.Expanded(&spec.ExpandOptions{RelativeBase: path})
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	return document
}
func getSortedPaths(spec *analysis.Spec) []string {
	paths := make([]string, len(spec.AllPaths()))
	i := 0
	for key := range spec.AllPaths() {
		paths[i] = key
		i++
	}
	sort.Strings(paths)
	return paths
}
