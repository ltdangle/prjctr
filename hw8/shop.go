package main

import (
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
	ordersCh := make(chan order)

	wg.Add(1)
	// Calculates order total
	go func() {
		defer wg.Done()
		for ordr := range ordersCh {
			fmt.Printf("\nReceived order %v", ordr)
		}
	}()

	wg.Add(1)
	// Calculates order total
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)

		orders := 5
		counter := 1

	ticker:
		for {
			select {
			case <-ticker.C:
				if counter == orders {
					break ticker
				}
				ordersCh <-createOrder()
				counter++
			}
		}
		close(ordersCh)
	}()

	wg.Wait()
	fmt.Printf("\nProgram exited.")
}
func createOrder() order {
	var products []product

	for i := 0; i < rand.Intn(5); i++ {
		products = append(products, createProduct())
	}

	id := rand.Intn(100)
	return order{id: id, name: "Product_" + strconv.Itoa(id), products: products}
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
