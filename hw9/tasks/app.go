package main

import (
	"net/http"
	"regexp"
	"time"
)

type App struct {
	taskList *TaskList
	rspndr   *Responder
}

func NewApp(taskList *TaskList, rspndr *Responder) *App {
	return &App{
		taskList: taskList,
		rspndr:   rspndr,
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
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		app.rspndr.Error(w, http.StatusBadRequest, "Incorrect date string.")
		return
	}

	// Find tasks on the specified date.
	tasks := app.taskList.findTasks(date)
	app.rspndr.Success(w,  tasks)
}
