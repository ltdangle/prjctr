package main

import (
	"bufio"
	"fmt"
	"os"
)

type gameLoop struct {
	game *game
}

func NewGameLoop(g *game) *gameLoop {
	return &gameLoop{game: g}
}

func (l *gameLoop) run() {
	scanner := bufio.NewScanner(os.Stdin)
	l.clearScreen()
	fmt.Println("Playing for (x) or (o) ?")
	scanner.Scan()
	fmt.Println("You are playing for " + scanner.Text())
}

func (l *gameLoop) clearScreen() {
	fmt.Print("\033[H\033[2J")
}
