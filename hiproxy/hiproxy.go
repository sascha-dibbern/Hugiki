package hiproxy

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"golang.org/x/text/encoding/charmap"
	//	 "github.com/gookit/config/v2"
	"github.com/sascha-dibbern/Hugiki/hiconfig"
)

type TextConversionRule interface {
	ConvertAll(input string) string
}

/**********************************
 * URI/URL helpers functions
 **********************************/

type TextConversionRuleDefinition struct {
	matchingregexp *regexp.Regexp
	replacement    string
}

func NewTextConversionRuleDefinition(matching string, replacement string) TextConversionRuleDefinition {
	return TextConversionRuleDefinition{
		regexp.MustCompile(matching),
		replacement,
	}
}

/**/

type hugoToHugikiUriRule struct {
	definition TextConversionRuleDefinition
}

func HugoToHugikiUriRule(matching string, replacement string) hugoToHugikiUriRule {
	// Todo: add check for URI
	return hugoToHugikiUriRule{
		NewTextConversionRuleDefinition(matching, replacement),
	}
}

func (rule hugoToHugikiUriRule) ConvertAll(hugoinput string) string {
	return rule.definition.matchingregexp.ReplaceAllString(hugoinput, rule.definition.replacement)
}

/**/

type hugoToHugikiUrlRule struct {
	definition TextConversionRuleDefinition
}

func HugoToHugikiUrlRule(matching_uri string, replacement_uri string) hugoToHugikiUrlRule {
	matching_url := hiconfig.AppConfig().BackendBaseUrl() + matching_uri
	replacement_url := replacement_uri
	return hugoToHugikiUrlRule{
		NewTextConversionRuleDefinition(matching_url, replacement_url),
	}
}

func (rule hugoToHugikiUrlRule) ConvertAll(hugoinput string) string {
	return rule.definition.matchingregexp.ReplaceAllString(hugoinput, rule.definition.replacement)
}

/**/

type hugikiToHugoUriRule struct {
	definition TextConversionRuleDefinition
}

func HugikiToHugoUriRule(matching string, replacement string) hugikiToHugoUriRule {
	// Todo: add check for URI
	return hugikiToHugoUriRule{
		NewTextConversionRuleDefinition(matching, replacement),
	}
}

func (rule hugikiToHugoUriRule) ConvertAll(hugoinput string) string {
	return rule.definition.matchingregexp.ReplaceAllString(hugoinput, rule.definition.replacement)
}

/**/

type hugikiUriToHugoUrlRule struct {
	definition TextConversionRuleDefinition
}

func HugikiUriToHugoUrlRule(hugiki_uri string, hugo_replacement_uri string) hugikiUriToHugoUrlRule {
	replacement_url := hiconfig.AppConfig().BackendBaseUrl() + "/" + hugo_replacement_uri
	return hugikiUriToHugoUrlRule{
		NewTextConversionRuleDefinition(hugiki_uri, replacement_url),
	}
}

func (rule hugikiUriToHugoUrlRule) ConvertAll(hugoinput string) string {
	return rule.definition.matchingregexp.ReplaceAllString(hugoinput, rule.definition.replacement)
}

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
 * Request manipulation
 **********************************/

type RequestManipulator interface {
	GenerateBackendUrl(request *http.Request) string
}

/**********************************
 * Proxy page generator
 **********************************/

type ProxyPageGenerator interface {
	GenerateHtml(proxiedcontent string, context *ProxyContext) string
}

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
	return TextResponseManipulator{
		context,
	}
}

func (manip TextResponseManipulator) pipe() string {
	context := manip.context
	backendResponse := context.backendResponse

	// Default text reader
	var reader io.Reader = backendResponse.Body

	// We always provide UTF-8 to front end
	context.ensureUtf8TextContentType()

	// Handle right usage of charset-mapping to UTF-8
	contenttype := getContentType(backendResponse)
	if !is_UTF_8_Response(contenttype) {
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
		//textstring = hiview.MakeHugikiWorkpageHtml(textstring,context.Request)
		proxypagegenerator := context.proxypagegenerator
		textstring = proxypagegenerator.GenerateHtml(textstring, context)
	}

	//fmt.Println(textstring)
	return textstring
}

type NonTextResponseManipulator struct {
	context *ProxyContext
}

func NewNonTextResponseManipulator(context *ProxyContext) NonTextResponseManipulator {
	return NonTextResponseManipulator{
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
	fmt.Fprintln(context.outputwriter, data)

	return data
}

/**********************************
 * Proxy
 **********************************/

type RequestObjectProxy struct {
	context             *ProxyContext
	responsemanipulator ResponseManipulator
}

func NewRequestObjectProxy(outputwriter http.ResponseWriter, request *http.Request, requestmanipulator RequestManipulator, proxypagegenerator ProxyPageGenerator) RequestObjectProxy {
	context := ProxyContext{
		request,
		nil,
		outputwriter,
		requestmanipulator,
		proxypagegenerator,
	}
	return RequestObjectProxy{
		&context,
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
	fmt.Fprintln(context.outputwriter, manipulatedstring)
}

func (proxy *RequestObjectProxy) requestBackend() {
	context := proxy.context
	url := proxy.context.Requestmanipulator.GenerateBackendUrl(context.Request)

	fmt.Println("GET: " + url)
	backendrequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	backendrequest.Header.Add("Accept-Charset", "utf-8")
	client := &http.Client{}

	backendresponse, err := client.Do(backendrequest)
	if err != nil {
		panic(err)
	}
	context.backendResponse = backendresponse
}
