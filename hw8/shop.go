package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Product struct {
	name  string
	price int
}
type Order struct {
	Customer string `json:"Customer"`
	products []Product
	Total    int `json:"Total"`
}

func main() {
	// Timeout for all jobs.
	jobsTimeout := 4

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(jobsTimeout)*time.Second)
	defer cancel()

	// Channel for errors from the worker goroutines.
	errc := make(chan error)

	order := randomOrder()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- ProductFinder(ctx, order)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errc <- OrderTotalCalculator(ctx, order)
	}()

	// Wait for any of the goroutines to return an error. If one does,
	// cancel the context to stop the others.
	select {
	case err := <-errc:
		if err != nil {
			fmt.Printf("\nOrder processing canceled: %s \n", err)
			cancel()
			return
		}
	case <-ctx.Done():
		fmt.Printf("\nOrder processing timed out.\n")
		return
	}

	wg.Wait()
}

// Generates random order.
func randomOrder() *Order {
	// Generate products.
	var products []Product
	for i := 0; i < rand.Intn(5); i++ {
		products = append(products, Product{name: "product " + strconv.Itoa(rand.Intn(100)), price: rand.Intn(100)})
	}
	return &Order{
		Customer: "Customer " + strconv.Itoa(rand.Intn(100)),
		products: products,
	}
}
