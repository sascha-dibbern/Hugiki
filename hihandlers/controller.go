package hihandlers

import (
	"fmt"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiview"
)

const uriPage_EditContent = hiview.UriPage_EditContent
const uriAction_ProxyContentPageBody = hiview.UriAction_ProxyContentPageBody
const uriAction_UpdateContent = hiview.UriAction_UpdateContent

func Setup(mux *http.ServeMux) {
	mux.HandleFunc("/", pipeThroughHandler)
	//mux.HandleFunc("/hugikidev/restart/", devRestart)
	mux.HandleFunc(uriPage_EditContent, PageHandler_EditContent)
	mux.HandleFunc(uriAction_ProxyContentPageBody, ActionHandler_ProxyContentPageBody)
	mux.HandleFunc(uriAction_UpdateContent, ActionHandler_UpdateContentText)
	mux.HandleFunc("/hugiki/action/edit-and-update/", editAndUpdateHandler)
	mux.HandleFunc("/hugiki/action/save-and-close/", saveAndCloseHandler)
	mux.HandleFunc("/hugiki/action/create-child-page/", createChildPageHandler)
	mux.HandleFunc("/hugiki/action/create-child-page-commit/", createChildPageCommitHandler)
	mux.HandleFunc("/hugiki/action/create-sibling-page/", createSiblingPageHandler)
	mux.HandleFunc("/hugiki/action/create-sibling-page-commit/", createChildPageCommithugikiHandler)
}

type DirectRequestManipulator struct {
}

func (manip DirectRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	backendBaseUrl := hiconfig.AppConfig().BackendBaseUrl()
	url := backendBaseUrl + request.URL.RequestURI()
	return url
}

func pipeThroughHandler(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := DirectRequestManipulator{}
	pagetemplate := hiview.StartPageGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}

type EditContentPageRequestManipulator struct {
}

// Transform (Hugiki)"/hugiki/page/edit/content/xyz.." to (Hugo)"/xyz..."
func (manip EditContentPageRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	hugikiUri := request.URL.RequestURI()
	return hiproxy.HugikiUriToHugoUrlRule(uriPage_EditContent, "/").ConvertAll(hugikiUri)
}

func PageHandler_EditContent(writer http.ResponseWriter, request *http.Request) {
	var requestmanipulator hiproxy.RequestManipulator = EditContentPageRequestManipulator{}
	var pagetemplate hiproxy.ProxyPageGenerator = hiview.EditContentPageGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}

type ProxyContentPageBodyRequestManipulator struct {
}

// Transform (Hugiki)"/hugiki/.../xyz.." to (Hugo)"/xyz..."
func (manip ProxyContentPageBodyRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	hugikiUri := request.URL.RequestURI()
	return hiproxy.HugikiUriToHugoUrlRule(uriAction_ProxyContentPageBody, "/").ConvertAll(hugikiUri)
}

func ActionHandler_ProxyContentPageBody(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := ProxyContentPageBodyRequestManipulator{}
	pagetemplate := hiview.ContentPageBodyGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}

func ActionHandler_UpdateContentText(writer http.ResponseWriter, request *http.Request) {
	newcontent := request.PostFormValue("text")
	//hugopath := hiproxy.HugikiToHugoUriRule(uriAction_UpdateContent, hiconfig.AppConfig().HugoProject()).ConvertAll(request.RequestURI)
	match := hiview.Filepath_From_UriAction_UpdateContent_Regexp.FindStringSubmatch(request.RequestURI)
	localhugopath := "content/" + match[1] + ".md"
	hugopath := match[1]
	himodel.SaveContentMarkdown(hugopath, newcontent)
	fmt.Fprintln(writer, hiview.Render_EditContentText(newcontent, localhugopath))
}

func editAndUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the editAndUpdateHandler")
}

func saveAndCloseHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the saveAndCloseHandler")
}

func createChildPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageHandler")
}

func createChildPageCommitHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageCommitHandler")
}

func createSiblingPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createSiblingPageHandler")
}

func createChildPageCommithugikiHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageCommithugikiHandler")
}
