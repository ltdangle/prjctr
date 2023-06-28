package main

import (
	"context"
	"fmt"
	"sync"
)

// OrderTotalCalculator gorutine
func OrderTotalCalculator(ctx context.Context, orders chan *Order, wg *sync.WaitGroup) {
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
			ordr.Total = calculateOrderTotal(ordr)
			fmt.Printf("\nNew order #%d from '%s'. Products %d, total: $%d", counter, ordr.Customer, len(ordr.products), ordr.Total)
			counter++
		}
	}
}

// Calculates order total.
func calculateOrderTotal(order *Order) int {
	var total int
	for _, product := range order.products {
		total += product.price
	}
	return total
}
