package main

import (
	"bufio"
	"fmt"
	"os"
)

var playerX *player
var playerO *player
var score map[string]int

func main() {
	// Game sore and players.
	score = make(map[string]int)
	playerX = &player{value: 10, name: "X", winningSumCols: Edge * 10}
	playerO = &player{value: 100, name: "O", winningSumCols: Edge * 100}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		game := NewGame()
		loop := NewGameLoop(game)
		loop.run()

		// Print score and continue?
		fmt.Println("Scores:")
		for player, score := range score {
			fmt.Printf("%s: %d", player, score)
		}
		fmt.Printf("\n\nContinue? (y)es :")
		scanner.Scan()
		choice := scanner.Text()
		if choice != "y" {
			break
		}
	}
}
