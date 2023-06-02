package main

// Animal.
type animal struct {
	name string
}

func (a animal) whoami() string {
	return "I am an animal."
}
func (a animal) setName(name string) {
	a.name = name
}
func (a animal) getName() string {
	return a.name
}

// Wolf.
type wolf struct {
	animal
}

func (w wolf) whoami() string {
	return "I am a wolf."
}

// Fox.
type fox struct {
	animal
}

func (f fox) whoami() string {
	return "I am a fox."
}

// Elephant.
type elephant struct {
	animal
}

func (a elephant) whoami() string {
	return "I am an elephant."
}

// Zebra.
type zebra struct {
	animal
}

func (z zebra) whoami() string {
	return "I am a zebra."
}

// Pantera.
type pantera struct {
	animal
}

func (p pantera) whoami() string {
	return "I am a pantera."
}
