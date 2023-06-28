package main

import (
	"context"
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

	order := randomOrder()

	wg.Add(1)
	go ProductFinder(ctx, order, &wg)
	wg.Add(1)
	go OrderTotalCalculator(ctx, order, &wg)

	wg.Wait()

	// jsonOrderQue, _ := json.MarshalIndent(order, "", " ")
	// fmt.Println(string(jsonOrderQue))
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