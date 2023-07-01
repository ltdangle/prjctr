package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Game sore.
	score := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	// Outer game loop.
	for {
		game := NewGame(
			&player{value: 10, name: "X", winningSumCols: Edge * 10},
			&player{value: 100, name: "O", winningSumCols: Edge * 100},
			NewGrid(),
		)
		loop := NewGameLoop(game, &score)
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
