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
	pipe3Go := func() {
		for t := range ticker.C {
			pipe2 <- fmt.Sprintf("Tick at %v", t)
		}
	}

	pipe2Go := func() {
		counter := 0
		for text := range pipe1 {
			pipe2 <- text + fmt.Sprintf(" p1:%d", counter)
			counter++
		}
	}

	pipe1Go := func() {
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

	go pipe3Go()
	go pipe2Go()
	go pipe1Go()
	go printerGo()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pipe1 <- scanner.Text()
	}

	close(pipe1)
	close(pipe2)
	close(pipe3)
}
