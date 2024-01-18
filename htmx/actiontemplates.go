package htmx

import (
	"fmt"
	"regexp"

	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
)

const upWithBodyStartTagRXPS = ".*<body.*>"

var upWithBodyStartTagRegexp = regexp.MustCompile(upWithBodyStartTagRXPS)

const fromWithBodyEndTagRXPS = "</body>.*"

var fromWithBodyEndTagRegexp = regexp.MustCompile(fromWithBodyEndTagRXPS)

// ...<body...>xyz</body>... -> xyz
type ContentPageBodyGenerator struct {
}

func (generator ContentPageBodyGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	clean_in_start := upWithBodyStartTagRegexp.ReplaceAllString(htmlInput, "")
	clean_also_after_end := fromWithBodyEndTagRegexp.ReplaceAllString(clean_in_start, "")
	return clean_also_after_end
}

func Render_EditContentText(contenttext string, localhugopath string) string {
	return fmt.Sprintf(`
	<button type="submit">update</button> </br>
	<textarea id="edittext" name="text" rows="30" cols="80">%s</textarea><br/>
	<br/>Saved at %s<br/>
	Git differences:<br/>
	<blockquote>%s</blockquote>
	`, contenttext, localhugopath, render_FileGitDiff(localhugopath))
}

func render_FileGitDiff(filepath string) string {
	gitdiff := newlineToLinebreakString(himodel.GitDiff(filepath))

	return gitdiff
}

func newlineToLinebreakString(s string) string {
	re := regexp.MustCompile(`\r?\n+`)
	return re.ReplaceAllString(s, "<br/>\n")
}
