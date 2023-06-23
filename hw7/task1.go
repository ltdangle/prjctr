package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	pipe1 := make(chan int)
	pipe2 := make(chan float64)

	var wg sync.WaitGroup
	// Яка створює 3 горутини.
	wg.Add(3)

	// Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину.
	rndGo := func(writePipe chan<- int, iterations int) {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < iterations; i++ {
			rndNum := rand.Intn(100)
			writePipe <- rndNum
		}
		close(writePipe)
	}

	// Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал.
	avgGo := func(readPipe <-chan int, writePipe chan<- float64) {
		defer wg.Done()
		var sum int
		counter := 1
		for num := range readPipe {
			sum += num
			avg := float64(sum) / float64(counter)
			writePipe <- avg
			counter++
		}
		close(writePipe)
	}

	//	Третя горутина виводить середнє значення на екран.
	printGo := func(readPipe <-chan float64) {
		defer wg.Done()
		for avg := range readPipe {
			fmt.Printf("\nAvg: %.2f", avg)
		}
	}

	go rndGo(pipe1, 10)
	go avgGo(pipe1, pipe2)
	go printGo(pipe2)

	wg.Wait()

	fmt.Println()
}
