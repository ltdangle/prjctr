package main

import "fmt"

func main() {
	// Create animals.
	wolf := wolf{}
	fox := fox{}
	elephant := elephant{}
	zebra := zebra{}
	pantera := pantera{}

	// Create cages.
	cage1 := &wolfCage{animal: wolf}
	cage1.setDimensions(3, 3, 3)
	cage1.animalCount = 3

	cage2 := &foxCage{animal: fox}
	cage2.setDimensions(2, 3, 3)
	cage2.animalCount = 2

	cage3 := &elephantCage{animal: elephant}
	cage3.setDimensions(10, 30, 30)
	cage3.animalCount = 1

	cage4 := &zebraCage{animal: zebra}
	cage4.setDimensions(5, 25, 9)
	cage4.animalCount = 2

	cage5 := &panteraCage{animal: pantera}
	cage5.setDimensions(5, 25, 9)
	cage5.animalCount = 3

	// Create zookeper.
	keeper := &zookeeper{name: "Semen"}

	//Create zoo.
	var zoo zoo = zoo{
		wolfCage:     cage1,
		foxCage:      cage2,
		elephantCage: cage3,
		zebraCage:    cage4,
		panteraCage:  cage5,
		zookeeper:    keeper,
	}

	fmt.Println(zoo)
}
