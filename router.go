package main

import (
	"net/http"
)

// Map of strings
// router (/api, /example, etc...) -> Handlers
// UPDATE -> Maps of maps, method -> route -> Handler
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Returns a handler or a boolean, the boolean is useful to know
// If the path exists or not, in our struct
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]                 // Checks if the route exists
	handler, methodExist := r.rules[path][method] // Checks if the method exists for that path
	return handler, methodExist, pathExist
}

// We have to implement this method, because if required when we pass this Router
// To the http.Handle method.The http.Handler interface rquires this method
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Compare handler from the request with our rules
	handler, methodExist, pathExist := r.FindHandler(request.URL.Path, request.Method)

	// URL doesn't exist
	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// URL exists but not with that method
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
