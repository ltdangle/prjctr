package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a new context with a timeout of 1 second
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel() // It's important to cancel when we're done to release resources

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello")
	}()

	// Here we wait until the context is done.
	<-ctx.Done()
	fmt.Println("Context expired, stopping operation")
}

