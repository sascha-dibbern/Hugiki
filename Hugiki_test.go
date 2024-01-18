package main

import "testing"

func TestMakeHugikiHtml(t *testing.T) {
	result := "" //makeHugikiHtml("<body></body>")
	want := "<body><hugiki/></body>"
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}
