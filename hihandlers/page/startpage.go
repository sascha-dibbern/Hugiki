package page

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
	"github.com/sascha-dibbern/Hugiki/hiview/pagegenerator"
)

var view = pagegenerator.NewStartPageView()

func StartPage(writer http.ResponseWriter, request *http.Request) {
	menustate := fragment.BuildMenuState(fragment.MenuItem_Start)
	startstate := fragment.StartStruct{
		Menu: menustate,
	}
	view.Render(writer, startstate)
}
