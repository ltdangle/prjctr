package main

// TransportI interface.
type TransportI interface {
	Name() string
	Move()
	Stop()
	ChangeSpeed()
	TakePassengers()
	DropPassengers()
}

// Car implementation.
type Car struct{}

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

// Train implementation.
type Train struct{}

func (t *Train) Name() string {
	return "Train"
}

func (t *Train) Move() {
	println("Train is moving")
}

func (t *Train) Stop() {
	println("Train stopped")
}

func (t *Train) ChangeSpeed() {
	println("Train changed speed")
}

func (t *Train) TakePassengers() {
	println("Train took passengers")
}

func (t *Train) DropPassengers() {
	println("Train dropped passengers")
}

// Plane implementation.
type Plane struct{}

func (p *Plane) Name() string {
	return "Plane"
}

func (p *Plane) Move() {
	println("Plane is moving")
}

func (p *Plane) Stop() {
	println("Plane stopped")
}

func (p *Plane) ChangeSpeed() {
	println("Plane changed speed")
}

func (p *Plane) TakePassengers() {
	println("Plane took passengers")
}

func (p *Plane) DropPassengers() {
	println("Plane dropped passengers")
}

// Route.
type Route struct {
	transport []TransportI
}

func (r *Route) AddTransport(t TransportI) {
	r.transport = append(r.transport, t)
}
func (r *Route) ShowTransport() string {
	var str string
	for _, t := range r.transport {
		str += t.Name() + "\n"
	}
	return str
}
