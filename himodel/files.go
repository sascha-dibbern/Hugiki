package himodel

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
)

type SlashPath interface {
	String() string
	Elements() []string
}

type ExtendableSlashPath interface {
	String() string
	Elements() []string
	Extend(entry string)
}

// An local slashpath starting without slash
type LocalPath struct {
	path  string
	empty bool
}

func NewLocalPath(path string) *LocalPath {
	empty := false
	if path == "" {
		empty = true
	}
	rexp, _ := regexp.Compile("^/")
	localpath := rexp.ReplaceAllString(path, "")
	return &LocalPath{
		path:  localpath,
		empty: empty,
	}
}

func (p *LocalPath) String() string {
	return p.path
}

func (p *LocalPath) Elements() []string {
	if p.empty {
		return []string{}
	}
	return strings.Split(p.path, "/")
}

func (p *LocalPath) Extend(item string) {
	p.path += "/" + item
	p.empty = false
}

// An absolute slashpath possbibly starting with slash
type AbsolutePath struct {
	path string
}

func NewAbsolutePath(path string) *AbsolutePath {
	return &AbsolutePath{
		path: path,
	}
}

func (p *AbsolutePath) String() string {
	return p.path
}

func (p *AbsolutePath) Elements() []string {
	rexp, _ := regexp.Compile("^/")
	localpath := rexp.ReplaceAllString(p.path, "")
	return strings.Split(localpath, "/")
}

func (p *AbsolutePath) Extend(item string) {
	p.path += "/" + item
}

type FullPath struct {
	basepath  AbsolutePath
	localpath LocalPath
}

func NewFullPath(basepath AbsolutePath, localpath LocalPath) *FullPath {
	return &FullPath{
		basepath:  basepath,
		localpath: localpath,
	}
}

func (p *FullPath) BasePath() AbsolutePath {
	return p.basepath
}

func (p *FullPath) LocalPath() LocalPath {
	return p.localpath
}

func (p *FullPath) String() string {
	return p.basepath.String() + "/" + p.localpath.String()
}

func (p *FullPath) Elements() []string {
	return strings.Split(p.String(), "/")
}

type SlashDirElement struct {
	Name      string
	IsDir     bool
	LocalPath LocalPath
	FullPath  FullPath
}

type SlashDirPathHelper interface {
	Exists() bool
	IsDir() bool
	IsSubPathOf(base AbsolutePath) bool
	ReadDirElements() ([]SlashDirElement, error)
}

type OsPathHelper struct {
	fullpath *FullPath
}

func NewOsPathHelper(fullpath *FullPath) OsPathHelper {
	return OsPathHelper{
		fullpath: fullpath,
	}
}

func (h OsPathHelper) Exists() bool {
	path := h.fullpath.String()
	_, err := os.Stat(filepath.Clean(path))
	return !os.IsNotExist(err)
}

func (h OsPathHelper) IsDir() bool {
	info, _ := os.Stat(filepath.Clean(h.fullpath.String()))
	return info.IsDir()
}

// base should be a full sub-path from the beginning of the tested path
func (h OsPathHelper) IsSubPathOf(base AbsolutePath) bool {
	basestring := base.String()
	baselen := len(basestring)
	fullpath := h.fullpath.String()
	fullpathbase := fullpath[0:baselen]
	return basestring == fullpathbase
}

func (h OsPathHelper) ReadDirElements() ([]SlashDirElement, error) {
	osdirpath := filepath.Clean(h.fullpath.String())
	basepath := h.fullpath.BasePath()
	baselocalpath := h.fullpath.LocalPath()

	entries, err := os.ReadDir(osdirpath)
	if err != nil {
		// Return no entries if error happened
		entries = make([]fs.DirEntry, 0)
	}

	result := make([]SlashDirElement, len(entries))
	for index, entry := range entries {
		entryname := entry.Name()
		entrylocalpath := baselocalpath
		entrylocalpath.Extend(entryname)
		entryfullpath := *NewFullPath(basepath, entrylocalpath)
		entrypath := filepath.Clean(entryfullpath.String())
		entryinfo, _ := os.Stat(entrypath)
		entryisdir := entryinfo.IsDir()
		result[index] = SlashDirElement{
			Name:      entryname,
			IsDir:     entryisdir,
			LocalPath: entrylocalpath,
			FullPath:  entryfullpath,
		}
	}
	return result, err
}

func LoadTextFromFile(localpath string) string {
	path := filepath.Clean(hiconfig.AppConfig().HugoProject() + "/" + localpath)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SaveTextToFile(localpath string, text string) {
	path := filepath.Clean(hiconfig.AppConfig().HugoProject() + localpath)
	err := os.WriteFile(path, []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveContentMarkdown(path_under_content string, text string) {
	SaveTextToFile("content/"+path_under_content, text)
}

func IsMarkdownFile(info fs.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	name := info.Name()
	namelen := len(name)
	filetype := name[namelen-3 : namelen]
	if strings.ToLower(filetype) == ".md" {
		return true
	}
	return false
}

func SearchContentFiles(searchregexp string) []string {
	var searchresult []string
	if searchregexp == "" {
		return searchresult
	}
	regexp := regexp.MustCompile(searchregexp)
	rootpath := filepath.Clean(hiconfig.AppConfig().HugoProject() + "/content")
	rootpathlen := len(rootpath)
	err := filepath.Walk(rootpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if IsMarkdownFile(info) {
			localpath := path[rootpathlen:]
			text := LoadTextFromFile("/content/" + localpath)
			found := regexp.FindString(text)
			if found != "" {
				cleanlocalpath := strings.ReplaceAll(localpath, "\\", "/")
				searchresult = append(searchresult, cleanlocalpath)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return searchresult
}
