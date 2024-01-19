package action

import (
	"fmt"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiview"
)

func UpdateContentText(writer http.ResponseWriter, request *http.Request) {
	newcontent := request.PostFormValue("text")
	//hugopath := hiproxy.HugikiToHugoUriRule(uriAction_UpdateContent, hiconfig.AppConfig().HugoProject()).ConvertAll(request.RequestURI)
	match := hiview.Filepath_From_UriAction_UpdateContent_Regexp.FindStringSubmatch(request.RequestURI)
	localhugopath := "content/" + match[1] + ".md"
	hugopath := match[1]
	himodel.SaveContentMarkdown(hugopath, newcontent)
	fmt.Fprintln(writer, hiview.Render_EditContentText(newcontent, localhugopath))
}
