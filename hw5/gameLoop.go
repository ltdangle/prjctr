package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strconv"
)

type gameLoop struct {
	game              *game
	score             *map[string]int
	player1           *player
	player2           *player
	currentTurnPlayer *player
}

func NewGameLoop(g *game, score *map[string]int) *gameLoop {
	return &gameLoop{
		game:  g,
		score: score,
	}
}

func (l *gameLoop) run() {
	l.choosePlayer1Side()
	l.makeMoves()
}

func (l *gameLoop) makeMoves() {
	for l.game.hasEmptyCells() {
		// Clear screen.
		l.clearScreen()

		// Draw game grid.
		l.drawGrid()

		// Exit loop if we have a winner.
		if winner := l.weHaveAWinner(); winner != nil {
			fmt.Println("The winner is player " + winner.name)
			break
		}

		// Get user input.
		l.userInput()

		// Pass turn to the other player.
		if l.currentTurnPlayer == l.player1 {
			l.currentTurnPlayer = l.player2
		} else if l.currentTurnPlayer == l.player2 {
			l.currentTurnPlayer = l.player1
		}
	}

}

func (l *gameLoop) userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Player " + l.currentTurnPlayer.name + " choose your cell (row,col): ")
	for {
		scanner.Scan()

		// Validate coordinates input.
		row, col, isValid := l.validateCoordInput(scanner.Text())
		if !isValid {
			fmt.Println("Error: Incorrect coordinates.")
			continue
		}

		// Check that coordinates are valid (in bounds).
		err := l.game.Set(l.currentTurnPlayer, row, col)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			continue
		}
		break
	}

}
func (l *gameLoop) weHaveAWinner() *player {

	if winner := l.game.WhoWon(); winner != nil {
		// Record the score.
		_, ok := (*l.score)[winner.name]
		if !ok {
			(*l.score)[winner.name] = 1
		} else {
			(*l.score)[winner.name]++
		}
		return winner
	}
	return nil
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
			case l.game.playerX.value:
				fmt.Printf(" %s |", l.game.playerX.name)
			case l.game.playerO.value:
				fmt.Printf(" %s |", l.game.playerO.name)
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
		fmt.Printf("Player1 (%s) or (%s) ? ", l.game.playerX.name, l.game.playerO.name)
		scanner.Scan()
		switch scanner.Text() {
		case l.game.playerX.name:
			l.player1 = l.game.playerX
			l.player2 = l.game.playerO
			break name
		case l.game.playerO.name:
			l.player1 = l.game.playerO
			l.player2 = l.game.playerX
			break name
		default:
			continue
		}
	}
	l.currentTurnPlayer = l.player1
}

func (l *gameLoop) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (l *gameLoop) validateCoordInput(s string) (int, int, bool) {
	r := regexp.MustCompile(`^(\d+),(\d+)$`)
	matches := r.FindStringSubmatch(s)

	if len(matches) == 3 {
		val1, _ := strconv.Atoi(matches[1])
		val2, _ := strconv.Atoi(matches[2])
		return val1, val2, true
	}

	return 0, 0, false
}
