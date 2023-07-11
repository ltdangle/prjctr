package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TranslateResponse struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare parameters.
	text := r.URL.Query().Get("text")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	translate := TranslateResponse{
		Text: text,
		From: from,
		To:   to,
	}
	json.NewEncoder(w).Encode(translate)
}

func main() {
	http.HandleFunc("/translate", translateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

