package main

import (
	"fmt"
	"net/http"
)

type App struct {
	rspndr *Responder
}

func NewApp(rspndr *Responder) *App {
	return &App{
		rspndr: rspndr,
	}
}

func (app *App) httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	app.rspndr.Success(w,"hi")
}
