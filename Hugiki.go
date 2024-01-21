package main

import (
	//    "fmt"
	"bytes"
	"io"
	"net/http"

	"github.com/sascha-dibbern/Hugiki/hiconfig"
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
	http.ListenAndServe(hiconfig.AppConfig().ServerBinding(), mux)
}
