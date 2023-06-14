package main

import (
	"bufio"
	"fmt"
	"os"
)

type gameLoop struct {
	game    *game
	player1 *player
	player2 *player
}

func NewGameLoop(g *game) *gameLoop {
	return &gameLoop{game: g}
}

func (l *gameLoop) run() {
	scanner := bufio.NewScanner(os.Stdin)

	// Choose player1 side.
name:
	for {
		l.clearScreen()
		fmt.Printf("Playing for (%s) or (%s) ? ", playerX.name, playerO.name)
		scanner.Scan()
		switch scanner.Text() {
		case playerX.name:
			l.player1 = playerX
			l.player2 = playerO
			break name
		case playerO.name:
			l.player1 = playerO
			l.player2 = playerX
			break name
		default:
			continue
		}
	}
	fmt.Println("You are playing for " + l.player1.name)
}

func (l *gameLoop) clearScreen() {
	fmt.Print("\033[H\033[2J")
}
