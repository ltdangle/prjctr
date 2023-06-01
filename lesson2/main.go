package main

import (
	"encoding/json"
	"fmt"
)

var (
	red   = "red"
	blue  = "blue"
	white = "white"
)

type Room struct {
	Name     string `json:"room_name"`
	Area     float64
	color    string
	NextRoom *Room
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
		NextRoom: &bathroom,
	}
	fmt.Println(bathroom, bedroom)
	printWhite()

	msg, _ := json.Marshal(bedroom)
	fmt.Println("Json:-----------------------------")
	fmt.Println(string(msg))
}

func printWhite() {
	fmt.Println(white)
}
