package main

import "net/http"

// Handler that will be verified
// It recieves a handler and returns a handler, because a middleware
// nest handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc
