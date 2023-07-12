package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// JSON response object.
type Response struct {
	Date    string      `json:"date"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func NewResponse() *Response {
	return &Response{
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}
}

// Responder.
type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (rspndr *Responder) Error(w http.ResponseWriter, statusCode int, error string) {
	// Build json response.
	r := NewResponse()
	r.Message = error

	// Send response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(r)
}

// func (rspndr *Responder) Success(w http.ResponseWriter, statusCode int, error string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"error": error,
// 	})
// }
