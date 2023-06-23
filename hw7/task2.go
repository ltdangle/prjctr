package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//   Перша горутина виводить найбільше й найменше числа на екран.

	// Яка створює 2 горутини.
	var wg sync.WaitGroup
	wg.Add(2)

	pipe1 := make(chan int)
	pipe2 := make(chan int)

	// Перша горутина генерує випадкові числа в заданому діапазоні й надсилає їх через канал у другу горутину.
	rndGo := func(writePipe chan<- int, readPipe <-chan int, iterations int) {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < iterations; i++ {
			rndNum := rand.Intn(100)
			writePipe <- rndNum
		}
		close(writePipe)

		fmt.Println("Min: ", <-readPipe)
		fmt.Println("M: ", <-readPipe)
	}

	// Друга горутина отримує випадкові числа і знаходить найбільше й найменше число, після чого надсилає його назад у першу горутину через канал.
	minMax := func(readPipe <-chan int, writePipe chan<- int) {
		defer wg.Done()
		writePipe <- 10
		writePipe <- 100
		close(writePipe)
	}

	go rndGo(pipe1, pipe2, 10)
	go minMax(pipe2, pipe1)

	wg.Wait()
}

