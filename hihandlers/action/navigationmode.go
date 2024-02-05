package action

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/himodel"
	"github.com/sascha-dibbern/Hugiki/hiuri"
	"github.com/sascha-dibbern/Hugiki/hiview/fragment"
)

func pathwidgetdata(dirlocalpath himodel.ExtendableSlashPath) []fragment.DirEntry {
	dirpathelements := dirlocalpath.Elements()
	entries := make([]fragment.DirEntry, len(dirpathelements)+1)

	// Hugo root-dir
	entries[0] = fragment.DirEntry{
		Name: "Project",
		Uri:  hiuri.UriAction_NavigationMode,
	}
	requestlink := ""
	hugolocalpath := ""
	for index, dirpathelement := range dirpathelements {
		hugolocalpath += "/" + dirpathelement
		requestlink = fmt.Sprintf("%s?path=%s", hiuri.UriAction_NavigationMode, hugolocalpath)
		entries[index+1] = fragment.DirEntry{
			Name: dirpathelement,
			Uri:  requestlink,
		}
	}
	return entries
}

func dirwidgetdata(dirlocalpath himodel.ExtendableSlashPath, direlements []himodel.SlashDirElement) []fragment.DirEntry {
	basepath := dirlocalpath.String()
	entries := make([]fragment.DirEntry, len(direlements))
	for index, direlement := range direlements {
		// Generate usable URI
		contentedit := false
		uri := ""
		hugolocalpath := basepath + "/" + direlement.Name
		rexp, _ := regexp.Compile("^/*content")
		if rexp.MatchString(hugolocalpath) {
			if direlement.IsDir {
				uri = fmt.Sprintf("%s?path=%s", hiuri.UriAction_NavigationMode, hugolocalpath)
			} else {
				contentedit = true
				uri = fmt.Sprintf("%s%s", hiuri.UriPage_Edit, hugolocalpath)
			}
		}

		entries[index] = fragment.DirEntry{
			Name:        direlement.Name,
			Uri:         uri,
			ContentEdit: contentedit,
		}
	}
	return entries
}

func NavigationMode(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	patharg := request.Form.Get("path")
	dirlocalpath := himodel.NewLocalPath(patharg)
	basepath := himodel.NewAbsolutePath(hiconfig.AppConfig().HugoProject())
	dirfullpath := himodel.NewFullPath(*basepath, *dirlocalpath)
	dirhelper := himodel.NewOsPathHelper(dirfullpath)

	// Assesment of path argument
	var err error = nil
	if !dirhelper.Exists() {
		err = fmt.Errorf("non-existing (local path %s): %s)", dirlocalpath.String(), dirfullpath.String())
	}
	if !dirhelper.IsDir() {
		err = fmt.Errorf("path %s (full path: %s) is not a directory", dirlocalpath.String(), dirfullpath.String())
	}
	if !dirhelper.IsSubPathOf(*basepath) {
		err = fmt.Errorf("%s is not a subpath of %s", dirfullpath.String(), basepath.String())
	}

	// Generate widget data
	menuwidget := fragment.BuildMenuState(fragment.MenuItem_Navigation)
	var dirdata []fragment.DirEntry
	var pathdata []fragment.DirEntry
	errortext := ""
	if err == nil {
		var direlements []himodel.SlashDirElement
		direlements, err = dirhelper.ReadDirElements()
		if err != nil {
			errortext = err.Error()
		}
		dirdata = dirwidgetdata(dirlocalpath, direlements)
		pathdata = pathwidgetdata(dirlocalpath)
	} else {
		errortext = err.Error()
	}

	mode := fragment.NavigationStruct{
		Menu:      menuwidget,
		Path:      &pathdata,
		Dir:       &dirdata,
		Errortext: errortext,
	}

	fragment.RenderNavigationMode(writer, mode)
}
