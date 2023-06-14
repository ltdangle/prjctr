package main

import "fmt"

func main() {
	game := NewGame()
	game.SetO(0, 0)
	game.SetX(1, 0)
	fmt.Println(game)
}
