package main

import "fmt"

func main() {
	// Game loop.
	var choice int
	scenario := scenario()
	for scenario.hasNextScene() {
		fmt.Println()
		fmt.Printf("%s\n\n", scenario.description)
		for i, next := range scenario.next {
			fmt.Printf("%d: %s\n", i, next.action)
		}

		fmt.Print("\nSelect next action or type -1 to go back: ")
		fmt.Scan(&choice)
		fmt.Print("\n")
		doubleLine()

		// Go back.
		if choice == -1 {
			if scenario.previous != nil {
				scenario = scenario.previous
				continue
			}
			continue
		}

		// Validate input.
		if choice < 0 || choice > scenario.countNextScene() {
			continue
		}

		scenario = scenario.next[choice]
	}

	fmt.Printf("%s\n", scenario.description)
	fmt.Println("The end!")

}
func doubleLine() {
	fmt.Println("======================================")
}
