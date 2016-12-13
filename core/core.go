package core

import (
	"fmt"
	"net/http"

	"github.com/fudali113/doob/core/router"
)

var (
	_doob = &doob{
		filters: make([]Filter, 0),
		router:  &router.SimpleRouter{},
	}
	urlPrefixs      = []string{}
	staticFileCache = map[string][]byte{}
)

func Listen(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), _doob)
}

func AddFilter(fs ...Filter) {
	_doob.addFilter(fs...)
}

// 注册一个handler
func AddHandlerFunc(url string, handler interface{}, methods ...HttpMethod) {
	methodHandlerMap := make(map[string]interface{}, 0)
	for _, method := range methods {
		methodStr := string(method)
		if checkMethodStr(methodStr) {
			methodHandlerMap[methodStr] = handler
		} else {
			logger.Notice("%s method is unsupport", methodStr)
		}
	}
	_doob.addRestHandler(url, router.GetSimpleRestHandler(methodHandlerMap, handler))
}

func AddStaticPrefix(prefixUrls ...string) {
	urlPrefixs = append(urlPrefixs, prefixUrls...)
}
