package action

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func GitMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Git)
	mode := fragment.GitStruct{
		Menu: menu,
	}
	fragment.RenderGitMode(writer, mode)
}
