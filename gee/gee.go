package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义request handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine  ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc //路由映射表
}

// New gee.Engine 的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 实现ServerHTPP接口
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUNT: %s\n", request.URL)
	}
}

// 添加路由function
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 封装GET请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 封装POST请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
