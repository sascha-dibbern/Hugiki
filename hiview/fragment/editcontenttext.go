package fragment

import (
	"fmt"

	converthtml "github.com/sascha-dibbern/Hugiki/hiconverters/html"
	"github.com/sascha-dibbern/Hugiki/himodel"
)

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
	gitdiff := himodel.GitDiff(filepath)
	return converthtml.NewlineToLinebreakString(gitdiff)
}
