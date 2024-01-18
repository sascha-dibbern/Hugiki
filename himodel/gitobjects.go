package himodel

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/sascha-dibbern/Hugiki/appconfig"
)

func runGit(arglist ...string) string {
	git := appconfig.AppConfig().GitCommand()

	cmd := exec.Command(git, arglist...)
	cmd.Dir = filepath.Clean(appconfig.AppConfig().HugoProject())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// error case : status code of command is different from 0
		log.Fatal("git checkout err:", err)
	}

	return out.String()
}

func CommitFiles(files ...string) {
	var args = append([]string{"comment"}, files...)
	runGit(args...)
}

func GitDiff(path string) string {
	return runGit("diff", path)
}
