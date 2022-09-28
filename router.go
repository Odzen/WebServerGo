package main

import (
	"net/http"
)

// Map of strings
// router (/api, /example, etc...) -> Handlers
type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

// Returns a handler or a boolean, the boolean is useful to know
// If the path exists or not, in our struct
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

// We have to implement this method, because if required when we pass this Router
// To the http.Handle method.The http.Handler interface rquires this method
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Compare handler from the request with our rules
	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}
