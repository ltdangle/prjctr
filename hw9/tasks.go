package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tasklist stuct.
type task struct {
	name string
	date time.Time
}

type taskList struct {
	tasks []*task
}

func NewTaskList() *taskList {
	return &taskList{}
}

func (t *taskList) add(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *taskList) findTasks(tm time.Time) []*task {
	var matches []*task

	for _, task := range t.tasks {
		if tm.Year() == task.date.Year() && tm.Month() == task.date.Month() && tm.Day() == task.date.Day() {
			matches = append(matches, task)
		}
	}

	return matches
}

func seedTasks(taskList *taskList, numTasks int) {
	for i := 0; i < numTasks; i++ {
		rnd := rand.Intn(10)
		task := &task{
			name: fmt.Sprintf("Task_%d", rnd),
			date: time.Now().AddDate(0, 0, rnd),
		}
		taskList.add(task)
	}
}
func main() {
	taskList := NewTaskList()
	taskList.add(&task{date: time.Now(), name: "First task"})
	taskList.add(&task{date: time.Now().AddDate(0, 0, 1), name: "Second task"})
	foundTasks := taskList.findTasks(time.Now())
	fmt.Printf("\nFound tasks: %v", foundTasks)
}
