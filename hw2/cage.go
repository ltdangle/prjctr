package main

// Cage.
type cage struct {
	Height      int `json:"height"`
	Length      int `json:"length"`
	Width       int `json:"width"`
	AnimalCount int `json:"animalCount"`
}

func (c *cage) setDimensions(height int, length int, width int) {
	c.Height = height
	c.Length = length
	c.Width = width
}

// Wolf cage.
type wolfCage struct {
	cage
	Animal wolf
}

// Fox cage.
type foxCage struct {
	cage
	Animal fox
}

// Elephant cage.
type elephantCage struct {
	cage
	Animal elephant
}

// Zebra cage.
type zebraCage struct {
	cage
	Animal zebra
}

// Pantera cage.
type panteraCage struct {
	cage
	Animal pantera
}
