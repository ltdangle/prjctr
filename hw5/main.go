package main

import (
	"os"
)

var playerX *player
var playerO *player

func main() {
	playerX = &player{value: 10, name: "X", winningSumCols: EDGE * 10}
	playerO = &player{value: 100, name: "O", winningSumCols: EDGE * 100}
	game := NewGame()
	loop := NewGameLoop(game)
	loop.run()
	os.Exit(0)
}
