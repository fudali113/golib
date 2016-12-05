package doob

import (
	"log"
	"net/http"

	"github.com/fudali113/doob/core"
)

const (
	GET     core.HttpMethod = "get"
	POST    core.HttpMethod = "post"
	PUT     core.HttpMethod = "put"
	DELETE  core.HttpMethod = "delete"
	OPTIONS core.HttpMethod = "options"
	HEAD    core.HttpMethod = "head"
)

/**
 * 启动server
 */
func Start(port int) {
	log.Printf("server is starting , listen port is %d", port)
	err := core.Listen(port)
	if err != nil {
		log.Printf("start is fail => %s", err.Error())
	}
}

/**
 * 注册一个handler
 */
func AddHandlerFunc(url string, handler http.HandlerFunc, tms ...core.HttpMethod) {
	core.AddHandlerFunc(url, handler, tms...)
}

func Get(url string, handler http.HandlerFunc) {
	core.AddHandlerFunc(url, handler, GET)
}
func Post(url string, handler http.HandlerFunc) {
	core.AddHandlerFunc(url, handler, POST)
}
func Put(url string, handler http.HandlerFunc) {
	core.AddHandlerFunc(url, handler, PUT)
}
func Delete(url string, handler http.HandlerFunc) {
	core.AddHandlerFunc(url, handler, DELETE)
}

/**
 * 添加一个过滤器
 */
func AddFilter(fs ...core.Filter) {
	core.AddFilter(fs...)
}