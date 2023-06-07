package main

import "fmt"

func main() {
	// Game loop.
	var choice int
	scenario := scenario()
	for scenario.hasNextScene() {
		// Print description.
		fmt.Println()
		fmt.Printf("%s\n\n", scenario.description)

		// Print next available actions.
		for i, next := range scenario.next {
			fmt.Printf("%d: %s\n", i, next.action)
		}

		// Get user input.
		fmt.Print("\nSelect next action or type -1 to go back: ")
		fmt.Scan(&choice)
		fmt.Print("\n")
		printDoubleLine()

		// Go back.
		if choice == -1 {
			if scenario.hasPreviousScene() {
				scenario = scenario.previous
				continue
			}
			continue
		}

		// Validate input.
		if choice < 0 || choice > scenario.nextSceneCount() {
			continue
		}

		scenario = scenario.next[choice]
	}

	fmt.Printf("%s\n", scenario.description)
	fmt.Println("The end!")

}
func printDoubleLine() {
	fmt.Println("======================================")
}
