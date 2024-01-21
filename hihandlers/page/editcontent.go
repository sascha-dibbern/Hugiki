package page

import (
	"net/http"

	convertpath "github.com/sascha-dibbern/Hugiki/hiconverters/path"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
	"github.com/sascha-dibbern/Hugiki/hiuri"
	"github.com/sascha-dibbern/Hugiki/hiview/pagegenerator"
)

type EditContentPageRequestManipulator struct {
}

// Transform (Hugiki)"/hugiki/page/edit/content/xyz.." to (Hugo)"/xyz..."
func (manip EditContentPageRequestManipulator) GenerateBackendUrl(request *http.Request) string {
	hugikiUri := request.URL.RequestURI()
	return convertpath.HugikiUriToHugoUrlRule(hiuri.UriPage_EditContent, "/").ConvertAll(hugikiUri)
}

func EditContent(writer http.ResponseWriter, request *http.Request) {
	var requestmanipulator hiproxy.RequestManipulator = EditContentPageRequestManipulator{}
	var pagetemplate hiproxy.ProxyPageGenerator = pagegenerator.EditContentPageGenerator{}
	proxy := hiproxy.NewRequestObjectProxy(writer, request, requestmanipulator, pagetemplate)
	proxy.GenericProxyRequest()
}
