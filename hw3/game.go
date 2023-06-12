package main

import "fmt"

// Game struct.
type game struct {
}

// Game loop.
func (g *game) loop(scene *scene) {

	// Game loop.
	var choice int
	for scene.hasNextScene() {
		fmt.Print(scene)

		// Get user input.
		fmt.Print("\nSelect next action or type -1 to go back: ")
		fmt.Scan(&choice)
		fmt.Print("\n")
		fmt.Println("===========================================")

		// Go back.
		if choice == -1 {
			if scene.hasPreviousScene() {
				scene = scene.previous
				continue
			}
			continue
		}

		// Validate input.
		if choice < 0 || choice > scene.nextSceneCount() {
			continue
		}

		// Goto next scene.
		scene = scene.gotoNextScene(choice)
	}

	// The end.
	fmt.Print(scene)
	fmt.Println("The end!")
}
