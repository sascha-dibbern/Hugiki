package fragment

import (
	"html/template"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiuri"
)

var SearchModeTemplate = MenuTemplate + `
{{- define "searchmode" -}}
{{ template "menu" .Menu }}
<!-- searchmode -->
<h1 class="hi-title">Search</h1>
<input class="hi-content-search" type="search" 
       name="search" placeholder="Begin Typing To Search content document..." 
       hx-post="` + hiuri.UriAction_SearchContent + `" 
       hx-trigger="input changed delay:500ms, search" 
       hx-target="#search-results"> <!-- TODO add hx-indicator -->

<table class="table">
	<thead>
	<tr>
		<th>Matching files</th>
	</tr>
	</thead>
	<tbody id="search-results">
	</tbody>
</table>
{{ end}}
`

type SearchStruct struct {
	Menu                MenuItems
	ContentSearchResult []string
}

func RenderSearchMode(writer http.ResponseWriter, searchstate SearchStruct, modeimplementation string) {
	if modeimplementation == "" {
		modeimplementation = `{{- template "searchmode" . -}}`
	}
	tpl := template.Must(template.New("ModeTemplate").Parse(SearchModeTemplate + modeimplementation))
	if err := tpl.Execute(writer, searchstate); err != nil {
		panic(err)
	}
}

type SearchResultEntry struct {
	Path string
	Uri  string
}

var SearchResultTemplate = `
{{- define "searchresults" -}}
	{{- range . -}}
		<tr><td><a href="{{- .Uri -}}" target="_blank" class="hi-navigation-editable-item">{{- .Path -}}</a></td></tr>
	{{- end -}}
{{- end -}}
{{- template "searchresults" . -}}
`

func RenderContentSearchResult(writer http.ResponseWriter, searchresults []SearchResultEntry) {
	tpl := template.Must(template.New("SearchResultTemplate").Parse(SearchResultTemplate))
	if err := tpl.Execute(writer, searchresults); err != nil {
		panic(err)
	}
}
