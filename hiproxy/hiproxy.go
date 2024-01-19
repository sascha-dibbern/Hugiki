package hiproxy

import (
	"fmt"
	"net/http"
)

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
