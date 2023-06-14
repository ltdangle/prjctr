package main

type player struct {
	value          int
	name           string
	winningSumRows int
	winningSumCols int
}

var playerX = &player{value: 10, name: "X", winningSumRows: ROWS * 10, winningSumCols: COLS * 10}
var playerO = &player{value: 100, name: "O", winningSumRows: ROWS * 100, winningSumCols: COLS * 100}
