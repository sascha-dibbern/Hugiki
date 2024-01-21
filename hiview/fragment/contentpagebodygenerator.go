package fragment

import (
	converthtml "github.com/sascha-dibbern/Hugiki/hiconverters/html"
	"github.com/sascha-dibbern/Hugiki/hiproxy"
)

// ...<body...>xyz</body>... -> xyz
type ContentPageBodyGenerator struct {
}

func (generator ContentPageBodyGenerator) GenerateHtml(htmlInput string, context *hiproxy.ProxyContext) string {
	return converthtml.ExtractBodycontent(htmlInput)
}
