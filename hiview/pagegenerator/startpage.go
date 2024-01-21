package pagegenerator

import (
	"fmt"

	"github.com/sascha-dibbern/Hugiki/hiproxy"
)

/*
 * Should not be called
 */
type StartPageGenerator struct {
}

func (generator StartPageGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	return "<html><head></head><body>Startpage for configuration</body></html>"
}

type EditConfigGenerator struct {
}

func (generator EditConfigGenerator) GenerateHtml(context *hiproxy.ProxyContext) string {
	result := fmt.Sprintf(`
	<html>
	<head>
	</head>
	<body>
	<h1>Configuration<h1>
	%s
	</body>
	</html>
	`, "x")
	return result
}
