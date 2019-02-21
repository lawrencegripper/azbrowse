package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
)

// Path represents an URL to output
type Path struct {
	// TODO Name
	Verbs    map[string]SwaggerPathVerb // TODO - create a new type that is easier to work with
	Endpoint *endpoints.EndpointInfo
	Children []*Path
	SubPaths []*Path
}

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

type swaggerOverride struct {
	TreatAsPath string
	TreatAsVerb string
}
type Config struct {
	// Path overrides is keyed on actual path. value is the logical path to use
	PathOverrides map[string]string
	// GetOverrides is keyed on url. value is the verb to use as a logical GET for the path
	GetOverrides map[string]string
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

	topFileInfos, err := ioutil.ReadDir("swagger-specs")
	if err != nil {
		panic(fmt.Errorf)
	}
	for _, topFileInfo := range topFileInfos {
		if topFileInfo.IsDir() {
			fmt.Printf("Processing folder: %s\n", topFileInfo.Name())
			fileInfos, err := ioutil.ReadDir("swagger-specs/" + topFileInfo.Name())
			if err != nil {
				panic(fmt.Errorf)
			}
			for _, fileInfo := range fileInfos {
				if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".json") {
					fmt.Printf("\tprocessing %s\n", fileInfo.Name())
					doc := loadDoc("swagger-specs/" + topFileInfo.Name() + "/" + fileInfo.Name())
					paths = mergeSwaggerDoc(paths, &config, &doc)
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

func getConfig() Config {

	pathOverrides := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list":           "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list":          "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list":  "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list":                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list":     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list":              "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list":          "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings",
	}
	getOverrides := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list":           "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list":          "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list":  "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list":                "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list":     "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list":              "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list": "post",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list":          "post",
	}
	config := Config{
		PathOverrides: pathOverrides,
		GetOverrides:  getOverrides,
	}
	return config
}
func mergeSwaggerDoc(paths []*Path, config *Config, doc *SwaggerDoc) []*Path {
	swaggerPaths := getSortedPaths(doc)
	for _, swaggerPath := range swaggerPaths {
		searchPath := config.PathOverrides[swaggerPath]
		if searchPath == "" {
			searchPath = swaggerPath
		}

		parent := findDeepestPath(paths, searchPath)
		endpoint, err := endpoints.GetEndpointInfoFromURL(swaggerPath, doc.Info.Version)
		if err != nil {
			panic(err)
		}
		path := Path{
			Endpoint: &endpoint,
			Verbs:    doc.Paths[swaggerPath],
		}
		if parent == nil {
			paths = append(paths, &path)
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
	return paths
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
	w.Write([]byte("package handlers\n"))
	w.Write([]byte("\n"))
	w.Write([]byte("func (e *SwaggerResourceExpander) getResourceTypes() []ResourceType {\n"))
	w.Write([]byte("\treturn []ResourceType{\n"))
}
func writeFooter(w io.Writer) {
	w.Write([]byte("\t}\n"))
	w.Write([]byte("}\n"))
}
func writePaths(w io.Writer, paths []*Path, config *Config, prefix string) {
	for _, path := range paths {
		pathVerb := config.GetOverrides[path.Endpoint.TemplateURL]
		if pathVerb == "" {
			pathVerb = "get"
		}
		if path.Verbs[pathVerb].OperationID != "" {
			fmt.Fprintf(w, "\t%s\tResourceType{\n", prefix)
			endpointForName := path.Endpoint
			if pathOverride := config.PathOverrides[path.Endpoint.TemplateURL]; pathOverride != "" {
				endpointOverride, err := endpoints.GetEndpointInfoFromURL(pathOverride, "")
				if err != nil {
					panic(fmt.Errorf("Error parsing pathOverride '%s': %s", pathOverride, err))
				}
				endpointForName = &endpointOverride
			}
			lastSegment := endpointForName.URLSegments[len(endpointForName.URLSegments)-1]
			name := lastSegment.Match
			if name == "" {
				name = "{" + lastSegment.Name + "}"
			}
			fmt.Fprintf(w, "\t%s\t\tDisplay:  \"%s\",\n", prefix, name)
			fmt.Fprintf(w, "\t%s\t\tEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, path.Endpoint.TemplateURL, path.Endpoint.APIVersion)
			if pathVerb != "get" {
				fmt.Fprintf(w, "\t%s\t\tVerb:     \"%s\"", prefix, strings.ToUpper(pathVerb))
			}
			if len(path.Children) > 0 {
				fmt.Fprintf(w, "\t%s\t\tChildren: []ResourceType {\n", prefix)
				writePaths(w, path.Children, config, prefix+"\t\t")
				fmt.Fprintf(w, "\t%s\t\t},\n", prefix)
			}
			if len(path.SubPaths) > 0 {
				fmt.Fprintf(w, "\t%s\t\tSubResources: []ResourceType {\n", prefix)
				writePaths(w, path.SubPaths, config, prefix+"\t\t")
				fmt.Fprintf(w, "\t%s\t\t},\n", prefix)
			}
			fmt.Fprintf(w, "\t%s\t},\n", prefix)
		}
	}
}
func dumpPaths(w io.Writer, paths []*Path, prefix string) {

	for _, path := range paths {
		fmt.Fprintf(w, "%s%s\n", prefix, path.Endpoint.TemplateURL)
		for verb, verbInfo := range path.Verbs {
			fmt.Printf("%s   - %v\t%v\n", prefix, verb, verbInfo.OperationID)
		}
		fmt.Printf("%s   * Children\n", prefix)
		dumpPaths(w, path.Children, prefix+"    ")
		fmt.Printf("%s   * SubPaths\n", prefix)
		dumpPaths(w, path.SubPaths, prefix+"    ")
	}
}

// findDeepestPath searches the endpoints tree to find the deepest point that the specified path can be nested at (used to build up the endpoint hierarchy)
func findDeepestPath(paths []*Path, pathString string) *Path {
	for _, path := range paths {
		if strings.HasPrefix(pathString, path.Endpoint.TemplateURL) &&
			pathString != path.Endpoint.TemplateURL { // short-circuit if we're overriding the path and we have a match. Temporary approach to put items on same parent (until we handle merging and tracking original urls)
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
	swaggerBuf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	var doc SwaggerDoc
	err = json.Unmarshal(swaggerBuf, &doc)
	if err != nil {
		log.Panicf("Error unmarshaling json: %s", err)
	}
	return doc
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
