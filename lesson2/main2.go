package main

import (
	"encoding/json"
	"fmt"
)

type Ptashka struct {
	Name string `json:"name"`
}

func (p Ptashka) Sing() string {
	return "Phewt"
}

type Gnizdo struct {
	Ptashka *Ptashka
}
type Tree struct {
	Name   string
	Gnizdo *Gnizdo
}

func main() {
	p := Ptashka{Name: "ptashka"}
	g := Gnizdo{
		Ptashka: &p,
	}
	t := Tree{Name: "oak", Gnizdo: &g}

	fmt.Println("Bird sings:------------------------")
	fmt.Println(p.Sing())

	msg, _ := json.Marshal(t)
	fmt.Println("Json tree:-----------------------------")
	fmt.Println(string(msg))

	msg2, _ := json.Marshal(Tree{})
	fmt.Println("Json empty tree:-----------------------------")
	fmt.Println(string(msg2))
}
