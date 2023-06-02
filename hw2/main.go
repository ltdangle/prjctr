package main

import (
	"fmt"
	"hr/hr"
)

func main() {
	e := hr.Employee{
		Person:  hr.Person{Name: "John Doe", Age: 30},
		Postion: "Software engineer",
	}
	fmt.Println(e.WhoAmI())
	fmt.Println(e.Person.WhoAmI())
}
