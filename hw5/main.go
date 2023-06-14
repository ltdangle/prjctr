package main

import (
	"os"
)

func main() {
	game := NewGame()
	game.SetO(0, 0)
	game.SetX(1, 0)
	game.SetX(1, 1)
	game.SetX(1, 2)
	loop := NewGameLoop(game)
	loop.run()
	os.Exit(0)
}
