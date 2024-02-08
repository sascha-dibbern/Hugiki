package fragment

import (
	"html/template"
	"net/http"
)

var StartModeTemplate = MenuTemplate + `
{{- define "startmode" -}}
{{ template "menu" .Menu }}
{{ end}}
`

type StartStruct struct {
	Menu MenuItems
}

func RenderStartMode(writer http.ResponseWriter, startstate StartStruct, modeimplementation string) {
	if modeimplementation == "" {
		modeimplementation = `{{- template "startmode" . -}}`
	}
	tpl := template.Must(template.New("ModeTemplate").Parse(StartModeTemplate + modeimplementation))
	if err := tpl.Execute(writer, startstate); err != nil {
		panic(err)
	}
}
