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

func (s *scene) hasNextScene() bool {
	if len(s.next) > 0 {
		return true
	}
	return false
}
func (s *scene) countNextScene() int {
	if len(s.next) == 0 {
		return 0
	}
	return len(s.next) - 1
}
func scenario() *scene {
	cave := &scene{
		name:        "Cave entrance",
		action:      "Go to the front of the cave",
		description: "You woke up near the cave entrance. You see a backpack nearby. There are matches, lighter, and a knife in the backpack. To the left you see a trail that leads into the forest. To the right you see a trail that disapears into meadow.",
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
	meadow := &scene{
		name:        "Meadow",
		action:      "Go to the meadow",
		description: "You see a well in the middle of the meadow. You notice that you haven't had water for a while and you are quite thirsty. You also see a showel to your left and marks of freshly dug earth.",
	}
	well := &scene{
		name:        "Well",
		action:      "Drink from the well",
		description: "You decide to quench your thirst by drinking from the well. Instead of satisfaction, you feel nauseaus and you pass out. The well was poisoned. ",
	}
	hole := &scene{
		name:        "Hole in the ground",
		action:      "Dig a hole in the ground where marks are",
		description: "You dig a whole in the ground using the showel you find. There is something at the bottom! It is a chest full of treasure!",
	}

	cave.addNextScene(forest)
	cave.addNextScene(meadow)
	meadow.addNextScene(well)
	meadow.addNextScene(hole)
	forest.addNextScene(camp)
	camp.addNextScene(tent)
	tent.addNextScene(safe)

	return cave
}
