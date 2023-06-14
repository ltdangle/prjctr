package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)

type gameLoop struct {
	game     *game
	player1  *player
	player2  *player
	nextTurn *player
}

func NewGameLoop(g *game) *gameLoop {
	return &gameLoop{game: g}
}

func (l *gameLoop) run() {
	l.choosePlayer1Side()
	l.makeMoves()
}

func (l *gameLoop) makeMoves() {
	scanner := bufio.NewScanner(os.Stdin)
	for l.game.hasEmptyCells() {
		l.drawGrid()
		fmt.Print("Player " + l.nextTurn.name + " choose your cell (row, col): ")
		scanner.Scan()
		fmt.Println("Player " + l.nextTurn.name + " chose " + scanner.Text())
	}
}

func (l *gameLoop) gridHeader() string {
	var header string = "    "
	for i := range l.game.grid[0] {
		header += fmt.Sprintf(" %d  ", i)
	}

	return header
}

func (l *gameLoop) drawGrid() {
	fmt.Println(l.gridHeader())
	for rowIndex, row := range l.game.grid {
		fmt.Printf(" %s |", strconv.Itoa(rowIndex))
		for _, cell := range row {
			switch cell.value {
			case playerX.value:
				fmt.Printf(" %s |", playerX.name)
			case playerO.value:
				fmt.Printf(" %s |", playerO.name)
			default:
				fmt.Printf(" %s |", " ")
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
	l.nextTurn = l.player1
}

func (l *gameLoop) clearScreen() {
	fmt.Print("\033[H\033[2J")
}
