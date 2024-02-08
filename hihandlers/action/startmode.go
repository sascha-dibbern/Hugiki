package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func StartMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Configuration)
	mode := fragment.StartStruct{
		Menu: menu,
	}
	fragment.RenderStartMode(writer, mode, "")
}
