package main

import "fmt"

func main() {
	// Game loop.
	var choice int
	scenario := scenario()
	for len(scenario.next) > 0 {

		fmt.Println(scenario.description)
		for i, next := range scenario.next {
			fmt.Printf("%d: %s\n", i, next.action)
		}

		fmt.Print("Select next action or type -1 to go back: ")
		fmt.Scan(&choice)

		if choice == -1 {
			if scenario.previous != nil {
				scenario = scenario.previous
				continue
			}
			continue
		}

		scenario = scenario.next[choice]
	}
	fmt.Printf("%s\n", scenario.description)
	fmt.Println("The end!")
}
