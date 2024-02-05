package pagegenerator

import (
	"fmt"
	"regexp"

	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiuri"
	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

// Todo: link to it like a popup: https://www.rapidtables.com/web/html/link/html-link-new-window.html

var UriPage_EditContentRegexp = regexp.MustCompile(hiuri.UriPage_EditContent)

var Filepath_From_UriPage_Edit_Regexp = regexp.MustCompile(hiuri.UriPage_Edit + "(.+)")
var Filepath_From_UriAction_UpdateContent_Regexp = regexp.MustCompile(hiuri.UriAction_UpdateContent + "(.+)")

var UriAction_ProxyContentPageBodyRegexp = regexp.MustCompile(hiuri.UriAction_ProxyContentPageBody)

const loadHtmxHtml = `<script src="https://unpkg.com/htmx.org@1.9.10"></script>`

const bodyStartTagRXPS = "<body.*>\n"

var bodyStartTagRegexp = regexp.MustCompile(bodyStartTagRXPS)

const bodyEndTagRXPS = "</body>"

var bodyEndTagRegexp = regexp.MustCompile(bodyEndTagRXPS)

/*
 * <body>...</body> -> <body><hugiki>...</hugiki><body>
 *  <div> <!-- Content polling -->
 *   ...
 *  </div>
 *  <div> <!-- Hugiki edit control -->
 *   ...
 *  </div>
 * </body>
 */

type EditContentPageGenerator struct {
}

func (generator EditContentPageGenerator) pollingUrl(context *hiproxy.ProxyContext) string {
	request := context.Request
	proxyContentPageBodyUri := UriPage_EditContentRegexp.ReplaceAllString(request.RequestURI, hiuri.UriAction_ProxyContentPageBody)
	return proxyContentPageBodyUri
}

func (generator EditContentPageGenerator) contentPathFragment(context *hiproxy.ProxyContext) string {
	request := context.Request
	proxyContentPageBodyUri := UriPage_EditContentRegexp.ReplaceAllString(request.RequestURI, "")
	return proxyContentPageBodyUri
}

func (generator EditContentPageGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	pollingUrl := generator.pollingUrl(context)
	matchedbodystart := bodyStartTagRegexp.FindString(htmlInput)
	modifiedpollingareastring := bodyStartTagRegexp.ReplaceAllString(
		htmlInput,
		fmt.Sprintf(`
			%s %s
			<!--Content polling area -->
			<div hx-get="%s" hx-trigger="every 1s">
		`, matchedbodystart, loadHtmxHtml, pollingUrl),
	)

	modifiedcontrolareastring := bodyEndTagRegexp.ReplaceAllString(
		modifiedpollingareastring,
		"</div><!-- Hugiki edit control area --><div>",
	)

	update_uri := hiuri.UriAction_UpdateContent + generator.contentPathFragment(context)
	match := Filepath_From_UriPage_Edit_Regexp.FindStringSubmatch(context.Request.RequestURI)
	filepath := match[1]
	markdown_content := himodel.LoadTextFromFile(filepath)
	result := modifiedcontrolareastring + fmt.Sprintf(`
	</br>	
	<div>
    	<form hx-post="%s">
		%s
		</form>
	</div>
	`, update_uri, fragment.Render_EditContentText(markdown_content, filepath))

	return result + "</div>" + bodyEndTagRXPS
}
