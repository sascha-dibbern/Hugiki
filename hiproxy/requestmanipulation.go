package hiproxy

import (
	"net/http"
)

/**********************************
 * Request manipulation
 **********************************/

type RequestManipulator interface {
	GenerateBackendUrl(request *http.Request) string
}
