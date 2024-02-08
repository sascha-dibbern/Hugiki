package fragment

import (
	"html/template"
	"net/http"
)

var StatusModeTemplate = MenuTemplate + `
{{- define "statusmode" -}}
{{ template "menu" .Menu }}
<!-- statusmode -->
<h1 class="hi-title">Status</h1>
{{ end}}
`

type StatusStruct struct {
	Menu MenuItems
}

func RenderStatusMode(writer http.ResponseWriter, statusstate StatusStruct, modeimplementation string) {
	if modeimplementation == "" {
		modeimplementation = `{{- template "statusmode" . -}}`
	}
	tpl := template.Must(template.New("ModeTemplate").Parse(StatusModeTemplate + modeimplementation))
	if err := tpl.Execute(writer, statusstate); err != nil {
		panic(err)
	}
}
