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

	ticker := time.NewTicker(100 * time.Millisecond)
	for t := range ticker.C {
		// Poll worker states and print statistics.
		fmt.Print("\033[H\033[2J")
		fmt.Println("Tick: ", t)
		for _, s := range workerStates {
			fmt.Printf("\nWorker: %s", s.name)
			fmt.Printf("\nLoad: %s\n", prettyLoad(s.load))
		}
	}
}

// Worker goroutine.
func workerRoutine(s *workerState) {
	ticker := time.NewTicker(10 * time.Millisecond)
	for _ = range ticker.C {
		// Set worker load randomly.
		s.load = rand.Intn(9)
	}
}
func prettyLoad(l int) string {
	str := "|"
	for i := 0; i < l; i++ {
		str += "|"
	}
	return str
}
