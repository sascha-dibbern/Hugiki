package hihandlers

import (
	"fmt"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hihandlers/action"
	"github.com/sascha-dibbern/Hugiki/hihandlers/definitions"
	"github.com/sascha-dibbern/Hugiki/hihandlers/page"
)

func Setup(mux *http.ServeMux) {
	mux.HandleFunc("/", pipeThroughHandler)
	//mux.HandleFunc("/hugikidev/restart/", devRestart)
	mux.HandleFunc(definitions.UriPage_EditContent, page.EditContent)
	mux.HandleFunc(definitions.UriAction_ProxyContentPageBody, action.ProxyContentPageBody)
	mux.HandleFunc(definitions.UriAction_UpdateContent, action.UpdateContentText)
	mux.HandleFunc("/hugiki/action/edit-and-update/", editAndUpdateHandler)
	mux.HandleFunc("/hugiki/action/save-and-close/", saveAndCloseHandler)
	mux.HandleFunc("/hugiki/action/create-child-page/", createChildPageHandler)
	mux.HandleFunc("/hugiki/action/create-child-page-commit/", createChildPageCommitHandler)
	mux.HandleFunc("/hugiki/action/create-sibling-page/", createSiblingPageHandler)
	mux.HandleFunc("/hugiki/action/create-sibling-page-commit/", createChildPageCommithugikiHandler)
}

func editAndUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the editAndUpdateHandler")
}

func saveAndCloseHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the saveAndCloseHandler")
}

func createChildPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageHandler")
}

func createChildPageCommitHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageCommitHandler")
}

func createSiblingPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createSiblingPageHandler")
}

func createChildPageCommithugikiHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the createChildPageCommithugikiHandler")
}
