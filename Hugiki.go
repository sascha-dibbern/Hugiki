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

const hugikiTag = "<hugiki/>"

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
 * 
 */



func hugikiHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Welcome to the hugikiHandler")

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

func replaceHugikiTag(htmlInput string) string {
	return htmlInput
}

func makeHugikiHtml (htmlInput string) string {
	matchHugikiTag, _ := regexp.MatchString(hugikiTag, htmlInput)
	var result string
	if ! matchHugikiTag {
		body := "</body>"
		rexp := regexp.MustCompile(body)
		result = rexp.ReplaceAllString(htmlInput,hugikiTag+body)
		fmt.Println(result)
	}
	return replaceHugikiTag(result)
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
		bodystring = makeHugikiHtml(bodystring)
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


func main() {
	getConfig()
	
	mux := http.NewServeMux()

	mux.HandleFunc("/hugiki/", hugikiHandler)
	mux.HandleFunc("/", pipeThroughHandler)
	
	http.ListenAndServe(":3000", mux)	
}