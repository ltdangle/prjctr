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
	ordersCh:=make(chan order)

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

		orders:=5
		counter:=1

		ticker:
		for {
			select {
			case <-ticker.C:
				if counter==orders{
					break ticker
				}
				id := rand.Intn(100)
				ordersCh <- order{id: id, name: "Order_" + strconv.Itoa(id)}
				counter++
			}
		}
		close(ordersCh)
	}()

	wg.Wait()
	fmt.Println("Program exited.")
}
