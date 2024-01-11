package main

import (
//    "fmt"
	"net/http"
    "bytes"
    "io"
//	"regexp"
	"github.com/sascha-dibbern/Hugiki/hihandlers"
)

// To be deleted
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func main() {
	mux := http.NewServeMux()
	hihandlers.Setup(mux)
	http.ListenAndServe(":3000", mux)	
}