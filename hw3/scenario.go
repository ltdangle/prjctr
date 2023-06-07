package main

type scene struct {
	// Scene name.
	name string
	// Scene action.
	action string
	// Scene description.
	description string
	// Previous scene.
	previous *scene
	// Next scene.
	next []*scene
}

func (s *scene) addNextScene(scene *scene) {
	scene.previous = s
	s.next = append(s.next, scene)
}

func scenario() *scene {
	cave := &scene{
		name:        "Cave entrance",
		action:      "Go to the front of the cave",
		description: "You woke up near the cave entrance. You see a backpack nearby. There are matches, lighter, and a knife in the backpack. To the left you see a trail that leads into the forest. To the right you see a trail that disapears into horizon.",
	}
	forest := &scene{
		name:        "Forest",
		action:      "Go into the forest",
		description: "In the forest you find a dead body of of unseen-before animal.",
	}
	camp := &scene{
		name:        "Abandoned camp",
		action:      "Go deeper into the forest",
		description: "Deep in the forest you find abandoned camp. The walk has been long, you are tired and decide to rest, rathere than continue.",
	}
	tent := &scene{
		name:        "Camp tent",
		action:      "Enter one of the tents in the camp",
		description: "You enter the nearest tent in the abandoned camp. To your suprise, you find a safe with a two-digit combination lock.",
	}
	safe := &scene{
		name:        "Safe in a tent",
		action:      "Pick a safe",
		description: "You try different letter combinations. Finally, the safe opens. A large insect crawls out of the safe, bites you, and runs away. You loose consciousness. It could've ended differently, very differently indeed.",
	}

	cave.addNextScene(forest)
	forest.addNextScene(camp)
	camp.addNextScene(tent)
	tent.addNextScene(safe)

	return cave
}
