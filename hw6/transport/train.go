package transport

import (
	"errors"
	"strconv"
)

type train struct {
	maxSpeed int
	speed    int
}

func NewTrain(maxSpeed int) *train {
	return &train{maxSpeed: maxSpeed}
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

func (t *train) ChangeSpeed(speedChange int) error {
	newSpeed := t.speed + speedChange

	if newSpeed > t.maxSpeed {
		return errors.New("speed is over max: " + strconv.Itoa(t.maxSpeed))
	}

	t.speed = newSpeed
	return nil
}

func (t *train) TakePassengers(*passenger) error {
	println("Train took passengers")
	return nil
}

func (t *train) DropPassengers() {
	println("Train dropped passengers")
}
