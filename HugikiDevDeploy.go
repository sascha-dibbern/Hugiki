package main

import (
    //"fmt"
    "os/exec"
	"log"
)

type customOutput struct{}

func (c customOutput) Write(p []byte) (int, error) {
	log.Print("> ", string(p))
	return len(p), nil
}

func main() {

    prg	:= "./Hugiki"
    //prg := "E:/Users/sascha/Documents/Projects/Hugiki/Hugiki.exe"
    arg1 := "--config"
    arg2 := "./hugiki.yml"
    arg3 := "--dev"
    arg4 := "true"
	
	for {
		log.Println("Child Hugiki starting")
		cmd := exec.Command(prg, arg1, arg2, arg3, arg4)
		cmd.Stdout = customOutput{}

		if err := cmd.Run(); err != nil {
			log.Println(err.Error())
		}
		log.Println("Child Hugiki stopped")
	}
}