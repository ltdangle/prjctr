package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	pipe1 := make(chan int)
	pipe2 := make(chan int)

	rndGo := func() {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			rndNum := rand.Intn(100)
			fmt.Printf("Generated %d: %d\n", i, rndNum)
			pipe1 <- rndNum
		}
		close(pipe1)
		min := <-pipe2
		max := <-pipe2
		fmt.Printf("Min: %d\n", min)
		fmt.Printf("Max: %d\n", max)
	}

	minMax := func() {
		defer wg.Done()
		min, max := 100, 0
		for num := range pipe1 {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		pipe2 <- min
		pipe2 <- max
	}

	wg.Add(2)
	go rndGo()
	go minMax()
	wg.Wait()
}
