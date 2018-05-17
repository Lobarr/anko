// +build !appengine

package packages

import (
	"net/http"
)

func init() {
	Packages["net/http"] = map[string]interface{}{
		"DefaultClient":    http.DefaultClient,
		"DefaultTransport": http.DefaultTransport,
		"NewRequest":       http.NewRequest,
	}
	PackageTypes["net/http"] = map[string]interface{}{
		"Client":   http.Client{},
		"Cookie":   http.Cookie{},
		"Request":  http.Request{},
		"Response": http.Response{},
	}
}
