package main

// Створити інтерфейс кавоварки і його імплементацію
// Навчити кавоварку варити на арабіці і на бленді (інтерфейс кави і дви стратегії)
func main() {
	machine := &HomeCoffeMachine{}
	machine.Prepare(&arabica{})
	machine.Prepare(&robusta{})
}
