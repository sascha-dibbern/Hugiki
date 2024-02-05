package action

import (
	"net/http"

	convertpath "github.com/sascha-dibbern/Hugiki/hiconverters/path"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiuri"
	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

type ProxyContentPageBodyRequestManipulator struct {
}

// Transform (Hugiki)"/hugiki/.../xyz.." to (Hugo)"/xyz..."
func (manip ProxyContentPageBodyRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	hugikiUri := request.URL.RequestURI()
	backendurl := convertpath.HugikiUriToHugoContentUrlRule(hiuri.UriAction_ProxyContentPageBody, "/").ConvertAll(hugikiUri)
	return backendurl
}

func ProxyContentPageBody(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := ProxyContentPageBodyRequestManipulator{}
	pagetemplate := fragment.ContentPageBodyGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}
