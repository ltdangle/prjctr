package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	rspndr := NewResponder("2006-01-02 15:04:05")

	dateInp := r.URL.Query().Get("date")

	// Validate date is in correct format.
	dateRegex := `^\d{4}-\d{2}-\d{2}$`
	rgxp, _ := regexp.Compile(dateRegex)
	dateStr := rgxp.FindString(dateInp)

	if dateStr == "" {
		rspndr.Error(w, http.StatusBadRequest, "Malformed date string.")
		return
	}

	// Validate date is a valid date.
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		rspndr.Error(w, http.StatusBadRequest, "Incorrect date string.")
		return
	}

}

func main() {
	port := ":8080"

	taskList := NewTaskList()
	seedTasks(taskList, 20)

	http.HandleFunc("/tasks", httpHandler)

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func seedTasks(taskList *taskList, numTasks int) {
	for i := 0; i < numTasks; i++ {
		rnd := rand.Intn(10)
		task := &task{
			name: fmt.Sprintf("Task_%d", rnd),
			due:  time.Now().AddDate(0, 0, rnd),
		}
		taskList.add(task)
	}
}
