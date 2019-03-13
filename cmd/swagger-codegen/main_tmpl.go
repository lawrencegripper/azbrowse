package main

const tmpl = `
{{ define "Path" -}}{{if .Operations.Get.Permitted}}
{ 
	Display: "{{ .Name}}",
	Endpoint: mustGetEndpointInfoFromURL("{{ .Operations.Get.Endpoint.TemplateURL }}", "{{ .Operations.Get.Endpoint.APIVersion}}"),
	{{- if ne .Operations.Get.Verb "" }}
	Verb: "{{upper .Operations.Get.Verb }}",{{end}}
	{{- if .Operations.Delete.Permitted }}
	DeleteEndpoint: mustGetEndpointInfoFromURL("{{ .Operations.Delete.Endpoint.TemplateURL }}", "{{ .Operations.Delete.Endpoint.APIVersion}}"),{{end}}
	{{- if .Operations.Patch.Permitted }}
	PatchEndpoint: mustGetEndpointInfoFromURL("{{ .Operations.Patch.Endpoint.TemplateURL }}", "{{ .Operations.Patch.Endpoint.APIVersion}}"),{{end}}
	{{- if .Operations.Put.Permitted }}
	PutEndpoint: mustGetEndpointInfoFromURL("{{ .Operations.Put.Endpoint.TemplateURL }}", "{{ .Operations.Put.Endpoint.APIVersion}}"),{{end}}
	{{- if .Children}}
	Children: {{template "PathList" .Children}},{{end}}
	{{- if .SubPaths}}
	SubResources: {{template "PathList" .SubPaths}},{{end}}
},{{end }}{{ end }}
{{define "PathList"}}[]SwaggerResourceType{ {{range .}}{{template "Path" .}}{{end}} } {{end}}
package handlers

func (e *SwaggerResourceExpander) getResourceTypes() []SwaggerResourceType {
	return  {{template "PathList" .Paths }}

}
`
