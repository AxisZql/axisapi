package axisapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// Context is a shortcut for record all information of the current HTTP session（只记录当前http会话）
// It's the most important part for axisapi
type Context struct {
	// origin object
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandleFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next 不是所有的handler都会调用 Next()。
//手工调用 Next()，一般用于在请求前后各实现一些行为。如果中间件只作用于请求前，可以省略调用Next()，算是一种兼容性比较好的写法吧。
func (ctx *Context) Next() {
	ctx.index++
	s := len(ctx.handlers)
	for ; ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

func (ctx *Context) Fail(code int, err string) {
	// 出现错误是直接跳过为执行的中间件
	ctx.index = len(ctx.handlers)
	ctx.JSON(code, H{"msg": err})
}

// Param 获取动态路由中的参数 Get the parameters in the dynamic route
func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]
	return value
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// Status is used to set the HTTP response code
func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

//--------------------------
//	Setting up different formats HTTP Responses
//--------------------------

func (ctx *Context) String(code int, format string, a ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
