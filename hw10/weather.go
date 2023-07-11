package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	fmt.Println("WEATHER_API_KEY", apiKey)

	city := r.URL.Query().Get("city")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", apiKey, city)
	fmt.Println("Api url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
	return
	// resp, err := http.Get(url)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// fmt.Println(body)
	//
	// var weatherResponse WeatherResponse
	// err = json.Unmarshal(body, &weatherResponse)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(weatherResponse)
	//
	// json.NewEncoder(w).Encode(weatherResponse)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
