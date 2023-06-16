package coffemachine

import "fmt"

type CoffeMachine interface {
	Prepare(coffe  lesson6.coffee.coffeInterface)
}

type HomeCoffeMachine struct{}

func (m *HomeCoffeMachine) Prepare(coffe coffe.coffeInterface) {
	fmt.Println("HomeCoffeMachine is preparing...")
	coffe.Brew()
}

type CafeCoffeMachine struct{}

func (m *CafeCoffeMachine) Prepare(coffe coffee.coffeInterface) {
	fmt.Println("CafeHomeCoffeMachine is preparing...")
	coffe.Brew()
}
