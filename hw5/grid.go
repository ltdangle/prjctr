package main

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
