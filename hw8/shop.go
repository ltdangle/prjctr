package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	CustomerName string
	Item         string
	Quantity     int
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	orders := make(chan Order)

	// Генеруємо випадкові замовлення
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(orders)
				return
			default:
				order := Order{
					CustomerName: fmt.Sprintf("Customer%d", rand.Intn(100)),
					Item:         "SomeItem",
					Quantity:     rand.Intn(5) + 1, // Ensure quantity is always at least 1
				}
				orders <- order
				time.Sleep(1 * time.Second) // Let's assume a new order comes in every second
			}
		}
	}()

	// Обробляємо замовлення
	go func() {
		defer wg.Done()
		const pricePerItem = 10 // Let's assume each item costs 10
		for order := range orders {
			orderTotal := order.Quantity * pricePerItem
			fmt.Printf("Processed order from %s: %d %s(s) totaling $%d\n", order.CustomerName, order.Quantity, order.Item, orderTotal)
		}
	}()

	// Зупиняємо генерування замовлень після 10 секунд
	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	wg.Wait()
}
