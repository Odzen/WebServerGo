package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Middlewares have the same structure that the handlers
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Checking Auth")

			// Check if the middleware passed
			// f -> the next middleware
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
