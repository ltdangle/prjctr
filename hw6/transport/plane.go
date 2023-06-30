package transport

import "errors"

type plane struct {
	registration *registration
	passengers   []*passenger
}

func NewPlane(registration *registration) *plane {
	return &plane{registration: registration}
}

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

func (pl *plane) TakePassengers(p *passenger) error {
	// Check if passenger is in registration list.
	found := false
	for _, pr := range pl.registration.passengers {
		if p == pr {
			found = true
		}
	}

	if !found {
		return errors.New("passenger not found in registration")
	}

	pl.passengers = append(pl.passengers, p)

	return nil
}

func (p *plane) DropPassengers() {
	println("Plane dropped passengers")
}
