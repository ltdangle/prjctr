package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WeatherApi struct {
	apiKey string
}

func NewWeatherApi(apiKey string) *WeatherApi {
	return &WeatherApi{apiKey: apiKey}
}
func (w *WeatherApi) Weather(city string) (error, *WeatherApiResponse) {
	// Call api.
	var apiResponse string
	// url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", w.apiKey, city)
	url := "https://api.weatherapi.com/v1/forecast.json?key=s&q=s&days=1"
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("WeatherApi url: " + err.Error()), nil
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("WeatherApi returned non-OK status %d", resp.StatusCode)), nil
	} else {
		respByte, _ := io.ReadAll(resp.Body)
		apiResponse = string(respByte)
	}

	fmt.Println(apiResponse)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(apiResponse), &result)
	if err != nil {
		return errors.New("WeatherApi unmarshal: " + err.Error()), nil
	}

	// Convert api result into our json response.
	// TODO:
	var weather = &WeatherApiResponse{
		Temp:     result["current"].(map[string]interface{})["temp_c"].(float64),
		Wind:     result["current"].(map[string]interface{})["wind_kph"].(float64),
		Humidity: result["current"].(map[string]interface{})["humidity"].(float64),
	}

	return nil, weather
}
