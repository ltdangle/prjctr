package main

import (
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

	wg.Add(1)
	go OrderGenerator(orders, &wg)
	wg.Add(1)
	go OrderTotalCalculator(orders, &wg)

	wg.Wait()
}

// Generates orders.
func OrderGenerator(orders chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 4; i++ {
		orders <- randomOrder()
		time.Sleep(1 * time.Second)
	}
	close(orders)
}

// Calculates order subtotal.
func OrderTotalCalculator(orders chan order, wg *sync.WaitGroup) {
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

// Calculates order total.
func calculateOrderTotal(order order) int {
	var total int
	for _, product := range order.products {
		total += product.price
	}
	return total
}
