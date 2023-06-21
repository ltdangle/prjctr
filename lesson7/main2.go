package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	myChan := make(chan int)

	// Wait for 1 gorutine to finish
	wg.Add(1)

	// First, init channel read via gorutine, otherwise it will block.
	go func() {
		defer wg.Done()
		v1 := <-myChan
		v2 := <-myChan
		v3 := <-myChan

		fmt.Println(v1, v2, v3)
	}()

	// Second, write to channel.
	myChan <- 1
	myChan <- 2
	myChan <- 3

	// Block until our 1 gorutine is finished.
	wg.Wait()
}
