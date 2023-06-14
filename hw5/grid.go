package main

const COLS = 3
const ROWS = 3

// Grid definition.
type cell struct{ value int }
type row [COLS]cell
type grid [ROWS]row

// Grid constructor.
func NewGrid() *grid {
	grid := &grid{}
	return grid
}
