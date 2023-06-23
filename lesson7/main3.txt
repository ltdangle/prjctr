package main

import (
	"fmt"
	"time"
)

// goroutine function
func addSuffix(c chan string, suffix string) {
	val := <-c
	val += suffix
	c <- val
}

func main() {
	ch := make(chan string) // Buffered channel

	// send data to channel
	ch <- "value1"

	// start goroutine to process channel values
	go addSuffix(ch, "-suffix1")
	go addSuffix(ch, "-suffix2")

	// give the goroutine time to process
	time.Sleep(time.Second)

	// receive data from channel
	fmt.Println(<-ch)

	// close the channel
	close(ch)
}
