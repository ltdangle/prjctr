package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type TranslationResponse struct {
	FromLng     string `json:"from"`
	ToLng       string `json:"to"`
	Source      string `json:"source"`
	Translation string `json:"translation"`
}

type Weather struct {
	City     string  `json:"city"`
	Temp     float64 `json:"temp"`
	Wind     float64 `json:"wind"`
	Humidity float64 `json:"humidity"`
}

func main() {
	port := ":8080"

	rspndr := NewResponder("2006-01-02 15:04:05")
	trnsltrApi := NewTranslatorApi(os.Getenv("LECTO_API_KEY"))
	weatherApi := NewWeatherApi(os.Getenv("WEATHER_API_KEY"))

	app := NewApp(weatherApi, trnsltrApi, rspndr)

	http.HandleFunc("/translate", app.translateHandler)
	http.HandleFunc("/weather", app.weatherHandler)
	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
