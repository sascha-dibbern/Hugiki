package fragment

import (
	"html/template"
	"net/http"
)

var HelpModeTemplate = MenuTemplate + `
{{- define "helpmode" -}}
{{ template "menu" .Menu }}<!-- helpmode -->
<h1 class="hi-title">Help</h1>
{{ end}}
{{ template "helpmode" . }}
`

type HelpStruct struct {
	Menu MenuItems
}

func RenderHelpMode(writer http.ResponseWriter, mode HelpStruct) {
	tpl := template.Must(template.New("ModeTemplate").Parse(HelpModeTemplate))
	if err := tpl.Execute(writer, mode); err != nil {
		panic(err)
	}
}
