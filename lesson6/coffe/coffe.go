package coffee

import "fmt"

type coffeInterface interface {
	Brew()
}

type arabica struct{}

func (a *arabica) Brew() {
	fmt.Println("Arabica is brewing...")
}

type robusta struct{}

func (a *robusta) Brew() {
	fmt.Println("Robusta is brewing...")
}
