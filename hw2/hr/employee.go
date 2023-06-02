package hr

// Person.
type Person struct {
	Name string
	Age  int
}

func (p Person) WhoAmI() string {
	return "I am a person!"
}

// Employee.
type Employee struct {
	Person
	Postion string
}

func (e Employee) WhoAmI() string {
	return "I am an employee!"
}
