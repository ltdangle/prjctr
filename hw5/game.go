package main

import (
	"errors"
)

type game struct{ grid *grid }

// Game constructor.
func NewGame() *game {
	return &game{grid: NewGrid()}
}

// Game methods.
func (g *game) Set(player *player, row int, col int) error {
	if err := g.validateCoords(row, col); err != nil {
		return err
	}

	if !g.hasEmptyCells() {
		return errors.New("No more empty cells. The game is finished.")
	}

	// Check if the cell is taken
	if g.grid[row][col].value != 0 {
		return errors.New("Cell is already taken.")
	}

	g.grid[row][col].value = player.value
	return nil
}

func (g *game) validateCoords(row int, col int) error {
	// Coordinates cannot be negative.
	if row < 0 || col < 0 {
		return errors.New("Coordinates cannot be negative.")
	}

	// Coordinateds must not be out of bounds.
	if row > ROWS-1 || col > COLS-1 {
		return errors.New("Coordinates cannot be out of bounds.")
	}

	return nil
}

func (g *game) hasEmptyCells() bool {
	for _, row := range g.grid {
		for _, cell := range row {
			if cell.value == 0 {
				return true
			}
		}
	}
	return false
}

func (g *game) WhoWon() *player {
	// Calculate rows.
	var rowSum int
	for _, row := range g.grid {
		for _, cell := range row {
			rowSum += cell.value
		}
		if rowSum == playerX.winningSumRows {
			return playerX
		}
		if rowSum == playerO.winningSumRows {
			return playerO
		}
		rowSum = 0
	}

	return nil
}
