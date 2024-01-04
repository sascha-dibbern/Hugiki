package main

import (
    "fmt"
	"net/http"
    "bytes"
    "io"
	"regexp"
    "golang.org/x/text/encoding/charmap"
    "github.com/gookit/config/v2"
    "github.com/gookit/config/v2/yaml"
)

const hugikiStartTag = "<hugiki>"
const hugikiEndTag = "</hugiki>"
const loadHtmxHtml = `<script src="https://unpkg.com/htmx.org@1.9.10"></script>`

var backendBaseUrl string
var hugoProject string


func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func getConfig() {
	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("hugiki.yml")
	if err != nil {
		panic(err)
	}

	if config.String("backendBaseUrl") == "" {
		panic("Configuration error: 'backendBaseUrl' is empty")
	}
	backendBaseUrl = config.String("backendBaseUrl")
	fmt.Println(backendBaseUrl)
	
	if config.String("hugoProject") == "" {
		panic("Configuration error: 'hugoProject' is empty")
	}
	hugoProject = config.String("hugoProject")
	fmt.Println(hugoProject)
}



/*
 * Proxy logic
 */


func requestBackend(frontendReq *http.Request) *http.Response {
	//url := backendBaseUrl+frontendReq.URL.Path
	url := backendBaseUrl+frontendReq.URL.RequestURI()

	fmt.Println("GET: "+url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
        panic(err)
	}

	//req.Header.Add("Accept-Charset","utf-8")
	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
        panic(err)
    }
	return resp
}

func isTextResponse(backendResp *http.Response) bool {
	contenttype := backendResp.Header.Get("Content-Type")
	match, _ := regexp.MatchString("text", contenttype)
	return match
}

func isHtmlResponse(backendResp *http.Response) bool {
	contenttype := backendResp.Header.Get("Content-Type")
	match, _ := regexp.MatchString("text/html", contenttype)
	return match
}

func is_ISO8859_1_Response(backendResp *http.Response) bool {
	contenttype := backendResp.Header.Get("Content-Type")
	match, _ := regexp.MatchString("ISO-8859-1", contenttype)
	return match
}

func is_UTF_8_Response(backendResp *http.Response) bool {
	contenttype := backendResp.Header.Get("Content-Type")
	match, _ := regexp.MatchString("UTF-8", contenttype)
	return match
}

func ensureUtf8TextContentType(backendResp *http.Response, w http.ResponseWriter) {
	rexp, _       := regexp.Compile("text/w+")
	oldContentTyp := backendResp.Header.Get("Content-Type")
	TextType      := rexp.FindString(oldContentTyp)
	newContentTyp := TextType+"; charset=UTF-8"
    w.Header().Set("Content-Type",newContentTyp)
}

func ensureSameContentType(backendResp *http.Response, w http.ResponseWriter) {
	oldContentTyp := backendResp.Header.Get("Content-Type")
    w.Header().Set("Content-Type",oldContentTyp)
}

func insertWorkareaHtml(htmlInput string,frontendReq *http.Request) string {
	workarea := `
	<h3>edit<h3></br>	
	<h3>create sibling page<h3></br>
	<h3>create child page<h3>
	`

	rexp   := regexp.MustCompile(hugikiEndTag)
	result := rexp.ReplaceAllString(htmlInput,hugikiEndTag+workarea)		
	
	return result
}

// <body>...</body> -> <body><hugiki>...</hugiki><body>
func makeHugikiHtml (htmlInput string,frontendReq *http.Request) string {
	//bodyStartTag    := "<body>"
	bodyStartTagRXP := "<body.*>\n"
	bodyEndTag      := "</body>"

	result1 := htmlInput
	matchHugikiStartTag, _ := regexp.MatchString(hugikiStartTag, htmlInput)
	if ! matchHugikiStartTag {
		rexp1 := regexp.MustCompile(bodyStartTagRXP)
		matchedBodyStartTag := rexp1.FindString(htmlInput)
		fmt.Println("matchBodyStartTag:"+matchedBodyStartTag)
		
		rexp2 := regexp.MustCompile(bodyStartTagRXP)
		result1 = rexp2.ReplaceAllString(htmlInput,matchedBodyStartTag+loadHtmxHtml+hugikiStartTag)
		//fmt.Println(result1)
	}

	result2 := result1
	matchHugikiEndTag, _ := regexp.MatchString(hugikiEndTag, result1)
	if ! matchHugikiEndTag {
		rexp := regexp.MustCompile(bodyEndTag)
		result2 = rexp.ReplaceAllString(result1,hugikiEndTag+bodyEndTag)
		//fmt.Println(result2)
	}
	return insertWorkareaHtml(result2,frontendReq)
}


func pipeText(backendResp *http.Response, w http.ResponseWriter, frontendReq *http.Request) {
	// Default text reader
	var reader io.Reader = backendResp.Body

	// We always provide UTF-8 to front end
	ensureUtf8TextContentType(backendResp,w)

	// Handle right usage of charset-mapping to UTF-8
	if (! is_UTF_8_Response(backendResp)) {
		if is_ISO8859_1_Response(backendResp) {
			reader = charmap.ISO8859_1.NewDecoder().Reader(backendResp.Body)
		} else {
			// Warn for undefined charset
			fmt.Println("Handling undefined charset:", backendResp.Header.Get("Content-Type"))
		}
	}
	
	responseBytes, err := io.ReadAll(reader) // Read response body as bytes
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}	
	
	bodystring := string(responseBytes) // Convert bytes to string
	if isHtmlResponse(backendResp) {
		bodystring = makeHugikiHtml(bodystring,frontendReq)
	}

	//fmt.Println(bodystring)
	fmt.Fprintln(w,bodystring)
  
	backendResp.Body.Close()
}

func pipeNontext(backendResp *http.Response, w http.ResponseWriter, frontendReq *http.Request) {
	// Pipe same content-type through
	ensureSameContentType(backendResp,w)
	responseBytes, err := io.ReadAll(backendResp.Body) // Read response body as bytes
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	
	bodystring := string(responseBytes) // Convert bytes to string
	//fmt.Println(bodystring)
	fmt.Fprintln(w,bodystring)
  
	backendResp.Body.Close()
}

func pipeThroughHandler(w http.ResponseWriter, frontendReq *http.Request) {
	
	backendResp := requestBackend(frontendReq)

	if isTextResponse(backendResp) {
		pipeText(backendResp,w,frontendReq)
	} else {
		pipeNontext(backendResp,w,frontendReq)
	}

}

/*
 * 
 */



func startAndEditHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the startAndEditHandler")
}

func editAndUpdateHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the editAndUpdateHandler")
}

func saveAndCloseHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the saveAndCloseHandler")
}

func createChildPageHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the createChildPageHandler")
}

func createChildPageCommitHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the createChildPageCommitHandler")
}

func createSiblingPageHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the createSiblingPageHandler")
}

func createChildPageCommithugikiHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the createChildPageCommithugikiHandler")
}

func main() {
	getConfig()
	
	mux := http.NewServeMux()

	mux.HandleFunc("/", pipeThroughHandler)
	mux.HandleFunc("/hugiki/start-and-edit/", startAndEditHandler)
	mux.HandleFunc("/hugiki/edit-and-update/", editAndUpdateHandler)
	mux.HandleFunc("/hugiki/save-and-close/", saveAndCloseHandler)
	mux.HandleFunc("/hugiki/create-child-page/", createChildPageHandler)
	mux.HandleFunc("/hugiki/create-child-page-commit/", createChildPageCommitHandler)
	mux.HandleFunc("/hugiki/create-sibling-page/", createSiblingPageHandler)
	mux.HandleFunc("/hugiki/create-sibling-page-commit/", createChildPageCommithugikiHandler)
	
	http.ListenAndServe(":3000", mux)	
}