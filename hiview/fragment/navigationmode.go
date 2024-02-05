package fragment

import (
	"html/template"
	"net/http"
)

var NavigationModeTemplate = MenuTemplate + `
{{- define "path" -}}
{{- range .Path -}}
<span hx-get="{{- .Uri -}}" hx-swap="outerHTML" hx-target="#hi-navigation" class="hi-navigation-pathelement">{{- .Name -}}</span>/
{{- end -}}
{{- end -}}

{{- define "dir" -}}
<table>
	<tbody>
    	{{- range .Dir -}}
		<tr>
			<td>
			{{- if eq .Uri "" -}}
				<div class="hi-navigation-inaccessible-item">{{- .Name -}}</div>
			{{- else -}}
				{{ if .ContentEdit }}
					<a href="{{- .Uri -}}" target="_blank" class="hi-navigation-editable-item">{{- .Name -}}</a>
				{{- else -}}
					<div hx-get="{{- .Uri -}}" hx-swap="outerHTML" hx-target="#hi-navigation" class="hi-navigation-accessible-item">{{- .Name -}}</div>
				{{- end -}}				
			{{- end -}}
			</td>
		</tr>	
    	{{- end -}}
	</tbody>
</table>
{{- end -}}

{{- define "navigationmode" -}}
<div id="hi-navigation">
	<!-- navigationmode -->
	{{ template "menu" .Menu }}
	<h1>Files</h1>
	{{ template "path" . }}
	</br>
	{{ template "dir" . }}
</div>
{{- end -}}

{{ template "navigationmode" . }}
`

type DirEntry struct {
	Name        string
	Uri         string
	ContentEdit bool
}

type NavigationStruct struct {
	Menu      MenuItems
	Path      *[]DirEntry
	Dir       *[]DirEntry
	Errortext string
}

func RenderNavigationMode(writer http.ResponseWriter, mode NavigationStruct) {
	tpl := template.Must(template.New("ModeTemplate").Parse(NavigationModeTemplate))
	if err := tpl.Execute(writer, mode); err != nil {
		// Todo errorpage
		panic(err)
	}
}
