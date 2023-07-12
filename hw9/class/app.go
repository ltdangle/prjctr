package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type App struct {
	class  *Class
	rspndr *Responder
}

func NewApp(class *Class, rspndr *Responder) *App {
	return &App{
		class:  class,
		rspndr: rspndr,
	}
}

func (app *App) httpHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL.
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		app.rspndr.Error(w, http.StatusBadRequest, "Invalid URL.")
		return
	}
	rawId := parts[2]

	// Error if student id is missing
	if rawId == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Missing student id.")
		return
	}

	// Error if student id is malformed
	id, err := strconv.Atoi(rawId)
	if err != nil {
		app.rspndr.Error(w, http.StatusBadRequest, "Malformed student id.")
		return
	}

	student := app.class.findById(id)
	app.rspndr.Success(w, student)
}
