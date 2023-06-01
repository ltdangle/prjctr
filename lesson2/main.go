package main

import "fmt"

var (
	red   = "red"
	blue  = "blue"
	white = "white"
)

type Room struct {
	Name     string
	Area     float64
	color    string
	nextRoom *Room
}

func main() {
	bathroom := Room{
		Name:  "Bathroom",
		Area:  15.3,
		color: red,
	}
	bedroom := Room{
		Name:     "Bedroom",
		color:    blue,
		nextRoom: &bathroom,
	}

	fmt.Println(bathroom, bedroom)

	printWhite()
}

func printWhite() {
	fmt.Println(white)
}
