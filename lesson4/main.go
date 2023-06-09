package main

import (
	"fmt"
	"math"
)

func main() {
	// Опишіть слайсами потяг, який проїжджає три станції, і на кожній до нього приєднується ще три вагона
	train := []string{"cart", "cart", "cart"}
	for i := 0; i < 3; i++ {
		train = append(train, "cart", "cart", "cart")
	}
	fmt.Println(train)

	// Опишіть слайсами вечірку, на яку всі прийшли парами. Усього 4 пари, але поступово вони пішли з вечірки (використовуй масив, ``як елемент слайса)
	party := [][2]string{
		{"john", "mary"},
		{"michael", "kate"},
		{"phillip", "joselyn"},
		{"igor", "svetlana"},
	}

	for i := range party {
		party[i] = [2]string{"", ""}
		fmt.Println("party: ", party)
	}
	// Опішіть слайсами зростання дерева. Спочатку у нього одна гілка, потім три, потім 6, потім 12, потім 24, і на прикінці - 48
	tree := []branch{}
	for i := 1; i <= 4; i++ {
		newBranchCount =math.Pow(x float64, y float64) 
		tree = append(tree)
	}
}
func createBranches(num int) []branch {
	var branches []branch
	for i := 0; i < num; i++ {
		branches = append(branches, branch{})
	}
	return branches
}

type branch struct{}
