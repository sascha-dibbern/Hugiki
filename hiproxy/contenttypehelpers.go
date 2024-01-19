/**********************************
 * Content-Type helper functions
 **********************************/
package hiproxy

import (
	"net/http"
	"regexp"
)

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
