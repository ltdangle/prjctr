package main

import (
	"os"
)

func main() {
	playerX = &player{value: 10, name: "X", winningSumCols: ROWS * 10}
	playerO = &player{value: 100, name: "O", winningSumCols: ROWS * 100}
	game := NewGame()
	game.Set(playerX, 1, 0)
	game.Set(playerO, 1, 1)
	game.Set(playerX, 1, 2)
	loop := NewGameLoop(game)
	loop.run()
	os.Exit(0)
}
