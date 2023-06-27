package main

import (
	"context"
	"encoding/json"
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
	orders := make(chan *Order)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go OrderTotalCalculator(ctx, orders, &wg)

	var orderQue []*Order
	for i := 0; i <= 4; i++ {
		rndOrder := randomOrder()
		orderQue = append(orderQue, rndOrder)

		orders <- rndOrder
		time.Sleep(1 * time.Second)
		if i == 2 {
			cancel()
			break
		}
	}
	wg.Wait()

	close(orders)

	// Print order que.
	jsonOrderQue, _ := json.MarshalIndent(orderQue, "", " ")
	fmt.Println(string(jsonOrderQue))
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
