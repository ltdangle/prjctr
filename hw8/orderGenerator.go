package main

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Gorutine to generates orders.
func OrderGenerator(ctx context.Context, orders chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 4; i++ {
		orders <- randomOrder()
		time.Sleep(1 * time.Second)
		if i == 2 {

		}
	}
	close(orders)
}

// Generates random order.
func randomOrder() order {
	// Generate products.
	var products []product
	for i := 0; i < rand.Intn(5); i++ {
		products = append(products, product{name: "product " + strconv.Itoa(rand.Intn(100)), price: rand.Intn(100)})
	}
	return order{
		customer: "Customer " + strconv.Itoa(rand.Intn(100)),
		products: products,
	}
}
