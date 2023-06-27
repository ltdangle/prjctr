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
	go func() {
		defer wg.Done()
		for i := 0; i <= 4; i++ {
			orders <- order{product: fmt.Sprintf("product %d", rand.Intn(100))}
			time.Sleep(1 * time.Second)
		}
		close(orders)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 1
		for ordr := range orders {
			fmt.Printf("\nNew order %d: %s", counter, ordr.product)
			counter++
		}
		fmt.Println()
	}()

	wg.Wait()
}
