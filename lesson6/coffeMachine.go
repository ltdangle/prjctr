package main

import "fmt"

type CoffeMachine interface {
	Prepare(coffe coffeInterface)
}

type HomeCoffeMachine struct{}

func (m *HomeCoffeMachine) Prepare(coffe coffeInterface) {
	fmt.Println("HomeCoffeMachine is preparing...")
	coffe.Brew()
}

type CafeCoffeMachine struct{}

func (m *CafeCoffeMachine) Prepare(coffe coffeInterface) {
	fmt.Println("CafeHomeCoffeMachine is preparing...")
	coffe.Brew()
}
