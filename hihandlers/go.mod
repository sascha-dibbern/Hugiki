module github.com/sascha-dibbern/Hugiki/hihandlers

go 1.21.1

replace github.com/sascha-dibbern/Hugiki/appconfig => ../appconfig

replace github.com/sascha-dibbern/Hugiki/hiproxy => ../hiproxy

replace github.com/sascha-dibbern/Hugiki/htmx => ../htmx

require (
	github.com/sascha-dibbern/Hugiki/appconfig v0.0.0-20240111213814-3341605cd241
	github.com/sascha-dibbern/Hugiki/hiproxy v0.0.0-20240111213814-3341605cd241
	github.com/sascha-dibbern/Hugiki/htmx v0.0.0-00010101000000-000000000000
)

require (
	github.com/fatih/color v1.14.1 // indirect
	github.com/goccy/go-yaml v1.11.2 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/gookit/config/v2 v2.2.4 // indirect
	github.com/gookit/goutil v0.6.14 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)
