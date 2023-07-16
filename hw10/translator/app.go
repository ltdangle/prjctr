package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	trnsltrApi *TranslatorApi
	weatherApi *WeatherApi
	rspndr     *Responder
}

func NewApp(weatherApi *WeatherApi, trnsltr *TranslatorApi, rspndr *Responder) *App {
	return &App{
		trnsltrApi: trnsltr,
		weatherApi: weatherApi,
		rspndr:     rspndr,
	}
}

func (app *App) translateHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	text := r.URL.Query().Get("text")

	if from == "" || to == "" || text == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Missing request parameters.")
		return
	}

	// Call api.
	err, translated := app.trnsltrApi.translate(from, to, text)
	if err != nil {
		app.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Convert api result to json.
	translatedJson, err := json.Marshal(translated)
	if err != nil {
		app.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return response.
	app.rspndr.Success(w, string(translatedJson))
}

func (app *App) weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare parameters.
	city := r.URL.Query().Get("city")

	if city == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Missing request parameters.")
		return
	}

	err, weather := app.weatherApi.Weather(city)
	if err != nil {
		app.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Convert api result to json.
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		app.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return response.
	app.rspndr.Success(w, string(weatherJson))
}
