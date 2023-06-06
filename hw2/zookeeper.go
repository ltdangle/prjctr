package main

type zookeeper struct {
	Name string
}

func (z *zookeeper) addAnimalToCage(cage *cage, animal animal) {
	cage.Animal = animal
}

func (z zookeeper) whoami() string {
	return z.Name
}
