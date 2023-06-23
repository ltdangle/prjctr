package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	pipe1 := make(chan string)
	pipe2 := make(chan string)
	pipe3 := make(chan string)

	ticker := time.NewTicker(3 * time.Second)
	tickerGo := func(pipe chan<- string) {
		for t := range ticker.C {
			pipe <- fmt.Sprintf("Tick at %v", t)
		}
	}

	pipe1Go := func(readPipe <-chan string, writePipe chan<- string) {
		counter := 0
		for text := range readPipe {
			writePipe <- text + fmt.Sprintf(" p1:%d", counter)
			counter++
		}
	}

	pipe2Go := func(readPipe <-chan string, writePipe chan<- string) {
		counter := 0
		for text := range pipe2 {
			pipe3 <- text + fmt.Sprintf(" p2:%d", counter)
			counter++
		}
	}

	printerGo := func() {
		for text := range pipe3 {
			fmt.Println(text)
		}
	}

	go pipe1Go(pipe1, pipe2)
	go pipe2Go(pipe2, pipe3)
	go tickerGo(pipe1)
	go printerGo()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pipe1 <- scanner.Text()
	}

	close(pipe1)
	close(pipe2)
	close(pipe3)
}
