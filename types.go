package main

import "net/http"

// Handler that will be verified
// It recieves a handler and returns a handler, because a middleware
// nest handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string
	Email string
	Phone string
}

type MetaData interface{}
