package main

import (
	"fmt"
	"hw4/editor"
	"hw4/rating"
	"os"
)

func main() {
	var choice string
	for {
		fmt.Print("Run (e)ditor, student (r)atings, or (q)uit? ")

		fmt.Scanln(&choice)
		switch choice {
		case "e":
			e := &editor.Editor{}
			e.Run()
			os.Exit(0)
		case "r":
			r := &rating.Rating{}
			r.Run()
			os.Exit(0)
		case "q":
			os.Exit(0)
		default:
			fmt.Println("You have to type \"e\" or \"r\" or \"q\".")
		}
	}
}
