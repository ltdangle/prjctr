package main

import "errors"

const COLS = 3
const ROWS = 3

const CROSS string = "X"
const ZERO = "O"

// Grid definition.
type cell struct{ value string }
type row [COLS]cell
type grid [ROWS]row

// Grid constructor.
func NewGrid() *grid {
	grid := &grid{}
	return grid
}

// Grid methods.
func (grid *grid) SetX(row int, col int) error {
	if err := grid.validateCoords(row, col); err != nil {
		return err
	}
	grid[row][col].value = CROSS
	return nil
}
func (grid *grid) SetO(row int, col int) error {
	if err := grid.validateCoords(row, col); err != nil {
		return err
	}
	grid[row][col].value = ZERO
	return nil
}

func (grid *grid) validateCoords(row int, col int) error {
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
