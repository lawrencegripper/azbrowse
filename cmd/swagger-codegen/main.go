package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

func showUsage() {
	fmt.Println("swagger-codegen")
	fmt.Println("===============")
	fmt.Println("")
	flag.Usage()
}

// The input folder structure is as below
// The bash script that generates this ensures that there is only a single version
// spec folder for each resource type. It is most likely to be `stable`, but it could be
// `preview` if no `stable` version exists for that type
//
//  swagger-specs
//   |- top-level
//          |-service1 (e.g. `cdn` or `compute`)
//          |   |-common   (want these)
//          |   |-quickstart-templates
//          |   |-data-plane
//          |   |-resource-manager (we're only interested in the contents of this folder)
//          |       |- resource-type1 (e.g. `Microsoft.Compute`)
//          |       |    |- common
//          |       |    |   |- *.json (want these)
//          |       |    |- stable (NB - may preview if no stable)
//          |       |    |    |- 2018-10-01
//          |       |    |        |- *.json   (want these)
//          |       |- misc files (e.g. readme)
//           ...

func main() {
	outputFileFlag := flag.String("output-file", "", "path to the file to output the generated code to")
	flag.Parse()
	if *outputFileFlag == "" {
		showUsage()
		return
	}

	config := getConfig()
	var paths []*swagger.Path

	serviceFileInfos, err := ioutil.ReadDir("swagger-specs")
	if err != nil {
		panic(err)
	}
	for _, serviceFileInfo := range serviceFileInfos {
		if serviceFileInfo.IsDir() && serviceFileInfo.Name() != "common-types" {
			fmt.Printf("Processing service folder: %s\n", serviceFileInfo.Name())
			// TODO - handle service folder/common folder
			resourceTypeFileInfos, err := ioutil.ReadDir(fmt.Sprintf("swagger-specs/%s/resource-manager", serviceFileInfo.Name()))
			if err != nil {
				panic(err)
			}
			for _, resourceTypeFileInfo := range resourceTypeFileInfos {
				if resourceTypeFileInfo.IsDir() && resourceTypeFileInfo.Name() != "common" {
					// TODO handle common
					swaggerPath := getFirstNonCommonPath(getFirstNonCommonPath(fmt.Sprintf("swagger-specs/%s/resource-manager/%s", serviceFileInfo.Name(), resourceTypeFileInfo.Name())))
					swaggerFileInfos, err := ioutil.ReadDir(swaggerPath)
					if err != nil {
						panic(err)
					}
					for _, swaggerFileInfo := range swaggerFileInfos {
						if !swaggerFileInfo.IsDir() && strings.HasSuffix(swaggerFileInfo.Name(), ".json") {
							fmt.Printf("\tprocessing %s/%s\n", swaggerPath, swaggerFileInfo.Name())
							doc := loadDoc(swaggerPath + "/" + swaggerFileInfo.Name())
							paths, err = swagger.MergeSwaggerDoc(paths, &config, doc, true)
							if err != nil {
								panic(err)
							}
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

// getFirstNonCommonPath returns the first subfolder under path that is not named 'common'
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

func getConfig() swagger.Config {
	config := swagger.Config{
		Overrides: map[string]swagger.PathOverride{
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

func writeTemplate(w io.Writer, paths []*swagger.Path, config *swagger.Config) {

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}
	t := template.Must(template.New("code-gen").Funcs(funcMap).Parse(tmpl))

	type Context struct {
		Paths []*swagger.Path
	}

	context := Context{
		Paths: paths,
	}

	err := t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func loadDoc(path string) *loads.Document {

	document, err := loads.Spec(path)
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	document, err = document.Expanded(&spec.ExpandOptions{RelativeBase: path})
	if err != nil {
		log.Panicf("Error expanding Swagger: %s", err)
	}

	return document
}
