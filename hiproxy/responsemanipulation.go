package hiproxy

import (
	"fmt"
	"io"

	"golang.org/x/text/encoding/charmap"
)

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
