package main

import (
	"context"
	"fmt"
	"sync"
)

func ProductFinder(ctx context.Context, orders chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nProduct finder cancelled.\n")
			return
		case _, ok := <-orders:
			if !ok {
				return
			}
			fmt.Printf("Product has been found!")
			counter++
		}
	}
}
