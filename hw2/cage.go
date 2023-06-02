package main

// Cage.
type cage struct {
	height      int
	length      int
	width       int
	animalCount int
}

func (c cage) setDimensions(height int, length int, width int) {
	c.height = height
	c.length = length
	c.width = width
}

// Wolf cage.
type wolfCage struct {
	cage
	animal wolf
}

// Fox cage.
type foxCage struct {
	cage
	animal fox
}

// Elephant cage.
type elephantCage struct {
	cage
	animal elephant
}

// Zebra cage.
type zebraCage struct {
	cage
	animal zebra
}

// Pantera cage.
type panteraCage struct {
	cage
	animal pantera
}
