package main

import (
	"context"
	"sync"
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
	ctx, _ := context.WithCancel(context.Background())
	wg.Add(1)
	go OrderGenerator(ctx, orders, &wg)
	wg.Add(1)
	go OrderTotalCalculator(ctx, orders, &wg)

	wg.Wait()
}
