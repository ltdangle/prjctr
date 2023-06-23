package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pipe1 := make(chan string)
	pipe2 := make(chan string)
	pipe3 := make(chan string)
	go func() {
		counter := 0
		for text := range pipe1 {
			pipe2 <- text + fmt.Sprintf(" p1:%d", counter)
			counter++
		}
	}()

	go func() {
		counter := 0
		for text := range pipe2 {
			pipe3 <- text + fmt.Sprintf(" p2:%d", counter)
			counter++
		}
	}()

	go func() {
		for text := range pipe3 {
			fmt.Println(text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pipe1 <- scanner.Text()
	}

	close(pipe1)
	close(pipe2)
	close(pipe3)
}
