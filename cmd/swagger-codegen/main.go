package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

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
// Swagger parsing models

// SwaggerDoc is a type to represent the relevant parts of a swagger document
type SwaggerDoc struct {
	Swagger string                      `json:"swagger"`
	Info    SwaggerInfo                 `json:"info"`
	Host    string                      `json:"host"`
	Paths   map[string]SwaggerPathVerbs `json:"paths"`
}

//SwaggerInfo represents the info for a swagger document
type SwaggerInfo struct {
	Version string `json:"version"`
	Title   string `json:"title"`
}

// SwaggerPathVerbs are verbs keyed on path
type SwaggerPathVerbs map[string]SwaggerPathVerb

// SwaggerPathVerb contains the properties for a verb + endpoint
type SwaggerPathVerb struct {
	OperationID string `json:"operationId"`
}

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
							doc2 := loadDoc2(swaggerPath + "/" + swaggerFileInfo.Name())
							// TODO merge!
							_ = doc2
							// doc := loadDoc("swagger-specs/" + serviceFileInfo.Name() + "/" + resourceTypeFileInfo.Name() + "/" + swaggerFileInfo.Name())
							// paths = mergeSwaggerDoc(paths, &config, &doc)
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

	writeHeader(writer)
	writePaths(writer, paths, &config, "")
	writeFooter(writer)
	// dumpPaths(writer, paths, "")
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
func mergeSwaggerDoc(paths []*Path, config *Config, doc *SwaggerDoc) []*Path {
	swaggerPaths := getSortedPaths(doc)
	for _, swaggerPath := range swaggerPaths {
		override := config.Overrides[swaggerPath]

		searchPath := override.Path
		if searchPath == "" {
			searchPath = swaggerPath
		}

		endpoint, err := endpoints.GetEndpointInfoFromURL(searchPath, doc.Info.Version) // logical path
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
		if doc.Paths[swaggerPath][getVerb].OperationID != "" {
			path.Operations.Get.Permitted = true
			if getVerb != "get" {
				path.Operations.Get.Verb = getVerb
			}
			if override.Path != "" {
				overriddenEndpoint, err := endpoints.GetEndpointInfoFromURL(swaggerPath, doc.Info.Version)
				if err != nil {
					panic(err)
				}
				path.Operations.Get.Endpoint = &overriddenEndpoint
			}
		}
		if doc.Paths[swaggerPath]["delete"].OperationID != "" && getVerb != "delete" {
			path.Operations.Delete.Permitted = true
		}
		if doc.Paths[swaggerPath]["patch"].OperationID != "" && getVerb != "patch" {
			path.Operations.Patch.Permitted = true
		}
		if doc.Paths[swaggerPath]["post"].OperationID != "" && getVerb != "post" {
			path.Operations.Post.Permitted = true
		}
		if doc.Paths[swaggerPath]["put"].OperationID != "" && getVerb != "put" {
			path.Operations.Put.Permitted = true
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

func writeHeader(w io.Writer) {
	w.Write([]byte("package handlers\n"))                                                             //nolint: errcheck
	w.Write([]byte("\n"))                                                                             //nolint: errcheck
	w.Write([]byte("func (e *SwaggerResourceExpander) getResourceTypes() []SwaggerResourceType {\n")) //nolint: errcheck
	w.Write([]byte("\treturn []SwaggerResourceType{\n"))                                              //nolint: errcheck
}
func writeFooter(w io.Writer) {
	w.Write([]byte("\t}\n")) //nolint: errcheck
	w.Write([]byte("}\n"))   //nolint: errcheck
}
func writePaths(w io.Writer, paths []*Path, config *Config, prefix string) { // TODO want to not need config here
	for _, path := range paths {
		var pathVerb string //nolint: gosimple
		pathVerb = config.Overrides[path.Endpoint.TemplateURL].GetVerb
		if pathVerb == "" {
			pathVerb = "get"
		}
		if path.Operations.Get.Permitted {
			fmt.Fprintf(w, "\t%s\tSwaggerResourceType{\n", prefix)
			getEndpoint := path.Operations.Get.Endpoint
			if getEndpoint == nil {
				getEndpoint = path.Endpoint
			}
			fmt.Fprintf(w, "\t%s\t\tDisplay:  \"%s\",\n", prefix, path.Name)

			fmt.Fprintf(w, "\t%s\t\tEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, getEndpoint.TemplateURL, getEndpoint.APIVersion)

			if path.Operations.Get.Verb != "" {
				fmt.Fprintf(w, "\t%s\t\tVerb:     \"%s\",\n", prefix, strings.ToUpper(path.Operations.Get.Verb))
			}
			if path.Operations.Delete.Permitted {
				deleteEndpoint := path.Endpoint
				if path.Operations.Delete.Endpoint != nil {
					deleteEndpoint = path.Operations.Delete.Endpoint
				}
				fmt.Fprintf(w, "\t%s\t\tDeleteEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, deleteEndpoint.TemplateURL, deleteEndpoint.APIVersion)
			}
			if path.Operations.Patch.Permitted {
				patchEndpoint := path.Endpoint
				if path.Operations.Patch.Endpoint != nil {
					patchEndpoint = path.Operations.Delete.Endpoint
				}
				fmt.Fprintf(w, "\t%s\t\tPatchEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, patchEndpoint.TemplateURL, patchEndpoint.APIVersion)
			}
			if path.Operations.Put.Permitted {
				putEndpoint := path.Endpoint
				if path.Operations.Put.Endpoint != nil {
					putEndpoint = path.Operations.Put.Endpoint
				}
				fmt.Fprintf(w, "\t%s\t\tPutEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, putEndpoint.TemplateURL, putEndpoint.APIVersion)
			}
			if len(path.Children) > 0 {
				fmt.Fprintf(w, "\t%s\t\tChildren: []SwaggerResourceType {\n", prefix)
				writePaths(w, path.Children, config, prefix+"\t\t")
				fmt.Fprintf(w, "\t%s\t\t},\n", prefix)
			}
			if len(path.SubPaths) > 0 {
				fmt.Fprintf(w, "\t%s\t\tSubResources: []SwaggerResourceType {\n", prefix)
				writePaths(w, path.SubPaths, config, prefix+"\t\t")
				fmt.Fprintf(w, "\t%s\t\t},\n", prefix)
			}
			fmt.Fprintf(w, "\t%s\t},\n", prefix)
		}
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
func loadDoc(path string) SwaggerDoc {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	reader := bufio.NewReader(file)

	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if rune != '\uFEFF' {
		reader.UnreadRune() // Not a BOM -- put the rune back
	}

	var doc SwaggerDoc
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&doc)

	// err = json.Unmarshal(swaggerBuf, &doc)
	if err != nil {
		log.Panicf("Error unmarshaling json: %s", err)
	}
	return doc
}
func loadDoc2(path string) *loads.Document {

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
func getSortedPaths(doc *SwaggerDoc) []string {
	paths := make([]string, len(doc.Paths))
	i := 0
	for key := range doc.Paths {
		paths[i] = key
		i++
	}
	sort.Strings(paths)
	return paths
}
