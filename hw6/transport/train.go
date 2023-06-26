package transport

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
