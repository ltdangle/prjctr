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
	l.choosePlayer1Side()
	// l.clearScreen()
	l.drawGrid()
}
func (l *gameLoop) drawGrid() {
	for _, row := range l.game.grid {
		for _, cell := range row {
			switch cell.value {
			case playerX.value:
				fmt.Print(playerX.name + " |")
			case playerO.value:
				fmt.Print(playerO.name + " |")
			default:
				fmt.Print("  |")
			}
		}
		fmt.Println()
	}
}
func (l *gameLoop) choosePlayer1Side() {
	scanner := bufio.NewScanner(os.Stdin)
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
}

func (l *gameLoop) clearScreen() {
	fmt.Print("\033[H\033[2J")
}
