package main

import "fmt"

const NUMBER = 100

func main() {
	diagnosis()
	seasonClothes()
	guessNumber()
}
func diagnosis() {
	var symptom string
	fmt.Println("Enter symptom: ")
	fmt.Scan(&symptom)

	if symptom == "fever" {
		fmt.Println("You have cold!")
	} else if symptom == "leaking nose" {
		fmt.Println("You have allergy!")
	} else {
		fmt.Println("No such symptom. You are healthy.")
	}
}

func seasonClothes() {
	var season string
	fmt.Println("Enter season: ")
	fmt.Scan(&season)

	switch season {
	case "summer":
		fmt.Println("Wear shorts.")
	case "winter":
		fmt.Println("Wear fur coat.")
	case "spring":
		fmt.Println("Wear light coat.")
	case "fall":
		fmt.Println("Wear raincoat.")
	default:
		fmt.Println("Wrong season!")
	}
}

func guessNumber() {
	var number int
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&number)
		if number < 0 {
			break
		}
		if number > NUMBER {
			fmt.Printf("%d is greater than %d", number, NUMBER)
			fmt.Println()
		} else if number < NUMBER {
			fmt.Printf("%d is smaller than %d", number, NUMBER)
			fmt.Println()
		}
	}
}
