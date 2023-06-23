package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	pipe1 := make(chan string)
	pipe2 := make(chan string)
	go func() {
		counter := 0
		for text := range pipe1 {
			// fmt.Printf("Pipe1 %d: %s\n", counter, text)
			pipe2 <- text + strconv.Itoa(counter)
			counter++
		}
	}()

	go func() {
		counter := 0
		for text := range pipe2 {
			fmt.Printf("Pipe2 %d: %s\n", counter, text)
			counter++
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pipe1 <- scanner.Text()
	}
}
