package hihandlers

import (
    "fmt"
	"net/http"
	"github.com/sascha-dibbern/Hugiki/appconfig"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
)

type HugikiRequestManipulator struct {   
}

func (manip *HugikiRequestManipulator) generateBackendUrl(request *http.Request) {
	backendBaseUrl := appconfig.AppConfig().BackendBaseUrl()
	url := backendBaseUrl+request.URL.RequestURI()
	// Todo "/content/"
	return url
}

func Setup(mux *http.ServeMux) {
	mux.HandleFunc("/", pipeThroughHandler)
	mux.HandleFunc("/hugikidev/restart/", devRestart)
	mux.HandleFunc("/hugiki/start-and-edit/", startAndEditHandler)
	mux.HandleFunc("/hugiki/edit-and-update/", editAndUpdateHandler)
	mux.HandleFunc("/hugiki/save-and-close/", saveAndCloseHandler)
	mux.HandleFunc("/hugiki/create-child-page/", createChildPageHandler)
	mux.HandleFunc("/hugiki/create-child-page-commit/", createChildPageCommitHandler)
	mux.HandleFunc("/hugiki/create-sibling-page/", createSiblingPageHandler)
	mux.HandleFunc("/hugiki/create-sibling-page-commit/", createChildPageCommithugikiHandler)
}


func pipeThroughHandler(writer http.ResponseWriter,request *http.Request) {	
	requestmanipulator = HugikiRequestManipulator()
	proxy := hiproxy.NewRequestObjectProxy(writer,request, requestmanipulator)
	proxy.GenericProxyRequest()
}

func startAndEditHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the startAndEditHandler")
}

func editAndUpdateHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the editAndUpdateHandler")
}

func saveAndCloseHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the saveAndCloseHandler")
}

func createChildPageHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the createChildPageHandler")
}

func createChildPageCommitHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the createChildPageCommitHandler")
}

func createSiblingPageHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the createSiblingPageHandler")
}

func createChildPageCommithugikiHandler(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Welcome to the createChildPageCommithugikiHandler")
}
