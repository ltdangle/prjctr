package main

type player struct {
	value int
	name  string
}

var playerX = &player{value: 10, name: "X"}
var playerO = &player{value: 100, name: "O"}
