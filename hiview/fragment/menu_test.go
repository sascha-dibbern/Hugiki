package fragment

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"testing"

	"github.com/sascha-dibbern/Hugiki/hiuri"
)

func TestBuildMenuState(t *testing.T) {
	menustate := BuildMenuState(MenuItem_Navigation)
	if menustate[0]["uri"] != "" {
		t.Error("Expected empty uri for selected menu item")
	}
}

func TestMenuTemplate(t *testing.T) {
	tpl := template.Must(template.New("MenuTemplate").Parse(MenuTemplate + "{{ template \"menu\" . }}"))
	menustate := BuildMenuState("")
	var b bytes.Buffer
	if err := tpl.Execute(&b, menustate); err != nil {
		panic(err)
	}
	s := b.String()
	fmt.Print(s)
	match1, _ := regexp.MatchString(MenuItem_Git, s)
	if !match1 {
		panic("Menu item 'Git' is not filled into the template")
	}
	match2, _ := regexp.MatchString(hiuri.UriAction_GitMode, s)
	if !match2 {
		panic("Menu item uri 'git' is not filled into the template")
	}

}
