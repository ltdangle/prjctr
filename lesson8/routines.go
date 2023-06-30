package main

import (
	"fmt"
	"math/rand"
	"time"
)

type workerState struct {
	name string
	load int
}
type workers []*workerState

func main() {
	// Setup array of worker states.
	workerStates := []*workerState{
		&workerState{name: "worker1"},
		&workerState{name: "worker2"},
		&workerState{name: "worker3"},
	}
	// Run workers.
	go workerRoutine(workerStates[0])
	go workerRoutine(workerStates[1])
	go workerRoutine(workerStates[2])

	// Poll worker states and print statistics.
	for i := 0; i < 20; i++ {
		// Clear screen.
		fmt.Print("\033[H\033[2J")
		for _, s := range workerStates {
			fmt.Printf("\nWorker: %s", s.name)
			fmt.Printf("\nLoad: %d\n\n", s.load)
		}
		time.Sleep(1 * time.Second)
	}
}

// Worker goroutine.
func workerRoutine(s *workerState) {
	for {
		// Set worker load randomly.
		s.load = rand.Intn(10)
		time.Sleep(1 * time.Second)
	}
}
