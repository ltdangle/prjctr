package transport

type Car struct{}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) Name() string {
	return "Car"
}

func (c *Car) Move() {
	println("Car is moving")
}

func (c *Car) Stop() {
	println("Car stopped")
}

func (c *Car) ChangeSpeed() {
	println("Car changed speed")
}

func (c *Car) TakePassengers() {
	println("Car took passengers")
}

func (c *Car) DropPassengers() {
	println("Car dropped passengers")
}
