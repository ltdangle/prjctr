package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type product struct {
	name  string
	price int
}
type order struct {
	customer string
	products []product
}

func main() {
	orders := make(chan order)
	var wg sync.WaitGroup
	ctx, _ := context.WithCancel(context.Background())

	go OrderTotalCalculator(ctx, orders, &wg)

	for i := 0; i <= 4; i++ {
		orders <- randomOrder()
		time.Sleep(1 * time.Second)
	}
	close(orders)

	wg.Wait()
	fmt.Println()
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
