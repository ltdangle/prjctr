package main

import "errors"

type game struct{ grid *grid }

// Game constructor.
func NewGame() *game {
	return &game{grid: NewGrid()}
}

// Grid methods.
func (g *game) SetX(row int, col int) error {
	if err := g.validateCoords(row, col); err != nil {
		return err
	}
	g.grid[row][col].value = CROSS
	return nil
}

func (g *game) SetO(row int, col int) error {
	if err := g.validateCoords(row, col); err != nil {
		return err
	}
	g.grid[row][col].value = ZERO
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
