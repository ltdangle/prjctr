package main

import "fmt"

func main() {
	// Game loop.
	var choice int
	scene := scenario()
	for scene.hasNextScene() {
		// Print description.
		fmt.Println()
		fmt.Println("You are here: " + breadcrumbs(scene))
		fmt.Println()
		fmt.Printf("%s\n\n", scene.description)

		// Print next available actions.
		for i, next := range scene.next {
			fmt.Printf("%d: %s\n", i, next.action)
		}

		// Get user input.
		fmt.Print("\nSelect next action or type -1 to go back: ")
		fmt.Scan(&choice)
		fmt.Print("\n")
		fmt.Println(doubleLine())

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

		scene = scene.gotoNextScene(choice)
	}

	// The end.
	fmt.Printf("%s\n", scene.description)
	fmt.Println("The end!")
}

func breadcrumbs(s *scene) string {
	var breadcrumbs string
	for s.hasPreviousScene() {
		breadcrumbs = s.name + " > " + breadcrumbs
		s = s.gotoPreviousScene()
	}
	breadcrumbs = s.name + " > " + breadcrumbs
	return breadcrumbs
}

func doubleLine() string {
	return "======================================"
}
