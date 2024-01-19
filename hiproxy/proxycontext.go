package hiproxy

import (
	"net/http"
	"regexp"
)

/**********************************
 * Proxy context
 **********************************/

type ProxyContext struct {
	Request            *http.Request
	backendResponse    *http.Response
	outputwriter       http.ResponseWriter
	Requestmanipulator RequestManipulator
	proxypagegenerator ProxyPageGenerator
}

func (context ProxyContext) ensureSameContentType() {
	oldcontenttype := context.backendResponse.Header.Get("Content-Type")
	context.outputwriter.Header().Set("Content-Type", oldcontenttype)
}

func (context ProxyContext) ensureUtf8TextContentType() {
	rexp, _ := regexp.Compile("text/w+")
	oldcontenttype := context.backendResponse.Header.Get("Content-Type")
	TextType := rexp.FindString(oldcontenttype)
	newcontenttype := TextType + "; charset=UTF-8"
	context.outputwriter.Header().Set("Content-Type", newcontenttype)
}
