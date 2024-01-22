package hihandlers

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiview/pagegenerator"
)

type DirectRequestManipulator struct {
}

func (manip DirectRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	backendBaseUrl := hiconfig.AppConfig().BackendBaseUrl()
	url := backendBaseUrl + request.URL.RequestURI()
	return url
}

func PipeThroughHandler(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := DirectRequestManipulator{}
	pagetemplate := pagegenerator.RootPageGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}
