package main

import (
	"fmt"
	"hw4/editor"
	"hw4/rating"
)

func main() {
	fmt.Print("Run (e)ditor, student (r)atings, or (q)uit? ")

	switch readInput() {
	case "e":
		e := &editor.Editor{}
		e.Run()
	case "r":
		r := &rating.Rating{}
		r.Run()
	case "q":
	}

}
func readInput() string {
	var choice string
	for {
		fmt.Scanln(&choice)
		if choice == "e" || choice == "r" || choice == "q" {
			return choice
		}
		fmt.Println("You have to type \"e\" or \"r\" or \"q\".")
	}
}
