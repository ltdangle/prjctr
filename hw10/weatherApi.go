package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Weather struct.
type Weather struct {
	City     string  `json:"city"`
	Temp     float64 `json:"temp"`
	Wind     float64 `json:"wind"`
	Humidity float64 `json:"humidity"`
}

// Weather api response.
type WeatherApiData struct {
	Current struct {
		TempC    float64 `json:"temp_c"`
		WindKph  float64 `json:"wind_kph"`
		Humidity float64 `json:"humidity"`
	} `json:"current"`
}

type WeatherApi struct {
	apiKey string
}

func NewWeatherApi(apiKey string) *WeatherApi {
	return &WeatherApi{apiKey: apiKey}
}
func (w *WeatherApi) Weather(city string) (error, *Weather) {
	// Call api.
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", w.apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("WeatherApi url: " + err.Error()), nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WeatherApi returned non-OK status %d", resp.StatusCode), nil
	}

	respByte, _ := io.ReadAll(resp.Body)

	// Unmarshall api response.
	var data WeatherApiData
	err = json.Unmarshal(respByte, &data)
	if err != nil {
		return errors.New("WeatherApi unmarshal: " + err.Error()), nil
	}

	// Convert api result into our json response.
	var weather = &Weather{
		Temp:     data.Current.TempC,
		Wind:     data.Current.WindKph,
		Humidity: data.Current.Humidity,
	}

	return nil, weather
}
