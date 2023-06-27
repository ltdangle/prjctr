package main

import (
	"context"
	"fmt"
	"sync"
)

func OrderTotalCalculator(ctx context.Context, orders chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nOrder calculator cancelled.\n")
			return
		case ordr, ok := <-orders:
			if !ok {
				return
			}
			fmt.Printf("\nNew order #%d from '%s'. Products %d, total: $%d", counter, ordr.customer, len(ordr.products), calculateOrderTotal(ordr))
			counter++
		}
	}
}

// Calculates order total.
func calculateOrderTotal(order order) int {
	var total int
	for _, product := range order.products {
		total += product.price
	}
	return total
}
