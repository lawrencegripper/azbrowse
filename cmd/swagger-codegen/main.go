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
)

// Endpoint represents an endpoint to output
type Endpoint struct {
	// TODO Name
	Verbs    map[string]SwaggerPathVerb // TODO - create a new type that is easier to work with
	Path     string
	Children []*Endpoint
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

	paths := getSortedPaths(doc)

	var endpoints []*Endpoint
	for _, path := range paths {
		parent := findPath(endpoints, path)
		endpoint := Endpoint{
			Path:  path,
			Verbs: doc.Paths[path],
		}
		if parent == nil {
			endpoints = append(endpoints, &endpoint)
		} else {
			parent.Children = append(parent.Children, &endpoint)
		}
	}

	dumpPaths(os.Stdout, endpoints, "")
}
func dumpPaths(w io.Writer, endpoints []*Endpoint, prefix string) {

	for _, endpoint := range endpoints {
		fmt.Fprintf(w, "%s%s\n", prefix, endpoint.Path)
		for verb, verbInfo := range endpoint.Verbs {
			fmt.Printf("%s   - %v\t%v\n", prefix, verb, verbInfo.OperationID)
		}
		dumpPaths(w, endpoint.Children, prefix+"  ")
	}
}
func findPath(endpoints []*Endpoint, path string) *Endpoint {
	for _, endpoint := range endpoints {
		if strings.HasPrefix(path, endpoint.Path) {
			// matches endpoint. Check children
			match := findPath(endpoint.Children, path)
			if match == nil {
				return endpoint
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
