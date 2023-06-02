package main

type zookeeper struct {
	name string
}

func (z zookeeper) whoami() string {
	return z.name
}
