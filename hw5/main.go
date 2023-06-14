package main

import "fmt"

func main() {
	grid := NewGrid()
	grid.SetO(0, 0)
	grid.SetX(1, 0)
	fmt.Println(grid)
}
