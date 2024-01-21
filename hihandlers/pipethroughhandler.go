package hihandlers

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
)

type DirectRequestManipulator struct {
}

func (manip DirectRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	backendBaseUrl := hiconfig.AppConfig().BackendBaseUrl()
	url := backendBaseUrl + request.URL.RequestURI()
	return url
}

/*
func pipeThroughHandler(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := DirectRequestManipulator{}
	pagetemplate := pagegenerator.StartPageGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}
*/
