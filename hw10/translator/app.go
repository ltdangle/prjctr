package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	trnsltr *Translator
	rspndr  *Responder
}

func NewApp(trnsltr *Translator, rspndr *Responder) *App {
	return &App{
		trnsltr: trnsltr,
		rspndr:  rspndr,
	}
}

func (app *App) httpHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	text := r.URL.Query().Get("text")

	if from == "" || to == "" || text == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Missing request parameters.")
		return
	}

	// Call api.
	err, translated := app.trnsltr.translate(from, to, text)
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
