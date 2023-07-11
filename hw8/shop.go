package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type product struct {
	id    int
	name  string
	price int
}
type order struct {
	id       int
	name     string
	products []product
	total    int
}

func main() {
	var wg sync.WaitGroup
	ordersCh := make(chan *order)

	orderCount := flag.Int("orders", 1, "Number of orders to generate")
	flag.Parse()

	wg.Add(1)
	go processOrdersRtn(&wg, ordersCh)

	wg.Add(1)
	go createOrdersRtn(&wg, ordersCh, *orderCount)

	wg.Wait()
	fmt.Printf("\nProgram exited.")
}

func processOrdersRtn(wg *sync.WaitGroup, ordersCh chan *order) {
	defer wg.Done()
	for ordr := range ordersCh {
		calculateTotals(ordr)
		fmt.Printf("\nReceived order %v", ordr)
	}
}

func createOrdersRtn(wg *sync.WaitGroup, ordersCh chan *order, numOrders int) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)

	for i := 0; i < numOrders; i++ {
		select {
		case <-ticker.C:
			ordersCh <- createOrder()
		}
	}
	close(ordersCh)
}

func createOrder() *order {
	var products []product

	for i := 0; i < rand.Intn(5); i++ {
		products = append(products, createProduct())
	}

	id := rand.Intn(100)
	return &order{id: id, name: "Product_" + strconv.Itoa(id), products: products}
}

func createProduct() product {
	id := rand.Intn(100)
	return product{id: id, name: "Product_" + strconv.Itoa(id), price: rand.Intn(100)}
}

func calculateTotals(order *order) {
	for _, product := range order.products {
		order.total += product.price
	}
}
