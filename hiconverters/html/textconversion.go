package html

import "regexp"

var newlineRegexp = regexp.MustCompile(`\r?\n+`)

func NewlineToLinebreakString(s string) string {
	return newlineRegexp.ReplaceAllString(s, "<br/>\n")
}
