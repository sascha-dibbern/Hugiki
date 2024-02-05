package hihandlers

import (
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/hihandlers/action"
	"github.com/sascha-dibbern/Hugiki/hihandlers/page"
	"github.com/sascha-dibbern/Hugiki/hiuri"
)

func Setup(mux *http.ServeMux) {
	// Page-handlers
	mux.HandleFunc(hiuri.UriPage_Root, PipeThroughHandler)
	mux.HandleFunc(hiuri.UriPage_HugikiRoot, page.StartPage)
	mux.HandleFunc(hiuri.UriPage_EditContent, page.EditContent)
	static := hiconfig.AppConfig().HugikiStatic()
	if static != "" {
		mux.Handle(hiuri.UriPage_Static, http.FileServer(http.Dir(static)))
	}

	// State-page action handlers
	mux.HandleFunc(hiuri.UriAction_ConfigMode, action.ConfigMode)
	mux.HandleFunc(hiuri.UriAction_NavigationMode, action.NavigationMode)
	mux.HandleFunc(hiuri.UriAction_GitMode, action.GitMode)
	mux.HandleFunc(hiuri.UriAction_HelpMode, action.HelpMode)

	// Edit-page action handlers
	mux.HandleFunc(hiuri.UriAction_ProxyContentPageBody, action.ProxyContentPageBody)
	mux.HandleFunc(hiuri.UriAction_UpdateContent, action.UpdateContentText)
}
