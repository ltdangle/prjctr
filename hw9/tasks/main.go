package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)
func main() {
	port := ":8080"

	taskList := NewTaskList()
	seedTasks(taskList, 5)

	// Dump data in console for debugging
	taskListJsonData, _ := json.MarshalIndent(taskList, "", " ")
	fmt.Println(string(taskListJsonData))

	rspndr:=NewResponder("2006-01-02 15:04:05")
	app := NewApp(taskList, rspndr)

	http.HandleFunc("/tasks", app.httpHandler)

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func seedTasks(taskList *TaskList, numTasks int) {
	for i := 0; i < numTasks; i++ {
		rnd := rand.Intn(100)
		task := &Task{
			Name: fmt.Sprintf("Task_%d", rnd),
			Due:  time.Now().AddDate(0, 0, rnd),
		}
		taskList.add(task)
	}
}
