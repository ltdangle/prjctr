package transport

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
