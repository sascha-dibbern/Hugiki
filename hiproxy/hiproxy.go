package hiproxy

import (
    "fmt"
	"io"
	"net/http"
	"regexp"
    "golang.org/x/text/encoding/charmap"	
//    "github.com/gookit/config/v2"
//	"github.com/sascha-dibbern/Hugiki/appconfig"
	"github.com/sascha-dibbern/Hugiki/htmx"
)

/**********************************
 * Content-Type helper functions
 **********************************/

func getContentType(response *http.Response) string {
	return response.Header.Get("Content-Type")
}

func isTextResponse(contenttype string) bool {
	match, _ := regexp.MatchString("text", contenttype)
	return match
}

func isHtmlResponse(contenttype string) bool {
	match, _ := regexp.MatchString("text/html", contenttype)
	return match
}

func is_ISO8859_1_Response(contenttype string) bool {
	match1, _ := regexp.MatchString("ISO-8859-1", contenttype)
	if match1 {
		return match1
	}
	match2, _ := regexp.MatchString("iso-8859-1", contenttype)
	return match2
}

func is_UTF_8_Response(contenttype string) bool {
	match1, _ := regexp.MatchString("UTF-8", contenttype)
	if match1 {
		return match1
	}
	match2, _ := regexp.MatchString("utf-8", contenttype)
	return match2
}

/**********************************
 * Proxy context
 **********************************/

type ProxyContext struct {   
	request *http.Request
	backendResponse *http.Response
	outputwriter http.ResponseWriter
}

func (context ProxyContext) ensureSameContentType() {
	oldcontenttype := context.backendResponse.Header.Get("Content-Type")
    context.outputwriter.Header().Set("Content-Type",oldcontenttype)
}

func (context ProxyContext) ensureUtf8TextContentType() {
	rexp, _       := regexp.Compile("text/w+")
	oldcontenttype:= context.backendResponse.Header.Get("Content-Type")
	TextType      := rexp.FindString(oldcontenttype)
	newcontenttype:= TextType+"; charset=UTF-8"
    context.outputwriter.Header().Set("Content-Type",newcontenttype)
}

/**********************************
 * Request manipulation
 **********************************/

type RequestManipulator interface {
	generateBackendUrl(request *http.Request) string
}

/**********************************
 * Response manipulation
 **********************************/

type ResponseManipulator interface {
	pipe() string
}

type TextResponseManipulator struct {   
	context *ProxyContext
}

func NewTextResponseManipulator(context *ProxyContext) TextResponseManipulator {
	return TextResponseManipulator {
		context,
	}
}

func (manip TextResponseManipulator) pipe() string {
	context        := manip.context
	backendResponse:= context.backendResponse
	
	// Default text reader
	var reader io.Reader = backendResponse.Body

	// We always provide UTF-8 to front end
	context.ensureUtf8TextContentType()

	// Handle right usage of charset-mapping to UTF-8
	contenttype := getContentType(backendResponse)
	if (! is_UTF_8_Response(contenttype)) {
		if is_ISO8859_1_Response(contenttype) {
			reader = charmap.ISO8859_1.NewDecoder().Reader(backendResponse.Body)
		} else {
			// Warn for undefined charset
			fmt.Println("Handling undefined charset:", backendResponse.Header.Get("Content-Type"))
		}
	}
	
	responsebytes, err := io.ReadAll(reader) // Read response body as bytes
	if err != nil {
		fmt.Println("Error reading response body:", err)
		panic(err)
	}	
	
	textstring := string(responsebytes) // Convert bytes to string	
	if isHtmlResponse(contenttype) {
		textstring = htmx.MakeHugikiWorkpageHtml(textstring,context.request)
	}

	//fmt.Println(textstring)
	return textstring
}

type NonTextResponseManipulator struct {   
	context *ProxyContext
}

func NewNonTextResponseManipulator(context *ProxyContext) NonTextResponseManipulator {
	return NonTextResponseManipulator {
		context,
	}
}

func (manip NonTextResponseManipulator) pipe() string {
	context := manip.context

	// Pipe same content-type through
	context.ensureSameContentType()

	responseBytes, err := io.ReadAll(context.backendResponse.Body) // Read response body as bytes
	if err != nil {
		fmt.Println("Error reading response body:", err)
		panic(err)
	}
	
	data := string(responseBytes) // Convert bytes to string
	//fmt.Println(data)
	fmt.Fprintln(context.outputwriter,data)
  
	return data
}


/**********************************
 * Proxy
 **********************************/
 
type RequestObjectProxy struct {   
	context *ProxyContext
	requestmanipulator RequestManipulator
	responsemanipulator ResponseManipulator
} 

func NewRequestObjectProxy(writer http.ResponseWriter, request *http.Request, requestmanipulator RequestManipulator) RequestObjectProxy {
	context := ProxyContext {
			request,
			nil,
			writer,
	}
	return RequestObjectProxy {
		&context,
		requestmanipulator,
		nil,	
	}
}

func (proxy *RequestObjectProxy) GenericProxyRequest() {
	proxy.requestBackend()
	context := proxy.context
	defer context.backendResponse.Body.Close()
	
	if isTextResponse(getContentType(context.backendResponse)) {
		proxy.responsemanipulator = NewTextResponseManipulator(context)
	} else {
		proxy.responsemanipulator = NewNonTextResponseManipulator(context)
	}
	
	manipulatedstring := proxy.responsemanipulator.pipe()
	
	//fmt.Println(manipulated)
	fmt.Fprintln(context.outputwriter,manipulatedstring)
}

func (proxy *RequestObjectProxy) requestBackend() {
	context := proxy.context
	url     := proxy.requestmanipulator.generateBackendUrl(context.request)

	fmt.Println("GET: "+url)
	backendrequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
        panic(err)
	}

	backendrequest.Header.Add("Accept-Charset","utf-8")
	client := &http.Client{}
	
	backendresponse, err := client.Do(backendrequest)
	if err != nil {
        panic(err)
    }
	context.backendResponse = backendresponse
}

