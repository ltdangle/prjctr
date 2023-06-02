package main

type zookeeper struct {
	Name string
}

func (z zookeeper) whoami() string {
	return z.Name
}
