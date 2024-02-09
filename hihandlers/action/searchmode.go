package action

import (
	"fmt"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiuri"
	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func SearchMode(writer http.ResponseWriter, request *http.Request) {
	menu := fragment.BuildMenuState(fragment.MenuItem_Configuration)
	mode := fragment.SearchStruct{
		Menu: menu,
	}
	fragment.RenderSearchMode(writer, mode, "")
}

func contentsearchresult(files []string) []fragment.SearchResultEntry {
	entries := make([]fragment.SearchResultEntry, len(files))

	for index, file := range files {
		uri := fmt.Sprintf("%s%s", hiuri.UriPage_EditContent, file)

		entries[index] = fragment.SearchResultEntry{
			Path: file,
			Uri:  uri,
		}
	}
	return entries
}

func ContentSearchResultAction(writer http.ResponseWriter, request *http.Request) {
	query := request.FormValue("search")
	queryresult := himodel.SearchContentFiles(query)
	result := contentsearchresult(queryresult)
	fragment.RenderContentSearchResult(writer, result)
}
