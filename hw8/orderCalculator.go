package main

import (
	"context"
	"fmt"
	"sync"
)

// Gorutine to calculate order subtotal.
func OrderTotalCalculator(ctx context.Context, orders chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 1
	for ordr := range orders {
		orderTotal := calculateOrderTotal(ordr)
		productsCount := len(ordr.products)
		fmt.Printf("\nNew order #%d from '%s'. Products %d, total: $%d", counter, ordr.customer, productsCount, orderTotal)
		counter++
	}
	fmt.Println()
}

// Calculates order total.
func calculateOrderTotal(order order) int {
	var total int
	for _, product := range order.products {
		total += product.price
	}
	return total
}
