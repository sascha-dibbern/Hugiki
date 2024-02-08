package fragment

import (
	"html/template"
	"net/http"
)

var GitModeTemplate = MenuTemplate + `
{{- define "gitmode" -}}
{{ template "menu" .Menu }}<!-- gitmode -->
<h1 class="hi-title">Git</h1>
{{ end}}
{{ template "gitmode" . }}
`

type GitStruct struct {
	Menu MenuItems
}

func RenderGitMode(writer http.ResponseWriter, mode GitStruct) {
	tpl := template.Must(template.New("ModeTemplate").Parse(GitModeTemplate))
	if err := tpl.Execute(writer, mode); err != nil {
		panic(err)
	}
}
