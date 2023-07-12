package main

import "time"

// Task stuct.
type Task struct {
	Name string    `json:"name"`
	Due  time.Time `json:"due"`
}

// Tasklist stuct.
type TaskList struct {
	tasks []*Task
}

func NewTaskList() *TaskList {
	return &TaskList{}
}

func (t *TaskList) add(task *Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) findTasks(tm time.Time) []*Task {
	var matches []*Task

	for _, task := range t.tasks {
		if tm.Year() == task.Due.Year() && tm.Month() == task.Due.Month() && tm.Day() == task.Due.Day() {
			matches = append(matches, task)
		}
	}

	return matches
}
