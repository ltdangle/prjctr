package main

import "fmt"

type CustomConstraint interface {
	~int | ~string
}

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(getKeys(m))
}

func getKeys[K CustomConstraint, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Node[T any] struct {
	val  T
	next *Node[T]
}
