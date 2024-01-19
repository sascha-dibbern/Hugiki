package hiview

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
)

const UriPage_Edit = "/hugiki/page/edit/"
const UriPage_EditContent = UriPage_Edit + "content/"
const UriAction_ProxyContentPageBody = "/hugiki/action/proxy-content-page-body/"
const UriAction_UpdateContent = "/hugiki/action/update/content/"

var Filepath_From_UriPage_Edit_Regexp = regexp.MustCompile(UriPage_Edit + "(.+)/")
var Filepath_From_UriAction_UpdateContent_Regexp = regexp.MustCompile(UriAction_UpdateContent + "(.+)/")
var UriPage_EditContentRegexp = regexp.MustCompile(UriPage_EditContent)
var UriAction_ProxyContentPageBodyRegexp = regexp.MustCompile(UriAction_ProxyContentPageBody)

const loadhiviewHtml = `<script src="https://unpkg.com/htmx.org@1.9.10"></script>`

const bodyStartTagRXPS = "<body.*>\n"

var bodyStartTagRegexp = regexp.MustCompile(bodyStartTagRXPS)

const bodyEndTagRXPS = "</body>"

var bodyEndTagRegexp = regexp.MustCompile(bodyEndTagRXPS)

/*
 * Should not be called
 */
type StartPageGenerator struct {
}

func (generator StartPageGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	return "<html><head></head><body>Startpage for configuration</body></html>"
}

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
	proxyContentPageBodyUri := UriPage_EditContentRegexp.ReplaceAllString(request.RequestURI, UriAction_ProxyContentPageBody)
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
		`, matchedbodystart, loadhiviewHtml, pollingUrl),
	)

	modifiedcontrolareastring := bodyEndTagRegexp.ReplaceAllString(
		modifiedpollingareastring,
		"</div><!-- Hugiki edit control area --><div>",
	)

	update_uri := UriAction_UpdateContent + generator.contentPathFragment(context)
	match := Filepath_From_UriPage_Edit_Regexp.FindStringSubmatch(context.Request.RequestURI)
	filepath := match[1] + ".md"
	markdown_content := himodel.LoadTextFromFile(filepath)
	result := modifiedcontrolareastring + fmt.Sprintf(`
	</br>	
	<div>
    	<form hx-post="%s">
		%s
		</form>
	</div>
	`, update_uri, Render_EditContentText(markdown_content, filepath))

	return result + "</div>" + bodyEndTagRXPS
}

/*
 * <body>...</body> -> <body><hugiki>...</hugiki><body>
 *  <div> <!-- Content polling -->
 *   ...
 *  </div>
 *  <div> <!-- Hugiki control -->
 *   ...
 *  </div>
 * </body>
 */

func MakeHugikiWorkpageHtml(htmlInput string, request *http.Request) string {
	backendBaseUrl := hiconfig.AppConfig().BackendBaseUrl()
	pollingUrl := backendBaseUrl + request.URL.RequestURI()

	bodyStartTagRXP := "<body.*>\n"
	bodyEndTag := "</body>"

	rexp1 := regexp.MustCompile(bodyStartTagRXP)
	matched := rexp1.FindString(htmlInput)
	result1 := rexp1.ReplaceAllString(htmlInput, matched+loadhiviewHtml+"<!--Content polling area --><div hx-get=\""+pollingUrl+"\" hx-trigger=\"every 1s\">")

	rexp2 := regexp.MustCompile(bodyEndTag)
	result2 := rexp2.ReplaceAllString(result1, "</div><!-- Hugiki control area --><div>")

	result3 := insertWorkareaHtml(result2, request)

	return result3 + "</div>" + bodyEndTag
}

func insertWorkareaHtml(htmlInput string, frontendReq *http.Request) string {
	workarea := `
	<h3>edit<h3></br>	
	<h3>create sibling page<h3></br>
	<h3>create child page<h3>
	`

	return htmlInput + workarea
}
