package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func ConfigMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Configuration)
	mode := fragment.ConfigStruct{
		Menu: menu,
	}
	fragment.RenderConfigMode(writer, mode)
}
