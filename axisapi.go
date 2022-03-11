package axisapi

import (
	"fmt"
	"net/http"
)

// HandleFunc defines the request handler used by axisapi
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

// Engine is the uni handler for all requests
// 定义ServeHTTP方法捕获所有request
func (eg *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "|" + req.URL.Path
	if handler, exist := eg.router[key]; exist {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}

// New is the constructor of axisapi.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// 包内函数，添加路由
func (eg *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "|" + pattern
	eg.router[key] = handler
}

// GET defines the method to add GET request
func (eg *Engine) GET(pattern string, handler HandleFunc) {
	eg.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (eg *Engine) POST(pattern string, handler HandleFunc) {
	eg.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http Server
func (eg *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, eg)
}
