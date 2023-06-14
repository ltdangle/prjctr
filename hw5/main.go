package main

import (
	"os"
)

func main() {
	game := NewGame()
	game.Set(playerX, 1, 0)
	game.Set(playerO, 1, 1)
	game.Set(playerX, 1, 2)
	loop := NewGameLoop(game)
	loop.run()
	os.Exit(0)
}
