package pagegenerator

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

var startpagetemplate = `
{{- define "startpagetemplate" -}}
<html>
<head>
	<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
	<meta http-equiv="Pragma" content="no-cache" />
	<meta http-equiv="Expires" content="0" />
	<link rel="stylesheet" href="/hugiki/static/hugiki.css">
</head>
<body>
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
<h1 class="hi-header">Hugiki</h1>
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
