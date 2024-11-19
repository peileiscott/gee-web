package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	routes map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		routes: make(map[string]HandlerFunc),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.routes[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
	}
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.routes[key] = handler
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
