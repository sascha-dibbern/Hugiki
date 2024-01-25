package pagegenerator

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

var startpagetemplate = `
{{- define "startpagetemplate" -}}
<html>
<head>
    
</head>
<body>
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
<div id="himode">
    <!-- startpage:startmode -->
	{{ template "startmode" . }}
</div>
</body>
</html>
{{- end -}}
{{- template "startpagetemplate" . -}}
`

type StartPageView struct {
}

func NewStartPageView() StartPageView {
	return StartPageView{}
}

func (generator StartPageView) Render(writer http.ResponseWriter, startstate fragment.StartStruct) {
	fragment.RenderStartMode(writer, startstate, startpagetemplate)
}
