package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func HelpMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Help)
	mode := fragment.HelpStruct{
		Menu: menu,
	}
	fragment.RenderHelpMode(writer, mode)
}
