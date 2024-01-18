package himodel

import (
	"log"
	"os"
	"path/filepath"

	"github.com/sascha-dibbern/Hugiki/appconfig"
)

func LoadTextFromFile(relativepath string) string {
	path := filepath.Clean(appconfig.AppConfig().HugoProject() + "/" + relativepath)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SaveTextToFile(relativepath string, text string) {
	path := filepath.Clean(appconfig.AppConfig().HugoProject() + relativepath)
	err := os.WriteFile(path, []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveContentMarkdown(path_under_content string, text string) {
	SaveTextToFile("content/"+path_under_content+".md", text)
}
