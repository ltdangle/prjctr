package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ProductFinder(ctx context.Context, order *Order, wg *sync.WaitGroup) error {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second) // check payments every 2 seconds
	defer ticker.Stop()
	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nProduct finder cancelled.\n")
			return ctx.Err()
		case <-ticker.C:
			fmt.Printf("\nProductFinder has been looking for product for %d seconds ", counter)
			counter++
			if rand.Intn(10) == 0 {
				fmt.Printf("\nProductFinder error\n")
				return errors.New("ProductFinder error")
			}
		}
	}
}
