package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
