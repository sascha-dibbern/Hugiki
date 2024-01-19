package action

import (
	"net/http"

	convertpath "github.com/sascha-dibbern/Hugiki/hiconverters/path"
	"github.com/sascha-dibbern/Hugiki/hihandlers/definitions"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiview"
)

type ProxyContentPageBodyRequestManipulator struct {
}

// Transform (Hugiki)"/hugiki/.../xyz.." to (Hugo)"/xyz..."
func (manip ProxyContentPageBodyRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	hugikiUri := request.URL.RequestURI()
	return convertpath.HugikiUriToHugoUrlRule(definitions.UriAction_ProxyContentPageBody, "/").ConvertAll(hugikiUri)
}

func ProxyContentPageBody(writer http.ResponseWriter, request *http.Request) {
	requestmanipulator := ProxyContentPageBodyRequestManipulator{}
	pagetemplate := hiview.ContentPageBodyGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}
