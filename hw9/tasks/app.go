package main

import (
	"net/http"
	"regexp"
	"time"
)

type App struct{
	rspndr Responder
}
func NewApp(rspndr *Responder)*App{
	return &App{
		rspndr: *rspndr,
	}
}

func (app *App) httpHandler(w http.ResponseWriter, r *http.Request) {

	dateInp := r.URL.Query().Get("date")

	// Validate date is in correct format.
	dateRegex := `^\d{4}-\d{2}-\d{2}$`
	rgxp, _ := regexp.Compile(dateRegex)
	dateStr := rgxp.FindString(dateInp)

	if dateStr == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Malformed date string.")
		return
	}

	// Validate date is a valid date.
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		app.rspndr.Error(w, http.StatusBadRequest, "Incorrect date string.")
		return
	}

}

