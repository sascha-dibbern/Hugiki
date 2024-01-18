module github.com/Hugiki

go 1.21.1

require github.com/sascha-dibbern/Hugiki/hihandlers v0.0.0-00010101000000-000000000000

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/goccy/go-yaml v1.11.2 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/gookit/config/v2 v2.2.5 // indirect
	github.com/gookit/goutil v0.6.15 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/sascha-dibbern/Hugiki/appconfig v0.0.0-20240111213814-3341605cd241 // indirect
	github.com/sascha-dibbern/Hugiki/hiproxy v0.0.0-20240111213814-3341605cd241 // indirect
	github.com/sascha-dibbern/Hugiki/htmx v0.0.0-00010101000000-000000000000 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

replace github.com/sascha-dibbern/Hugiki/hiproxy => ./hiproxy

replace github.com/sascha-dibbern/Hugiki/hihandlers => ./hihandlers

replace github.com/sascha-dibbern/Hugiki/htmx => ./htmx

replace github.com/sascha-dibbern/Hugiki/appconfig => ./appconfig
