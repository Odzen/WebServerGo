package main

import (
	"encoding/json"
	"net/http"
)

// Handler that will be verified
// It recieves a handler and returns a handler, because a middleware
// nest handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc

// `json:"name"` -> It changes the field names if I decode this to JSON
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}
