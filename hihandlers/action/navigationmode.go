package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func NavigationMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Navigation)
	mode := fragment.NavigationStruct{
		Menu: menu,
	}
	fragment.RenderNavigationMode(writer, mode)
}
