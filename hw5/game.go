package main

import (
	"errors"
	"fmt"
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
	// Calculate win over rows.
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

	// Calculate win over columns.
	var colSum int
	rows := len(g.grid)
	cols := len(g.grid[0])
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			colSum += g.grid[r][c].value
		}
		if colSum == playerX.winningSumCols {
			return playerX
		}
		if colSum == playerO.winningSumCols {
			return playerO
		}
		colSum = 0
	}

	// Calculate win over first horizontal (left - right).
	var diag1Sum int
	rows = len(g.grid)
	cols = len(g.grid[0])
	for c := 0; c < cols; c++ {
		diag1Sum += g.grid[c][c].value
	}
	if diag1Sum == playerX.winningSumCols {
		return playerX
	}
	if diag1Sum == playerO.winningSumCols {
		return playerO
	}

	// Calculate win over second horizontal (right - left).
	var diag2Sum int
	rows = len(g.grid)
	cols = len(g.grid[0])
	c := 0
	for r := cols - 1; r >= 0; r-- {
		diag2Sum += g.grid[r][c].value
		c++
	}
	if diag2Sum == playerX.winningSumCols {
		return playerX
	}
	if diag2Sum == playerO.winningSumCols {
		return playerO
	}

	return nil
}
