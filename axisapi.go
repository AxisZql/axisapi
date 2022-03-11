package axisapi

/*
*Author：AxisZql
*Date: 2022-3-10
 */

import (
	"net/http"
)

// HandleFunc defines the request handler used by axisapi
type HandleFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// Engine is the uni handler for all requests
// 定义ServeHTTP方法捕获所有request
func (eg *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	eg.router.handle(ctx)
}

// New is the constructor of axisapi.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 包内函数，添加路由
func (eg *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	eg.router.addRoute(method, pattern, handler)
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
