package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	u "net/url"
)

type TranslateResponse struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

// https://translate.googleapis.com/translate_a/single?client=gtx&sl=eng&tl=ru&dt=t&q=hi
func translateHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare parameters.
	text := r.URL.Query().Get("text")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	// Call api.
	var apiResponse string
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + from + "&tl=" + to + "&dt=t&q=" + u.QueryEscape(text)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		respByte, _ := io.ReadAll(resp.Body)
		apiResponse = string(respByte)
	}

	fmt.Println(apiResponse)

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
