package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// OrderTotalCalculator gorutine
func OrderTotalCalculator(ctx context.Context, order *Order) error {
	ticker := time.NewTicker(1 * time.Second)
	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nOrder calculator cancelled.\n")
			return ctx.Err()
		case <-ticker.C:
			fmt.Printf("\nOrder calculator has been calculating  %d seconds ", counter)
			counter++
			if rand.Intn(10) == 0 {
				fmt.Printf("\nOrder calculator error\n")
				return errors.New("Order calculator error")
			}
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
