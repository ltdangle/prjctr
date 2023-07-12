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
	return &Response{}
}

// Responder.
type Responder struct {
	dateFormat string
}

func NewResponder(dateFormat string) *Responder {
	return &Responder{dateFormat: dateFormat}
}

func (rspndr *Responder) Error(w http.ResponseWriter, statusCode int, error string) {
	// Build json response.
	r := NewResponse()
	r.Date = time.Now().Format(rspndr.dateFormat)
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
