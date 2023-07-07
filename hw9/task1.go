package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "Missing city name", http.StatusBadRequest)
			return
		}

		resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=YOUR_API_KEY", city))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var weatherResponse WeatherResponse
		err = json.Unmarshal(body, &weatherResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(weatherResponse)
	})

	http.ListenAndServe(":8080", nil)
}
