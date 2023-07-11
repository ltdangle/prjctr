package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	u "net/url"
	"strings"
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

	translated := TranslateResponse{
		Text: "",
		From: from,
		To:   to,
	}

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

	// Parse translated text from api response.
	var result []interface{}
	var ctext []string

	err = json.Unmarshal([]byte(apiResponse), &result)
	if err != nil {
		fmt.Println("Error unmarshaling data")
	}

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				ctext = append(ctext, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		translated.Text = strings.Join(ctext, "")
	}

	json.NewEncoder(w).Encode(translated)
}

func main() {
	http.HandleFunc("/translate", translateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
