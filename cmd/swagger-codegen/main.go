package main

import (
	"encoding/json"
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

	fmt.Println("*******************************************")
	fmt.Println("  Processing ARM Specs ")
	fmt.Println("*******************************************")
	config := getARMConfig()
	paths := loadARMSwagger(config)
	writeOutput(paths, config, "./internal/pkg/expanders/swagger-armspecs.generated.go", "SwaggerAPISetARMResources")
	fmt.Println()

	fmt.Println("*******************************************")
	fmt.Println("  Processing Azure Search Data-plane Specs ")
	fmt.Println("*******************************************")
	config = getAzureSearchDataPlaneConfig()
	paths = loadAzureSearchDataPlaneSpecs(config)
	writeOutput(paths, config, "./internal/pkg/expanders/search.generated.go", "AzureSearchServiceExpander")
	fmt.Println()

	fmt.Println("*******************************************")
	fmt.Println("  Processing Databricks Data-plane Specs ")
	fmt.Println("*******************************************")
	config = getDatabricksDataPlaneConfig()
	paths = loadDatabricksDataPlaneSpecs(config)
	writeOutput(paths, config, "./internal/pkg/expanders/databricks.generated.go", "AzureDatabricksExpander")
	fmt.Println()
}

// APISet matches the structure of `api-set.json` files from swagger-update
type APISet struct {
	Name       string   `json:"name"`
	InputFiles []string `json:"input_files"`
}

func print(s string) {
	if os.Getenv("VERBOSE") == "true" {
		fmt.Print(s)
	} else {
		fmt.Print(".")
	}
}

func loadARMSwagger(config *swagger.Config) []*swagger.Path {
	var paths []*swagger.Path
	processAPISet := func(apiSetFolder string) bool {
		apiSetPath := fmt.Sprintf("%s/api-set.json", apiSetFolder)
		buf, err := ioutil.ReadFile(apiSetPath)
		if err != nil {
			return false
		}
		var apiSet APISet
		err = json.Unmarshal(buf, &apiSet)
		if err != nil {
			panic(err)
		}
		// Build up paths for all files in the APISet
		folderPaths := []swagger.Path{}
		for _, inputFile := range apiSet.InputFiles {
			print(fmt.Sprintf("\tprocessing %s/%s\n", apiSetFolder, inputFile))
			doc := loadDoc(apiSetFolder + "/" + inputFile)
			filePaths, err := swagger.GetPathsFromSwagger(doc, config, "")
			if err != nil {
				panic(err)
			}
			folderPaths = append(folderPaths, filePaths...)
		}
		if len(folderPaths) > 0 {
			paths, err = swagger.MergeSwaggerPaths(paths, config, folderPaths, true, "")
			if err != nil {
				panic(err)
			}
			return true
		}
		return false
	}

	resourceProviderFileInfos, err := ioutil.ReadDir("swagger-specs")
	if err != nil {
		panic(err)
	}
	for _, resourceProviderFileInfo := range resourceProviderFileInfos {
		if resourceProviderFileInfo.IsDir() && resourceProviderFileInfo.Name() != "common-types" {
			print(fmt.Sprintf("Processing resource provider folder: %s\n", resourceProviderFileInfo.Name()))
			resourceProviderFolderPath := fmt.Sprintf("swagger-specs/%s/resource-manager", resourceProviderFileInfo.Name())
			resourceTypeFileInfos, err := ioutil.ReadDir(resourceProviderFolderPath)
			_ = resourceTypeFileInfos
			if err != nil {
				continue // may just be data-plane folder
			}
			processed := processAPISet(resourceProviderFolderPath)
			if processed {
				print("Got api-set.json")
			} else {
				// Didn't get an api-set.json in the resource provider - check in resource types
				for _, resourceTypeFileInfo := range resourceTypeFileInfos {
					if resourceTypeFileInfo.IsDir() && resourceTypeFileInfo.Name() != "common" {
						resourceTypeFolderPath := fmt.Sprintf("%s/%s", resourceProviderFolderPath, resourceTypeFileInfo.Name())
						print(fmt.Sprintf("\tProcessing resource type folder: %s\n", resourceTypeFolderPath))
						processed = processAPISet(resourceTypeFolderPath)
						if processed {
							print("Got api-set.json")
						}
					}
				}
			}
		}
	}
	return paths
}

// getARMConfig returns the config for ARM Swagger processing
func getARMConfig() *swagger.Config {
	config := &swagger.Config{
		Overrides: map[string]swagger.PathOverride{
			// App Service patches
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
			// Search patches
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys", // no change to path
				GetVerb: "post",
			},
			// VM Scale Sets patches
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/instanceView": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/instanceView",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/publicipaddresses",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses",
				RewritePath: true,
			},
			// DNS
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}": {
				Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets/{recordType}/{relativeRecordSetName}",
			},
		},
	}
	return config
}

// getARMConfig returns the config for ARM Swagger processing
func getAzureSearchDataPlaneConfig() *swagger.Config {
	return &swagger.Config{
		Overrides: map[string]swagger.PathOverride{
			"/indexes('{indexName}')/docs('{key}')": {
				PutPath:    "/indexes('{indexName}')/docs/index", // update requires POST, but the APISet will handle that
				DeletePath: "/indexes('{indexName}')/docs/index", // update requires POST, but the APISet will handle that
			},
		},
	}
}
func loadAzureSearchDataPlaneSpecs(config *swagger.Config) []*swagger.Path {
	var paths []*swagger.Path

	directoryNames := []string{"Microsoft.Azure.Search.Service", "Microsoft.Azure.Search.Data"} // need to control the document load order
	for _, directoryName := range directoryNames {
		swaggerPath := getFirstNonCommonPath(getFirstNonCommonPath(fmt.Sprintf("swagger-specs/search/data-plane/%s", directoryName)))
		swaggerFileInfos, err := ioutil.ReadDir(swaggerPath)
		if err != nil {
			panic(err)
		}
		for _, swaggerFileInfo := range swaggerFileInfos {
			if !swaggerFileInfo.IsDir() && strings.HasSuffix(swaggerFileInfo.Name(), ".json") {
				print(fmt.Sprintf("\tprocessing %s/%s\n", swaggerPath, swaggerFileInfo.Name()))
				doc := loadDoc(swaggerPath + "/" + swaggerFileInfo.Name())
				pathPrefix := ""
				if swaggerFileInfo.Name() == "searchindex.json" {
					// searchindex.json uses a custom property to set a base URL that the paths in that file are relative to
					// I couldn't find a way to retrieve it with the swagger library so adding some config here
					pathPrefix = "/indexes('{indexName}')"
				}
				paths, err = swagger.MergeSwaggerDoc(paths, config, doc, true, pathPrefix)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return paths
}

// getARMConfig returns the config for ARM Swagger processing
func getDatabricksDataPlaneConfig() *swagger.Config {
	return &swagger.Config{
		SuppressAPIVersion: true,
		AdditionalPaths: []swagger.AdditionalPath{
			// add as a missing path - also overridden to map to the actual endpoint that exists!
			{Name: "{scope}", Path: "/api/2.0/secrets/{scope}", DeletePath: "/api/2.0/secrets/scopes/delete", FixedContent: "Choose a node to expand..."},
			// Add extra point for runs listing
			{Name: "runs", Path: "/api/2.0/runs", GetPath: "/api/2.0/jobs/runs/list"},
			{Name: "{run_id}", Path: "/api/2.0/runs/{run_id}", GetPath: "/api/2.0/jobs/runs/get", DeletePath: "/api/2.0/jobs/runs/delete"},
		},
		// Override the some paths to relate them to each other better
		Overrides: map[string]swagger.PathOverride{
			"/api/2.0/clusters/get": { // push cluster get under cluster list
				Path:       "/api/2.0/clusters/list/{cluster_id}",
				PutPath:    "/api/2.0/clusters/edit",
				DeletePath: "/api/2.0/clusters/permanent-delete",
			},

			"/api/2.0/instance-pools/list": {
				Path: "/api/2.0/instance-pools",
			},
			"/api/2.0/instance-pools/get": {
				Path:       "/api/2.0/instance-pools/{instance_pool_id}",
				DeletePath: "/api/2.0/instance-pools/delete",
			},

			"/api/2.0/jobs/list": { // tweak job list path to be sorted before others!
				Path: "/api/2.0/jobs",
			},
			"/api/2.0/jobs/get": { // push job get under job list
				Path:       "/api/2.0/jobs/{job_id}",
				PutPath:    "/api/2.0/jobs/reset",
				DeletePath: "/api/2.0/jobs/delete",
			},
			"/api/2.0/jobs/runs/list": { // push run list under job get
				Path: "/api/2.0/jobs/{job_id}/runs",
			},
			"/api/2.0/jobs/runs/get": { // push run get under run list
				Path:       "/api/2.0/jobs/{job_id}/runs/{run_id}",
				DeletePath: "/api/2.0/jobs/runs/delete",
			},

			"/api/2.0/secrets/scopes/list": { // Ensure secret scope list is sorted top for secrets urls
				Path: "/api/2.0/secrets",
			},
			"/api/2.0/secrets/list": { // push secrets under secret scopes
				Path: "/api/2.0/secrets/{scope}/secrets",
			},
			"/api/2.0/secrets/acls/list": { // push secret acls under secret scopes
				Path:       "/api/2.0/secrets/{scope}/acls",
				PutPath:    "/api/2.0/secrets/acls/put",
				DeletePath: "/api/2.0/secrets/acls/delete",
			},
			"/api/2.0/secrets/acls/get": { // push secret acls under secret scopes
				Path: "/api/2.0/secrets/{scope}/acls/{principal}",
			},
			"/api/2.0/dbfs/list": {
				Path:       "/api/2.0/dbfs",
				DeletePath: "/api/2.0/dbfs/delete",
			},
			"/api/2.0/workspace/list": {
				Path:       "/api/2.0/workspace",
				DeletePath: "/api/2.0/workspace/delete",
			},
		},
	}
}
func loadDatabricksDataPlaneSpecs(config *swagger.Config) []*swagger.Path {
	var paths []*swagger.Path

	doc := loadDoc("custom-swagger-specs/databricks.json")
	pathPrefix := ""
	paths, err := swagger.MergeSwaggerDoc(paths, config, doc, true, pathPrefix)
	if err != nil {
		panic(err)
	}
	return paths
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
func writeOutput(paths []*swagger.Path, config *swagger.Config, filename string, structName string) {
	writer, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("Error opening file: %s", err))
	}
	defer func() {
		err := writer.Close()
		if err != nil {
			panic(fmt.Errorf("Failed to close output file: %s", err))
		}
	}()

	writeTemplate(writer, paths, config, structName)
}
func writeTemplate(w io.Writer, paths []*swagger.Path, config *swagger.Config, structName string) {

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}
	t := template.Must(template.New("code-gen").Funcs(funcMap).Parse(tmpl))

	type Context struct {
		Paths      []*swagger.Path
		StructName string
	}

	context := Context{
		Paths:      paths,
		StructName: structName,
	}

	err := t.Execute(w, context)
	if err != nil {
		panic(err)
	}
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
