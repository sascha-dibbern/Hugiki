package fragment

import (
	"html/template"
	"net/http"
)

var NavigationModeTemplate = MenuTemplate + `
{{- define "navigationmode" -}}
{{ template "menu" .Menu }}<!-- navigationmode -->
<h1>Navigation</h1>
{{ end}}
{{ template "navigationmode" . }}
`

type NavigationStruct struct {
	Menu MenuItems
}

func RenderNavigationMode(writer http.ResponseWriter, mode NavigationStruct) {
	tpl := template.Must(template.New("ModeTemplate").Parse(NavigationModeTemplate))
	if err := tpl.Execute(writer, mode); err != nil {
		panic(err)
	}
}
