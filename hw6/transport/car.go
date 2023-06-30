package transport

import (
	"errors"
	"strconv"
)

type Car struct {
	maxPassengers int
	passengers    []*Passenger
}

func NewCar(maxPassengers int) *Car {
	return &Car{maxPassengers: maxPassengers}
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

func (c *Car) TakePassengers(p *Passenger) error {
	if len(c.passengers) == c.maxPassengers {
		return errors.New("Cannot take more than " + strconv.Itoa(c.maxPassengers))
	}
	c.passengers = append(c.passengers, p)
	return nil
}

func (c *Car) DropPassengers() {
	println("Car dropped passengers")
}
