package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func StatusMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Configuration)
	mode := fragment.StatusStruct{
		Menu: menu,
	}
	fragment.RenderStatusMode(writer, mode, "")
}
