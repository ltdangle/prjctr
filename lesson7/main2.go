package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	myChan := make(chan int)

	go func() {
		defer wg.Done()
		v1 := <-myChan
		v2 := <-myChan
		v3 := <-myChan

		fmt.Println(v1, v2, v3)
	}()

	myChan <- 1
	myChan <- 2
	myChan <- 3

	wg.Wait()
}
