package pagegenerator

import "github.com/sascha-dibbern/Hugiki/hiproxy"

type RootPageGenerator struct {
}

func (generator RootPageGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	return "<html><head><meta http-equiv=\"refresh\" content=\"0; url='/hugiki/'\"/></head><body>x</body></html>"
}
