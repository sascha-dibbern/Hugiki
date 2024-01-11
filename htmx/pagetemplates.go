package htmx

import (
//    "fmt"
	"net/http"
	"regexp"
	"github.com/sascha-dibbern/Hugiki/appconfig"	
)

const loadHtmxHtml = `<script src="https://unpkg.com/htmx.org@1.9.10"></script>`

/* 
 * <body>...</body> -> <body><hugiki>...</hugiki><body>
 *  <div> <!-- Content polling -->
 *   ...
 *  </div> 
 *  <div> <!-- Hugiki control -->
 *   ...
 *  </div> 
 * </body>
 */ 
 
func MakeHugikiWorkpageHtml (htmlInput string,request *http.Request) string {
	backendBaseUrl := appconfig.AppConfig().BackendBaseUrl()	
	pollingUrl := backendBaseUrl+request.URL.RequestURI()

	bodyStartTagRXP := "<body.*>\n"
	bodyEndTag      := "</body>"

	
	rexp1 := regexp.MustCompile(bodyStartTagRXP)
	matched := rexp1.FindString(htmlInput)
	result1 := rexp1.ReplaceAllString(htmlInput,matched+loadHtmxHtml+"<!--Content polling area --><div hx-get=\""+pollingUrl+"\" hx-trigger=\"every 1s\">")
	//fmt.Println(result1)

	rexp2 := regexp.MustCompile(bodyEndTag)
	result2 := rexp2.ReplaceAllString(result1,"</div><!-- Hugiki control area --><div>")
	//fmt.Println(result2)

	result3 := insertWorkareaHtml(result2,request)

	return result3+"</div>"+bodyEndTag
}

func insertWorkareaHtml(htmlInput string,frontendReq *http.Request) string {
	workarea := `
	<h3>edit<h3></br>	
	<h3>create sibling page<h3></br>
	<h3>create child page<h3>
	`
	
	return htmlInput+workarea
}
