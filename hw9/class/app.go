package main

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type App struct {
	school *School
	rspndr *Responder
}

func NewApp(school *School, rspndr *Responder) *App {
	return &App{
		school: school,
		rspndr: rspndr,
	}
}

func (app *App) httpHandler(w http.ResponseWriter, r *http.Request) {
	teacher := r.Context().Value("teacher").(*Teacher)

	// Parse URL.
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		app.rspndr.Error(w, http.StatusBadRequest, "Invalid URL.")
		return
	}
	rawId := parts[2]

	// Error if student id is missing.
	if rawId == "" {
		app.rspndr.Error(w, http.StatusBadRequest, "Missing student id.")
		return
	}

	// Error if student id is malformed.
	id, err := strconv.Atoi(rawId)
	if err != nil {
		app.rspndr.Error(w, http.StatusBadRequest, "Malformed student id.")
		return
	}

	student := app.school.FindStudentById(id)

	if student == nil {
		app.rspndr.Error(w, http.StatusBadRequest, "No student with this id.")
		return
	}

	// Check that teacher hass access to student's class.
	if student.class != teacher.Class {
		app.rspndr.Error(w, http.StatusBadRequest, "Teacher not authorized to access the class.")
		return
	}

	app.rspndr.Success(w, student)
}

func (app *App) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		teacher := app.school.FindTeacher(username)

		// Teacher not found in school, unauthorize.
		if teacher == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Teacher provided wrong password, unauthorize.
		if teacher.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Pass teacher downstream via context.
		ctx := context.WithValue(r.Context(), "teacher", teacher)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
