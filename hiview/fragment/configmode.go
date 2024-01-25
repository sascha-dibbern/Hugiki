package fragment

import (
	"html/template"
	"net/http"
)

var ConfigModeTemplate = MenuTemplate + `
{{- define "configmode" -}}
{{ template "menu" .Menu }}<!-- configmode -->
<h1>Configuration</h1>
{{ end}}
{{ template "configmode" . }}
`

type ConfigStruct struct {
	Menu MenuItems
}

func RenderConfigMode(writer http.ResponseWriter, mode ConfigStruct) {
	tpl := template.Must(template.New("ModeTemplate").Parse(ConfigModeTemplate))
	if err := tpl.Execute(writer, mode); err != nil {
		panic(err)
	}
}
