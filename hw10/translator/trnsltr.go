package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type LectoItem struct {
	To         string   `json:"to"`
	Translated []string `json:"translated"`
}

type LectoResponse struct {
	Translations         []LectoItem `json:"translations"`
	From                 string      `json:"from"`
	TranslatedCharacters int         `json:"translated_characters"`
}

func translate(from string, to string, text string) (error, *Translation) {
	url := "https://api.lecto.ai/v1/translate/text"
	data := fmt.Sprintf(`{
	"texts": ["%s"],
	"to": ["%s"],
	"from": "%s"
}`, text, to, from)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err, nil
	}

	req.Header.Set("X-API-Key", os.Getenv("LECTO_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	responseJSON := string(body)

	var lectoResponse LectoResponse
	err = json.Unmarshal([]byte(responseJSON), &lectoResponse)
	if err != nil {
		return err, nil
	}

	translation := &Translation{FromLng: from, ToLng: to, Source: text, Translation: lectoResponse.Translations[0].Translated[0]}
	return nil, translation
}
