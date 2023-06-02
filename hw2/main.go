package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create animals.
	wolf := wolf{}
	fox := fox{}
	elephant := elephant{}
	zebra := zebra{}
	pantera := pantera{}

	// Create cages.
	cage1 := &wolfCage{Animal: wolf}
	cage1.setDimensions(3, 3, 3)
	cage1.AnimalCount = 3

	cage2 := &foxCage{Animal: fox}
	cage2.setDimensions(2, 3, 3)
	cage2.AnimalCount = 2

	cage3 := &elephantCage{Animal: elephant}
	cage3.setDimensions(10, 30, 30)
	cage3.AnimalCount = 1

	cage4 := &zebraCage{Animal: zebra}
	cage4.setDimensions(5, 25, 9)
	cage4.AnimalCount = 2

	cage5 := &panteraCage{Animal: pantera}
	cage5.setDimensions(5, 25, 9)
	cage5.AnimalCount = 3

	// Create zookeper.
	keeper := &zookeeper{name: "Semen"}

	//Create zoo.
	var zoo zoo = zoo{
		Name:         "My Zoo",
		WolfCage:     cage1,
		FoxCage:      cage2,
		ElephantCage: cage3,
		ZebraCage:    cage4,
		PanteraCage:  cage5,
		Zookeeper:    keeper,
	}

	//Print zoo json.
	zooJson, _ := json.MarshalIndent(zoo, "", " ")
	fmt.Println(string(zooJson))
}
