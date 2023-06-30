package transport

type train struct{}

func NewTrain() *train {
	return &train{}
}

func (t *train) Name() string {
	return "Train"
}

func (t *train) Move() {
	println("Train is moving")
}

func (t *train) Stop() {
	println("Train stopped")
}

func (t *train) ChangeSpeed() {
	println("Train changed speed")
}

func (t *train) TakePassengers(*Passenger) error {
	println("Train took passengers")
	return nil
}

func (t *train) DropPassengers() {
	println("Train dropped passengers")
}
