package main

import "time"

// Task stuct.
type task struct {
	name string
	due time.Time
}

// Tasklist stuct.
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
		if tm.Year() == task.due.Year() && tm.Month() == task.due.Month() && tm.Day() == task.due.Day() {
			matches = append(matches, task)
		}
	}

	return matches
}
