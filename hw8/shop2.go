package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type order struct {
	product string
	price   int
}

func main() {
	orders := make(chan order)
	var wg sync.WaitGroup

	wg.Add(1)
	func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			orders <- order{product: fmt.Sprintf("product %d", rand.Intn(100))}
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	func() {
		defer wg.Done()
		for ordr := range orders {
			fmt.Printf("\nNew order: %s", ordr.product)
		}
	}()

	wg.Wait()
}
