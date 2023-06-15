package main

const EDGE = 3

// Grid definition.
type cell struct{ value int }
type row [EDGE]cell
type grid [EDGE]row

// Grid constructor.
func NewGrid() *grid {
	grid := &grid{}
	return grid
}
