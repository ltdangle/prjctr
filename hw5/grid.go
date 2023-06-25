package main

const Edge = 3

// Grid definition.
type cell struct{ value int }
type row [Edge]cell
type grid [Edge]row

// Grid constructor.
func NewGrid() *grid {
	grid := &grid{}
	return grid
}
