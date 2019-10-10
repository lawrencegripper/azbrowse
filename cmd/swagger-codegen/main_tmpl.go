package main

const tmpl = `
{{ define "Path" -}}{{if .Operations.Get.Permitted}}
{ 
	Display: "{{ .Name}}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("{{ .Operations.Get.Endpoint.TemplateURL }}", "{{ .Operations.Get.Endpoint.APIVersion}}"),
	{{- if ne .Operations.Get.Verb "" }}
	Verb: "{{upper .Operations.Get.Verb }}",{{end}}
	{{- if .Operations.Delete.Permitted }}
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("{{ .Operations.Delete.Endpoint.TemplateURL }}", "{{ .Operations.Delete.Endpoint.APIVersion}}"),{{end}}
	{{- if .Operations.Patch.Permitted }}
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("{{ .Operations.Patch.Endpoint.TemplateURL }}", "{{ .Operations.Patch.Endpoint.APIVersion}}"),{{end}}
	{{- if .Operations.Put.Permitted }}
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("{{ .Operations.Put.Endpoint.TemplateURL }}", "{{ .Operations.Put.Endpoint.APIVersion}}"),{{end}}
	{{- if .Children}}
	Children: {{template "PathList" .Children}},{{end}}
	{{- if .SubPaths}}
	SubResources: {{template "PathList" .SubPaths}},{{end}}
},{{end }}{{ end }}
{{define "PathList"}}[]swagger.SwaggerResourceType{ {{range .}}{{template "Path" .}}{{end}} } {{end}}
package handlers

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"	
	"github.com/lawrencegripper/azbrowse/pkg/swagger"	
)

func (e *SwaggerConfigARMResources) loadResourceTypes() []swagger.SwaggerResourceType {
	return  {{template "PathList" .Paths }}

}
`
