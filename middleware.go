package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Middlewares have the same structure that the handlers
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Auth")

			// Check if the middleware passed
			// f -> the next middleware
			if flag {
				// Allows the middleware to jump up to the next middleware or Handler
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			// Allows the middleware to jump up to the next middleware or Handler
			f(w, r)
		}
	}
}
