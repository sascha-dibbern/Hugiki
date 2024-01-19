package html

import "regexp"

const upWithBodyStartTagRXPS = ".*<body.*>"

var upWithBodyStartTagRegexp = regexp.MustCompile(upWithBodyStartTagRXPS)

const fromWithBodyEndTagRXPS = "</body>.*"

var fromWithBodyEndTagRegexp = regexp.MustCompile(fromWithBodyEndTagRXPS)
