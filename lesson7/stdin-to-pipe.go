package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pipe := make(chan string)
	go func() {
		counter := 0
		for text := range pipe {
			fmt.Printf("%d: %s\n", counter, text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pipe <- scanner.Text()
	}

}
