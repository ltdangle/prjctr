package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)
func main() {
	port := ":8080"

	taskList := NewTaskList()
	seedTasks(taskList, 20)

	rspndr:=NewResponder("2006-01-02 15:04:05")
	app := NewApp(rspndr)

	http.HandleFunc("/tasks", app.httpHandler)

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
