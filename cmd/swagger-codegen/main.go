package main

import (
	"encoding/json"
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

func main() {
	// TODO - stream input to json decoder rather than loading full doc in memory
	doc := loadDoc()

	swaggerPaths := getSortedPaths(doc)

	var paths []*Path
	for _, swaggerPath := range swaggerPaths {
		parent := findDeepestPath(paths, swaggerPath)
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

	writer := os.Stdout // TODO - take filename as argument to write to

	writeHeader(writer)
	writePaths(writer, paths, "")
	writeFooter(writer)
	// dumpPaths(writer, paths, "")
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
	w.Write([]byte("func (e *AppServiceResourceExpander) getResourceTypes() []ResourceType {\n"))
	w.Write([]byte("\treturn []ResourceType{\n"))

}
func writeFooter(w io.Writer) {
	w.Write([]byte("\t}\n"))
	w.Write([]byte("}\n"))
}
func writePaths(w io.Writer, paths []*Path, prefix string) {
	for _, path := range paths {
		if path.Verbs["get"].OperationID != "" {
			fmt.Fprintf(w, "\t%s\tResourceType{\n", prefix)
			lastSegment := path.Endpoint.URLSegments[len(path.Endpoint.URLSegments)-1]
			name := lastSegment.Match
			if name == "" {
				name = "{" + lastSegment.Name + "}"
			}
			fmt.Fprintf(w, "\t%s\t\tDisplay: \"%s\",\n", prefix, name)
			fmt.Fprintf(w, "\t%s\t\tEndpoint: mustGetEndpointInfoFromURL(\"%s\", \"%s\"),\n", prefix, path.Endpoint.TemplateURL, path.Endpoint.APIVersion)
			if len(path.Children) > 0 {
				fmt.Fprintf(w, "\t%s\t\tChildren: []ResourceType {\n", prefix)
				writePaths(w, path.Children, prefix+"\t\t")
				fmt.Fprintf(w, "\t%s\t\t},\n", prefix)
			}
			if len(path.SubPaths) > 0 {
				fmt.Fprintf(w, "\t%s\t\tSubResources: []ResourceType {\n", prefix)
				writePaths(w, path.SubPaths, prefix+"\t\t")
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
func loadDoc() SwaggerDoc {
	path := "cmd/swagger-codegen/sample-specs/WebApps.json"
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
func getSortedPaths(doc SwaggerDoc) []string {
	paths := make([]string, len(doc.Paths))
	i := 0
	for key := range doc.Paths {
		paths[i] = key
		i++
	}
	sort.Strings(paths)
	return paths
}
