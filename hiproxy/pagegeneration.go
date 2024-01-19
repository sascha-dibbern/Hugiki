package hiproxy

/**********************************
 * Proxy page generator
 **********************************/

type ProxyPageGenerator interface {
	GenerateHtml(proxiedcontent string, context *ProxyContext) string
}
