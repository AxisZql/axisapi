package axisapi

/*
*Author：AxisZql
*Date: 2022-3-10
 */

import (
	"log"
	"net/http"
)

// HandleFunc defines the request handler used by axisapi
type HandleFunc func(*Context)

// RouterGroup defines the router group
type RouterGroup struct {
	prefix      string       // router prefix
	middlewares []HandleFunc //support middleware
	parent      *RouterGroup // support nesting (支持嵌套）
	engine      *Engine
}

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup //store all groups
}

// Engine is the uni handler for all requests
// 定义ServeHTTP方法捕获所有request,本框架的核心（在go中实现接口方法的struct都可以强制转化为接口类型）
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

// Group is defined to create a new RouterGroup
// all groups share the same Engine instance,engine can use the method
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group, //支持嵌套路由组
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	// 调用engine实现的addRouter方法，保证之前的路由逻辑不会受影响
	group.engine.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http Server
func (eg *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, eg)
}
