package transport

type plane struct{}

func (p *plane) Name() string {
	return "Plane"
}

func (p *plane) Move() {
	println("Plane is moving")
}

func (p *plane) Stop() {
	println("Plane stopped")
}

func (p *plane) ChangeSpeed() {
	println("Plane changed speed")
}

func (p *plane) TakePassengers() {
	println("Plane took passengers")
}

func (p *plane) DropPassengers() {
	println("Plane dropped passengers")
}
func NewPlane() *plane {
	return &plane{}
}
